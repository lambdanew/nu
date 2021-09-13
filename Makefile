.PHONY: all
all: test

.PHONY: test
test:
	cd pkg/parser && go test
	cd pkg/runtime && go test
	
.PHONY: test_coverage
test_coverage:
	go test `go list ./... | grep -v 'hack\|cmd'` -coverprofile=coverage.txt -covermode=atomic
