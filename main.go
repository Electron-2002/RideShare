package main

import (
	db "app/postgres"
	"fmt"
	"net/http"
	"os"
)

func getDrivers(w http.ResponseWriter, req *http.Request) {
	rows, err := db.Connection.Query("SELECT name FROM drivers")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	data := ""
	for rows.Next() {
		var name string
		err = rows.Scan(&name)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(name)
		data += fmt.Sprintf("%s ", name)
	}

	err = rows.Err()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprintf(w, data)
}

func main() {
	db.InitDB()
	defer db.Connection.Close()

	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/drivers", getDrivers)

	serverEnv := os.Getenv("SERVER_ENV")

	if serverEnv == "DEV" {
		fmt.Println("Server is running on port 8080")
		http.ListenAndServe(":8080", nil)
	} else if serverEnv == "PROD" {
		fmt.Println("Server is running on port 443")
		http.ListenAndServeTLS(":443",
			"/etc/letsencrypt/live/rideshare.satejbidvai.tech/fullchain.pem",
			"/etc/letsencrypt/live/rideshare.satejbidvai.tech/privkey.pem",
			nil,
		)
	}
}
