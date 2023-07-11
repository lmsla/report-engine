INSERT INTO `instances` (
    `created_at`,
    `updated_at`,
    `type`,`name`,`url`,`user`,`password`,`auth`)
VALUES
(UNIX_TIMESTAMP(), UNIX_TIMESTAMP(),'kibana','kibana_110','http://10.99.1.241:5601/','elastic','12345678',1),
(UNIX_TIMESTAMP(), UNIX_TIMESTAMP(),'grafana','grafana_241','http://10.99.1.241:3000','admin','12345678',1);

