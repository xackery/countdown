build:
	@mkdir -p bin
	@GOOS=darwin go build -o bin/countdown-osx
	@GOOS=linux go build -o bin/countdown-linux
	@GOOS=windows go build -o bin/countdown.exe