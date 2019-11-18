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

 Date: 18/11/2019 23:36:59
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for article
-- ----------------------------
DROP TABLE IF EXISTS `article`;
CREATE TABLE `article` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` text DEFAULT NULL,
  `subtitle` text DEFAULT NULL,
  `content` text DEFAULT NULL,
  `publisher` varchar(150) DEFAULT NULL,
  `published` date DEFAULT NULL,
  `created_on` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of article
-- ----------------------------
BEGIN;
INSERT INTO `article` VALUES (1, 'Oh My Gig Esport Tournament, Turnamen ML dan PUBGM Berhadiah Total 50 Juta Rupiah!', 'Yuk uji kemampuan main ML dan PUBGM dengan ikut Oh My Gig Esport Tournament. Total hadiah Rp50 juta menunggu kamu!', 'Uji kemampuan main ML dan PUBGM kamu disini', 'admin', '2019-11-18', '2019-11-18 19:26:02');
INSERT INTO `article` VALUES (2, 'Patahkan Kutukan, OG Keluar Sebagai Pemenang The International 2019!', 'Kemenangan OG di The International 2019 mematahkan kutukan ‘tak ada tim yang menang dua kali di The International Dota 2’!', 'Penggemar Dota 2 di seluruh penjuru dunia riuh setelah OG berhasil mengangkat Aegis Shield tanda mereka memenangkan The International 2019 dengan skor 3-1 melawan Team Liquid. Terlebihnya lagi, OG keluar sebagai juara TI dua kali berturut-turut dengan roster tim yang sama! Berbicara dengan ‘kutukan’, Kemenangan OG ini mematahkan beberapa ketentuan tak tertulis yang sudah dikenal di komunitas Dota 2. Ketentuan yang dikenal sebagai ‘kutukan’ ini adalah (1) tim barat tak akan menang TI di tahun ganjil; (2) tak ada tim yang menang dua kali; dan (3) tak ada tim yang menang berturut-turut. Hal ini sangat menarik karena OG kurang menunjukkan kemampuannya pada awal-awal TI9, namun bisa keluar sebagai sang jawara.\nInilah para pemain terbaik OG di The International 2019 :\nAnathan ‘ana’ Pham\nTopias ‘Topson’ Taavitsainen\nSabastien ‘Ceb’ Debs\nJesses ‘JerAx’ Vainikka\nJohan ‘N0tail’ Sundstein\n\nThe International 2019 Final Placements\nRank	Tim	Hadiah\n1	OG	Rp221,9 miliar\n2	Team Liquid	Rp63,4 miliar\n3	PSG.LGD	Rp43,9 miliar\n4	Team Secret	Rp29,2 miliar\n5-6	Evil Geniuses	Rp17 miliar\n5-6	ViCi Gaming	Rp17 miliar\n7-8	Royal Never Give Up	Rp12,2 miliar\n7-8	Infamous	Rp12,2 miliar\n9-12	Virtus.Pro	Rp9,7 miliar\n9-12	TNC Predator	Rp9,7 miliar\n9-12	Newbee	Rp9,7 miliar\n9-12	Mineski	Rp9,7 miliar\n13-16	Alliance	Rp7,3 miliar\n13-16	Fnatic	Rp7,3 miliar\n13-16	Keen Gaming	Rp7,3 miliar\n13-16	Natus Vincere	Rp7,3 miliar\n17-18	Ninjas in Pyjamas	Rp1,2 miliar\n17-18	Chaos Esports Club	Rp1,2 miliar\nSelanjutnya, The International 2020 sudah menetapkan tempat pelaksanaannya, yaitu di Stockholm, Swiss.\n\nSumber: Dexerto\nSumber gambar utama: youtube.com', 'admin', '2019-11-18', '2019-11-18 19:29:30');
INSERT INTO `article` VALUES (3, 'article baru', 'article baru lagi', 'ini article baru', 'mahmudin', '2019-11-18', '2019-11-18 23:27:43');
COMMIT;

-- ----------------------------
-- Table structure for contact
-- ----------------------------
DROP TABLE IF EXISTS `contact`;
CREATE TABLE `contact` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) DEFAULT NULL,
  `email` varchar(100) DEFAULT NULL,
  `phone` varchar(100) DEFAULT NULL,
  `message` text DEFAULT NULL,
  `created_on` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of contact
-- ----------------------------
BEGIN;
INSERT INTO `contact` VALUES (1, 'Fauzan', 'fauzanrab@gmail.com', '081297555987', 'hallo', '2019-11-18 19:14:25');
COMMIT;

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
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=latin1;

-- ----------------------------
-- Records of users
-- ----------------------------
BEGIN;
INSERT INTO `users` VALUES (1, 'ojanganteng', 'Fauzan Rabbani', 'fauzanrab@gmail.com', '$2a$10$lM0B/oerHtGVtUKrywNPd.Ir6RoUnIwGBKOFff/G8WXg4lOj5vop6', 'admin');
INSERT INTO `users` VALUES (2, 'cimong', 'Valerie Stephanie', 'valkrie@gmail.com', '$2a$10$cicDss4PTafB97jeJZDOD.d5etO2f4AOJnEpFdz4qnLAsN7jarUfu', 'user');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
