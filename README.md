![Go Report Card](https://goreportcard.com/badge/github.com/joaoepj/learning-go)
![CircleCI](https://img.shields.io/circleci/build/gh/joaoepj/learning-go)

# learningo-go

## Quick start

```
make build
```
```
./learning-go
NAME:
   learning-go - a repository to learn the go programming language.

   Go projects are made of modules, and modules are made of packages.
   This is a simple project designed to contain a single module and some packages.
   the learning-go module uses the urfave/cli go library to provide a CLI allowing
   a single executable to test different snippets of code and organize your 
   own packages.

USAGE:
   learning-go command [command options]
   learning-go [global options]

VERSION:
   None-None

COMMANDS:
   problems, p  Some HackerRank problem solutions implemented in golang.
   add, a       add a task to the list
   snmp, s      Test the SNMP library gosnmp
   misc, m      Miscelaneous code
   help, h      Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help (default: false)
   --version, -v  print the version (default: false) 
```

# How to test the code
```
make test
```
or

```
make verbose_test
```

Golang provides the "testing" package if you need to speed up the writing of your unit tests.
After you write test functions just call go-test tool and it will find and run yout test code. I have provided a Makefile just to document some usefull go tool commands you can use to test your code.

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
