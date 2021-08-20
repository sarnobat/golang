#echo "      5 DOCUMENT_FREQUENCY_TOTAL: component" | go run main.go

#cat /Volumes/git/github/mahout/shell_simplification/cat_with_filename.50.out | go run main.go 2> /dev/null
cat cat_with_filename.full.out | go run main.go  2> /dev/null | sort -n | grep -v '\t0' | grep -v '\t1' | grep -v '2.0' | uniq | edit -
