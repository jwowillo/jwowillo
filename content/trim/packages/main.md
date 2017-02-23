# Packages

`trim` is modularly broken into several packages. The root package contains only
examples. The rest contain the actual functionality. These are:

* `application`: Package `application` provides a basic `Application` which is
  a collection of `Controller`s and other nested `Application`s. More complex
  `Application`s are also provided.
* `controller`: Package `controller` contains the `Controller` interface which
  is a `handler.Handler` with a path.
* `handler`: Package `handler` contains the `Handler` interface which describes
  an interaction started by a `Request` and finished by a `Response`.
* `logger`: Package `logger` contains several `Logger` implementations for
  documenting `Server` interactions in different ways.
* `request`: Package `request` contains a `Request` type which is sent from a
  client to a `handler.Handler` and a `Builder` for `Request`s.
* `response`: Package `response` contains several `Response` implementations
  which are given during interactions with `Request`s.
* `server`: Package `server` contains a `Server` which serves
  `application.Application`s.
* `trimming`: Package `trimming` consists of implementations of `Trimming`s
   which are pieces of extra functionality for `handler.Handler`s along with a
  `Trimmed` interface which types that have `Trimming`s satisfy.
* `url`: Package url contains a `URL` primitive along with a `Builder` for
  `URL`s. A `Matcher` is also provided that makes `Match`s from a `URL` and a
  `Pattern` that the `URL` matches.
