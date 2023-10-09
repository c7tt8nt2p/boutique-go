// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.3
// source: v1/product.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type CreateProductReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string  `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Stock       int32   `protobuf:"varint,3,opt,name=stock,proto3" json:"stock,omitempty"`
	UnitPrice   float64 `protobuf:"fixed64,4,opt,name=unit_price,json=unitPrice,proto3" json:"unit_price,omitempty"`
}

func (x *CreateProductReq) Reset() {
	*x = CreateProductReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_product_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateProductReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProductReq) ProtoMessage() {}

func (x *CreateProductReq) ProtoReflect() protoreflect.Message {
	mi := &file_v1_product_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProductReq.ProtoReflect.Descriptor instead.
func (*CreateProductReq) Descriptor() ([]byte, []int) {
	return file_v1_product_proto_rawDescGZIP(), []int{0}
}

func (x *CreateProductReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateProductReq) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateProductReq) GetStock() int32 {
	if x != nil {
		return x.Stock
	}
	return 0
}

func (x *CreateProductReq) GetUnitPrice() float64 {
	if x != nil {
		return x.UnitPrice
	}
	return 0
}

type CreateProductResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string  `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Stock       int32   `protobuf:"varint,4,opt,name=stock,proto3" json:"stock,omitempty"`
	UnitPrice   float64 `protobuf:"fixed64,5,opt,name=unit_price,json=unitPrice,proto3" json:"unit_price,omitempty"`
}

func (x *CreateProductResp) Reset() {
	*x = CreateProductResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_product_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateProductResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProductResp) ProtoMessage() {}

func (x *CreateProductResp) ProtoReflect() protoreflect.Message {
	mi := &file_v1_product_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProductResp.ProtoReflect.Descriptor instead.
func (*CreateProductResp) Descriptor() ([]byte, []int) {
	return file_v1_product_proto_rawDescGZIP(), []int{1}
}

func (x *CreateProductResp) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateProductResp) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateProductResp) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateProductResp) GetStock() int32 {
	if x != nil {
		return x.Stock
	}
	return 0
}

func (x *CreateProductResp) GetUnitPrice() float64 {
	if x != nil {
		return x.UnitPrice
	}
	return 0
}

type GetProductByIdReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetProductByIdReq) Reset() {
	*x = GetProductByIdReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_product_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProductByIdReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductByIdReq) ProtoMessage() {}

func (x *GetProductByIdReq) ProtoReflect() protoreflect.Message {
	mi := &file_v1_product_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductByIdReq.ProtoReflect.Descriptor instead.
func (*GetProductByIdReq) Descriptor() ([]byte, []int) {
	return file_v1_product_proto_rawDescGZIP(), []int{2}
}

func (x *GetProductByIdReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetProductByIdResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string  `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Stock       int32   `protobuf:"varint,4,opt,name=stock,proto3" json:"stock,omitempty"`
	UnitPrice   float64 `protobuf:"fixed64,5,opt,name=unit_price,json=unitPrice,proto3" json:"unit_price,omitempty"`
}

func (x *GetProductByIdResp) Reset() {
	*x = GetProductByIdResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_product_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProductByIdResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductByIdResp) ProtoMessage() {}

func (x *GetProductByIdResp) ProtoReflect() protoreflect.Message {
	mi := &file_v1_product_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductByIdResp.ProtoReflect.Descriptor instead.
func (*GetProductByIdResp) Descriptor() ([]byte, []int) {
	return file_v1_product_proto_rawDescGZIP(), []int{3}
}

func (x *GetProductByIdResp) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetProductByIdResp) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetProductByIdResp) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *GetProductByIdResp) GetStock() int32 {
	if x != nil {
		return x.Stock
	}
	return 0
}

func (x *GetProductByIdResp) GetUnitPrice() float64 {
	if x != nil {
		return x.UnitPrice
	}
	return 0
}

type ValidatePurchasableProductReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id  string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Qty int32  `protobuf:"varint,2,opt,name=qty,proto3" json:"qty,omitempty"`
}

func (x *ValidatePurchasableProductReq) Reset() {
	*x = ValidatePurchasableProductReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_product_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidatePurchasableProductReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidatePurchasableProductReq) ProtoMessage() {}

func (x *ValidatePurchasableProductReq) ProtoReflect() protoreflect.Message {
	mi := &file_v1_product_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidatePurchasableProductReq.ProtoReflect.Descriptor instead.
func (*ValidatePurchasableProductReq) Descriptor() ([]byte, []int) {
	return file_v1_product_proto_rawDescGZIP(), []int{4}
}

func (x *ValidatePurchasableProductReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ValidatePurchasableProductReq) GetQty() int32 {
	if x != nil {
		return x.Qty
	}
	return 0
}

type ValidatePurchasableProductResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Validated         bool   `protobuf:"varint,1,opt,name=validated,proto3" json:"validated,omitempty"`
	ValidationMessage string `protobuf:"bytes,2,opt,name=validation_message,json=validationMessage,proto3" json:"validation_message,omitempty"`
}

func (x *ValidatePurchasableProductResp) Reset() {
	*x = ValidatePurchasableProductResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_product_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidatePurchasableProductResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidatePurchasableProductResp) ProtoMessage() {}

