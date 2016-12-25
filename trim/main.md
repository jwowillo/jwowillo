## trim

trim is a web framework based on decorators.

<center>![Turkey Leg](http://{{ cdn }}/trim/leg.png)</center>

Making an API which tells a user the length of their name and that caches
responses is simple using trim. Three main components will be needed: a
Controller, an Application, and a Server. Along with these, Decorators will
be applied.

These three components correspond to different parts of a URL. For example, if a
URL has structure \<subdomain\>.\<domain\>.\<tld\>/\<path\>, the Server maps to
the entire URL, Applications match to subdomains, and Controllers map to paths.

The Controller can be implemented like

```
package main

import (
	"net/http"
	"time"

	"github.com/jwowillo/trim"
	"github.com/jwowillo/decorators"
	"github.com/jwowillo/responses"
)

// LengthController is a trim.Controller which returns the lengths of names
// sent in trim.Requests.
type LengthController struct {
	trim.Decorated // Embedding trim.Decorated means no
                       // trim.Decoratable::AddDecorator method needs to be
                       // added.
}

// NewLengthController creates a LengthController.
type NewLengthController() *LengthController {
	return &LengthController{}
}

// Decorators is a list of a single trim.Decorator which caches trim.Responses
// for an hour at a time.
//
// This way, if someone asks how long their name is and that name has already
// been seen, the trim.Request doesn't need to be rehandled.
func (c *LengthController) Decorators() []trim.Decorator {
	return []trim.Decorator{decorators.CacheDecorator(time.Hour)}
}

// Path the controller is located at.
//
// <name> matches any string after the first /. It is placed at key "name" in
// the trim.Request's PathArguments.
func (c *LengthController) Path() string {
	return "/<name>"
}

// Handle the trim.Request by returning a JSON trim.Response containing the
// length of the passed name.
func (c *LengthController) Handle(r *trim.Request) trim.Response {
	name := r.PathArguments()["name"]
	return responses.NewPackagedJSONResponse(
		responses.AnyMap{"length": len(name)},
		http.StatusOK,
	)
}
```

The Controller then needs to added to an Application. This one will be located
at subdomain "api".

```
import "github.com/jwowillo/handlers"

// Application containing a LengthController located at "api".
//
// Returns a JSON trim.Response for 404s instead of an HTML trim.Response.
func Application() *trim.Application {
	app := trim.NewApplication("api")
	app.AddController(NewLengthController())
	app.SetHandle404(handlers.HandleJSON404)
	return app
}
```

Finally, the Application needs to be added to a Server. This will be located
at namelength.com.

```
// Server containing the API trim.Application.
//
// Expects to be run at namelength.com.
func Server() *trim.Server {
	server := trim.NewServer("namelength.com")
	server.AddApplication(Application())
	return server
}
```

Running the Server is simple with the above steps completed.

```
// main runs the namelength.com web server.
func main() {
	Server().Run(80) // Listen on port 80.
}
```

Assuming the above is located in a file called server.go, it can be ran with
`go run server.go`.

Requests can then be made like

```
$ curl -i "api.namelength.com/jim"

HTTP/1.1 200 OK
Access-Control-Allow-Origin: *
Content-Type: application/json
Date: Sun, 25 Sep 2016 13:49:48 GMT
Content-Length: 21

{"data":{"length":3}}

$ curl -i "api.namelength.com/joseph"

HTTP/1.1 200 OK
Access-Control-Allow-Origin: *
Content-Type: application/json
Date: Sun, 25 Sep 2016 13:49:48 GMT
Content-Length: 21

{"data":{"length":6}}
```

The framework has many other features and details which are described in the
rest of the documentation.
