[![Go Report Card](https://goreportcard.com/badge/github.com/joaoepj/learning-go)](https://goreportcard.com/report/github.com/joaoepj/learning-go)
![CircleCI](https://img.shields.io/circleci/build/gh/joaoepj/learning-go)

# learningo-go

Go projects are made of modules, and modules are made of packages.
This is a simple project designed to contain a single module and some packages.
the learning-go module uses the urfave/cli go library to provide a CLI allowing
 a single executable to test different snippets of code and organize your 
 own packages. 


## How to add code to git repo

create or modify file

... write code

... check if it builds

git add .

git commit -m "commit message"

git push


## How to add a module
$ mkdir learning-go

$ cd learning-go

$ git init

$ go mod init github.com/joaoepj/learning-go (create the module - go.mod file )

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
