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
$ bin/learning-go -h
NAME:
   learning-go - a repository to learn the go programming language.

USAGE:
   learning-go command [command options]
   learning-go [global options]

VERSION:
   main-b42c6ff0533a37cf47f521903bc656e995b4ce41

COMMANDS:
   problems, p  Some HackerRank problem solutions implemented in golang.
   misc, m      Miscelaneous code
   help, h      Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help (default: false)
   --version, -v  print the version (default: false)
```


#### Listing problems
```
$ bin/learning-go p
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
   pickingnumbers, pn        Picking Numbers Problem
   help, h                   Shows a list of commands or help for one command

OPTIONS:
   --help, -h  show help (default: false)
```

#### Running a problem
```
$ bin/learning-go p pn
PickingNumbers1
0
0
PickingNumbers2
sbl:  [1 2 2 1]
sbl:  [4 5]
PickingNumbers3
map[4:[4 5 3 3 1]]
map[4:[4 5 3 3 1] 6:[6 5 3 3 1]]
map[4:[4 5 3 3 1] 5:[5 3 3 1] 6:[6 5 3 3 1]]
map[3:[3 3 1] 4:[4 5 3 3 1] 5:[5 3 3 1] 6:[6 5 3 3 1]]
map[3:[3 3 1] 4:[4 5 3 3 1] 5:[5 3 3 1] 6:[6 5 3 3 1]]
map[1:[1] 3:[3 3 1] 4:[4 5 3 3 1] 5:[5 3 3 1] 6:[6 5 3 3 1]]
5
map[1:[1 1 2 2]]
map[1:[1 1 2 2]]
map[1:[1 1 2 2] 2:[2 2]]
map[1:[1 1 2 2] 2:[2 2]]
map[1:[1 1 2 2] 2:[2 2] 4:[4 4 5 5 5]]
map[1:[1 1 2 2] 2:[2 2] 4:[4 4 5 5 5]]
map[1:[1 1 2 2] 2:[2 2] 4:[4 4 5 5 5] 5:[5 5 5]]
map[1:[1 1 2 2] 2:[2 2] 4:[4 4 5 5 5] 5:[5 5 5]]
map[1:[1 1 2 2] 2:[2 2] 4:[4 4 5 5 5] 5:[5 5 5]]
5
```


#### How to test the code
```
make test
```
or

```
make verbose_test
```

Golang provides the "testing" package if you need to speed up the writing of your unit tests.
After you write test functions just call go-test tool and it will find and run yout test code. I have provided a Makefile just to document some usefull go tool commands you can use to test your code.

#### How to add code to git repo

create or modify file

... write code

... check if it builds

git add .

git commit -m "commit message"

git push


#### How to add package

learning-go module is comprised of golang packages lg_misc and lg_problems. This section explains how to add packages.

$ mkdir lg_newpackage

$ cd lg_newpackage

... write code

public functions must start with capital letters

$ go build (test if package builds)

import package in other packages, build module

#### How to add a module

learning-go its a golang module itself. So that you don't need to add modules to it. But in case you are curious about how to create a golang module, 

$ mkdir learning-go

$ cd learning-go

$ git init

$ go mod init github.com/joaoepj/learning-go (create the module - go.mod file )

... write code

How to output the executable to build/bin when issuing `go build`?

Set GOBIN environment variable


#### Tips

Search for issues at Github. Go to Issues and type text below in search field.

is:open is:issue archived:false language:go label:"good first issue" 


#### How to debug

* Install Go language support extension into Visual Studio
* Add a Launch Configurantion File (usually it lives at /home/learning-go/.vscode/launch.json, see content example below)
* Set the breakpoints, launch code, explore variables, have fun

```
{
    // Use IntelliSense to learn about possible attributes.
    // Hover to view descriptions of existing attributes.
    // For more information, visit: https://go.microsoft.com/fwlink/?linkid=830387
    "version": "0.2.0",
    "configurations": [
        
        {
            "name": "Launch file",
            "type": "go",
            "request": "launch",
            "mode": "debug",
            "program": "/home/learning-go/main.go",
            "args": ["a", "dp"]
        },

    ]
}
```