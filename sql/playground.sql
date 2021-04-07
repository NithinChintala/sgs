SELECT * FROM papers;
SELECT * FROM `references`;
SELECT * FROM users;

SELECT first_name, last_name
FROM users, authors
WHERE paper_id = 
(SELECT id FROM papers WHERE title = "Protein Measurement With The Folin Phenol Reagent");

SELECT * FROM `sys`.sys_config;