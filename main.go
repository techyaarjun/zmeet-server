package main

import (
	_ "github.com/joho/godotenv/autoload"
	server "zmeet/cmd"
)

func main() {
	server.Start()
}
