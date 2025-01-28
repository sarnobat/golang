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
echo "revitalizing" | go run lemmatize.go
cat /Volumes/git/mwk.git/snippets/*mwk | perl -pe 's{\b}{\n}g' | go run /Volumes/git/github/2024/golang/26_stemmer_lemmatizer/lemmatize.go 2> /dev/null | sort | uniq -c | sort -n
cat /Volumes/git/mwk.git/snippets/*mwk | perl -pe 's{\b}{\n}g' | /Volumes/git/github/2024/golang/26_stopword_remover/stopwordrm  2> /dev/null | cat | /Volumes/git/github/2024/golang/27_stemmer_lemmatizer/lemmatize 2> /dev/null
EOF
