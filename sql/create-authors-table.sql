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