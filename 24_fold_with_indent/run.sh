#I don't know why the builtin unix fold doesn't have this. It would be good for printing code that has long lines
set -e

NAME=foldindent

env GOOS=windows GOARCH=amd64 go build main.go
mv main.exe $NAME.exe

env GOOS=linux GOARCH=amd64 go build main.go
mv main $NAME.linux

env GOOS=darwin GOARCH=amd64 go build main.go
mv main $NAME.mac.intel

env GOOS=darwin GOARCH=arm64 go build main.go
mv main $NAME.mac.m1

cp -a -v ${NAME}.linux 		~/github/binaries/linux/bin/${NAME}
cp -a -v ${NAME}.mac.intel 	~/github/binaries/mac.intel/bin/${NAME}
cp -a -v ${NAME}.mac.m1		~/github/binaries/mac.m1/bin/${NAME}
cp -a -v ${NAME}.exe			~/github/binaries/windows/bin/${NAME}.exe

