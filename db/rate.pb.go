// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: rate.proto

package db

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ConversionRate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FromCurrency       string  `protobuf:"bytes,1,opt,name=from_currency,json=fromCurrency,proto3" json:"from_currency,omitempty"`
	ToCurrency         string  `protobuf:"bytes,2,opt,name=to_currency,json=toCurrency,proto3" json:"to_currency,omitempty"`
	Rate               float64 `protobuf:"fixed64,3,opt,name=rate,proto3" json:"rate,omitempty"`
	DateUpdatedUnixUtc int64   `protobuf:"varint,4,opt,name=date_updated_unix_utc,json=dateUpdatedUnixUtc,proto3" json:"date_updated_unix_utc,omitempty"`
	Uuid               string  `protobuf:"bytes,5,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *ConversionRate) Reset() {
	*x = ConversionRate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConversionRate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConversionRate) ProtoMessage() {}

func (x *ConversionRate) ProtoReflect() protoreflect.Message {
	mi := &file_rate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConversionRate.ProtoReflect.Descriptor instead.
func (*ConversionRate) Descriptor() ([]byte, []int) {
	return file_rate_proto_rawDescGZIP(), []int{0}
}

func (x *ConversionRate) GetFromCurrency() string {
	if x != nil {
		return x.FromCurrency
	}
	return ""
}

func (x *ConversionRate) GetToCurrency() string {
	if x != nil {
		return x.ToCurrency
	}
	return ""
}

func (x *ConversionRate) GetRate() float64 {
	if x != nil {
		return x.Rate
	}
	return 0
}

func (x *ConversionRate) GetDateUpdatedUnixUtc() int64 {
	if x != nil {
		return x.DateUpdatedUnixUtc
	}
	return 0
}

func (x *ConversionRate) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type CreateConversionRateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConversionRate []*ConversionRate `protobuf:"bytes,1,rep,name=conversion_rate,json=conversionRate,proto3" json:"conversion_rate,omitempty"`
}

func (x *CreateConversionRateRequest) Reset() {
	*x = CreateConversionRateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rate_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateConversionRateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateConversionRateRequest) ProtoMessage() {}

func (x *CreateConversionRateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rate_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateConversionRateRequest.ProtoReflect.Descriptor instead.
func (*CreateConversionRateRequest) Descriptor() ([]byte, []int) {
	return file_rate_proto_rawDescGZIP(), []int{1}
}

func (x *CreateConversionRateRequest) GetConversionRate() []*ConversionRate {
	if x != nil {
		return x.ConversionRate
	}
	return nil
}

type CreateConversionRateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConversionRate []*ConversionRate `protobuf:"bytes,1,rep,name=conversion_rate,json=conversionRate,proto3" json:"conversion_rate,omitempty"`
}

func (x *CreateConversionRateResponse) Reset() {
	*x = CreateConversionRateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rate_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateConversionRateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateConversionRateResponse) ProtoMessage() {}

func (x *CreateConversionRateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rate_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateConversionRateResponse.ProtoReflect.Descriptor instead.
func (*CreateConversionRateResponse) Descriptor() ([]byte, []int) {
	return file_rate_proto_rawDescGZIP(), []int{2}
}

func (x *CreateConversionRateResponse) GetConversionRate() []*ConversionRate {
	if x != nil {
		return x.ConversionRate
	}
	return nil
}

type ReadConversionRateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FromUnixUtc int64 `protobuf:"varint,1,opt,name=from_unix_utc,json=fromUnixUtc,proto3" json:"from_unix_utc,omitempty"`
	ToUnixUtc   int64 `protobuf:"varint,2,opt,name=to_unix_utc,json=toUnixUtc,proto3" json:"to_unix_utc,omitempty"`
}

func (x *ReadConversionRateRequest) Reset() {
	*x = ReadConversionRateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rate_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadConversionRateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadConversionRateRequest) ProtoMessage() {}

func (x *ReadConversionRateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rate_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadConversionRateRequest.ProtoReflect.Descriptor instead.
func (*ReadConversionRateRequest) Descriptor() ([]byte, []int) {
	return file_rate_proto_rawDescGZIP(), []int{3}
}

func (x *ReadConversionRateRequest) GetFromUnixUtc() int64 {
	if x != nil {
		return x.FromUnixUtc
	}
	return 0
}

func (x *ReadConversionRateRequest) GetToUnixUtc() int64 {
	if x != nil {
		return x.ToUnixUtc
	}
	return 0
}

type ReadConversionRateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConversionRate []*ConversionRate `protobuf:"bytes,1,rep,name=conversion_rate,json=conversionRate,proto3" json:"conversion_rate,omitempty"`
}

func (x *ReadConversionRateResponse) Reset() {
	*x = ReadConversionRateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rate_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadConversionRateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadConversionRateResponse) ProtoMessage() {}

func (x *ReadConversionRateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rate_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadConversionRateResponse.ProtoReflect.Descriptor instead.
func (*ReadConversionRateResponse) Descriptor() ([]byte, []int) {
	return file_rate_proto_rawDescGZIP(), []int{4}
}

func (x *ReadConversionRateResponse) GetConversionRate() []*ConversionRate {
	if x != nil {
		return x.ConversionRate
	}
	return nil
}

type UpdateConversionRateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConversionRate *ConversionRate `protobuf:"bytes,1,opt,name=conversion_rate,json=conversionRate,proto3" json:"conversion_rate,omitempty"`
}

func (x *UpdateConversionRateRequest) Reset() {
	*x = UpdateConversionRateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rate_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateConversionRateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateConversionRateRequest) ProtoMessage() {}

func (x *UpdateConversionRateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rate_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateConversionRateRequest.ProtoReflect.Descriptor instead.
func (*UpdateConversionRateRequest) Descriptor() ([]byte, []int) {
	return file_rate_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateConversionRateRequest) GetConversionRate() *ConversionRate {
	if x != nil {
		return x.ConversionRate
	}
	return nil
}

type UpdateConversionRateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateConversionRateResponse) Reset() {
	*x = UpdateConversionRateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rate_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateConversionRateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateConversionRateResponse) ProtoMessage() {}

func (x *UpdateConversionRateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rate_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateConversionRateResponse.ProtoReflect.Descriptor instead.
func (*UpdateConversionRateResponse) Descriptor() ([]byte, []int) {
	return file_rate_proto_rawDescGZIP(), []int{6}
}

type DeleteConversionRateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *DeleteConversionRateRequest) Reset() {
	*x = DeleteConversionRateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rate_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteConversionRateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteConversionRateRequest) ProtoMessage() {}

func (x *DeleteConversionRateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rate_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteConversionRateRequest.ProtoReflect.Descriptor instead.
func (*DeleteConversionRateRequest) Descriptor() ([]byte, []int) {
	return file_rate_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteConversionRateRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type DeleteConversionRateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteConversionRateResponse) Reset() {
	*x = DeleteConversionRateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rate_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteConversionRateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteConversionRateResponse) ProtoMessage() {}

func (x *DeleteConversionRateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rate_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteConversionRateResponse.ProtoReflect.Descriptor instead.
func (*DeleteConversionRateResponse) Descriptor() ([]byte, []int) {
	return file_rate_proto_rawDescGZIP(), []int{8}
}

var File_rate_proto protoreflect.FileDescriptor

var file_rate_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x64, 0x62,
	0x22, 0xb1, 0x01, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52,
	0x61, 0x74, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x66, 0x72, 0x6f, 0x6d,
	0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x5f, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74,
	0x6f, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x61, 0x74,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x72, 0x61, 0x74, 0x65, 0x12, 0x31, 0x0a,
	0x15, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x75, 0x6e,
	0x69, 0x78, 0x5f, 0x75, 0x74, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x12, 0x64, 0x61,
	0x74, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x55, 0x6e, 0x69, 0x78, 0x55, 0x74, 0x63,
	0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x75, 0x75, 0x69, 0x64, 0x22, 0x5a, 0x0a, 0x1b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f,
	0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x3b, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x64,
	0x62, 0x2e, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x61, 0x74, 0x65,
	0x52, 0x0e, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x61, 0x74, 0x65,
	0x22, 0x5b, 0x0a, 0x1c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x3b, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x72,
	0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x64, 0x62, 0x2e, 0x43,
	0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x61, 0x74, 0x65, 0x52, 0x0e, 0x63,
	0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x61, 0x74, 0x65, 0x22, 0x5f, 0x0a,
	0x19, 0x52, 0x65, 0x61, 0x64, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0d, 0x66, 0x72,
	0x6f, 0x6d, 0x5f, 0x75, 0x6e, 0x69, 0x78, 0x5f, 0x75, 0x74, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0b, 0x66, 0x72, 0x6f, 0x6d, 0x55, 0x6e, 0x69, 0x78, 0x55, 0x74, 0x63, 0x12, 0x1e,
	0x0a, 0x0b, 0x74, 0x6f, 0x5f, 0x75, 0x6e, 0x69, 0x78, 0x5f, 0x75, 0x74, 0x63, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x6f, 0x55, 0x6e, 0x69, 0x78, 0x55, 0x74, 0x63, 0x22, 0x59,
	0x0a, 0x1a, 0x52, 0x65, 0x61, 0x64, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x0f,
	0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x64, 0x62, 0x2e, 0x43, 0x6f, 0x6e, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x61, 0x74, 0x65, 0x52, 0x0e, 0x63, 0x6f, 0x6e, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x61, 0x74, 0x65, 0x22, 0x5a, 0x0a, 0x1b, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3b, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x64, 0x62, 0x2e, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x61, 0x74, 0x65, 0x52, 0x0e, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x61, 0x74, 0x65, 0x22, 0x1e, 0x0a, 0x1c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43,
	0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x31, 0x0a, 0x1b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43,
	0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x22, 0x1e, 0x0a, 0x1c, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x70, 0x70, 0x73, 0x4c, 0x61, 0x62, 0x2d, 0x4b,
	0x45, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x65, 0x76, 0x65, 0x72, 0x79, 0x73,
	0x68, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x73, 0x2f, 0x64, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_rate_proto_rawDescOnce sync.Once
	file_rate_proto_rawDescData = file_rate_proto_rawDesc
)

func file_rate_proto_rawDescGZIP() []byte {
	file_rate_proto_rawDescOnce.Do(func() {
		file_rate_proto_rawDescData = protoimpl.X.CompressGZIP(file_rate_proto_rawDescData)
	})
	return file_rate_proto_rawDescData
}

var file_rate_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_rate_proto_goTypes = []interface{}{
	(*ConversionRate)(nil),               // 0: db.ConversionRate
	(*CreateConversionRateRequest)(nil),  // 1: db.CreateConversionRateRequest
	(*CreateConversionRateResponse)(nil), // 2: db.CreateConversionRateResponse
	(*ReadConversionRateRequest)(nil),    // 3: db.ReadConversionRateRequest
	(*ReadConversionRateResponse)(nil),   // 4: db.ReadConversionRateResponse
	(*UpdateConversionRateRequest)(nil),  // 5: db.UpdateConversionRateRequest
	(*UpdateConversionRateResponse)(nil), // 6: db.UpdateConversionRateResponse
	(*DeleteConversionRateRequest)(nil),  // 7: db.DeleteConversionRateRequest
	(*DeleteConversionRateResponse)(nil), // 8: db.DeleteConversionRateResponse
}
var file_rate_proto_depIdxs = []int32{
	0, // 0: db.CreateConversionRateRequest.conversion_rate:type_name -> db.ConversionRate
	0, // 1: db.CreateConversionRateResponse.conversion_rate:type_name -> db.ConversionRate
	0, // 2: db.ReadConversionRateResponse.conversion_rate:type_name -> db.ConversionRate
	0, // 3: db.UpdateConversionRateRequest.conversion_rate:type_name -> db.ConversionRate
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_rate_proto_init() }
func file_rate_proto_init() {
	if File_rate_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConversionRate); i {
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
		file_rate_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateConversionRateRequest); i {
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
		file_rate_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateConversionRateResponse); i {
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
		file_rate_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadConversionRateRequest); i {
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
		file_rate_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadConversionRateResponse); i {
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
		file_rate_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateConversionRateRequest); i {
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
		file_rate_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateConversionRateResponse); i {
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
		file_rate_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteConversionRateRequest); i {
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
		file_rate_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteConversionRateResponse); i {
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
			RawDescriptor: file_rate_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rate_proto_goTypes,
		DependencyIndexes: file_rate_proto_depIdxs,
		MessageInfos:      file_rate_proto_msgTypes,
	}.Build()
	File_rate_proto = out.File
	file_rate_proto_rawDesc = nil
	file_rate_proto_goTypes = nil
	file_rate_proto_depIdxs = nil
}
