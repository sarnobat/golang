cat <<EOF
# Native binary
env GOOS=linux GOARCH=amd64 go build playlists.go
mv playlists playlists.linux

GOPATH=/Users/sarnobat/2021/gopath/ env GOOS=darwin GOARCH=amd64 go build playlists.go
cp playlists playlists.osx

GOPATH=/Users/sarnobat/2021/gopath/ env GOOS=darwin GOARCH=amd64 go build playlist_items.go
cp playlist_items playlist_items.osx

# Note: watch later playlist is not supported anymore by the API. Back to greasemonkey :(

# TODO - try to do this with a Dockerfile for experience, and to see what is still not automatic
# minimal example
playlists --clientSecret client_secret_803470544206-ni4anh43sh69athsll8s32qkb5dn0des.apps.googleusercontent.com.json
GOPATH=/Users/sarnobat/2021/gopath/  go run playlists.go --clientSecret client_secret_803470544206-ni4anh43sh69athsll8s32qkb5dn0des.apps.googleusercontent.com.json | tee playlists.out.txt

./playlist_items --clientSecret client_secret_803470544206-ni4anh43sh69athsll8s32qkb5dn0des.apps.googleusercontent.com.json | tee playlist_items.out.txt
GOPATH=/Users/sarnobat/2021/gopath/  go run playlist_items.go --clientSecret client_secret_803470544206-ni4anh43sh69athsll8s32qkb5dn0des.apps.googleusercontent.com.json | tee playlist_items.out.txt

## Setting up the credentials is a pain (but worth it? I haven't decided)
EOF

cat <<EOF
Easiest (for public urls)
-------
(I'm also storing the output here: https://docs.google.com/spreadsheets/d/1n05xdlaG4oPppcpkJUmvrVjVz2TK43ZHiMuKiMk0F6c/edit?gid=435927657#gid=435927657)
sh /Volumes/git/computers.git/2022/mac/bin/2022/2021/youtube_playlists.sh | tee ~/db.git/youtube_playlists.public.`date -I`.txt

Easiest (for private urls)
-------
/Volumes/git/src.git/python/youtube_playlists_python/yt.py
EOF