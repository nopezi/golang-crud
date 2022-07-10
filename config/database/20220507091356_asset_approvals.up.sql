CREATE TABLE IF NOT EXISTS `asset_approvals` (
  `id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,  
  `asset_id` INT(10) UNSIGNED NOT NULL ,
  `approval_id` INT(10) UNSIGNED NOT NULL ,
  `created_at` DATETIME NULL,
  `status` VARCHAR(100) NOT NULL,
  `updated_at` DATETIME NULL,
  PRIMARY KEY (`id`)  
)ENGINE = InnoDB DEFAULT CHARSET=utf8mb4;
