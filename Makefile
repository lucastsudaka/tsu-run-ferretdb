build:
	GOOS=linux GOARCH=amd64 go build   ./tsu-run-ferretdb.go

buildAndRun:
	GOOS=linux GOARCH=amd64 go build   ./tsu-run-ferretdb.go && ./tsu-run-ferretdb