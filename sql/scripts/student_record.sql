CREATE TABLE IF NOT EXISTS `student_record`
(
    `id`           INT(11) UNSIGNED    NOT NULL AUTO_INCREMENT COMMENT '自增ID',
    `student_id`   VARCHAR(16)         NOT NULL COMMENT '学号',
    `student_name` VARCHAR(16)         NOT NULL COMMENT '名字',
    `id_card`      VARCHAR(16)         NOT NULL COMMENT '身份证号',
    `code`         CHAR(8)             NOT NULL COMMENT '高校',
    `campus`       VARCHAR(10)         NOT NULL COMMENT '校区',
    `college`      VARCHAR(10)         NOT NULL COMMENT '学院',
    `grade`        TINYINT(2) UNSIGNED NOT NULL COMMENT '年级',
    `major`        VARCHAR(10)         NOT NULL COMMENT '专业',
    `class`        TINYINT(3) UNSIGNED NOT NULL COMMENT '班级',
    `status`       TINYINT(1) UNSIGNED NOT NULL COMMENT '学生状态，休学等',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;
