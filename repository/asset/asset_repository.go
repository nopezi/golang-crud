package asset

import (
	"fmt"
	"infolelang/lib"
	models "infolelang/models/assets"
	"math"
	"time"

	elastic "gitlab.com/golang-package-library/elasticsearch"
	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type AssetDefinition interface {
	WithTrx(trxHandle *gorm.DB) AssetRepository
	GetAll() (responses []models.AssetsResponse, err error)
	GetOne(id int64) (responses models.AssetsResponse, err error)
	GetOneAsset(id int64) (responses models.AssetsResponse, err error)
	Store(request *models.Assets) (responses *models.Assets, err error)
	Update(request *models.AssetsRequest) (responses bool, err error)
	Delete(id int64) (err error)
	StoreElastic(request *models.AssetsResponse) (response bool, err error)
	GetApproval(request models.AssetsRequestMaintain) (responses []models.AssetsResponseMaintain, totalRows int, err error)
	GetMaintain(request models.AssetsRequestMaintain) (responses []models.AssetsResponseMaintain, totalRows int, err error)
}
type AssetRepository struct {
	db      lib.Database
	db2     lib.Databases
	elastic elastic.Elasticsearch
	logger  logger.Logger
	timeout time.Duration
}

func NewAssetReporitory(
	db lib.Database,
	db2 lib.Databases,
	elastic elastic.Elasticsearch,
	logger logger.Logger) AssetDefinition {
	return AssetRepository{
		db:      db,
		db2:     db2,
		elastic: elastic,
		logger:  logger,
		timeout: time.Second * 100,
	}
}

// WithTrx implements AssetDefinition
func (asset AssetRepository) WithTrx(trxHandle *gorm.DB) AssetRepository {
	if trxHandle == nil {
		asset.logger.Zap.Error("transaction Database not found in gin context. ")
		return asset
	}
	asset.db.DB = trxHandle
	return asset
}

// GetAll implements AssetDefinition
func (asset AssetRepository) GetAll() (responses []models.AssetsResponse, err error) {
	return responses, asset.db.DB.Find(&responses).Error
}

// GetOne implements AssetDefinition
func (asset AssetRepository) GetOne(id int64) (responses models.AssetsResponse, err error) {
	return responses, asset.db.DB.Where("id = ?", id).Find(&responses).Error
}

// GetOneAsset implements AssetDefinition
func (asset AssetRepository) GetOneAsset(id int64) (responses models.AssetsResponse, err error) {
	return responses, asset.db.DB.Where("asset_id = ?", id).Find(&responses).Error
}

// Store implements AssetDefinition
func (asset AssetRepository) Store(request *models.Assets) (responses *models.Assets, err error) {
	return request, asset.db.DB.Save(&request).Error
}

// Update implements AssetDefinition
func (asset AssetRepository) Update(request *models.AssetsRequest) (responses bool, err error) {
	return true, asset.db.DB.Save(&responses).Error
}

// Delete implements AssetDefinition
func (asset AssetRepository) Delete(id int64) (err error) {
	return asset.db.DB.Where("id = ?", id).Delete(&models.AssetsResponse{}).Error
}

func (asset AssetRepository) StoreElastic(request *models.AssetsResponse) (response bool, err error) {
	return true, err
}

func (asset AssetRepository) GetApproval(request models.AssetsRequestMaintain) (responses []models.AssetsResponseMaintain, totalRows int, err error) {
	where := " WHERE 1+1 "
	whereCount := " WHERE 1+1 "
	if request.CheckerID != "" {
		where += " AND a.checker_id = '" + request.CheckerID + "'"
		whereCount += " AND a.checker_id = '" + request.CheckerID + "'"
	}

	if request.SignerID != "" {
		if request.CheckerID == "" {
			where += " AND a.signer_id = '" + request.SignerID + "'"
			whereCount += " AND a.signer_id = '" + request.SignerID + "'"
		} else {
			where += " OR a.signer_id = '" + request.SignerID + "'"
			whereCount += " OR a.signer_id = '" + request.SignerID + "'"
		}
	}

	if request.Name != "" {
		where += " AND ast.name LIKE '%" + request.Name + "%'"
		whereCount += " AND ast.name LIKE '%" + request.Name + "%'"
	}

	query := `SELECT ast.id, ast.type,
			c.name category,sc.name sub_category, 
			ast.name,ast.price, ast.status, c2.pic_name, ast.published,
			a.checker_id, a.signer_id
			from assets ast 
			join categories c on ast.category_id = c.id 
			join sub_categories sc on ast.sub_category_id = sc.id
			join ref_status rs on ast.status  = rs.kodeStatus
			join contacts c2 on ast.id = c2.asset_id
			join approvals a on ast.id = a.asset_id ` + where + ` LIMIT ? OFFSET ?`
	asset.logger.Zap.Info(query)
	rows, err := asset.db2.DB.Query(query, request.Limit, request.Offset)

	asset.logger.Zap.Info("rows ", rows)
	if err != nil {
		return responses, totalRows, err
	}

	response := models.AssetsResponseMaintain{}
	for rows.Next() {
		_ = rows.Scan(
			&response.ID,
			&response.Type,
			&response.Category,
			&response.SubCategory,
			&response.Name,
			&response.Price,
			&response.Status,
			&response.PicName,
			&response.Published,
			&response.CheckerID,
			&response.SignerID,
		)
		responses = append(responses, response)
	}

	if err = rows.Err(); err != nil {
		return responses, totalRows, err
	}

	paginateQuery := `SELECT count(*) as total from assets  ast 
					join categories c on ast.category_id = c.id 
					join sub_categories sc on ast.sub_category_id = sc.id
					join ref_status rs on ast.status  = rs.kodeStatus
					join contacts c2 on ast.id = c2.asset_id
					join approvals a on ast.id = a.asset_id ` + whereCount
	// fmt.Println("paginateQuery", paginateQuery)
	err = asset.db2.DB.QueryRow(paginateQuery).Scan(&totalRows)

	result := float64(totalRows) / float64(request.Limit)
	resultFinal := int(math.Ceil(result))

	// fmt.Println("OK=>", responses, resultFinal)

	return responses, resultFinal, err
}

func (asset AssetRepository) GetMaintain(request models.AssetsRequestMaintain) (responses []models.AssetsResponseMaintain, totalRows int, err error) {
	where := ""
	whereCount := ""

	if request.Name != "" {
		where += " AND ast.name LIKE '%" + request.Name + "%'"
		whereCount += " AND ast.name LIKE '%" + request.Name + "%'"
	}

	query := `SELECT ast.id, ast.type,
			c.name category,sc.name sub_category, 
			ast.name,ast.price, ast.status, c2.pic_name, ast.published
			from assets ast 
			join categories c on ast.category_id = c.id 
			join sub_categories sc on ast.sub_category_id = sc.id
			join ref_status rs on ast.status  = rs.kodeStatus
			join contacts c2 on ast.id = c2.asset_id ` + where + ` LIMIT ? OFFSET ?`

	rows, err := asset.db2.DB.Query(query, request.Limit, request.Offset)

	if err != nil {
		return responses, totalRows, err
	}

	response := models.AssetsResponseMaintain{}
	for rows.Next() {
		_ = rows.Scan(
			&response.ID,
			&response.Type,
			&response.Category,
			&response.SubCategory,
			&response.Name,
			&response.Price,
			&response.Status,
			&response.PicName,
			&response.Published,
		)
		responses = append(responses, response)
	}

	if err = rows.Err(); err != nil {
		return responses, totalRows, err
	}

	paginateQuery := `SELECT count(*) as total from assets  ast 
						join categories c on ast.category_id = c.id 
						join sub_categories sc on ast.sub_category_id = sc.id
						join ref_status rs on ast.status  = rs.kodeStatus
						join contacts c2 on ast.id = c2.asset_id ` + whereCount

	err = asset.db2.DB.QueryRow(paginateQuery).Scan(&totalRows)

	result := float64(totalRows) / float64(request.Limit)
	resultFinal := int(math.Ceil(result))

	fmt.Println("OK=>", responses, resultFinal)

	return responses, resultFinal, err

}
