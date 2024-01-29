CREATE TABLE `reports_schedules` (
  `report_id` bigint NOT NULL,
  `schedule_id` bigint NOT NULL,
  PRIMARY KEY (`report_id`,`schedule_id`),
  KEY `fk_schedules_reports` (`schedule_id`),
  CONSTRAINT `fk_reports_schedules` FOREIGN KEY (`report_id`) REFERENCES `reports` (`id`),
  CONSTRAINT `fk_schedules_reports` FOREIGN KEY (`schedule_id`) REFERENCES `schedules` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci