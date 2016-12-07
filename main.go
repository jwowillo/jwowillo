package main

import (
	"log"

	"github.com/jwowillo/mancalai/app"
	"github.com/jwowillo/md2web"
)

func main() {
	server, err := md2web.NewServer("localhost", 5000, []string{"GET"})
	if err != nil {
		log.Fatal(err)
	}
	server.AddApplication(app.NewApplication("mancalai"))
	server.Run(5000)
}
