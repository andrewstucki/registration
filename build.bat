@echo off 
SET BUILDHOME=%~dp0
SET GOPATH=%~dp0
IF EXIST %BUILDHOME%\bin\register.exe DEL %BUILDHOME%\bin\register.exe
pushd %BUILDHOME%\client
echo ---------------
echo Building client
echo ---------------
set start=%time%
CALL grunt > NUL 2>&1
set end=%time%
set options="tokens=1-4 delims=:."
for /f %options% %%a in ("%start%") do set start_h=%%a&set /a start_m=100%%b %% 100&set /a start_s=100%%c %% 100&set /a start_ms=100%%d %% 100
for /f %options% %%a in ("%end%") do set end_h=%%a&set /a end_m=100%%b %% 100&set /a end_s=100%%c %% 100&set /a end_ms=100%%d %% 100
set /a hours=%end_h%-%start_h%
set /a mins=%end_m%-%start_m%
set /a secs=%end_s%-%start_s%
set /a ms=%end_ms%-%start_ms%
if %hours% lss 0 set /a hours = 24%hours%
if %mins% lss 0 set /a hours = %hours% - 1 & set /a mins = 60%mins%
if %secs% lss 0 set /a mins = %mins% - 1 & set /a secs = 60%secs%
if %ms% lss 0 set /a secs = %secs% - 1 & set /a ms = 100%ms%
if 1%ms% lss 100 set ms=0%ms%
:: mission accomplished
set /a totalsecs = %hours%*3600 + %mins%*60 + %secs% 
echo Client build took %hours%:%mins%:%secs%.%ms% (%totalsecs%.%ms%s total)
popd
pushd %BUILDHOME%\server
echo.
echo ---------------
echo Building server
echo ---------------
set start=%time%
go build -o %BUILDHOME%\bin\register.exe > NUL 2>&1
popd
pushd %BUILDHOME%\server\commands\run
%BUILDHOME%\bin\rice.exe append --exec %BUILDHOME%\bin\register.exe > NUL 2>&1
popd
zip -A %BUILDHOME%\bin\register.exe > NUL 2>&1
set end=%time%
set options="tokens=1-4 delims=:."
for /f %options% %%a in ("%start%") do set start_h=%%a&set /a start_m=100%%b %% 100&set /a start_s=100%%c %% 100&set /a start_ms=100%%d %% 100
for /f %options% %%a in ("%end%") do set end_h=%%a&set /a end_m=100%%b %% 100&set /a end_s=100%%c %% 100&set /a end_ms=100%%d %% 100
set /a hours=%end_h%-%start_h%
set /a mins=%end_m%-%start_m%
set /a secs=%end_s%-%start_s%
set /a ms=%end_ms%-%start_ms%
if %hours% lss 0 set /a hours = 24%hours%
if %mins% lss 0 set /a hours = %hours% - 1 & set /a mins = 60%mins%
if %secs% lss 0 set /a mins = %mins% - 1 & set /a secs = 60%secs%
if %ms% lss 0 set /a secs = %secs% - 1 & set /a ms = 100%ms%
if 1%ms% lss 100 set ms=0%ms%
:: mission accomplished
set /a totalsecs = %hours%*3600 + %mins%*60 + %secs% 
echo Server build took %hours%:%mins%:%secs%.%ms% (%totalsecs%.%ms%s total)
echo.
echo -----------------------------
echo Finished building application
echo -----------------------------