package data

import (
	"database/sql"
)

type TopTrainer struct {
	UserID    int64  `json:"user_id"`
	Username  string `json:"username"`
	PostCount int64  `json:"post_count"`
	Position  int    `json:"position"`
}

func QueryTopTrainers(userID int) (result [6]TopTrainer, err error) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/forum?parseTime=true")
	if err != nil {
		return result, err
	}
	defer db.Close()

	rows, err := db.Query(`SELECT u.username, u.id, 
	(SELECT COUNT(*) FROM posts WHERE posts.user_id = u.id) AS "count"
    FROM posts AS p 
    RIGHT JOIN users AS u ON p.user_id = u.id
    GROUP BY u.id
    ORDER BY count desc;`)
	if err != nil {
		return result, err
	}
	defer rows.Close()

	i := 1
	for rows.Next() {
		var temp TopTrainer
		err = rows.Scan(&temp.Username, &temp.UserID, &temp.PostCount)
		temp.Position = i
		if i <= 5 {
			result[i-1] = temp
		}
		if i > 5 && int(temp.UserID) == userID {
			result[5] = temp

		}
		i++
	}
	return result, nil
}
