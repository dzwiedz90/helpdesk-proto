// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: services/tickets/ticket.proto

package tickets

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

// TicketServiceClient is the client API for TicketService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TicketServiceClient interface {
	CreateTicket(ctx context.Context, in *CreateTicketRequest, opts ...grpc.CallOption) (*CreateTicketResponse, error)
	GetAllTickets(ctx context.Context, in *GetAllTicketsRequest, opts ...grpc.CallOption) (*GetAllTicketsResponse, error)
	GetAllTicketsByUser(ctx context.Context, in *GetAllTicketsByUserRequest, opts ...grpc.CallOption) (*GetAllTicketsByUserResponse, error)
	GetAllTicketsByAgent(ctx context.Context, in *GetAllTicketsByAgentRequest, opts ...grpc.CallOption) (*GetAllTicketsByAgentResponse, error)
	GetTicket(ctx context.Context, in *GetTicketRequest, opts ...grpc.CallOption) (*GetTicketResponse, error)
	UpdateTicket(ctx context.Context, in *UpdateTicketRequest, opts ...grpc.CallOption) (*UpdateTicketResponse, error)
	DeleteTicket(ctx context.Context, in *DeleteTicketRequest, opts ...grpc.CallOption) (*DeleteTicketResponse, error)
}

type ticketServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTicketServiceClient(cc grpc.ClientConnInterface) TicketServiceClient {
	return &ticketServiceClient{cc}
}

