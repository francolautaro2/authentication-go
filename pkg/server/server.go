package server

import (
	"auth-go/pkg/handlers"
	"auth-go/pkg/token"
	"log"
	"net/http"
)

type Server struct {
	addr string
	port string
}

func CreateServer(addr string, port string) *Server {
	return &Server{
		addr: addr,
		port: port,
	}
}

func (s *Server) Listen() {
	mux := http.NewServeMux()

	mux.HandleFunc("/Login", handlers.Login)
	mux.HandleFunc("/Singup", handlers.Signup)
	mux.HandleFunc("/Home", token.ValidateToken(handlers.Home))

	err := http.ListenAndServe(s.port, mux)
	if err != nil {
		log.Fatal(err)
	}
}
