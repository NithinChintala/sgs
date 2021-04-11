-- Creates the tags table
DROP TABLE IF EXISTS `sgs`.`tags`;
CREATE TABLE `sgs`.`tags` (
	`id`          INT NOT NULL AUTO_INCREMENT,
    `word`        VARCHAR(45) NOT NULL,
    `searches`    INT NOT NULL DEFAULT 0,
    `last_search` DATETIME NULL,
    PRIMARY KEY (`id`)
);