func (c *ticketServiceClient) CreateTicket(ctx context.Context, in *CreateTicketRequest, opts ...grpc.CallOption) (*CreateTicketResponse, error) {
	out := new(CreateTicketResponse)
	err := c.cc.Invoke(ctx, "/TicketService/CreateTicket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketServiceClient) GetAllTickets(ctx context.Context, in *GetAllTicketsRequest, opts ...grpc.CallOption) (*GetAllTicketsResponse, error) {
	out := new(GetAllTicketsResponse)
	err := c.cc.Invoke(ctx, "/TicketService/GetAllTickets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketServiceClient) GetAllTicketsByUser(ctx context.Context, in *GetAllTicketsByUserRequest, opts ...grpc.CallOption) (*GetAllTicketsByUserResponse, error) {
	out := new(GetAllTicketsByUserResponse)
	err := c.cc.Invoke(ctx, "/TicketService/GetAllTicketsByUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketServiceClient) GetAllTicketsByAgent(ctx context.Context, in *GetAllTicketsByAgentRequest, opts ...grpc.CallOption) (*GetAllTicketsByAgentResponse, error) {
	out := new(GetAllTicketsByAgentResponse)
	err := c.cc.Invoke(ctx, "/TicketService/GetAllTicketsByAgent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketServiceClient) GetTicket(ctx context.Context, in *GetTicketRequest, opts ...grpc.CallOption) (*GetTicketResponse, error) {
	out := new(GetTicketResponse)
	err := c.cc.Invoke(ctx, "/TicketService/GetTicket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketServiceClient) UpdateTicket(ctx context.Context, in *UpdateTicketRequest, opts ...grpc.CallOption) (*UpdateTicketResponse, error) {
	out := new(UpdateTicketResponse)
	err := c.cc.Invoke(ctx, "/TicketService/UpdateTicket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketServiceClient) DeleteTicket(ctx context.Context, in *DeleteTicketRequest, opts ...grpc.CallOption) (*DeleteTicketResponse, error) {
	out := new(DeleteTicketResponse)
	err := c.cc.Invoke(ctx, "/TicketService/DeleteTicket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TicketServiceServer is the server API for TicketService service.
// All implementations must embed UnimplementedTicketServiceServer
// for forward compatibility
type TicketServiceServer interface {
	CreateTicket(context.Context, *CreateTicketRequest) (*CreateTicketResponse, error)
	GetAllTickets(context.Context, *GetAllTicketsRequest) (*GetAllTicketsResponse, error)
	GetAllTicketsByUser(context.Context, *GetAllTicketsByUserRequest) (*GetAllTicketsByUserResponse, error)
	GetAllTicketsByAgent(context.Context, *GetAllTicketsByAgentRequest) (*GetAllTicketsByAgentResponse, error)
	GetTicket(context.Context, *GetTicketRequest) (*GetTicketResponse, error)
	UpdateTicket(context.Context, *UpdateTicketRequest) (*UpdateTicketResponse, error)
	DeleteTicket(context.Context, *DeleteTicketRequest) (*DeleteTicketResponse, error)
	mustEmbedUnimplementedTicketServiceServer()
}

// UnimplementedTicketServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTicketServiceServer struct {
}

func (UnimplementedTicketServiceServer) CreateTicket(context.Context, *CreateTicketRequest) (*CreateTicketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTicket not implemented")
}
func (UnimplementedTicketServiceServer) GetAllTickets(context.Context, *GetAllTicketsRequest) (*GetAllTicketsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllTickets not implemented")
}
func (UnimplementedTicketServiceServer) GetAllTicketsByUser(context.Context, *GetAllTicketsByUserRequest) (*GetAllTicketsByUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllTicketsByUser not implemented")
}
func (UnimplementedTicketServiceServer) GetAllTicketsByAgent(context.Context, *GetAllTicketsByAgentRequest) (*GetAllTicketsByAgentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllTicketsByAgent not implemented")
}
func (UnimplementedTicketServiceServer) GetTicket(context.Context, *GetTicketRequest) (*GetTicketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTicket not implemented")
}
func (UnimplementedTicketServiceServer) UpdateTicket(context.Context, *UpdateTicketRequest) (*UpdateTicketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTicket not implemented")
}
func (UnimplementedTicketServiceServer) DeleteTicket(context.Context, *DeleteTicketRequest) (*DeleteTicketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTicket not implemented")
}
func (UnimplementedTicketServiceServer) mustEmbedUnimplementedTicketServiceServer() {}

// UnsafeTicketServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TicketServiceServer will
// result in compilation errors.
type UnsafeTicketServiceServer interface {
	mustEmbedUnimplementedTicketServiceServer()
}

func RegisterTicketServiceServer(s grpc.ServiceRegistrar, srv TicketServiceServer) {
	s.RegisterService(&TicketService_ServiceDesc, srv)
}

func _TicketService_CreateTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTicketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketServiceServer).CreateTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TicketService/CreateTicket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketServiceServer).CreateTicket(ctx, req.(*CreateTicketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TicketService_GetAllTickets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllTicketsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketServiceServer).GetAllTickets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TicketService/GetAllTickets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketServiceServer).GetAllTickets(ctx, req.(*GetAllTicketsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TicketService_GetAllTicketsByUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllTicketsByUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketServiceServer).GetAllTicketsByUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TicketService/GetAllTicketsByUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketServiceServer).GetAllTicketsByUser(ctx, req.(*GetAllTicketsByUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TicketService_GetAllTicketsByAgent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllTicketsByAgentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketServiceServer).GetAllTicketsByAgent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TicketService/GetAllTicketsByAgent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketServiceServer).GetAllTicketsByAgent(ctx, req.(*GetAllTicketsByAgentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TicketService_GetTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTicketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketServiceServer).GetTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TicketService/GetTicket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketServiceServer).GetTicket(ctx, req.(*GetTicketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TicketService_UpdateTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTicketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketServiceServer).UpdateTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TicketService/UpdateTicket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketServiceServer).UpdateTicket(ctx, req.(*UpdateTicketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TicketService_DeleteTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTicketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketServiceServer).DeleteTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TicketService/DeleteTicket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketServiceServer).DeleteTicket(ctx, req.(*DeleteTicketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TicketService_ServiceDesc is the grpc.ServiceDesc for TicketService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TicketService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "TicketService",
	HandlerType: (*TicketServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTicket",
			Handler:    _TicketService_CreateTicket_Handler,
		},
		{
			MethodName: "GetAllTickets",
			Handler:    _TicketService_GetAllTickets_Handler,
		},
		{
			MethodName: "GetAllTicketsByUser",
			Handler:    _TicketService_GetAllTicketsByUser_Handler,
		},
		{
			MethodName: "GetAllTicketsByAgent",
			Handler:    _TicketService_GetAllTicketsByAgent_Handler,
		},
		{
			MethodName: "GetTicket",
			Handler:    _TicketService_GetTicket_Handler,
		},
		{
			MethodName: "UpdateTicket",
			Handler:    _TicketService_UpdateTicket_Handler,
		},
		{
			MethodName: "DeleteTicket",
			Handler:    _TicketService_DeleteTicket_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services/tickets/ticket.proto",
}
