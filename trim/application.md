## Application

Application is the portion of a trim program that encapsulates a common set of
functionalities. These are routed based on subdomain.

```
NewApplication(String) Application

Application
  AddController(Controller) Error
  AddDecorator(Decorator)
  Decorators() List<Decorator>
  RouteToController(Request) (Controller, Arguments)
  SetHandle404(Handler)
  SubDomain() string
```

NewApplication creates an Application at the given subdomain.

AddController binds a Controller to the Application. An Error is returned if the
path the Controller binds to has already been used in the Application.
AddDecorator and Decorator fulfill the Decorated interface. RouteToController
routes a Request to a Controller by returning the matched Controller along with
the Arguments in the path which caused it to match. SetHandle404 sets the 404
handler for the Controller. This Handler gets called when no controller is
matched. SubDomain returns the subdomain the Application is bound to.

A common way to use Application is returned from some factory function. This
allows the Application to encapsulate its own creation and lets the Server view
it as a black box.

```
func APIApplication() *trim.Application {
	app := trim.NewApplication("api")
	app.AddDecorator(decorators.AllowDecorator([]string{"GET"}))
	app.AddDecorator(decorators.CacheDecorator(time.Hour))
	app.SetHandle404(handlers.HandleJSON404)
	app.AddController(NewController1())
	app.AddController(NewController2())
	app.AddController(NewController3())
	return app
}
```

The subdomain binding supports a small degree of freedom in how the subdomain is
declared. Instead of putting a string literal, a variable in '<>' characters can
be used. This matches any subdomain part into the variable in the Request. If
the matching needs to expand beyond the subdomain delimiters, a period can be
added before the closing angle bracket. Some examples are

* literal: Matched if the subdomain is 'literal'.
* \<name\>: Matches any subdomain containing no periods into the name variable.
* \<name\>.account: Matches any subdomain part containing no periods into the
  name variable as long as it is followed by '.account'.
* \<userApplication .\>: Matches any subdomain into the userApplication
  variable, even across periods.
