package main

import (
	"gin/cmd/server"
	"gin/db"
	"log"
)

func main() {
	if err := db.Connect(); err != nil {
		log.Fatalf("Ошибка подключения к БД: %v", err)
	}

	if err := server.StartServer(); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
