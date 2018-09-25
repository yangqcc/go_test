package main

import (
	"go_test/controller"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", controller.SayHelloName)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
