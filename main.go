// filename: main.go
package main

import (
	"log"
	"net/http"
	"time"
)

// create a handler function
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hi, my name is Addie Vasquez. \n"))
	w.Write([]byte("Aside from getting my asociates degree in I.T, I have a good amount of hobbbies. One would call it grandma hobbies.\n"))
	w.Write([]byte("I like baking, sewing, napping, etc. \n"))

}

func greeting(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now()
	w.Write([]byte("Hello, the time is:" + currentTime.Format("2006-01-02 15:04:05 Monday")))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	log.Println("starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
