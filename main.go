package main

import (
	"fmt"
	"golang-crud/helpers"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	fmt.Println("Start server")

	routes := httprouter.New()

	routes.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "Welcome home page")
	})

	server := http.Server{Addr: "localhost:8888", Handler: routes}

	err := server.ListenAndServe()
	helpers.PanicIfError(err)
}
