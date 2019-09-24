#!/usr/bin/env bash
protoc -Iproto --go_out=plugins=grpc:./goproto ./proto/msg.proto
protoc -Iproto --go_out=plugins=grpc:./goproto ./proto/backendsvr1.proto
protoc -Iproto --go_out=plugins=grpc:./goproto ./proto/backendsvr2.proto
protoc -Iproto --grpc-tcpgw_out=logtostderr=true:./goproto ./proto/tcpgate.proto