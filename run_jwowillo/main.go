package main

import (
	"flag"
	"fmt"
	"log"
	"path/filepath"

	"github.com/jwowillo/jwowillo/project"
	"github.com/jwowillo/md2web"
	"github.com/jwowillo/trim/server"
)

var (
	// host to run at.
	host string
	// port to run at.
	port int
)

// main runs the server on the given host and port.
func main() {
	var app *md2web.MD2Web
	ignores := []string{"README.md"}
	h := host
	if port == 80 {
		app = md2web.New(h, "content", ignores)
	} else {
		if host == "localhost" || port != 80 {
			h += fmt.Sprintf(":%d", port)
		}
		app = md2web.NewDebug(h, "content", ignores)
	}
	for repo, constructor := range project.Map {
		if err := app.AddApplication(constructor(
			project.Name(repo),
			h,
			filepath.Join(project.StaticFolder(repo), "dist"),
		).Application); err != nil {
			log.Fatal(err)
		}

	}
	s := server.New(host, port)
	s.AddHeader("Access-Control-Allow-Origin", "*")
	s.Serve(app.Application)
}

// init parses the host and port.
func init() {
	flag.StringVar(&host, "host", "localhost", "host to run on")
	flag.IntVar(&port, "port", 5000, "port to run on")
	flag.Parse()
}
