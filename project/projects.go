package project

import landgrab "github.com/jwowillo/landgrab/web"

// Map is a mapping from Github repository URLs in go repository format
// style to trim.Application constructors for a web project.
//
// These will be attached to the main trim.Application at a subdomain named
// after the project's Githup repository.
var Map = map[string]Constructor{
	"github.com/jwowillo/landgrab": landgrab.New,
}
