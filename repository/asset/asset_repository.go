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
	StoreElastic(request *models.AssetsResponse) (response bool, err error)
	GetApproval(request models.AssetsRequestMaintain) (responses []models.AssetsResponseMaintain, totalRows int, err error)
	GetMaintain(request models.AssetsRequestMaintain) (responses []models.AssetsResponseMaintain, totalRows int, err error)
	UpdateApproval(request *models.AssetsUpdateApproval) (responses bool, err error)
	UpdatePublish(request *models.AssetsUpdatePublish) (responses bool, err error)
	UpdateMaintain(request *models.AssetsRequest) (responses bool, err error)
	Delete(request *models.AssetsUpdateDelete) (responses bool, err error)
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
	// db.Raw("SELECT * FROM users WHERE name1 = @name OR name2 = @name2 OR name3 = @name",
	//    sql.Named("name", "jinzhu1"), sql.Named("name2", "jinzhu2")).Find(&user)
	// return responses, asset.db.DB.Where("id = ?", id).Find(&responses).Error
	fmt.Println("===>ID ", id)
	err = asset.db.DB.Raw(`
	SELECT 
		ast.id,
		ast.type,
		ast.kpknl_id,
		ast.auction_date,
		ast.auction_time,
		ast.auction_link,
		ast.category_id,
		ast.sub_category_id,
		ast.name,
		ast.price,
		ast.description,
		ast.maker_id,
		ast.maker_desc,
		ast.maker_date,
		ast.last_maker_id,
		ast.last_maker_desc,
		ast.last_maker_date,
		ast.published,
		ast.deleted,
		ast.publish_date,
		ast.expired_date,
		ast.status,
		ast.action,
		ast.updated_at,
		ast.created_at,
		rk.desc  kpknl_name,
		c.name category_name,
		sc.name sub_category_name,
		rs.namaStatus status_name
		FROM assets ast 
		LEFT JOIN categories c on ast.category_id = c.id 
		LEFT JOIN sub_categories sc on ast.sub_category_id = sc.id
		LEFT JOIN ref_status rs on ast.status  = rs.kodeStatus
		LEFT JOIN ref_kpknl rk  on ast.kpknl_id  = rk.id where ast.id = ?`, id).Find(&responses).Error

	if err != nil {
		asset.logger.Zap.Error(err)
		return responses, err
	}
	return responses, err
}

// GetOneAsset implements AssetDefinition
func (asset AssetRepository) GetOneAsset(id int64) (responses models.AssetsResponse, err error) {
	return responses, asset.db.DB.Where("asset_id = ?", id).Find(&responses).Error
}

// Store implements AssetDefinition
func (asset AssetRepository) Store(request *models.Assets) (responses *models.Assets, err error) {
	return request, asset.db.DB.Save(&request).Error
}

// UpdateApproval implements AssetDefinition
func (asset AssetRepository) UpdateApproval(request *models.AssetsUpdateApproval) (responses bool, err error) {
	return true, asset.db.DB.Save(&request).Error
}

// UpdatePublish implements AssetDefinition
func (asset AssetRepository) UpdatePublish(request *models.AssetsUpdatePublish) (responses bool, err error) {
	return true, asset.db.DB.Save(&request).Error
}

// Update implements AssetDefinition
func (asset AssetRepository) UpdateMaintain(request *models.AssetsRequest) (responses bool, err error) {
	return true, asset.db.DB.Save(&request).Error
}

// Delete implements AssetDefinition
func (asset AssetRepository) Delete(request *models.AssetsUpdateDelete) (responses bool, err error) {
	return true, asset.db.DB.Save(&request).Error
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
			left join categories c on ast.category_id = c.id 
			left join sub_categories sc on ast.sub_category_id = sc.id
			left join ref_status rs on ast.status  = rs.kodeStatus
			left join contacts c2 on ast.id = c2.asset_id
			left join approvals a on ast.id = a.asset_id ` + where + ` order by id desc LIMIT ? OFFSET ?`
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
					left join categories c on ast.category_id = c.id 
					left join sub_categories sc on ast.sub_category_id = sc.id
					left join ref_status rs on ast.status  = rs.kodeStatus
					left join contacts c2 on ast.id = c2.asset_id
					left join approvals a on ast.id = a.asset_id ` + whereCount
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
			join contacts c2 on ast.id = c2.asset_id ` + where + ` order by id desc LIMIT ? OFFSET ?`

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
