package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/jwowillo/md2web"
	"github.com/jwowillo/trim/server"
)

var (
	// host the trim.Server runs at.
	host string
	// port the trim.Server runs at.
	port int
)

// main runs the server on the given host and port.
func main() {
	var app *md2web.MD2Web
	ignores := []string{"README.md"}
	h := host
	if port == 80 {
		app = md2web.New(h, ignores)
	} else {
		if host == "localhost" || port != 80 {
			h += fmt.Sprintf(":%d", port)
		}
		app = md2web.NewDebug(h, ignores)
	}
	server.New(host, port).Serve(app)
}

// init parses the host and port.
func init() {
	message := "Usage: jwowillo <host> <port:int>\n"
	if len(os.Args) != 3 {
		log.Fatal(message)
	}
	host = os.Args[1]
	portArg := os.Args[2]
	portVal, err := strconv.Atoi(portArg)
	if err != nil {
		log.Fatal(err)
	}
	port = portVal
}
