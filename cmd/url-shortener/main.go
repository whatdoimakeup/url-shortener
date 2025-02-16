package main

import (
	"log"

	"github.com/whatdoimakeup/url-shortener/internal/app"
)

func main() {
	log.Println("Hello, i am here")
	log.Fatal(app.Run())
}