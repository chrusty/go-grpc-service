go-grpc-service
===============

A template GRPC service in Go

Features
--------

* Adheres to this [GoLang project layout](https://github.com/golang-standards/project-layout)
* Uses a [Makefile](Makefile) to trigger build and test steps
* Uses versioned proto definitions (/api/v*)
* Generated GoLang code is also versioned, importable by other projects (under the /pkg directory)
* Service logic and packages are kept under /internal (to prevent other projects from depending on them in ways we didn't intend them to)

Not covered
-----------

* GRPC Middleware
* K8S probes
* Config
* Actual implementations of interfaces
* Tests
