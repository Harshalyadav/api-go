package main

import (
	"fmt"
	"log"
	"net/http"
)

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,"welcom to the home page")
	fmt.Println("Endpoint hit :homepage")
}

func foopages(w http.ResponseWriter, r *http.Request){
   fmt.Fprintf(w,"Welcom to the foo page")
   log.Fatalln("Foo page")
}

func main() {

	http.HandleFunc("/", homepage)
	http.HandleFunc("/foo", foopages)
	//  running on port number 10000
	http.ListenAndServe("localhost:10000", nil)
}
