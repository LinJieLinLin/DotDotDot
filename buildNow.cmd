@echo off

cd %cd%

set GOPATH=%GOPATH%;%cd%

go build  main

echo build finished

pause /s
main.exe
