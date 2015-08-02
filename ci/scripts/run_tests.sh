#!/bin/bash

# change to root of release
DIR=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )
cd $DIR/../..

cat > ~/.lockpick << EOF
app_id: foobar
user_id: myuserid
key: mykeyid
output_file: ~/mygitcrypt.key
vault_address: "https://127.0.0.1:8200"
EOF

godep restore

go vet -x ./...
golint ./...
go test -v ./...