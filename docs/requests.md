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

``
