package main

import (
	"log"

	"github.com/whatdoimakeup/url-shortener/internal/app"
	"github.com/whatdoimakeup/url-shortener/internal/database"
)

func main() {
	log.Println("Hello, i am here")
	if err := database.Connect(); err != nil {
		log.Fatal("не удалось подключиться к базе данных, завершаю работу.")
	}
	log.Fatal(app.Run())
}