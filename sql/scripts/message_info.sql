CREATE TABLE IF NOT EXISTS `message_info`
(
    `id`         INT(11) UNSIGNED    NOT NULL AUTO_INCREMENT COMMENT '自增ID',
    `uid`        INT(11) UNSIGNED    NOT NULL COMMENT '用户ID',
    `mid`        VARCHAR(12) DEFAULT '' COMMENT '消息ID',
    `type`       TINYINT(1) UNSIGNED NOT NULL COMMENT '消息类型',
    `title`      VARCHAR(64)         NOT NULL COMMENT '标题',
    `content`    VARCHAR(2083)       NOT NULL COMMENT '文本内容的URL',
    `constraint` VARCHAR(65535)      NOT NULL COMMENT '消息约束',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;
