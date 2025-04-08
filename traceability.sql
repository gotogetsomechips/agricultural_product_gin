/*
 Navicat Premium Data Transfer

 Source Server         : mysql
 Source Server Type    : MySQL
 Source Server Version : 80033
 Source Host           : localhost:3306
 Source Schema         : traceability

 Target Server Type    : MySQL
 Target Server Version : 80033
 File Encoding         : 65001

 Date: 26/03/2025 14:16:53
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for company
-- ----------------------------
DROP TABLE IF EXISTS `company`;
CREATE TABLE `company`  (
  `com_id` int NOT NULL AUTO_INCREMENT,
  `com_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '物流公司名',
  `com_address` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '地址',
  `com_administrator` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '负责人',
  `com_phone` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '联系电话',
  PRIMARY KEY (`com_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for logistics
-- ----------------------------
DROP TABLE IF EXISTS `logistics`;
CREATE TABLE `logistics`  (
  `log_id` int NOT NULL AUTO_INCREMENT,
  `product_info_id` int NULL DEFAULT NULL COMMENT '生产信息id',
  `company_id` int NULL DEFAULT NULL COMMENT '物流公司id',
  `start_location` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '起点',
  `destination` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '目的地',
  `start_time` datetime NULL DEFAULT NULL COMMENT '出发时间',
  `end_time` datetime NULL DEFAULT NULL COMMENT '到达时间',
  PRIMARY KEY (`log_id`) USING BTREE,
  INDEX `product_info_id`(`product_info_id`) USING BTREE,
  INDEX `company_id`(`company_id`) USING BTREE,
  CONSTRAINT `logistics_ibfk_1` FOREIGN KEY (`product_info_id`) REFERENCES `product_info` (`pi_id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `logistics_ibfk_2` FOREIGN KEY (`company_id`) REFERENCES `company` (`com_id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for product
-- ----------------------------
DROP TABLE IF EXISTS `product`;
CREATE TABLE `product`  (
  `pd_id` int NOT NULL AUTO_INCREMENT,
  `pd_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '名称',
  `type` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '类别',
  `image` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '图片',
  `pd_description` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '具体描述',
  PRIMARY KEY (`pd_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for product_info
-- ----------------------------
DROP TABLE IF EXISTS `product_info`;
CREATE TABLE `product_info`  (
  `pi_id` int NOT NULL AUTO_INCREMENT,
  `product_id` int NULL DEFAULT NULL COMMENT '产品id',
  `product_place_id` int NULL DEFAULT NULL COMMENT '生产地id',
  `seed` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '种子来源',
  `pi_description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '记录生产内容',
  `planting_date` datetime NULL DEFAULT NULL COMMENT '播种时间',
  `harvest_date` datetime NULL DEFAULT NULL COMMENT '收获时间',
  PRIMARY KEY (`pi_id`) USING BTREE,
  INDEX `product_id`(`product_id`) USING BTREE,
  INDEX `product_place_id`(`product_place_id`) USING BTREE,
  CONSTRAINT `product_info_ibfk_1` FOREIGN KEY (`product_id`) REFERENCES `product` (`pd_id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `product_info_ibfk_2` FOREIGN KEY (`product_place_id`) REFERENCES `product_place` (`pp_id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for product_place
-- ----------------------------
DROP TABLE IF EXISTS `product_place`;
CREATE TABLE `product_place`  (
  `pp_id` int NOT NULL AUTO_INCREMENT,
  `pp_address` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '生产地址',
  `pp_administrator` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '负责人',
  `pp_phone` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '联系电话',
  PRIMARY KEY (`pp_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for sale_info
-- ----------------------------
DROP TABLE IF EXISTS `sale_info`;
CREATE TABLE `sale_info`  (
  `si_id` int NOT NULL AUTO_INCREMENT,
  `logistics_id` int NULL DEFAULT NULL COMMENT '物流信息id',
  `sale_place_id` int NULL DEFAULT NULL COMMENT '销售地id',
  `si_description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '销售相关内容',
  `sale_time` datetime NULL DEFAULT NULL COMMENT '销售时间',
  PRIMARY KEY (`si_id`) USING BTREE,
  INDEX `logistics_id`(`logistics_id`) USING BTREE,
  INDEX `sale_place_id`(`sale_place_id`) USING BTREE,
  CONSTRAINT `sale_info_ibfk_1` FOREIGN KEY (`logistics_id`) REFERENCES `logistics` (`log_id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `sale_info_ibfk_2` FOREIGN KEY (`sale_place_id`) REFERENCES `sale_place` (`sp_id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for sale_place
-- ----------------------------
DROP TABLE IF EXISTS `sale_place`;
CREATE TABLE `sale_place`  (
  `sp_id` int NOT NULL AUTO_INCREMENT,
  `sp_address` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '销售地址',
  `sp_administrator` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '负责人',
  `sp_phone` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '联系电话',
  PRIMARY KEY (`sp_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `username` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '用户名',
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '密码',
  `sex` varchar(2) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '性别',
  `name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '昵称',
  `phone` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '电话',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
