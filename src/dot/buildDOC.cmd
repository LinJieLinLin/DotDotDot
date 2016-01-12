@echo off
cd ../../
cd %cd%

set GOPATH=%GOPATH%;%cd%
echo http://localhost:1231/pkg/dot/
explorer http://localhost:1231/pkg/dot/
godoc -http :1231
