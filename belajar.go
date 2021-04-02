package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/afistapratama12/belajar-go-web/handler"
)

func main() {

	// create server mux
	mux := http.NewServeMux()

	mux.HandleFunc("/form", handler.HandleForm)
	mux.HandleFunc("/post-get", handler.PostGet)
	mux.HandleFunc("/list", handler.HandelListProduct)
	mux.HandleFunc("/tables", handler.HandleTables)
	mux.HandleFunc("/author", handler.HandlerAuthor)
	mux.HandleFunc("/product", handler.HandlerProduct)
	mux.HandleFunc("/process", handler.HandleProcess)
	mux.HandleFunc("/", handler.HandlerIndex)

	// get static file directory path
	fileServer := http.FileServer(http.Dir("assets"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// port address
	port := "localhost:3030"

	log.Println("Server started at", port)

	// excecute http mux
	err := http.ListenAndServe(port, mux)

	if err != nil {
		fmt.Println(err.Error())
	}
}
