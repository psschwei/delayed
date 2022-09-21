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
	log.Print("Starting to handle request\n")
	log.Print(r.Header)
	doDelay("REQUEST")
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
	log.Print("starting server")
	doDelay("STARTUP")
	log.Print("server ready")

	http.HandleFunc("/", handler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
