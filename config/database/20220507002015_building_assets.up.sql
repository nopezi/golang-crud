CREATE TABLE IF NOT EXISTS `building_assets` (
  `id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,  
  `asset_id` INT(10) UNSIGNED NOT NULL ,
  `certificate_type` VARCHAR(100) NOT NULL ,
  `certificate_number` VARCHAR(100) NOT NULL ,
  `build_year` VARCHAR(100) NOT NULL ,
  `surface_area` VARCHAR(100) NOT NULL ,
  `building_area` VARCHAR(100) NOT NULL ,
  `direction` VARCHAR(100) NOT NULL ,
  `number_of_floors` VARCHAR(100) NOT NULL ,
  `number_of_bedrooms` VARCHAR(100) NOT NULL ,
  `number_of_bathrooms` VARCHAR(100) NOT NULL ,
  `electrical_power` VARCHAR(100) NOT NULL ,
  `carport` VARCHAR(100) NOT NULL ,
  `created_at` DATETIME NOT NULL,
  `updated_at` DATETIME NOT NULL,
  PRIMARY KEY (`id`),
  CONSTRAINT name_unique UNIQUE(name)
)ENGINE = InnoDB DEFAULT CHARSET=utf8mb4;
