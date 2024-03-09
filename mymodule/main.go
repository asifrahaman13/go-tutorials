package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)


func greet(){
	fmt.Println("Hey there mod users.")
	
}

func serveHome(w http.ResponseWriter, r *http.Request){
    w.Write([]byte("<h1>Welcome to the golang series</h1>"))
}

func main(){
	fmt.Println("Hello mod in golang")
	greet()
	r:=mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")
	log.Fatal(http.ListenAndServe(":4000",r))
}