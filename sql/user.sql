SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `account_id` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '账户ID，全局唯一',
  `user_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户名',
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户密码md5值',
  `password_salt` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '密码salt值',
  `sex` tinyint(1) DEFAULT NULL COMMENT '性别，1-男，2-女',
  `phone` char(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '手机号',
  `country_code` char(5) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '手机区号',
  `email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '邮箱',
  `state` tinyint(1) NOT NULL DEFAULT '3' COMMENT '状态，0-未激活，1-审核中，2-审核未通过，3-已审核',
  `id_card_no` char(18) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '身份证号',
  `inviter` bigint DEFAULT NULL COMMENT '邀请人uid',
  `invite_code` char(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '邀请码',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  -- `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `contact_addr` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '联系地址',
  `age` int DEFAULT NULL COMMENT '年龄',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `account_id_index` (`account_id`) USING BTREE COMMENT '账户ID索引',
  UNIQUE KEY `country_code_phone_index` (`country_code`,`phone`) USING BTREE COMMENT '手机号索引',
  KEY `user_name_index` (`user_name`) USING BTREE COMMENT '用户名索引',
  KEY `email_index` (`email`) USING BTREE COMMENT '邮箱索引',
  KEY `id_card_no_index` (`id_card_no`) USING BTREE COMMENT '身份证号索引',
  KEY `invite_code_index` (`invite_code`) USING BTREE COMMENT '邀请码索引'
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户信息表';