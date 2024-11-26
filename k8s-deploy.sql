/*
 Navicat Premium Data Transfer

 Source Server         : localhost-3309
 Source Server Type    : MySQL
 Source Server Version : 50731
 Source Host           : localhost:3309
 Source Schema         : k8s-deploy

 Target Server Type    : MySQL
 Target Server Version : 50731
 File Encoding         : 65001

 Date: 08/11/2024 12:36:29
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for admin_auth
-- ----------------------------
DROP TABLE IF EXISTS `admin_auth`;
CREATE TABLE `admin_auth`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `pid` int(11) NULL DEFAULT NULL COMMENT '上级ID',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '节点名',
  `key` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '权限标识',
  `is_menu` int(11) NULL DEFAULT NULL COMMENT '是否是菜单栏 0：否 1：是',
  `api` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '接口',
  `action` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '操作方法',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `delete_time` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `key`(`key`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 44 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '权限' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of admin_auth
-- ----------------------------
INSERT INTO `admin_auth` VALUES (1, 0, '后台管理', 'admin', 1, '', '', NULL, NULL, NULL);
INSERT INTO `admin_auth` VALUES (2, 1, '用户管理', 'admin:user', 1, '', '', NULL, NULL, NULL);
INSERT INTO `admin_auth` VALUES (3, 2, '用户列表', 'admin:user:list', 0, '/user', 'get', NULL, NULL, NULL);
INSERT INTO `admin_auth` VALUES (4, 2, '添加用户', 'admin:user:add', 0, '/user', 'post', NULL, NULL, NULL);
INSERT INTO `admin_auth` VALUES (7, 1, '角色管理', 'admin:role', 1, '', '', NULL, NULL, NULL);
INSERT INTO `admin_auth` VALUES (8, 7, '角色列表', 'admin:role:list', 0, '/role', 'get', NULL, NULL, NULL);
INSERT INTO `admin_auth` VALUES (9, 2, '编辑用户', 'admin:user:edit', 0, '/user/:id', 'put', NULL, NULL, NULL);
INSERT INTO `admin_auth` VALUES (10, 2, '获取用户详情', 'admin:user:detail', 0, '/user/:id', 'get', NULL, NULL, NULL);
INSERT INTO `admin_auth` VALUES (11, 2, '删除用户', 'admin:user:del', 0, '/user/:id', 'delete', NULL, NULL, NULL);
INSERT INTO `admin_auth` VALUES (12, 7, '角色详情', 'admin:role:detail', 0, '/role/:id', 'get', NULL, NULL, NULL);
INSERT INTO `admin_auth` VALUES (13, 7, '添加角色', 'admin:role:add', 0, '/role', 'post', NULL, NULL, NULL);
INSERT INTO `admin_auth` VALUES (14, 7, '编辑角色', 'admin:role:edit', 0, '/role/:id', 'put', NULL, NULL, NULL);
INSERT INTO `admin_auth` VALUES (15, 7, '删除角色', 'admin:role:del', 0, '/role/:id', 'delete', NULL, NULL, NULL);
INSERT INTO `admin_auth` VALUES (16, 1, '权限管理', 'admin:auth', 1, '', '', NULL, NULL, NULL);
INSERT INTO `admin_auth` VALUES (17, 16, '权限列表', 'admin:auth:list', 0, '/auth', 'get', NULL, NULL, NULL);
INSERT INTO `admin_auth` VALUES (18, 16, '添加权限', 'admin:auth:add', 0, '/auth', 'post', NULL, NULL, NULL);
INSERT INTO `admin_auth` VALUES (19, 16, '编辑权限', 'admin:auth:edit', 0, '/auth/:id', 'put', NULL, NULL, NULL);
INSERT INTO `admin_auth` VALUES (20, 16, '删除权限', 'admin:auth:del', 0, '/auth/:id', 'delete', NULL, NULL, NULL);
INSERT INTO `admin_auth` VALUES (23, 0, '命名空间管理', 'namespace', 1, '', '', NULL, NULL, NULL);
INSERT INTO `admin_auth` VALUES (24, 23, '命名空间列表', 'namespace:list', 0, '/namespace', 'get', NULL, NULL, NULL);
INSERT INTO `admin_auth` VALUES (25, 23, '添加命名空间', 'namespace:add', 0, '/namespace', 'post', NULL, NULL, NULL);
INSERT INTO `admin_auth` VALUES (26, 23, '删除命名空间', 'namespace:del', 0, '/namespace/:id', 'delete', NULL, NULL, NULL);
INSERT INTO `admin_auth` VALUES (27, 0, '项目管理', 'project', 1, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `admin_auth` VALUES (28, 27, '项目列表', 'project:list', 0, '/project', 'get', NULL, NULL, NULL);
INSERT INTO `admin_auth` VALUES (29, 27, '添加项目', 'project:add', 0, '/project', 'post', NULL, NULL, NULL);
INSERT INTO `admin_auth` VALUES (30, 27, '编辑项目', 'project:edit', 0, '/project/:id', 'put', NULL, NULL, NULL);
INSERT INTO `admin_auth` VALUES (31, 27, '删除项目', 'project:del', 0, '/project/:id', 'delete', NULL, NULL, NULL);
INSERT INTO `admin_auth` VALUES (33, 0, '模板管理', 'template', 1, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `admin_auth` VALUES (34, 33, '模板列表', 'template:list', 0, '/template', 'get', NULL, NULL, NULL);
INSERT INTO `admin_auth` VALUES (35, 33, '添加模板', 'template:add', 0, '/template', 'post', NULL, NULL, NULL);
INSERT INTO `admin_auth` VALUES (36, 33, '编辑模板', 'template:edit', 0, '/template/:id', 'put', NULL, NULL, NULL);
INSERT INTO `admin_auth` VALUES (37, 33, '删除模板', 'template:del', 0, '/template/:id', 'delete', NULL, NULL, NULL);
INSERT INTO `admin_auth` VALUES (38, 0, '部署管理', 'deploy', 1, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `admin_auth` VALUES (39, 38, '上线单列表', 'deploy:list', 0, '/deploy', 'get', NULL, NULL, NULL);
INSERT INTO `admin_auth` VALUES (40, 38, '添加上线单', 'deploy:add', 0, '/deploy', 'post', NULL, NULL, NULL);
INSERT INTO `admin_auth` VALUES (41, 38, '上线', 'deploy:deploy', 0, '/deploy/deploy/:id', 'post', NULL, NULL, NULL);
INSERT INTO `admin_auth` VALUES (42, 38, '删除上线单', 'deploy:del', 0, '/deploy/:id', 'delete', NULL, NULL, NULL);
INSERT INTO `admin_auth` VALUES (43, 38, '获取项目详情', 'deploy:project:detail', 0, '/deploy/peoject/:id', 'get', NULL, NULL, NULL);
INSERT INTO `admin_auth` VALUES (44, 0, '调试运行', 'run', 1, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `admin_auth` VALUES (45, 44, '调试运行提交', 'run:index', 0, '/run/deploy', 'post', NULL, NULL, NULL);

-- ----------------------------
-- Table structure for admin_casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS `admin_casbin_rule`;
CREATE TABLE `admin_casbin_rule`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `ptype` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v0` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v1` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v2` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v3` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v4` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v5` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v6` varchar(25) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v7` varchar(25) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_admin_casbin_rule`(`ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 15 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = 'casbin 权限管理' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for admin_role
-- ----------------------------
DROP TABLE IF EXISTS `admin_role`;
CREATE TABLE `admin_role`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '角色名',
  `auth` json NULL COMMENT '权限ID',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `delete_time` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '角色' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of admin_role
-- ----------------------------
INSERT INTO `admin_role` VALUES (1, '超级管理员', NULL, NULL, NULL, NULL);

-- ----------------------------
-- Table structure for admin_user
-- ----------------------------
DROP TABLE IF EXISTS `admin_user`;
CREATE TABLE `admin_user`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '用户名',
  `realname` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '真实姓名',
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '密码',
  `phone` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '手机号',
  `role_id` int(11) NULL DEFAULT NULL COMMENT '角色ID',
  `status` int(11) NULL DEFAULT 0 COMMENT '状态 0:未启用 1:正常',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `delete_time` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '后台用户' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of admin_user
-- ----------------------------
INSERT INTO `admin_user` VALUES (1, 'admin', 'admin', '$2a$10$jlxETsTS1zLajnxqhojtIuMdGHGvEX5vKlLHzKx4LLx4Qj2vujMKq', '10086', 1, 1, NULL, NULL, NULL);

-- ----------------------------
-- Table structure for deploy
-- ----------------------------
DROP TABLE IF EXISTS `deploy`;
CREATE TABLE `deploy`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '部署标题',
  `project_id` int(11) NULL DEFAULT NULL COMMENT '项目ID',
  `project` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '项目',
  `fingerprint` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '模板指纹',
  `template_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '模板名',
  `template_content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '模板原始值',
  `template_parse` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '模板解析值',
  `params` json NULL COMMENT '模板变量',
  `status` tinyint(4) NULL DEFAULT 0 COMMENT '状态 0：等待上线 1：上线中 2：上线成功 3：上线失败',
  `create_time` datetime NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `delete_time` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of deploy
-- ----------------------------

-- ----------------------------
-- Table structure for deploy_log
-- ----------------------------
DROP TABLE IF EXISTS `deploy_log`;
CREATE TABLE `deploy_log`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `pid` int(11) NULL DEFAULT NULL COMMENT '部署ID',
  `type` tinyint(4) NULL DEFAULT NULL COMMENT '消息登记 0：默认 1：成功 2：警告 3：错误',
  `message` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '消息',
  `create_time` datetime NULL DEFAULT CURRENT_TIMESTAMP COMMENT '时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '部署日志' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of deploy_log
-- ----------------------------

-- ----------------------------
-- Table structure for k8s_template
-- ----------------------------
DROP TABLE IF EXISTS `k8s_template`;
CREATE TABLE `k8s_template`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '模板名',
  `desc` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '模板描述',
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '模板内容',
  `create_time` datetime NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `delete_time` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '模板管理' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of k8s_template
-- ----------------------------

-- ----------------------------
-- Table structure for project
-- ----------------------------
DROP TABLE IF EXISTS `project`;
CREATE TABLE `project`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '项目名',
  `desc` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '项目描述',
  `git` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'git地址',
  `user_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '用户名',
  `token` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '令牌/密码',
  `use_tag` tinyint(4) NULL DEFAULT 0 COMMENT '是否使用 TAG',
  `params` json NULL COMMENT '预设变量',
  `template` json NULL COMMENT '本项目模板',
  `create_time` datetime NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `delete_time` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '项目表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of project
-- ----------------------------

SET FOREIGN_KEY_CHECKS = 1;
