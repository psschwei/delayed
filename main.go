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
	fmt.Fprintf(w, "Starting to handle request\n")
	log.Print("Starting to handle request\n")
	doDelay("REQUEST")
	log.Print(r.Header)
	log.Print(r.Body)
	fmt.Fprintf(w, "Request handled\n")
	log.Print("Request handled\n")
}

// doDelay takes the name of an environmental variable and sleeps that amount of time
func doDelay(envvar string) {
	delay, err := strconv.Atoi(os.Getenv(envvar))
	if err != nil {
		delay = 0
	}
	if delay > 0 {
		for i := 0; i < delay; i++ {
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	fmt.Println("starting server")
	log.Print("starting server")
	doDelay("STARTUP")
	fmt.Println("server ready")
	log.Print("server ready")

	http.HandleFunc("/", handler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
