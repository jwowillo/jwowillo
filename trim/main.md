## trim

* Tutorial: https://godoc.org/github.com/jwowillo/trim/tutorial
* Examples: https://godoc.org/github.com/jwowillo/trim/examples
* API Documentation: https://godoc.org/github.com/jwowillo/trim
* Repository: https://github.com/jwowillo/trim

trim is a golang web framework which simplifies application by maximizing
reuse, testability, and power while minimizing boilerplate and magic.

<center>![Turkey Leg](http://{{ static }}/trim/leg.png)</center>

### Overview

trim is a powerful web application framework due to several design features:

* Trimmings
* Responses
* Application and Controller separation
* Nestable Applications
* Default implementations

These allow web application to look like:

```
package main

// main serves the Echo web application.
//
// Since the premade application.Web is used, the client, API, and static file
// serving are separated at different subdomains. The client also will cache by
// default and serve a template file. The API will have ping and schema
// controllers already added and will package response from other controllers.
// The Static allows static files to also be templated.
func main() {
	// Assume these are read in from somewhere else.
	var (
		domain string
		port int
	)
	web := application.NewWebWithConfig(
		// Custom Client configuration allows home template to exist at
		// project root.
		application.ClientConfig{
			ClientTemplateHome: "index.html",
			// Inject the name of the web application into the
			// template so that it can be dynamically changed.
			Args: func(r *trim.Request) trim.AnyMap {
				return trim.AnyMap{"name": "Echo Client"}
			},
		},
		application.APIDefault, application.StaticDefault,
	)
	web.API().AddController(NewEchoController())
	trim.NewServer(domain, port).Serve(web.Application)
}

// EchoController echos strings to a name.
type EchoController struct {}

// NewEchoController creates an EchoController.
func NewEchoController() *EchoController { return &EchoController{} }

// Path matches a single name to be echoed to.
func (c *EchoController) Path() string { return "/:name" }

// Trimmings only allows GET requests.
func (c *EchoController) Trimmings() []trim.Trimming {
	return []trim.Trimming{trimmings.NewAllow([]string{"GET"})}
}

// Handle the request by parsing the name and string to echo from the request
// and returning a message to the user.
func (c *EchoController) Handle(r *trim.Request) trim.Response {
	msg := fmt.Sprintf("%s: %s", r.URLArg("name"), r.FormArg("string"))
	return response.NewJSON(
		trim.AnyMap{"message": msg},
		trim.Code(http.StatusOK),
	)
}
```

### Architecture

trim's core architecture is marked by Server and Application separation,
nestable Applications, and Trimmed types.

<center>![Architecture](http://{{ static }}/trim/arch.png)</center>

Shown in the diagram, trim serves Applications through very loosely coupled
Servers. Servers can provide access to Application's without any knowledge of
the specific Application. Applications represent a tree of functionality groups
accessbile through subdomains and base paths. An Application has its own
functionalities and also has child Applications which have their own
functionalities. These functionalities are provided through the Application's
Controllers. Controllers are simply Handlers with an associated path.
Applications also have Handler's for error situations.

Application's and Handlers are Trimmed types. This means that decorator
functionality can be added to them to change behavior without convoluting
Application and Handler logic through Trimmings. The Trimmings are instantiated
and listed and the Request handling process correctly wraps the functionalities
around the Trimmed types.

Request handling is a process which first involves the Server receiving a
communication and using a RequestBuilder to build a Request from it. The Server
then asks the Application for a Response to the built Request. To satisfy this,
the Appilcation finds the correct Handler for the Request by navigating the
Application and Controller tree. If errors happen during the process, the
Application chooses the correct error Handler. Once the correct Handler is
found, it is wrapped with the correct Trimmings and the Response is retrieved.
The Application then returns the Response to the Server.

trim provides several customizable types. All of these are interfaces which can
be satisfied except Application. All customizable types have default
implementations in subpackages named after themselves. The types are:

* Application: Custom Applications already have Trimmings and Controllers
  attached and can be attached to other Applications or run from Servers.
* Trimming: Custom Trimmings provide a functionality that would normally be
  included in a Controller but is not related to the actual logic of the
  Controller.
* Controller: Custom Controllers allow functionalities to be delivered by the
  web Application.
* Handler: Custom Handlers help decide the proper behavior in error situations
  or special circumstances.
* Response: Custom Responses allow more specialized and powerfulr messages to be
  delivered back to the client which made the Request.

More detailed discussions and more types are discussed in the API documentation,
examples, and tutorial.
