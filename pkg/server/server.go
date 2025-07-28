package server

import (
	"fmt"
	"net/http"
	"os"
)

func Run() error {

	// назначим порт (8080 - чтобы доступ был на localhost:8080)
	// порт 80 github actions не разрешает (порты ниже 1024 требуют привилегий суперпользователя)
	port, ok := os.LookupEnv("SQNStest_PORT")
	if !ok {
		port = "8080"
	}

	http.Handle("/", http.FileServer(http.Dir("web")))

	return http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
}
