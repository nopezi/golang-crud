CREATE TABLE `access_places` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100)  DEFAULT NULL,
  `icon` varchar(100)  DEFAULT NULL,
  `description` varchar(100)  DEFAULT NULL,
  `status` tinyint(1) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) 19 ;

CREATE TABLE `addresses` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `asset_id` int DEFAULT NULL,
  `postalcode_id` int unsigned NOT NULL,
  `address` varchar(100) NOT NULL,
  `longitude` varchar(100) NOT NULL,
  `langitude` varchar(100) NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
);

CREATE TABLE `approvals` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `asset_id` int DEFAULT NULL,
  `checker_id` varchar(100) NOT NULL,
  `checker_desc` varchar(100) NOT NULL,
  `checker_comment` varchar(100)  DEFAULT NULL,
  `checker_date` datetime DEFAULT NULL,
  `signer_id` varchar(100) NOT NULL,
  `signer_desc` varchar(100) NOT NULL,
  `signer_comment` varchar(100)  DEFAULT NULL,
  `signer_date` datetime DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
);

CREATE TABLE `asset_access_places` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `asset_id` int unsigned NOT NULL,
  `access_place_id` int unsigned NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
);

CREATE TABLE `asset_approvals` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `asset_id` int unsigned NOT NULL,
  `approval_id` int unsigned NOT NULL,
  `created_at` datetime NOT NULL,
  `status` varchar(100) NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ;

CREATE TABLE `asset_facilities` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `asset_id` int unsigned NOT NULL,
  `facility_id` int unsigned NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
);

CREATE TABLE `asset_images` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `asset_id` int unsigned NOT NULL,
  `image_id` int unsigned NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ;

CREATE TABLE `assets` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `type` varchar(100) NOT NULL,
  `kpknl_id` int unsigned NOT NULL,
  `auction_date` date NOT NULL,
  `auction_time` time NOT NULL,
  `auction_link` varchar(100) NOT NULL,
  `category_id` int unsigned NOT NULL,
  `sub_category_id` int unsigned NOT NULL,
  `name` varchar(100) NOT NULL,
  `price` int NOT NULL,
  `description` varchar(100) NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `status` varchar(100) DEFAULT NULL,
  `published` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ;

CREATE TABLE `auction_schedule` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `asset_id` int unsigned DEFAULT NULL,
  `date` datetime DEFAULT NULL,
  `kpknl_id` int unsigned DEFAULT NULL,
  `kanca` varchar(100)  DEFAULT NULL,
  `contact_id` int unsigned DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ;

CREATE TABLE `building_assets` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `asset_id` int DEFAULT NULL,
  `certificate_type` varchar(180) NOT NULL,
  `certificate_number` varchar(180) NOT NULL,
  `build_year` int DEFAULT NULL,
  `surface_area` int DEFAULT NULL,
  `building_area` int DEFAULT NULL,
  `direction` varchar(180) NOT NULL,
  `number_of_floors` int DEFAULT NULL,
  `number_of_bedrooms` int DEFAULT NULL,
  `number_of_bathrooms` int DEFAULT NULL,
  `electrical_power` int DEFAULT NULL,
  `carport` int DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
);

CREATE TABLE `categories` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `status` tinyint(1) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ;

CREATE TABLE `certificate_type` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100)  DEFAULT NULL,
  `icon` varchar(100)  DEFAULT NULL,
  `description` varchar(100)  DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ;

CREATE TABLE `cities` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `province_id` int unsigned NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ;

CREATE TABLE `contacts` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `asset_id` int DEFAULT NULL,
  `debitur_name` varchar(100) NOT NULL,
  `pic_name` varchar(100) NOT NULL,
  `pic_phone` varchar(100) NOT NULL,
  `pic_email` varchar(100) NOT NULL,
  `cif` varchar(100) NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ;

CREATE TABLE `districts` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `province_id` int unsigned NOT NULL,
  `city_id` int unsigned NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ;

CREATE TABLE `facilities` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100)  DEFAULT NULL,
  `icon` varchar(100)  DEFAULT NULL,
  `description` varchar(100)  DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ;

CREATE TABLE `faqs` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `question` varchar(100) NOT NULL,
  `answer` varchar(100) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ;

CREATE TABLE `images` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `filename` varchar(100)  DEFAULT NULL,
  `path` varchar(100) NOT NULL,
  `extension` varchar(100)  DEFAULT NULL,
  `size` varchar(100) NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ;

CREATE TABLE `map_approver` (
  `tipeUker` varchar(5)  NOT NULL,
  `hilfm` varchar(3)  NOT NULL,
  `checker` int NOT NULL DEFAULT '1',
  `signer` int NOT NULL DEFAULT '1',
  PRIMARY KEY (`tipeUker`,`hilfm`) USING BTREE
) ;

CREATE TABLE `map_maker` (
  `tipeUker` varchar(5)  NOT NULL,
  `levelId` varchar(8)  NOT NULL,
  PRIMARY KEY (`tipeUker`,`levelId`) USING BTREE
) ;

CREATE TABLE `map_special_level_id` (
  `kostl` varchar(10)  NOT NULL,
  `orgeh` varchar(8)  NOT NULL,
  `hilfm` varchar(3)  NOT NULL,
  `levelId` varchar(8)  NOT NULL,
  PRIMARY KEY (`kostl`,`orgeh`,`hilfm`,`levelId`) USING BTREE
) ;

CREATE TABLE `mst_access_menu` (
  `LevelUker` varchar(3)  NOT NULL,
  `LevelID` varchar(8)  NOT NULL,
  `IDMenu` int NOT NULL,
  PRIMARY KEY (`LevelUker`,`LevelID`,`IDMenu`) USING BTREE
) ;

