cat tree_rohidekar.mwk | go run main.go

# Native binary
env GOOS=linux GOARCH=amd64 go build main.go
mv main tree2fs.linux

env GOOS=darwin GOARCH=amd64 go build main.go
cp main tree2fs.osx
mv main tree2fs

#sort du_jpg_only_reduced_md5sum.txt  | main | sort -n | tee /tmp/out.txt
#echo "cat /tmp/out.txt | tail -500 "

#cat ~/trash/out.txt | go run main.go
