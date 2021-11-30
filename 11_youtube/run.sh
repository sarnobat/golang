# Native binary
env GOOS=linux GOARCH=amd64 go build playlists.go
mv playlists playlists.linux

env GOOS=darwin GOARCH=amd64 go build playlists.go
cp playlists playlists.osx

# Note: watch later playlist is not supported anymore by the API. Back to greasemonkey :(

# TODO - try to do this with a Dockerfile for experience, and to see what is still not automatic
# minimal example
go run playlists.go --clientSecret client_secret_803470544206-ni4anh43sh69athsll8s32qkb5dn0des.apps.googleusercontent.com.json
playlists --clientSecret go run playlists.go --clientSecret client_secret_803470544206-ni4anh43sh69athsll8s32qkb5dn0des.apps.googleusercontent.com.json
