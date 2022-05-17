/*
MySQL Backup
Source Server Version: 5.7.23
Source Database: issue
Date: 2020/11/8 11:53:52
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
--  Table structure for `article`
-- ----------------------------
DROP TABLE IF EXISTS `article`;
CREATE TABLE `article` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(255) NOT NULL DEFAULT '',
  `keywords` varchar(255) NOT NULL,
  `category` varchar(255) NOT NULL,
  `content` longtext NOT NULL,
  `status` tinyint(3) unsigned NOT NULL DEFAULT '1',
  `created_at` datetime NOT NULL,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=38 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
--  Table structure for `article_log`
-- ----------------------------
DROP TABLE IF EXISTS `article_log`;
CREATE TABLE `article_log` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
	`article_id` int(10) UNSIGNED NOT NULL,
	`mark` VARCHAR(20), 
  `created_at` datetime NOT NULL,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
--  Table structure for `category`
-- ----------------------------
DROP TABLE IF EXISTS `category`;
CREATE TABLE `category` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(10) NOT NULL DEFAULT '',
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
--  Table structure for `groups`
-- ----------------------------
DROP TABLE IF EXISTS `groups`;
CREATE TABLE `groups` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `role_id` int(11) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_user_group` (`user_id`,`role_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=51 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
--  Table structure for `job_count`
-- ----------------------------
DROP TABLE IF EXISTS `job_count`;
CREATE TABLE `job_count` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `job_title` varchar(30) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '职位名称，开发语言',
  `region` varchar(10) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '上海' COMMENT '地区',
  `amount` int(11) NOT NULL DEFAULT '0' COMMENT '职位数',
  `created_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_title` (`job_title`),
  KEY `idx_region` (`region`)
) ENGINE=InnoDB AUTO_INCREMENT=959 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
--  Table structure for `log`
-- ----------------------------
DROP TABLE IF EXISTS `log`;
CREATE TABLE `log` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `mark` varchar(255) DEFAULT NULL COMMENT '日志说明',
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=983 DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `picture`
-- ----------------------------
DROP TABLE IF EXISTS `picture`;
CREATE TABLE `picture` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `pic_url` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `status` tinyint(1) NOT NULL DEFAULT '1',
  `tag` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT '',
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
--  Table structure for `role`
-- ----------------------------
DROP TABLE IF EXISTS `role`;
CREATE TABLE `role` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `groups` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '权限列表',
  `description` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '说明',
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_id` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='可以给一个角色很多权限，也可以通过很多角色组合来拥有很多权限';

-- ----------------------------
--  Table structure for `user`
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `mobile` varchar(255) NOT NULL DEFAULT '',
  `user_name` varchar(255) NOT NULL DEFAULT '',
  `password` varchar(255) NOT NULL DEFAULT '',
  `gender` varchar(255) NOT NULL DEFAULT '',
  `email` varchar(255) NOT NULL DEFAULT '',
  `addr` varchar(255) NOT NULL DEFAULT '',
  `status` int(11) NOT NULL DEFAULT '1',
  `description` varchar(255) NOT NULL DEFAULT '',
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_mobile_uniq` (`mobile`) USING BTREE,
  UNIQUE KEY `idx_id` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=33 DEFAULT CHARSET=utf8;

-- ----------------------------
--  Records 
-- ----------------------------

INSERT INTO `groups` VALUES ('6','2','3','2018-04-28 17:35:30','2018-04-28 17:35:32'), ('7','2','1','2018-04-28 13:16:33','2018-04-28 13:16:36'), ('8','2','2','2018-06-11 13:10:12','2018-06-11 13:10:15'), ('9','2','4','2018-06-11 13:10:25','2018-06-11 13:10:28'), ('10','2','5','2018-06-11 13:10:38','2018-06-11 13:10:40'), ('11','2','6','2018-06-11 13:10:50','2018-06-11 13:10:55'), ('12','2','7','2019-02-15 15:25:02','2019-02-15 15:25:02'), ('15','2','8','2019-02-16 14:21:49','2019-02-16 14:21:49'), ('16','1','1','2019-05-08 14:58:10','2019-05-08 14:58:10'), ('17','25','2','2019-05-08 14:58:37','2019-05-08 14:58:37'), ('20','23','6','2019-05-19 23:10:19','2019-05-19 23:10:19'), ('21','13','7','2019-06-28 14:40:07','2019-06-28 14:40:07'), ('22','13','6','2019-06-28 14:40:07','2019-06-28 14:40:07'), ('23','14','8','2019-07-28 19:22:27','2019-07-28 19:22:27'), ('24','14','7','2019-07-28 19:22:27','2019-07-28 19:22:27'), ('25','14','6','2019-07-28 19:22:27','2019-07-28 19:22:27'), ('26','1','8','2019-08-16 10:08:42','2019-08-16 10:08:42'), ('27','1','7','2019-08-16 10:08:42','2019-08-16 10:08:42'), ('29','5','8','2019-08-16 10:08:59','2019-08-16 10:08:59'), ('30','5','3','2019-08-16 10:08:59','2019-08-16 10:08:59'), ('33','8','8','2019-08-16 10:11:18','2019-08-16 10:11:18'), ('34','8','3','2019-08-16 10:11:18','2019-08-16 10:11:18'), ('36','28','8','2019-09-03 14:50:28','2019-09-03 14:50:28'), ('43','13','8','2019-09-17 10:40:41','2019-09-17 10:40:41');
INSERT INTO `helpers` VALUES ('1','张三','18734587454','zhangsan@qq.com',NULL,'教学工作站,销售工作站','2017-10-24 23:11:10','2017-10-24 23:11:14'), ('2','李四','12222222222','12222222222@163.com',NULL,'财务工作站,用户管理','2017-11-18 15:30:06','2017-11-18 15:30:09');

INSERT INTO `picture` VALUES ('1','//static.cnodejs.org/Ft685Ah4vM0Z3QLB_Kht2YnTDNp9','0','Microtask,Macrotask','2019-03-20 10:06:55','2019-03-20 10:06:55'), ('2','http://dockone.io/uploads/article/20150913/e31fd491a376d3398fd4ca2dfcc98a9c.png','1','','2019-05-15 16:42:35','2019-05-15 16:42:35'), ('3','','0','','2019-05-19 23:17:32','2019-05-19 23:17:32'), ('4','','0','','2019-05-19 23:18:35','2019-05-19 23:18:35'), ('5','','0','','2019-09-30 14:38:08','2019-09-30 14:38:08'), ('6','','0','','2019-10-24 13:39:29','2019-10-24 13:39:29'), ('7','','0','','2019-10-24 13:40:07','2019-10-24 13:40:07'), ('8','','0','','2019-10-24 13:40:09','2019-10-24 13:40:09'), ('9','','0',',,','2019-10-24 13:40:48','2019-10-24 13:40:48'), ('10','','0',',','2019-10-24 13:43:25','2019-10-24 13:43:25'), ('11','','0','','2019-11-13 18:07:05','2019-11-13 18:07:05'), ('12','','0','','2019-12-14 17:06:59','2019-12-14 17:06:59'), ('13','http://5b0988e595225.cdn.sohucs.com/images/20190807/a1061be84397443ea273736378690cdd.png','1','持续集成（CI）,持续交付和持续部署（CD）,CI/CD的优势','2020-03-25 16:33:12','2020-03-25 16:33:12');
INSERT INTO `role` VALUES ('1','AdminController:UserList,AdminController:UserListRoute','管理员','2018-04-21 14:00:18','2018-04-21 14:00:23'), ('2','AdminController:UserListRoute','普通管理员','2018-04-21 14:00:15','2018-04-21 14:00:21'), ('3','AdminController:POST','添加用户','2018-04-20 13:29:49','2018-04-20 13:29:52'), ('4','AdminController:PUT','编辑用户','2018-06-11 13:05:02','2018-06-11 13:05:05'), ('5','AdminController:DeleteUser','删除用户','2018-06-11 13:07:02','2018-06-11 13:07:07'), ('6','AdminController:UserList','用户列表','2018-06-11 13:09:09','2018-06-11 13:09:12'), ('7','AdminController:ArticleEdit','编辑文章','2019-02-15 13:37:51','2019-02-15 13:37:55'), ('8','AdminController:ArticleDelete','删除文章','2019-02-16 14:21:04','2019-02-16 14:21:07');
INSERT INTO `user` VALUES ('1','13477889900','犬夜叉33','960232f4a37f948b480a3f8a5512c6f8','1','13477889900@139.com','日暮神社','0','半妖','2018-03-17 20:46:31','2020-04-21 13:38:10'), ('2','18701897513','戈薇','abc72b24857be42850f67d3160f8710e','1','18701897513@139.com','日暮神社','1','博主我感觉你好牛批','2018-03-17 20:49:44','2019-10-21 17:11:58'), ('5','18701893513','桔梗','abc72b24857be42850f67d3160f8710e','0','18611118146@139.com','看见的任何司空见惯和','0','奈落都害怕的女人','2017-07-27 03:25:01','2019-08-16 09:41:42'), ('8','10701897527','弥勒','8fa2952fff72d92c98f9f43e46dfc6bd','0','huo@gmail.com','吉林大街好地方','0','而喝了酒而温柔你感觉','2017-07-27 09:00:43','2018-03-19 11:10:50'), ('9','10706597527','七宝1','8fa2952fff72d92c98f9f43e46dfc6bd','0','438473@qq.com','发顺丰','0','收到了架构过人家饿啊人工','2017-07-29 10:38:06','2019-09-02 15:03:46'), ('13','18701497527','杀生丸','8fa2952fff72d92c98f9f43e46dfc6bd','1','hp@sina.com','送就送山东黄金人数','1','视频国际投行饿哦日后我如何进入','2018-02-05 04:20:37','2018-12-02 20:38:37'), ('14','12345678909','珊瑚','5335412ee0f17806e1017e607149336a','0','ligh@163.com','很快就都大佛开盘后具体要','1','我感觉哦过仁和堂撒今天','2018-02-05 08:07:37','2018-04-06 17:03:04'), ('21','18721897527','奈落','8fa2952fff72d92c98f9f43e46dfc6bd','1','18611118146@139.com','送就送山东黄金人数几乎是丢改好','1','就fdfda','2018-02-11 07:53:18','2018-10-24 11:43:07'), ('22','18766464985','云母','442ba06a1ac9ad299865c11234b9c492','0','ligh@163.com','送就送山东黄金人数几乎是丢改好看机会','1','','2018-02-11 15:58:10','2019-02-24 21:01:36'), ('23','12335678909','邪见','5335412ee0f17806e1017e607149336a','1','13344442929@163.com','看见的任何','0','','2018-02-11 16:05:53','2019-02-24 21:01:51'), ('25','18111897528','铃','033554527363fed57bacfcab7c77c5fb','0','lig@163.com','dfpkgipniu','0','当人看了韩国人都','2018-02-28 10:34:39','2019-02-24 21:02:23'), ('26','18765464985','神乐','442ba06a1ac9ad299865c11234b9c492','0','13344442929@163.com','很快就都大佛开盘后具体要','0','客人很多事让她二炮还叫人','2018-02-28 10:34:58','2019-02-24 21:02:58'), ('27','18701817525','神无','67c70763ce38919105acc783fa5e834d','0','18765464985','了PDF你好逗','1','哦【但是若干年后天赋','2018-02-28 10:35:20','2019-02-24 21:03:14'), ('28','15800699208','梦幻之白夜','dfda39bb37573e74338338642162d85b','0','8974@qq.com','工时费','0','好回家哦个积极破解','2018-04-05 23:19:33','2019-02-24 21:03:53'), ('29','11111111111','admin111','d384ece233ea14a88e887d2e0b363753','0','111111111','1111111111111','0','11111111111111','2019-07-25 12:58:58','2019-07-25 12:58:58'), ('31','1231312313','1111','14a5f308f01e71018e9531b2f43dbeb8','0','123123@qq.com','啊手动阀','0','阿迪斯','2019-09-15 12:41:15','2019-09-15 12:41:15'), ('32','13927459802','admin','0838bbf98cc735eda3cab04d098387be','1','330490409@qq.com','','0','','2020-01-27 14:48:44','2020-01-27 14:48:44');
