# MySQL-Front 5.1  (Build 4.2)

/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE */;
/*!40101 SET SQL_MODE='STRICT_TRANS_TABLES,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES */;
/*!40103 SET SQL_NOTES='ON' */;


# Host: 127.0.0.1    Database: goblog
# ------------------------------------------------------
# Server version 5.0.41-community-nt

DROP DATABASE IF EXISTS `goblog`;
CREATE DATABASE `goblog` /*!40100 DEFAULT CHARACTER SET utf8 */;
USE `goblog`;

#
# Source for table tb_option
#

DROP TABLE IF EXISTS `tb_option`;
CREATE TABLE `tb_option` (
  `id` mediumint(8) unsigned NOT NULL auto_increment,
  `name` varchar(30) NOT NULL default '',
  `value` text NOT NULL,
  PRIMARY KEY  (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

#
# Dumping data for table tb_option
#

INSERT INTO `tb_option` VALUES (1,'sitename','测试网站');
INSERT INTO `tb_option` VALUES (2,'siteurl','http://www.lisijie.org');
INSERT INTO `tb_option` VALUES (3,'subtitle','adad');
INSERT INTO `tb_option` VALUES (4,'pagesize','ada');
INSERT INTO `tb_option` VALUES (5,'keywords','dad');
INSERT INTO `tb_option` VALUES (6,'description','sdasda');
INSERT INTO `tb_option` VALUES (7,'email','lisijie86@gmail.com');
INSERT INTO `tb_option` VALUES (8,'theme','default');

#
# Source for table tb_post
#

DROP TABLE IF EXISTS `tb_post`;
CREATE TABLE `tb_post` (
  `id` mediumint(8) unsigned NOT NULL auto_increment,
  `userid` mediumint(8) unsigned NOT NULL default '0' COMMENT '用户id',
  `author` varchar(15) NOT NULL default '' COMMENT '作者',
  `title` varchar(100) NOT NULL default '' COMMENT '标题',
  `color` varchar(7) NOT NULL default '' COMMENT '标题颜色',
  `urlname` varchar(100) NOT NULL default '' COMMENT 'url名',
  `urltype` tinyint(3) NOT NULL default '0' COMMENT 'url访问形式',
  `excerpt` mediumtext NOT NULL COMMENT '内容摘要',
  `content` mediumtext NOT NULL COMMENT '内容',
  `tags` varchar(100) NOT NULL default '' COMMENT '标签',
  `views` mediumint(8) unsigned NOT NULL default '0' COMMENT '查看次数',
  `status` tinyint(1) NOT NULL default '0' COMMENT '状态{0:正常,1:草稿,2:回收站}',
  `posttime` datetime NOT NULL default '0000-00-00 00:00:00' COMMENT '发布时间',
  `updated` datetime NOT NULL default '0000-00-00 00:00:00' COMMENT '更新时间',
  `istop` tinyint(3) NOT NULL default '0' COMMENT '是否置顶',
  PRIMARY KEY  (`id`),
  KEY `userid` (`userid`),
  KEY `posttime` (`posttime`),
  KEY `urlname` (`urlname`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

#
# Dumping data for table tb_post
#

INSERT INTO `tb_post` VALUES (12,0,'','测试文章','','',0,'','<p>啊是大三的阿萨德撒发射点发是否艾丝凡啊是大三的阿萨德撒发射点发是否艾丝凡啊是大三的阿萨德撒发射点发是否艾丝凡啊是大三的阿萨德撒发射点发是否艾丝凡</p>','标签1,标签2',0,0,'2013-12-19 09:22:36','0000-00-00 00:00:00',0);
INSERT INTO `tb_post` VALUES (14,0,'','asdada','','',0,'','<p>dasdasda</p>','',0,0,'2013-12-20 01:44:39','0000-00-00 00:00:00',0);
INSERT INTO `tb_post` VALUES (15,0,'','asddadas','','',0,'','<p>dasdadasads</p>','',0,0,'2013-12-20 01:44:43','0000-00-00 00:00:00',0);
INSERT INTO `tb_post` VALUES (16,0,'','dadadasd','','',0,'','','',0,0,'2013-12-20 01:44:46','0000-00-00 00:00:00',0);
INSERT INTO `tb_post` VALUES (18,1,'admin','asdsadasd','','',1,'','<p>sadasd</p>','aaaxxx,asda1',0,0,'2013-12-23 19:12:11','2013-12-25 05:15:34',0);
INSERT INTO `tb_post` VALUES (19,1,'admin','asadads','#CC0000','das',2,'','<p>dasssss</p>','asd',0,0,'2013-12-24 20:05:25','2013-12-25 04:53:05',1);

#
# Source for table tb_tag
#

DROP TABLE IF EXISTS `tb_tag`;
CREATE TABLE `tb_tag` (
  `id` mediumint(8) unsigned NOT NULL auto_increment,
  `name` varchar(20) NOT NULL default '' COMMENT '标签名',
  `count` mediumint(8) unsigned NOT NULL default '0' COMMENT '使用次数',
  PRIMARY KEY  (`id`),
  KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

#
# Dumping data for table tb_tag
#

INSERT INTO `tb_tag` VALUES (1,'标签1',1);
INSERT INTO `tb_tag` VALUES (2,'标签2',1);
INSERT INTO `tb_tag` VALUES (3,'aaaxxx',1);
INSERT INTO `tb_tag` VALUES (4,'asda1',1);
INSERT INTO `tb_tag` VALUES (5,'asd',1);

#
# Source for table tb_tag_post
#

DROP TABLE IF EXISTS `tb_tag_post`;
CREATE TABLE `tb_tag_post` (
  `id` int(11) unsigned NOT NULL auto_increment,
  `tagid` mediumint(8) unsigned NOT NULL default '0' COMMENT '标签id',
  `postid` mediumint(8) unsigned NOT NULL default '0' COMMENT '内容id',
  `poststatus` tinyint(3) NOT NULL default '0' COMMENT '内容状态',
  `posttime` datetime NOT NULL default '0000-00-00 00:00:00' COMMENT '发布时间',
  PRIMARY KEY  (`id`),
  KEY `tagid` (`tagid`),
  KEY `postid` (`postid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

#
# Dumping data for table tb_tag_post
#

INSERT INTO `tb_tag_post` VALUES (1,1,12,0,'0000-00-00 00:00:00');
INSERT INTO `tb_tag_post` VALUES (2,2,12,0,'0000-00-00 00:00:00');
INSERT INTO `tb_tag_post` VALUES (6,5,19,0,'0000-00-00 00:00:00');
INSERT INTO `tb_tag_post` VALUES (7,3,18,0,'0000-00-00 00:00:00');
INSERT INTO `tb_tag_post` VALUES (8,4,18,0,'0000-00-00 00:00:00');

#
# Source for table tb_user
#

DROP TABLE IF EXISTS `tb_user`;
CREATE TABLE `tb_user` (
  `id` mediumint(8) unsigned NOT NULL auto_increment,
  `username` varchar(15) NOT NULL default '' COMMENT '用户名',
  `password` varchar(32) NOT NULL default '' COMMENT '密码',
  `email` varchar(50) NOT NULL default '' COMMENT '邮箱',
  `logincount` mediumint(8) unsigned NOT NULL default '0' COMMENT '登录次数',
  `lastip` varchar(15) NOT NULL default '0' COMMENT '最后登录ip',
  `lastlogin` datetime NOT NULL default '0000-00-00 00:00:00' COMMENT '最后登录时间',
  `authkey` char(10) NOT NULL default '' COMMENT '登录key',
  `active` tinyint(3) NOT NULL default '0' COMMENT '是否激活',
  PRIMARY KEY  (`id`),
  UNIQUE KEY `username` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

#
# Dumping data for table tb_user
#

INSERT INTO `tb_user` VALUES (1,'admin','7fef6171469e80d32c0559f88b377245','admin@admin.com',4,'[','2013-12-25 05:48:17','',1);

/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
