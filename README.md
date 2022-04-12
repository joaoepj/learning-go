# learningo-go

Go projects are made of modules, and modules are made of packages.

## How to add a module
$ mkdir learning-go

$ cd learning-go

$ git init

$ go mod init github.com/joaoepj/learning-go

... write code

How to output the executable to build/bin when issuing `go build`?

Set GOBIN environment variable


## How to add package

$ mkdir lg_snmp

$ cd lg_snmp

... write code

public functions must start with capital letters

$ go build (test if package builds)

import package in other packages, build module