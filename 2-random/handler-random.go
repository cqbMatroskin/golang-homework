package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

type RandomHandler struct {}

func NewRandomHandler(router *http.ServeMux) {
	handler := &RandomHandler{}
	router.HandleFunc("/random", handler.Random())
}

func (handler *RandomHandler) Random() http.HandlerFunc {
	return func (w http.ResponseWriter, req *http.Request)  {
		randomNum := rand.Intn(6) + 1
		w.Write([]byte(fmt.Sprint(randomNum)))
	}
}