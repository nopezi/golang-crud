CREATE TABLE IF NOT EXISTS `facilities` (
  `id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,  
  `name` VARCHAR(100) NOT NULL,
  `icon` VARCHAR(100) NOT NULL,
  `description` VARCHAR(100) NOT NULL,
  `created_at` DATETIME NOT NULL,
  `updated_at` DATETIME NOT NULL,
  PRIMARY KEY (`id`)
)ENGINE = InnoDB DEFAULT CHARSET=utf8mb4;