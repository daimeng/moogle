DIR=./cmd/moogle
BINARY_NAME=moogle

all: build
build:
	go build -o $(BINARY_NAME) -v $(DIR)
clean:
	go clean
	rm -f $(BINARY_NAME)
run:
	go run $(DIR)
run_debug:
	go run $(DIR) -step=true
