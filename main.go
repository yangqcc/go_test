package main

import (
	"go_test/domain"
	"fmt"
	"go_test/util"
	"time"
)

func main() {
	domain.ShowUserName()
	fmt.Println(util.Add(6, 6))
	time.Sleep(10000 * time.Second)
}
