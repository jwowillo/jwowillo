# jwowillo

Documentation for projects at https://github.com/jwowillo.

## Description

`jwowillo` contains documentation for several projects. Descriptions of these
 are:

* `alexaharvester`: Package `alexaharvester` defines an Alexa skill which allows
  `harvester` to be used with voice.
* `apidrone`: Package `apidrone` contains a web-application which allows a
  crazyflie drone to be controlled from an API.
* `harvester`: Package `harvester` is a web-application which returns
  descriptions to users based on a set of URLs and tags entered by them and
  provided by machine-learned recommendations.
* `landgrab`:
* `md2web`: Package `md2web` offers a simple way to turn a nested directory of
  markdown files into a web-application.
* `pack`: Package `pack` contains general thread-safe data-structure
  implementations along with interfaces the data-structures satisfy.
* `seeds`: Package `seeds` contains common project starters along with scripts
  to configure them.
* `trim`: Package `trim` is a web application framework which allows efficient
  development by removing boilerplate and providing commonly used functionality.

This website was created using `trim` and `md2web`. `trim` simplifies adding
applications at subdomains and `md2web` allows for easy displaying of
markdown-file directories. `jwowillo` itself simplifies the deployment of the
static files for the subdomain applications and markdown-files.

## Subdomains

* http://harvester.jwowillo.com: `Harvester` application as described in
  `Harvester` project.
* http://sample.jwowillo.com: Sample application for the `web` seed in `seeds`.
