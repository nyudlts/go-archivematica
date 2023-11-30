tidy:
	go mod tidy

build:
	go mod tidy
	go build -o am_services main/main.go