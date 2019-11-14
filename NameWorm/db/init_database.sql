SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for t_first_name
-- ----------------------------
DROP TABLE IF EXISTS `t_first_name`;
CREATE TABLE `t_first_name`  (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `first_name` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

SET FOREIGN_KEY_CHECKS = 1;

-- ----------------------------
-- Table structure for t_suffix_name
-- ----------------------------
DROP TABLE IF EXISTS `t_suffix_name`;
CREATE TABLE `t_suffix_name`  (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `first_name_id` int(10) NOT NULL,
  `suffix_name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `sex` int(1) NOT NULL COMMENT '性别',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `fk_first_name`(`first_name_id`) USING BTREE,
  CONSTRAINT `fk_first_name` FOREIGN KEY (`first_name_id`) REFERENCES `t_first_name` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB AUTO_INCREMENT = 1439 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

SET FOREIGN_KEY_CHECKS = 1;
