set -e

NAME=groupby

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

# ------------------------------

exit
echo "hello" | go run main.go

# Native binary
env GOOS=linux GOARCH=amd64 go build main.go
mv main main.linux

env GOOS=darwin GOARCH=amd64 go build main.go
cp main main.osx

sort du_jpg_only_reduced_md5sum.txt  | head -100 | main | sort -n | tee /tmp/out.txt
echo "cat /tmp/out.txt | tail -500 "

#cat ~/trash/out.txt | go run main.go
