#!/bin/bash
pwd=`pwd`

export GOPATH=$pwd

go get github.com/russross/blackfriday
go fmt md2html.go
go build md2html.go

echo "finished"
