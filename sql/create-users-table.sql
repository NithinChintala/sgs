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
