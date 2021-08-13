CREATE TABLE IF NOT EXISTS `user_info`
(
    `id`            INT(11) UNSIGNED    NOT NULL AUTO_INCREMENT COMMENT '用户ID',
    `identify_type` VARCHAR(64)         NOT NULL COMMENT '证件类型',
    `identify_no`   VARCHAR(36)         NOT NULL COMMENT '证件编号',
    `telephone`     CHAR(11)            NOT NULL COMMENT '国内手机号',
    `email`         VARCHAR(50)         NOT NULL COMMENT '电子邮箱',
    `birthday`      DATE                NOT NULL COMMENT '生日',
    `sex`           CHAR(1)             NOT NULL COMMENT '性别',
    `avatar`        VARCHAR(2083)       NOT NULL COMMENT '头像URL',
    `user_type`     TINYINT(1) UNSIGNED NOT NULL COMMENT '用户类别',
    `user_rank`     TINYINT(1) UNSIGNED NOT NULL COMMENT '用户组级别',
    `user_sub_rank` TINYINT(1) UNSIGNED NOT NULL COMMENT '用户阶级别',
    `followers`     INT(9) UNSIGNED     NOT NULL COMMENT '粉丝数：亿',
    `following`     INT(9) UNSIGNED     NOT NULL COMMENT '关注数：亿',
    `username`      VARCHAR(64)         NOT NULL COMMENT '用户名称',
    `password`      CHAR(44)            NOT NULL COMMENT '哈希密码',
    `salt`          CHAR(8)             NOT NULL COMMENT '盐',
    `status`        TINYINT(1) UNSIGNED NOT NULL COMMENT '用户状态',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 100000
  DEFAULT CHARSET = utf8mb4;
