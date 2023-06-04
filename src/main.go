package main

import (
	"database/sql"
	"fmt"
	"forum/packages/users"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func dbTest() {
	var allUsers []*users.User
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/forum?parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	fmt.Println("Success!")

	rows, _ := db.Query("SELECT id, username, password, birthdate, creation_date, lastvisit_date FROM users")
	if err != nil {
		panic(err.Error())
	}

	for rows.Next() {
		u := new(users.User)
		err := rows.Scan(&u.ID, &u.Username, &u.Password, &u.Birthdate, &u.CreationDate, &u.LastVistDate)
		if err != nil {
			log.Fatal(err)
		}
		allUsers = append(allUsers, u)
		fmt.Println(u)
	}
	fmt.Println(len(allUsers))
}
func main() {
	dbTest()
	r := mux.NewRouter()
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))

	// Handles routing:
	r.HandleFunc("/", indexHandler)
	r.HandleFunc("/register", registerHandler)
	r.HandleFunc("/success", successHandler)
	r.HandleFunc("/error", errorHandler)
	r.NotFoundHandler = http.HandlerFunc(notFoundHandler)

	// Launches the server:
	preferredPort := ":8080"
	fmt.Printf("Starting server at port %v\n", preferredPort)
	if err := http.ListenAndServe(preferredPort, r); err != nil {
		log.Fatal(err)
	}
}
