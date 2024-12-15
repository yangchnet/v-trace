// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: ca/v1/ca.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type UserType int32

const (
	UserType_root      UserType = 0
	UserType_ca        UserType = 1
	UserType_admin     UserType = 2
	UserType_client    UserType = 3
	UserType_consensus UserType = 4
	UserType_common    UserType = 5
)

// Enum value maps for UserType.
var (
	UserType_name = map[int32]string{
		0: "root",
		1: "ca",
		2: "admin",
		3: "client",
		4: "consensus",
		5: "common",
	}
	UserType_value = map[string]int32{
		"root":      0,
		"ca":        1,
		"admin":     2,
		"client":    3,
		"consensus": 4,
		"common":    5,
	}
)

func (x UserType) Enum() *UserType {
	p := new(UserType)
	*p = x
	return p
}

func (x UserType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserType) Descriptor() protoreflect.EnumDescriptor {
	return file_ca_v1_ca_proto_enumTypes[0].Descriptor()
}

func (UserType) Type() protoreflect.EnumType {
	return &file_ca_v1_ca_proto_enumTypes[0]
}

func (x UserType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserType.Descriptor instead.
func (UserType) EnumDescriptor() ([]byte, []int) {
	return file_ca_v1_ca_proto_rawDescGZIP(), []int{0}
}

type CertUsage int32

const (
	CertUsage_sign    CertUsage = 0
	CertUsage_tls     CertUsage = 1
	CertUsage_tls_enc CertUsage = 2
)

// Enum value maps for CertUsage.
var (
	CertUsage_name = map[int32]string{
		0: "sign",
		1: "tls",
		2: "tls_enc",
	}
	CertUsage_value = map[string]int32{
		"sign":    0,
		"tls":     1,
		"tls_enc": 2,
	}
)

func (x CertUsage) Enum() *CertUsage {
	p := new(CertUsage)
	*p = x
	return p
}

func (x CertUsage) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CertUsage) Descriptor() protoreflect.EnumDescriptor {
	return file_ca_v1_ca_proto_enumTypes[1].Descriptor()
}

func (CertUsage) Type() protoreflect.EnumType {
	return &file_ca_v1_ca_proto_enumTypes[1]
}

func (x CertUsage) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CertUsage.Descriptor instead.
func (CertUsage) EnumDescriptor() ([]byte, []int) {
	return file_ca_v1_ca_proto_rawDescGZIP(), []int{1}
}

type GenCertRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 组织id
	OrgId string `protobuf:"bytes,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	// 唯一用户名
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	// 用户类型(root, ca, admin, client, consensus, common)
	UserType UserType `protobuf:"varint,3,opt,name=user_type,json=userType,proto3,enum=ca.v1.UserType" json:"user_type,omitempty"`
	// 证书用途(sign, tls, tls-enc)
	CertUsage []CertUsage `protobuf:"varint,4,rep,packed,name=cert_usage,json=certUsage,proto3,enum=ca.v1.CertUsage" json:"cert_usage,omitempty"`
	// 密钥密码
	PrivateKeyPwd string `protobuf:"bytes,5,opt,name=private_key_pwd,json=privateKeyPwd,proto3" json:"private_key_pwd,omitempty"`
	// 证书字段-国家
	Country string `protobuf:"bytes,6,opt,name=country,proto3" json:"country,omitempty"`
	// 证书字段-城市
	Locality string `protobuf:"bytes,7,opt,name=locality,proto3" json:"locality,omitempty"`
	// 证书字段-省份
	Province string `protobuf:"bytes,8,opt,name=province,proto3" json:"province,omitempty"`
}

func (x *GenCertRequest) Reset() {
	*x = GenCertRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ca_v1_ca_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenCertRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenCertRequest) ProtoMessage() {}

func (x *GenCertRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ca_v1_ca_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenCertRequest.ProtoReflect.Descriptor instead.
func (*GenCertRequest) Descriptor() ([]byte, []int) {
	return file_ca_v1_ca_proto_rawDescGZIP(), []int{0}
}

func (x *GenCertRequest) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

func (x *GenCertRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *GenCertRequest) GetUserType() UserType {
	if x != nil {
		return x.UserType
	}
	return UserType_root
}

func (x *GenCertRequest) GetCertUsage() []CertUsage {
	if x != nil {
		return x.CertUsage
	}
	return nil
}

func (x *GenCertRequest) GetPrivateKeyPwd() string {
	if x != nil {
		return x.PrivateKeyPwd
	}
	return ""
}

func (x *GenCertRequest) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *GenCertRequest) GetLocality() string {
	if x != nil {
		return x.Locality
	}
	return ""
}

func (x *GenCertRequest) GetProvince() string {
	if x != nil {
		return x.Province
	}
	return ""
}

type GenCertResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 证书内容
	Cert *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=cert,proto3" json:"cert,omitempty"`
	// 密钥内容
	PrivateKey *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=private_key,json=privateKey,proto3" json:"private_key,omitempty"`
	// 用户名
	Username *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	// tls证书
	TlsCert *wrapperspb.StringValue `protobuf:"bytes,4,opt,name=tls_cert,json=tlsCert,proto3" json:"tls_cert,omitempty"`
	// tls密钥
	TlsKey *wrapperspb.StringValue `protobuf:"bytes,5,opt,name=tls_key,json=tlsKey,proto3" json:"tls_key,omitempty"`
}

