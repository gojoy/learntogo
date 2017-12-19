package main

import (
	"net/http"
	"fmt"
	"log"
)

func main() {
	http.HandleFunc("/",handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,"url.path = %v\n",r.URL.Path)
}

