
.PHONY: build fmt ci

LDFLAGS += -X "main.LearningGoGitHash=$(shell git rev-parse HEAD)"
LDFLAGS += -X "main.LearningGoGitBranch=$(shell git rev-parse --abbrev-ref HEAD)"

GOPATH ?= $(go env GOPATH)
GOBIN ?= $(GOPATH)/bin
GO := go


test:
	$(GO) test ./lg_misc
	$(GO) test ./lg_problems
verbose_test:
	$(GO) test -v ./lg_misc ./lg_problems

function_test:
	$(GO) test -v -run "TestCuriosity" ./lg_misc

case_function_test:
	$(GO) test -v -run "TestCuriosity/Curiosity=1" ./lg_misc

build:
	$(GO) build -ldflags '$(LDFLAGS)' -o bin/learning-go ./main.go

run:
	$(GO) run ./main.go

fmt:
	$(GO) fmt -x ./lg_misc
	$(GO) fmt -x ./lg_problems
	$(GO) fmt -x ./lg_snmp

ci:
	$(MAKE) fmt
	$(MAKE) build
