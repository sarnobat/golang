#I don't know why the builtin unix fold doesn't have this. It would be good for printing code that has long lines
set -e

NAME=myprog

env GOOS=windows GOARCH=amd64 go build main.go
mv main.exe $NAME.exe

env GOOS=linux GOARCH=amd64 go build main.go
mv main $NAME.linux

env GOOS=darwin GOARCH=amd64 go build main.go
cp main $NAME.mac.intel

env GOOS=darwin GOARCH=arm64 go build main.go
cp main $NAME.mac.m1

