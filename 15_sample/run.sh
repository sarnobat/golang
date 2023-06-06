# Native binary
env GOOS=linux GOARCH=amd64 go build sample.go
mv sample sample.linux

env GOOS=darwin GOARCH=amd64 go build sample.go
cp sample sample.mac

