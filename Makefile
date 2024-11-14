APP_NAME=vara
APP_PATH=.
APP_BIN_PATH=./bin

.PHONY: all
all: build run

.PHONY: run
run:
	go run ${APP_PATH}

.PHONY: build
build:
	GOOS=linux GOARCH=amd64 go build -o ${APP_BIN_PATH}/linux/${APP_NAME} ${APP_PATH}
	GOOS=darwin GOARCH=amd64 go build -o ${APP_BIN_PATH}/darwin/${APP_NAME} ${APP_PATH}
	GOOS=windows GOARCH=amd64 go build -o ${APP_BIN_PATH}/windows/${APP_NAME}.exe ${APP_PATH}