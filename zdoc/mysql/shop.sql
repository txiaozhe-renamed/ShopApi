--
CREATE DATABASE IF NOT EXISTS `shop`;
USE `shop`;

-- ----------------------------------------------------------

CREATE TABLE IF NOT EXISTS `carts` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `productid` int(11) unsigned NOT NULL,
  `name` varchar(200) NOT NULL,
  `count` int(11) unsigned NOT NULL,
  `size`    varchar(50) DEFAULT '0',
  `color`   varchar(50) DEFAULT '0',
  `imageid` int(11) unsigned NOT NULL,
  `userid` int(11) NOT NULL,
  `status`  int(11) NOT NULL,
  `created` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1000 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ----------------------------------------------------------

CREATE TABLE IF NOT EXISTS `categories` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(200) NOT NULL DEFAULT '',
  `pid` int(11) NOT NULL DEFAULT '0',
  `status` int(11) NOT NULL,
  `remark` varchar(1000) DEFAULT NULL,
  `created` datetime NOT NULL DEFAULT current_timestamp,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1000 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ----------------------------------------------------------

CREATE TABLE IF NOT EXISTS `contact` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL DEFAULT '',
  `userid` int(11) NOT NULL,
  `phone` varchar(20) DEFAULT '',
  `province` varchar(100) NOT NULL,
  `city` varchar(100) NOT NULL,
  `street` varchar(100) NOT NULL,
  `address` varchar(200) NOT NULL DEFAULT '',
  `created` datetime NOT NULL DEFAULT current_timestamp,
  `isdefault` TINYINT(1) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1000 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ----------------------------------------------------------


CREATE TABLE IF NOT EXISTS `images` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `url` varchar(200) NOT NULL DEFAULT '',
  `image` varchar(200) NOT NULL,
  `type` int(11) NOT NULL,
  `title` varchar(100) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1000 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ----------------------------------------------------------


CREATE TABLE IF NOT EXISTS `orders` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `userid` int(11) unsigned NOT NULL,
  `totalprice` double NOT NULL,
  `payment` double NOT NULL,
  `freight` double DEFAULT '0' COMMENT '运费',
  `remark` text COMMENT '备注',
  `discount` int(11) DEFAULT '0',
  `size`    varchar(50) DEFAULT '0',
  `color`   varchar(50) DEFAULT '0',
  `status` int(11) NOT NULL,
  `created` datetime NOT NULL DEFAULT current_timestamp,
  `payway` INT  NOT NULL ,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1000 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ----------------------------------------------------------


CREATE TABLE IF NOT EXISTS `products` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(200) NOT NULL DEFAULT '',
  `totalsale` INT(10) NOT NULL DEFAULT '0',
  `category` int(11) NOT NULL,
  `price` double NOT NULL,
  `originalprice` double NOT NULL,
  `status` int(11) NOT NULL,
  `size`  varchar(200),
  `color` varchar(200),
  `imageid` int(11) unsigned NOT NULL COMMENT '商品封面图片',
  `imageids` varchar(200) NOT NULL DEFAULT '' COMMENT '商品图片集',
  `remark` varchar(1000) DEFAULT '',
  `detail` longtext NOT NULL,
  `created` datetime NOT NULL DEFAULT current_timestamp,
  `inventory` int(11) unsigned NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1000 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ----------------------------------------------------------


CREATE TABLE IF NOT EXISTS `users` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `openid` text,
  `name` varchar(100) DEFAULT NULL UNIQUE ,
  `password` varchar(128) NOT NULL DEFAULT '',
  `status` int(11) DEFAULT NULL,
  `type` INT(11)  NOT NULL,
  `created` datetime NOT NULL DEFAULT current_timestamp,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1000 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ----------------------------------------------------------


CREATE TABLE IF NOT EXISTS `userinfo` (
  `userid`   INT(11),
  `avatar` text,
  `nickname` VARCHAR(100)         DEFAULT NULL,
  `email`    VARCHAR(100)         DEFAULT NULL,
  `phone`    VARCHAR(20) NOT NULL DEFAULT '',
  `sex`      TINYINT(1)           DEFAULT NULL COMMENT '0:男;1:女',
  PRIMARY KEY (`userid`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
