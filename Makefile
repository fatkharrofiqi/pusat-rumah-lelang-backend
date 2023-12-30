
build:
	go build -o cmd/web/main
	./main

pull-build:
	git pull
	go build -o cmd/web/main
	./main