#### Checklist Module

- [ ] FrontWeb
    - [ ] GetAllAssetsDefault elasticsearch pagination di frontend
    - [ ] GetAllAssetsSearch elasticsearch pagination di frontend
        - params:
            - property
                - name
                - lokasi
                - tipe
                - harga
            - otomotif
                - name
                - merk
                - seri
                - harga
            - mesin
                - name
                - lokasi
                - merk
                - harga
            - lainnya
                - name
                - lokasi
                - merk
    
- [ ] Admin 
    - [ ] Assets
        - [x] GetAll
        - [x] Store 
            type => 
            - form_b1| form bangunan 1: rumah, ruko, apartemen, condotel, kos atau kontrakan, toko, vila , kios
            - form_k1| form kendaraan 1: Kapal
            - form_k2| form kendaraan 2: Pesawat, Alat Berat
            - form_k3| from kendaraan 3: Mobil, Sepeda Motor, Truk, Bus
        - [x] GetOne
        - [x] Update
        - [x] Delete
    - [ ] BuildingAssets [repoonly]
        - [x] GetAll
        - [x] Store
        - [x] GetOne
        - [x] Update
        - [x] Delete
    - [ ] VehicleAssets [repoonly]
        - [x] GetAll
        - [x] Store
        - [x] GetOne
        - [x] Update
        - [x] Delete
    - [ ] Jadwal Lelang
        - [ ] GetAll
        - [ ] Store
        - [ ] GetOne
        - [ ] Update
        - [ ] Delete
    - [ ] Categories
        - [x] GetAll
        - [x] Store
        - [x] GetOne
        - [x] Update
        - [x] Delete
    - [ ] SubCategories
        - [x] GetAll
        - [x] Store
        - [x] GetOne
        - [x] Update
        - [x] Delete
    - [x] AssetAccessPlaces ketika create asset -> create ini [repoonly]
        - [x] GetAll
        - [x] Store
        - [x] GetOne
        - [x] Update
        - [x] Delete
          - repository not yet change to *sql.DB
    - [ ] Facilities ketika create asset -> create ini [repoonly]
        - [x] GetAll
        - [x] Store
        - [x] GetOne
        - [x] Update
        - [x] Delete
    - [ ] AssetFacilities ketika create asset -> create ini [repoonly]
        - [x] GetAll
        - [x] Store
        - [x] GetOne
        - [x] Update
        - [x] Delete
    - [ ] Contacts ketika create asset -> create contact [repoonly]
        - [x] Store
        - [x] GetOne
        - [x] Update
        - [x] Delete
    - [ ] Addresses ketika create asset -> create address [repoonly]
        - [x] Store
        - [x] GetOne
        - [x] Update
        - [x] Delete
    - [ ] Images ketika create asset -> create images [repoonly]
        - [x] Store
        - [x] GetOne
        - [x] Update
        - [x] Delete
    - [ ] AssetImages ketika create asset -> create images [repoonly]
        - [x] Store
        - [x] GetOne
        - [x] Update
        - [x] Delete
    - [ ] Approvals ketika create asset -> create approvals [repoonly]
        - [x] Store
        - [x] GetOne
        - [x] Update
        - [x] Delete
    - [ ] AssetApprovals ketika create asset -> create approvals [repoonly]
        - [x] Store
        - [x] GetOne
        - [x] Update
        - [x] Delete
    - [ ] Kpknl
        - [x] GetAll
        - [x] Store
        - [x] GetOne
        - [x] Update
        - [x] Delete
        - repository not yet change to *sql.DB
    - [ ] Faqs
        - [x] GetAll
        - [x] Store
        - [x] GetOne
        - [x] Update
        - [x] Delete
          - repository not yet change to *sql.DB
    - [ ] brands
        - [ ] GetAll
        - [ ] Store
        - [ ] GetOne
        - [ ] Update
        - [ ] Delete
    - [ ] sub_brands
        - [ ] GetAll
        - [ ] Store
        - [ ] GetOne
        - [ ] Update
        - [ ] Delete
    - [ ] type_of_certificate
        - [ ] GetAll
        - [ ] Store
        - [ ] GetOne
        - [ ] Update
        - [ ] Delete
    - [ ] machine_capacity
        - [ ] GetAll
        - [ ] Store
        - [ ] GetOne
        - [ ] Update
        - [ ] Delete
    - [ ] vehicle_colors
        - [ ] GetAll
        - [ ] Store
        - [ ] GetOne
        - [ ] Update
        - [ ] Delete
    - [ ] transmissions
        - [ ] GetAll
        - [ ] Store
        - [ ] GetOne
        - [ ] Update
        - [ ] Delete
    - [ ] loginUserByPN ldap mas tomi
    - [ ] create asset multiple form
    - [ ] handle upload witbh minio
    - [ ] scheduler buat update asset status to unpublish dan hapus asset di elastic

    
- Kpknl list
- category list
- sub category list where category_id
- postalcode like postalcode
- certificate_type
- facilities
- access_place
- checker list
- signer list

# Elastic trial 14Day
https://infolelang.es.us-central1.gcp.cloud.es.io
username,
elastic,
password 
CIEkxlGYVzr83ql1Lvwp4C1N

docker exec df2152ab9116 /usr/bin/mysqldump -u root --password=P@ssw0rd infolelang2 > infolelang2.sql 

-- ==========================================================
## PR Backend
-- auth onegate => done ,( incognito not done)
-- getMCS => done 
    -- - checker list
    -- - signer list

  - handle get table error, 
  - get data empty, belum semua
  - pagination done
  - api inquiry cif ke esb = done

   - create vehicle
        - certificate_type to certificate_type_id
        - series_id
        - brand_id
        - transmission_id
        - machineCapacity_id to machine_capacity_id
        - color_id
        - get one vehicle join details above
   - api get menu by role
   - create to elastic
   - create api auction_schedule/ jadwal lelang,
   - create get all asset to elastic
   - create get all with filter to elastic
   - api crud maintain banners
   - activity logger
   - runtime logger
# next development
   - module user integration register with google or email,
   - create whitelist asset
   - mobile infolelang
   - notification created asset to approvals
   - integration geotaging google maps
   - get all status, name and desc
 ==========================================================


 ### ORM
 - Tipe Update
 ```
 -- only include field to update
 func (asset AssetRepository) UpdateApproval(request *models.AssetsUpdateApproval, include []string) (responses bool, err error) {
	return true, asset.db.DB.Select(include).Updates(&request).Error

}

exclude := []string{
 	"last_maker_id",
 		"last_maker_desc",
 		"last_maker_date",
 		"published",
 		"publish_date",
 		"expired_date",
 }
-- Exclude field
func (asset AssetRepository) UpdateApproval(request *models.AssetsUpdateApproval,exclude []string) (responses bool, err error) {
     	return true, asset.db.DB.Omit(strings.Join(exclude, ","))Updates(&request).Error
 }

func (asset AssetRepository) UpdateApproval(request *models.AssetsUpdateApproval) (responses bool, err error) {
     	return true, asset.db.DB.Save(&request).Error
 }
 ``` 