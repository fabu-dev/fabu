/*
 Navicat Premium Data Transfer

 Source Server         : local.32.master
 Source Server Type    : MySQL
 Source Server Version : 50716
 Source Host           : 192.168.3.132:8805
 Source Schema         : fabu.dev

 Target Server Type    : MySQL
 Target Server Version : 50716
 File Encoding         : 65001

 Date: 14/10/2020 17:02:30
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

CREATE DATABASE IF NOT EXISTS `fabu.dev`;
USE `fabu.dev`

-- ----------------------------
-- Table structure for app
-- ----------------------------
DROP TABLE IF EXISTS `app`;
CREATE TABLE `app`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'appID',
  `name` varchar(25) NOT NULL DEFAULT '' COMMENT 'app名称',
  `team_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT 'app归属的团队ID',
  `platform` tinyint(3) UNSIGNED NOT NULL DEFAULT 1 COMMENT '平台（1：iOS，2：Android）',
  `icon` varchar(150) NOT NULL DEFAULT '' COMMENT 'app的图标',
  `short_url` varchar(15) NOT NULL DEFAULT '' COMMENT '短连接',
  `bundle_id` varchar(100) NOT NULL DEFAULT '' COMMENT '包名',
  `current_version` varchar(10) NOT NULL DEFAULT '' COMMENT '当前版本',
  `key` varchar(30) NOT NULL DEFAULT '' COMMENT 'APP KEY',
  `browse_count` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '浏览次数',
  `download_count` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '下载次数',
  `status` tinyint(3) UNSIGNED NOT NULL DEFAULT 1 COMMENT '状态（1：删除，2有效）',
  `updated_at` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '更新时间',
  `updated_by` varchar(50) NOT NULL DEFAULT '' COMMENT '修改人',
  `created_at` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `created_by` varchar(50) NOT NULL DEFAULT '' COMMENT '添加人',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COMMENT = 'app主表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for download_log
-- ----------------------------
DROP TABLE IF EXISTS `app_download_log`;
CREATE TABLE `app_download_log`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `app_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT 'APP ID',
  `version_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '版本ID',
  `ip` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '下载人IP',
  `download_date` date NOT NULL DEFAULT '2000-01-01' COMMENT '下载日期时间',
  `created_at` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `created_by` varchar(50) NOT NULL DEFAULT '' COMMENT '添加人',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for version
-- ----------------------------
DROP TABLE IF EXISTS `app_version`;
CREATE TABLE `app_version`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'APP版本ID',
  `app_id` bigint(20) UNSIGNED NOT NULL COMMENT 'APP ID',
  `tag` varchar(10) NOT NULL DEFAULT '' COMMENT '版本tag',
  `code` varchar(10) NOT NULL DEFAULT '' COMMENT '版本号',
  `description` varchar(300) NOT NULL DEFAULT '' COMMENT '更新说明',
  `size` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '包大小',
  `hash` varchar(80) NOT NULL DEFAULT '' COMMENT 'hash',
  `path` varchar(200) NOT NULL DEFAULT '' COMMENT '文件存放路径（如果上传到云平台这里是url）',
  `is_publish` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否发布（0：未发布，1：发布）',
  `status` tinyint(1) UNSIGNED DEFAULT 1 COMMENT '状态（1：删除，2有效）',
  `updated_at` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '更新时间',
  `updated_by` varchar(50) NOT NULL DEFAULT '' COMMENT '修改人',
  `created_at` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `created_by` varchar(50) NOT NULL DEFAULT '' COMMENT '添加人',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COMMENT = 'app版本表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for members
-- ----------------------------
DROP TABLE IF EXISTS `member`;
CREATE TABLE `member`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `mobile` varchar(15) NOT NULL DEFAULT '' COMMENT '手机',
  `account` varchar(50) NOT NULL DEFAULT '' COMMENT '登录账户',
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '姓名',
  `password` varchar(80) NOT NULL DEFAULT '' COMMENT '密码',
  `avatar` varchar(255) NOT NULL DEFAULT '' COMMENT '头像',
  `email` varchar(80) NOT NULL DEFAULT '' COMMENT '邮箱',
  `token` varchar(80) NOT NULL DEFAULT '' COMMENT 'API Token',
  `status` tinyint(1) NOT NULL COMMENT '状态（1：删除，2有效）',
  `updated_at` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '更新时间',
  `updated_by` varchar(50) NOT NULL DEFAULT '' COMMENT '修改人',
  `created_at` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `created_by` varchar(50) NOT NULL DEFAULT '' COMMENT '添加人',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `uk_token` (`token`) USING BTREE,
  UNIQUE KEY `uk_email` (`email`) USING BTREE,
  UNIQUE KEY `uk_account` (`account`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for member_statistics
-- ----------------------------
DROP TABLE IF EXISTS `member_statistics`;
CREATE TABLE `member_statistics` (
  `member_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '会员ID',
  `count_team` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '参与的团队数量',
  `count_app` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '维护的app数量',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`member_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='用户的信息统计表';

-- ----------------------------
-- Table structure for team
-- ----------------------------
DROP TABLE IF EXISTS `team`;
CREATE TABLE `team`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `owner` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '团队所有人',
  `name` varchar(25) DEFAULT NULL COMMENT '团队名',
  `status` tinyint(1) UNSIGNED DEFAULT 2 COMMENT '状态（1：删除，2有效）',
  `updated_at` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '更新时间',
  `updated_by` varchar(50) NOT NULL DEFAULT '' COMMENT '修改人',
  `created_at` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `created_by` varchar(50) NOT NULL DEFAULT '' COMMENT '添加人',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COMMENT = '团队表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for team_member
-- ----------------------------
DROP TABLE IF EXISTS `team_member`;
CREATE TABLE `team_member`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `team_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '团队ID',
  `member_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '会员ID',
  `role` tinyint(3) UNSIGNED NOT NULL DEFAULT 1 COMMENT '角色（1 普通成员 2 管理员 3 所有人）',
  `updated_at` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '更新时间',
  `updated_by` varchar(50) NOT NULL DEFAULT '' COMMENT '修改人',
  `created_at` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `created_by` varchar(50) NOT NULL DEFAULT '' COMMENT '添加人',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `uk_team_member_id` (`team_id`,`member_id`) USING BTREE,
  INDEX `idx_member_id`(`member_id`) USING BTREE,
  INDEX `idx_team_id`(`team_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COMMENT = '团队成员表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for team_statistics
-- ----------------------------
DROP TABLE IF EXISTS `team_statistics`;
CREATE TABLE `team_statistics` (
  `team_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '团队ID',
  `count_member` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '团队成员数量',
  `count_app` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '维护的app数量',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`team_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='团队的信息统计表';

SET FOREIGN_KEY_CHECKS = 1;
