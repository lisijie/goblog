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

#
# Source for table tb_post
#

DROP TABLE IF EXISTS `tb_post`;
CREATE TABLE `tb_post` (
  `id` mediumint(8) unsigned NOT NULL auto_increment,
  `userid` mediumint(8) unsigned NOT NULL default '0',
  `author` varchar(15) NOT NULL default '',
  `title` varchar(100) NOT NULL default '',
  `urlname` varchar(100) NOT NULL default '',
  `content` mediumtext NOT NULL,
  `tags` varchar(100) NOT NULL default '',
  `views` mediumint(8) unsigned NOT NULL default '0',
  `status` tinyint(1) NOT NULL default '0',
  `posttime` datetime NOT NULL default '0000-00-00 00:00:00',
  `updated` datetime NOT NULL default '0000-00-00 00:00:00',
  PRIMARY KEY  (`id`),
  KEY `userid` (`userid`),
  KEY `posttime` (`posttime`),
  KEY `urlname` (`urlname`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

#
# Dumping data for table tb_post
#

INSERT INTO `tb_post` VALUES (12,0,'','测试文章','','<p>啊是大三的阿萨德撒发射点发是否艾丝凡啊是大三的阿萨德撒发射点发是否艾丝凡啊是大三的阿萨德撒发射点发是否艾丝凡啊是大三的阿萨德撒发射点发是否艾丝凡</p>','标签1,标签2',0,0,'2013-12-19 09:22:36','0000-00-00 00:00:00');
INSERT INTO `tb_post` VALUES (13,0,'','asdsdsdsdsdsdsdsd','','<p>sdsdsdsdsdsdsdsdsdsdsdsdsdsdsdsd</p>','',0,0,'2013-12-20 01:44:36','0000-00-00 00:00:00');
INSERT INTO `tb_post` VALUES (14,0,'','asdada','','<p>dasdasda</p>','',0,0,'2013-12-20 01:44:39','0000-00-00 00:00:00');
INSERT INTO `tb_post` VALUES (15,0,'','asddadas','','<p>dasdadasads</p>','',0,0,'2013-12-20 01:44:43','0000-00-00 00:00:00');
INSERT INTO `tb_post` VALUES (16,0,'','dadadasd','','','',0,0,'2013-12-20 01:44:46','0000-00-00 00:00:00');
INSERT INTO `tb_post` VALUES (17,0,'','adsdaswwwwwwwwwww','ww\' or 1=1','<p>daswwwwwwwwwwwwwww</p>','',0,0,'2013-12-19 09:44:51','2013-12-24 02:53:40');
INSERT INTO `tb_post` VALUES (18,1,'admin','asdsadasd','','<p>sadasd</p>','aaaxxx,asda1',0,0,'2013-12-24 03:12:11','2013-12-24 03:12:11');

#
# Source for table tb_tag
#

DROP TABLE IF EXISTS `tb_tag`;
CREATE TABLE `tb_tag` (
  `id` mediumint(8) unsigned NOT NULL auto_increment,
  `name` varchar(20) NOT NULL default '',
  `count` mediumint(8) unsigned NOT NULL default '0',
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

#
# Source for table tb_tag_post
#

DROP TABLE IF EXISTS `tb_tag_post`;
CREATE TABLE `tb_tag_post` (
  `id` int(11) unsigned NOT NULL auto_increment,
  `tagid` mediumint(8) unsigned NOT NULL,
  `postid` mediumint(8) unsigned NOT NULL default '0',
  PRIMARY KEY  (`id`),
  KEY `tagid` (`tagid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

#
# Dumping data for table tb_tag_post
#

INSERT INTO `tb_tag_post` VALUES (1,1,12);
INSERT INTO `tb_tag_post` VALUES (2,2,12);
INSERT INTO `tb_tag_post` VALUES (3,3,18);
INSERT INTO `tb_tag_post` VALUES (4,4,18);

#
# Source for table tb_user
#

DROP TABLE IF EXISTS `tb_user`;
CREATE TABLE `tb_user` (
  `id` mediumint(8) unsigned NOT NULL auto_increment,
  `username` varchar(15) NOT NULL default '',
  `password` varchar(32) NOT NULL default '',
  `email` varchar(50) NOT NULL default '',
  `logincount` mediumint(8) unsigned NOT NULL default '0',
  `lastip` varchar(15) NOT NULL default '0',
  `lastlogin` datetime NOT NULL default '0000-00-00 00:00:00',
  `authkey` char(10) NOT NULL default '',
  `active` tinyint(3) NOT NULL default '0',
  PRIMARY KEY  (`id`),
  UNIQUE KEY `username` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

#
# Dumping data for table tb_user
#

INSERT INTO `tb_user` VALUES (1,'admin','7fef6171469e80d32c0559f88b377245','admin@admin.com',1,'127.0.0.1','2013-12-24 06:47:10','',1);

/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
