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