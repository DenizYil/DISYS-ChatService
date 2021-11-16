// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PeerClient is the client API for Peer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PeerClient interface {
	Broadcast(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Empty, error)
	Join(ctx context.Context, in *JoinMessage, opts ...grpc.CallOption) (Peer_JoinClient, error)
	Publish(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Empty, error)
	Retrieve(ctx context.Context, in *RetrieveMessage, opts ...grpc.CallOption) (Peer_RetrieveClient, error)
	Release(ctx context.Context, in *ReleaseMessage, opts ...grpc.CallOption) (*Empty, error)
}

type peerClient struct {
	cc grpc.ClientConnInterface
}

func NewPeerClient(cc grpc.ClientConnInterface) PeerClient {
	return &peerClient{cc}
}

func (c *peerClient) Broadcast(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/api.Peer/Broadcast", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *peerClient) Join(ctx context.Context, in *JoinMessage, opts ...grpc.CallOption) (Peer_JoinClient, error) {
	stream, err := c.cc.NewStream(ctx, &Peer_ServiceDesc.Streams[0], "/api.Peer/Join", opts...)
	if err != nil {
		return nil, err
	}
	x := &peerJoinClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Peer_JoinClient interface {
	Recv() (*Message, error)
	grpc.ClientStream
}

type peerJoinClient struct {
	grpc.ClientStream
}

func (x *peerJoinClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *peerClient) Publish(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/api.Peer/Publish", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *peerClient) Retrieve(ctx context.Context, in *RetrieveMessage, opts ...grpc.CallOption) (Peer_RetrieveClient, error) {
	stream, err := c.cc.NewStream(ctx, &Peer_ServiceDesc.Streams[1], "/api.Peer/Retrieve", opts...)
	if err != nil {
		return nil, err
	}
	x := &peerRetrieveClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Peer_RetrieveClient interface {
	Recv() (*RetrieveReply, error)
	grpc.ClientStream
}

type peerRetrieveClient struct {
	grpc.ClientStream
}

func (x *peerRetrieveClient) Recv() (*RetrieveReply, error) {
	m := new(RetrieveReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *peerClient) Release(ctx context.Context, in *ReleaseMessage, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/api.Peer/Release", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PeerServer is the server API for Peer service.
// All implementations must embed UnimplementedPeerServer
// for forward compatibility
type PeerServer interface {
	Broadcast(context.Context, *Message) (*Empty, error)
	Join(*JoinMessage, Peer_JoinServer) error
	Publish(context.Context, *Message) (*Empty, error)
	Retrieve(*RetrieveMessage, Peer_RetrieveServer) error
	Release(context.Context, *ReleaseMessage) (*Empty, error)
	mustEmbedUnimplementedPeerServer()
}

// UnimplementedPeerServer must be embedded to have forward compatible implementations.
type UnimplementedPeerServer struct {
}

func (UnimplementedPeerServer) Broadcast(context.Context, *Message) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Broadcast not implemented")
}
func (UnimplementedPeerServer) Join(*JoinMessage, Peer_JoinServer) error {
	return status.Errorf(codes.Unimplemented, "method Join not implemented")
}
func (UnimplementedPeerServer) Publish(context.Context, *Message) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Publish not implemented")
}
func (UnimplementedPeerServer) Retrieve(*RetrieveMessage, Peer_RetrieveServer) error {
	return status.Errorf(codes.Unimplemented, "method Retrieve not implemented")
}
func (UnimplementedPeerServer) Release(context.Context, *ReleaseMessage) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Release not implemented")
}
func (UnimplementedPeerServer) mustEmbedUnimplementedPeerServer() {}

// UnsafePeerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PeerServer will
// result in compilation errors.
type UnsafePeerServer interface {
	mustEmbedUnimplementedPeerServer()
}

func RegisterPeerServer(s grpc.ServiceRegistrar, srv PeerServer) {
	s.RegisterService(&Peer_ServiceDesc, srv)
}

func _Peer_Broadcast_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeerServer).Broadcast(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Peer/Broadcast",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeerServer).Broadcast(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _Peer_Join_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(JoinMessage)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PeerServer).Join(m, &peerJoinServer{stream})
}

type Peer_JoinServer interface {
	Send(*Message) error
	grpc.ServerStream
}

type peerJoinServer struct {
	grpc.ServerStream
}

func (x *peerJoinServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func _Peer_Publish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeerServer).Publish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Peer/Publish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeerServer).Publish(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _Peer_Retrieve_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RetrieveMessage)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PeerServer).Retrieve(m, &peerRetrieveServer{stream})
}

type Peer_RetrieveServer interface {
	Send(*RetrieveReply) error
	grpc.ServerStream
}

type peerRetrieveServer struct {
	grpc.ServerStream
}

func (x *peerRetrieveServer) Send(m *RetrieveReply) error {
	return x.ServerStream.SendMsg(m)
}

func _Peer_Release_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReleaseMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeerServer).Release(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Peer/Release",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeerServer).Release(ctx, req.(*ReleaseMessage))
	}
	return interceptor(ctx, in, info, handler)
}

// Peer_ServiceDesc is the grpc.ServiceDesc for Peer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Peer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.Peer",
	HandlerType: (*PeerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Broadcast",
			Handler:    _Peer_Broadcast_Handler,
		},
		{
			MethodName: "Publish",
			Handler:    _Peer_Publish_Handler,
		},
		{
			MethodName: "Release",
			Handler:    _Peer_Release_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Join",
			Handler:       _Peer_Join_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Retrieve",
			Handler:       _Peer_Retrieve_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "ChatService/api/chat.proto",
}
