all: clean compile test

clean:
	@echo "==> Cleaning up previous builds."
	@rm -rf ./bin/godoku

compile:
	@echo "==> Compiling source code."
	@go build -v -o ./bin/godoku $(go list ./... | grep -v /vendor/)
	cp ./bin/godoku ./bin/application

coverage:
	@go test -coverprofile cover.out
	@go tool cover -html=cover.out

deps:
	@echo "==> Downloading dependencies."
	@godep save $(go list ./... | grep -v /vendor/)

fmt:
	@echo "==> Formatting source code."
	@gofmt -w ./

race_compile:
	@echo "==> Compiling source code."
	@go build -v -race -o ./bin/godoku $(go list ./... | grep -v /vendor/)
	cp ./bin/godoku ./bin/application

test: fmt vet
	@echo "==> Running tests."
	@go test -cover $(go list ./... | grep -v /vendor/)
	@echo "==> Tests complete."

vet:
	@go vet $(go list ./... | grep -v /vendor/)

help:
	@echo "clean\t\tremove previous builds"
	@echo "compile\t\tbuild the code"
	@echo "coverage\tgenerate and view code coverage"
	@echo "deps\t\tdownload dependencies"
	@echo "fmt\t\tformat the code"
	@echo "race_compile\tbuild the code with race detection"
	@echo "test\t\ttest the code"
	@echo "vet\t\tvet the code"
	@echo ""
	@echo "default will test, format, and compile the code"

.PNONY: all clean deps fmt help test
