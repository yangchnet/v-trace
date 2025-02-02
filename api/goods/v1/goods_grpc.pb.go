// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: goods/v1/goods.proto

package v1

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

// GoodsServiceClient is the client API for GoodsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GoodsServiceClient interface {
	// CreateClass 创建产品类型
	CreateClass(ctx context.Context, in *CreateClassRequest, opts ...grpc.CallOption) (*CreateClassResponse, error)
	// CreateSerial 创建产品批次
	CreateSerial(ctx context.Context, in *CreateSerialRequest, opts ...grpc.CallOption) (*CreateSerialResponse, error)
	// CreateGoods 批量创建产品
	BatchCreateGoods(ctx context.Context, in *BatchCreateGoodsRequest, opts ...grpc.CallOption) (*BatchCreateGoodsResponse, error)
	// GetGoods 根据id获取商品信息
	GetGoods(ctx context.Context, in *GetGoodsRequest, opts ...grpc.CallOption) (*GetGoodsResponse, error)
	// GetClass 根据id获取商品类型
	GetClass(ctx context.Context, in *GetClassRequest, opts ...grpc.CallOption) (*GetClassResponse, error)
	// GetSerial 根据id获取商品批次
	GetSerial(ctx context.Context, in *GetSerialRequest, opts ...grpc.CallOption) (*GetSerialResponse, error)
	// ListGoodsClass 列出产品类型
	ListGoodsClass(ctx context.Context, in *ListGoodsClassRequest, opts ...grpc.CallOption) (*ListGoodsClassResponse, error)
	// ListGoodsSerial 列出产品批次
	ListGoodsSerial(ctx context.Context, in *ListGoodsSerialRequest, opts ...grpc.CallOption) (*ListGoodsSerialResponse, error)
	// ListGoods 列出产品
	ListGoods(ctx context.Context, in *ListGoodsRequest, opts ...grpc.CallOption) (*ListGoodsResponse, error)
	// 更新产品类型信息
	UpdateGoodsClass(ctx context.Context, in *UpdateGoodsClassRequest, opts ...grpc.CallOption) (*UpdateGoodsClassResponse, error)
	// UpdateGoodsSerial 更新产品批次信息
	UpdateGoodsSerial(ctx context.Context, in *UpdateGoodsSerialRequest, opts ...grpc.CallOption) (*UpdateGoodsSerialResponse, error)
	// UpdateGoods 更新产品
	UpdateGoods(ctx context.Context, in *UpdateGoodsRequest, opts ...grpc.CallOption) (*UpdateGoodsResponse, error)
	// GetOrgOfX 获取产品/类型/批次所属企业
	GetOrgOfX(ctx context.Context, in *GetOrgOfXRequest, opts ...grpc.CallOption) (*GetOrgOfXResponse, error)
}

type goodsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGoodsServiceClient(cc grpc.ClientConnInterface) GoodsServiceClient {
	return &goodsServiceClient{cc}
}

