package data

import (
	"database/sql"
	"time"
)

type User struct {
	ID           int64        `json:"id"`
	Username     string       `json:"username"`
	Password     string       `json:"password"`
	Email        string       `json:"email"`
	Birthdate    sql.NullTime `json:"birthdate"`
	CreationDate time.Time    `json:"creation_date"`
	LastVistDate sql.NullTime `json:"lastvisit_date"`
}
