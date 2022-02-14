package main

import (
	"golang/di"
	"log"
	"net/http"
	"os"
)

func main() {
	di.Greet(os.Stdout, "Foo")
	log.Fatal(http.ListenAndServe(":8000", http.HandlerFunc(di.MyGreetHandler)))
}
