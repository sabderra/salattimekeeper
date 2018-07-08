# Salat Time Keeper
Go based REST server for calculating Salat times.

The Prayer time calculations are based on the Python code from http://praytimes.org/. The license can be found [here](License_praytime.org). Some of the algorithms have been updated to use those from United States Naval Observatory (USNO)  see references below.

## Install

Install go version 1.10.3 or later visit https://golang.org/doc/install for instructions.

Setup GOROOT and GOPATH

```
export GOROOT=/usr/local/go
export GOPATH=$HOME/workspaces/go
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
```

Install build dependencies. This project used [gometalinter](https://github.com/alecthomas/gometalinter) for linting and [Dep](https://github.com/golang/dep) for dependency management. Addition dependencies will be installed when building the project.

```
go get -u github.com/alecthomas/gometalinter
go get -u github.com/golang/dep/cmd/dep
```

Retrieve and build this project

```
go get github.com/sabderra/salattimekeeper
cd $GOPATH/src/github.com/sabderra/salattimekeeper
make
```

## References:
* http://praytimes.org/calculation
* http://aa.usno.navy.mil/faq/docs/JD_Formula.php
* http://aa.usno.navy.mil/faq/docs/SunApprox.php
* https://www.moonsighting.com/how-we.html
