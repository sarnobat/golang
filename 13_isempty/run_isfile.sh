#!/bin/sh

#----------------------------------------------------------------------------
# DESCRIPTION		
# DATE				[:VIM_EVAL:]strftime('%Y-%m-%d')[:END_EVAL:]
# AUTHOR			ss401533@gmail.com                                           
#----------------------------------------------------------------------------

set -o errexit
set -o nounset

PROGRAM_NAME="isfile"

cat <<EOF | tee /tmp/run.sh | \batcat --style=plain --paging=never --language sh --theme TwoDark
set -e
GOOS=linux GOARCH=amd64 go build $PROGRAM_NAME.go
mv $PROGRAM_NAME $PROGRAM_NAME.linux
GOOS=darwin GOARCH=arm64 go build $PROGRAM_NAME.go
mv $PROGRAM_NAME $PROGRAM_NAME.mac.m1
GOOS=darwin GOARCH=amd64 go build $PROGRAM_NAME.go
mv $PROGRAM_NAME $PROGRAM_NAME.mac.intel
GOOS=windows GOARCH=amd64 go build $PROGRAM_NAME.go
mv ${PROGRAM_NAME}.exe ${PROGRAM_NAME}.exe

rsync -a -v *m1					~/github/binaries/mac.m1/bin
rsync -a -v *linux 				~/github/binaries/linux/bin
rsync -a -v *intel 				~/github/binaries/mac.intel/bin
rsync -a -v ${PROGRAM_NAME}.exe	~/github/binaries/windows/bin/${PROGRAM_NAME}.exe

rename -f -v 's{.linux$}{}g' /Volumes/git/github/binaries/linux/bin/*.linux
rename -f -v 's{.mac.intel$}{}g' /Volumes/git/github/binaries/mac.intel/bin/*.intel
rename -f -v 's{.mac.m1$}{}g' /Volumes/git/github/binaries/mac.m1/bin/*.m1
EOF

cat <<EOF | \batcat --plain --paging=never --language sh --theme TwoDark

cd /Volumes/git/github/binaries/

To build all:
sh /tmp/run.sh
EOF
