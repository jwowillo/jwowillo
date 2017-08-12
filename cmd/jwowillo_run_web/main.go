package main

import (
	"flag"
	"fmt"
	"log"
	"path/filepath"

	"github.com/jwowillo/md2web"
	"github.com/jwowillo/trim/application"
	"github.com/jwowillo/trim/server"

	landgrab "github.com/jwowillo/landgrab/web"
)

// Constructor ...
type Constructor func(string, string, string) *application.Application

var (
	// url to run at.
	url string
	// port to run at.
	port     int
	projects = map[string]Constructor{
		"landgrab": landgrab.New,
	}
)

// main runs the server on the given URL and port.
func main() {
	var app *md2web.MD2Web
	ignores := []string{"README.md"}
	h := url
	if port == 80 {
		app = md2web.New(h, "content", ignores)
	} else {
		if url == "localhost" || port != 80 {
			h += fmt.Sprintf(":%d", port)
		}
		app = md2web.NewDebug(h, "content", ignores)
	}
	for name, constructor := range projects {
		if err := app.AddApplication(constructor(
			name,
			h,
			filepath.Join("build_"+name, "web"),
		)); err != nil {
			log.Fatal(err)
		}

	}
	s := server.New(url, port)
	s.AddHeader("Access-Control-Allow-Origin", "*")
	s.AddHeader("Access-Control-Allow-Headers", "Authorization")
	s.Serve(app)
}

// init parses the URL and port.
func init() {
	flag.StringVar(&url, "url", "localhost", "URL to listen from")
	flag.IntVar(&port, "port", 5000, "port to run on")
	flag.Parse()
}
