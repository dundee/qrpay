
test:
	go test ./...

coverage:
	go test -v -race -coverprofile=coverage.txt -covermode=atomic ./...

coverage-html: coverage
	go tool cover -html=coverage.txt