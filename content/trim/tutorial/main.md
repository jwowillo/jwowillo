# Tutorial

<center>![fibcalc]({{ static }}/trim/tutorial/fibcalc.png)</center>

This tutorial will describe how `trim` is used and also explain tools that
simplify `trim` development. A web-application named `fibcalc` will be created
which has a client, API, and static-file server. The static-file server will
return static-files required by the client. The API will return the nth
fibonacci number. The client will render a more friendly way of interacting with
the application.

The completed application can be viewed at http://fibcalc.jwowillo.com. The
source code for the completed application can be viewed at
https://github.com/jwowillo/fibcalc.

## Getting Started

`fibcalc` will have several dependencies:

* [`Go`](https://golang.org/doc/install)
* [`Angular-CLI`](https://github.com/angular/angular-cli)
* [`Node.js`](https://nodejs.org/en/)
* [`NPM`](https://docs.npmjs.com/getting-started/installing-node)
* [`Make`](https://www.gnu.org/software/make/)
* [`Seeds`](https://github.com/jwowillo/seeds)

First, install all these dependencies if needed. Then, create the project with:

```
$ mkdir fibcalc && cd fibcalc && plant $SEED_PATH/web
Package name:
fibcalc
```

`Seeds` will created a sample project and initialize the `Angular` dependencies
according to the web seed documentation. A `Makefile` which can be run with
`make` will be placed at the project root that has several targets. These are:

* `all`: Builds entire project for development.
* `production`: Builds entire project for production.
* `fibcalc`: Builds `Go` part of the project.
* `angular`: Builds `Angular` folder for development.
* `angular_production`: Builds `Angular` folder for production.
* `clean`: Removes `dist` folder.

To view the project in its current state, run:

```
$ make && run_fibcalc
```

and navigate to localhost:5000.

The runnable `run_fibcalc` also accepts optional command line arguments. These
are:

* `--host`: Host the application will listen on.
* `--port`: Port the application will listen on.

Part 1 of the tutorial will give a more thorough overview of the generated
project, how to use the geneated application, and more details about `fibcalc`'s
requirements. Part 2 will demonstrate the creation of the API. Part 3 will show
how a client is made for the API. Part 4 will introduce a simple way to test
`trim` applications.  Throughout the tutorial, it should be seen that `trim`
emphasizes removing
boilerplate and allows the programmer to only write the necessary code.
