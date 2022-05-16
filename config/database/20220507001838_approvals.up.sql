CREATE TABLE IF NOT EXISTS `approvals` (
  `id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,  
  `checker_id` VARCHAR(100) NOT NULL,
  `checker_desc` VARCHAR(100) NOT NULL,
  `checker_comment` VARCHAR(100) NOT NULL,
  `checker_date` DATETIME NOT NULL,
  `signer_id` VARCHAR(100) NOT NULL,
  `signer_desc` VARCHAR(100) NOT NULL,
  `signer_comment` VARCHAR(100) NOT NULL,
  `signer_date` VARCHAR(100) NOT NULL,
  `created_at` DATETIME NOT NULL,
  `updated_at` DATETIME NOT NULL,
  PRIMARY KEY (`id`)  
)ENGINE = InnoDB DEFAULT CHARSET=utf8mb4;