func (x *GenCertResponse) Reset() {
	*x = GenCertResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ca_v1_ca_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenCertResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenCertResponse) ProtoMessage() {}

func (x *GenCertResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ca_v1_ca_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenCertResponse.ProtoReflect.Descriptor instead.
func (*GenCertResponse) Descriptor() ([]byte, []int) {
	return file_ca_v1_ca_proto_rawDescGZIP(), []int{1}
}

func (x *GenCertResponse) GetCert() *wrapperspb.StringValue {
	if x != nil {
		return x.Cert
	}
	return nil
}

func (x *GenCertResponse) GetPrivateKey() *wrapperspb.StringValue {
	if x != nil {
		return x.PrivateKey
	}
	return nil
}

func (x *GenCertResponse) GetUsername() *wrapperspb.StringValue {
	if x != nil {
		return x.Username
	}
	return nil
}

func (x *GenCertResponse) GetTlsCert() *wrapperspb.StringValue {
	if x != nil {
		return x.TlsCert
	}
	return nil
}

func (x *GenCertResponse) GetTlsKey() *wrapperspb.StringValue {
	if x != nil {
		return x.TlsKey
	}
	return nil
}

type GetCertRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 组织id
	OrgId string `protobuf:"bytes,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	// 唯一用户名
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	// 用户类型(root, ca, admin, client, consensus, common)
	UserType UserType `protobuf:"varint,3,opt,name=user_type,json=userType,proto3,enum=ca.v1.UserType" json:"user_type,omitempty"`
	// 证书用途(sign, tls, tls-enc)
	CertUsage []CertUsage `protobuf:"varint,4,rep,packed,name=cert_usage,json=certUsage,proto3,enum=ca.v1.CertUsage" json:"cert_usage,omitempty"`
}

func (x *GetCertRequest) Reset() {
	*x = GetCertRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ca_v1_ca_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCertRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCertRequest) ProtoMessage() {}

func (x *GetCertRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ca_v1_ca_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCertRequest.ProtoReflect.Descriptor instead.
func (*GetCertRequest) Descriptor() ([]byte, []int) {
	return file_ca_v1_ca_proto_rawDescGZIP(), []int{2}
}

func (x *GetCertRequest) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

func (x *GetCertRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *GetCertRequest) GetUserType() UserType {
	if x != nil {
		return x.UserType
	}
	return UserType_root
}

func (x *GetCertRequest) GetCertUsage() []CertUsage {
	if x != nil {
		return x.CertUsage
	}
	return nil
}

type GetCertResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 证书内容
	Cert *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=cert,proto3" json:"cert,omitempty"`
	// 密钥内容
	PrivateKey *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=private_key,json=privateKey,proto3" json:"private_key,omitempty"`
	// 用户名
	Username *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	// tls证书
	TlsCert *wrapperspb.StringValue `protobuf:"bytes,4,opt,name=tls_cert,json=tlsCert,proto3" json:"tls_cert,omitempty"`
	// tls密钥
	TlsKey *wrapperspb.StringValue `protobuf:"bytes,5,opt,name=tls_key,json=tlsKey,proto3" json:"tls_key,omitempty"`
}

func (x *GetCertResponse) Reset() {
	*x = GetCertResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ca_v1_ca_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCertResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCertResponse) ProtoMessage() {}

