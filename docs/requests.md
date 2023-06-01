<!-- Afficher un profil -->

`SELECT * FROM users WHERE id = [id];`
`SELECT * FROM users WHERE username = [username];`

<!-- Afficher le nombre de topics crées par un utilisateur -->

`SELECT COUNT(*) FROM topics AS t WHERE t.user_id = [id];`
`SELECT COUNT(*) FROM posts AS p WHERE p.user_id = [id];`

<!-- Afficher les 5 derniers topics crées par un utilisateur -->

`SELECT * FROM topics WHERE user_id = 3 ORDER BY creation_date DESC LIMIT 5;`

<!-- Nombre d'utilisateur et affichage dans la limite de 25 utilisateurs -->

`SELECT COUNT(*) FROM users;`
`SELECT * FROM users ORDER BY username ASC;`




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

``
