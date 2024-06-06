# binary name
APP_NAME=pismo

# set additional test parameters
TEST_PARAMS=-v

all: clean fmt test build

build:
	@go build -o build/$(APP_NAME) cmd/$(APP_NAME)/main.go

test:
	@go test $(TEST_PARAMS) ./...

fmt:
	@go fmt ./...

clean:
	@go clean
	@rm -f $(APP_NAME)
