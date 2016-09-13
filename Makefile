BIN_DIR=./bin
VERSION=0.1.0

build-cross:
	GOOS=linux GOARCH=amd64 go build -o $BIN_DIR/goodluck_chatwork_linux_amd64-$VERSION
	GOOS=darwin GOARCH=amd64 go build -o $BIN_DIR/goodluck_chatwork_darwin_amd64-$VERSION
	GOOS=windows GOARCH=amd64 go build -o $BIN_DIR/goodluck_chatwork-$VERSION.exe
