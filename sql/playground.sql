SELECT * FROM papers;
SELECT * FROM `references`;
SELECT * FROM users;
SELECT * FROM tags;

UPDATE users SET id = id - 1 WHERE id > 0;

DELETE FROM users WHERE id = 1;

SELECT first_name, last_name
FROM users, authors
WHERE paper_id = 
(SELECT id FROM papers WHERE title = "Protein Measurement With The Folin Phenol Reagent");

SELECT * FROM `sys`.sys_config;