echo "hello" | go run main.go

# Native binary
env GOOS=linux GOARCH=amd64 go build main.go
mv main main.linux

env GOOS=darwin GOARCH=amd64 go build main.go
cp main main.osx

cat keywords.full.out | perl -pe 's{^phase ...(.*)\t(.*)\t(.*)$}{$3\t$2\t$1}g' | sort | tee /tmp/out.txt
echo "cat /tmp/out.txt | tail -500 "

#cat ~/trash/out.txt | go run main.go
