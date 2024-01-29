CREATE TABLE `histories` (
  `created_at` bigint unsigned DEFAULT NULL,
  `updated_at` bigint unsigned DEFAULT NULL,
  `deleted_at` bigint DEFAULT NULL,
  `schedule_id` bigint DEFAULT NULL,
  `schedule_name` varchar(50) DEFAULT NULL,
  `to` varchar(50) DEFAULT NULL,
  `cc` varchar(50) DEFAULT NULL,
  `bcc` varchar(50) DEFAULT NULL,
  `execute_time` bigint unsigned DEFAULT NULL,
  `email_time` bigint unsigned DEFAULT NULL,
  `success` varchar(50) DEFAULT NULL,
  KEY `fk_histories_schedule` (`schedule_id`),
  CONSTRAINT `fk_histories_schedule` FOREIGN KEY (`schedule_id`) REFERENCES `schedules` (`id`) ON DELETE RESTRICT ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci