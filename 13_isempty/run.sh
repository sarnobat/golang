#!/bin/sh

#----------------------------------------------------------------------------
# DESCRIPTION		
# DATE				[:VIM_EVAL:]strftime('%Y-%m-%d')[:END_EVAL:]
# AUTHOR			ss401533@gmail.com                                           
#----------------------------------------------------------------------------

set -o errexit
set -o nounset

test $# -gt 0 && echo "args given" || echo "no args"

cat <<EOF | \batcat --plain --paging=never --language sh --theme TwoDark
GOOS=linux GOARCH=amd64 go build isempty.go
mv isempty isempty.linux
go build isempty.go
mv isempty isempty.mac

GOOS=linux GOARCH=amd64 go build isdir.go
mv isdir isdir.linux
go build isdir.go
mv isdir isdir.mac

GOOS=linux GOARCH=amd64 go build isfile.go
mv isfile isfile.linux
go build isdir.go
mv isdir isdir.mac

EOF



