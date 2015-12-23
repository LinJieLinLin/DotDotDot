@echo off
cd ../../
cd %cd%

set GOPATH=%GOPATH%;%cd%
cd src/dot
bee run

