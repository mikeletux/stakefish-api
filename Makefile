build:
	 go build -o bin/stakefish-api github.com/mikeletux/stakefish-api/cmd/main

test:
	go test ./...

lint:
	go vet ./...

run:
	go run github.com/mikeletux/stakefish-api/cmd/main