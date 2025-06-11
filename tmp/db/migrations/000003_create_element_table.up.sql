CREATE TABLE `elements` (
  `created_at` bigint unsigned DEFAULT NULL,
  `updated_at` bigint unsigned DEFAULT NULL,
  `deleted_at` bigint DEFAULT NULL,
  `id` bigint NOT NULL AUTO_INCREMENT,
  `report_id` bigint DEFAULT NULL,
  `type` varchar(50) DEFAULT NULL,
  `name` varchar(50) DEFAULT NULL,
  `uid` varchar(50) DEFAULT NULL,
  `row_num` bigint DEFAULT NULL,
  `column_type` varchar(50) DEFAULT NULL,
  `instance_id` bigint DEFAULT NULL,
  `space_name` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_elements_id` (`id`),
  KEY `idx_elements_report_id` (`report_id`),
  KEY `fk_elements_instance` (`instance_id`),
  CONSTRAINT `fk_elements_instance` FOREIGN KEY (`instance_id`) REFERENCES `instances` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=40 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci