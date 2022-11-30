#!/bin/sh

#----------------------------------------------------------------------------
# DESCRIPTION		
# DATE				[:VIM_EVAL:]strftime('%Y-%m-%d')[:END_EVAL:]
# AUTHOR			ss401533@gmail.com                                           
#----------------------------------------------------------------------------


cat <<EOF
mkdir /tmp/gomount
go run main.go
rmdir /tmp/gomount
EOF

ARGS="$@"

