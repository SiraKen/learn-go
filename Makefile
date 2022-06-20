NAME = example
BIN := bin/$(NAME)

LDFLAGS := -w \
		   -s

.PHONY: build
build:
	@echo "Building..."
	@go build -o $(BIN)/$(NAME) -ldflags "$(LDFLAGS)"

