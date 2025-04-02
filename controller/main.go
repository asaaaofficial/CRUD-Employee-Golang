package controller

import (
	"net/http"
	"log"
)

func NewHelloWorld() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request)  {
		log.Println("==> Handler NewHelloWorld DIPANGGIL")
		w.Write([]byte("Hello World"))
	}
}