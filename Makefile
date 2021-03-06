PACKAGE=subscription-api
.PHONY: test run build build_docker clean
test:
	go test ./...
run:
	go run main.go
build:
	go build -o $(PACKAGE) .
build_docker:
	docker build -t $(PACKAGE) .
clean:
	rm -f ./$(PACKAGE)