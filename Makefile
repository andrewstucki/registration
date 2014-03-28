export GOPATH := $(shell pwd)
all: client embed build

deps:
	go get github.com/mattn/go-sqlite3
	go get github.com/GeertJohan/go.rice
	go get github.com/GeertJohan/go.rice/rice
	go get github.com/emicklei/go-restful
	cd client && npm install && bower install

client:
	cd client && grunt

embed:
	cd server/commands/run && rice embed

build:
	cd server && go build -o ../bin/register

clean:
	cd client && grunt clean
	cd server && rice clean
	rm -rf bin/register

.PHONY: clean client