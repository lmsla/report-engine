CREATE TABLE `histories` (
  `created_at` bigint unsigned DEFAULT NULL,
  `updated_at` bigint unsigned DEFAULT NULL,
  `deleted_at` bigint DEFAULT NULL,
  `id` bigint NOT NULL AUTO_INCREMENT,
  `schedule_id` bigint DEFAULT NULL,
  `schedule_name` varchar(50) DEFAULT NULL,
  `to` text,
  `cc` text,
  `bcc` text,
  `execute_time` bigint unsigned DEFAULT NULL,
  `email_time` bigint unsigned DEFAULT NULL,
  `success` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_histories_schedule` (`schedule_id`),
  CONSTRAINT `fk_histories_schedule` FOREIGN KEY (`schedule_id`) REFERENCES `schedules` (`id`) ON DELETE SET NULL ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=1964 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci