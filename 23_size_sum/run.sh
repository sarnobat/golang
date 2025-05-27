cat <<EOF
TODO 2024-10: add the size of every file supplied on stdin (outputted from find, NOT du)
TODO 2024-10: show running total because it may take time to find the full list of input files.

This is useful if,  .e.g you want to find the sum of all files of type mts

Limitation: doesn't work for small sizes (e.g. total is reported in 0G, 1G or more)
EOF

env GOOS=linux GOARCH=amd64 go build main.go
mv main size_sum.linux

env GOOS=darwin GOARCH=amd64 go build main.go
mv main size_sum.mac.intel

