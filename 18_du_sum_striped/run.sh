#!/bin/sh

#----------------------------------------------------------------------------
# DESCRIPTION		
# DATE				[:VIM_EVAL:]strftime('%Y-%m-%d')[:END_EVAL:]
# AUTHOR			ss401533@gmail.com                                           
#----------------------------------------------------------------------------

#set -e
set -o errexit

test $# -gt 0 && echo "args given" || echo "no args"
# TODO: string comparison check (both ways)

cat <<EOF | batcat --plain --paging=never --language sh --theme TwoDark
TODO: port to golang for easier cold running:
/Volumes/git/repos_personal.git/src.git/2021/python/du_sum_striped/du_sum_striped.py
EOF



