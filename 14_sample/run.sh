#!/bin/sh

#----------------------------------------------------------------------------
# DESCRIPTION		
# DATE				[:VIM_EVAL:]strftime('%Y-%m-%d')[:END_EVAL:]
# AUTHOR			ss401533@gmail.com                                           
#----------------------------------------------------------------------------

set -o errexit
set -o nounset

cat <<EOF | tee /tmp/run.sh | \batcat --plain --paging=never --language sh --theme TwoDark
GOOS=linux GOARCH=amd64 go build sample.go
mv sample sample.linux
GOOS=darwin GOARCH=arm64 go build sample.go
mv sample sample.mac.m1
GOOS=darwin GOARCH=amd64 go build sample.go
mv sample sample.mac.intel


rsync -a -v *m1		/Volumes/git/github/binaries/mac.m1/bin
rsync -a -v *linux 	/Volumes/git/github/binaries/ubuntu/bin
rsync -a -v *intel 	/Volumes/git/github/binaries/mac.intel/bin
EOF

cat <<EOF | \batcat --plain --paging=never --language sh --theme TwoDark

cd /Volumes/git/github/binaries/

To build all:
sh /tmp/run.sh
EOF