export GOPATH := $(CURDIR)
all: build

deps:
	go get github.com/mattn/go-sqlite3
	go get github.com/andrewstucki/go.rice
	go get github.com/andrewstucki/go.rice/rice
	go get github.com/emicklei/go-restful
	cd client; bundle install; npm install; bower install

tools:
	gem install bundler
	npm install -g grunt-cli
	npm install -g bower

client:
	cd client && grunt

build: client
	cd server && go build -o ../bin/register && cd commands/run && ../../../bin/rice append --exec ../../../bin/register

clean:
	cd client; grunt clean
	cd server; ../bin/rice clean
	rm -rf bin/register

.PHONY: clean client build append
