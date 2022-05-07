CREATE TABLE IF NOT EXISTS `contacts` (
  `id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,  
  `debitur_name` VARCHAR(100) NOT NULL,
  `pic_name` VARCHAR(100) NOT NULL,
  `pic_phone` VARCHAR(100) NOT NULL,
  `pic_email` VARCHAR(100) NOT NULL,
  `cif` VARCHAR(100) NOT NULL,
  `created_at` DATETIME NOT NULL,
  `updated_at` DATETIME NOT NULL,
  PRIMARY KEY (`id`),
  CONSTRAINT name_unique UNIQUE(name)
)ENGINE = InnoDB DEFAULT CHARSET=utf8mb4;
