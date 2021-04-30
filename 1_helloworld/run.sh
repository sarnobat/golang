go run helloworld.go

# static linking?
go tool compile helloworld.go

# static linking
env CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -a -o golang_tutorial