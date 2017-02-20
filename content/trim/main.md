# trim

* [API Documentation and Examples](https://godoc.org/github.com/jwowillo/trim)
* [Source](https://github.com/jwowillo/trim)

<center>![Turkey Leg]({{ static }}/trim/leg.png)</center>

`trim` is a web application framework which allows efficient development by
removing boilerplate and providing commonly used functionality.

```
func HelloWorld(_ *request.Request) response.Response {
	return response.NewPlainText(
		response.Body("Hello World"),
		response.CodeOK,
	)
}

func main() {
	app := application.New()
	app.AddController(controller.NewFunc("/", handler.NewFunc(HelloWorld)))
	server.New(host, port).Serve(app)
}
```

## Description

`trim` works by removing unrelated logic from applications through the use of
"tirmmings". Trimmings are decorators which wrap a specific piece of
functionality in other necessary, but not directly associted, functionality. For
example, a piece of functionality that returns an expensive operation could have
a cache trimming which remembers results. Caching isn't directly related to the
expensive operation but is good to have for avoiding recomputation.

This concept is extended through the entire framework. For example, application
nesting is made easy in `trim`. Applications are groups of related functionality
that exist at a subdomain or base path and are organized into a tree structure.
Trimmings applied to an application in the tree affect all child applications
and controllers. At a small level, a trimming which only allows GET requests
could be applied to a static file serving application which will force all
attached controllers to only allow GET requests. At a higher level, a context
injecting trimming could be applied to the root application, injecting a context
into every application and controller below it.

Trimmings allow the repetitive tasks of web development to be easily separated
and refactored into their own specific functionality. These trimmings can then
by more broadly applied. The actual code of substance is also more effectively
able to convey its purpose without the unrelated logic. Premade trimmings can
also be easily provided allowing the functionality to be written once in a
library and used anywhere.

## Architecture

`trim` is organized into several modular packages, including:

* url
* request
* response
* handler
* trimming
* controller
* application
* server
* logger

Each of these serves a specific purpose and is documented in the links above.
The packages interact with eachother according to the below diagram:

This shows that, from the top down, servers serve applications. These
applications are composed of other applications and have associated controllers. 
Applications and controllers can have trimmings which apply to the entire
application subtree. Controllers are handlers with associated paths. Specific
applications and controllers are requested through urls. The server uses loggers
to log interactions, which are composed or a request to a controller and the
controller's response. A model of the structure of the framework is:
