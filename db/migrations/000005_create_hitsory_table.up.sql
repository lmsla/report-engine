CREATE TABLE `histories` (
  `created_at` bigint unsigned DEFAULT NULL,
  `updated_at` bigint unsigned DEFAULT NULL,
  `deleted_at` bigint DEFAULT NULL,
  `schedule_id` bigint DEFAULT NULL,
  `schedule_name` varchar(50) DEFAULT NULL,
  `report_id` bigint DEFAULT NULL,
  `to` varchar(50) DEFAULT NULL,
  `cc` varchar(50) DEFAULT NULL,
  `bcc` varchar(50) DEFAULT NULL,
  `execute_time` bigint unsigned DEFAULT NULL,
  `email_time` bigint unsigned DEFAULT NULL,
  `success` tinyint(1) DEFAULT '0',
  KEY `fk_histories_schedule` (`schedule_id`),
  KEY `fk_histories_report` (`report_id`),
  CONSTRAINT `fk_histories_report` FOREIGN KEY (`report_id`) REFERENCES `reports` (`id`) ON DELETE SET NULL ON UPDATE CASCADE,
  CONSTRAINT `fk_histories_schedule` FOREIGN KEY (`schedule_id`) REFERENCES `schedules` (`id`) ON DELETE SET NULL ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci