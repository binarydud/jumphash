test:
	go get github.com/stretchr/testify/assert
	go test ./server
build:
	go get github.com/labstack/echo
	go get github.com/tylerb/graceful
	go build

run: build
	./jumphash
