lint:
	golangci-lint run --config .golangci-lint.yml ./...

test:
	go test -race -v -coverprofile="c.out" ./...
	go tool cover -func="c.out"

scan:
	gosec -no-fail -fmt sarif -out security.sarif ./...

bench:
	go test -bench=. ./...