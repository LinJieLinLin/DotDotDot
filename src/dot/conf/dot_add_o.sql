/*
Navicat MySQL Data Transfer

Source Server         : localhost
Source Server Version : 50625
Source Host           : localhost:3306
Source Database       : dot

Target Server Type    : MYSQL
Target Server Version : 50625
File Encoding         : 65001

Date: 2016-01-22 16:06:37
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for d_goods
-- ----------------------------
DROP TABLE IF EXISTS `d_goods`;
CREATE TABLE `d_goods` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `no` varchar(255) DEFAULT NULL COMMENT '商品编号',
  `name` varchar(255) DEFAULT NULL COMMENT '商品名称',
  `depict` varchar(255) DEFAULT NULL COMMENT '商品描述',
  `status` int(3) DEFAULT NULL COMMENT '状态（20~39）',
  `e_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '编辑时间',
  `c_time` datetime DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for d_goods_detail
-- ----------------------------
DROP TABLE IF EXISTS `d_goods_detail`;
CREATE TABLE `d_goods_detail` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `pic` varchar(255) DEFAULT NULL COMMENT '图片',
  `d_name` varchar(255) DEFAULT NULL COMMENT '属性名称',
  `num` varchar(5) DEFAULT NULL COMMENT '属性编号',
  `price` double(13,2) DEFAULT NULL COMMENT '价格',
  `def` int(1) DEFAULT '1' COMMENT '是否是默认属性',
  `status` int(3) DEFAULT NULL COMMENT '状态（20~39）',
  `e_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '编辑时间',
  `c_time` datetime DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for d_name
-- ----------------------------
DROP TABLE IF EXISTS `d_name`;
CREATE TABLE `d_name` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `k` varchar(255) DEFAULT NULL COMMENT 'key',
  `v` varchar(255) DEFAULT NULL COMMENT 'value',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=52 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for d_store
-- ----------------------------
DROP TABLE IF EXISTS `d_store`;
CREATE TABLE `d_store` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `num` varchar(255) DEFAULT NULL COMMENT '商家编号',
  `name` varchar(255) DEFAULT NULL COMMENT '商家名称',
  `owner_id` int(11) DEFAULT NULL COMMENT '所有者ID',
  `s_type` int(3) DEFAULT NULL COMMENT '商家类型(50~59)',
  `status` int(3) DEFAULT NULL COMMENT '状态（20~39）',
  `icon` varchar(255) DEFAULT NULL COMMENT '商标',
  `about` varchar(255) DEFAULT NULL COMMENT '关于商铺',
  `address` varchar(255) DEFAULT NULL COMMENT '地址',
  `e_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '编辑时间',
  `c_time` datetime DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for d_user
-- ----------------------------
DROP TABLE IF EXISTS `d_user`;
CREATE TABLE `d_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `account` varchar(50) NOT NULL COMMENT '帐号',
  `password` varchar(255) NOT NULL COMMENT '密码',
  `name` varchar(50) DEFAULT NULL COMMENT '用户名',
  `tel` int(11) DEFAULT NULL COMMENT '手机号',
  `pic` varchar(255) DEFAULT NULL COMMENT '头像',
  `email` varchar(50) DEFAULT NULL COMMENT '邮件',
  `role_id` int(11) DEFAULT NULL COMMENT '角色ID',
  `status` int(11) DEFAULT NULL COMMENT '状态',
  `e_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '编辑时间',
  `c_time` datetime DEFAULT NULL COMMENT '创建时间',
  `last_ip` varchar(30) DEFAULT NULL COMMENT '最后登陆IP',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for tem_buy
-- ----------------------------
DROP TABLE IF EXISTS `tem_buy`;
CREATE TABLE `tem_buy` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `uId` int(11) NOT NULL,
  `mId` int(11) NOT NULL,
  `time` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for tem_menu
-- ----------------------------
DROP TABLE IF EXISTS `tem_menu`;
CREATE TABLE `tem_menu` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `type` int(2) DEFAULT NULL,
  `price` decimal(10,2) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for tem_name
-- ----------------------------
DROP TABLE IF EXISTS `tem_name`;
CREATE TABLE `tem_name` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8;
