CREATE TABLE IF NOT EXISTS `video_info`
(
    `id`       INT(11) UNSIGNED    NOT NULL AUTO_INCREMENT COMMENT '自增ID',
    `uid`      INT(11) UNSIGNED    NOT NULL COMMENT '用户ID',
    `vid`      VARCHAR(12) DEFAULT '' COMMENT '视频ID',
    `title`    VARCHAR(64)         NOT NULL COMMENT '标题',
    `tags`     VARCHAR(72)         NOT NULL COMMENT '视频标签',
    `cover`    VARCHAR(2083)       NOT NULL COMMENT '封面地址',
    `video`    VARCHAR(2083)       NOT NULL COMMENT '视频地址',
    `location` VARCHAR(64) DEFAULT '' COMMENT '位置',
    `status`   TINYINT(1) UNSIGNED NOT NULL COMMENT '视频状态',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;
