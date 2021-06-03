go build
mv 3_gedcom gedcom

# no args - use default values
gedcom

gedcom --name Sridhar --file rohidekar.ged hi

cat rohidekar.ged | gedcom_indent

env GOOS=linux GOARCH=amd64 go build gedcom2mwk.go
mv gedcom2mwk gedcom2mwk.linux

env GOOS=darwin GOARCH=amd64 go build gedcom2mwk.go
cp gedcom2mwk gedcom2mwk.osx

cat rohidekar.ged | gedcom2mwk