-- --------------------------------------------------------
-- Host:                         127.0.0.1
-- Server version:               8.0.29 - MySQL Community Server - GPL
-- Server OS:                    Win64
-- HeidiSQL Version:             12.0.0.6468
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

-- Dumping structure for table contents
CREATE TABLE IF NOT EXISTS `contents` (
  `id` int NOT NULL AUTO_INCREMENT,
  `id_user` int DEFAULT NULL,
  `title` varchar(200) CHARACTER SET armscii8 COLLATE armscii8_bin DEFAULT NULL,
  `description` text CHARACTER SET armscii8 COLLATE armscii8_bin,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=armscii8 COLLATE=armscii8_bin;

-- Dumping data for table contents: ~2 rows (approximately)
INSERT INTO `contents` (`id`, `id_user`, `title`, `description`, `created_at`, `updated_at`) VALUES
	(4, 3, 'judul', 'isi deskripsi', NULL, NULL),
	(5, 3, 'artikel', 'isi artikel', '2022-09-18 21:23:47', '2022-09-18 21:23:47'),
	(7, NULL, 'judul pertama', 'ini adalah judul', '2022-09-29 21:31:30', NULL),
	(8, NULL, 'judul pertama', 'ini adalah judul', '2022-09-29 21:32:55', NULL),
	(10, NULL, 'judul pertama update', 'ini adalah judul update', '0000-00-00 00:00:00', NULL);

-- Dumping structure for table users
CREATE TABLE IF NOT EXISTS `users` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(200) COLLATE armscii8_bin DEFAULT NULL,
  `email` varchar(200) COLLATE armscii8_bin DEFAULT NULL,
  `password` varchar(200) COLLATE armscii8_bin DEFAULT NULL,
  `birthday` date DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `level_user` int DEFAULT NULL COMMENT '1 = admin, 2 = user',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=armscii8 COLLATE=armscii8_bin;

-- Dumping data for table users: ~2 rows (approximately)
INSERT INTO `users` (`id`, `name`, `email`, `password`, `birthday`, `created_at`, `updated_at`, `level_user`) VALUES
	(1, 'admin', 'admin@gmail.com', '$2a$12$cPcfzi2UDyQEk6FqewS.buc9KOzOGHGcQXF1gDaJIjcg7yZsrE/Aq', NULL, '2022-09-15 22:53:25', '2022-09-15 22:53:25', 1),
	(3, 'jonocoba', 'jonocoba@gmail.com', '123', NULL, '2022-09-18 18:58:41', '2022-09-18 18:58:41', 2);

/*!40103 SET TIME_ZONE=IFNULL(@OLD_TIME_ZONE, 'system') */;
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;
