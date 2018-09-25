package controller

import (
	"net/http"
	"fmt"
)

type Controller struct {
	Data interface{}
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello,World!")
}

func ShowName() {
	fmt.Print("this is name!")
}
