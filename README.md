# jwowillo

jwowillo is a location to host personal projects. The root web application
serves project documentation. Personal projects are located at subdomains.

## Adding a Project

In `projects.go`, add the projects Github URL as a key and constructor as a
value. When jwowillo is reinstalled, it will automatically be added.

## Repository

The repository is located at github.com/jwowillo/jwowillo.

## Installing

Run `make`.

## Running

Run `jwowillo`. The command accepts optional flags `--url` and `--port`.
