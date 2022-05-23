.PHONY: test benchmark

test:
	go test -v ./...

benchmark:
	go test -v ./... -bench . -run ^$$
