package server

import (
	"fmt"
	"net/http"
)

func Run() error {

	// назначим порт (80 - чтобы доступ был на localhost)
	//	port, ok := os.LookupEnv("SQNStest_PORT")
	//	if !ok {
	port := "80"
	//	}

	http.Handle("/", http.FileServer(http.Dir("web")))

	return http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
}
