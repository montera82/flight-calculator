package main

import (
	"flights-calculator/pkg/handler"
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/v1/calculate", handler.CalculateHandler)
	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
