CREATE TABLE IF NOT EXISTS `message_archive`
(
    `id`             INT(11) UNSIGNED    NOT NULL AUTO_INCREMENT COMMENT '自增ID',
    `mid`            VARCHAR(8)          NOT NULL COMMENT '消息ID',
    `get_counter`    SMALLINT(5) UNSIGNED  DEFAULT 0 COMMENT '收到数量',
    `follow_counter` SMALLINT(5) UNSIGNED  DEFAULT 0 COMMENT '关注数量',
    `watch_counter`  MEDIUMINT(8) UNSIGNED DEFAULT 0 COMMENT '查看次数',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;
