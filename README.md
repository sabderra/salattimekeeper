# Salat Time Keeper
Go based REST server for calculating Salat times.

--- NOTE still a work-in-progress ---

The Prayer time calculations are based on the Python code from http://praytimes.org/. The license can be found [here](License_praytime.org). Some of the algorithms have been updated to use those from United States Naval Observatory (USNO)  see references below.

## Install

Install go version 1.12.6 or later visit https://golang.org/doc/install for instructions.

Setup GOROOT and GOPATH

```
export GOROOT=/usr/local/go
export GOPATH=$HOME/workspaces/go
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
```

Install build dependencies. This project used [golangci-lint](https://github.com/golangci/golangci-lint) for linting. 

```
curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $(go env GOPATH)/bin v1.17.1
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
