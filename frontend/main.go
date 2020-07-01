package main

import (
	"fmt"

	"github.com/SimonTanner/simple-grpc-app/frontend/frontend-app"
)

func main() {
	frontend := frontend.New()

	err := frontend.Start(":8080")

	if err != nil {
		fmt.Println(err)
	}

}
