CREATE DATABASE minigameDB;
USE minigameDB;
DROP TABLE IF EXISTS `user_info`;
CREATE TABLE `user_info` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `user_name` varchar(64) NOT NULL,
  `account_id` varchar(64) NOT NULL,
  `password` varchar(64) NOT NULL,
  `created` datetime NOT NULL DEFAULT now(),
  `updated` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

INSERT INTO `user_info` VALUES ('1', '第一推动力', '123456789','123456', '2021-09-28', '2021-09-28');

INSERT INTO `user_info`(user_name,account_id,password,created,updated) VALUES ('第二推动力', '123123','1111111', '2021-09-28', '2021-09-28');



