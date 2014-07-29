@echo off

setlocal

if exist make.bat goto run
echo make.bat must be run from its folder
goto END

: run

chcp 437

SET GOROOT=D:\server\go
SET GOPATH=%GOPATH%;d:\git\tools\md2html

go get github.com/russross/blackfriday
go fmt md2html.go
go build md2html.go


:end
echo finished
pause
