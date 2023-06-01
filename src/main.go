package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func dbTest() {
	var (
		id       int
		username string
	)
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/Forum")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	fmt.Println("Success!")

	result, _ := db.Query("SELECT id, username FROM users")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		err := result.Scan(&id, &username)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, username)
	}
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
