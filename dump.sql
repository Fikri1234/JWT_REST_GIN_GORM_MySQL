-- MySQL dump 10.13  Distrib 8.0.22, for Win64 (x86_64)
--
-- Host: localhost    Database: app_data
-- ------------------------------------------------------
-- Server version	8.0.22

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `m_branches`
--

DROP TABLE IF EXISTS `m_branches`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `m_branches` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `code` varchar(255) DEFAULT NULL,
  `name` varchar(255) NOT NULL,
  `owner` varchar(255) DEFAULT NULL,
  `address` varchar(255) DEFAULT NULL,
  `npwp` varchar(255) DEFAULT NULL,
  `flag` tinyint(1) DEFAULT '1',
  `created_by` varchar(255) DEFAULT NULL,
  `updated_by` varchar(255) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_m_branches_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `m_branches`
--

LOCK TABLES `m_branches` WRITE;
/*!40000 ALTER TABLE `m_branches` DISABLE KEYS */;
INSERT INTO `m_branches` VALUES (1,'','P.D Parahyangan Djaya','Budiman','Jl Terusan Kiaracondong','123456789',1,'','','2021-01-11 00:04:53.124','2021-01-11 00:04:53.124',NULL);
/*!40000 ALTER TABLE `m_branches` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `m_menu_roles`
--

DROP TABLE IF EXISTS `m_menu_roles`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `m_menu_roles` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `m_menu_id` bigint NOT NULL,
  `m_user_role_id` bigint NOT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_m_menu_roles_m_menu` (`m_menu_id`),
  KEY `fk_m_menu_roles_m_user_role` (`m_user_role_id`),
  CONSTRAINT `fk_m_menu_roles_m_menu` FOREIGN KEY (`m_menu_id`) REFERENCES `m_menus` (`id`),
  CONSTRAINT `fk_m_menu_roles_m_user_role` FOREIGN KEY (`m_user_role_id`) REFERENCES `m_user_roles` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `m_menu_roles`
--

LOCK TABLES `m_menu_roles` WRITE;
/*!40000 ALTER TABLE `m_menu_roles` DISABLE KEYS */;
INSERT INTO `m_menu_roles` VALUES (1,1,1),(2,2,1),(3,3,1),(4,4,1),(5,5,1),(6,6,1),(7,7,1),(8,8,1),(9,9,1);
/*!40000 ALTER TABLE `m_menu_roles` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `m_menus`
--

DROP TABLE IF EXISTS `m_menus`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `m_menus` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `menu_name` varchar(255) DEFAULT NULL,
  `parent_menu_id` bigint DEFAULT NULL,
  `access_link` longtext,
  `created_by` varchar(255) DEFAULT NULL,
  `updated_by` varchar(255) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_m_menus_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `m_menus`
--

LOCK TABLES `m_menus` WRITE;
/*!40000 ALTER TABLE `m_menus` DISABLE KEYS */;
INSERT INTO `m_menus` VALUES (1,'Penmbelian',0,'#','','','2021-01-11 06:00:02.317','2021-01-11 06:00:02.317',NULL),(2,'Penjualan',0,'#','','','2021-01-11 06:00:14.996','2021-01-11 06:00:14.996',NULL),(3,'Form Operasional',1,'#','','','2021-01-11 06:00:54.570','2021-01-11 06:00:54.570',NULL),(4,'Form Penjualan',1,'#','','','2021-01-11 06:01:00.495','2021-01-11 06:01:00.495',NULL),(5,'Report',0,'#','','','2021-01-11 06:01:11.614','2021-01-11 06:01:11.614',NULL),(6,'Setting',0,'#','','','2021-01-11 06:01:23.645','2021-01-11 06:01:23.645',NULL),(7,'Tahunan',5,'#','','','2021-01-11 06:01:36.903','2021-01-11 06:01:36.903',NULL),(8,'Bulanan',5,'#','','','2021-01-11 06:01:41.203','2021-01-11 06:01:41.203',NULL),(9,'Harian',5,'#','','','2021-01-11 06:01:46.131','2021-01-11 06:01:46.131',NULL);
/*!40000 ALTER TABLE `m_menus` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `m_params`
--

DROP TABLE IF EXISTS `m_params`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `m_params` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `param_key` varchar(255) NOT NULL,
  `param_value` varchar(255) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `created_by` varchar(255) DEFAULT NULL,
  `updated_by` varchar(255) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_m_params_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `m_params`
--

LOCK TABLES `m_params` WRITE;
/*!40000 ALTER TABLE `m_params` DISABLE KEYS */;
INSERT INTO `m_params` VALUES (1,'EXPIRES_JWT','900','','','','2021-01-11 06:18:25.626','2021-01-11 06:18:25.626',NULL),(2,'EXPIRES_REFRESH_JWT','1800','','','','2021-01-11 06:18:39.664','2021-01-11 06:18:39.664',NULL);
/*!40000 ALTER TABLE `m_params` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `m_user_roles`
--

DROP TABLE IF EXISTS `m_user_roles`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `m_user_roles` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `role_code` varchar(255) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `flag` tinyint(1) DEFAULT '1',
  `created_by` varchar(255) DEFAULT NULL,
  `updated_by` varchar(255) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_m_user_roles_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `m_user_roles`
--

LOCK TABLES `m_user_roles` WRITE;
/*!40000 ALTER TABLE `m_user_roles` DISABLE KEYS */;
INSERT INTO `m_user_roles` VALUES (1,'ROLE_ADMIN','',1,'','','2021-01-11 00:55:28.367','2021-01-11 00:55:28.367',NULL);
/*!40000 ALTER TABLE `m_user_roles` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `m_users`
--

DROP TABLE IF EXISTS `m_users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `m_users` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_name` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `account_non_expired` tinyint(1) DEFAULT '1',
  `account_non_locked` tinyint(1) DEFAULT '1',
  `credentials_non_expired` tinyint(1) DEFAULT '1',
  `enabled` tinyint(1) DEFAULT '1',
  `email` longtext,
  `branch_id` bigint NOT NULL,
  `user_role_id` bigint NOT NULL,
  `change_pwd_counter` int DEFAULT NULL,
  `fail_counter` int DEFAULT NULL,
  `last_login` datetime(3) DEFAULT NULL,
  `reset_token` longtext,
  `created_by` varchar(255) DEFAULT NULL,
  `updated_by` varchar(255) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_name` (`user_name`),
  KEY `idx_m_users_deleted_at` (`deleted_at`),
  KEY `fk_m_users_m_branch` (`branch_id`),
  KEY `fk_m_users_m_user_role` (`user_role_id`),
  CONSTRAINT `fk_m_users_m_branch` FOREIGN KEY (`branch_id`) REFERENCES `m_branches` (`id`),
  CONSTRAINT `fk_m_users_m_user_role` FOREIGN KEY (`user_role_id`) REFERENCES `m_user_roles` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `m_users`
--

LOCK TABLES `m_users` WRITE;
/*!40000 ALTER TABLE `m_users` DISABLE KEYS */;
INSERT INTO `m_users` VALUES (1,'admin','$2a$10$adN.2gSmO0pTo6T3CMlsKeAm7gYHvypS6Mk7lK7RC4Xl2Vxqxhrte',1,1,1,1,'adsf@gmail.com',1,1,0,0,'0000-00-00 00:00:00.000','','','','2021-01-11 00:59:06.622','2021-01-11 05:54:44.398',NULL);
/*!40000 ALTER TABLE `m_users` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping routines for database 'app_data'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2021-01-11 19:02:10
