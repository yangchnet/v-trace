// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: goods/v1/model.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/durationpb"
	_ "google.golang.org/protobuf/types/known/structpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Class struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id
	ID *wrapperspb.Int64Value `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	// 产品类别名称
	Name *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// 产品类别描述(json)
	Des *wrapperspb.BytesValue `protobuf:"bytes,3,opt,name=des,proto3" json:"des,omitempty"`
	// 状态
	State *wrapperspb.StringValue `protobuf:"bytes,4,opt,name=state,proto3" json:"state,omitempty"`
	// 创建者
	Creator *wrapperspb.StringValue `protobuf:"bytes,5,opt,name=creator,proto3" json:"creator,omitempty"`
	// 生产企业
	OrgId *wrapperspb.Int64Value `protobuf:"bytes,6,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	// 商标
	Tm *wrapperspb.StringValue `protobuf:"bytes,7,opt,name=tm,proto3" json:"tm,omitempty"`
	// 原料
	MaterialId *wrapperspb.Int32Value `protobuf:"bytes,8,opt,name=material_id,json=materialId,proto3" json:"material_id,omitempty"`
}

func (x *Class) Reset() {
	*x = Class{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goods_v1_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Class) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Class) ProtoMessage() {}

func (x *Class) ProtoReflect() protoreflect.Message {
	mi := &file_goods_v1_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Class.ProtoReflect.Descriptor instead.
func (*Class) Descriptor() ([]byte, []int) {
	return file_goods_v1_model_proto_rawDescGZIP(), []int{0}
}

func (x *Class) GetID() *wrapperspb.Int64Value {
	if x != nil {
		return x.ID
	}
	return nil
}

func (x *Class) GetName() *wrapperspb.StringValue {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *Class) GetDes() *wrapperspb.BytesValue {
	if x != nil {
		return x.Des
	}
	return nil
}

func (x *Class) GetState() *wrapperspb.StringValue {
	if x != nil {
		return x.State
	}
	return nil
}

func (x *Class) GetCreator() *wrapperspb.StringValue {
	if x != nil {
		return x.Creator
	}
	return nil
}

func (x *Class) GetOrgId() *wrapperspb.Int64Value {
	if x != nil {
		return x.OrgId
	}
	return nil
}

func (x *Class) GetTm() *wrapperspb.StringValue {
	if x != nil {
		return x.Tm
	}
	return nil
}

func (x *Class) GetMaterialId() *wrapperspb.Int32Value {
	if x != nil {
		return x.MaterialId
	}
	return nil
}

type Serial struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id
	ID *wrapperspb.Int64Value `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	// 生产日期
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// 状态
	State *wrapperspb.StringValue `protobuf:"bytes,4,opt,name=state,proto3" json:"state,omitempty"`
	// 创建者
	Creator *wrapperspb.StringValue `protobuf:"bytes,5,opt,name=creator,proto3" json:"creator,omitempty"`
	// 类别
	Class *Class `protobuf:"bytes,6,opt,name=class,proto3" json:"class,omitempty"`
}

func (x *Serial) Reset() {
	*x = Serial{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goods_v1_model_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Serial) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Serial) ProtoMessage() {}

func (x *Serial) ProtoReflect() protoreflect.Message {
	mi := &file_goods_v1_model_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Serial.ProtoReflect.Descriptor instead.
func (*Serial) Descriptor() ([]byte, []int) {
	return file_goods_v1_model_proto_rawDescGZIP(), []int{1}
}

func (x *Serial) GetID() *wrapperspb.Int64Value {
	if x != nil {
		return x.ID
	}
	return nil
}

func (x *Serial) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *Serial) GetState() *wrapperspb.StringValue {
	if x != nil {
		return x.State
	}
	return nil
}

func (x *Serial) GetCreator() *wrapperspb.StringValue {
	if x != nil {
		return x.Creator
	}
	return nil
}

func (x *Serial) GetClass() *Class {
	if x != nil {
		return x.Class
	}
	return nil
}

type Goods struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id
	ID *wrapperspb.Int64Value `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	// 状态
	State *wrapperspb.StringValue `protobuf:"bytes,4,opt,name=state,proto3" json:"state,omitempty"`
	// 创建者
	Creator *wrapperspb.StringValue `protobuf:"bytes,5,opt,name=creator,proto3" json:"creator,omitempty"`
	// 批次
	Serial *Serial `protobuf:"bytes,6,opt,name=serial,proto3" json:"serial,omitempty"`
}

func (x *Goods) Reset() {
	*x = Goods{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goods_v1_model_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Goods) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Goods) ProtoMessage() {}

