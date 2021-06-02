go build
mv 3_gedcom gedcom

# no args - use default values
gedcom

gedcom --name Sridhar --file rohidekar.ged hi

cat rohidekar.ged | gedcom_indent

env GOOS=linux GOARCH=amd64 go build list_families.go
mv list_families list_families.linux

env GOOS=darwin GOARCH=amd64 go build list_families.go
cp list_families list_families.osx

cat rohidekar.ged | list_families