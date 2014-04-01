STARTTIME = @date '+%s' > $@_time
ENDTIME = @read st < $@_time ; st=$$((`date '+%s'`-$$st-68400)) ; echo Total time to run task: `date -r $$st '+%M'` minutes `date -r $$st '+%S'` seconds

export GOPATH := $(CURDIR)
all: client build
	@echo ---------------
	@echo Build Complete!
	@echo ---------------
	@echo ""

deps:
	@echo ---------------------------
	@echo Getting server dependencies
	@echo ---------------------------
	$(STARTTIME)
	@echo Getting go-sqlite3
	@go get github.com/mattn/go-sqlite3
	@echo Getting go.rice
	@go get github.com/andrewstucki/go.rice
	@echo Getting rice binary
	@go get github.com/andrewstucki/go.rice/rice
	@echo Getting go-restful
	@go get github.com/emicklei/go-restful
	$(ENDTIME)
	@echo ""
	@echo ---------------------------
	@echo Getting client dependencies
	@echo ---------------------------
	$(STARTTIME)
	@echo Running "bundle install"
	@cd client && bundle install > /dev/null 2>&1
	@echo Running "npm install"
	@cd client && npm install > /dev/null 2>&1
	@echo Running "bower install"
	@cd client && bower install > /dev/null 2>&1
	$(ENDTIME)
	@echo ""
	@echo -----------------------------
	@echo Finished getting dependencies
	@echo -----------------------------

tools:
	@echo -------------------------
	@echo Installing required tools
	@echo -------------------------
	@echo Installing bundler
	@gem install bundler > /dev/null 2>&1
	@echo Installing grunt-cli
	@npm install -g grunt-cli > /dev/null 2>&1
	@echo Installing bower
	@npm install -g bower > /dev/null 2>&1

client:
	@echo ---------------
	@echo Building client
	@echo ---------------
	$(STARTTIME)
	@cd client && grunt > /dev/null 2>&1
	$(ENDTIME)

build:
	@echo ---------------
	@echo Building server
	@echo ---------------
	$(STARTTIME)
	@cd server && go build -o ../bin/register > /dev/null 2>&1
	@cd server/commands/run && ../../../bin/rice append --exec ../../../bin/register
	@zip -A bin/register > /dev/null
	$(ENDTIME)

clean:
	@echo --------------
	@echo Cleaning files
	@echo --------------
	@cd client
	@grunt clean
	@cd server
	@../bin/rice clean
	@rm -rf bin/register
	@echo Done.

.PHONY: clean client build append
