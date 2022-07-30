#!/usr/bin/env just --justfile

# run application
run action="":
    go run . {{action}}

update:
  go get -u
  go mod tidy -v
 
 
init:
     go install github.com/spf13/cobra-cli@latest
     $GOPATH/bin/cobra-cli --config .cobra.yaml init
# add command
add command:
    $GOPATH/bin/cobra-cli --config .cobra.yaml add {{command}}