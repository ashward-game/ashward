.PHONY: all test clean

setup:
	go mod download

quick_test:
	go test -count=1 -p=1 ./...

test:
	test -z "$(go fmt ./...)"
	go vet ./...
	./test_coverage.sh

coverage:
	go tool cover -html=profile.cov
