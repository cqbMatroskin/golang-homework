package main

import (
	"fmt"
	"net/http"
)

func main()  {
	router := http.NewServeMux()
	NewRandomHandler(router)

	server := http.Server{
		Addr: ":8081",
		Handler: router,
	}

	fmt.Println("Server is run on port 8081")
	server.ListenAndServe()
}