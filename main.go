// filename: main.go
package main

import (
	"log"
	"math/rand"
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
	w.Write([]byte("Hello, the time is: " + currentTime.Format("2006-01-02 15:04:05, Monday")))
}

type K string
type V string

func pick(m map[K]V) V {
	k := rand.Intn(len(m))
	for _, x := range m {
		if k == 0 {
			return x
		}
		k--
	}
	panic("unreachable")
}

func random(w http.ResponseWriter, r *http.Request) {
	m := map[K]V{
		"one":   "You're doing great! Keep going :)",
		"two":   "Start where you are. Use what you have. Do what you can.- Arthur Ashe, tennis star",
		"three": "The biggest adventure you can take is to live the life of your dreams. - Oprah Winfrey, talk show host",
		"four":  "Practice isn’t the thing you do once you’re good. It’s the thing you do that makes you good. - Malcolm Gladwell, author",
		"five":  "Why fit in when you were born to stand out? -Dr. Seuss, author",
	}

	x := pick(m)

	w.Write([]byte(x))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/greeting", greeting)
	mux.HandleFunc("/random", random)

	log.Println("starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
