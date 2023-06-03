<!-- Afficher un profil -->

`SELECT * FROM users WHERE id = [id];`
`SELECT * FROM users WHERE username = [username];`

<!-- Afficher le nombre de topics crées par un utilisateur -->

`SELECT COUNT(*) FROM topics AS t WHERE t.user_id = [id];`
`SELECT COUNT(*) FROM posts AS p WHERE p.user_id = [id];`

<!-- Afficher les 5 derniers topics crées par un utilisateur -->

`SELECT * FROM topics WHERE user_id = 3 ORDER BY creation_date DESC LIMIT 5;`

<!-- Nombre d'utilisateur -->

`SELECT COUNT(*) FROM users;`

<!-- Affichage dans la limite de 25 utilisateurs, ordre de rôle et alphabétique -->

`SELECT * FROM users ORDER BY role_id ASC, username ASC LIMIT 25;`
`SELECT * FROM users ORDER BY role_id ASC, username ASC LIMIT 25 OFFSET 26;`

<!-- EVZ : Récupérer toutes les réactions sur le post 1 -->

`SELECT * FROM `post_reactions` WHERE post_id = 1;`

<!-- EVZ : Récupérer tous les pseudos des utilisateurs ayant réagi sur le post 1 -->

`SELECT username FROM users 
    JOIN `post_reactions`AS pr 
    WHERE pr.post_id = 1 AND pr.user_id = id;`

<!-- EVZ : Récupérer tous les pseudos des utilisateurs && leur réaction, sur le post 1 -->

`SELECT username, r.name FROM users AS u
    JOIN `post_reactions`AS pr 
    JOIN`reactions` AS r
    WHERE pr.post_id = 1 AND pr.user_id = u.id AND pr.reaction_id = r.id
    ORDER BY r.name DESC;`

<!-- La même, en partant de la table réactions -->

`SELECT u.username, r.name FROM post_reactions AS pr
	JOIN 'users'as u
    JOIN 'reactions' as r
    WHERE post_id = 1 AND user_id = u.id AND reaction_id = r.id 
    ORDER BY r.name DESC;`

<!-- Récupérer le nom du rôle d'un utilisateur -->

SELECT (username, profile_picture, birthdate, isActive, creation_date, lastvisit_date, r.name) FROM users
JOIN LEFT `roles` as r ON role_id = r.id;

<!-- Récupérer les informations d'un topic, son premier post et son OP -->

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
WHERE t.id = 1;

<!-- Récupérer l'ID du premier post depuis un topic pour l'exclure des réponses, récupère les réponses et les -->
<!-- ordonne par total de (likes - dislikes) ou date de post -->

SELECT post_id INTO @firstPostID FROM topic_first_posts where topic_id = 1;

DROP TABLE IF EXISTS temp;
CREATE TEMPORARY TABLE temp SELECT p.content, u.username, p.creation_date, p.modification_date,
(SELECT COUNT(pr.post_id) from post_reactions as pr where pr.post_id = p.id and pr.reaction_id = 1) as "likes",
(SELECT COUNT(pr.post_id) from post_reactions as pr where pr.post_id = p.id and pr.reaction_id = 2) as "dislikes"
FROM posts AS p
LEFT JOIN users AS u ON p.user_id = u.id
WHERE p.topic_id = 1 AND p.id != @firstPostID;

SELECT _ FROM temp ORDER BY (likes - dislikes) DESC;
SELECT _ FROM temp ORDER BY creation_date ASC;

<!-- Créer un topic et son premier post -->

`INSERT INTO topics (subcategory_id, user_id, title) VALUES
(1, 4, "Why is datetime format so painful in JS");
SELECT @topicID := LAST_INSERT_ID();

INSERT INTO topic_tags (topic_id, tag_id) VALUES (@topicID, 1);

INSERT INTO posts (topic_id, user_id, content) VALUES
(@topicID, 4, "This is so bullshit please help me (this is an almost auto-generated post).");
SELECT @postID := LAST_INSERT_ID();

INSERT INTO topic_first_posts(topic_id, post_id) VALUES (@topicID, @post_ID);

<!-- Afficher tous les topics d'une sous-categorie -->

`SELECT * FROM topics WHERE subcategory_id = [subcategory_id];`

<!-- Creer un nouveau topic -->

`INSERT INTO topics (subcategory_id, user_id, title) VALUES ([subcategory_id], [user_id], [title]);`

<!-- Creer un nouveau post sur un topic existant -->

`INSERT INTO posts (topic_id, user_id, content) VALUES ([topic_id], [user_id], [content]);`

<!-- Afficher tous les posts sur un topic -->

`SELECT * FROM posts WHERE topic_id = [topic_id];`

<!-- Mise a jour du contenu d'un post -->

`UPDATE posts SET content = [new_content] WHERE id = [post_id];`

<!-- Suppression d'un post -->

`DELETE FROM posts WHERE id = [post_id];`

-- GET GENERAL DATA : Users, Categories, subcategories, topics, answers, reactions

SELECT COUNT(_) AS "categories",
(SELECT COUNT(_) FROM subcategories) as "subcategories",
(SELECT COUNT(_) FROM users) as "users",
(SELECT COUNT(_) FROM topics) as "topics",
(SELECT COUNT(_) FROM posts as p
LEFT JOIN topic_first_posts as tfp ON p.topic_id = tfp.topic_id
WHERE p.id != tfp.post_id) as "answers",
(SELECT COUNT(_) FROM post_reactions) as "reactions"
FROM categories;
