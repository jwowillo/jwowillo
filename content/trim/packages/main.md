# Packages

`trim` is modularly broken into several packages. The root package contains
interfaces, general types, and examples only. The rest contain the actual
functionality. These are:

* `application`: Package `application` provides a basic `Application` which is
  a collection of `Controller`s and other nested `Application`s. More complex
  `Application`s are also provided.
* `controller`: Package `controller` contains a `trim.Controller` implementation
  which creates `trim.Controller`s from a `trim.Handler`, path, and
  `trim.Trimming`s along with an embeddable type for `trim.Controller`s with no
  `trim.Trimming`s.
* `handler`: Package `handler` contains implementations of the `trim.Handler`
  interface.
* `logger`: Package `logger` contains several `trim.Logger` implementations for
  documenting `trim.Server` interactions in different ways.
* `request`: Package `request` contains a `trim.Request` implementation which is
  sent from a client to a `trim.Handler` and a `Builder` for `Request`s.
* `response`: Package `response` contains several `trim.Response`
  implementations which are given during interactions with `trim.Request`s.
* `server`: Package `server` contains a `trim.Server` implementation which
  serves `trim.Application`s.
* `trimming`: Package `trimming` consists of implementations of `trim.Trimming`s
   which are pieces of extra functionality for `trim.Handler`s.
* `url`: Package url contains a `trim.URL` implementation along with a `Builder`
  for `URL`s. A `Matcher` is also provided that makes `Match`s from a `trim.URL`
  and a `Pattern` that the `trim.URL` matches.
