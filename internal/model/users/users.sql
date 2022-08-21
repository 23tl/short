CREATE TABLE `users`
(
    `id`          bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',
    `name`        varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '用户名',
    `mobile`      varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '手机号',
    `password`    varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci          DEFAULT NULL COMMENT '密码',
    `avatar`      varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '头像',
    `status`      tinyint                                                       NOT NULL DEFAULT '1' COMMENT '状态',
    `extra`       json                                                          DEFAULT NULL COMMENT '额外信息',
    `createdTime` int                                                           NOT NULL COMMENT '创建时间',
    `updatedTime` int                                                           NOT NULL COMMENT '更新时间',
    `deletedTime` int                                                                    DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `mobile` (`mobile`) USING BTREE,
    KEY           `idx_name` (`name`) USING BTREE,
    KEY           `idx_status` (`status`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=25 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC;