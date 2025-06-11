CREATE TABLE `tables` (
  `created_at` bigint unsigned DEFAULT NULL,
  `updated_at` bigint unsigned DEFAULT NULL,
  `deleted_at` bigint DEFAULT NULL,
  `id` bigint NOT NULL AUTO_INCREMENT,
  `type` varchar(50) DEFAULT NULL,
  `name` varchar(50) DEFAULT NULL,
  `data_view` varchar(50) DEFAULT NULL,
  `uid` varchar(50) DEFAULT NULL,
  `row_num` int DEFAULT NULL,
  `instance_id` bigint DEFAULT NULL,
  `space_name` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_tables_id` (`id`),
  KEY `fk_tables_instance` (`instance_id`),
  CONSTRAINT `fk_tables_instance` FOREIGN KEY (`instance_id`) REFERENCES `instances` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=30 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci