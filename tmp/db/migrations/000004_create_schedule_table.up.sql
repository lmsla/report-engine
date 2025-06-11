CREATE TABLE `schedules` (
  `created_at` bigint unsigned DEFAULT NULL,
  `updated_at` bigint unsigned DEFAULT NULL,
  `deleted_at` bigint DEFAULT NULL,
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` varchar(50) DEFAULT NULL,
  `cron_time` varchar(50) DEFAULT NULL,
  `to` text,
  `cc` text,
  `bcc` text,
  `cron_id` bigint DEFAULT NULL,
  `subject` varchar(255) DEFAULT NULL,
  `body` varchar(255) DEFAULT NULL,
  `enable` tinyint DEFAULT '1',
  PRIMARY KEY (`id`),
  KEY `idx_schedules_id` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci


