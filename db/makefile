BINARY = main.exe

pre:
	go mod init db
	go mod tidy

build-win: 
	go env -w GOOS=windows
	go build -o ${BINARY} .

run:	
	go env -w GOOS=windows
	go build -o ${BINARY} .
	./${BINARY}

clean:
	go clean 