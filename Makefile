# Makefile at the root of your project

APP_NAME=gbb
MAIN_PATH=./cmd/gbb

.PHONY: build run clean

build:
	go build -o $(APP_NAME) $(MAIN_PATH)

run:
	go run $(MAIN_PATH)/main.go

clean:
	rm -f $(APP_NAME)
