package server

import (
	"fmt"
	"net/http"
)

type Server struct {
}

type H map[string]any

func NewServer() Server {
	return Server{}
}

func (s Server) GET(pattern string, callback func(c Context)) {
	handle := func(w http.ResponseWriter, r *http.Request) {
		c := Context{Writer: w, Request: r}
		if r.Method != http.MethodGet {
			c.JSON(http.StatusBadRequest, H{"message": "Bad Request"})
			return
		}
		callback(c)
	}
	http.HandleFunc(pattern, handle)
}

func (s Server) POST(pattern string, callback func(c Context)) {

	handle := func(w http.ResponseWriter, r *http.Request) {
		c := Context{Writer: w, Request: r}
		if r.Method != http.MethodPost {
			c.JSON(http.StatusBadRequest, H{"message": "Bad Request"})
			return
		}
		callback(c)
	}
	http.HandleFunc(pattern, handle)
}

func (s Server) Run(addr ...string) {

	var tmp string

	switch len(addr) {
	case 0:
		tmp = ":8080"
	case 1:
		tmp = addr[0]
	default:
		panic("too many parameters")
	}

	err := http.ListenAndServe(tmp, nil)
	if err != nil {
		fmt.Println(err)
	}
}
