export GOPATH := $(shell pwd)
all: client embed build

deps:
	go get github.com/mattn/go-sqlite3
	go get github.com/GeertJohan/go.rice
	go get github.com/GeertJohan/go.rice/rice
	go get github.com/emicklei/go-restful
	gem install bundler && npm install -g grunt-cli && npm install -g bower
	cd client && bundle install && npm install && bower install

client:
	cd client && grunt

embed:
	cd server/commands/run && ../../../bin/rice embed

build:
	cd server && go build -o ../bin/register

clean:
	cd client && grunt clean
	cd server && ../bin/rice clean
	rm -rf bin/register

.PHONY: clean client
