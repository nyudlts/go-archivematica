tidy:
	go mod tidy

build:
	go mod tidy
	go build -o am-services main/main.go

install:
	cp am-services /usr/local/bin