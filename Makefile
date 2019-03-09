GO=go
GOTOOL=go tool
GORUN=$(GO) run
GOTEST=$(GO) test
GOLINT=$(GO)lint
GOFMT=$(GO)fmt
GODOC=$(GO)doc
run:
	$(GOTEST) ./... -cover
	$(GOLINT) ./...
coverage:
	$(GOTEST) ./... -coverprofile=test/coverage.out
	$(GOTOOL) cover -html=test/coverage.out
godoc:
	$(GODOCK) -http=":8080"