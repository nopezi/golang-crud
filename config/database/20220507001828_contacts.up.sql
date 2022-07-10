CREATE TABLE IF NOT EXISTS `contacts` (
  `id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,  
  `asset_id` INT(10),
  `debitur_name` VARCHAR(100) NOT NULL,
  `pic_name` VARCHAR(100) NOT NULL,
  `pic_phone` VARCHAR(100) NOT NULL,
  `pic_email` VARCHAR(100) NOT NULL,
  `cif` VARCHAR(100) NOT NULL,
  `created_at` DATETIME NULL,
  `updated_at` DATETIME NULL,
  PRIMARY KEY (`id`)
)ENGINE = InnoDB DEFAULT CHARSET=utf8mb4;
