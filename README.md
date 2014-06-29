# Recruitment_Backend

## Servers

TBD: Google App Engine
A little [background](http://blog.golang.org/go-and-google-app-engine) on the Go runtime for App Engine.

A few things to note:

* Go doesn't need to be installed on the system for development; the App Engine SDK is self-contained. (Unless you want to download 3rd party libraries automatically, which we do.)
* App Engine programs are run on a single thread. Goroutines and channels are supported, but they can't be parallelized. Multiple goroutines may run concurrently if the current routine is waiting on an external resource. (Multi-threaded support is expected in the future.)
* The source is uploaded to the App Engine and compiled there. (Go is the only compiled language on the App Engine.)

## Local Build Instructions


* Install [Go](http://golang.org/doc/install) (Don't forget to add to PATH.)
  
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
  
  If your github account has SSH set up:
  ```
  git clone git@github.com:SellJamHere/Recruitment_Backend.git src
  ```

  Otherwise:
  ```
  git clone https://github.com/SellJamHere/Recruitment_Backend.git src
  ```

* Set GOPATH

  ```
  cd ../
  export GOPATH=`pwd`
  ```

* Get Dependencies (Only needed the first time unless you add a new dep)

  ```
  go get ./...
  ```

* Run

  ```
  goapp serve ./src
  ```

## Important Notes

### Workspace Structure

Your workspace should be structured as follows:

```
backend-api/
  |bin/
  |pkg/
  |src/
    |app.yaml
    |backend/         <-- this source directory
    |github.com/      <-- libs from github
    |other_site.com/  <-- libs from another site
```

`app.yaml` contains runtime information for App Engine. It lives one directory above the source files. Source files you write are in the `backend/` directory. When running `go get`, code will be downloaded into an appropriate directory, ie github libraries to `github.com`. 

### 3rd Party Libs and .gitignore

Downloaded 3rd party libs need to be ignored. Each directory containing 3rd party libs must be added to `.gitignore`.
