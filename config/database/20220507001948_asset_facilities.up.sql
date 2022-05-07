CREATE TABLE IF NOT EXISTS `asset_facilities` (
  `id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,  
  `asset_id` INT(10) UNSIGNED NOT NULL ,
  `facility_id` INT(10) UNSIGNED NOT NULL ,
  `created_at` DATETIME NOT NULL,
  `updated_at` DATETIME NOT NULL,
  PRIMARY KEY (`id`),
  CONSTRAINT name_unique UNIQUE(name)
)ENGINE = InnoDB DEFAULT CHARSET=utf8mb4;
