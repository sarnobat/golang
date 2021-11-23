# Native binary
env GOOS=linux GOARCH=amd64 go build playlists.go
mv playlists playlists.linux

env GOOS=darwin GOARCH=amd64 go build playlists.go
cp playlists playlists.osx

# Note: watch later playlist is not supported anymore by the API. Back to greasemonkey
go run playlists.go --clientSecret /Users/sarnobat/trash/client_secret_803470544206-ni4anh43sh69athsll8s32qkb5dn0des.apps.googleusercontent.com.json
playlists --clientSecret /Users/sarnobat/trash/client_secret_803470544206-lvipsulouuemapfptubh5uq2qop22fj5.apps.googleusercontent.com.json | tee playlists.out.txt