ALTER TABLE  `tb_post` CHANGE  `alias`  `urlname` VARCHAR( 200 ) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT  ''

ALTER TABLE  `tb_post` ADD INDEX (  `posttime` )

ALTER TABLE  `tb_user` ADD  `authkey` CHAR( 10 ) NOT NULL DEFAULT  ''