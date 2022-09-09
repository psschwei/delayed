package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	target := "delayed"
	fmt.Fprintf(w, "Hello %s!\n", target)
}

func main() {
	log.Print("delayed: simulating slow startup...")

	delay, err := strconv.Atoi(os.Getenv("SLEEP"))
	if err != nil {
		delay = 30
	}
	for i := 0; i < delay; i++ {
		time.Sleep(1 * time.Second)
	}

	log.Print("delayed: starting server...")
	http.HandleFunc("/", handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("delayed: listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
