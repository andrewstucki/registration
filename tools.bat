@echo off
echo ------------------------------
echo Installing client dependencies
echo ------------------------------
echo Installing bundler
CALL gem install bundler > NUL 2>&1
echo Installing grunt-cli
CALL npm install -g grunt-cli > NUL 2>&1
echo Installing bower
CALL npm install -g bower> NUL 2>&1
echo.
echo -------------------------
echo Finished installing tools
echo -------------------------