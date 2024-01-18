package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, r *http.Request) {
		fmt.Println("server open")
		writer.Write([]byte("hello world"))
	})
	fmt.Println("huafoa")
	log.Fatal(http.ListenAndServe(":8097", nil))
	fmt.Println("test")
}