package asset

import (
	"fmt"
	"infolelang/lib"
	access "infolelang/models/access_places"
	address "infolelang/models/addresses"
	models "infolelang/models/assets"

	// approval "infolelang/models/approvals"
	building "infolelang/models/building_assets"
	contact "infolelang/models/contacts"
	facility "infolelang/models/facilities"
	image "infolelang/models/images"
	vehicle "infolelang/models/vehicle_assets"
	"log"
	"math"
	"strings"
	"time"

	// "github.com/elastic/go-elasticsearch/v8"
	elastic "gitlab.com/golang-package-library/elasticsearch"
	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type AssetDefinition interface {
	WithTrx(trxHandle *gorm.DB) AssetRepository
	GetAll() (responses []models.AssetsResponse, err error)
	GetAssetElastic(request models.AssetRequestElastic) (responses []models.AssetsResponseGetOne, err error)
	GetAuctionSchedule(request models.AuctionSchedule) (responses []models.AuctionScheduleResponse, totalRows int64, totalData int64, err error)
	GetOne(id int64) (responses models.AssetsResponse, err error)
	GetOneAsset(id int64) (responses models.AssetsResponse, err error)
	Store(request *models.Assets) (responses *models.Assets, err error)
	StoreElastic(request models.AssetsResponseGetOne) (response bool, err error)
	DeleteElastic(request models.AssetsResponseGetOne) (response bool, err error)
	GetApproval(request models.AssetsRequestMaintain) (responses []models.AssetsResponseMaintain, totalRows int, totalData int, err error)
	GetMaintain(request models.AssetsRequestMaintain) (responses []models.AssetsResponseMaintain, totalRows int, totalData int, err error)
	UpdateApproval(request *models.AssetsUpdateApproval, include []string) (responses bool, err error)
	UpdatePublish(request *models.AssetsUpdatePublish, include []string) (responses bool, err error)
	UpdateMaintain(request *models.Assets, include []string) (responses *models.Assets, err error)
	Delete(request *models.AssetsUpdateDelete, include []string) (responses bool, err error)
	UpdateDocumentID(request *models.AssetsRequestUpdateElastic, include []string) (responses bool, err error)
	UpdateRemoveDocumentID(request *models.AssetsRequestUpdateElastic, include []string) (responses bool, err error)
	DeleteAssetImage(id int64) (err error)
}
type AssetRepository struct {
	db            lib.Database
	dbRaw         lib.Databases
	elastic       elastic.Elasticsearch
	elasticsearch lib.Elasticsearch
	logger        logger.Logger
	timeout       time.Duration
}

func NewAssetReporitory(
	db lib.Database,
	dbRaw lib.Databases,
	elastic elastic.Elasticsearch,
	elasticsearch lib.Elasticsearch,
	logger logger.Logger) AssetDefinition {
	return AssetRepository{
		db:            db,
		dbRaw:         dbRaw,
		elastic:       elastic,
		elasticsearch: elasticsearch,
		logger:        logger,
		timeout:       time.Second * 100,
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

// GetAll implements AssetDefinition
func (asset AssetRepository) GetAssetElastic(request models.AssetRequestElastic) (responses []models.AssetsResponseGetOne, err error) {
	result, err := asset.elasticsearch.Search(lib.RequestElastic{
		Index: "assets",
		Body:  request,
	})
	if err != nil {
		asset.logger.Zap.Error(err)
		return responses, err
	}
	// fmt.Println("elastic", result)
	if result == nil {
		return responses, err
	}

	for _, hit := range result.([]interface{}) {
		// log.Printf(" * ID=%s, %s", hit.(map[string]interface{})["_id"], hit.(map[string]interface{})["_source"])
		log.Println(strings.Repeat("=>", 37))

		source := hit.(map[string]interface{})["_source"]
		// fmt.Println(source.(map[string]interface{})["addresses"].(map[string]interface{})["id"])

		addresses := source.(map[string]interface{})["addresses"]
		buildingAssets := source.(map[string]interface{})["building_assets"]
		vehicleAssets := source.(map[string]interface{})["vehicle_assets"]
		facilities := source.(map[string]interface{})["facilities"]      // array
		accessPlaces := source.(map[string]interface{})["access_places"] // array
		contacts := source.(map[string]interface{})["contacts"]
		images := source.(map[string]interface{})["images"] // array
		fmt.Println("================================")
		fmt.Println("================================")

		Addresses := address.AddressesResponse{
			ID:           int64(addresses.(map[string]interface{})["id"].(float64)),
			AssetID:      int64(addresses.(map[string]interface{})["asset_id"].(float64)),
			Address:      addresses.(map[string]interface{})["address"].(string),
			Longitude:    addresses.(map[string]interface{})["longitude"].(string),
			Langitude:    addresses.(map[string]interface{})["langitude"].(string),
			Longlat:      addresses.(map[string]interface{})["longlat"].(string),
			PostalcodeID: int64(addresses.(map[string]interface{})["postalcode_id"].(float64)),
			PostalCode:   addresses.(map[string]interface{})["postal_code"].(string),
			Region:       addresses.(map[string]interface{})["region"].(string),
			District:     addresses.(map[string]interface{})["district"].(string),
			City:         addresses.(map[string]interface{})["city"].(string),
			Province:     addresses.(map[string]interface{})["province"].(string),
		}
		// fmt.Println(Addresses)

		BuildingAssets := building.BuildingAssetsResponse{
			AssetID:           int64(buildingAssets.(map[string]interface{})["asset_id"].(float64)),
			CertificateTypeID: int64(buildingAssets.(map[string]interface{})["certificate_type_id"].(float64)),
			CertificateNumber: buildingAssets.(map[string]interface{})["certificate_number"].(string),
			BuildYear:         int64(buildingAssets.(map[string]interface{})["build_year"].(float64)),
			SurfaceArea:       int64(buildingAssets.(map[string]interface{})["burface_area"].(float64)),
			BuildingArea:      int64(buildingAssets.(map[string]interface{})["building_area"].(float64)),
			Direction:         buildingAssets.(map[string]interface{})["direction"].(string),
			NumberOfFloors:    int64(buildingAssets.(map[string]interface{})["number_of_floors"].(float64)),
			NumberOfBedrooms:  int64(buildingAssets.(map[string]interface{})["number_of_bedrooms"].(float64)),
			NumberOfBathrooms: int64(buildingAssets.(map[string]interface{})["number_of_bathrooms"].(float64)),
			ElectricalPower:   int64(buildingAssets.(map[string]interface{})["electrical_power"].(float64)),
			Carport:           int64(buildingAssets.(map[string]interface{})["carport"].(float64)),
		}
		// fmt.Println(BuildingAssets)

		VehicleAssets := vehicle.VehicleAssetsResponse{
			ID:                  int64(vehicleAssets.(map[string]interface{})["id"].(float64)),
			AssetID:             int64(vehicleAssets.(map[string]interface{})["asset_id"].(float64)),
			VehicleTypeID:       int64(vehicleAssets.(map[string]interface{})["vehicle_type_id"].(float64)),
			CertificateTypeID:   int64(vehicleAssets.(map[string]interface{})["certificate_type_id"].(float64)),
			CertificateNumber:   vehicleAssets.(map[string]interface{})["certificate_number"].(string),
			Series:              vehicleAssets.(map[string]interface{})["series"].(string),
			BrandID:             int64(vehicleAssets.(map[string]interface{})["brand_id"].(float64)),
			Type:                vehicleAssets.(map[string]interface{})["type"].(string),
			ProductionYear:      vehicleAssets.(map[string]interface{})["production_year"].(string),
			TransmissionID:      int64(vehicleAssets.(map[string]interface{})["transmission_id"].(float64)),
			MachineCapacityID:   int64(vehicleAssets.(map[string]interface{})["machine_capacity_id"].(float64)),
			ColorID:             int64(vehicleAssets.(map[string]interface{})["color_id"].(float64)),
			NumberOfSeat:        int64(vehicleAssets.(map[string]interface{})["number_of_seat"].(float64)),
			NumberOfUsage:       vehicleAssets.(map[string]interface{})["number_of_usage"].(string),
			MachineNumber:       vehicleAssets.(map[string]interface{})["machine_number"].(string),
			BodyNumber:          vehicleAssets.(map[string]interface{})["body_number"].(string),
			LicenceDate:         vehicleAssets.(map[string]interface{})["licence_date"].(string),
			BrandName:           vehicleAssets.(map[string]interface{})["brand_name"].(string),
			TransmissionName:    vehicleAssets.(map[string]interface{})["transmission_name"].(string),
			MachineCapacityName: vehicleAssets.(map[string]interface{})["machine_capacity_name"].(string),
			ColorName:           vehicleAssets.(map[string]interface{})["color_name"].(string),
		}
		// fmt.Println(VehicleAssets)

		Facilities := []facility.FacilitiesResponse{}
		for _, row := range facilities.([]interface{}) {
			// rowFacility := row.(map[string]interface{})["facilities"]
			Facilities = append(Facilities, facility.FacilitiesResponse{
				ID:          int64(row.(map[string]interface{})["id"].(float64)),
				Name:        row.(map[string]interface{})["name"].(string),
				Icon:        row.(map[string]interface{})["icon"].(string),
				Status:      row.(map[string]interface{})["status"].(bool),
				Description: row.(map[string]interface{})["description"].(string),
			})
		}
		// fmt.Println(Facilities)

		AccessPlaces := []access.AccessPlacesResponse{}
		for _, row := range accessPlaces.([]interface{}) {
			AccessPlaces = append(AccessPlaces, access.AccessPlacesResponse{
				ID:          int64(row.(map[string]interface{})["id"].(float64)),
				Name:        row.(map[string]interface{})["name"].(string),
				Icon:        row.(map[string]interface{})["icon"].(string),
				Status:      row.(map[string]interface{})["status"].(bool),
				Description: row.(map[string]interface{})["description"].(string),
			})
		}
		// fmt.Println(AccessPlaces)

		Contacts := contact.ContactsResponse{
			ID:          int64(contacts.(map[string]interface{})["id"].(float64)),
			AssetID:     int64(contacts.(map[string]interface{})["asset_id"].(float64)),
			DebiturName: contacts.(map[string]interface{})["debitur_name"].(string),
			PicName:     contacts.(map[string]interface{})["pic_name"].(string),
			PicPhone:    contacts.(map[string]interface{})["pic_phone"].(string),
			PicEmail:    contacts.(map[string]interface{})["pic_email"].(string),
			Cif:         contacts.(map[string]interface{})["cif"].(string),
		}

		// fmt.Println(Contacts)

		Images := []image.ImagesResponses{}
		for _, row := range images.([]interface{}) {
			Images = append(Images, image.ImagesResponses{
				ID:        int64(row.(map[string]interface{})["id"].(float64)),
				Filename:  row.(map[string]interface{})["filename"].(string),
				Path:      row.(map[string]interface{})["path"].(string),
				Extension: row.(map[string]interface{})["extension"].(string),
				Size:      int64(row.(map[string]interface{})["size"].(float64)),
			})
		}

		// fmt.Println(Images)

		responses = append(responses, models.AssetsResponseGetOne{
			ID:              int64(source.(map[string]interface{})["id"].(float64)),
			FormType:        source.(map[string]interface{})["form_type"].(string),
			Type:            source.(map[string]interface{})["type"].(string),
			KpknlID:         int64(source.(map[string]interface{})["kpknl_id"].(float64)),
			AuctionDate:     source.(map[string]interface{})["auction_date"].(*string),
			AuctionTime:     source.(map[string]interface{})["auction_time"].(string),
			AuctionLink:     source.(map[string]interface{})["auction_link"].(string),
			CategoryID:      int64(source.(map[string]interface{})["category_id"].(float64)),
			SubCategoryID:   int64(source.(map[string]interface{})["sub_category_id"].(float64)),
			Name:            source.(map[string]interface{})["name"].(string),
			Price:           float32(source.(map[string]interface{})["price"].(float64)),
			Description:     source.(map[string]interface{})["description"].(string),
			Status:          source.(map[string]interface{})["status"].(string),
			KpknlName:       source.(map[string]interface{})["kpknl_name"].(string),
			CategoryName:    source.(map[string]interface{})["category_name"].(string),
			SubCategoryName: source.(map[string]interface{})["sub_category_name"].(string),
			StatusName:      source.(map[string]interface{})["status_name"].(string),
			Addresses:       Addresses,
			BuildingAssets:  BuildingAssets,
			VehicleAssets:   VehicleAssets,
			Facilities:      Facilities,
			AccessPlaces:    AccessPlaces,
			Contacts:        Contacts,
			Images:          Images,
		})

	}

	if err != nil {
		asset.logger.Zap.Error(err)
		return responses, err
	}
	return responses, err
}

// GetOne implements AssetDefinition
func (asset AssetRepository) GetOne(id int64) (responses models.AssetsResponse, err error) {
	err = asset.db.DB.Raw(`
	SELECT 
		ast.id,
		ast.form_type,
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
		rs.namaStatus status_name,
		ast.document_id
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

// GetOne implements AssetDefinition
func (asset AssetRepository) GetAuctionSchedule(request models.AuctionSchedule) (responses []models.AuctionScheduleResponse, totalRows int64, totalData int64, err error) {
	where := " WHERE 1+1 "
	whereCount := " 1+1 "
	if request.Name != "" {
		where += " AND a.name LIKE '%" + request.Name + "%'"
		whereCount += " AND name LIKE '%" + request.Name + "%'"
	}

	if request.KpknlID != 0 {
		where += " AND a.kpknl_id = " + request.AuctionDate
		whereCount += " AND kpknl_id = " + request.AuctionDate
	}

	if request.AuctionDate != "" {
		where += " AND MONTH(a.auction_date) = " + request.AuctionDate
		whereCount += " AND MONTH(auction_date) = " + request.AuctionDate
	}

	query := `
	SELECT a.id, a.name, 
	a.auction_date, 
	a.auction_time, 
	a.kpknl_id,
	rk.desc kpknl_name,
	c.pic_name pic_lelang,
	a2.address from assets a
	LEFT JOIN ref_kpknl rk on a.kpknl_id = rk.id
	LEFT JOIN contacts c on a.id = c.asset_id
	LEFT JOIN addresses a2 on a.id = a2.asset_id ` + where

	rows, err := asset.db.DB.Raw(query).Rows()
	defer rows.Close()

	var auctionScheduleResponse models.AuctionScheduleResponse
	for rows.Next() {
		asset.db.DB.ScanRows(rows, &auctionScheduleResponse)
		responses = append(responses, auctionScheduleResponse)
	}

	if err != nil {
		asset.logger.Zap.Error(err)
		return responses, totalRows, totalData, err
	}

	asset.db.DB.Table("assets").Where(whereCount).Count(&totalData)
	result := float64(totalData) / float64(request.Limit)
	totalRows = int64(math.Ceil(result))

	return responses, totalRows, totalData, err
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
func (asset AssetRepository) UpdateApproval(request *models.AssetsUpdateApproval, include []string) (responses bool, err error) {
	return true, asset.db.DB.Select(include).Updates(&request).Error

}

// UpdatePublish implements AssetDefinition
func (asset AssetRepository) UpdatePublish(request *models.AssetsUpdatePublish, include []string) (responses bool, err error) {
	return true, asset.db.DB.Save(&request).Error
}

// Update implements AssetDefinition
func (asset AssetRepository) UpdateMaintain(request *models.Assets, include []string) (responses *models.Assets, err error) {
	return request, asset.db.DB.Save(&request).Error
}

// Delete implements AssetDefinition
func (asset AssetRepository) Delete(request *models.AssetsUpdateDelete, include []string) (responses bool, err error) {
	return true, asset.db.DB.Save(&request).Error
}

func (asset AssetRepository) GetApproval(request models.AssetsRequestMaintain) (responses []models.AssetsResponseMaintain, totalRows int, totalData int, err error) {
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

	if request.Status != "" {
		where += " AND ast.status = '" + request.Status + "'"
		whereCount += " AND ast.status = '" + request.Status + "'"
	}

	if request.Published != "" {
		where += " AND ast.published = '" + request.Published + "'"
		whereCount += " AND ast.published = '" + request.Published + "'"
	}

	if request.Deleted != "" {
		where += " AND ast.deleted = '" + request.Deleted + "'"
		whereCount += " AND ast.deleted = '" + request.Deleted + "'"
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
	rows, err := asset.dbRaw.DB.Query(query, request.Limit, request.Offset)

	asset.logger.Zap.Info("rows ", rows)
	if err != nil {
		return responses, totalRows, totalData, err
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
		return responses, totalRows, totalData, err
	}

	paginateQuery := `SELECT count(*) as total from assets  ast 
					left join categories c on ast.category_id = c.id 
					left join sub_categories sc on ast.sub_category_id = sc.id
					left join ref_status rs on ast.status  = rs.kodeStatus
					left join contacts c2 on ast.id = c2.asset_id
					left join approvals a on ast.id = a.asset_id ` + whereCount
	err = asset.dbRaw.DB.QueryRow(paginateQuery).Scan(&totalRows)

	result := float64(totalRows) / float64(request.Limit)
	resultFinal := int(math.Ceil(result))

	return responses, resultFinal, totalRows, err
}

func (asset AssetRepository) GetMaintain(request models.AssetsRequestMaintain) (responses []models.AssetsResponseMaintain, totalRows int, totalData int, err error) {
	where := ""
	whereCount := ""

	if request.MakerID != "" {
		where += " AND ast.last_maker_id = '" + request.MakerID + "'"
		whereCount += " AND ast.last_maker_id = '" + request.MakerID + "'"
	}

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

	rows, err := asset.dbRaw.DB.Query(query, request.Limit, request.Offset)

	if err != nil {
		return responses, totalRows, totalData, err
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
		return responses, totalRows, totalData, err
	}

	paginateQuery := `SELECT count(*) as total from assets  ast 
						join categories c on ast.category_id = c.id 
						join sub_categories sc on ast.sub_category_id = sc.id
						join ref_status rs on ast.status  = rs.kodeStatus
						join contacts c2 on ast.id = c2.asset_id ` + whereCount

	err = asset.dbRaw.DB.QueryRow(paginateQuery).Scan(&totalRows)

	result := float64(totalRows) / float64(request.Limit)
	resultFinal := int(math.Ceil(result))

	fmt.Println("OK=>", responses, resultFinal)

	return responses, resultFinal, totalRows, err

}

func (asset AssetRepository) StoreElastic(request models.AssetsResponseGetOne) (response bool, err error) {
	documentID := lib.UUID(false)
	// fmt.Println(request)
	store, err := asset.elasticsearch.Store(lib.RequestElastic{
		DocumentID: documentID,
		Index:      "assets",
		Body:       request,
	})
	// fmt.Println(err)
	if err != nil {
		asset.logger.Zap.Error(err)
		return false, err
	}

	if !store {
		asset.logger.Zap.Error(err)
		return false, err
	}

	if store {
		update, err := asset.UpdateDocumentID(&models.AssetsRequestUpdateElastic{
			ID:         request.ID,
			DocumentID: documentID,
		},
			[]string{
				"document_id",
			})
		if !update || err != nil {
			return false, err
		}
		return true, err
	} else {
		return false, err
	}

}

func (asset AssetRepository) DeleteElastic(request models.AssetsResponseGetOne) (response bool, err error) {
	store, err := asset.elastic.Delete(elastic.RequestElastic{
		DocumentID: request.DocumentID,
		Index:      "assets",
	})

	if err != nil {
		asset.logger.Zap.Error(err)
		return false, err
	}

	if !store {
		asset.logger.Zap.Error(err)
		return false, err
	}

	if store {
		update, err := asset.UpdateDocumentID(&models.AssetsRequestUpdateElastic{
			ID:         request.ID,
			DocumentID: "",
		},
			[]string{
				"document_id",
			})
		if !update || err != nil {
			return false, err
		}
	}

	return true, err
}

// UpdateDocumentID implements AssetDefinition
func (asset AssetRepository) UpdateDocumentID(request *models.AssetsRequestUpdateElastic, include []string) (responses bool, err error) {
	return true, asset.db.DB.Save(&request).Error
}

// UpdateRemoveDocumentID implements AssetDefinition
func (asset AssetRepository) UpdateRemoveDocumentID(request *models.AssetsRequestUpdateElastic, include []string) (responses bool, err error) {
	return true, asset.db.DB.Save(&request).Error
}

// DeleteAssetImage implements ImageDefinition
func (asset AssetRepository) DeleteAssetImage(id int64) (err error) {
	return asset.db.DB.Where("id = ?", id).Delete(&models.AssetImagesRequest{}).Error
}
