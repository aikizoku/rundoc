.PHONY: h server run statik

server:
	go run ./server/main.go

run:
	go run main.go

statik:
	statik -src=./src/template -dest=./src
