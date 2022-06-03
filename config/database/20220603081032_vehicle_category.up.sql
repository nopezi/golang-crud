  CREATE TABLE `vehicle_category` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100)  DEFAULT NULL,
  `icon` varchar(100)  DEFAULT NULL,
  `description` varchar(100)  DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1;
