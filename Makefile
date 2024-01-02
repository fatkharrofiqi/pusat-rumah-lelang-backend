
build:
	go build -o main cmd/web/main
	./main

pull-build:
	git pull
	go build -o main cmd/web/main.go
	./main