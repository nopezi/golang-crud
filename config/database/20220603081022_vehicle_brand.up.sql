CREATE TABLE `vehicle_brand` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `vehicle_category_id` int DEFAULT NULL,
  `name` varchar(100) DEFAULT NULL,
  `icon` varchar(100)  DEFAULT NULL,
  `description` varchar(100) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1;
