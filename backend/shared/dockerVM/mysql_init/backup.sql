-- MySQL dump 10.13  Distrib 8.0.31, for Linux (x86_64)
--
-- Host: localhost    Database: car_rental
-- ------------------------------------------------------
-- Server version       8.0.31

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
-- Table structure for table `client`
--

DROP TABLE IF EXISTS `client`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `client` (
  `ID` int NOT NULL AUTO_INCREMENT,
  `name` varchar(30) NOT NULL,
  `surname` varchar(30) NOT NULL,
  `PESEL` varchar(11) NOT NULL,
  `phone_number` varchar(9) NOT NULL,
  `user_id` int NOT NULL,
  PRIMARY KEY (`ID`),
  UNIQUE KEY `PESEL` (`PESEL`),
  UNIQUE KEY `phone_number` (`phone_number`),
  UNIQUE KEY `user_id` (`user_id`),
  CONSTRAINT `FKclient245213` FOREIGN KEY (`user_id`) REFERENCES `user` (`ID`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=38 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `client`
--

LOCK TABLES `client` WRITE;
/*!40000 ALTER TABLE `client` DISABLE KEYS */;
INSERT INTO `client` VALUES (4,'asd','aasdsd','a12asd11','a1asd22',11),(5,'asd','aasdsd','a12asd11a','a1asd22a',12),(7,'asd','aasdsd','admin','admin',14),(8,'asd','aasdsd','owner1','owner1',15),
(9,'asd','aasdsd','owner12','owner12',16),(22,'Szymon','Nowak','11112311','1231111',30),(23,'Szymon','Nowak','1111231111','123111111',31),(24,'Szymon','Nowak','333','333',32),(25,'Szymon','Nowak
','3333','3333',33),(26,'admin','admin','admin1','admin1',34),(28,'admin','admin','admin2','admin2',36),(29,'Szymon','Nowak','12312312312','123123123',37),(30,'admin','admin','test','test',38),(
33,'działajtj kurczee','admin','test1','test1',41),(34,'admin','admin','user','user',42),(36,'client','client','client','client',44),(37,'tt','tt','tt','tt',45);
/*!40000 ALTER TABLE `client` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `employee`
--

DROP TABLE IF EXISTS `employee`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `employee` (
  `ID` int NOT NULL AUTO_INCREMENT,
  `name` varchar(30) NOT NULL,
  `surname` varchar(30) DEFAULT NULL,
  `salary` int NOT NULL,
  `PESEL` varchar(255) NOT NULL,
  `user_id` int NOT NULL,
  PRIMARY KEY (`ID`),
  UNIQUE KEY `user_id` (`user_id`),
  CONSTRAINT `FKemployee939388` FOREIGN KEY (`user_id`) REFERENCES `user` (`ID`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `employee`
--

LOCK TABLES `employee` WRITE;
/*!40000 ALTER TABLE `employee` DISABLE KEYS */;
/*!40000 ALTER TABLE `employee` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `notification`
--

DROP TABLE IF EXISTS `notification`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `notification` (
  `ID` int NOT NULL AUTO_INCREMENT,
  `description` varchar(255) NOT NULL,
  `employee_id` int NOT NULL,
  `client_id` int DEFAULT NULL,
  `vechicle_id` int DEFAULT NULL,
  PRIMARY KEY (`ID`),
  UNIQUE KEY `description` (`description`),
  KEY `FKnotificati466077` (`vechicle_id`),
  KEY `FKnotificati174904` (`employee_id`),
  KEY `FKnotificati402236` (`client_id`),
  CONSTRAINT `FKnotificati174904` FOREIGN KEY (`employee_id`) REFERENCES `employee` (`ID`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `FKnotificati402236` FOREIGN KEY (`client_id`) REFERENCES `client` (`ID`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `FKnotificati466077` FOREIGN KEY (`vechicle_id`) REFERENCES `vehicle` (`ID`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `notification`
--

LOCK TABLES `notification` WRITE;
/*!40000 ALTER TABLE `notification` DISABLE KEYS */;
/*!40000 ALTER TABLE `notification` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `rental`
--

DROP TABLE IF EXISTS `rental`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `rental` (
  `ID` int NOT NULL AUTO_INCREMENT,
  `start_date` date NOT NULL,
  `end_date` date NOT NULL,
  `pickup_address` varchar(255) DEFAULT NULL,
  `driver_id` int DEFAULT NULL,
  `client_id` int NOT NULL,
  `vehicle_id` int NOT NULL,
  PRIMARY KEY (`ID`),
  KEY `FKrental154746` (`vehicle_id`),
  KEY `FKrental590843` (`client_id`),
  KEY `FKrental398356` (`driver_id`),
  CONSTRAINT `FKrental154746` FOREIGN KEY (`vehicle_id`) REFERENCES `vehicle` (`ID`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `FKrental398356` FOREIGN KEY (`driver_id`) REFERENCES `employee` (`ID`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `FKrental590843` FOREIGN KEY (`client_id`) REFERENCES `client` (`ID`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=42 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `rental`
--

LOCK TABLES `rental` WRITE;
/*!40000 ALTER TABLE `rental` DISABLE KEYS */;
INSERT INTO `rental` VALUES (25,'2009-11-10','2009-11-11','owner',NULL,7,5),(26,'2009-11-10','2009-11-11','owner',NULL,7,13),(27,'2009-11-10','2009-11-11','owner',NULL,36,5),(28,'2009-11-10','20
09-11-11','owner',NULL,28,5),(29,'2009-11-10','2009-11-11','owner',NULL,36,5),(31,'2009-11-10','2009-11-11','owner',NULL,36,13),(32,'2009-11-10','2009-11-11','owner',NULL,36,13),(33,'2009-11-10'
,'2009-11-11','owner',NULL,36,13),(34,'2019-10-12','2019-10-13','owner',NULL,36,14),(35,'2019-10-12','2019-10-13','owner',NULL,36,14),(36,'2019-10-12','2019-10-13','owner',NULL,7,14),(37,'2019-1
0-12','2019-10-13','owner',NULL,7,14),(38,'2019-10-12','2019-10-13','owner',NULL,7,14),(39,'2019-10-12','2019-10-13','owner',NULL,7,14),(40,'2019-10-12','2019-10-13','owner',NULL,7,18),(41,'2019
-10-12','2019-10-13','owner',NULL,7,17);
/*!40000 ALTER TABLE `rental` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `repairs`
--

DROP TABLE IF EXISTS `repairs`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `repairs` (
  `ID` int NOT NULL AUTO_INCREMENT,
  `cost` int NOT NULL,
  `approved` bit(1) DEFAULT NULL,
  `vehicle_id` int NOT NULL,
  `notification_id` int DEFAULT NULL,
  PRIMARY KEY (`ID`),
  UNIQUE KEY `notification_id` (`notification_id`),
  KEY `FKrepairs601556` (`vehicle_id`),
  CONSTRAINT `FKrepairs601556` FOREIGN KEY (`vehicle_id`) REFERENCES `vehicle` (`ID`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `FKrepairs934942` FOREIGN KEY (`notification_id`) REFERENCES `notification` (`ID`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `repairs`
--

LOCK TABLES `repairs` WRITE;
/*!40000 ALTER TABLE `repairs` DISABLE KEYS */;
/*!40000 ALTER TABLE `repairs` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `salary`
--

DROP TABLE IF EXISTS `salary`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `salary` (
  `date` date NOT NULL,
  `amount` int NOT NULL,
  `employee_id` int NOT NULL,
  KEY `FKsalary843084` (`employee_id`),
  CONSTRAINT `FKsalary843084` FOREIGN KEY (`employee_id`) REFERENCES `employee` (`ID`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `salary`
--

LOCK TABLES `salary` WRITE;
/*!40000 ALTER TABLE `salary` DISABLE KEYS */;
/*!40000 ALTER TABLE `salary` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user` (
  `ID` int NOT NULL AUTO_INCREMENT,
  `email` varchar(50) NOT NULL,
  `password` varchar(255) NOT NULL,
  `role` varchar(20) NOT NULL,
  PRIMARY KEY (`ID`),
  UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=46 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES (1,'asd@asd.com','$2a$14$w4UulzcBH8Hzxo./iEOKY.DKE7MPRptuLjwn0Iiap0W6sS1q0BgnK','admin'),(4,'asd@asd.buziacezee3dk','incorrrect','cos'),(7,'1233','incorrect','cos'),(8,
'12133','incorrect','cos'),(9,'asd','incorrect','cos'),(10,'owner','owner','owner'),(11,'ownerr','$2a$14$7ykbYtXhvGtnPDD95UatPOoe5f5kcuqHuIYDiXqZs9ltf8fewGebu','owner'),(12,'ownerra','$2a$14$R8U
VupjtavVlprJs7uNN7eWNqgcrKsXMaw4AAExcJNrfFoktqVZmG','owner'),(14,'admin','$2a$14$IXi4nWupEqDanhNvmJj03O0m19fHsT7R12RCrV1DYHJZjhdZ8rElu','admin'),(15,'owner1','$2a$14$dNN/G09jj3zaerBk3SkSzuTPOZRm
uhLJn9MZvkWZuWur2UaIepM12','owner1'),(16,'owner2','$2a$14$D7/ykNHLIW4XXfkhWOX1lu3mv5twQ0aZxbqkDQlvVqRPsGPFZ83EK','owner'),(30,'szymek.nowak00@gmail.com','$2a$14$dFpbkqLIA5r6BUadP/Pu3urM/DVyo.nMO
MMtYMWIZ5ECP/Kpr3EFe',''),(31,'szymek.nowaak00@gmail.com','$2a$14$olpgSGXyd9og3gRR3y04O.51IeKNcKXnBssmAH3.EGfAjdproxnQq',''),(32,'Nikodem.dudek5@gmail.com','$2a$14$4P4nboV66RlERXw0hk9UA.xEHGCtEC
29MmJ2Kt4gIbjZOAp88tYku',''),(33,'Nikodem.dudek25@gmail.com','$2a$14$AnBgkSQ/LBJygcT2CFU8dugJ5ebAUYRofeOS7JyayeOq1xuThLmnS','client'),(34,'admin1','$2a$14$Q2Mmsda3IjwgiPQNw8jaq.UcOpnuoZx2dxwenZk
c2NLTUEUzdEiii','admin'),(36,'admin2','$2a$14$L7laa0bCt7fsCEKEj2IIOO1uHtv3cS2mhUU1Zz1ONqWvNWwCDoBS6','admin'),(37,'szymeek.nowak00@gmail.com','$2a$14$buni/yDeLS5G1Yq0eMYkguZMMkdLQDa2uwkShMdQ5gq6
4w3auqRF.',''),(38,'test','$2a$14$Zth2mUMwOtGW1Pw5sOgATeuCT4uA2.sxNelejes7J84bPuav9VlSu','test'),(41,'test12ww','$2a$14$.F3kr8o9j2B/rTAM82QTVO5lyysZQBZqpiWh51eYjcU.N17hVlqDW','testt1w'),(42,'use
r','$2a$14$9XJHJv8SZwyoEuyFR67dBujBuwbk/fuWgUjaa47m5AD1CHY6v5LKm','client'),(44,'client','$2a$14$LWi19YOnyTaa06IO.tlsn.ZQO6zyPgAn5V6tt/UibWb7HDQpGPQpy','client'),(45,'tt','$2a$14$0DO4B4vhxrNDo/l
Xf3uC4.B3KZQk.sY.elqi0oEYafPvrstBO5w/S','client');
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `vehicle`
--

DROP TABLE IF EXISTS `vehicle`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `vehicle` (
  `ID` int NOT NULL AUTO_INCREMENT,
  `registration_number` varchar(7) NOT NULL,
  `brand` varchar(50) NOT NULL,
  `model` varchar(50) NOT NULL,
  `type` varchar(30) NOT NULL,
  `color` varchar(20) NOT NULL,
  `fuel_consumption` float NOT NULL,
  `daily_cost` int NOT NULL,
  PRIMARY KEY (`ID`),
  UNIQUE KEY `registration_number` (`registration_number`)
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `vehicle`
--

LOCK TABLES `vehicle` WRITE;
/*!40000 ALTER TABLE `vehicle` DISABLE KEYS */;
INSERT INTO `vehicle` VALUES (5,'3','Opel','Astra','hatchback','czarny',5.8,140),(13,'9','Porche','911','sportowy','czerwony',11.2,250),(14,'99','Porche','911','sportowy','czerwony',11.2,250),(1
5,'11','Porche','911','sportowy','czerwony',11.2,250),(16,'12','Porche','911','sportowy','czerwony',11.2,250),(17,'13','Porche','911','sportowy','czerwony',11.2,250),(18,'15','Volkswagen','Golf'
,'hatchback','czarny',0.3,120);
/*!40000 ALTER TABLE `vehicle` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-12-28 13:59:54