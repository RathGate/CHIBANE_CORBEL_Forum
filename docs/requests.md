<!-- Afficher un profil -->

`SELECT * FROM users WHERE id = [id];`
`SELECT * FROM users WHERE username = [username];`

<!-- Afficher le nombre de topics crées par un utilisateur -->

`SELECT COUNT(*) FROM topics AS t WHERE t.user_id = 8;`
`SELECT COUNT(*) FROM posts AS p WHERE p.user_id = 8;`

<!-- Afficher les 5 derniers topics crées par un utilisateur -->

`SELECT * FROM topics WHERE user_id = 3 ORDER BY creation_date DESC LIMIT 5;`

<!-- Nombre d'utilisateur et affichage dans la limite de 25 utilisateurs -->
