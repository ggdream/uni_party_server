CREATE TABLE IF NOT EXISTS `video_archive`
(
    `id`              INT(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增ID',
    `vid`             VARCHAR(8)       NOT NULL COMMENT '视频ID',
    `star_counter`    MEDIUMINT(7) UNSIGNED DEFAULT 0 COMMENT '点赞数量',
    `collect_counter` MEDIUMINT(6) UNSIGNED DEFAULT 0 COMMENT '收藏数量',
    `watch_counter`   INT(9) UNSIGNED       DEFAULT 0 COMMENT '播放次数',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;