func (x *Goods) ProtoReflect() protoreflect.Message {
	mi := &file_goods_v1_model_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Goods.ProtoReflect.Descriptor instead.
func (*Goods) Descriptor() ([]byte, []int) {
	return file_goods_v1_model_proto_rawDescGZIP(), []int{2}
}

func (x *Goods) GetID() *wrapperspb.Int64Value {
	if x != nil {
		return x.ID
	}
	return nil
}

func (x *Goods) GetState() *wrapperspb.StringValue {
	if x != nil {
		return x.State
	}
	return nil
}

func (x *Goods) GetCreator() *wrapperspb.StringValue {
	if x != nil {
		return x.Creator
	}
	return nil
}

func (x *Goods) GetSerial() *Serial {
	if x != nil {
		return x.Serial
	}
	return nil
}

var File_goods_v1_model_proto protoreflect.FileDescriptor

var file_goods_v1_model_proto_rawDesc = []byte{
	0x0a, 0x14, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x2e, 0x76, 0x31,
	0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xa1, 0x03, 0x0a, 0x05, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x12, 0x2b, 0x0a, 0x02, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x02, 0x49, 0x44, 0x12, 0x30, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2d, 0x0a, 0x03, 0x64, 0x65, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x79, 0x74, 0x65, 0x73, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x03, 0x64, 0x65, 0x73, 0x12, 0x32, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x36, 0x0a, 0x07, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x6f, 0x72, 0x12, 0x32, 0x0a, 0x06, 0x6f, 0x72, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x05, 0x6f, 0x72, 0x67, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x02, 0x74, 0x6d, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x02, 0x74, 0x6d, 0x12, 0x3c, 0x0a, 0x0b, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61,
	0x6c, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74,
	0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61,
	0x6c, 0x49, 0x64, 0x22, 0x82, 0x02, 0x0a, 0x06, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x12, 0x2b,
	0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74,
	0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x02, 0x49, 0x44, 0x12, 0x38, 0x0a, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x32, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x36, 0x0a, 0x07, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f,
	0x72, 0x12, 0x25, 0x0a, 0x05, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x61, 0x73,
	0x73, 0x52, 0x05, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x22, 0xca, 0x01, 0x0a, 0x05, 0x47, 0x6f, 0x6f,
	0x64, 0x73, 0x12, 0x2b, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x02, 0x49, 0x44, 0x12,
	0x32, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x36, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x28, 0x0a, 0x06, 0x73,
	0x65, 0x72, 0x69, 0x61, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x67, 0x6f,
	0x6f, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x52, 0x06, 0x73,
	0x65, 0x72, 0x69, 0x61, 0x6c, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x65, 0x65, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x71, 0x63, 0x69, 0x69, 0x70, 0x2d, 0x69, 0x63, 0x70, 0x2f, 0x76, 0x2d, 0x74,
	0x72, 0x61, 0x63, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x2f, 0x76,
	0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_goods_v1_model_proto_rawDescOnce sync.Once
	file_goods_v1_model_proto_rawDescData = file_goods_v1_model_proto_rawDesc
)

func file_goods_v1_model_proto_rawDescGZIP() []byte {
	file_goods_v1_model_proto_rawDescOnce.Do(func() {
		file_goods_v1_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_goods_v1_model_proto_rawDescData)
	})
	return file_goods_v1_model_proto_rawDescData
}

var file_goods_v1_model_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_goods_v1_model_proto_goTypes = []interface{}{
	(*Class)(nil),                  // 0: goods.v1.Class
	(*Serial)(nil),                 // 1: goods.v1.Serial
	(*Goods)(nil),                  // 2: goods.v1.Goods
	(*wrapperspb.Int64Value)(nil),  // 3: google.protobuf.Int64Value
	(*wrapperspb.StringValue)(nil), // 4: google.protobuf.StringValue
	(*wrapperspb.BytesValue)(nil),  // 5: google.protobuf.BytesValue
	(*wrapperspb.Int32Value)(nil),  // 6: google.protobuf.Int32Value
	(*timestamppb.Timestamp)(nil),  // 7: google.protobuf.Timestamp
}
var file_goods_v1_model_proto_depIdxs = []int32{
	3,  // 0: goods.v1.Class.ID:type_name -> google.protobuf.Int64Value
	4,  // 1: goods.v1.Class.name:type_name -> google.protobuf.StringValue
	5,  // 2: goods.v1.Class.des:type_name -> google.protobuf.BytesValue
	4,  // 3: goods.v1.Class.state:type_name -> google.protobuf.StringValue
	4,  // 4: goods.v1.Class.creator:type_name -> google.protobuf.StringValue
	3,  // 5: goods.v1.Class.org_id:type_name -> google.protobuf.Int64Value
	4,  // 6: goods.v1.Class.tm:type_name -> google.protobuf.StringValue
	6,  // 7: goods.v1.Class.material_id:type_name -> google.protobuf.Int32Value
	3,  // 8: goods.v1.Serial.ID:type_name -> google.protobuf.Int64Value
	7,  // 9: goods.v1.Serial.timestamp:type_name -> google.protobuf.Timestamp
	4,  // 10: goods.v1.Serial.state:type_name -> google.protobuf.StringValue
	4,  // 11: goods.v1.Serial.creator:type_name -> google.protobuf.StringValue
	0,  // 12: goods.v1.Serial.class:type_name -> goods.v1.Class
	3,  // 13: goods.v1.Goods.ID:type_name -> google.protobuf.Int64Value
	4,  // 14: goods.v1.Goods.state:type_name -> google.protobuf.StringValue
	4,  // 15: goods.v1.Goods.creator:type_name -> google.protobuf.StringValue
	1,  // 16: goods.v1.Goods.serial:type_name -> goods.v1.Serial
	17, // [17:17] is the sub-list for method output_type
	17, // [17:17] is the sub-list for method input_type
	17, // [17:17] is the sub-list for extension type_name
	17, // [17:17] is the sub-list for extension extendee
	0,  // [0:17] is the sub-list for field type_name
}

func init() { file_goods_v1_model_proto_init() }
func file_goods_v1_model_proto_init() {
	if File_goods_v1_model_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_goods_v1_model_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Class); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_goods_v1_model_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Serial); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_goods_v1_model_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Goods); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_goods_v1_model_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_goods_v1_model_proto_goTypes,
		DependencyIndexes: file_goods_v1_model_proto_depIdxs,
		MessageInfos:      file_goods_v1_model_proto_msgTypes,
	}.Build()
	File_goods_v1_model_proto = out.File
	file_goods_v1_model_proto_rawDesc = nil
	file_goods_v1_model_proto_goTypes = nil
	file_goods_v1_model_proto_depIdxs = nil
}
