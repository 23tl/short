CREATE TABLE `short`
(
    `id`          bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',
    `old_url`     varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '旧链接',
    `new_url`     varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '生成的短连接',
    `status`      tinyint                                                        NOT NULL DEFAULT '1' COMMENT '状态',
    `visits`      int                                                            NOT NULL DEFAULT '1' COMMENT '访问次数',
    `extra`       json                                                                    DEFAULT NULL COMMENT '额外信息',
    `createdTime` int                                                            NOT NULL COMMENT '创建时间',
    `updatedTime` int                                                            NOT NULL COMMENT '更新时间',
    `deletedTime` int                                                                     DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`) USING BTREE,
    KEY           `idx_new_url` (`new_url`) USING BTREE,
    KEY           `idx_status` (`status`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=25 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC;