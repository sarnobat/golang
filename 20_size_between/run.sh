#!/bin/sh

#----------------------------------------------------------------------------
# DESCRIPTION		
# DATE				[:VIM_EVAL:]strftime('%Y-%m-%d')[:END_EVAL:]
# AUTHOR			ss401533@gmail.com                                           
#----------------------------------------------------------------------------

set -o errexit
set -o nounset

FILE="sizeis"
cat <<EOF | tee /tmp/run.sh | \batcat --plain --paging=never --language sh --theme TwoDark
set -e
GOOS=windows GOARCH=amd64 go build ${FILE}.go
mv ${FILE}.exe ${FILE}.exe
GOOS=linux GOARCH=amd64 go build ${FILE}.go
mv ${FILE} ${FILE}.linux
GOOS=darwin GOARCH=arm64 go build ${FILE}.go
mv ${FILE} ${FILE}.mac.m1
GOOS=darwin GOARCH=amd64 go build ${FILE}.go
mv ${FILE} ${FILE}.mac.intel


rsync -a -v *m1		/Volumes/git/github/binaries/mac.m1/bin
rsync -a -v *linux 	/Volumes/git/github/binaries/linux/bin
rsync -a -v *intel 	/Volumes/git/github/binaries/mac.intel/bin
rsync -a -v *exe 	/Volumes/git/github/binaries/windows/bin
EOF

cat <<EOF | \batcat --plain --paging=never --language sh --theme TwoDark

cd /Volumes/git/github/binaries/

To build all:
sh -x /tmp/run.sh
EOF
