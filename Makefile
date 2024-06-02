include .env.golang

default: run_test_server

run_test_server:
	go run main.go -config ./config.toml

build:
	go build -o pacgen main.go

