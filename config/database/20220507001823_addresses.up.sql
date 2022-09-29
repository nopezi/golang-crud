CREATE TABLE IF NOT EXISTS `addresses` (
  `id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,  
  `asset_id` INT(10),
  `postalcode_id` INT(10) UNSIGNED NOT NULL ,
  `address` VARCHAR(100) NOT NULL,
  `longitude` VARCHAR(100) NOT NULL,
  `langitude` VARCHAR(100) NOT NULL,
  `created_at` DATETIME NULL,
  `updated_at` DATETIME NULL,
  PRIMARY KEY (`id`)
)ENGINE = InnoDB DEFAULT CHARSET=utf8mb4;
