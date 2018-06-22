test:
	go test -v -cover

vet:
	go vet

lint:
	golint -set_exit_status
	gofmt -s -d *.go

