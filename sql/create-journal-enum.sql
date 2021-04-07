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
