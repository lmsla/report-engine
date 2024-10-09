CREATE TABLE `report_tables` (
  `report_id` bigint NOT NULL,
  `table_id` bigint NOT NULL,
  PRIMARY KEY (`report_id`,`table_id`),
  KEY `fk_report_tables_table` (`table_id`),
  CONSTRAINT `fk_report_tables_report` FOREIGN KEY (`report_id`) REFERENCES `reports` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `fk_report_tables_table` FOREIGN KEY (`table_id`) REFERENCES `tables` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci