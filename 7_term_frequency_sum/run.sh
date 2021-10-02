#echo "      5 DOCUMENT_FREQUENCY_TOTAL: component" | go run main.go

# TODO: this isn't quite correct - they should be percentages, not frequencies

#cat /Volumes/git/github/mahout/shell_simplification/cat_with_filename.50.out | go run main.go 2> /dev/null
cat cat_with_filename.full.out | go run main.go  2> /dev/null | sort -n | grep -v '\t0.0'  | uniq | tee keywords.out | edit -
cat keywords.out | grep -v http | perl -pe 's{phase.3..([^\s]+)\s+(.+)\s+(.+)$}{$3\t| $2 |\t$1\n}g'  | sort -n | tee word_to_file_mappings.txt
