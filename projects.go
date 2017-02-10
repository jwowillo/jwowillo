package main

import (
	"github.com/jwowillo/harvester"
	"github.com/jwowillo/trim"
)

// projects is a mapping from Github repository URLs in go repository format
// style to trim.Application constructors for a web project.
//
// These will be attached to the main trim.Application at a subdomain named
// after the project's Githup repository.
var projects = map[string]applicationConstructor{
	"github.com/jwowillo/harvester": harvester.New,
}

// applicationConstructor constructs a trim.Application which expects to exist
// at the given subdomain and host and have its static files located in the
// given static folder.
type applicationConstructor func(
	subdomain, host, staticFolder string,
) trim.Application
