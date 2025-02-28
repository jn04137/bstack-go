-- Adminer 4.8.1 MySQL 11.5.2-MariaDB-ubu2404 dump

SET NAMES utf8;
SET time_zone = '+00:00';
SET foreign_key_checks = 0;
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

SET NAMES utf8mb4;

DROP TABLE IF EXISTS `player`;
CREATE TABLE `player` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `nano_id` varchar(22) NOT NULL,
  `username` varchar(64) NOT NULL,
  `password` varchar(128) NOT NULL,
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `nano_id` (`nano_id`),
  UNIQUE KEY `username` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;

INSERT INTO `player` (`id`, `nano_id`, `username`, `password`, `created_at`, `updated_at`) VALUES
(1,	'wYg4xSVwfkQjDuqXkM6dR',	'wthunder',	'$2a$10$hiY.6LcMex0JTqQ.7.MqLugUuWU6.5YgpJZuhW6eIRDZ2Mv8Z3wTa',	'2025-02-10 04:11:20',	'2025-02-10 04:11:20'),
(2,	'DVYC5v1wzs1K92FYZL0e7',	'prime',	'$2a$10$Dp9mIL80D.65aFbqPzoODu/xcmO/dknRGL5N.EQ049q6sjW4pzttC',	'2025-02-13 03:58:09',	'2025-02-17 16:32:14'),
(3,	'yv95bTG-X6a13-OYqLnaG',	'cheeto',	'$2a$10$hiY.6LcMex0JTqQ.7.MqLugUuWU6.5YgpJZuhW6eIRDZ2Mv8Z3wTa',	'2025-02-17 15:27:00',	'2025-02-17 15:28:20'),
(4,	'Qsu-WUpK1Sp-vZjpAQ7t3',	'dorito',	'$2a$10$hiY.6LcMex0JTqQ.7.MqLugUuWU6.5YgpJZuhW6eIRDZ2Mv8Z3wTa',	'2025-02-17 15:27:28',	'2025-02-17 15:28:14'),
(5,	'zsJK8OM1kWAUfwhPhyPNr',	'optimus',	'$2a$10$hiY.6LcMex0JTqQ.7.MqLugUuWU6.5YgpJZuhW6eIRDZ2Mv8Z3wTa',	'2025-02-17 15:28:03',	'2025-02-17 15:28:03');

DROP TABLE IF EXISTS `player_details`;
CREATE TABLE `player_details` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `player` int(11) NOT NULL,
  `details` text DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `player` (`player`),
  CONSTRAINT `player_details_ibfk_1` FOREIGN KEY (`player`) REFERENCES `player` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;


DROP TABLE IF EXISTS `schema_migrations`;
CREATE TABLE `schema_migrations` (
  `version` bigint(20) NOT NULL,
  `dirty` tinyint(1) NOT NULL,
  PRIMARY KEY (`version`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;

INSERT INTO `schema_migrations` (`version`, `dirty`) VALUES
(3,	0);

DROP TABLE IF EXISTS `team`;
CREATE TABLE `team` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `nano_id` varchar(22) NOT NULL,
  `team_owner` int(11) DEFAULT NULL,
  `team_name` varchar(64) NOT NULL,
  `team_desc` text DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `nano_id` (`nano_id`),
  UNIQUE KEY `team_name` (`team_name`),
  KEY `team_owner` (`team_owner`),
  CONSTRAINT `team_ibfk_1` FOREIGN KEY (`team_owner`) REFERENCES `player` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;


-- 2025-02-28 03:54:04
