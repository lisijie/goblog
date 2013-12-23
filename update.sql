ALTER TABLE  `tb_post` CHANGE  `alias`  `urlname` VARCHAR( 200 ) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT  '';

ALTER TABLE  `tb_post` ADD INDEX (  `posttime` );

ALTER TABLE  `tb_user` ADD  `authkey` CHAR( 10 ) NOT NULL DEFAULT  '';

ALTER TABLE  `tb_user` DROP  `lastlogin`;

ALTER TABLE  `tb_user` ADD  `lastlogin` DATETIME NOT NULL AFTER  `lastip`;

ALTER TABLE  `tb_post` ADD  `updated` DATETIME NOT NULL DEFAULT  '0000-00-00 00:00:00'