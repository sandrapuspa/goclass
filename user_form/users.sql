/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 100408
 Source Host           : localhost:3306
 Source Schema         : test

 Target Server Type    : MySQL
 Target Server Version : 100408
 File Encoding         : 65001

 Date: 15/11/2019 21:03:14
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(50) DEFAULT NULL,
  `name` varchar(200) NOT NULL,
  `email` varchar(200) NOT NULL,
  `password` varchar(120) DEFAULT NULL,
  `role` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=latin1;

-- ----------------------------
-- Records of users
-- ----------------------------
BEGIN;
INSERT INTO `users` VALUES (1, 'ojanganteng', 'Fauzan Rabbani', 'fauzanrab@gmail.com', '$2a$10$lM0B/oerHtGVtUKrywNPd.Ir6RoUnIwGBKOFff/G8WXg4lOj5vop6', 'admin');
INSERT INTO `users` VALUES (2, 'cimong', 'Valerie Stephanie', 'valkrie@gmail.com', '$2a$10$cicDss4PTafB97jeJZDOD.d5etO2f4AOJnEpFdz4qnLAsN7jarUfu', 'user');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
