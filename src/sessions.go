package main

import (
	"database/sql"
	"fmt"
	"forum/packages/data"
	"net/http"

	"github.com/gorilla/sessions"
)

var (
	key        = []byte("super-secret-key")
	store      = sessions.NewCookieStore(key)
	cookieName = "auth"
)

func setSession(r *http.Request, w *http.ResponseWriter, userID int) error {
	session, _ := store.Get(r, cookieName)
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/forum?parseTime=true")
	if err != nil {
		return err
	}
	defer db.Close()

	var tempUsername string
	err = db.QueryRow(`SELECT username FROM users WHERE id = ?`, userID).Scan(&tempUsername)
	if err != nil || tempUsername == "" {
		fmt.Println(err)
		return err
	}
	session.Values["authenticated"] = true
	session.Values["id"] = userID
	session.Values["username"] = tempUsername
	session.Save(r, *w)
	return nil
}

func clearSession(r *http.Request, w *http.ResponseWriter) {
	session, _ := store.Get(r, cookieName)
	session.Values["authenticated"] = false
	session.Save(r, *w)
}

func getSession(r *http.Request) (tData data.TemplateData) {
	session, _ := store.Get(r, cookieName)
	if (session.Values["authenticated"] == nil || !session.Values["authenticated"].(bool)) || !(session.Values["id"].(int) > 0) {
		return tData
	}

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/forum?parseTime=true")
	if err != nil {
		fmt.Println(err)
		return tData
	}
	defer db.Close()

	var tempUsername string
	err = db.QueryRow(`SELECT username FROM users WHERE id = ?`, session.Values["id"].(int)).Scan(&tempUsername)
	if err != nil || tempUsername == "" {
		fmt.Println(err)
		return tData
	}

	tData.User = data.TemplateUser{
		ID:              session.Values["id"].(int),
		Username:        tempUsername,
		IsAuthenticated: true,
	}

	return tData
}
