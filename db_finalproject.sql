CREATE DATABASE  IF NOT EXISTS `db_finalproject` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `db_finalproject`;
-- MySQL dump 10.13  Distrib 8.0.27, for Win64 (x86_64)
--
-- Host: localhost    Database: db_finalproject
-- ------------------------------------------------------
-- Server version	8.0.27

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `bookmarks`
--

DROP TABLE IF EXISTS `bookmarks`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `bookmarks` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `game_name` longtext,
  `id_game` bigint DEFAULT NULL,
  `ratings` bigint DEFAULT NULL,
  `image_url` longtext,
  `user_id` bigint unsigned DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_bookmarks_user` (`user_id`),
  CONSTRAINT `fk_bookmarks_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `bookmarks`
--

LOCK TABLES `bookmarks` WRITE;
/*!40000 ALTER TABLE `bookmarks` DISABLE KEYS */;
INSERT INTO `bookmarks` VALUES (1,'Example Games: Reborn',2,0,'http://example-game-2-logo.jpg',3,'2022-04-03 22:34:49.087','2022-04-03 22:34:49.087'),(2,'Coding Games',4,0,'http://coding-games-logo.jpg',3,'2022-04-03 22:36:26.644','2022-04-03 22:36:26.644');
/*!40000 ALTER TABLE `bookmarks` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `categories`
--

DROP TABLE IF EXISTS `categories`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `categories` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` longtext NOT NULL,
  `description` longtext,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `categories`
--

LOCK TABLES `categories` WRITE;
/*!40000 ALTER TABLE `categories` DISABLE KEYS */;
INSERT INTO `categories` VALUES (1,'Single-player','Played by One Player','2022-04-03 22:03:31.387','2022-04-03 22:03:31.387'),(2,'Multiplayer','More than one person can play in the same game environment at the same time','2022-04-03 22:04:04.538','2022-04-03 22:04:04.538'),(3,'LAN','Local area network games','2022-04-03 22:04:26.763','2022-04-03 22:04:26.763'),(4,'MMO','Game that enables hundreds or thousands of players to simultaneously interact in a game world they are connected to via the Internet','2022-04-03 22:04:45.708','2022-04-03 22:04:45.708'),(5,'Co-operative','Game that allows players to work together as teammates, usually against one or more non-player character opponents (PvE)','2022-04-03 22:05:12.253','2022-04-03 22:05:12.253');
/*!40000 ALTER TABLE `categories` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `games`
--

DROP TABLE IF EXISTS `games`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `games` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` longtext NOT NULL,
  `ratings` bigint DEFAULT NULL,
  `ratings_counter` bigint DEFAULT NULL,
  `release_date` datetime(3) DEFAULT NULL,
  `description` longtext,
  `image_url` longtext,
  `genre_id` bigint DEFAULT NULL,
  `category_id` bigint DEFAULT NULL,
  `publisher_id` bigint DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_publishers_game` (`publisher_id`),
  KEY `fk_genres_game` (`genre_id`),
  KEY `fk_categories_game` (`category_id`),
  CONSTRAINT `fk_categories_game` FOREIGN KEY (`category_id`) REFERENCES `categories` (`id`),
  CONSTRAINT `fk_genres_game` FOREIGN KEY (`genre_id`) REFERENCES `genres` (`id`),
  CONSTRAINT `fk_publishers_game` FOREIGN KEY (`publisher_id`) REFERENCES `publishers` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `games`
--

LOCK TABLES `games` WRITE;
/*!40000 ALTER TABLE `games` DISABLE KEYS */;
INSERT INTO `games` VALUES (1,'Example Games',4,2,'2022-04-03 22:23:17.524','Games example created by example publisher','http://example-game-logo.jpg',3,1,1,'2022-04-03 22:23:17.524','2022-04-03 22:45:50.411'),(2,'Example Games: Reborn',0,0,'2022-04-03 22:24:35.939','Second games example created by example publisher','http://example-game-2-logo.jpg',6,1,1,'2022-04-03 22:24:35.939','2022-04-03 22:24:35.939'),(3,'Golang Games',0,0,'2022-04-03 22:26:34.651','Third games example created by example publisher','http://example-game-3-logo.jpg',3,4,1,'2022-04-03 22:26:34.651','2022-04-03 22:26:34.651'),(4,'Coding Games',0,0,'2022-04-03 22:27:50.613','Fourth games example created by example publisher','http://coding-games-logo.jpg',1,2,1,'2022-04-03 22:27:50.613','2022-04-03 22:27:50.613'),(5,'Last Games',0,0,'2022-04-03 22:28:25.019','Last games example created by example publisher','http://last-games-logo.jpg',1,3,1,'2022-04-03 22:28:25.019','2022-04-03 22:28:25.019');
/*!40000 ALTER TABLE `games` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `genres`
--

DROP TABLE IF EXISTS `genres`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `genres` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` longtext NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `genres`
--

LOCK TABLES `genres` WRITE;
/*!40000 ALTER TABLE `genres` DISABLE KEYS */;
INSERT INTO `genres` VALUES (1,'Action','2022-04-03 22:06:08.124','2022-04-03 22:06:08.124'),(2,'Adventure','2022-04-03 22:06:26.532','2022-04-03 22:06:26.532'),(3,'Sports','2022-04-03 22:06:30.980','2022-04-03 22:06:30.980'),(4,'Racing','2022-04-03 22:06:41.178','2022-04-03 22:06:41.178'),(5,'Music','2022-04-03 22:06:50.421','2022-04-03 22:06:50.421'),(6,'RPG','2022-04-03 22:06:58.128','2022-04-03 22:06:58.128');
/*!40000 ALTER TABLE `genres` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `installed_games`
--

DROP TABLE IF EXISTS `installed_games`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `installed_games` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `game_name` longtext,
  `id_game` bigint DEFAULT NULL,
  `ratings` bigint DEFAULT NULL,
  `image_url` longtext,
  `user_id` bigint unsigned DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_installed_games_user` (`user_id`),
  CONSTRAINT `fk_installed_games_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `installed_games`
--

LOCK TABLES `installed_games` WRITE;
/*!40000 ALTER TABLE `installed_games` DISABLE KEYS */;
INSERT INTO `installed_games` VALUES (1,'Example Games',1,5,'http://example-game-logo.jpg',3,'2022-04-03 22:38:12.133','2022-04-03 22:38:12.133');
/*!40000 ALTER TABLE `installed_games` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `publishers`
--

DROP TABLE IF EXISTS `publishers`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `publishers` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` longtext,
  `logo_url` longtext,
  `user_id` bigint unsigned DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_publishers_user` (`user_id`),
  CONSTRAINT `fk_publishers_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `publishers`
--

LOCK TABLES `publishers` WRITE;
/*!40000 ALTER TABLE `publishers` DISABLE KEYS */;
INSERT INTO `publishers` VALUES (1,'publisher','http://publisher-logo.jpg',2,'2022-04-03 22:11:29.920','2022-04-03 22:11:29.920');
/*!40000 ALTER TABLE `publishers` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `reviews`
--

DROP TABLE IF EXISTS `reviews`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `reviews` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `rate` bigint DEFAULT NULL,
  `content` longtext,
  `game_id` bigint DEFAULT NULL,
  `user_id` bigint unsigned DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_games_review` (`game_id`),
  KEY `fk_users_review` (`user_id`),
  CONSTRAINT `fk_games_review` FOREIGN KEY (`game_id`) REFERENCES `games` (`id`),
  CONSTRAINT `fk_users_review` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `reviews`
--

LOCK TABLES `reviews` WRITE;
/*!40000 ALTER TABLE `reviews` DISABLE KEYS */;
INSERT INTO `reviews` VALUES (1,5,'This game is amazing!!',1,3,'2022-04-03 22:35:27.320','2022-04-03 22:35:27.320'),(2,3,'So many bug in this game :(',1,4,'2022-04-03 22:45:50.405','2022-04-03 22:45:50.405');
/*!40000 ALTER TABLE `reviews` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(191) NOT NULL,
  `email` varchar(191) NOT NULL,
  `password` longtext NOT NULL,
  `role` longtext,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`),
  UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (1,'admin','admin@email.com','$2a$10$F1rbFzL1gB/YOQKqid/wguxhWw5JqFz04m.TJZJ0XRmHr7oBFGNYm','admin','2022-04-03 21:57:42.854','2022-04-03 22:00:17.965'),(2,'publisher','publisher@email.com','$2a$10$3iLeHcRF.hMTs/w2RbUsB.R5YxVlV07HVsoPZ5P9DW36p0zpGoiYC','publisher','2022-04-03 22:09:27.738','2022-04-03 22:11:29.911'),(3,'user','user@email.com','$2a$10$2mQJbXSQRYr/LskyrMGEceX4./eGA3YZ0//WeYdMyiEvenCKNLv0y','user','2022-04-03 22:30:25.719','2022-04-03 22:30:25.719'),(4,'fajar','fajar@email.com','$2a$10$/mD/LDh6UhBJKwoy6JC/LeY6FjoCfpnRKV3NDy4ryuvm9MK3kS.6S','user','2022-04-03 22:39:44.100','2022-04-03 22:40:52.903');
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-04-04  4:15:26
