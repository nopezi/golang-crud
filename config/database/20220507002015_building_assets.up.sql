CREATE TABLE IF NOT EXISTS `building_assets` (
  `id` INT(18) UNSIGNED NOT NULL AUTO_INCREMENT, 
  `asset_id` INT(10), 
  `certificate_type` VARCHAR(180) NOT NULL ,
  `certificate_number` VARCHAR(180) NOT NULL ,
  `build_year` int(18) NULL ,
  `surface_area` int(18)  NULL ,
  `building_area`int(18)  NULL ,
  `direction` VARCHAR(180) NOT NULL ,
  `number_of_floors` int(18) NULL ,
  `number_of_bedrooms` int(18) NULL ,
  `number_of_bathrooms` int(18) NULL ,
  `electrical_power` int(18) NULL ,
  `carport` int(18) NULL ,
  `created_at` DATETIME NULL,
  `updated_at` DATETIME NULL,
  PRIMARY KEY (`id`)  
)ENGINE = InnoDB DEFAULT CHARSET=utf8mb4;
