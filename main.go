package main

import (
	"log"
	"os"
	"strconv"

	"github.com/jwowillo/md2web"
	"github.com/jwowillo/trim"
)

var (
	// Domain the trim.Server runs at.
	Domain = "localhost"
	// Port the trim.Server runs at.
	Port = 5000
)

// main runs the server on the given domain and port.
func main() {
	app := md2web.New([]string{"README.md"}).Application
	trim.NewServer(Domain, Port).Serve(app)
}

// init parses the domain and port.
func init() {
	message := []byte("Usage: jwowillo <domain> <port:int>\n")
	if len(os.Args) != 3 {
		log.Fatal(message)
	}
	Domain = os.Args[1]
	portArg := os.Args[2]
	portVal, err := strconv.Atoi(portArg)
	if err != nil {
		log.Fatal(err)
	}
	Port = portVal
}
