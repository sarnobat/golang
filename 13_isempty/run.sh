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
GOARCH=amd64 go build

EOF



