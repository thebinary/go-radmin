build: *.go ../libradmin/*.go ../libradmin/stats/*.go
	mkdir -p build
	GOOS=darwin GOARCH=amd64 go build -o build/goradmin-darwin-amd64 *.go
	GOOS=linux GOARCH=amd64 go build -o build/goradmin-linux-amd64 *.go

clean:
	rm -rf build
