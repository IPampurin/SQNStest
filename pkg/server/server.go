package server

import (
	"fmt"
	"net/http"
)

func Run() error {

	// назначим порт
	port := "80"

	http.Handle("/", http.FileServer(http.Dir("web")))

	return http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
}
