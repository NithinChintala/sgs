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
