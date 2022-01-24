VERSION = 1.0.1
build:
	@mkdir -p bin
	@GOOS=darwin go build -ldflags="-X 'main.Version=${VERSION}'" -o bin/countdown-osx
	@GOOS=linux go build -ldflags="-X 'main.Version=${VERSION}'" -o bin/countdown-linux
	@GOOS=windows go build -ldflags="-X 'main.Version=${VERSION}'" -o bin/countdown.exe