run:
	go run .

build:
	ifeq($(OS), Windows_NT)
		go build -o library-api.exe
		./library-api.exe
	else
		go build -o library-api
		./library-api
	endif
	