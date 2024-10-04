BINARY_NAME=amd-redfish-exporter

.PHONY: build run test clean

build:
	cd api; make; cd ../
	go build -o $(BINARY_NAME)
	cd api; make clean; cd ../

run: build
	./$(BINARY_NAME)

test:
	go test -v ./...

clean:
	go clean
	rm -f $(BINARY_NAME)

integration-test: build
	./integration_test.sh