func (c *goodsServiceClient) CreateClass(ctx context.Context, in *CreateClassRequest, opts ...grpc.CallOption) (*CreateClassResponse, error) {
	out := new(CreateClassResponse)
	err := c.cc.Invoke(ctx, "/goods.v1.GoodsService/CreateClass", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsServiceClient) CreateSerial(ctx context.Context, in *CreateSerialRequest, opts ...grpc.CallOption) (*CreateSerialResponse, error) {
	out := new(CreateSerialResponse)
	err := c.cc.Invoke(ctx, "/goods.v1.GoodsService/CreateSerial", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsServiceClient) BatchCreateGoods(ctx context.Context, in *BatchCreateGoodsRequest, opts ...grpc.CallOption) (*BatchCreateGoodsResponse, error) {
	out := new(BatchCreateGoodsResponse)
	err := c.cc.Invoke(ctx, "/goods.v1.GoodsService/BatchCreateGoods", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsServiceClient) GetGoods(ctx context.Context, in *GetGoodsRequest, opts ...grpc.CallOption) (*GetGoodsResponse, error) {
	out := new(GetGoodsResponse)
	err := c.cc.Invoke(ctx, "/goods.v1.GoodsService/GetGoods", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsServiceClient) GetClass(ctx context.Context, in *GetClassRequest, opts ...grpc.CallOption) (*GetClassResponse, error) {
	out := new(GetClassResponse)
	err := c.cc.Invoke(ctx, "/goods.v1.GoodsService/GetClass", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsServiceClient) GetSerial(ctx context.Context, in *GetSerialRequest, opts ...grpc.CallOption) (*GetSerialResponse, error) {
	out := new(GetSerialResponse)
	err := c.cc.Invoke(ctx, "/goods.v1.GoodsService/GetSerial", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsServiceClient) ListGoodsClass(ctx context.Context, in *ListGoodsClassRequest, opts ...grpc.CallOption) (*ListGoodsClassResponse, error) {
	out := new(ListGoodsClassResponse)
	err := c.cc.Invoke(ctx, "/goods.v1.GoodsService/ListGoodsClass", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsServiceClient) ListGoodsSerial(ctx context.Context, in *ListGoodsSerialRequest, opts ...grpc.CallOption) (*ListGoodsSerialResponse, error) {
	out := new(ListGoodsSerialResponse)
	err := c.cc.Invoke(ctx, "/goods.v1.GoodsService/ListGoodsSerial", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsServiceClient) ListGoods(ctx context.Context, in *ListGoodsRequest, opts ...grpc.CallOption) (*ListGoodsResponse, error) {
	out := new(ListGoodsResponse)
	err := c.cc.Invoke(ctx, "/goods.v1.GoodsService/ListGoods", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsServiceClient) UpdateGoodsClass(ctx context.Context, in *UpdateGoodsClassRequest, opts ...grpc.CallOption) (*UpdateGoodsClassResponse, error) {
	out := new(UpdateGoodsClassResponse)
	err := c.cc.Invoke(ctx, "/goods.v1.GoodsService/UpdateGoodsClass", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsServiceClient) UpdateGoodsSerial(ctx context.Context, in *UpdateGoodsSerialRequest, opts ...grpc.CallOption) (*UpdateGoodsSerialResponse, error) {
	out := new(UpdateGoodsSerialResponse)
	err := c.cc.Invoke(ctx, "/goods.v1.GoodsService/UpdateGoodsSerial", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsServiceClient) UpdateGoods(ctx context.Context, in *UpdateGoodsRequest, opts ...grpc.CallOption) (*UpdateGoodsResponse, error) {
	out := new(UpdateGoodsResponse)
	err := c.cc.Invoke(ctx, "/goods.v1.GoodsService/UpdateGoods", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsServiceClient) GetOrgOfX(ctx context.Context, in *GetOrgOfXRequest, opts ...grpc.CallOption) (*GetOrgOfXResponse, error) {
	out := new(GetOrgOfXResponse)
	err := c.cc.Invoke(ctx, "/goods.v1.GoodsService/GetOrgOfX", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GoodsServiceServer is the server API for GoodsService service.
// All implementations must embed UnimplementedGoodsServiceServer
// for forward compatibility
type GoodsServiceServer interface {
	// CreateClass 创建产品类型
	CreateClass(context.Context, *CreateClassRequest) (*CreateClassResponse, error)
	// CreateSerial 创建产品批次
	CreateSerial(context.Context, *CreateSerialRequest) (*CreateSerialResponse, error)
	// CreateGoods 批量创建产品
	BatchCreateGoods(context.Context, *BatchCreateGoodsRequest) (*BatchCreateGoodsResponse, error)
	// GetGoods 根据id获取商品信息
	GetGoods(context.Context, *GetGoodsRequest) (*GetGoodsResponse, error)
	// GetClass 根据id获取商品类型
	GetClass(context.Context, *GetClassRequest) (*GetClassResponse, error)
	// GetSerial 根据id获取商品批次
	GetSerial(context.Context, *GetSerialRequest) (*GetSerialResponse, error)
	// ListGoodsClass 列出产品类型
	ListGoodsClass(context.Context, *ListGoodsClassRequest) (*ListGoodsClassResponse, error)
	// ListGoodsSerial 列出产品批次
	ListGoodsSerial(context.Context, *ListGoodsSerialRequest) (*ListGoodsSerialResponse, error)
	// ListGoods 列出产品
	ListGoods(context.Context, *ListGoodsRequest) (*ListGoodsResponse, error)
	// 更新产品类型信息
	UpdateGoodsClass(context.Context, *UpdateGoodsClassRequest) (*UpdateGoodsClassResponse, error)
	// UpdateGoodsSerial 更新产品批次信息
	UpdateGoodsSerial(context.Context, *UpdateGoodsSerialRequest) (*UpdateGoodsSerialResponse, error)
	// UpdateGoods 更新产品
	UpdateGoods(context.Context, *UpdateGoodsRequest) (*UpdateGoodsResponse, error)
	// GetOrgOfX 获取产品/类型/批次所属企业
	GetOrgOfX(context.Context, *GetOrgOfXRequest) (*GetOrgOfXResponse, error)
	mustEmbedUnimplementedGoodsServiceServer()
}

// UnimplementedGoodsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGoodsServiceServer struct {
}

func (UnimplementedGoodsServiceServer) CreateClass(context.Context, *CreateClassRequest) (*CreateClassResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateClass not implemented")
}
func (UnimplementedGoodsServiceServer) CreateSerial(context.Context, *CreateSerialRequest) (*CreateSerialResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSerial not implemented")
}
func (UnimplementedGoodsServiceServer) BatchCreateGoods(context.Context, *BatchCreateGoodsRequest) (*BatchCreateGoodsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchCreateGoods not implemented")
}
func (UnimplementedGoodsServiceServer) GetGoods(context.Context, *GetGoodsRequest) (*GetGoodsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGoods not implemented")
}
func (UnimplementedGoodsServiceServer) GetClass(context.Context, *GetClassRequest) (*GetClassResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClass not implemented")
}
func (UnimplementedGoodsServiceServer) GetSerial(context.Context, *GetSerialRequest) (*GetSerialResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSerial not implemented")
}
func (UnimplementedGoodsServiceServer) ListGoodsClass(context.Context, *ListGoodsClassRequest) (*ListGoodsClassResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListGoodsClass not implemented")
}
func (UnimplementedGoodsServiceServer) ListGoodsSerial(context.Context, *ListGoodsSerialRequest) (*ListGoodsSerialResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListGoodsSerial not implemented")
}
func (UnimplementedGoodsServiceServer) ListGoods(context.Context, *ListGoodsRequest) (*ListGoodsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListGoods not implemented")
}
func (UnimplementedGoodsServiceServer) UpdateGoodsClass(context.Context, *UpdateGoodsClassRequest) (*UpdateGoodsClassResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGoodsClass not implemented")
}
func (UnimplementedGoodsServiceServer) UpdateGoodsSerial(context.Context, *UpdateGoodsSerialRequest) (*UpdateGoodsSerialResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGoodsSerial not implemented")
}
func (UnimplementedGoodsServiceServer) UpdateGoods(context.Context, *UpdateGoodsRequest) (*UpdateGoodsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGoods not implemented")
}
func (UnimplementedGoodsServiceServer) GetOrgOfX(context.Context, *GetOrgOfXRequest) (*GetOrgOfXResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrgOfX not implemented")
}
func (UnimplementedGoodsServiceServer) mustEmbedUnimplementedGoodsServiceServer() {}

// UnsafeGoodsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GoodsServiceServer will
// result in compilation errors.
type UnsafeGoodsServiceServer interface {
	mustEmbedUnimplementedGoodsServiceServer()
}

func RegisterGoodsServiceServer(s grpc.ServiceRegistrar, srv GoodsServiceServer) {
	s.RegisterService(&GoodsService_ServiceDesc, srv)
}

func _GoodsService_CreateClass_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateClassRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServiceServer).CreateClass(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goods.v1.GoodsService/CreateClass",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServiceServer).CreateClass(ctx, req.(*CreateClassRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoodsService_CreateSerial_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSerialRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServiceServer).CreateSerial(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goods.v1.GoodsService/CreateSerial",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServiceServer).CreateSerial(ctx, req.(*CreateSerialRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoodsService_BatchCreateGoods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchCreateGoodsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServiceServer).BatchCreateGoods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goods.v1.GoodsService/BatchCreateGoods",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServiceServer).BatchCreateGoods(ctx, req.(*BatchCreateGoodsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoodsService_GetGoods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGoodsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServiceServer).GetGoods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goods.v1.GoodsService/GetGoods",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServiceServer).GetGoods(ctx, req.(*GetGoodsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoodsService_GetClass_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClassRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServiceServer).GetClass(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goods.v1.GoodsService/GetClass",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServiceServer).GetClass(ctx, req.(*GetClassRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoodsService_GetSerial_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSerialRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServiceServer).GetSerial(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goods.v1.GoodsService/GetSerial",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServiceServer).GetSerial(ctx, req.(*GetSerialRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoodsService_ListGoodsClass_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListGoodsClassRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServiceServer).ListGoodsClass(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goods.v1.GoodsService/ListGoodsClass",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServiceServer).ListGoodsClass(ctx, req.(*ListGoodsClassRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoodsService_ListGoodsSerial_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListGoodsSerialRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServiceServer).ListGoodsSerial(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goods.v1.GoodsService/ListGoodsSerial",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServiceServer).ListGoodsSerial(ctx, req.(*ListGoodsSerialRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoodsService_ListGoods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListGoodsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServiceServer).ListGoods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goods.v1.GoodsService/ListGoods",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServiceServer).ListGoods(ctx, req.(*ListGoodsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoodsService_UpdateGoodsClass_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGoodsClassRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServiceServer).UpdateGoodsClass(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goods.v1.GoodsService/UpdateGoodsClass",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServiceServer).UpdateGoodsClass(ctx, req.(*UpdateGoodsClassRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoodsService_UpdateGoodsSerial_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGoodsSerialRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServiceServer).UpdateGoodsSerial(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goods.v1.GoodsService/UpdateGoodsSerial",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServiceServer).UpdateGoodsSerial(ctx, req.(*UpdateGoodsSerialRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoodsService_UpdateGoods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGoodsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServiceServer).UpdateGoods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goods.v1.GoodsService/UpdateGoods",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServiceServer).UpdateGoods(ctx, req.(*UpdateGoodsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoodsService_GetOrgOfX_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrgOfXRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServiceServer).GetOrgOfX(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goods.v1.GoodsService/GetOrgOfX",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServiceServer).GetOrgOfX(ctx, req.(*GetOrgOfXRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GoodsService_ServiceDesc is the grpc.ServiceDesc for GoodsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GoodsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "goods.v1.GoodsService",
	HandlerType: (*GoodsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateClass",
			Handler:    _GoodsService_CreateClass_Handler,
		},
		{
			MethodName: "CreateSerial",
			Handler:    _GoodsService_CreateSerial_Handler,
		},
		{
			MethodName: "BatchCreateGoods",
			Handler:    _GoodsService_BatchCreateGoods_Handler,
		},
		{
			MethodName: "GetGoods",
			Handler:    _GoodsService_GetGoods_Handler,
		},
		{
			MethodName: "GetClass",
			Handler:    _GoodsService_GetClass_Handler,
		},
		{
			MethodName: "GetSerial",
			Handler:    _GoodsService_GetSerial_Handler,
		},
		{
			MethodName: "ListGoodsClass",
			Handler:    _GoodsService_ListGoodsClass_Handler,
		},
		{
			MethodName: "ListGoodsSerial",
			Handler:    _GoodsService_ListGoodsSerial_Handler,
		},
		{
			MethodName: "ListGoods",
			Handler:    _GoodsService_ListGoods_Handler,
		},
		{
			MethodName: "UpdateGoodsClass",
			Handler:    _GoodsService_UpdateGoodsClass_Handler,
		},
		{
			MethodName: "UpdateGoodsSerial",
			Handler:    _GoodsService_UpdateGoodsSerial_Handler,
		},
		{
			MethodName: "UpdateGoods",
			Handler:    _GoodsService_UpdateGoods_Handler,
		},
		{
			MethodName: "GetOrgOfX",
			Handler:    _GoodsService_GetOrgOfX_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "goods/v1/goods.proto",
}
