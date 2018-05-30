package main

//pemanggilan package
import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func helloworld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World") //yg tampil di web
}

func handleRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", helloworld).Methods("GET") //ambil methods get dari helworld
	myRouter.HandleFunc("/user", AllUsers).Methods("GET")
	myRouter.HandleFunc("/user/{name}/{email}", NewUser).Methods("POST")
	myRouter.HandleFunc("/user/{name}", DeleteUser).Methods("DELETE")
	myRouter.HandleFunc("/user/{name}/{email}", UpdateUser).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8081", myRouter)) //penampilan di port 8081 dengan package http
}

func main() {
	fmt.Println("GO Tutorial ORM") //yg tampil di terminal

	InitialMigration()

	handleRequest() //menjalankan local server
}
