CREATE TABLE IF NOT EXISTS `approvals` (
  `id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,  
  `asset_id` INT(10),
  `checker_id` VARCHAR(100) NOT NULL,
  `checker_desc` VARCHAR(100) NOT NULL,
  `checker_comment` VARCHAR(100) NULL,
  `checker_date` DATETIME  NULL,
  `signer_id` VARCHAR(100) NOT NULL,
  `signer_desc` VARCHAR(100) NOT NULL,
  `signer_comment` VARCHAR(100) NULL,
  `signer_date` DATETIME NULL,
  `created_at` DATETIME NULL,
  `updated_at` DATETIME NULL,
  PRIMARY KEY (`id`)  
)ENGINE = InnoDB DEFAULT CHARSET=utf8mb4;
