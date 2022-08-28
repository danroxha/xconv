BINARY_NAME=xc

build:
# GOARCH=amd64 GOOS=darwin go build -o ${BINARY_NAME}-darwin
	GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME}
	GOARCH=amd64 GOOS=windows go build -o ${BINARY_NAME}.exe

run:
	./${BINARY_NAME} version

clean:
	go clean
#	rm ${BINARY_NAME}-darwin
	rm -f ${BINARY_NAME}
	rm -f ${BINARY_NAME}.exe