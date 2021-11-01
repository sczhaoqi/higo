-- 数据库表结构
CREATE TABLE IF NOT EXISTS test_user (
    `id` INT(3) NOT NULL AUTO_INCREMENT,
    `name` varchar(20) NOT NULL DEFAULT '',
    `age` smallint(3) unsigned NOT NULL DEFAULT '0',
    `gender` tinyint(1) unsigned NOT NULL DEFAULT '0',
    PRIMARY KEY (`id`)
    )
    ENGINE=InnoDB
    DEFAULT CHARACTER SET = utf8;