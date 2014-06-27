backend
=======

Recruitment Backend

Servers
========
TBD: Google App Engine
A little [background](http://blog.golang.org/go-and-google-app-engine) on the Go runtime for App Engine.

A few things to note:

* Go doesn't need to be installed on the system for development; the App Engine SDK is self-contained. (Unless you want to download 3rd party libraries automatically, which we do.)
* App Engine programs are run on a single thread. Goroutines and channels are supported, but they can't be parallelized. Multiple goroutines may run concurrently if the current routine is waiting on an external resource. (Multi-threaded support is expected in the future.)
* The source is uploaded to the App Engine and compiled there. (Go is the only compiled language on the App Engine.)

Build Instructions
==================

* Install [Go](http://golang.org/doc/install)
  
  ```
  brew install go
  ```
  
* Install [Google App Engine SDK for Go](https://developers.google.com/appengine/downloads) (Don't forget to add to PATH.)

* Create working dir

  ```
  mkdir backend-api
  cd backend-api
  mkdir src bin pkg
  cd src
  ```

* Clone this repo
  
  ```
  git clone repo_name_here
  ```

* Get Dependencies (Only needed the first time unless you add a new dep)

  ```
  go get ./...
  ```

* Install

  ```
  go install ./...
  ```

* Run

  Prod:
  ```
  ./bin/backend --conf=../backend-deploy/conf/prod/prod.conf
  ```
  Dev:
  ```
  ./bin/backend --conf=../backend-deploy/conf/dev/dev.conf
  ```

API Documentation
==================
Documentation is available under dev configuration, located at /api-docs (when running locally: http://localhost:8080/api-docs)

