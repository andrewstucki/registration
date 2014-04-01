@echo off
SET GOPATH=%~dp0
echo ---------------------------
echo Getting server dependencies
echo ---------------------------
echo Getting go-sqlite3
go get github.com/mattn/go-sqlite3
echo Getting go.rice
go get github.com/andrewstucki/go.rice
echo Getting rice binary
go get github.com/andrewstucki/go.rice/rice
echo Getting go-restful
go get github.com/emicklei/go-restful
pushd %GOPATH%\client
echo.
echo ------------------------------
echo Installing client dependencies
echo ------------------------------
echo Running "bundle install"
CALL bundle install > NUL 2>&1
echo Running "npm install"
CALL npm install > NUL 2>&1
echo Running "bower install"
CALL bower install > NUL 2>&1
echo.
echo -----------------------------
echo Finished getting dependencies
echo -----------------------------
popd