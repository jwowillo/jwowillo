## Controller

Controllers encapsulate a single functionality. They are routed based on path.

```
interface Controller
  Decorated
  Path() String
  Handle(Request) Response
```

Implementing a Controller involves adding methods in the Decorated interface,
setting a path the Controller is bound to, and specifying how the Controller
Handles turning a Request to a Response.

The Decorated methods in Controller times are empty a lot of times. Because of
this, a Controller without Decorators can embed a BareController and not have to
provide custom Decorated implementations.

```
type SampleController struct {
	trim.BareController
}
```

Even if BareController is embedded, Decorated methods can be overrided. A common
case for this is if a Controller never needs to add Decorators, but does have
default decorators.

```
func (c *SampleController) Decorators() []trim.Decorator {
	return []trim.Decorator{decorators.AllowDecorator([]string{"GET"})}
}
```

Variable paths are also supported. These are done by placing variable names
inside '<>' along with an optional forward slash to match across path
components.

Examples are

* /\<name\>: Match a single name after a forward slash until a forward slash.
* /\<path /\>: Match into path across forward slashes after an initial forward
  slash.
* /user/\<name\>: Match a name after a '/user/' string.

Finally, the Handle method is meant to interpret a Request in some way to create
a Response. A sample implementation is

```
func (c *SampleController) Handle(r *trim.Request) trim.Response {
	path := r.Path()
	return responses.NewPackagedJSONResponse(
		responses.AnyMap{"requestPath": path},
		http.StatusOK,
	)
}
```
