SELECT * FROM papers;
SELECT * FROM `references`;

DELETE FROM papers WHERE id=3;

SELECT * FROM authors WHERE paper_id =3;

UPDATE PAPERS SET
`year`=2001, title='Structure of the Ku heterodimer bound to DNA and its implications for double-strand break repair', 
journal='Nature', volume=412, issue=2, pages='607–614', doi='10.1038/35088000'
WHERE id=1;
SELECT * FROM papers;

SELECT * FROM users;
SELECT * FROM tags;
SELECT * FROM authors;
SELECT * FROM (SELECT tag_id, count(paper_id) FROM keywords GROUP BY tag_id) test ORDER BY test.`count(paper_id)` DESC;

SELECT * FROM (SELECT user_id, count(paper_id) FROM authors
GROUP BY user_id) test
ORDER BY `count(paper_id)` DESC;

SELECT `year`, title, user_id FROM papers, (SELECT * FROM authors WHERE user_id = 22) test
WHERE papers.id = test.paper_id
ORDER BY `year` DESC;

SELECT * FROM `references` WHERE citee_id = citer_id;
DELETE FROM `references` WHERE id > 0;

UPDATE `references` SET id = id - 395 WHERE id > 0;
SELECT * FROM `references`;

-- papers that cite paper 32
SELECT papers.* 
FROM papers,  (SELECT * FROM `references` WHERE citee_id = 32) sub_query
WHERE papers.id = sub_query.citer_id;

SELECT * FROM papers where id=32;

-- references of paper 32
SELECT papers.* 
FROM papers,  (SELECT * FROM `references` WHERE citer_id = 32) sub_query
WHERE papers.id = sub_query.citee_id;

SELECT * FROM papers, `references` 
WHERE papers.id = citee_id
AND citer_id = 1;

SELECT * FROM papers ORDER BY `year` DESC;

UPDATE users SET id = id - 1 WHERE id > 0;

DELETE FROM users WHERE id > 75;
DELETE FROM papers WHERE id > 200;

SELECT first_name, last_name
FROM users, authors
WHERE paper_id = 
(SELECT id FROM papers WHERE title = "Protein Measurement With The Folin Phenol Reagent");

SELECT papers.*
FROM papers, `references`
WHERE papers.id = `references`.citer_id
AND `references`.citer_id = 7;

SELECT * FROM `sys`.sys_config;
