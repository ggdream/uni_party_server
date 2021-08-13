CREATE TABLE IF NOT EXISTS `reply_ancestor`
(
    `id`                INT(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增ID',
    `uid`               INT(11) UNSIGNED NOT NULL COMMENT '用户ID',
    `vid`               VARCHAR(8)       NOT NULL COMMENT '视频ID',
    `rid`               VARCHAR(8)       NOT NULL COMMENT '视频ID',
    `content`           VARCHAR(65535)   NOT NULL COMMENT '评论内容',
    `star_counter`      MEDIUMINT(7) UNSIGNED DEFAULT 0 COMMENT '点赞数量',
    `son_reply_counter` SMALLINT(5) UNSIGNED  DEFAULT 0 COMMENT '子评数量',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;
