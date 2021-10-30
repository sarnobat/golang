# Native binary
env GOOS=linux GOARCH=amd64 go build playlists.go
mv playlists playlists.linux

env GOOS=darwin GOARCH=amd64 go build playlists.go
cp playlists playlists.osx

go run playlists.go --clientSecret /Users/sarnobat/trash/client_secret_803470544206-lvipsulouuemapfptubh5uq2qop22fj5.apps.googleusercontent.com.json
