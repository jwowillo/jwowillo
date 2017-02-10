// Package main allows multiple trim.Applications to be joined to a
// documentation trim.Application.
//
// trim.Applications are defined by a mapping of repository URL to
// trim.Application constructor that accepts subdomain and static folder
// arguments. Static folders are expected to be attached to the repository's
// root in a folder called 'static' if they exist and Makefiles are expected to
// be in the same place and named 'Makefile' if they exist.
package main

import (
	"flag"
	"fmt"
	"log"
	"sync"

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
		app = md2web.New(h, ignores)
	} else {
		if host == "localhost" || port != 80 {
			h += fmt.Sprintf(":%d", port)
		}
		app = md2web.NewDebug(h, ignores)
	}
	var wg sync.WaitGroup
	for repo, constructor := range projects {
		fmt.Println("Adding Application at", repo)
		if err := app.AddApplication(constructor(
			projectName(repo),
			h,
			staticFolder(repo),
		)); err != nil {
			log.Fatal(err)
		}
		wg.Add(1)
		go func() {
			fmt.Println("Downloading repository at", repo)
			if err := buildRepo(repo); err != nil {
				log.Fatal(err)
			}
			fmt.Println("Done downloading repository at", repo)
			wg.Done()
		}()

	}
	wg.Wait()
	s := server.New(host, port)
	s.AddHeader("Access-Control-Allow-Origin", "*")
	s.Serve(app)
}

// init parses the host and port.
func init() {
	flag.StringVar(&host, "host", "localhost", "host to run on")
	flag.IntVar(&port, "port", 5000, "port to run on")
	flag.Parse()
}
