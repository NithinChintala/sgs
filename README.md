# Scuffed Google Scholar
Scuffed Google Scholar (SGS) is an attempt to create a small version
of an application that indexes academic journals. It includes tables
to store papers, authors and tags.

## Team
I worked alone for this project. My group on Canvas is `Database Design 23`,
i'm in `CS3200 Section 04`, and my name is `Nithin Chintala`.

# Description
Please view the [uml](https://github.com/NithinChintala/sgs/blob/main/database/db_design_final_project_UML.pdf)

SGS attempts to solve the problem of having a local database of paper to search through academic papers. There
is a plethura of academic work online, but same may want to search through a manually currated set that can
be created and maintained.

SGS implements a simple relational database that can store relation between the core object: papers. SGS
provides a simple UI to insert papers and view their relations with one another.

The typical user for SGS would be someone in academia looking to having a personal copy of papers they
need for their research. Other users simply looking to research papers are also a prime target user.

The two core domain objects are a `paper` and `tag`. A `papaer` describes an academic work that may
be published in a Journal, have multiple references, have multiple tags and have multiple authors.
A tag is simply a descriptor. It can be anything really, it simply described how a group of `papers`
can be related. So the tag `proteins` may be used on papers about protein science.

## Problem Statements / Solution
1. What are the most cited papers in the database?
```sql
CREATE VIEW most_cited AS
SELECT * 
FROM (SELECT citee_id, count(citer_id) AS num_cited FROM `references` GROUP BY citee_id) subquery
ORDER BY subquery.num_cited DESC;
```

2. What papers cited this paper in their work?
```sql
-- citee_id = <id> is refering to "this paper"
SELECT papers.*
FROM papers, `references`
WHERE papers.id = subquery.citer_id
AND `references`.citee_id = 5;
```

3. What are popular papers within a certain tag?
```sql
-- keywords.tag_id = <id> refers to a "certain tag"
SELECT *
FROM most_cited, keywords
WHERE most_cited.citee_id = keywords.paper_id
AND keywords.tag_id = 4
```

## User Data Model
A user represents someone who may write papers. A user stores the `first name`, `last name`, `username`, 
`password`, `email` and `data of birth`. The `email` was added because it is something necessary
when contacting acadaemics.

## Domain Objects
There are two domain object, a `paper` and a `tag`. 

A `paper` represents an academic writing. It contains information to uniquely identify an academic
writing and optionally provide a doi.

A `tag` represents categories or other broad ideas that can group multiple papers togethr. 

## User to Domain Relationship
The `users` and `papers` table are related to one another in a many to many relationship.
That is, one `paper` can have many `authors`, and one `author` can have many `papers`. This is
reflective of how actual academic relationships are as well.

Since MySQL was used to implement this, the relationship is reified in an `authors` table.
Whenever a `paper` or `author` is deleted all of its respective mappings in the `authors` table
are also deleted.

## Domain to Domain Relationship
A `paper` can be cited by any other `paper`. This relationship is best understood with the idea 
of a `citeer` and `citee`. A `citer` cites a `citee`.  That is, if `paper` A used `paper` B in 
its finding, then A would be a `citer` citing `paper` B, the `citee`. To implement this in sql,
the `papers` table has a reified `references` table to describe the self-referential many-to-many
relationship. Deleting a `paper` deletes all the records in the mapping table that refers to the
deleted `paper`.

Finally, the `papers` and `tags` tables have a many-to-many relationship. One `paper` can have
many `tags`, and one `tag` can have many `papers`. This is reified in the `keywords` table. Deleting
a `paper` or a `tag` removes all the mapping records that relate to it.

## Portable Enum
The portable enum is the `journal` table that is used
in the `papers` table. A journal is simply an official place where `papers` are published.
Currently the available `journals` in the databse are: "Cell", "Nature" and "Science". A 
`journal` can be `NULL` meaning the paper has not been published in an approved journal.

## UI
There is a list and edit page for `papers`, `tags` and `users`. The list pages can be accessed
from the top header anywhere on the website. Once on the list page you can navigate to the edit
page by clicking on the associated title for each entry in the list. Furthermore, on the list
and edit page for any table a link to the related records in other tables are provided.

Lastly, the homepage allows you to search for papers related to a certain tag serving as an
extremely simple version of Google Scholar's search engine.

# Other
## Motivation
This was my CS3200 Final project. Also I use Google Scholar a lot and it
has a simple UI and seemed like a cool way to practice making relational
databases.

## Why Go?
I just wanted to get better with the language, so I decided to use it
to make this. Also Go is a pretty cool language, though I have to say 
not having generics in CRUD based applications feels AWFUL.

## Install
You can create all the tables by running the sql in `sql/all.sql`.
To build the server just run `go build`.
To populate the databse run `go run scripts/populate.go`.

## Scraping
I intially wanted to scrape BibTeX files from Google Scholar to populate
my databse, but quickly learned that this is explicitily made incredibly difficult.
There is a way, but it essentially gets flagged and banned immediately.

So, instead of actual academic papers I just made random data in `scripts/populate.go`
