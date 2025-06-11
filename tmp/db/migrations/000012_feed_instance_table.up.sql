INSERT INTO `instances` (
    `created_at`,
    `updated_at`,
    `deleted_at`,
    `type`,`name`,`url`,`es_url`,`user`,`password`,`auth`)
VALUES
(UNIX_TIMESTAMP(), UNIX_TIMESTAMP(),NULL,'kibana','kibana93','http://10.99.1.93:5601','https://10.99.1.93:9200','elastic','12345678',1);