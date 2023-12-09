
build:
	go build -o main
	./main

pull-build:
	git pull
	go build -o main
	./main