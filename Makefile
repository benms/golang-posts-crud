SHELL := /bin/bash

run-server:
	go run . -env dev
run-migrations:
	cd migrations && go run *.go up
get-migrations:
	cd migrations && go run *.go version
run-tests:
	source .env_test && cd migrations && go run *.go init; go run *.go up; cd ../ && go test -p 1 ./services/...