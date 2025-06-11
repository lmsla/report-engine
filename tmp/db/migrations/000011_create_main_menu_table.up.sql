CREATE TABLE `main_menus` (
  `id` int NOT NULL AUTO_INCREMENT,
  `router_path` varchar(100) DEFAULT NULL,
  `icon` varchar(30) DEFAULT NULL,
  `title` varchar(30) DEFAULT NULL,
  `sort` bigint unsigned DEFAULT NULL,
  `module` varchar(20) DEFAULT NULL,
  `only_admin` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci