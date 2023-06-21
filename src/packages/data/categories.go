package data

import (
	"database/sql"
	"forum/packages/utils"
)

type Category struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	TopicCount   int64  `json:"topic_count"`
	PostCount    int64  `json:"post_count"`
	MinReadRole  int64  `json:"min_read_role"`
	MinWriteRole int64  `json:"min_write_role"`
}

func GetCategoryData(dba utils.DB_Access, categoryID int) (category Category, err error) {
	db, err := sql.Open("mysql", dba.ToString())
	if err != nil {
		return category, err
	}
	defer db.Close()

	err = db.QueryRow("SELECT * FROM categories WHERE id = ?;", categoryID).Scan(&category.ID, &category.Name, &category.MinReadRole, &category.MinWriteRole)
	if err != nil {
		return category, err
	}

	return category, nil
}

func GetCategories(dba utils.DB_Access, roleID int) ([]Category, error) {
	if roleID == 0 {
		roleID = 4
	}
	db, err := sql.Open("mysql", dba.ToString())
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT    c.id,    c.name AS category_name,  COUNT(DISTINCT t.id) AS topic_count, COUNT(DISTINCT p.id), c.min_read_role, c.min_write_role AS post_count FROM    categories c LEFT JOIN    topics t ON c.id = t.category_id LEFT JOIN   posts p ON t.id = p.topic_id WHERE ? <= c.min_read_role GROUP BY     c.id, c.name;", roleID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []Category
	for rows.Next() {
		var category Category
		err := rows.Scan(&category.ID, &category.Name, &category.TopicCount, &category.PostCount, &category.MinReadRole, &category.MinWriteRole)
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
