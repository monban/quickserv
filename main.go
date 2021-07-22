package main

import (
	"log"
	"net/http"
)

type server struct {
	fs http.Handler
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("%v %v %v%v", r.Proto, r.Method, r.Host, r.URL)
	s.fs.ServeHTTP(w, r)
}

func main() {
	s := server{
		fs: http.FileServer(http.Dir(".")),
	}
	log.Print("Server starting on port 8000")
	log.Fatal(http.ListenAndServe(":8000", &s))
}
