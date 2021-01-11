#!/bin/sh
protoc -I $GOPATH/game_framework -I . --go_out=plugins=grpc:. --validator_out=. --service_out=plugins=http+clientproxy:. mail.proto
#protoc-go-inject-tag -input=./mail.pb.go -XXX_skip=bson


