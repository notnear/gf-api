SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for api_user
-- ----------------------------
DROP TABLE IF EXISTS `api_user`;
CREATE TABLE `api_user` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `mobile` char(11) COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '电话号码',
  `nickname` varchar(128) COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '昵称',
  `avatar` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '头像',
  `sex` tinyint(1) DEFAULT '0' COMMENT '性别 0 未知 1男 2女',
  `status` tinyint(1) DEFAULT '0' COMMENT '状态',
  `create_at` datetime DEFAULT NULL COMMENT '添加时间',
  `update_at` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of api_user
-- ----------------------------
BEGIN;
INSERT INTO `api_user` VALUES (1, '13076045652', '昵称123', 'https://baidu.com/1.png', 1, 0, '2020-11-24 11:27:31', '2020-11-24 11:27:31');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
