BIN_DIR=./bin
VERSION=0.1.2

build-cross:
	rm -rf ${BIN_DIR}/*
	GOOS=linux GOARCH=amd64 go build -o ${BIN_DIR}/goodluck_chatwork-${VERSION}
	GOOS=darwin GOARCH=amd64 go build -o ${BIN_DIR}/goodluck_chatwork_darwin-${VERSION}
	GOOS=windows GOARCH=amd64 go build -o ${BIN_DIR}/goodluck_chatwork-${VERSION}.exe
