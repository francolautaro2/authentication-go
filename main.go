package main

import (
	"auth-go/pkg/server"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()
	port := os.Getenv("PORT")
	addr := "127.0.0.1"

	s := server.CreateServer(addr, port)

	s.Listen()
}
