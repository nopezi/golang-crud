CREATE TABLE IF NOT EXISTS `assets` (
  `id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,  
  `type` VARCHAR(100) NOT NULL,
  `kpknl_id` INT(10) UNSIGNED NOT NULL ,
  `auction_date` DATE NOT NULL,
  `auction_time` TIME NOT NULL,
  `auction_link` VARCHAR(100) NOT NULL,
  `category_id` INT(10) UNSIGNED NOT NULL ,
  `sub_category_id` INT(10) UNSIGNED NOT NULL ,
  `name` VARCHAR(100) NOT NULL,
  `price` int(100) NOT NULL,
  `description` VARCHAR(100) NOT NULL,
  `created_at` DATETIME NULL,
  `updated_at` DATETIME NULL,
  PRIMARY KEY (`id`)
)ENGINE = InnoDB DEFAULT CHARSET=utf8mb4;
