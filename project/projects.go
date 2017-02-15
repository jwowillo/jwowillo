package project

import (
	"github.com/jwowillo/harvester"
	"github.com/jwowillo/sample"
)

// Map is a mapping from Github repository URLs in go repository format
// style to trim.Application constructors for a web project.
//
// These will be attached to the main trim.Application at a subdomain named
// after the project's Githup repository.
var Map = map[string]Constructor{
	"github.com/jwowillo/harvester": harvester.New,
	"github.com/jwowillo/sample":    sample.New,
}
