// Code generated by protoc-gen-grpc-tcpway. DO NOT EDIT.
// source: tcpgate.proto

/*
Package gwProto is a tcp/ws proxy.

It translates protobuf/Json packet into gRPC APIs.
*/
package gwProto

import (
	"context"
	"encoding/json"
	"errors"
	"strings"
	"time"

	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// define func
type registerHandler func(args *TransmitArgs) (err error)

type transmit_Backendsvr2_Handler func(*TransmitArgs, Backendsvr2Client) (proto.Message, error)

type transmit_Backendsvr1_Handler func(*TransmitArgs, Backendsvr1Client) (proto.Message, error)

type TransmitArgs struct {
	Method       string
	Endpoint     string
	Conn         *grpc.ClientConn
	MD           metadata.MD
	Data         []byte
	Codec        uint16
	Opts         []grpc.DialOption
	DoneCallback func(proto.Message)
	ctx          context.Context
}

var (
	// tag @id to package.TargetService/Method map
	id2meth = map[uint16]string{

		4: "gwProto.Backendsvr2/Method2",
		2: "gwProto.Backendsvr1/Method1",
	}

	meth2id = map[string]uint16{}

	id2struct = map[uint16]proto.Message{

		4: &Method2Request{},
		5: &Method2Reply{},
		2: &Method1Request{},
		3: &Method1Reply{},
	}

	transmit_Backendsvr2_Map = map[string]transmit_Backendsvr2_Handler{}

	transmit_Backendsvr1_Map = map[string]transmit_Backendsvr1_Handler{}

	serviceMap = map[string]registerHandler{}
)

func init() {
	for k, v := range id2meth {
		meth2id[v] = k
	}

	// todo something handler

	transmit_Backendsvr2_Map["gwProto.Backendsvr2/Method2"] = request_Backendsvr2_Method2

	transmit_Backendsvr1_Map["gwProto.Backendsvr1/Method1"] = request_Backendsvr1_Method1

	serviceMap["gwProto.Backendsvr2"] = register_Backendsvr2_Transmitor

	serviceMap["gwProto.Backendsvr1"] = register_Backendsvr1_Transmitor

}

func decodeBytes(data []byte, codec uint16, inst proto.Message) error {
	if codec == 0 {
		if err := proto.Unmarshal(data, inst); err != nil {
			return err
		}
	} else if codec == 1 {
		if err := json.Unmarshal(data, inst); err != nil {
			return err
		}
	}
	return nil
}

// get meth(package.TargetService/Method) by id(cmdid)
func GetMethById(id uint16) string {
	return id2meth[id]
}

func GetIdByMeth(meth string) uint16 {
	return meth2id[meth]
}

// 根据@id/@upid/@downid标签获取对应方法的请求参数对象
func GetMsgObjById(id uint16) proto.Message {
	return id2struct[id]
}

func ParseMethod(method string) (string, string, string, error) {
	method = strings.Trim(method, "/")
	dotIdx := strings.Index(method, ".")
	slashIdx := strings.Index(method, "/")
	if dotIdx < 1 || slashIdx < 1 || dotIdx > slashIdx {
		return "", "", "", errors.New("method must be type of 'package.ServiceName/Method'")
	}
	packageName := method[:dotIdx]
	serviceName := strings.Trim(method[dotIdx:slashIdx], ".")
	methodName := strings.Trim(method[slashIdx:], "/")
	return packageName, serviceName, methodName, nil
}

// define call enter point
func RegisterTransmitor(args *TransmitArgs) error {
	if len(args.Method) < 1 || len(args.Endpoint) < 1 || len(args.MD) < 1 || args.DoneCallback == nil {
		return errors.New("transmit args empty")
	}

	packageName, serviceName, _, err := ParseMethod(args.Method)
	if err != nil {
		return err
	}
	packageService := packageName + "." + serviceName
	if handler, ok := serviceMap[packageService]; ok {
		err := handler(args)
		return err
	}
	return errors.New("method not register yet")
}

// registor single service enter point

// *********************************************************************************
// 注册TcpGateway传输转换入口

func register_Backendsvr2_Transmitor(args *TransmitArgs) (err error) {
	if args.Conn == nil {
		conn, err := grpc.Dial(args.Endpoint, args.Opts...)
		if err != nil {
			return err
		}
		defer conn.Close()
		args.Conn = conn
	}

	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, 5*time.Second)
	ctx = metadata.NewOutgoingContext(ctx, args.MD)
	args.ctx = ctx
	//
	client := NewBackendsvr2Client(args.Conn)
	handler, ok := transmit_Backendsvr2_Map[args.Method]
	if !ok {
		return errors.New("method error")
	}
	res, err := handler(args, client)
	if err != nil {
		return err
	}
	args.DoneCallback(res)
	return nil
}

// 注册Backendsvr2/Method2 传输方法入口
// 后端服务2
// 注释
func request_Backendsvr2_Method2(args *TransmitArgs, client Backendsvr2Client) (proto.Message, error) {
	protoReq := &Method2Request{}
	if err := decodeBytes(args.Data, args.Codec, protoReq); err != nil {
		return nil, errors.New("codec err[" + err.Error() + "]")
	}
	reply, err := client.Method2(args.ctx, protoReq)
	if err != nil {
		return nil, errors.New("call err[" + err.Error() + "]")
	}
	return reply, nil
}

// *********************************************************************************
// 注册TcpGateway传输转换入口

func register_Backendsvr1_Transmitor(args *TransmitArgs) (err error) {
	if args.Conn == nil {
		conn, err := grpc.Dial(args.Endpoint, args.Opts...)
		if err != nil {
			return err
		}
		defer conn.Close()
		args.Conn = conn
	}

	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, 5*time.Second)
	ctx = metadata.NewOutgoingContext(ctx, args.MD)
	args.ctx = ctx
	//
	client := NewBackendsvr1Client(args.Conn)
	handler, ok := transmit_Backendsvr1_Map[args.Method]
	if !ok {
		return errors.New("method error")
	}
	res, err := handler(args, client)
	if err != nil {
		return err
	}
	args.DoneCallback(res)
	return nil
}

// 注册Backendsvr1/Method1 传输方法入口
// 后端服务1
// 注释
func request_Backendsvr1_Method1(args *TransmitArgs, client Backendsvr1Client) (proto.Message, error) {
	protoReq := &Method1Request{}
	if err := decodeBytes(args.Data, args.Codec, protoReq); err != nil {
		return nil, errors.New("codec err[" + err.Error() + "]")
	}
	reply, err := client.Method1(args.ctx, protoReq)
	if err != nil {
		return nil, errors.New("call err[" + err.Error() + "]")
	}
	return reply, nil
}
