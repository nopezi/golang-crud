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
    - [ ] handle upload with minio
    - [ ] scheduler buat update asset status to unpublish dan hapus asset di elastic

    
    



  CREATE TABLE `certificate_type` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `icon` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `description` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `vehicle_transmission` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `icon` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `description` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `vehicle_capacity` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `icon` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `description` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `vehicle_color` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `icon` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `description` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `vehicle_brand` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `icon` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `description` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `vehicle_category` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `icon` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `description` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `auction_schedule` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `asset_id` int unsigned DEFAULT NULL,
  `date` datetime DEFAULT NULL,
  `kpknl_id` int unsigned DEFAULT NULL,
  `kanca` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `contact_id` int unsigned DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

