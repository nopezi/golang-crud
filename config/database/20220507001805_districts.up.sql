CREATE TABLE IF NOT EXISTS `districts` (
  `id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,  
  `name` VARCHAR(100) NOT NULL,
  `province_id` INT(10) UNSIGNED NOT NULL ,
  `city_id` INT(10) UNSIGNED NOT NULL ,
  `created_at` DATETIME NULL,
  `updated_at` DATETIME NULL,
  PRIMARY KEY (`id`)
)ENGINE = InnoDB DEFAULT CHARSET=utf8mb4;
