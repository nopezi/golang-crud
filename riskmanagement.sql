-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Sep 22, 2022 at 03:52 AM
-- Server version: 10.4.24-MariaDB
-- PHP Version: 8.1.6

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `riskmanagement`
--

-- --------------------------------------------------------

--
-- Table structure for table `activity`
--

CREATE TABLE `activity` (
  `id` int(10) UNSIGNED NOT NULL,
  `kode_activity` varchar(10) NOT NULL,
  `name` varchar(100) NOT NULL,
  `create_at` datetime DEFAULT NULL,
  `update_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `activity`
--

INSERT INTO `activity` (`id`, `kode_activity`, `name`, `create_at`, `update_at`) VALUES
(1, '01', 'MANAJEMEN RISIKO', '2022-08-15 07:53:04', NULL),
(2, '02', 'PERKREDITAN', '2022-08-15 07:59:22', NULL),
(3, '03', 'PENDANAAN', '2022-08-15 08:03:01', NULL),
(4, '04', 'JASA', '2022-08-15 08:04:52', NULL),
(5, '05', 'TRADE FINANCE', '2022-08-15 08:06:38', NULL),
(6, '06', 'OPERASIONAL', '2022-08-15 08:14:13', NULL),
(7, '07', 'PELAYANAN', '2022-08-15 08:14:22', NULL),
(8, '08', 'SUPPORT', '2022-08-15 08:14:36', NULL),
(9, '09', 'STRATEGIS', '2022-08-15 08:14:49', NULL),
(10, '10', 'TREASURY', '2022-08-15 08:15:18', NULL),
(11, '11', 'IT', '2022-08-15 08:15:34', NULL),
(12, '12', 'Invesment Service', '2022-08-15 08:15:46', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `briefing`
--

CREATE TABLE `briefing` (
  `id` int(10) UNSIGNED NOT NULL,
  `no_pelaporan` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `unit_kerja` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `peserta` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `jumlah_peserta` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `maker_id` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `maker_desc` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `maker_date` datetime DEFAULT NULL,
  `last_maker_id` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `last_maker_desc` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `last_maker_date` datetime DEFAULT NULL,
  `status` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `action` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `deleted` tinyint(1) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `briefing`
--

INSERT INTO `briefing` (`id`, `no_pelaporan`, `unit_kerja`, `peserta`, `jumlah_peserta`, `maker_id`, `maker_desc`, `maker_date`, `last_maker_id`, `last_maker_desc`, `last_maker_date`, `status`, `action`, `deleted`, `created_at`, `updated_at`) VALUES
(1, 'BR-50046567-100522-0001', 'Kantor Cabang Khusus', 'Semua', '27', '00304155', '00304155 | Feb\'hana Faradilla Bimantari | Teller', '2022-08-24 12:06:36', '00304155', '00304155 | Feb\'hana Faradilla Bimantari | Teller', '2022-08-24 12:06:36', '02b', 'updateDelete', 1, '2022-08-24 12:06:36', '2022-08-25 08:25:23'),
(2, 'BR-50046567-250822-0001', 'Kantor Cabang Khusus', 'Semua', '27', '00304155', '00304155 | Feb\'hana Faradilla Bimantari | Teller', '2022-08-24 12:06:36', '00304155', '00304155 | NANA FEBRI | Teller', '2022-08-25 15:35:36', '02b', 'Selesai', 0, '2022-08-25 08:25:23', '2022-08-25 15:35:36'),
(5, 'BR-50046567-100522-0003', 'Kantor Cabang Khusus', 'Semua', '27', '00304155', '00304155 | Feb\'hana Faradilla Bimantari | Teller', '2022-08-25 15:35:36', '00304155', '00304155 | NANA FEBRI | Teller', '2022-08-25 15:35:36', '02b', 'Update', 0, '2022-08-25 15:35:36', '2022-08-25 15:35:36'),
(7, 'BR-50046567-100522-0004', 'Kantor Cabang Medan', 'Semua', '27', '00304155', '00304155 | Feb\'hana Faradilla Bimantari | Teller', '2022-09-05 11:00:18', '00304155', '00304155 | Feb\'hana Faradilla Bimantari | Teller', '2022-09-05 11:00:18', '02b', 'updateDelete', 1, '2022-09-05 11:00:18', '2022-09-05 11:00:18'),
(8, 'BR-00304155-100522-0001', 'Kantor Cabang Khusus', 'Semua', '27', '00304155', '00304155 | Feb\'hana Faradilla Bimantari | Teller', '2022-09-19 09:28:32', '00304155', '00304155 | Feb\'hana Faradilla Bimantari | Teller', '2022-09-19 09:28:32', '01a', 'Draft', 0, '2022-09-19 09:28:32', NULL),
(9, 'BR-00304155-200922-0004', 'Kantor Cabang Khusus', 'Semua', '27', '00304155', '00304155 | Feb\'hana Faradilla Bimantari | Teller', '2022-09-20 09:09:45', '00304155', '00304155 | Feb\'hana Faradilla Bimantari | Teller', '2022-09-20 09:09:45', '01a', 'Draft', 0, '2022-09-20 09:09:45', NULL),
(12, 'BR-00304155-200922-0005', 'jhasdgjh', 'hasgdjasdg', '0', '00304155', '00304155 | Feb\'hana Faradilla Bimantari | Teller', '2022-09-20 16:12:08', '00304155', '00304155 | Feb\'hana Faradilla Bimantari | Teller', '2022-09-20 16:12:08', '01a', 'Draft', 0, '2022-09-20 16:12:08', NULL),
(13, 'BR-00304155-200922-0005', 'asdasd', 'asdasd', '0', '00304155', '00304155 | Feb\'hana Faradilla Bimantari | Teller', '2022-09-20 16:12:08', '00304155', '00304155 | Feb\'hana Faradilla Bimantari | Teller', '2022-09-20 16:12:08', '01a', 'Draft', 0, '2022-09-20 16:12:08', NULL),
(14, 'BR-00304155-210922-0006', 'Kantor Cabang Khusus', '1,2', '2', '00304155', '00304155 | Feb\'hana Faradilla Bimantari | Teller', '2022-09-21 08:33:30', '00304155', '00304155 | Feb\'hana Faradilla Bimantari | Teller', '2022-09-21 08:33:30', '01a', 'Draft', 0, '2022-09-21 08:33:30', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `briefing_materis`
--

CREATE TABLE `briefing_materis` (
  `id` int(10) UNSIGNED NOT NULL,
  `briefing_id` int(10) NOT NULL,
  `activity_id` int(10) UNSIGNED NOT NULL,
  `sub_activity_id` int(10) UNSIGNED NOT NULL,
  `product_id` int(10) UNSIGNED NOT NULL,
  `judul_materi` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `rekomendasi_materi` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `materi_tambahan` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `briefing_materis`
--

INSERT INTO `briefing_materis` (`id`, `briefing_id`, `activity_id`, `sub_activity_id`, `product_id`, `judul_materi`, `rekomendasi_materi`, `materi_tambahan`, `created_at`, `updated_at`) VALUES
(2, 1, 1, 1, 1, 'PDP 1 - RISIKO STRATEGIS', 'Risk Awarness', 'Materi yang disampaikan pada pertemuaan briefing', '2022-08-24 12:06:36', NULL),
(3, 2, 1, 1, 1, 'PDP 1 - RISIKO STRATEGIS', 'Risk Awarness', 'Materi yang disampaikan pada pertemuaan briefing', NULL, '2022-08-25 15:35:36'),
(4, 2, 1, 1, 1, 'PDP 1 - RISIKO STRATEGIS', 'Risk Awarness', 'Materi yang disampaikan pada pertemuaan briefing', NULL, '2022-08-25 15:35:36'),
(7, 5, 1, 1, 1, 'PDP 1 - RISIKO STRATEGIS', 'Risk Awarness | JUKLAK', 'Materi yang disampaikan pada pertemuaan briefing', NULL, '2022-08-25 15:35:36'),
(8, 5, 1, 1, 1, 'PDP 1 - RISIKO STRATEGIS', 'JUKLAK', 'Materi yang disampaikan pada pertemuaan briefing', NULL, '2022-08-25 15:35:36'),
(9, 5, 1, 1, 1, 'PDP 1 - RISIKO STRATEGIS', 'Juklak | Risk Awarness', 'Materi yang disampaikan pada pertemuaan briefing', '2022-08-25 15:35:36', NULL),
(12, 7, 1, 1, 1, 'PDP 1 - RISIKO STRATEGIS', 'Risk Awarness', 'Materi yang disampaikan pada pertemuaan briefing', NULL, '2022-09-05 11:00:18'),
(14, 7, 2, 1, 2, 'PDP 1 - RISIKO STRATEGIS', 'Risk Awarness', 'Materi yang disampaikan pada pertemuaan briefing', NULL, '2022-09-05 11:00:18'),
(15, 7, 2, 1, 2, 'PDP 1 - RISIKO STRATEGIS', 'Risk Awarness', 'Materi yang disampaikan pada pertemuaan briefing', NULL, '2022-09-05 11:00:18'),
(16, 8, 1, 1, 1, 'Penegasan Protokol Penangan COVID-19 bagi Pekerja BRI', '1,2,3', 'Materi yang disampaikan pada pertemuaan briefing', '2022-09-19 09:28:32', NULL),
(17, 9, 1, 1, 1, 'Penegasan Protokol Penangan COVID-19 bagi Pekerja BRI', '1,2,3', 'Materi yang disampaikan pada pertemuaan briefing', '2022-09-20 09:09:45', NULL),
(19, 12, 0, 0, 0, '', '', '', '2022-09-20 16:12:08', NULL),
(20, 13, 1, 1, 1, 'asdasd', 'asdasd', 'asdasda', '2022-09-20 16:12:08', NULL),
(21, 14, 1, 1, 1, 'Penegasan Protokol Penangan COVID-19 bagi Pekerja BRI', '1,2', 'Materi yang disampaikan pada pertemuaan briefing', '2022-09-21 08:33:30', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `coaching`
--

CREATE TABLE `coaching` (
  `id` int(10) UNSIGNED NOT NULL,
  `no_pelaporan` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `unit_kerja` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `peserta` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `jumlah_peserta` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `activity_id` int(10) UNSIGNED NOT NULL,
  `sub_activity_id` int(10) UNSIGNED NOT NULL,
  `maker_id` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `maker_desc` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `maker_date` datetime DEFAULT NULL,
  `last_maker_id` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `last_maker_desc` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `last_maker_date` datetime DEFAULT NULL,
  `status` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `action` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `deleted` tinyint(1) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `coaching`
--

INSERT INTO `coaching` (`id`, `no_pelaporan`, `unit_kerja`, `peserta`, `jumlah_peserta`, `activity_id`, `sub_activity_id`, `maker_id`, `maker_desc`, `maker_date`, `last_maker_id`, `last_maker_desc`, `last_maker_date`, `status`, `action`, `deleted`, `created_at`, `updated_at`) VALUES
(1, 'CO-50046567-100522-0001', 'Kantor Cabang Khusus', 'Semua', '27', 1, 1, '00304155', '00304155 | Feb\'hana Faradilla Bimantari | Teller', '2022-08-30 10:40:21', '00304155', '00304155 | Feb\'hana Faradilla Bimantari | Teller', '2022-08-30 13:47:23', '02b', 'updateDelete', 1, '0000-00-00 00:00:00', '2022-08-30 13:47:23'),
(2, 'CO-50046567-100522-0001', 'Kantor Cabang Khusus', 'Semua', '27', 1, 1, '00304155', '00304155 | Feb\'hana Faradilla Bimantari | Teller', '2022-08-30 13:17:42', '00304155', '00304155 | NANA FEBRI | Teller', '2022-08-30 14:27:36', '02b', 'Update', 0, '2022-08-30 13:17:42', '2022-08-30 14:27:36'),
(3, 'BR-50046567-100522-0003', 'Kantor Cabang Medan Utara', 'Semua', '27', 1, 1, '00304155', '00304155 | Feb\'hana Faradilla Bimantari | Teller', '2022-09-05 11:00:18', '00304155', '00304155 | Feb\'hana Faradilla Bimantari | Teller', '2022-09-05 11:00:18', '02b', 'updateDelete', 1, '2022-09-05 11:00:18', '2022-09-05 11:00:18'),
(4, 'BR-50046567-100522-0003', 'Kantor Cabang Medan', 'Semua', '27', 1, 1, '00304155', '00304155 | Feb\'hana Faradilla Bimantari | Teller', '2022-09-05 11:00:18', '00304155', '00304155 | Feb\'hana Faradilla Bimantari | Teller', NULL, '01a', 'Draft', 0, '2022-09-05 11:00:18', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `coaching_activity`
--

CREATE TABLE `coaching_activity` (
  `id` int(10) UNSIGNED NOT NULL,
  `coaching_id` int(10) NOT NULL,
  `risk_issue_id` int(10) UNSIGNED NOT NULL,
  `judul_materi` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `rekomendasi_materi` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `materi_tambahan` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `coaching_activity`
--

INSERT INTO `coaching_activity` (`id`, `coaching_id`, `risk_issue_id`, `judul_materi`, `rekomendasi_materi`, `materi_tambahan`, `created_at`, `updated_at`) VALUES
(2, 1, 1, 'PDP 2 - RISIKO STRATEGIS', 'Risk Awarness', 'Materi yang disampaikan pada pertemuaan briefing', '2022-08-30 10:40:21', NULL),
(3, 2, 1, 'RISIKO STRATEGIS', 'Risk Awarness | Juklak', 'Materi yang disampaikan pada pertemuaan Coaching', '2022-08-30 10:40:21', '2022-08-30 14:27:36'),
(4, 2, 1, 'RISIKO STRATEGIS', 'JUKLAK', 'Materi yang disampaikan pada pertemuaan Coaching', '2022-08-30 10:40:21', '2022-08-30 14:27:36'),
(5, 2, 1, 'RISIKO STRATEGIS 3', 'JUKLAK JUNGKLIK', 'Materi yang disampaikan pada pertemuaan Coaching', '2022-08-30 10:40:21', '2022-08-30 14:27:36'),
(6, 3, 1, 'PDP 1 - RISIKO STRATEGIS', 'Risk Awarness', 'Materi yang disampaikan pada pertemuaan briefing', NULL, '2022-09-05 11:00:18'),
(7, 3, 1, 'PDP 2 - RISIKO STRATEGIS', 'Risk Awarness', 'Materi yang disampaikan pada pertemuaan briefing', NULL, '2022-09-05 11:00:18'),
(8, 4, 1, 'PDP 1 - RISIKO STRATEGIS', 'Risk Awarness', 'Materi yang disampaikan pada pertemuaan briefing', '2022-09-05 11:00:18', NULL),
(9, 4, 1, 'PDP 2 - RISIKO STRATEGIS', 'Risk Awarness', 'Materi yang disampaikan pada pertemuaan briefing', '2022-09-05 11:00:18', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `files`
--

CREATE TABLE `files` (
  `id` int(10) UNSIGNED NOT NULL,
  `filename` varchar(100) NOT NULL,
  `path` varchar(100) NOT NULL,
  `extension` varchar(100) NOT NULL,
  `size` varchar(100) NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `files`
--

INSERT INTO `files` (`id`, `filename`, `path`, `extension`, `size`, `created_at`, `updated_at`) VALUES
(1, 'Risk Awarness.txt', 'verifikasi/2022/9/21/Risk Awarness.txt', 'plain/text', '50', NULL, '2022-09-21 14:34:12'),
(2, 'Risk Awarness.txt', 'materi/2022/9/20/Risk Awarness.txt', 'text document', '6748', '2022-09-20 09:09:45', NULL),
(3, 'Risk Awarness.txt', 'verifikasi/2022/9/21/Risk Awarness.txt', 'plain/text', '6748', '2022-09-21 14:04:40', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `incident_cause`
--

CREATE TABLE `incident_cause` (
  `id` int(10) UNSIGNED NOT NULL,
  `kode_kejadian` varchar(100) NOT NULL,
  `penyebab_kejadian1` varchar(100) NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `incident_cause`
--

INSERT INTO `incident_cause` (`id`, `kode_kejadian`, `penyebab_kejadian1`, `created_at`, `updated_at`) VALUES
(1, 'PK1.MOP.0001', 'Sumber Daya Manusia', '2022-08-16 11:41:02', NULL),
(2, 'PK1.MOP.0002', 'Proses Bisnis', '2022-08-16 11:41:36', NULL),
(3, 'PK1.MOP.0003', 'Sistem IT atau Teknologi & Sistem', '2022-08-16 11:42:16', NULL),
(4, 'PK1.MOP.0004', 'Gangguan Eksternal', '2022-08-16 11:42:36', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `materi`
--

CREATE TABLE `materi` (
  `id` int(10) UNSIGNED NOT NULL,
  `name` varchar(100) NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `materi`
--

INSERT INTO `materi` (`id`, `name`, `created_at`, `updated_at`) VALUES
(1, 'JUKLAK', '2022-08-31 09:22:13', NULL),
(2, 'Risk Awarness', '2022-09-20 09:09:45', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `materi_files`
--

CREATE TABLE `materi_files` (
  `id` int(10) UNSIGNED NOT NULL,
  `materi_id` int(10) NOT NULL,
  `files_id` int(10) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `materi_files`
--

INSERT INTO `materi_files` (`id`, `materi_id`, `files_id`) VALUES
(1, 1, 1),
(2, 2, 4);

-- --------------------------------------------------------

--
-- Table structure for table `mst_access_menu`
--

CREATE TABLE `mst_access_menu` (
  `LevelUker` varchar(10) COLLATE utf8mb4_unicode_ci NOT NULL,
  `LevelID` varchar(8) COLLATE utf8mb4_unicode_ci NOT NULL,
  `IDMenu` int(11) NOT NULL,
  `Keterangan` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `mst_access_menu`
--

INSERT INTO `mst_access_menu` (`LevelUker`, `LevelID`, `IDMenu`, `Keterangan`) VALUES
('KC', '014', 1, 'Approver KC  Pinca'),
('KC', '014', 6, 'Approver KC  Pinca'),
('KC', '041', 1, 'Maker KC RM SME'),
('KC', '041', 2, 'Maker KC RM SME'),
('KC', '041', 3, 'Maker KC RM SME'),
('KC', '041', 5, 'Maker KC RM SME'),
('KC', '147', 1, 'Maker KC RM CRR'),
('KC', '147', 2, 'Maker KC RM CRR'),
('KC', '147', 3, 'Maker KC RM CRR'),
('KC', '147', 5, 'Maker KC RM CRR'),
('KC', '159', 1, 'Approver KC Asistant Manager'),
('KC', '159', 6, 'Approver KC Asistant Manager'),
('KC', '160', 1, 'Approver KC Manager'),
('KC', '160', 6, 'Approver KC Manager'),
('KCK', '041', 1, 'Maker KCK RM SME'),
('KCK', '041', 2, 'Maker KCK RM SME'),
('KCK', '041', 3, 'Maker KCK RM SME'),
('KCK', '041', 5, 'Maker KCK RM SME'),
('KCK', '147', 1, 'Maker KCK RM CRR'),
('KCK', '147', 2, 'Maker KCK RM CRR'),
('KCK', '147', 3, 'Maker KCK RM CRR'),
('KCK', '147', 5, 'Maker KCK RM CRR'),
('KCK', '157', 1, 'Maker KCK Assistant'),
('KCK', '157', 2, 'Maker KCK Assistant'),
('KCK', '157', 3, 'Maker KCK Assistant'),
('KCK', '157', 5, 'Maker KCK Assistant'),
('KCK', '158', 1, 'Maker KCK Officer'),
('KCK', '158', 2, 'Maker KCK Officer'),
('KCK', '158', 3, 'Maker KCK Officer'),
('KCK', '158', 5, 'Maker KCK Officer'),
('KCK', '159', 1, 'Maker KCK Assistant Manager'),
('KCK', '159', 2, 'Maker KCK Assistant Manager'),
('KCK', '159', 3, 'Maker KCK Assistant Manager'),
('KCK', '159', 5, 'Maker KCK Assistant Manager'),
('KCK', '160', 1, 'Maker KCK Manager'),
('KCK', '160', 2, 'Maker KCK Manager'),
('KCK', '160', 3, 'Maker KCK Manager'),
('KCK', '160', 5, 'Maker KCK Manager'),
('KCK', '160', 6, 'Approver Maker KCK Manager'),
('KCK', '161', 1, 'Approver KCK Senior Manager'),
('KCK', '161', 6, 'Approver KCK Senior Manager'),
('KCK', '162', 1, 'Approver KCK AVP'),
('KCK', '162', 6, 'Approver KCK AVP'),
('KCK', '163', 1, 'Approver KCK VP'),
('KCK', '163', 6, 'Approver KCK VP'),
('KCK', '164', 1, 'Approver KCK EVP'),
('KCK', '164', 6, 'Approver KCK EVP'),
('KCK', '166', 1, 'Approver KCK SVP'),
('KCK', '166', 6, 'Approver KCK SVP'),
('KW', '041', 1, 'Maker KW RM SME'),
('KW', '041', 2, 'Maker KW RM SME'),
('KW', '041', 3, 'Maker KW RM SME'),
('KW', '041', 5, 'Maker KW RM SME'),
('KW', '147', 1, 'Maker KW RM CRR'),
('KW', '147', 2, 'Maker KW RM CRR'),
('KW', '147', 3, 'Maker KW RM CRR'),
('KW', '147', 5, 'Maker KW RM CRR'),
('KW', '157', 1, 'Maker KW Assistant'),
('KW', '157', 2, 'Maker KW Assistant'),
('KW', '157', 3, 'Maker KW Assistant'),
('KW', '157', 5, 'Maker KW Assistant'),
('KW', '158', 1, 'Maker KW Officer'),
('KW', '158', 2, 'Maker KW Officer'),
('KW', '158', 3, 'Maker KW Officer'),
('KW', '158', 5, 'Maker KW Officer'),
('KW', '159', 1, 'Maker KW Assistant Manager'),
('KW', '159', 2, 'Maker KW Assistant Manager'),
('KW', '159', 3, 'Maker KW Assistant Manager'),
('KW', '159', 5, 'Maker KW Assistant Manager'),
('KW', '160', 1, 'Maker KW Manager'),
('KW', '160', 2, 'Maker KW Manager'),
('KW', '160', 3, 'Maker KW Manager'),
('KW', '160', 5, 'Maker KW Manager'),
('KW', '171', 1, 'Approver KW Dep Head'),
('KW', '171', 6, 'Approver KW Dep Head'),
('KW', '173', 1, 'Approver KW Dep Head'),
('KW', '173', 6, 'Approver KW Dep Head'),
('PS80000', '157', 1, 'Maker KP CRR Assistant'),
('PS80000', '157', 2, 'Maker KP CRR Assistant'),
('PS80000', '157', 3, 'Maker KP CRR Assistant'),
('PS80000', '157', 4, 'Maker KP CRR Assistant'),
('PS80000', '157', 5, 'Maker KP CRR Assistant'),
('PS80000', '157', 7, 'Maker KP CRR Assistant'),
('PS80000', '158', 1, 'Maker KP CRR Officer'),
('PS80000', '158', 2, 'Maker KP CRR Officer'),
('PS80000', '158', 3, 'Maker KP CRR Officer'),
('PS80000', '158', 4, 'Maker KP CRR Officer'),
('PS80000', '158', 5, 'Maker KP CRR Officer'),
('PS80000', '158', 7, 'Maker KP CRR Officer'),
('PS80000', '159', 1, 'Maker KP CRR Assist Manager'),
('PS80000', '159', 2, 'Maker KP CRR Assist Manager'),
('PS80000', '159', 3, 'Maker KP CRR Assist Manager'),
('PS80000', '159', 4, 'Maker KP CRR Assist Manager'),
('PS80000', '159', 5, 'Maker KP CRR Assist Manager'),
('PS80000', '159', 7, 'Maker KP CRR Assist Manager'),
('PS80000', '160', 1, 'Maker Approver  KP CRR'),
('PS80000', '160', 2, 'Maker Approver  KP CRR'),
('PS80000', '160', 3, 'Maker Approver  KP CRR'),
('PS80000', '160', 4, 'Maker Approver  KP CRR'),
('PS80000', '160', 5, 'Maker Approver  KP CRR'),
('PS80000', '160', 6, 'Maker Approver  KP CRR'),
('PS80000', '160', 7, 'Maker Approver  KP CRR'),
('PS80000', '161', 1, 'Approver KP CRR SM'),
('PS80000', '161', 6, 'Approver KP CRR SM'),
('PS80000', '162', 1, 'Approver KP CRR AVP'),
('PS80000', '162', 6, 'Approver KP CRR AVP'),
('PS80000', '163', 1, 'Approver KP CRR VP'),
('PS80000', '163', 6, 'Approver KP CRR VP'),
('PS80000', '164', 1, 'Approver KP CRR EVP'),
('PS80000', '164', 6, 'Approver KP CRR EVP'),
('PS80000', '166', 1, 'Approver KP CRR SVP'),
('PS80000', '166', 6, 'Approver KP CRR SVP'),
('PS98400', '157', 1, 'Maker KP Assistant'),
('PS98400', '157', 2, 'Maker KP Assistant'),
('PS98400', '157', 3, 'Maker KP Assistant'),
('PS98400', '157', 4, 'Maker KP Assistant'),
('PS98400', '157', 5, 'Maker KP Assistant'),
('PS98400', '157', 7, 'Maker KP Assistant'),
('PS98400', '158', 1, 'Maker KP Officer'),
('PS98400', '158', 2, 'Maker KP Officer'),
('PS98400', '158', 3, 'Maker KP Officer'),
('PS98400', '158', 4, 'Maker KP Officer'),
('PS98400', '158', 5, 'Maker KP Officer'),
('PS98400', '158', 7, 'Maker KP Officer'),
('PS98400', '159', 1, 'Maker KP Assist Manager'),
('PS98400', '159', 2, 'Maker KP Assist Manager'),
('PS98400', '159', 3, 'Maker KP Assist Manager'),
('PS98400', '159', 4, 'Maker KP Assist Manager'),
('PS98400', '159', 5, 'Maker KP Assist Manager'),
('PS98400', '159', 7, 'Maker KP Assist Manager'),
('PS98400', '160', 1, 'Maker Approver  KP'),
('PS98400', '160', 2, 'Maker Approver  KP'),
('PS98400', '160', 3, 'Maker Approver  KP'),
('PS98400', '160', 4, 'Maker Approver  KP'),
('PS98400', '160', 5, 'Maker Approver  KP'),
('PS98400', '160', 6, 'Maker Approver  KP'),
('PS98400', '160', 7, 'Maker Approver  KP'),
('PS98400', '161', 1, 'Approver KP'),
('PS98400', '161', 6, 'Approver KP'),
('PS98400', '162', 1, 'Approver KP'),
('PS98400', '162', 6, 'Approver KP'),
('PS98400', '163', 1, 'Approver KP'),
('PS98400', '163', 6, 'Approver KP'),
('PS98400', '164', 1, 'Approver KP'),
('PS98400', '164', 6, 'Approver KP'),
('PS98400', '166', 1, 'Approver KP'),
('PS98400', '166', 6, 'Approver KP');

-- --------------------------------------------------------

--
-- Table structure for table `mst_menu`
--

CREATE TABLE `mst_menu` (
  `IDMenu` int(11) NOT NULL,
  `Title` varchar(500) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `Url` varchar(200) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `Deskripsi` varchar(8000) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `Icon` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `svgIcon` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `fontIcon` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `Atribut` varchar(4000) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `Badge` int(11) NOT NULL DEFAULT 0 COMMENT '0 = ya, 1 = tidak',
  `IDParent` bigint(20) NOT NULL,
  `Target` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `Urutan` int(11) NOT NULL,
  `RoleAccess` int(11) NOT NULL DEFAULT 0 COMMENT '0 = Tanpa role akses/public, 1 = Dengan role akses',
  `KanpusOnly` int(11) DEFAULT NULL COMMENT '0 = Tidak, 1 = Ya',
  `Jenis` int(11) DEFAULT 0 COMMENT '0 = Umum\r\n1 = Pemimpin Uker',
  `Posisi` int(11) NOT NULL COMMENT '0 = Atas Kiri, 1 = Atas Kanan, 2 = Sidebar Kiri',
  `Status` int(11) NOT NULL DEFAULT 0 COMMENT '0 = Non aktif, 1 = Aktif'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `mst_menu`
--

INSERT INTO `mst_menu` (`IDMenu`, `Title`, `Url`, `Deskripsi`, `Icon`, `svgIcon`, `fontIcon`, `Atribut`, `Badge`, `IDParent`, `Target`, `Urutan`, `RoleAccess`, `KanpusOnly`, `Jenis`, `Posisi`, `Status`) VALUES
(1, 'Beranda', '/beranda', 'Beranda', 'fas fa-home', 'media/icons/duotune/art/art002.svg', 'bi-person', NULL, 0, 0, NULL, 1, 1, 0, 0, 1, 1),
(2, 'Kelola Aset', '', 'Kelola Aset', 'mdi-clipboard-list-outline', NULL, NULL, NULL, 0, 0, NULL, 2, 1, 0, 0, 1, 1),
(3, 'Tambah Aset', '/aset/upload-aset', 'Tambah Aset', 'fas fa-file-upload', 'media/icons/duotune/general/gen022.svg', 'bi-archive', NULL, 0, 2, NULL, 3, 1, 0, 0, 2, 1),
(4, 'Ubah Data Aset', '/aset/maintain-aset', 'Ubah Data Aset admin', 'fas fa-file-signature', 'media/icons/duotune/general/gen022.svg', 'bi-archive', NULL, 0, 2, NULL, 3, 1, 0, 0, 3, 1),
(5, 'Ubah Data Aset', '/aset/maintain-aset', 'Ubah Data Aset maker', 'fas fa-file-signature', 'media/icons/duotune/general/gen022.svg', 'bi-archive', NULL, 0, 2, NULL, 3, 1, 0, 0, 3, 1),
(6, 'Persetujuan Data Aset', '/aset/approval-aset', 'Persetujuan Data Aset', 'fas fa-check-square', 'media/icons/duotune/general/gen022.svg', 'bi-archive', NULL, 0, 2, NULL, 1, 1, 0, 0, 4, 1),
(7, 'Pengaturan', '', 'Pengaturan', '', NULL, NULL, NULL, 0, 0, NULL, 2, 1, 0, 0, 1, 1),
(8, 'Banner', '/pengaturan/maintain-banner', 'Persetujuan Data Aset', 'fas fa-images', 'media/icons/duotune/general/gen022.svg', 'bi-archive', NULL, 0, 5, NULL, 3, 1, 0, 0, 2, 1);

-- --------------------------------------------------------

--
-- Table structure for table `product`
--

CREATE TABLE `product` (
  `id` int(10) UNSIGNED NOT NULL,
  `kode_product` varchar(20) NOT NULL,
  `product` varchar(100) NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `product`
--

INSERT INTO `product` (`id`, `kode_product`, `product`, `created_at`, `updated_at`) VALUES
(1, 'PROD.1', 'Operasional Kas', '2022-09-14 11:38:36', NULL),
(2, 'PROD.2', 'Operasional ATM dan CRM', '2022-09-14 11:38:36', NULL),
(4, 'PROD.3', 'AKUNTANSI & PELAPORAN', '2022-09-15 10:19:25', '2022-09-15 14:48:14'),
(6, 'PROD.5', 'Kartu Kredit', '2022-09-20 09:39:16', '2022-09-20 09:41:24');

-- --------------------------------------------------------

--
-- Table structure for table `ref_cronjobs`
--

CREATE TABLE `ref_cronjobs` (
  `time` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `method` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `action` varchar(10) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'add, remove',
  `flag` varchar(10) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '1 : tambah job ke service, any: tidak di eksekusi',
  `status` varchar(10) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `updated_at` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `ref_cronjobs`
--

INSERT INTO `ref_cronjobs` (`time`, `method`, `action`, `flag`, `status`, `updated_at`) VALUES
('0 0 0 * * *', 'updateDataMasterSDM', 'add', '2', 'RUN', '2022-02-16 14:45:40'),
('* * * * *', 'updateDataDashboard', 'add', '2', 'RUN', '2022-02-16 14:45:40'),
('0* * * * *', 'dioRemainder', 'add', '2', 'RUN', '2022-02-16 14:45:40');

-- --------------------------------------------------------

--
-- Table structure for table `risk_control`
--

CREATE TABLE `risk_control` (
  `id` int(11) NOT NULL,
  `kode` varchar(10) NOT NULL,
  `risk_control` text NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `risk_control`
--

INSERT INTO `risk_control` (`id`, `kode`, `risk_control`, `created_at`, `updated_at`) VALUES
(1, 'C1', 'Terdapat panduan/kebijakan/prosedur yang terkait', '2022-09-21 11:16:10', NULL),
(2, 'C3', 'Komunikasi dan koordinasi dengan pihak-pihak terkait', '2022-09-21 11:16:10', '2022-09-21 11:31:33');

-- --------------------------------------------------------

--
-- Table structure for table `risk_indicator`
--

CREATE TABLE `risk_indicator` (
  `id` int(10) UNSIGNED NOT NULL,
  `risk_indicator_code` varchar(100) NOT NULL,
  `risk_indicator` varchar(100) NOT NULL,
  `activity_id` int(11) NOT NULL,
  `product_id` int(11) NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `risk_indicator`
--

INSERT INTO `risk_indicator` (`id`, `risk_indicator_code`, `risk_indicator`, `activity_id`, `product_id`, `created_at`, `updated_at`) VALUES
(5, 'RNC1', 'Kas Kantor', 6, 1, '2022-09-14 16:08:43', NULL),
(6, 'RNC2', 'Kas Dalam Perjalanan', 6, 1, '2022-09-14 16:09:10', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `risk_issue`
--

CREATE TABLE `risk_issue` (
  `id` int(10) UNSIGNED NOT NULL,
  `risk_type_id` int(11) NOT NULL,
  `risk_issue_code` varchar(100) NOT NULL,
  `risk_issue` varchar(100) NOT NULL,
  `deskripsi` text NOT NULL,
  `kategori_risiko` varchar(100) NOT NULL,
  `status` varchar(20) NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `risk_issue`
--

INSERT INTO `risk_issue` (`id`, `risk_type_id`, `risk_issue_code`, `risk_issue`, `deskripsi`, `kategori_risiko`, `status`, `created_at`, `updated_at`) VALUES
(3, 1, 'AHP1', 'Kontrak/perjanjian dengan pihak ketiga tidak sesuai dengan kaidah hukum dan/atau memperlemah posisi ', 'Kontrak/perjanjian dengan pihak ketiga tidak sesuai dengan kaidah hukum dan/atau memperlemah posisi BRI	\n', 'BCM', 'Aktif', NULL, '2022-09-14 10:14:28'),
(4, 1, 'AHP2', 'Advise hukum yang diberikan belum memadai	\n', 'Advise hukum yang diberikan belum memadai	\n', 'BCM', 'Aktif', '2022-09-13 16:20:03', '0000-00-00 00:00:00');

-- --------------------------------------------------------

--
-- Table structure for table `risk_type`
--

CREATE TABLE `risk_type` (
  `id` int(10) UNSIGNED NOT NULL,
  `risk_type_code` varchar(100) NOT NULL,
  `risk_type` varchar(100) NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `risk_type`
--

INSERT INTO `risk_type` (`id`, `risk_type_code`, `risk_type`, `created_at`, `updated_at`) VALUES
(1, 'HUK0005', 'Risiko Hukum', '2022-08-18 11:41:35', NULL),
(2, 'KEP0007', 'Risiko Kepatuhan', '2022-08-18 11:42:38', '2022-09-15 11:49:06'),
(3, 'KRE0001', 'Risiko Kredit', '2022-09-15 11:58:59', '2022-09-15 14:11:29'),
(4, 'LIK0003', 'Risiko Likuiditas', '2022-09-15 14:14:47', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `sub_activity`
--

CREATE TABLE `sub_activity` (
  `id` int(10) UNSIGNED NOT NULL,
  `kode_sub_activity` varchar(10) NOT NULL,
  `activity_id` int(10) NOT NULL,
  `name` varchar(100) NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `sub_activity`
--

INSERT INTO `sub_activity` (`id`, `kode_sub_activity`, `activity_id`, `name`, `created_at`, `updated_at`) VALUES
(1, '1.1', 1, 'MR PASAR & LIKUIDITAS', '2022-08-15 11:09:49', '2022-08-23 09:39:29'),
(2, '1.2', 1, 'MR KREDIT', '2022-08-15 11:19:56', '2022-08-15 13:48:06'),
(3, '1.3', 1, 'MR OPERASIONAL & RISIKO LAIN', '2022-08-15 11:16:05', '2022-08-15 13:53:40'),
(4, '1.4', 1, 'ERM', '2022-08-15 11:16:05', '2022-08-15 13:54:15');

-- --------------------------------------------------------

--
-- Table structure for table `sub_incident_cause`
--

CREATE TABLE `sub_incident_cause` (
  `id` int(10) UNSIGNED NOT NULL,
  `kode_kejadian` varchar(100) NOT NULL,
  `kode_sub_kejadian` varchar(100) NOT NULL,
  `kriteria_penyebab_kejadian` varchar(100) NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `sub_incident_cause`
--

INSERT INTO `sub_incident_cause` (`id`, `kode_kejadian`, `kode_sub_kejadian`, `kriteria_penyebab_kejadian`, `created_at`, `updated_at`) VALUES
(1, 'PK1.MOP.0001', 'PK3.MOP.0001', 'Training tidak sesuai dengan kebutuhan', '2022-09-16 14:39:05', NULL),
(2, 'PK1.MOP.0001', 'PK3.MOP.0002', 'Perencanaan training tidak memadai', '2022-09-16 14:39:51', NULL),
(3, 'PK1.MOP.0001', 'PK3.MOP.0003', 'Frekuensi training tidak memadai', NULL, '2022-09-16 14:42:01');

-- --------------------------------------------------------

--
-- Table structure for table `unit_kerja`
--

CREATE TABLE `unit_kerja` (
  `id` int(10) UNSIGNED NOT NULL,
  `kode_uker` int(11) NOT NULL,
  `nama_uker` varchar(100) NOT NULL,
  `kode_cabang` int(11) NOT NULL,
  `nama_cabang` varchar(200) NOT NULL,
  `kanwil_id` int(11) NOT NULL,
  `kode_kanwil` varchar(10) NOT NULL,
  `kanwil` varchar(200) NOT NULL,
  `status` smallint(1) NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `unit_kerja`
--

INSERT INTO `unit_kerja` (`id`, `kode_uker`, `nama_uker`, `kode_cabang`, `nama_cabang`, `kanwil_id`, `kode_kanwil`, `kanwil`, `status`, `created_at`, `updated_at`) VALUES
(1, 37, 'KC BANDA ACEH CUT MEUTIA', 201, 'KANWIL MEDAN', 13, 'B', 'MEDAN', 1, '2022-08-22 09:55:48', '2022-08-23 07:53:26'),
(2, 42, 'KC LANGSA', 201, 'KANWIL MEDAN', 13, 'B', 'MEDAN', 1, '2022-08-23 07:54:41', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `verifikasi`
--

CREATE TABLE `verifikasi` (
  `id` int(10) UNSIGNED NOT NULL,
  `no_pelaporan` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `unit_kerja` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `activity_id` int(10) NOT NULL,
  `sub_activity_id` int(10) NOT NULL,
  `product_id` int(10) NOT NULL,
  `risk_issue_id` int(10) NOT NULL,
  `risk_indicator_id` int(10) NOT NULL,
  `incident_cause_id` int(10) NOT NULL,
  `sub_incident_cause_id` int(10) NOT NULL,
  `application_id` int(10) NOT NULL,
  `hasil_verifikasi` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `kunjungan_nasabah` tinyint(1) DEFAULT NULL,
  `indikasi_fraud` tinyint(1) DEFAULT NULL,
  `jenis_kerugian_finansial` tinyint(1) DEFAULT NULL,
  `jumlah_perkiraan_kerugian` int(10) DEFAULT NULL,
  `jenis_kerugian_non_finansial` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `rekomendasi_tindak_lanjut` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `rencana_tindak_lanjut` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `risk_type_id` int(10) NOT NULL,
  `tanggal_ditemukan` datetime DEFAULT NULL,
  `tanggal_mulai_rtl` datetime DEFAULT NULL,
  `tanggal_target_selesai` datetime DEFAULT NULL,
  `maker_id` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `maker_desc` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `maker_date` datetime DEFAULT NULL,
  `last_maker_id` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `last_maker_desc` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `last_maker_date` datetime DEFAULT NULL,
  `status` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `action` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `deleted` tinyint(1) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `verifikasi`
--

INSERT INTO `verifikasi` (`id`, `no_pelaporan`, `unit_kerja`, `activity_id`, `sub_activity_id`, `product_id`, `risk_issue_id`, `risk_indicator_id`, `incident_cause_id`, `sub_incident_cause_id`, `application_id`, `hasil_verifikasi`, `kunjungan_nasabah`, `indikasi_fraud`, `jenis_kerugian_finansial`, `jumlah_perkiraan_kerugian`, `jenis_kerugian_non_finansial`, `rekomendasi_tindak_lanjut`, `rencana_tindak_lanjut`, `risk_type_id`, `tanggal_ditemukan`, `tanggal_mulai_rtl`, `tanggal_target_selesai`, `maker_id`, `maker_desc`, `maker_date`, `last_maker_id`, `last_maker_desc`, `last_maker_date`, `status`, `action`, `deleted`, `created_at`, `updated_at`) VALUES
(1, 'VER-50046567-010922-0001', 'Kantor Cabang Khusus', 1, 1, 1, 1, 1, 1, 1, 1, 'Verifikasi ininih', 1, 1, 1, 200000000, 'ada deh', 'apaan ya', 'apaan ya', 1, '2022-09-01 13:37:00', '2022-09-01 13:37:00', '2022-09-30 00:00:00', '', '', NULL, '00304155', '00304155 | Feb\'hana Faradilla Bimantari | Teller', '2022-09-21 14:34:12', '02b', 'Update', 0, '2022-09-21 14:04:40', '2022-09-21 14:34:12');

-- --------------------------------------------------------

--
-- Table structure for table `verifikasi_data_anomali`
--

CREATE TABLE `verifikasi_data_anomali` (
  `id` int(10) UNSIGNED NOT NULL,
  `verifikasi_id` int(10) NOT NULL,
  `tanggal_kejadian` datetime DEFAULT NULL,
  `nomor_rekening` varchar(100) NOT NULL,
  `nominal` int(10) NOT NULL,
  `keterangan` text DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `verifikasi_data_anomali`
--

INSERT INTO `verifikasi_data_anomali` (`id`, `verifikasi_id`, `tanggal_kejadian`, `nomor_rekening`, `nominal`, `keterangan`) VALUES
(1, 1, '2022-09-01 13:37:00', '54321123456', 200000000, 'Data transfer Tidak Normal');

-- --------------------------------------------------------

--
-- Table structure for table `verifikasi_lampiran`
--

CREATE TABLE `verifikasi_lampiran` (
  `id` int(10) UNSIGNED NOT NULL,
  `verifikasi_id` int(10) UNSIGNED NOT NULL,
  `files_id` int(10) UNSIGNED NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `verifikasi_lampiran`
--

INSERT INTO `verifikasi_lampiran` (`id`, `verifikasi_id`, `files_id`) VALUES
(5, 1, 1);

-- --------------------------------------------------------

--
-- Table structure for table `verifikasi_pic_tindak_lanjut`
--

CREATE TABLE `verifikasi_pic_tindak_lanjut` (
  `id` int(10) UNSIGNED NOT NULL,
  `verifikasi_id` int(10) UNSIGNED NOT NULL,
  `pic_id` int(10) UNSIGNED NOT NULL,
  `tanggal_tindak_lanjut` datetime NOT NULL,
  `deskripsi_tindak_lanjut` text NOT NULL,
  `status` varchar(100) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `verifikasi_pic_tindak_lanjut`
--

INSERT INTO `verifikasi_pic_tindak_lanjut` (`id`, `verifikasi_id`, `pic_id`, `tanggal_tindak_lanjut`, `deskripsi_tindak_lanjut`, `status`) VALUES
(1, 1, 1, '2022-09-30 00:00:00', 'Akan dilakukan coaching lanjutan', '');

-- --------------------------------------------------------

--
-- Table structure for table `verifikasi_risk_control`
--

CREATE TABLE `verifikasi_risk_control` (
  `id` int(10) UNSIGNED NOT NULL,
  `verifikasi_id` int(10) UNSIGNED NOT NULL,
  `risk_control_id` int(10) UNSIGNED NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `verifikasi_risk_control`
--

INSERT INTO `verifikasi_risk_control` (`id`, `verifikasi_id`, `risk_control_id`) VALUES
(1, 1, 2);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `activity`
--
ALTER TABLE `activity`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `briefing`
--
ALTER TABLE `briefing`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `briefing_materis`
--
ALTER TABLE `briefing_materis`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `coaching`
--
ALTER TABLE `coaching`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `coaching_activity`
--
ALTER TABLE `coaching_activity`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `files`
--
ALTER TABLE `files`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `incident_cause`
--
ALTER TABLE `incident_cause`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `materi`
--
ALTER TABLE `materi`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `materi_files`
--
ALTER TABLE `materi_files`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `mst_access_menu`
--
ALTER TABLE `mst_access_menu`
  ADD PRIMARY KEY (`LevelUker`,`LevelID`,`IDMenu`) USING BTREE;

--
-- Indexes for table `mst_menu`
--
ALTER TABLE `mst_menu`
  ADD PRIMARY KEY (`IDMenu`) USING BTREE;

--
-- Indexes for table `product`
--
ALTER TABLE `product`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `risk_control`
--
ALTER TABLE `risk_control`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `risk_indicator`
--
ALTER TABLE `risk_indicator`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `risk_issue`
--
ALTER TABLE `risk_issue`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `risk_type`
--
ALTER TABLE `risk_type`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `sub_activity`
--
ALTER TABLE `sub_activity`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `sub_incident_cause`
--
ALTER TABLE `sub_incident_cause`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `unit_kerja`
--
ALTER TABLE `unit_kerja`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `verifikasi`
--
ALTER TABLE `verifikasi`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `verifikasi_data_anomali`
--
ALTER TABLE `verifikasi_data_anomali`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `verifikasi_lampiran`
--
ALTER TABLE `verifikasi_lampiran`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `verifikasi_pic_tindak_lanjut`
--
ALTER TABLE `verifikasi_pic_tindak_lanjut`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `verifikasi_risk_control`
--
ALTER TABLE `verifikasi_risk_control`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `activity`
--
ALTER TABLE `activity`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=26;

--
-- AUTO_INCREMENT for table `briefing`
--
ALTER TABLE `briefing`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=15;

--
-- AUTO_INCREMENT for table `briefing_materis`
--
ALTER TABLE `briefing_materis`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=22;

--
-- AUTO_INCREMENT for table `coaching`
--
ALTER TABLE `coaching`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- AUTO_INCREMENT for table `coaching_activity`
--
ALTER TABLE `coaching_activity`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=11;

--
-- AUTO_INCREMENT for table `files`
--
ALTER TABLE `files`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT for table `incident_cause`
--
ALTER TABLE `incident_cause`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=8;

--
-- AUTO_INCREMENT for table `materi`
--
ALTER TABLE `materi`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT for table `materi_files`
--
ALTER TABLE `materi_files`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT for table `mst_menu`
--
ALTER TABLE `mst_menu`
  MODIFY `IDMenu` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=24;

--
-- AUTO_INCREMENT for table `product`
--
ALTER TABLE `product`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;

--
-- AUTO_INCREMENT for table `risk_control`
--
ALTER TABLE `risk_control`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT for table `risk_indicator`
--
ALTER TABLE `risk_indicator`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;

--
-- AUTO_INCREMENT for table `risk_issue`
--
ALTER TABLE `risk_issue`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- AUTO_INCREMENT for table `risk_type`
--
ALTER TABLE `risk_type`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- AUTO_INCREMENT for table `sub_activity`
--
ALTER TABLE `sub_activity`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=9;

--
-- AUTO_INCREMENT for table `sub_incident_cause`
--
ALTER TABLE `sub_incident_cause`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT for table `unit_kerja`
--
ALTER TABLE `unit_kerja`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT for table `verifikasi`
--
ALTER TABLE `verifikasi`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `verifikasi_data_anomali`
--
ALTER TABLE `verifikasi_data_anomali`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `verifikasi_lampiran`
--
ALTER TABLE `verifikasi_lampiran`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- AUTO_INCREMENT for table `verifikasi_pic_tindak_lanjut`
--
ALTER TABLE `verifikasi_pic_tindak_lanjut`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `verifikasi_risk_control`
--
ALTER TABLE `verifikasi_risk_control`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
