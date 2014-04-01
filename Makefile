STARTTIME = @date '+%s' > $@_time
ENDTIME = @read st < $@_time ; st=$$((`date '+%s'`-$$st-68400)) ; echo Total time: `date -r $$st '+%H:%M:%S'`

export GOPATH := $(CURDIR)
all: client build

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
	@cd client
	@echo ""
	@echo ---------------------------
	@echo Getting client dependencies
	@echo ---------------------------
	$(STARTTIME)
	@bundle install > /dev/null 2>&1
	@npm install > /dev/null 2>&1
	@bower install > /dev/null 2>&1
	$(ENDTIME)
	@echo ""
	@echo -----------------------------
	@echo Finished getting dependencies
	@echo -----------------------------

tools:
	@echo -------------------------
	@echo Installing required tools
	@echo -------------------------
	@gem install bundler > /dev/null 2>&1
	@npm install -g grunt-cli > /dev/null 2>&1
	@npm install -g bower > /dev/null 2>&1

client:
	@echo ---------------
	@echo Building client
	@echo ---------------
	@cd client
	$(STARTTIME)
	@grunt > /dev/null 2>&1
	$(ENDTIME)

build:
	@echo ---------------
	@echo Building server
	@echo ---------------
	@cd server
	$(STARTTIME)
	@go build -o ../bin/register > /dev/null 2>&1
	@cd commands/run
	@../../../bin/rice append --exec ../../../bin/register
	@zip -A ../../../bin/register
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