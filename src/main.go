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

//func dbGetUsers() {
//	var allUsers []*structs.User
//	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/forum?parseTime=true")
//	if err != nil {
//		panic(err.Error())
//	}
//	defer db.Close()
//	fmt.Println("Success!")
//
//	rows, _ := db.Query("SELECT id, username, password, birthdate, creation_date, lastvisit_date FROM users")
//	if err != nil {
//		panic(err.Error())
//	}
//
//	for rows.Next() {
//		u := new(structs.User)
//		err := rows.Scan(&u.ID, &u.Username, &u.Password, &u.Birthdate, &u.CreationDate, &u.LastVistDate)
//		if err != nil {
//			log.Fatal(err)
//		}
//		allUsers = append(allUsers, u)
//		fmt.Println(u)
//	}
//	fmt.Println(len(allUsers))
//}

func dbGetCategories() ([]structs.Category, error) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/forum?parseTime=true")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, name FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []structs.Category
	for rows.Next() {
		var category structs.Category
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

func main() {
	//dbGetUsers()
	//dbGetCategories()
	r := mux.NewRouter()
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))

	// Handles routing:
	r.HandleFunc("/", indexHandler)
	r.HandleFunc("/categories", catHandler)
	r.HandleFunc("/register", registerHandler)
	r.HandleFunc("/login", loginHandler)
	r.HandleFunc("/profile", profileHandler)
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
