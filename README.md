![Go Report Card](https://goreportcard.com/badge/github.com/joaoepj/learning-go)
![CircleCI](https://img.shields.io/circleci/build/gh/joaoepj/learning-go)

# learningo-go

## Quick start

#### Compiling
```
make build
```

#### Getting some help
```
bin/learning-go -h
NAME:
   learning-go - a repository to learn the go programming language.

USAGE:
   learning-go command [command options]
   learning-go [global options]

VERSION:
   main-03e4c88cd396f232d53f126f598a81dfe7c90fda

COMMANDS:
   problems, p  Some HackerRank problem solutions implemented in golang.
   misc, m      Miscelaneous code
   help, h      Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help (default: false)
   --version, -v  print the version (default: false)
```

#### Running a problem
```
bin/learning-go p
NAME:
   learning-go problems - Some HackerRank problem solutions implemented in golang.

USAGE:
   learning-go problems command [command options] [arguments...]

COMMANDS:
   arraymanipulation, am     Array Manipulation Problem
   bigsum, bs                A Very Big Sum Problem
   catmouse, cm              Cats and a Mouse Problem
   checkmagazine, cm2        Check Magazine Problem
   climbinglearderboard, cl  Climbing the Leaderboard Problem
   comparethetriplets, ctt   Compare The Triplets Problem
   designerpdfviewer, dpv    Designer Pdf Viewer Problem
   formingmagicsquare, fms   Forming A Magic Square Problem
   hourglasssum, hgs         Hourglass Sum Problem
   hurdlerace, hr            Hurdle Race Problem
   letfrotation, lr          Left Rotation Problem
   migratorybirds, mb        Migratory Birds Problem
   minimumbribes, mb2        Minimun Bribes Problem
   minimumswaps2, ms2        Minimun Swaps 2 Problem
   minmaxsum, mms            Min-Max Sum Problem
   missinginteger, mi        Missing Integer Problem
   staircase, s              Staircase Problem
   twostrings, ts            Two String Problem
   countswaps, cs            Count Swaps Problem
   twosum, tw                Two Sum Problem
   help, h                   Shows a list of commands or help for one command

OPTIONS:
   --help, -h  show help (default: false)
```

```
bin/learning-go p ms2
minimum-swaps2.go:28: [7 2 3 4 5 6 1]
minimum-swaps2.go:28: [7 1 3 4 5 6 2]
minimum-swaps2.go:28: [7 1 3 2 5 6 4]
minimum-swaps2.go:28: [7 1 3 2 4 6 5]
minimum-swaps2.go:28: [7 1 3 2 4 5 6]
minimum-swaps2.go:33: 5
minimum-swaps2.go:28: [1 3 2 4 5 6 7]
minimum-swaps2.go:28: [1 3 5 4 2 6 7]
minimum-swaps2.go:28: [1 3 5 2 4 6 7]
minimum-swaps2.go:33: 3
minimum-swaps2.go:28: [2 1 3 4 5]
minimum-swaps2.go:28: [2 3 1 4 5]
minimum-swaps2.go:28: [2 3 4 1 5]
minimum-swaps2.go:33: 3
minimum-swaps2.go:28: [4 2 3 1]
minimum-swaps2.go:28: [4 3 2 1]
minimum-swaps2.go:28: [4 3 1 2]
minimum-swaps2.go:33: 3
```


## How to test the code
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

## Tips

Search for issues at Github. Go to Issues and type text below in search field.

is:open is:issue archived:false language:go label:"good first issue" 
