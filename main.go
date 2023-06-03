package main

import (
	"fmt"
	"net/http"
	"os"
)

func getData(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/data", getData)

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
