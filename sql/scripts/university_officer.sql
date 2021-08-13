# Redis
CREATE TABLE IF NOT EXISTS `university_officer`
(
    `id`       INT(11) UNSIGNED    NOT NULL AUTO_INCREMENT COMMENT '自增ID',
    `uid`      INT(11) UNSIGNED    NOT NULL COMMENT '用户ID',
    `rank`     TINYINT(1) UNSIGNED NOT NULL COMMENT '组级别',
    `sub_rank` TINYINT(1) UNSIGNED NOT NULL COMMENT '阶级别',
    `code`     CHAR(8)             NOT NULL COMMENT '高校',
    `campus`   VARCHAR(10)         NOT NULL COMMENT '校区',
    `college`  VARCHAR(10)         NOT NULL COMMENT '学院',
    `grade`    TINYINT(2) UNSIGNED NOT NULL COMMENT '年级',
    `major`    VARCHAR(10)         NOT NULL COMMENT '专业',
    `class`    TINYINT(3) UNSIGNED NOT NULL COMMENT '班级',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;
