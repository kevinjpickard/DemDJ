[![Coverage Status](https://coveralls.io/repos/github/kevinjpickard/DemDJ/badge.svg?branch=master)](https://coveralls.io/github/kevinjpickard/demdj?branch=master)
[![Build Status](https://travis-ci.org/kevinjpickard/demdj.svg?branch=master)](https://travis-ci.org/kevinjpickard/demdj)
[![Windows Build status](https://ci.appveyor.com/api/projects/status/5grajhscy718wc02?svg=true)](https://ci.appveyor.com/project/kevinjpickard/demdj)
# DemDJ
DemDJ: Music By the People, For the People

## Dependencies
This repo uses [go modules](https://github.com/golang/go/wiki/Modules#how-to-use-modules) to manage dependencies. This should be automatic. However, you must be running Go 1.11 or later. In addition, until the release of Go 1.12 you need to enable it:
```
export GO111MODULE='on'
```
To add a dependency, do the following anywhere in the project repo:
```
go get <dependency>
go mod vendor
```

## Testing
Test are run using the [Ginko Testing Framework](https://github.com/onsi/ginkgo). To run all tests, run:
```
go get -u github.com/onsi/ginkgo/ginkgo  # installs the ginkgo CLI
go get -u github.com/onsi/gomega/...     # fetches the matcher library
ginko
```
or 
```
go test
```
When writing tests, please use the [Gomega Matcher Library](https://github.com/onsi/gomega) ([Gomega documentation](http://onsi.github.io/gomega/#provided-matchers))
