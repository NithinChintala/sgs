DROP DATABASE IF EXISTS `sgs`;
CREATE SCHEMA `sgs`;

-- Create the jounral portable enum
DROP TABLE IF EXISTS `sgs`.`journal`;
CREATE TABLE `sgs`.`journal` (
	`title` VARCHAR(45) NOT NULL DEFAULT "None",
    PRIMARY KEY (`title`)
);

INSERT INTO `sgs`.`journal`(title) VALUES("Nature");
INSERT INTO `sgs`.`journal`(title) VALUES("Science");
INSERT INTO `sgs`.`journal`(title) VALUES("Cell");
INSERT INTO `sgs`.`journal`(title) VALUES("The Journal of Biological Chemistry");

-- Create the users table
DROP TABLE IF EXISTS `sgs`.`users`;
CREATE TABLE `sgs`.`users` (
	`id`            INT NOT NULL AUTO_INCREMENT,
    `first_name`    VARCHAR(45) NOT NULL,
    `last_name`     VARCHAR(45) NOT NULL,
    `username`      VARCHAR(45) NOT NULL,
    `password`      VARCHAR(45) NOT NULL,
    `email`         VARCHAR(45) NOT NULL,
    `date_of_birth` DATE NULL,
    PRIMARY KEY (`id`)
);

-- Create the papers table
DROP TABLE IF EXISTS `sgs`.`papers`;
CREATE TABLE `sgs`.`papers` (
	`id`      INT NOT NULL AUTO_INCREMENT,
    `year`    YEAR NOT NULL,
    `title`   VARCHAR(100) NOT NULL,
    `journal` VARCHAR(45) DEFAULT NULL,
    `volume`  INT NULL,
    `issue`   INT NULL,
    `pages`   VARCHAR(20) NULL,
    `doi`     VARCHAR(100) NULL,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`journal`) REFERENCES `sgs`.`journal` (`title`) 
);

-- Creates the tags table
DROP TABLE IF EXISTS `sgs`.`tags`;
CREATE TABLE `sgs`.`tags` (
	`id`          INT NOT NULL AUTO_INCREMENT,
    `word`        VARCHAR(45) NOT NULL,
    `searches`    INT NOT NULL DEFAULT 0,
    `last_search` DATETIME NULL,
    PRIMARY KEY (`id`)
);

-- Create the references table
DROP TABLE IF EXISTS `sgs`.`references`;
CREATE TABLE `sgs`.`references` (
	`id`        INT NOT NULL AUTO_INCREMENT,
    `citer_id`  INT NOT NULL,
    `citee_id`   INT NOT NULL,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`citer_id`) REFERENCES `sgs`.`papers` (`id`),
    FOREIGN KEY (`citee_id`) REFERENCES `sgs`.`papers` (`id`)
);

-- Create the authors table
DROP TABLE IF EXISTS `sgs`.`authors`;
CREATE TABLE `sgs`.`authors` (
	`id`       INT NOT NULL AUTO_INCREMENT,
    `user_id`  INT NOT NULL,
    `paper_id` INT NOT NULL,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`user_id`) REFERENCES `sgs`.`users` (`id`),
    FOREIGN KEY (`paper_id`) REFERENCES `sgs`.`papers` (`id`)
);

-- Create the keywords table
DROP TABLE IF EXISTS `sgs`.`keywords`;
CREATE TABLE `sgs`.`keywords` (
	`id`       INT NOT NULL AUTO_INCREMENT,
    `paper_id` INT NOT NULL,
    `tag_id`   INT NOT NULL,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`paper_id`) REFERENCES `sgs`.`papers` (`id`),
    FOREIGN KEY (`tag_id`)   REFERENCES `sgs`.`tags` (`id`)
);
