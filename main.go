package main

import (
	"SQNStest/pkg/server"
	"fmt"

	"github.com/joho/godotenv"
)

func main() {

	var err error

	err = godotenv.Load()
	if err != nil {
		fmt.Printf("ошибка загрузки .env файла: %v\n", err)
		return
	}

	err = server.Run()
	if err != nil {
		fmt.Printf("ошибка запуска сервера: %v\n", err)
		return
	}

}
