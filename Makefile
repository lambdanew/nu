.PHONY: all
all: test

.PHONY: gomod
gomod:
	go mod verify
	go mod tidy

.PHONY: test
test:
	cd pkg/common && go test
	cd pkg/crypto/sr25519 && go test
	cd pkg/crypto/secp256k1 && go test
	cd pkg/crypto/ed25519 && go test
	cd pkg/crypto && go test
	cd pkg/parser && go test
	cd pkg/runtime && go test
	
.PHONY: test_coverage
test_coverage:
	go test `go list ./... | grep -v 'hack\|cmd'` -coverprofile=coverage.txt -covermode=atomic
