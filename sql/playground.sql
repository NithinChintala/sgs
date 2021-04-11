SELECT * FROM papers;
SELECT * FROM `references`;
SELECT * FROM users;
SELECT * FROM tags;

SELECT * FROM `references` WHERE citee_id = citer_id;
DELETE FROM `references` WHERE id > 0;

UPDATE `references` SET id = id - 395 WHERE id > 0;
SELECT * FROM `references`;

SELECT * FROM `references` WHERE citee_id = 32;


SELECT * FROM papers ORDER BY `year` DESC;

UPDATE users SET id = id - 1 WHERE id > 0;

DELETE FROM users WHERE id > 75;
DELETE FROM papers WHERE id > 200;

SELECT first_name, last_name
FROM users, authors
WHERE paper_id = 
(SELECT id FROM papers WHERE title = "Protein Measurement With The Folin Phenol Reagent");

SELECT * FROM `sys`.sys_config;