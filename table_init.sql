/*
 Navicat Premium Data Transfer

 Source Server         : 127.0.0.1
 Source Server Type    : MySQL
 Source Server Version : 80022
 Source Host           : 127.0.0.1:3306
 Source Schema         : names

 Target Server Type    : MySQL
 Target Server Version : 80022
 File Encoding         : 65001

 Date: 13/11/2020 10:09:39
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for t_first_name
-- ----------------------------
DROP TABLE IF EXISTS `t_first_name`;
CREATE TABLE `t_first_name`  (
  `id` int(0) NOT NULL AUTO_INCREMENT,
  `first_name` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `last_page` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 388 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for t_name_extends
-- ----------------------------
DROP TABLE IF EXISTS `t_name_extends`;
CREATE TABLE `t_name_extends`  (
  `id` bigint(0) NOT NULL AUTO_INCREMENT,
  `destiny` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '天格',
  `destiny_content` varchar(2000) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '天格内容',
  `ground` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '地格',
  `ground_content` varchar(2000) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '地格内容',
  `people` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '人格',
  `people_content` varchar(2000) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '人格内容',
  `other` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '外格',
  `other_content` varchar(2000) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '外格内容',
  `all` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '总格',
  `all_content` varchar(2000) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '总格内容',
  `talents` varchar(2000) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '三才内容',
  `influences` varchar(2000) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '影响内容',
  `name_id` bigint(0) DEFAULT NULL COMMENT '所属名字',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_name`(`name_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 40508 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for t_name_score
-- ----------------------------
DROP TABLE IF EXISTS `t_name_score`;
CREATE TABLE `t_name_score`  (
  `id` bigint(0) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '名称',
  `score` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '分数',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_name`(`name`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 40523 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for t_suffix_name
-- ----------------------------
DROP TABLE IF EXISTS `t_suffix_name`;
CREATE TABLE `t_suffix_name`  (
  `id` int(0) NOT NULL AUTO_INCREMENT,
  `first_name_id` int(0) NOT NULL,
  `suffix_name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `sex` int(0) NOT NULL COMMENT '性别',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `fk_first_name`(`first_name_id`) USING BTREE,
  CONSTRAINT `fk_first_name` FOREIGN KEY (`first_name_id`) REFERENCES `t_first_name` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB AUTO_INCREMENT = 2439468 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

SET FOREIGN_KEY_CHECKS = 1;
