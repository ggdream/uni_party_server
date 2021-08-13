CREATE TABLE IF NOT EXISTS `reply_ancestor`
(
    `id`                INT(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增ID',
    `uid`               INT(11) UNSIGNED NOT NULL COMMENT '用户ID',
    `aid`               VARCHAR(8)       NOT NULL COMMENT '祖评论ID',
    `pid`               VARCHAR(8)       NOT NULL COMMENT '父评论ID',
    `rrid`              VARCHAR(8)       NOT NULL COMMENT '评论ID',
    `content`           VARCHAR(65535)   NOT NULL COMMENT '评论内容',
    `star_counter`      MEDIUMINT(7) UNSIGNED DEFAULT 0 COMMENT '点赞数量',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;
