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
	$(GO) build -o learning-go ./main.go

run:
	$(GO) run ./main.go

fmt:
	$(GO) fmt ./lg_misc
	$(GO) fmt ./lg_problems
	$(GO) fmt ./lg_snmp