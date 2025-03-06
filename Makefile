run:
	go run .

test:
	go test -v ./... -race

.PHONY: run test