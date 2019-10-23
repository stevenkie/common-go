package main

import (
	"log"
	"time"

	api "github.com/sleey/common-go/api"
	cb "github.com/sleey/common-go/circuitbreaker"
)

func main() {
	// Create CB with errorThreshold: 2, successThreshold: 1, CB Open 1000 milliseconds
	cbreaker := cb.New(2, 1, 1000)

	req := api.Request{
		Method:  "GET",
		URL:     "http://api.example.com/example",
		Timeout: 1 * time.Millisecond,
		CB:      &cbreaker,
	}

	for i := 0; i < 10; i++ {
		_, err := req.Do()
		if err != nil {
			log.Println(err)
			continue
		}
	}
}
