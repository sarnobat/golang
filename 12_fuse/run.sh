#!/bin/sh

#----------------------------------------------------------------------------
# DESCRIPTION		
# DATE				[:VIM_EVAL:]strftime('%Y-%m-%d')[:END_EVAL:]
# AUTHOR			ss401533@gmail.com                                           
#----------------------------------------------------------------------------


cat <<EOF | tee /tmp/run.sh
mkdir /tmp/gomount
cd ~/github/go-fuse/example/hello/
go run main.go /tmp/gomount
umount /tmp/gomount
rmdir /tmp/gomount
EOF

ARGS="$@"

