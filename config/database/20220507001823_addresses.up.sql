CREATE TABLE IF NOT EXISTS `addresses` (
  `id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,  
  `province_id` INT(10) UNSIGNED NOT NULL ,
  `city_id` INT(10) UNSIGNED NOT NULL ,
  `district_id` INT(10) UNSIGNED NOT NULL ,
  `address` VARCHAR(100) NOT NULL,
  `longitude` VARCHAR(100) NOT NULL,
  `langitude` VARCHAR(100) NOT NULL,
  `created_at` DATETIME NOT NULL,
  `updated_at` DATETIME NOT NULL,
  PRIMARY KEY (`id`),
  CONSTRAINT name_unique UNIQUE(name)
)ENGINE = InnoDB DEFAULT CHARSET=utf8mb4;
