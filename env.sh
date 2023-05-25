#! /bin/sh

export GOPATH=$(go env | grep GOPATH | cut -d \" -f2)
export PATH=$PATH:$GOPATH"/bin"

