### Get all general data

```
SELECT COUNT(*) as "categories",
    (SELECT COUNT(*) FROM subcategories) as "subcategories",
    (SELECT COUNT(*) FROM users) as "users",
    (SELECT COUNT(*) FROM topics) as "topics",
    (SELECT COUNT(*) FROM posts as p
    	LEFT JOIN topic_first_posts as tfp ON p.id
     	WHERE p.id != tfp.post_id) as "answers",
    (SELECT COUNT(*) FROM post_reactions) as "reactions"
    FROM categories
```

**Output**:

```
+------------+---------------+-------+--------+---------+-----------+
| categories | subcategories | users | topics | answers | reactions |
+------------+---------------+-------+--------+---------+-----------+
|          1 |             3 |     5 |      1 |       2 |         8 |
+------------+---------------+-------+--------+---------+-----------+
```

### Get the list of all usernames associated to a reaction on a given post

```
SELECT r.name as "reaction", GROUP_CONCAT(u.username ORDER BY u.username SEPARATOR ";") as "users"
	FROM post_reactions AS pr
	JOIN users AS u ON pr.user_id = u.id
    JOIN reactions AS r ON pr.reaction_id = r.id
    WHERE pr.post_id = [[1]]
    GROUP BY pr.reaction_id;
```

**Output**:

```
+-----------+--------------+
| reaction  |    users     |
+-----------+--------------+
| Like      | evz;RathGate |
| Dislike   | RandomUser1  |
+-----------+--------------+
```

### Get the total number of each reaction on a given post

```
SELECT r.name as "reaction", COUNT(u.username) as "users"
	FROM post_reactions AS pr
	JOIN users AS u ON pr.user_id = u.id
    JOIN reactions AS r ON pr.reaction_id = r.id
    WHERE pr.post_id = [[1]]
    GROUP BY pr.reaction_id;
```

**Output**:

```
+-----------+-------+
| reaction  | users |
+-----------+-------+
| Like      |     2 |
| Dislike   |     1 |
+-----------+-------+
```

### Gets the essential information of a topic and its first post

```
SELECT title, GROUP_CONCAT(DISTINCT tags.name SEPARATOR ";") as "tags", p.content, u.username,
t.creation_date, t.modification_date,
	(SELECT COUNT(pr.post_id) from post_reactions as pr where pr.post_id = p.id and pr.reaction_id = 1) as "likes",
	(SELECT COUNT(pr.post_id) from post_reactions as pr where pr.post_id = p.id and pr.reaction_id = 2) as "dislikes"
	FROM topics as t
	LEFT JOIN topic_tags AS tag ON t.id = tag.topic_id
	JOIN tags ON tag.tag_id = tags.id
	JOIN topic_first_posts AS fp ON t.id = fp.topic_id
	JOIN posts AS p ON fp.post_id = p.id
	LEFT JOIN users as u ON t.user_id = u.id
	WHERE t.id = [[1]];
```

**Output**:

```
+---------------------------+------------+----------------------+----------+---------------------+-------------------+--------+----------+
|           title           |    tags    |       content        | username |    creation_date    | modification_date | likes  | dislikes |
+---------------------------+------------+----------------------+----------+---------------------+-------------------+--------+----------+
| Why is datetime format... | Golang;SQL | This is so bullsh... | Kuha     | 2023-06-02 17:16:06 | NULL              |      2 |        1 |
+---------------------------+------------+----------------------+----------+---------------------+-------------------+--------+----------+
```

### Gets the essential informations of all the posts from a given topic

```
SELECT post_id INTO @firstPostID FROM topic_first_posts where topic_id = [[1]];

DROP TABLE IF EXISTS temp;
CREATE TEMPORARY TABLE temp SELECT p.content, u.username, p.creation_date, p.modification_date,
    (SELECT COUNT(pr.post_id) from post_reactions as pr where pr.post_id = p.id and pr.reaction_id = 1) as "likes",
    (SELECT COUNT(pr.post_id) from post_reactions as pr where pr.post_id = p.id and pr.reaction_id = 2) as "dislikes"
    FROM posts AS p
    LEFT JOIN users AS u ON p.user_id = u.id
    WHERE p.topic_id = [[1]] AND p.id != @firstPostID;

SELECT * FROM temp ORDER BY (likes - dislikes) DESC;
```

**Output**:

```
+---------+----------+---------------------+-------------------+-------+----------+
| content | username |    creation_date    | modification_date | likes | dislikes |
+---------+----------+---------------------+-------------------+-------+----------+
| Ratio   | RathGate | 2023-06-02 22:47:51 | NULL              |     2 |        0 |
| Feur    | evz      | 2023-06-02 22:48:07 | NULL              |     2 |        1 |
+---------+----------+---------------------+-------------------+-------+----------+
```

### Create a topic and its first post, with a given category ID, subcategory ID and user ID

```
INSERT INTO topics (subcategory_id, user_id, title) VALUES
(1, 2, "Hi, this is a test!");
SELECT @topicID := LAST_INSERT_ID();

INSERT INTO topic_tags (topic_id, tag_id) VALUES (@topicID, 3);

INSERT INTO posts (topic_id, user_id, content) VALUES
(@topicID, 2, "Hope this is gonna work properly!");
SELECT @postID := LAST_INSERT_ID();

INSERT INTO topic_first_posts(topic_id, post_id) VALUES (@topicID, @postID);
```

**Output**:

```
+---------------------+------------+-----------------------------------+----------+---------------------------------+-------+----------+
|        title        |    tags    |              content              | username | creation_date modification_date | likes | dislikes |
+---------------------+------------+-----------------------------------+----------+---------------------------------+-------+----------+
| Hi, this is a test! | Javascript | Hope this is gonna work properly! | RathGate | 2023-06-03 23:19:36             |     0 |        0 |
+---------------------+------------+-----------------------------------+----------+---------------------------------+-------+----------+
```
