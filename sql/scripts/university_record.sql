CREATE TABLE IF NOT EXISTS `university_info`
(
    `id`       INT(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增ID',
    `uid`      INT(11) UNSIGNED NOT NULL COMMENT '用户ID',
    `index`    CHAR(5)          NOT NULL COMMENT '国家高校代码',
    `code`     CHAR(8)          NOT NULL COMMENT '自设高校编码',
    `name`     VARCHAR(16)      NOT NULL COMMENT '高校名称',
    `province` VARCHAR(10)      NOT NULL COMMENT '省份',
    `city`     VARCHAR(8)       NOT NULL COMMENT '城市',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;
