package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("../../assets"))))
	fmt.Println("asdf")
}
