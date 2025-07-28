package main

import (
	"SQNStest/pkg/server"
	"fmt"
)

func main() {

	var err error

	err = server.Run()
	if err != nil {
		fmt.Printf("ошибка запуска сервера: %v\n", err)
		return
	}

}
