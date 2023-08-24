
.PHONY: build fmt ci

LDFLAGS += -X "main.LearningGoGitHash=$(shell git rev-parse HEAD)"
LDFLAGS += -X "main.LearningGoGitBranch=$(shell git rev-parse --abbrev-ref HEAD)"

GOPATH ?= $(go env GOPATH)
GOBIN ?= $(GOPATH)/bin
GO := go




test:
	$(GO) test -skip="TestCuriosity/Curiosity=2" ./lg_misc
	$(GO) test ./lg_problems
	$(GO) test ./lg_algorithms
verbose_test:
	$(GO) test -v ./lg_misc ./lg_problems ./lg_algorithms

function_test:
	$(GO) test -v -run "TestCuriosity" ./lg_misc

benchmark_test:
	$(GO) test -v  -benchmem -bench ^Benchmark ./lg_algorithms

case_function_test:
	$(GO) test -v -run "TestCuriosity/Curiosity=1" ./lg_misc

build:
	$(GO) build -ldflags '$(LDFLAGS)' -o bin/learning-go ./main.go


# If the first argument is "run"...
ifeq (run,$(firstword $(MAKECMDGOALS)))
  # use the rest as arguments for "run"
  RUN_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
  # ...and turn them into do-nothing targets
  $(eval $(RUN_ARGS):;@:)
endif

run:
	$(GO) run ./main.go $(RUN_ARGS)

fmt:
	$(GO) fmt -x ./lg_misc
	$(GO) fmt -x ./lg_problems
	$(GO) fmt -x ./lg_algorithms

ci:
	$(MAKE) fmt
	$(MAKE) build
