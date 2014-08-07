.PHONY: all binary build default test 

default: run

all: run

get:
	go get ./...

test: get 
	go test ./... -v 


files=$(wildcard 2014*/*/*.go) 

run: $(files)

$(files): get
	go run $@


