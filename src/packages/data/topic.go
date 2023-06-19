package data

import (
	"database/sql"
	"fmt"
)

func TopicExists(id int) (bool, error) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/forum?parseTime=true")
	if err != nil {
		return false, err
	}
	defer db.Close()

	var result int
	err = db.QueryRow(`SELECT COUNT(*) FROM topics WHERE id = ?`, id).Scan(&result)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	fmt.Println(result)
	return (result > 0), nil
}
