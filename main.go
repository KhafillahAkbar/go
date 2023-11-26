package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/KhafillahAkbar/go/controller/productcontroller"
	"github.com/KhafillahAkbar/go/models"
	"github.com/gorilla/mux"
)

func main() {
	models.ConnectDatabase()

	r := mux.NewRouter()

	r.HandleFunc("/products", productcontroller.Index).Methods("GET")
	r.HandleFunc("/product/{id}", productcontroller.Show).Methods("GET")
	r.HandleFunc("/product", productcontroller.Create).Methods("POST")
	r.HandleFunc("/product/{id}", productcontroller.Update).Methods("PUT")
	r.HandleFunc("/product", productcontroller.Delete).Methods("DELETE")

	fmt.Println("Server at 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