CREATE TABLE `mst_menu` (
  `IDMenu` int NOT NULL AUTO_INCREMENT,
  `Title` varchar(500)  DEFAULT NULL,
  `Url` varchar(200)  DEFAULT NULL,
  `Deskripsi` varchar(8000)  DEFAULT NULL,
  `Icon` varchar(100)  NOT NULL,
  `Atribut` varchar(4000)  DEFAULT NULL,
  `Badge` int NOT NULL DEFAULT '0' COMMENT '0 = ya, 1 = tidak',
  `IDParent` bigint NOT NULL,
  `Target` varchar(100)  DEFAULT NULL,
  `Urutan` int NOT NULL,
  `RoleAccess` int NOT NULL DEFAULT '0' COMMENT '0 = Tanpa role akses/public, 1 = Dengan role akses',
  `KanpusOnly` int DEFAULT NULL COMMENT '0 = Tidak, 1 = Ya',
  `Jenis` int DEFAULT '0' COMMENT '0 = Umum\r\n1 = Pemimpin Uker',
  `Posisi` int NOT NULL COMMENT '0 = Atas Kiri, 1 = Atas Kanan, 2 = Sidebar Kiri',
  `Status` int NOT NULL DEFAULT '0' COMMENT '0 = Non aktif, 1 = Aktif',
  PRIMARY KEY (`IDMenu`) USING BTREE
) ;

CREATE TABLE `provinces` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ;

CREATE TABLE `ref_cronjobs` (
  `time` varchar(100)  DEFAULT NULL,
  `method` varchar(100)  DEFAULT NULL,
  `action` varchar(10)  DEFAULT NULL COMMENT 'add, remove',
  `flag` varchar(10)  DEFAULT NULL COMMENT '1 : tambah job ke service, any: tidak di eksekusi',
  `status` varchar(10)  DEFAULT NULL,
  `updated_at` varchar(255)  DEFAULT NULL
) ;

CREATE TABLE `ref_kpknl` (
  `id` int NOT NULL,
  `desc` varchar(255) DEFAULT NULL,
  `uodated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) DEFAULT CHARSET=latin1;

CREATE TABLE `ref_postal_code` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `postal_code` varchar(100) NOT NULL,
  `region` varchar(100) NOT NULL,
  `district` varchar(100) NOT NULL,
  `city` varchar(100) NOT NULL,
  `province` varchar(100) NOT NULL,
  `enabled` int NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ;

CREATE TABLE `ref_status` (
  `kodeStatus` varchar(5)  NOT NULL,
  `namaStatus` varchar(255)  DEFAULT NULL,
  `namaStatusAlias` varchar(255)  DEFAULT NULL,
  `textClassStatus` varchar(255)  DEFAULT NULL,
  `textClassStatusAlias` varchar(255)  DEFAULT NULL,
  `group` varchar(255)  DEFAULT NULL,
  PRIMARY KEY (`kodeStatus`) USING BTREE,
  KEY `kode_status` (`kodeStatus`) USING BTREE
) ;

CREATE TABLE `schema_migrations` (
  `version` bigint NOT NULL,
  `dirty` tinyint(1) NOT NULL,
  PRIMARY KEY (`version`)
) ;

CREATE TABLE `sub_categories` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `category_id` int unsigned NOT NULL,
  `name` varchar(100) NOT NULL,
  `form` varchar(100) DEFAULT NULL,
  `status` tinyint(1) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ;

CREATE TABLE `users` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `pernr` varchar(100)  NOT NULL,
  `password` varchar(100) NOT NULL,
  `name` varchar(20)  DEFAULT NULL,
  `status` tinyint(1) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `email_unique` (`pernr`)
) ;

CREATE TABLE `vehicle_assets` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `asset_id` int unsigned NOT NULL,
  `vehicle_type` varchar(100) NOT NULL,
  `certificate_type_id` int NOT NULL,
  `certificate_number` varchar(100) NOT NULL,
  `series` varchar(100) NOT NULL,
  `brand_id` int NOT NULL,
  `type` varchar(100) NOT NULL,
  `production_year` varchar(100) NOT NULL,
  `transmission_id` int NOT NULL,
  `machine_capacity_id` int NOT NULL,
  `color_id` int NOT NULL,
  `number_of_seat` varchar(100) NOT NULL,
  `number_of_usage` varchar(100) NOT NULL,
  `machine_number` varchar(100) NOT NULL,
  `body_number` varchar(100) NOT NULL,
  `licence_date` datetime NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ;

CREATE TABLE `vehicle_brand` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `vehicle_category_id` int DEFAULT NULL,
  `name` varchar(100)  DEFAULT NULL,
  `icon` varchar(100)  DEFAULT NULL,
  `description` varchar(100)  DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ;

CREATE TABLE `vehicle_capacity` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100)  DEFAULT NULL,
  `icon` varchar(100)  DEFAULT NULL,
  `description` varchar(100)  DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ;

CREATE TABLE `vehicle_category` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100)  DEFAULT NULL,
  `icon` varchar(100)  DEFAULT NULL,
  `description` varchar(100)  DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ;

CREATE TABLE `vehicle_color` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100)  DEFAULT NULL,
  `icon` varchar(100)  DEFAULT NULL,
  `description` varchar(100)  DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) 8 ;

CREATE TABLE `vehicle_transmission` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100)  DEFAULT NULL,
  `icon` varchar(100)  DEFAULT NULL,
  `description` varchar(100)  DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ;
