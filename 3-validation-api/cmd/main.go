package main

import (
	"fmt"
	"net/http"
	"validation-api/configs"
	"validation-api/internal/auth"
	"validation-api/internal/verify"
)

func main()  {
	conf := configs.LoadConfig()

	router := http.NewServeMux()

	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		AuthConfig: &conf.Auth,
	})

	verify.NewVerifyHandler(router, verify.VerifyHandlerDeps{
		EmailVerifyConfig: &conf.EmailVerify,
	})

	server := http.Server{
		Addr: ":8081",
		Handler: router,
	}

	fmt.Println("Server is listening on port 8081")
	server.ListenAndServe()
}