CREATE TABLE IF NOT EXISTS `ref_postal_code` (
  `id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,  
  `postal_code` VARCHAR(100) NOT NULL,
  `region` VARCHAR(100) NOT NULL,
  `district` VARCHAR(100) NOT NULL,
  `city` VARCHAR(100) NOT NULL,
  `province` VARCHAR(100) NOT NULL ,
  `enabled` int(2) NOT NULL,
  `created_at` DATETIME NOT NULL,
  `updated_at` DATETIME NOT NULL,
  PRIMARY KEY (`id`)  
)ENGINE = InnoDB DEFAULT CHARSET=utf8mb4;
