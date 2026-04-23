# Merged statement coverage across the whole module (includes main).
COVERPKG=./...
.PHONY: test cover
test:
	go test ./...

cover:
	go test ./... -coverpkg=$(COVERPKG) -coverprofile=coverage.out
	go tool cover -func=coverage.out | grep '^total'
	@echo "HTML: go tool cover -html=coverage.out"
