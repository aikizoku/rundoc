GOPHER = 'ʕ◔ϖ◔ʔ'

.PHONY: build

hello:
	@echo Hello go project ${GOPHER}

file_to_bin:
	@statik -src=src/template

build:
	@go build -a -o rundoc main.go
