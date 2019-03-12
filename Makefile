GOPHER = 'ʕ◔ϖ◔ʔ'

.PHONY: build

hello:
	@echo Hello go project ${GOPHER}

build:
	@go build -a -o rundoc main.go
