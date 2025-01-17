// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.2
// source: message/proto/pythonServicer.proto

package message

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

const (
	APMService_MainServiceRequest_FullMethodName        = "/protoData.APMService/MainServiceRequest"
	APMService_SubordinateServiceRequest_FullMethodName = "/protoData.APMService/SubordinateServiceRequest"
)

// APMServiceClient is the client API for APMService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type APMServiceClient interface {
	// Upload MainServiceStructure to Python Server
	// returned with pre-defined data structure and its formatted data for game use
	MainServiceRequest(ctx context.Context, in *MainServicerRequest, opts ...grpc.CallOption) (*ServiceResponse, error)
	SubordinateServiceRequest(ctx context.Context, in *SubordinateServicerRequest, opts ...grpc.CallOption) (*ServiceResponse, error)
}

type aPMServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAPMServiceClient(cc grpc.ClientConnInterface) APMServiceClient {
	return &aPMServiceClient{cc}
}

func (c *aPMServiceClient) MainServiceRequest(ctx context.Context, in *MainServicerRequest, opts ...grpc.CallOption) (*ServiceResponse, error) {
	out := new(ServiceResponse)
	err := c.cc.Invoke(ctx, APMService_MainServiceRequest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPMServiceClient) SubordinateServiceRequest(ctx context.Context, in *SubordinateServicerRequest, opts ...grpc.CallOption) (*ServiceResponse, error) {
	out := new(ServiceResponse)
	err := c.cc.Invoke(ctx, APMService_SubordinateServiceRequest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// APMServiceServer is the server API for APMService service.
// All implementations must embed UnimplementedAPMServiceServer
// for forward compatibility
type APMServiceServer interface {
	// Upload MainServiceStructure to Python Server
	// returned with pre-defined data structure and its formatted data for game use
	MainServiceRequest(context.Context, *MainServicerRequest) (*ServiceResponse, error)
	SubordinateServiceRequest(context.Context, *SubordinateServicerRequest) (*ServiceResponse, error)
	mustEmbedUnimplementedAPMServiceServer()
}

// UnimplementedAPMServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAPMServiceServer struct {
}

func (UnimplementedAPMServiceServer) MainServiceRequest(context.Context, *MainServicerRequest) (*ServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MainServiceRequest not implemented")
}
func (UnimplementedAPMServiceServer) SubordinateServiceRequest(context.Context, *SubordinateServicerRequest) (*ServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubordinateServiceRequest not implemented")
}
func (UnimplementedAPMServiceServer) mustEmbedUnimplementedAPMServiceServer() {}

// UnsafeAPMServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to APMServiceServer will
// result in compilation errors.
type UnsafeAPMServiceServer interface {
	mustEmbedUnimplementedAPMServiceServer()
}

func RegisterAPMServiceServer(s grpc.ServiceRegistrar, srv APMServiceServer) {
	s.RegisterService(&APMService_ServiceDesc, srv)
}

func _APMService_MainServiceRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MainServicerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APMServiceServer).MainServiceRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: APMService_MainServiceRequest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APMServiceServer).MainServiceRequest(ctx, req.(*MainServicerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _APMService_SubordinateServiceRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubordinateServicerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APMServiceServer).SubordinateServiceRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: APMService_SubordinateServiceRequest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APMServiceServer).SubordinateServiceRequest(ctx, req.(*SubordinateServicerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// APMService_ServiceDesc is the grpc.ServiceDesc for APMService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var APMService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protoData.APMService",
	HandlerType: (*APMServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MainServiceRequest",
			Handler:    _APMService_MainServiceRequest_Handler,
		},
		{
			MethodName: "SubordinateServiceRequest",
			Handler:    _APMService_SubordinateServiceRequest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "message/proto/pythonServicer.proto",
}

const (
	SubFunctionalService_EmbeddingNounChunks_FullMethodName = "/protoData.SubFunctionalService/EmbeddingNounChunks"
	SubFunctionalService_EmbeddingSentence_FullMethodName   = "/protoData.SubFunctionalService/EmbeddingSentence"
	SubFunctionalService_EmbeddingList_FullMethodName       = "/protoData.SubFunctionalService/EmbeddingList"
	SubFunctionalService_EmbeddingTopic_FullMethodName      = "/protoData.SubFunctionalService/EmbeddingTopic"
)

// SubFunctionalServiceClient is the client API for SubFunctionalService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SubFunctionalServiceClient interface {
	EmbeddingNounChunks(ctx context.Context, in *RequestPrompt, opts ...grpc.CallOption) (*WordList, error)
	EmbeddingSentence(ctx context.Context, in *RequestPrompt, opts ...grpc.CallOption) (*SentenceVec, error)
	EmbeddingList(ctx context.Context, in *RequestList, opts ...grpc.CallOption) (*WordList, error)
	EmbeddingTopic(ctx context.Context, in *RequestPrompt, opts ...grpc.CallOption) (*WordList, error)
}

type subFunctionalServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSubFunctionalServiceClient(cc grpc.ClientConnInterface) SubFunctionalServiceClient {
	return &subFunctionalServiceClient{cc}
}

func (c *subFunctionalServiceClient) EmbeddingNounChunks(ctx context.Context, in *RequestPrompt, opts ...grpc.CallOption) (*WordList, error) {
	out := new(WordList)
	err := c.cc.Invoke(ctx, SubFunctionalService_EmbeddingNounChunks_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subFunctionalServiceClient) EmbeddingSentence(ctx context.Context, in *RequestPrompt, opts ...grpc.CallOption) (*SentenceVec, error) {
	out := new(SentenceVec)
	err := c.cc.Invoke(ctx, SubFunctionalService_EmbeddingSentence_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subFunctionalServiceClient) EmbeddingList(ctx context.Context, in *RequestList, opts ...grpc.CallOption) (*WordList, error) {
	out := new(WordList)
	err := c.cc.Invoke(ctx, SubFunctionalService_EmbeddingList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subFunctionalServiceClient) EmbeddingTopic(ctx context.Context, in *RequestPrompt, opts ...grpc.CallOption) (*WordList, error) {
	out := new(WordList)
	err := c.cc.Invoke(ctx, SubFunctionalService_EmbeddingTopic_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SubFunctionalServiceServer is the server API for SubFunctionalService service.
// All implementations must embed UnimplementedSubFunctionalServiceServer
// for forward compatibility
type SubFunctionalServiceServer interface {
	EmbeddingNounChunks(context.Context, *RequestPrompt) (*WordList, error)
	EmbeddingSentence(context.Context, *RequestPrompt) (*SentenceVec, error)
	EmbeddingList(context.Context, *RequestList) (*WordList, error)
	EmbeddingTopic(context.Context, *RequestPrompt) (*WordList, error)
	mustEmbedUnimplementedSubFunctionalServiceServer()
}

// UnimplementedSubFunctionalServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSubFunctionalServiceServer struct {
}

func (UnimplementedSubFunctionalServiceServer) EmbeddingNounChunks(context.Context, *RequestPrompt) (*WordList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EmbeddingNounChunks not implemented")
}
func (UnimplementedSubFunctionalServiceServer) EmbeddingSentence(context.Context, *RequestPrompt) (*SentenceVec, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EmbeddingSentence not implemented")
}
func (UnimplementedSubFunctionalServiceServer) EmbeddingList(context.Context, *RequestList) (*WordList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EmbeddingList not implemented")
}
func (UnimplementedSubFunctionalServiceServer) EmbeddingTopic(context.Context, *RequestPrompt) (*WordList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EmbeddingTopic not implemented")
}
func (UnimplementedSubFunctionalServiceServer) mustEmbedUnimplementedSubFunctionalServiceServer() {}

// UnsafeSubFunctionalServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SubFunctionalServiceServer will
// result in compilation errors.
type UnsafeSubFunctionalServiceServer interface {
	mustEmbedUnimplementedSubFunctionalServiceServer()
}

func RegisterSubFunctionalServiceServer(s grpc.ServiceRegistrar, srv SubFunctionalServiceServer) {
	s.RegisterService(&SubFunctionalService_ServiceDesc, srv)
}

func _SubFunctionalService_EmbeddingNounChunks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestPrompt)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubFunctionalServiceServer).EmbeddingNounChunks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SubFunctionalService_EmbeddingNounChunks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubFunctionalServiceServer).EmbeddingNounChunks(ctx, req.(*RequestPrompt))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubFunctionalService_EmbeddingSentence_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestPrompt)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubFunctionalServiceServer).EmbeddingSentence(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SubFunctionalService_EmbeddingSentence_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubFunctionalServiceServer).EmbeddingSentence(ctx, req.(*RequestPrompt))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubFunctionalService_EmbeddingList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubFunctionalServiceServer).EmbeddingList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SubFunctionalService_EmbeddingList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubFunctionalServiceServer).EmbeddingList(ctx, req.(*RequestList))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubFunctionalService_EmbeddingTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestPrompt)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubFunctionalServiceServer).EmbeddingTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SubFunctionalService_EmbeddingTopic_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubFunctionalServiceServer).EmbeddingTopic(ctx, req.(*RequestPrompt))
	}
	return interceptor(ctx, in, info, handler)
}

// SubFunctionalService_ServiceDesc is the grpc.ServiceDesc for SubFunctionalService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SubFunctionalService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protoData.SubFunctionalService",
	HandlerType: (*SubFunctionalServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EmbeddingNounChunks",
			Handler:    _SubFunctionalService_EmbeddingNounChunks_Handler,
		},
		{
			MethodName: "EmbeddingSentence",
			Handler:    _SubFunctionalService_EmbeddingSentence_Handler,
		},
		{
			MethodName: "EmbeddingList",
			Handler:    _SubFunctionalService_EmbeddingList_Handler,
		},
		{
			MethodName: "EmbeddingTopic",
			Handler:    _SubFunctionalService_EmbeddingTopic_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "message/proto/pythonServicer.proto",
}
