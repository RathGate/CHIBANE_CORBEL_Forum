package main

import (
	"database/sql"
	"fmt"
	"forum/packages/structs"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func dbTest() {
	var allUsers []*structs.User
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/forum?parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	fmt.Println("Success!")

	rows, _ := db.Query("SELECT id, username, password, creation_date FROM users")
	if err != nil {
		panic(err.Error())
	}

	for rows.Next() {
		u := new(structs.User)
		err := rows.Scan(&u.ID, &u.Username, &u.Password, &u.CreationDate)
		if err != nil {
			log.Fatal(err)
		}
		allUsers = append(allUsers, u)
	}
	fmt.Println(len(allUsers))
}
func main() {
	dbTest()
	r := mux.NewRouter()
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))

	// Handles routing:
	r.HandleFunc("/", indexHandler)
	r.NotFoundHandler = http.HandlerFunc(notFoundHandler)

	// Launches the server:
	preferredPort := ":8080"
	fmt.Printf("Starting server at port %v\n", preferredPort)
	if err := http.ListenAndServe(preferredPort, r); err != nil {
		log.Fatal(err)
	}
}
