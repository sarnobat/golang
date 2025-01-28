#!/bin/sh

#----------------------------------------------------------------------------
# DESCRIPTION		
# DATE				2024
# AUTHOR			ss401533@gmail.com                                           
#----------------------------------------------------------------------------
# template found at ~/.vim/sh_header.temp

set -o errexit
echo "$0" >> ~/db.git/command_history.txt | ts >> ~/db.git/command_history_timestamped.txt

cat <<EOF | batcat --style=plain --paging=never --language sh --theme TwoDark
cat /Volumes/git/mwk.git/snippets/*mwk | perl -pe 's{\b}{\n}g' | /Volumes/git/github/2024/golang/27_stopword_remover/stopwordrm | sort | uniq -c | sort -n
EOF

