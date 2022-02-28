.PHONY: .build
.build: BIN?=./bin
.build:
	go build -o $(BIN) ./cmd/re-finder

# build app
.PHONY: build
build: .build