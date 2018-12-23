SHELL = bash
GOLIST=$(shell go list ./... | grep -v '/vendor/')

fmt:
	go fmt $(GOLIST)

vet:
	go vet $(GOLIST)

test:
	go test -v $(GOLIST)
