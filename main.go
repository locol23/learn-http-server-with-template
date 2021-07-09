package main

import (
	"html/template"
	"math/rand"
	"net"
	"net/http"
)

type Fortune struct {
	Name   string
	Result string
}

type Server struct {
	fortuneTemplate *template.Template
}

func (s Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	num := rand.Int()
	result := [...]string{"lucky", "awesome", "good", "nice", "great"}

	fortune := &Fortune{Name: r.FormValue("p"), Result: result[num%len(result)]}
	s.fortuneTemplate.Execute(w, fortune)
	s.logger(fortune.Result)
}

func (s Server) logger(message string) {
	s.logger(message)
}

func main() {
	rand.Seed(10)

	var err error
	listener, err := net.Listen("tcp", ":8080")

	if err != nil {
		return
	}

	server := Server{}
	server.fortuneTemplate, err = template.Must(template.New("msg"), nil).Parse("<html><body>{{.Name}}'s fortune is <b>{{.Result}}</b>!</body></html>")

	if err != nil {
		return
	}

	http.Handle("/", server)
	http.Serve(listener, server)
}
