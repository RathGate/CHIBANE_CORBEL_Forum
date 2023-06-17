package data

import (
	"database/sql"
	"time"
)

type File struct {
	UserID sql.NullInt64  `json:"user_id"`
	Name   string         `json:"name"`
	Type   sql.NullString `json:"type"`
}

type User struct {
	ID             int64        `json:"id"`
	Username       string       `json:"username"`
	Password       string       `json:"password"`
	Email          string       `json:"email"`
	ProfilePicture File         `json:"profile_picture"`
	Birthdate      sql.NullTime `json:"birthdate"`
	CreationDate   time.Time    `json:"creation_date"`
	LastVistDate   sql.NullTime `json:"lastvisit_date"`
}

type Category struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
type Tag struct {
	Name  string
	Color string
}
