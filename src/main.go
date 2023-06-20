package main

import (
	"fmt"
	"forum/packages/utils"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var DATABASE_ACCESS = utils.DB_Access{
	User:     "root",
	Password: "",
	Port:     3306,
	Type:     "sql",
	DBName:   "forum",
}

func main() {
	r := mux.NewRouter()
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))

	// Handles routing:
	r.HandleFunc("/", indexHandler)
	r.HandleFunc("/register", registerHandler)
	r.HandleFunc("/login", loginHandler)
	r.HandleFunc("/logout", logoutHandler)
	r.HandleFunc("/topics", topicsHandler)
	r.HandleFunc("/topic/{id}", topicHandler)
	r.HandleFunc("/new-topic", newTopicHandler)
	r.NotFoundHandler = http.HandlerFunc(notFoundHandler)

	// Launches the server:
	preferredPort := ":8080"
	fmt.Printf("Starting server at port %v\n", preferredPort)
	if err := http.ListenAndServe(preferredPort, r); err != nil {
		log.Fatal(err)
	}
}
