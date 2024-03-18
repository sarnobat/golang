go run helloworld.go

# static linking (2021)
go build helloworld.go

# static linking (2020) - not necessary anymore apparently
#go tool compile helloworld.go
#env CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -a -o golang_tutorial