func (x *ValidatePurchasableProductResp) ProtoReflect() protoreflect.Message {
	mi := &file_v1_product_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidatePurchasableProductResp.ProtoReflect.Descriptor instead.
func (*ValidatePurchasableProductResp) Descriptor() ([]byte, []int) {
	return file_v1_product_proto_rawDescGZIP(), []int{5}
}

func (x *ValidatePurchasableProductResp) GetValidated() bool {
	if x != nil {
		return x.Validated
	}
	return false
}

func (x *ValidatePurchasableProductResp) GetValidationMessage() string {
	if x != nil {
		return x.ValidationMessage
	}
	return ""
}

var File_v1_product_proto protoreflect.FileDescriptor

var file_v1_product_proto_rawDesc = []byte{
	0x0a, 0x10, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0f, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa8, 0x01, 0x0a,
	0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65,
	0x71, 0x12, 0x1b, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x29,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x05, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x05, 0x73, 0x74, 0x6f,
	0x63, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x1a, 0x02, 0x20,
	0x00, 0x52, 0x05, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x2d, 0x0a, 0x0a, 0x75, 0x6e, 0x69, 0x74,
	0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x42, 0x0e, 0xfa, 0x42,
	0x0b, 0x12, 0x09, 0x21, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x52, 0x09, 0x75, 0x6e,
	0x69, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x22, 0x8e, 0x01, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x6e, 0x69,
	0x74, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x75,
	0x6e, 0x69, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x22, 0x2d, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03,
	0xb0, 0x01, 0x01, 0x52, 0x02, 0x69, 0x64, 0x22, 0x8f, 0x01, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x6e,
	0x69, 0x74, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09,
	0x75, 0x6e, 0x69, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x22, 0x54, 0x0a, 0x1d, 0x56, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x61, 0x62, 0x6c, 0x65,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x03, 0x71, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x1a, 0x02, 0x20, 0x00, 0x52, 0x03, 0x71, 0x74, 0x79, 0x22,
	0x6d, 0x0a, 0x1e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x50, 0x75, 0x72, 0x63, 0x68,
	0x61, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x1c, 0x0a, 0x09, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12,
	0x2d, 0x0a, 0x12, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xc1,
	0x02, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x58, 0x0a, 0x0d, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x21, 0x2e, 0x63, 0x61,
	0x72, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x22,
	0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x22, 0x00, 0x12, 0x5b, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x22, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x23, 0x2e, 0x63, 0x61, 0x72,
	0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x22,
	0x00, 0x12, 0x7f, 0x0a, 0x1a, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x50, 0x75, 0x72,
	0x63, 0x68, 0x61, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12,
	0x2e, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61,
	0x73, 0x61, 0x62, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x1a,
	0x2f, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61,
	0x73, 0x61, 0x62, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x22, 0x00, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6b, 0x78, 0x2d, 0x62, 0x6f, 0x75, 0x74, 0x69, 0x71, 0x75, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_product_proto_rawDescOnce sync.Once
	file_v1_product_proto_rawDescData = file_v1_product_proto_rawDesc
)

func file_v1_product_proto_rawDescGZIP() []byte {
	file_v1_product_proto_rawDescOnce.Do(func() {
		file_v1_product_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_product_proto_rawDescData)
	})
	return file_v1_product_proto_rawDescData
}

var file_v1_product_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_v1_product_proto_goTypes = []interface{}{
	(*CreateProductReq)(nil),               // 0: cart.service.v1.CreateProductReq
	(*CreateProductResp)(nil),              // 1: cart.service.v1.CreateProductResp
	(*GetProductByIdReq)(nil),              // 2: cart.service.v1.GetProductByIdReq
	(*GetProductByIdResp)(nil),             // 3: cart.service.v1.GetProductByIdResp
	(*ValidatePurchasableProductReq)(nil),  // 4: cart.service.v1.ValidatePurchasableProductReq
	(*ValidatePurchasableProductResp)(nil), // 5: cart.service.v1.ValidatePurchasableProductResp
}
var file_v1_product_proto_depIdxs = []int32{
	0, // 0: cart.service.v1.Product.CreateProduct:input_type -> cart.service.v1.CreateProductReq
	2, // 1: cart.service.v1.Product.GetProductById:input_type -> cart.service.v1.GetProductByIdReq
	4, // 2: cart.service.v1.Product.ValidatePurchasableProduct:input_type -> cart.service.v1.ValidatePurchasableProductReq
	1, // 3: cart.service.v1.Product.CreateProduct:output_type -> cart.service.v1.CreateProductResp
	3, // 4: cart.service.v1.Product.GetProductById:output_type -> cart.service.v1.GetProductByIdResp
	5, // 5: cart.service.v1.Product.ValidatePurchasableProduct:output_type -> cart.service.v1.ValidatePurchasableProductResp
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_v1_product_proto_init() }
func file_v1_product_proto_init() {
	if File_v1_product_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_product_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateProductReq); i {
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
		file_v1_product_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateProductResp); i {
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
		file_v1_product_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProductByIdReq); i {
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
		file_v1_product_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProductByIdResp); i {
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
		file_v1_product_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidatePurchasableProductReq); i {
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
		file_v1_product_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidatePurchasableProductResp); i {
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
			RawDescriptor: file_v1_product_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_product_proto_goTypes,
		DependencyIndexes: file_v1_product_proto_depIdxs,
		MessageInfos:      file_v1_product_proto_msgTypes,
	}.Build()
	File_v1_product_proto = out.File
	file_v1_product_proto_rawDesc = nil
	file_v1_product_proto_goTypes = nil
	file_v1_product_proto_depIdxs = nil
}
