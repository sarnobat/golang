#!/bin/sh

#----------------------------------------------------------------------------
# DESCRIPTION		
# DATE				[:VIM_EVAL:]strftime('%Y-%m-%d')[:END_EVAL:]
# AUTHOR			ss401533@gmail.com                                           
#----------------------------------------------------------------------------

set -o errexit
set -o nounset

cat <<EOF | tee /tmp/run.sh | \batcat --plain --paging=never --language sh --theme TwoDark
GOOS=linux GOARCH=amd64 go build isempty.go
mv isempty isempty.linux
GOOS=darwin GOARCH=arm64 go build isempty.go
mv isempty isempty.mac.m1
GOOS=darwin GOARCH=amd64 go build isempty.go
mv isempty isempty.mac.intel

GOOS=linux GOARCH=amd64 go build isdir.go
mv isdir isdir.linux
GOOS=darwin GOARCH=arm64 go build isdir.go
mv isdir isdir.mac.m1
GOOS=darwin GOARCH=amd64  go build isdir.go
mv isdir isdir.mac.intel

GOOS=linux GOARCH=amd64 go build isfile.go
mv isfile isfile.linux
GOOS=darwin GOARCH=amd64 go build isfile.go
mv isfile isfile.mac.intel
GOOS=darwin GOARCH=arm64 go build isfile.go
mv isfile isfile.mac.m1

EOF
cat <<EOF
To build all:
sh /tmp/run.sh
EOF