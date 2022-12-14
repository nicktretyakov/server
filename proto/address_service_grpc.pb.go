// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.2
// source: address_service.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AddressServiceClient is the client API for AddressService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AddressServiceClient interface {
	AddAttachment(ctx context.Context, in *CreateAttachmentRequest, opts ...grpc.CallOption) (*CreateAttachmentResponse, error)
	RemoveAttachment(ctx context.Context, in *RemoveAttachmentRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	RenameAttachment(ctx context.Context, in *RenameAttachmentRequest, opts ...grpc.CallOption) (*CreateAttachmentResponse, error)
	AddLink(ctx context.Context, in *AddLinkRequest, opts ...grpc.CallOption) (*AddLinkResponse, error)
	UpdateLink(ctx context.Context, in *UpdateLinkRequest, opts ...grpc.CallOption) (*UpdateLinkResponse, error)
	RemoveLink(ctx context.Context, in *RemoveLinkRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	AddParticipant(ctx context.Context, in *AddParticipantRequest, opts ...grpc.CallOption) (*AddParticipantResponse, error)
	UpdateParticipant(ctx context.Context, in *UpdateParticipantRequest, opts ...grpc.CallOption) (*UpdateParticipantResponse, error)
	RemoveParticipant(ctx context.Context, in *RemoveParticipantRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type addressServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAddressServiceClient(cc grpc.ClientConnInterface) AddressServiceClient {
	return &addressServiceClient{cc}
}

func (c *addressServiceClient) AddAttachment(ctx context.Context, in *CreateAttachmentRequest, opts ...grpc.CallOption) (*CreateAttachmentResponse, error) {
	out := new(CreateAttachmentResponse)
	err := c.cc.Invoke(ctx, "/booking.AddressService/AddAttachment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addressServiceClient) RemoveAttachment(ctx context.Context, in *RemoveAttachmentRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/booking.AddressService/RemoveAttachment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addressServiceClient) RenameAttachment(ctx context.Context, in *RenameAttachmentRequest, opts ...grpc.CallOption) (*CreateAttachmentResponse, error) {
	out := new(CreateAttachmentResponse)
	err := c.cc.Invoke(ctx, "/booking.AddressService/RenameAttachment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addressServiceClient) AddLink(ctx context.Context, in *AddLinkRequest, opts ...grpc.CallOption) (*AddLinkResponse, error) {
	out := new(AddLinkResponse)
	err := c.cc.Invoke(ctx, "/booking.AddressService/AddLink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addressServiceClient) UpdateLink(ctx context.Context, in *UpdateLinkRequest, opts ...grpc.CallOption) (*UpdateLinkResponse, error) {
	out := new(UpdateLinkResponse)
	err := c.cc.Invoke(ctx, "/booking.AddressService/UpdateLink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addressServiceClient) RemoveLink(ctx context.Context, in *RemoveLinkRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/booking.AddressService/RemoveLink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addressServiceClient) AddParticipant(ctx context.Context, in *AddParticipantRequest, opts ...grpc.CallOption) (*AddParticipantResponse, error) {
	out := new(AddParticipantResponse)
	err := c.cc.Invoke(ctx, "/booking.AddressService/AddParticipant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addressServiceClient) UpdateParticipant(ctx context.Context, in *UpdateParticipantRequest, opts ...grpc.CallOption) (*UpdateParticipantResponse, error) {
	out := new(UpdateParticipantResponse)
	err := c.cc.Invoke(ctx, "/booking.AddressService/UpdateParticipant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addressServiceClient) RemoveParticipant(ctx context.Context, in *RemoveParticipantRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/booking.AddressService/RemoveParticipant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AddressServiceServer is the server API for AddressService service.
// All implementations must embed UnimplementedAddressServiceServer
// for forward compatibility
type AddressServiceServer interface {
	AddAttachment(context.Context, *CreateAttachmentRequest) (*CreateAttachmentResponse, error)
	RemoveAttachment(context.Context, *RemoveAttachmentRequest) (*emptypb.Empty, error)
	RenameAttachment(context.Context, *RenameAttachmentRequest) (*CreateAttachmentResponse, error)
	AddLink(context.Context, *AddLinkRequest) (*AddLinkResponse, error)
	UpdateLink(context.Context, *UpdateLinkRequest) (*UpdateLinkResponse, error)
	RemoveLink(context.Context, *RemoveLinkRequest) (*emptypb.Empty, error)
	AddParticipant(context.Context, *AddParticipantRequest) (*AddParticipantResponse, error)
	UpdateParticipant(context.Context, *UpdateParticipantRequest) (*UpdateParticipantResponse, error)
	RemoveParticipant(context.Context, *RemoveParticipantRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedAddressServiceServer()
}

// UnimplementedAddressServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAddressServiceServer struct {
}

func (UnimplementedAddressServiceServer) AddAttachment(context.Context, *CreateAttachmentRequest) (*CreateAttachmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddAttachment not implemented")
}
func (UnimplementedAddressServiceServer) RemoveAttachment(context.Context, *RemoveAttachmentRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveAttachment not implemented")
}
func (UnimplementedAddressServiceServer) RenameAttachment(context.Context, *RenameAttachmentRequest) (*CreateAttachmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RenameAttachment not implemented")
}
func (UnimplementedAddressServiceServer) AddLink(context.Context, *AddLinkRequest) (*AddLinkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddLink not implemented")
}
func (UnimplementedAddressServiceServer) UpdateLink(context.Context, *UpdateLinkRequest) (*UpdateLinkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateLink not implemented")
}
func (UnimplementedAddressServiceServer) RemoveLink(context.Context, *RemoveLinkRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveLink not implemented")
}
func (UnimplementedAddressServiceServer) AddParticipant(context.Context, *AddParticipantRequest) (*AddParticipantResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddParticipant not implemented")
}
func (UnimplementedAddressServiceServer) UpdateParticipant(context.Context, *UpdateParticipantRequest) (*UpdateParticipantResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateParticipant not implemented")
}
func (UnimplementedAddressServiceServer) RemoveParticipant(context.Context, *RemoveParticipantRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveParticipant not implemented")
}
func (UnimplementedAddressServiceServer) mustEmbedUnimplementedAddressServiceServer() {}

// UnsafeAddressServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AddressServiceServer will
// result in compilation errors.
type UnsafeAddressServiceServer interface {
	mustEmbedUnimplementedAddressServiceServer()
}

func RegisterAddressServiceServer(s grpc.ServiceRegistrar, srv AddressServiceServer) {
	s.RegisterService(&AddressService_ServiceDesc, srv)
}

func _AddressService_AddAttachment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAttachmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddressServiceServer).AddAttachment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/booking.AddressService/AddAttachment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddressServiceServer).AddAttachment(ctx, req.(*CreateAttachmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AddressService_RemoveAttachment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveAttachmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddressServiceServer).RemoveAttachment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/booking.AddressService/RemoveAttachment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddressServiceServer).RemoveAttachment(ctx, req.(*RemoveAttachmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AddressService_RenameAttachment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RenameAttachmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddressServiceServer).RenameAttachment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/booking.AddressService/RenameAttachment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddressServiceServer).RenameAttachment(ctx, req.(*RenameAttachmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AddressService_AddLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddressServiceServer).AddLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/booking.AddressService/AddLink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddressServiceServer).AddLink(ctx, req.(*AddLinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AddressService_UpdateLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddressServiceServer).UpdateLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/booking.AddressService/UpdateLink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddressServiceServer).UpdateLink(ctx, req.(*UpdateLinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AddressService_RemoveLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddressServiceServer).RemoveLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/booking.AddressService/RemoveLink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddressServiceServer).RemoveLink(ctx, req.(*RemoveLinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AddressService_AddParticipant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddParticipantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddressServiceServer).AddParticipant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/booking.AddressService/AddParticipant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddressServiceServer).AddParticipant(ctx, req.(*AddParticipantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AddressService_UpdateParticipant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateParticipantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddressServiceServer).UpdateParticipant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/booking.AddressService/UpdateParticipant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddressServiceServer).UpdateParticipant(ctx, req.(*UpdateParticipantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AddressService_RemoveParticipant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveParticipantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddressServiceServer).RemoveParticipant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/booking.AddressService/RemoveParticipant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddressServiceServer).RemoveParticipant(ctx, req.(*RemoveParticipantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AddressService_ServiceDesc is the grpc.ServiceDesc for AddressService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AddressService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "booking.AddressService",
	HandlerType: (*AddressServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddAttachment",
			Handler:    _AddressService_AddAttachment_Handler,
		},
		{
			MethodName: "RemoveAttachment",
			Handler:    _AddressService_RemoveAttachment_Handler,
		},
		{
			MethodName: "RenameAttachment",
			Handler:    _AddressService_RenameAttachment_Handler,
		},
		{
			MethodName: "AddLink",
			Handler:    _AddressService_AddLink_Handler,
		},
		{
			MethodName: "UpdateLink",
			Handler:    _AddressService_UpdateLink_Handler,
		},
		{
			MethodName: "RemoveLink",
			Handler:    _AddressService_RemoveLink_Handler,
		},
		{
			MethodName: "AddParticipant",
			Handler:    _AddressService_AddParticipant_Handler,
		},
		{
			MethodName: "UpdateParticipant",
			Handler:    _AddressService_UpdateParticipant_Handler,
		},
		{
			MethodName: "RemoveParticipant",
			Handler:    _AddressService_RemoveParticipant_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "address_service.proto",
}