func (x *GetCertResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ca_v1_ca_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCertResponse.ProtoReflect.Descriptor instead.
func (*GetCertResponse) Descriptor() ([]byte, []int) {
	return file_ca_v1_ca_proto_rawDescGZIP(), []int{3}
}

func (x *GetCertResponse) GetCert() *wrapperspb.StringValue {
	if x != nil {
		return x.Cert
	}
	return nil
}

func (x *GetCertResponse) GetPrivateKey() *wrapperspb.StringValue {
	if x != nil {
		return x.PrivateKey
	}
	return nil
}

func (x *GetCertResponse) GetUsername() *wrapperspb.StringValue {
	if x != nil {
		return x.Username
	}
	return nil
}

func (x *GetCertResponse) GetTlsCert() *wrapperspb.StringValue {
	if x != nil {
		return x.TlsCert
	}
	return nil
}

func (x *GetCertResponse) GetTlsKey() *wrapperspb.StringValue {
	if x != nil {
		return x.TlsKey
	}
	return nil
}

var File_ca_v1_ca_proto protoreflect.FileDescriptor

var file_ca_v1_ca_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x63, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x05, 0x63, 0x61, 0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9c, 0x02, 0x0a, 0x0e, 0x47, 0x65, 0x6e, 0x43,
	0x65, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x6f, 0x72,
	0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x72, 0x67, 0x49,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2c, 0x0a,
	0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x0f, 0x2e, 0x63, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2f, 0x0a, 0x0a, 0x63,
	0x65, 0x72, 0x74, 0x5f, 0x75, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0e, 0x32,
	0x10, 0x2e, 0x63, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x65, 0x72, 0x74, 0x55, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x09, 0x63, 0x65, 0x72, 0x74, 0x55, 0x73, 0x61, 0x67, 0x65, 0x12, 0x26, 0x0a, 0x0f,
	0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x70, 0x77, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65,
	0x79, 0x50, 0x77, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x1a,
	0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x22, 0xac, 0x02, 0x0a, 0x0f, 0x47, 0x65, 0x6e, 0x43, 0x65,
	0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x04, 0x63, 0x65,
	0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x63, 0x65, 0x72, 0x74, 0x12, 0x3d, 0x0a, 0x0b,
	0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x0a, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x38, 0x0a, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x37, 0x0a, 0x08, 0x74, 0x6c, 0x73, 0x5f, 0x63, 0x65, 0x72,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x07, 0x74, 0x6c, 0x73, 0x43, 0x65, 0x72, 0x74, 0x12, 0x35,
	0x0a, 0x07, 0x74, 0x6c, 0x73, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x74,
	0x6c, 0x73, 0x4b, 0x65, 0x79, 0x22, 0xa2, 0x01, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x43, 0x65, 0x72,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x6f, 0x72, 0x67, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x72, 0x67, 0x49, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x09, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f,
	0x2e, 0x63, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2f, 0x0a, 0x0a, 0x63, 0x65, 0x72,
	0x74, 0x5f, 0x75, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x10, 0x2e,
	0x63, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x65, 0x72, 0x74, 0x55, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x09, 0x63, 0x65, 0x72, 0x74, 0x55, 0x73, 0x61, 0x67, 0x65, 0x22, 0xac, 0x02, 0x0a, 0x0f, 0x47,
	0x65, 0x74, 0x43, 0x65, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30,
	0x0a, 0x04, 0x63, 0x65, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x63, 0x65, 0x72, 0x74,
	0x12, 0x3d, 0x0a, 0x0b, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x0a, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x12,
	0x38, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x37, 0x0a, 0x08, 0x74, 0x6c, 0x73,
	0x5f, 0x63, 0x65, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x07, 0x74, 0x6c, 0x73, 0x43, 0x65,
	0x72, 0x74, 0x12, 0x35, 0x0a, 0x07, 0x74, 0x6c, 0x73, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x06, 0x74, 0x6c, 0x73, 0x4b, 0x65, 0x79, 0x2a, 0x4e, 0x0a, 0x08, 0x55, 0x73, 0x65,
	0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x72, 0x6f, 0x6f, 0x74, 0x10, 0x00, 0x12,
	0x06, 0x0a, 0x02, 0x63, 0x61, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x10, 0x03, 0x12, 0x0d,
	0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x10, 0x04, 0x12, 0x0a, 0x0a,
	0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x10, 0x05, 0x2a, 0x2b, 0x0a, 0x09, 0x43, 0x65, 0x72,
	0x74, 0x55, 0x73, 0x61, 0x67, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x73, 0x69, 0x67, 0x6e, 0x10, 0x00,
	0x12, 0x07, 0x0a, 0x03, 0x74, 0x6c, 0x73, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x74, 0x6c, 0x73,
	0x5f, 0x65, 0x6e, 0x63, 0x10, 0x02, 0x32, 0x83, 0x01, 0x0a, 0x09, 0x43, 0x41, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x3a, 0x0a, 0x07, 0x47, 0x65, 0x6e, 0x43, 0x65, 0x72, 0x74, 0x12,
	0x15, 0x2e, 0x63, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x6e, 0x43, 0x65, 0x72, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x63, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x6e, 0x43, 0x65, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x3a, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x43, 0x65, 0x72, 0x74, 0x12, 0x15, 0x2e, 0x63, 0x61,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x65, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x63, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x65,
	0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2a, 0x5a, 0x28,
	0x67, 0x69, 0x74, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x71, 0x63, 0x69, 0x69, 0x70, 0x2d,
	0x69, 0x63, 0x70, 0x2f, 0x76, 0x2d, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x63, 0x61, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ca_v1_ca_proto_rawDescOnce sync.Once
	file_ca_v1_ca_proto_rawDescData = file_ca_v1_ca_proto_rawDesc
)

