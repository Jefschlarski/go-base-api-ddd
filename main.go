package main

import (
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello, World!")
	r := router.GenRouter()

	log.Fatal(http.ListenAndServe(":8080", r))
}
