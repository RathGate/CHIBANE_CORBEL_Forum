# Trainers' Corner (Forum)

![Topics page of the website](https://i.postimg.cc/X7YSqZtt/Capture-d-cran-2024-10-15-182121.png)
![Mobile version](https://i.postimg.cc/tgyzR6Mx/Capture-d-cran-2024-10-15-182419.png)

## About

Project submitted for a full-stack course, realised by [evzs](https://github.com/evzs) and [RathGate](https://github.com/RathGate). This project is intended to be a forum about the Pokémon video-games series and the subjects that revolve around it. Users can log in, create topics, post answers and react to other users' posts !

Technical documentation: [HERE](https://github.com/RathGate/CHIBANE_CORBEL_Forum/tree/main/docs) ( /src/documentation )

How we organised our work: [HERE (Asana)](https://app.asana.com/0/1204809319000403/timeline) 

## Technical Specifications

 - Back-end: Golang 
 - Front-end: HTML, CSS, JS (with Axios library).
 - DMBS: MySQL (used with WAMP Server)

**COMPATIBILITY:** The website has been entirely tested on the latest versions of Chrome, Firefox and Edge. It should also appear responsive on Chrome, Safari and Firefox mobile browsers !

## How to use the program

For now, the website has not been hosted anywhere, though we might find a way to host it in the future. In order to visit it on your computer, you must run both the database server and the golang server concurrently.

To clone the repository:

    git clone https://github.com/RathGate/CHIBANE_CORBEL_Forum

### Database Server

The given database script has been tested both on XAMPP (MariaDB) and XAMP (MySQL). However, you might encounter some compatibility issues if you choose to use another DBMS than the one that generated the script (MySQL). You will find the said script in `docs/script`.

For now, database access values have been hardcoded into the back-end code. If you happen to have settings different than the plain default-not-secure-root-whatever-values, you can change them in `src/main.go`:

      
      var DATABASE_ACCESS = utils.DB_Access{
        User:  "root",
        Password:  "",
        Port:  3306,
        Type:  "sql",
        DBName:  "forum",
      }

### Back-end Server

Launch a terminal in the `/src` folder:

    go run .

**NOTE:** `go run main.go` will not work properly as the `package main` is divided into three separate files and not just contained in `main.go`.

The website should be available on `localhost:8080`. In case of port collision, look for this line of code at the bottom of `/src/main.go` :

    preferredPort := ":8080"

Change the numerical value after the colon with any other port number.

Enjoy ♫
