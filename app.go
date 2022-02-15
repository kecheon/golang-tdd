package main

import (
	"golang/di"
	"golang/mocking"
	"os"
)

func main() {
	di.Greet(os.Stdout, "Foo")
	// log.Fatal(http.ListenAndServe(":8000", http.HandlerFunc(di.MyGreetHandler)))
	mocking.Countdown(os.Stdout)
}
