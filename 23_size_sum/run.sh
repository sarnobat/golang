cat <<EOF
TODO: add the size of every file supplied on stdin (outputted from find, NOT du)

This is useful if,  .e.g you want to find the sum of all files of type mts

Limitation: doesn't work for small sizes (e.g. total is reported in 0G, 1G or more)
EOF

env GOOS=linux GOARCH=amd64 go build main.go
mv main size_sum.linux

env GOOS=darwin GOARCH=amd64 go build main.go
mv main size_sum.mac.intel

