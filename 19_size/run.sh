#!/bin/sh

#----------------------------------------------------------------------------
# DESCRIPTION		
# DATE				::DATE_CREATED::
# AUTHOR			ss401533@gmail.com                                           
#----------------------------------------------------------------------------

set -o errexit
set -o nounset

echo `date -I`"\t$0" >> ~/computers.git/analytics_stats.sh.txt

test $# -gt 0 && echo "args given" || echo "no args"

cat <<EOF | \batcat --plain --paging=never --language sh --theme TwoDark
TODO (useful for locate -b output)

EOF



