package data

import (
	"database/sql"
	"forum/packages/utils"
)

type Category struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func GetCategories(dba utils.DB_Access) ([]Category, error) {
	db, err := sql.Open("mysql", dba.ToString())
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, name FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []Category
	for rows.Next() {
		var category Category
		err := rows.Scan(&category.ID, &category.Name)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return categories, nil
}
