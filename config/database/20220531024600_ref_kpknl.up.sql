CREATE TABLE IF NOT EXISTS `ref_kpknl` (
  `id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,  
  `kpknl` VARCHAR(100) NOT NULL ,
  `address` VARCHAR(100) NOT NULL ,
  `phone` VARCHAR(100) NOT NULL ,
  `fax` VARCHAR(100) NOT NULL ,
  `created_at` DATETIME NOT NULL,
  `updated_at` DATETIME NOT NULL,
  PRIMARY KEY (`id`)  
)ENGINE = InnoDB DEFAULT CHARSET=utf8mb4;
