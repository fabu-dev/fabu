/*
 Navicat Premium Data Transfer

 Source Server         : 87.中台
 Source Server Type    : MySQL
 Source Server Version : 80016
 Source Host           : 10.32.5.87:3306
 Source Schema         : fabu.dev

 Target Server Type    : MySQL
 Target Server Version : 80016
 File Encoding         : 65001

 Date: 11/09/2020 15:32:08
*/


CREATE DATABASE IF NOT EXISTS `fabu.dev`;

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for app
-- ----------------------------
DROP TABLE IF EXISTS `app`;
CREATE TABLE `app`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'appID',
  `name` varchar(25) NOT NULL COMMENT 'app名称',
  `group_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT 'app归属的团队ID',
  `platform` tinyint(3) UNSIGNED NOT NULL DEFAULT 1 COMMENT '平台（1：iOS，2：Android）',
  `icon` varchar(150) NOT NULL COMMENT 'app的图标',
  `short_url` varchar(15) NOT NULL COMMENT '短连接',
  `bunde_id` varchar(100) NOT NULL COMMENT '包名',
  `current_version` varchar(10) NOT NULL COMMENT '当前版本',
  `key` varchar(30) NOT NULL COMMENT 'APP KEY',
  `browse_count` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '浏览次数',
  `donwload_count` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '下载次数',
  `status` tinyint(3) UNSIGNED NOT NULL DEFAULT 1 COMMENT '状态（ 1 正常 ）',
  `updated_at` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '更新时间',
  `updated_by` varchar(50) NOT NULL DEFAULT '' COMMENT '修改人',
  `created_at` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `created_by` varchar(50) NOT NULL DEFAULT '' COMMENT '添加人',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 COMMENT = 'app主表';

-- ----------------------------
-- Table structure for download_log
-- ----------------------------
DROP TABLE IF EXISTS `download_log`;
CREATE TABLE `download_log`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `app_id` int(10) UNSIGNED NOT NULL,
  `version_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0,
  `ip` bigint(20) UNSIGNED NOT NULL DEFAULT 0,
  `download_date` date NOT NULL DEFAULT '2000-01-01',
  `created_at` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `created_by` varchar(50) NOT NULL DEFAULT '' COMMENT '添加人',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB;

-- ----------------------------
-- Table structure for group
-- ----------------------------
DROP TABLE IF EXISTS `group`;
CREATE TABLE `group`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `owner` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '团队所有人',
  `name` varchar(25) DEFAULT NULL COMMENT '团队名',
  `updated_at` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '更新时间',
  `updated_by` varchar(50) NOT NULL DEFAULT '' COMMENT '修改人',
  `created_at` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `created_by` varchar(50) NOT NULL DEFAULT '' COMMENT '添加人',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 COMMENT = '团队表';

-- ----------------------------
-- Table structure for group_member
-- ----------------------------
DROP TABLE IF EXISTS `group_member`;
CREATE TABLE `group_member`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `group_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '团队ID',
  `member_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '会员ID',
  `role` tinyint(3) UNSIGNED NOT NULL DEFAULT 1 COMMENT '角色（1 普通成员 2 管理员 3 所有人）',
  `updated_at` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '更新时间',
  `updated_by` varchar(50) NOT NULL DEFAULT '' COMMENT '修改人',
  `created_at` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `created_by` varchar(50) NOT NULL DEFAULT '' COMMENT '添加人',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_member_id`(`member_id`) USING BTREE,
  INDEX `idx_group_id`(`group_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 COMMENT = '团队成员表';

-- ----------------------------
-- Table structure for member
-- ----------------------------
DROP TABLE IF EXISTS `member`;
CREATE TABLE `member`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `mobile` varchar(15) NOT NULL,
  `account` varchar(50) DEFAULT NULL COMMENT '登录账户',
  `password` varchar(80) DEFAULT NULL COMMENT '密码',
  `email` varchar(80) DEFAULT NULL COMMENT '邮箱',
  `token` varchar(80) DEFAULT NULL COMMENT 'APIToken',
  `updated_at` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '更新时间',
  `updated_by` varchar(50) NOT NULL DEFAULT '' COMMENT '修改人',
  `created_at` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `created_by` varchar(50) NOT NULL DEFAULT '' COMMENT '添加人',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uk_token`(`token`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1;

-- ----------------------------
-- Table structure for version
-- ----------------------------
DROP TABLE IF EXISTS `version`;
CREATE TABLE `version`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'APP版本ID',
  `app_id` bigint(20) UNSIGNED NOT NULL COMMENT 'APPID',
  `tag` varchar(10) NOT NULL COMMENT '版本tag',
  `code` varchar(10) NOT NULL COMMENT '版本号',
  `description` varchar(300) NOT NULL COMMENT '更新说明',
  `size` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '包大小',
  `hash` varchar(80) NOT NULL COMMENT 'hash',
  `path` varchar(200) NOT NULL COMMENT '文件存放路径（如果上传到云平台这里是url）',
  `is_publish` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否发布（0：未发布，1：发布）',
  `updated_at` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '更新时间',
  `updated_by` varchar(50) NOT NULL DEFAULT '' COMMENT '修改人',
  `created_at` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `created_by` varchar(50) NOT NULL DEFAULT '' COMMENT '添加人',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 COMMENT = 'app版本表';

SET FOREIGN_KEY_CHECKS = 1;
