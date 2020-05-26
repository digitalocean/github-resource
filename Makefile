mkfile := $(abspath $(lastword $(MAKEFILE_LIST)))
dir := $(dir $(mkfile))

.PHONY: gotidy
gotidy:
	go mod tidy && go mod vendor
