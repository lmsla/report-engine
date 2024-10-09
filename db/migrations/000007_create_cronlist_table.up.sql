CREATE TABLE `cron_lists` (
  `created_at` bigint unsigned DEFAULT NULL,
  `updated_at` bigint unsigned DEFAULT NULL,
  `deleted_at` bigint DEFAULT NULL,
  `schedule_id` bigint NOT NULL,
  `entry_id` bigint NOT NULL,
  PRIMARY KEY (`schedule_id`),
  KEY `idx_cron_list_id` (`schedule_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci