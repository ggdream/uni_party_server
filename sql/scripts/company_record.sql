CREATE TABLE IF NOT EXISTS `company_record`
(
    `id`       INT(11) UNSIGNED    NOT NULL AUTO_INCREMENT COMMENT '自增ID',
    `uid`      INT(11) UNSIGNED    NOT NULL COMMENT '用户ID',
    `index`    VARCHAR(16)         NOT NULL COMMENT '工商行政管理注册号',
    `name`     VARCHAR(16)         NOT NULL COMMENT '公司名称',
    `industry` TINYINT(3) UNSIGNED NOT NULL COMMENT '行业类别',
    `province` VARCHAR(10)         NOT NULL COMMENT '省份',
    `city`     VARCHAR(8)          NOT NULL COMMENT '城市',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;