func file_ca_v1_ca_proto_rawDescGZIP() []byte {
	file_ca_v1_ca_proto_rawDescOnce.Do(func() {
		file_ca_v1_ca_proto_rawDescData = protoimpl.X.CompressGZIP(file_ca_v1_ca_proto_rawDescData)
	})
	return file_ca_v1_ca_proto_rawDescData
}

var file_ca_v1_ca_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_ca_v1_ca_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_ca_v1_ca_proto_goTypes = []interface{}{
	(UserType)(0),                  // 0: ca.v1.UserType
	(CertUsage)(0),                 // 1: ca.v1.CertUsage
	(*GenCertRequest)(nil),         // 2: ca.v1.GenCertRequest
	(*GenCertResponse)(nil),        // 3: ca.v1.GenCertResponse
	(*GetCertRequest)(nil),         // 4: ca.v1.GetCertRequest
	(*GetCertResponse)(nil),        // 5: ca.v1.GetCertResponse
	(*wrapperspb.StringValue)(nil), // 6: google.protobuf.StringValue
}
var file_ca_v1_ca_proto_depIdxs = []int32{
	0,  // 0: ca.v1.GenCertRequest.user_type:type_name -> ca.v1.UserType
	1,  // 1: ca.v1.GenCertRequest.cert_usage:type_name -> ca.v1.CertUsage
	6,  // 2: ca.v1.GenCertResponse.cert:type_name -> google.protobuf.StringValue
	6,  // 3: ca.v1.GenCertResponse.private_key:type_name -> google.protobuf.StringValue
	6,  // 4: ca.v1.GenCertResponse.username:type_name -> google.protobuf.StringValue
	6,  // 5: ca.v1.GenCertResponse.tls_cert:type_name -> google.protobuf.StringValue
	6,  // 6: ca.v1.GenCertResponse.tls_key:type_name -> google.protobuf.StringValue
	0,  // 7: ca.v1.GetCertRequest.user_type:type_name -> ca.v1.UserType
	1,  // 8: ca.v1.GetCertRequest.cert_usage:type_name -> ca.v1.CertUsage
	6,  // 9: ca.v1.GetCertResponse.cert:type_name -> google.protobuf.StringValue
	6,  // 10: ca.v1.GetCertResponse.private_key:type_name -> google.protobuf.StringValue
	6,  // 11: ca.v1.GetCertResponse.username:type_name -> google.protobuf.StringValue
	6,  // 12: ca.v1.GetCertResponse.tls_cert:type_name -> google.protobuf.StringValue
	6,  // 13: ca.v1.GetCertResponse.tls_key:type_name -> google.protobuf.StringValue
	2,  // 14: ca.v1.CAService.GenCert:input_type -> ca.v1.GenCertRequest
	4,  // 15: ca.v1.CAService.GetCert:input_type -> ca.v1.GetCertRequest
	3,  // 16: ca.v1.CAService.GenCert:output_type -> ca.v1.GenCertResponse
	5,  // 17: ca.v1.CAService.GetCert:output_type -> ca.v1.GetCertResponse
	16, // [16:18] is the sub-list for method output_type
	14, // [14:16] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_ca_v1_ca_proto_init() }
func file_ca_v1_ca_proto_init() {
	if File_ca_v1_ca_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ca_v1_ca_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenCertRequest); i {
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
		file_ca_v1_ca_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenCertResponse); i {
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
		file_ca_v1_ca_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCertRequest); i {
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
		file_ca_v1_ca_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCertResponse); i {
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
			RawDescriptor: file_ca_v1_ca_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ca_v1_ca_proto_goTypes,
		DependencyIndexes: file_ca_v1_ca_proto_depIdxs,
		EnumInfos:         file_ca_v1_ca_proto_enumTypes,
		MessageInfos:      file_ca_v1_ca_proto_msgTypes,
	}.Build()
	File_ca_v1_ca_proto = out.File
	file_ca_v1_ca_proto_rawDesc = nil
	file_ca_v1_ca_proto_goTypes = nil
	file_ca_v1_ca_proto_depIdxs = nil
}
