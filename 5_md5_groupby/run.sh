echo "hello" | go run main.go

# Native binary
env GOOS=linux GOARCH=amd64 go build main.go
mv main main.linux

env GOOS=darwin GOARCH=amd64 go build main.go
cp main main.osx

sort du_jpg_only_reduced_md5sum.txt  | main | sort -n

#cat ~/trash/out.txt | go run main.go
