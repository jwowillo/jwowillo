# trim

* [API Documentation and Examples](https://godoc.org/github.com/jwowillo/trim)
* [Source](https://github.com/jwowillo/trim)

<center>![Turkey Leg]({{ static }}/trim/leg.png)</center>

Package `trim` is a web-application framework which allows efficient development
by removing boilerplate and providing commonly used functionality.

```
func HelloWorld(_ trim.Request) trim.Response {
	return response.NewPlainText(trim.Body("Hello World"), trim.CodeOK)
}

func main() {
	app := application.New()
	app.AddController(controller.NewFunc("/", handler.NewFunc(HelloWorld)))
	server.New(host, port).Serve(app)
}
```

## Description

`trim` works by removing unrelated logic from `Application`s through the use of
`Trimmings`. `Trimming`s are decorators which wrap a specific piece of
functionality in other necessary, but not directly associted, functionality. For
example, a piece of functionality that returns an expensive operation could have
a cache trimming which remembers results. Caching isn't directly related to the
expensive operation but is good to have for avoiding recomputation.

This concept is extended through the entire framework. For example, application
nesting is made easy in `trim`. `Application`s, as defined in the `application`
package, are groups of related functionality that exist at a subdomain or
base-path and are organized into a tree structure. `Trimming`s applied to an
`Application` in the tree affect all child `Application`s and `Controller`s. At
a small level, a `Trimming` which only allows GET requests could be applied to a
static-file `Application` which will force all attached `Controller`s to only
allow GET requests. At a higher level, a context injecting `Trimming` could be
applied to the root `Application` injecting a context into every `Application`
and `Controller` below it.

`Trimming`s allow the repetitive tasks of web-development to be easily separated
and refactored into their own specific functionality. These `Trimming`s can then
by more broadly applied. The actual code of substance is also more effectively
able to convey its purpose without the unrelated logic. Premade `Trimming`s can
also be easily provided allowing the functionality to be written once in a
library and used anywhere.

## Architecture

`trim` is organized into several modular packages, including:

* `url`
* `request`
* `response`
* `handler`
* `trimming`
* `controller`
* `application`
* `server`
* `logger`

Each of these implements an interface or extends a type defined in the `trim`
package. The interaction of these interfaces and types is explained according to
the digram:

This shows that, from the top down, `Server`s serve `Application`s. These
`Application`s have associated `Controller`s. `Application`s and `Controller`s
can have `Trimming`s which apply to all of the `Application`'s `Controller`s or
just the `Controller` itself. `Controller`s are `Handler`s with associated
paths. Specific `Application`s and `Controller`s are requested through
`trim.URLs`. The `Server` uses `Logger`s to log interactions which are composed
of a `Request` to a `Handler` and the `Handler`'s `Response`. A model of the
structure of the framework is:
