// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package REDACTEDapi

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// ACLServiceClient is the client API for ACLService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ACLServiceClient interface {
	Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error)
	Write(ctx context.Context, in *WriteRequest, opts ...grpc.CallOption) (*WriteResponse, error)
	Check(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*CheckResponse, error)
	ContentChangeCheck(ctx context.Context, in *ContentChangeCheckRequest, opts ...grpc.CallOption) (*CheckResponse, error)
	Expand(ctx context.Context, in *ExpandRequest, opts ...grpc.CallOption) (*ExpandResponse, error)
}

type aCLServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewACLServiceClient(cc grpc.ClientConnInterface) ACLServiceClient {
	return &aCLServiceClient{cc}
}

func (c *aCLServiceClient) Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error) {
	out := new(ReadResponse)
	err := c.cc.Invoke(ctx, "/ACLService/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aCLServiceClient) Write(ctx context.Context, in *WriteRequest, opts ...grpc.CallOption) (*WriteResponse, error) {
	out := new(WriteResponse)
	err := c.cc.Invoke(ctx, "/ACLService/Write", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aCLServiceClient) Check(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*CheckResponse, error) {
	out := new(CheckResponse)
	err := c.cc.Invoke(ctx, "/ACLService/Check", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aCLServiceClient) ContentChangeCheck(ctx context.Context, in *ContentChangeCheckRequest, opts ...grpc.CallOption) (*CheckResponse, error) {
	out := new(CheckResponse)
	err := c.cc.Invoke(ctx, "/ACLService/ContentChangeCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aCLServiceClient) Expand(ctx context.Context, in *ExpandRequest, opts ...grpc.CallOption) (*ExpandResponse, error) {
	out := new(ExpandResponse)
	err := c.cc.Invoke(ctx, "/ACLService/Expand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ACLServiceServer is the server API for ACLService service.
// All implementations must embed UnimplementedACLServiceServer
// for forward compatibility
type ACLServiceServer interface {
	Read(context.Context, *ReadRequest) (*ReadResponse, error)
	Write(context.Context, *WriteRequest) (*WriteResponse, error)
	Check(context.Context, *CheckRequest) (*CheckResponse, error)
	ContentChangeCheck(context.Context, *ContentChangeCheckRequest) (*CheckResponse, error)
	Expand(context.Context, *ExpandRequest) (*ExpandResponse, error)
	mustEmbedUnimplementedACLServiceServer()
}

// UnimplementedACLServiceServer must be embedded to have forward compatible implementations.
type UnimplementedACLServiceServer struct {
}

func (UnimplementedACLServiceServer) Read(context.Context, *ReadRequest) (*ReadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}
func (UnimplementedACLServiceServer) Write(context.Context, *WriteRequest) (*WriteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Write not implemented")
}
func (UnimplementedACLServiceServer) Check(context.Context, *CheckRequest) (*CheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Check not implemented")
}
func (UnimplementedACLServiceServer) ContentChangeCheck(context.Context, *ContentChangeCheckRequest) (*CheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ContentChangeCheck not implemented")
}
func (UnimplementedACLServiceServer) Expand(context.Context, *ExpandRequest) (*ExpandResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Expand not implemented")
}
func (UnimplementedACLServiceServer) mustEmbedUnimplementedACLServiceServer() {}

// UnsafeACLServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ACLServiceServer will
// result in compilation errors.
type UnsafeACLServiceServer interface {
	mustEmbedUnimplementedACLServiceServer()
}

func RegisterACLServiceServer(s grpc.ServiceRegistrar, srv ACLServiceServer) {
	s.RegisterService(&_ACLService_serviceDesc, srv)
}

func _ACLService_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ACLServiceServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ACLService/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ACLServiceServer).Read(ctx, req.(*ReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ACLService_Write_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WriteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ACLServiceServer).Write(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ACLService/Write",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ACLServiceServer).Write(ctx, req.(*WriteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ACLService_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ACLServiceServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ACLService/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ACLServiceServer).Check(ctx, req.(*CheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ACLService_ContentChangeCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContentChangeCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ACLServiceServer).ContentChangeCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ACLService/ContentChangeCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ACLServiceServer).ContentChangeCheck(ctx, req.(*ContentChangeCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ACLService_Expand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExpandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ACLServiceServer).Expand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ACLService/Expand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ACLServiceServer).Expand(ctx, req.(*ExpandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ACLService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ACLService",
	HandlerType: (*ACLServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Read",
			Handler:    _ACLService_Read_Handler,
		},
		{
			MethodName: "Write",
			Handler:    _ACLService_Write_Handler,
		},
		{
			MethodName: "Check",
			Handler:    _ACLService_Check_Handler,
		},
		{
			MethodName: "ContentChangeCheck",
			Handler:    _ACLService_ContentChangeCheck_Handler,
		},
		{
			MethodName: "Expand",
			Handler:    _ACLService_Expand_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "REDACTEDapi/api/acl_service.proto",
}