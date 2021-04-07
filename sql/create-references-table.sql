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