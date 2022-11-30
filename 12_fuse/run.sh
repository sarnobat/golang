#!/bin/sh

#----------------------------------------------------------------------------
# DESCRIPTION		
# DATE				[:VIM_EVAL:]strftime('%Y-%m-%d')[:END_EVAL:]
# AUTHOR			ss401533@gmail.com                                           
#----------------------------------------------------------------------------


cat <<EOF | tee /tmp/run.sh
# The required code is extensive, you can't use go get apparently
cd ~/github/go-fuse/example/hello/
mkdir /tmp/gomount
go run main.go /tmp/gomount
umount /tmp/gomount
rmdir /tmp/gomount
EOF

ARGS="$@"

