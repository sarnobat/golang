#!/bin/sh

#----------------------------------------------------------------------------
# DESCRIPTION		
# DATE				[:VIM_EVAL:]strftime('%Y-%m-%d')[:END_EVAL:]
# AUTHOR			ss401533@gmail.com                                           
#----------------------------------------------------------------------------


# seq "$1" "$2"
find ~/mwk/snippets/ -type f | xargs grep --no-filename '^#' | sort | uniq | perl -pe 's{^#}{}g'


