protoc -Iproto --go_out=plugins=grpc:./goproto ./proto/tcpgate.proto
protoc -Iproto --grpc-tcpgw_out=logtostderr=true:./goproto ./proto/tcpgate.proto