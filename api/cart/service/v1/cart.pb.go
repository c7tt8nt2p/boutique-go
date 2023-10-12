// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.3
// source: v1/cart.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type NewCartReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *NewCartReq) Reset() {
	*x = NewCartReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_cart_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewCartReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewCartReq) ProtoMessage() {}

func (x *NewCartReq) ProtoReflect() protoreflect.Message {
	mi := &file_v1_cart_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewCartReq.ProtoReflect.Descriptor instead.
func (*NewCartReq) Descriptor() ([]byte, []int) {
	return file_v1_cart_proto_rawDescGZIP(), []int{0}
}

func (x *NewCartReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type NewCartResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *NewCartResp) Reset() {
	*x = NewCartResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_cart_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewCartResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewCartResp) ProtoMessage() {}

func (x *NewCartResp) ProtoReflect() protoreflect.Message {
	mi := &file_v1_cart_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewCartResp.ProtoReflect.Descriptor instead.
func (*NewCartResp) Descriptor() ([]byte, []int) {
	return file_v1_cart_proto_rawDescGZIP(), []int{1}
}

func (x *NewCartResp) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type AddItemToCartReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId string `protobuf:"bytes,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Qty       int32  `protobuf:"varint,2,opt,name=qty,proto3" json:"qty,omitempty"`
}

func (x *AddItemToCartReq) Reset() {
	*x = AddItemToCartReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_cart_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddItemToCartReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddItemToCartReq) ProtoMessage() {}

func (x *AddItemToCartReq) ProtoReflect() protoreflect.Message {
	mi := &file_v1_cart_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddItemToCartReq.ProtoReflect.Descriptor instead.
func (*AddItemToCartReq) Descriptor() ([]byte, []int) {
	return file_v1_cart_proto_rawDescGZIP(), []int{2}
}

func (x *AddItemToCartReq) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *AddItemToCartReq) GetQty() int32 {
	if x != nil {
		return x.Qty
	}
	return 0
}

type AddItemToCartResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId string `protobuf:"bytes,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Qty       int32  `protobuf:"varint,2,opt,name=qty,proto3" json:"qty,omitempty"`
}

func (x *AddItemToCartResp) Reset() {
	*x = AddItemToCartResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_cart_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddItemToCartResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddItemToCartResp) ProtoMessage() {}

func (x *AddItemToCartResp) ProtoReflect() protoreflect.Message {
	mi := &file_v1_cart_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddItemToCartResp.ProtoReflect.Descriptor instead.
func (*AddItemToCartResp) Descriptor() ([]byte, []int) {
	return file_v1_cart_proto_rawDescGZIP(), []int{3}
}

func (x *AddItemToCartResp) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *AddItemToCartResp) GetQty() int32 {
	if x != nil {
		return x.Qty
	}
	return 0
}

type ViewCartResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CartId string               `protobuf:"bytes,1,opt,name=cart_id,json=cartId,proto3" json:"cart_id,omitempty"`
	Items  []*ViewCartResp_Item `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ViewCartResp) Reset() {
	*x = ViewCartResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_cart_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ViewCartResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ViewCartResp) ProtoMessage() {}

func (x *ViewCartResp) ProtoReflect() protoreflect.Message {
	mi := &file_v1_cart_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ViewCartResp.ProtoReflect.Descriptor instead.
func (*ViewCartResp) Descriptor() ([]byte, []int) {
	return file_v1_cart_proto_rawDescGZIP(), []int{4}
}

func (x *ViewCartResp) GetCartId() string {
	if x != nil {
		return x.CartId
	}
	return ""
}

func (x *ViewCartResp) GetItems() []*ViewCartResp_Item {
	if x != nil {
		return x.Items
	}
	return nil
}

type ViewCartResp_Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Price    float64 `protobuf:"fixed64,3,opt,name=price,proto3" json:"price,omitempty"`
	Quantity int64   `protobuf:"varint,4,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *ViewCartResp_Item) Reset() {
	*x = ViewCartResp_Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_cart_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ViewCartResp_Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ViewCartResp_Item) ProtoMessage() {}

func (x *ViewCartResp_Item) ProtoReflect() protoreflect.Message {
	mi := &file_v1_cart_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ViewCartResp_Item.ProtoReflect.Descriptor instead.
func (*ViewCartResp_Item) Descriptor() ([]byte, []int) {
	return file_v1_cart_proto_rawDescGZIP(), []int{4, 0}
}

func (x *ViewCartResp_Item) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ViewCartResp_Item) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ViewCartResp_Item) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *ViewCartResp_Item) GetQuantity() int64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

var File_v1_cart_proto protoreflect.FileDescriptor

var file_v1_cart_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0f, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2f, 0x0a, 0x0a, 0x4e, 0x65, 0x77, 0x43, 0x61, 0x72,
	0x74, 0x52, 0x65, 0x71, 0x12, 0x21, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x1d, 0x0a, 0x0b, 0x4e, 0x65, 0x77, 0x43, 0x61,
	0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x56, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x49, 0x74, 0x65,
	0x6d, 0x54, 0x6f, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x12, 0x27, 0x0a, 0x0a, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08,
	0xfa, 0x42, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x03, 0x71, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x42, 0x07, 0xfa, 0x42, 0x04, 0x1a, 0x02, 0x20, 0x00, 0x52, 0x03, 0x71, 0x74, 0x79, 0x22, 0x44,
	0x0a, 0x11, 0x41, 0x64, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x54, 0x6f, 0x43, 0x61, 0x72, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x71, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x03, 0x71, 0x74, 0x79, 0x22, 0xbf, 0x01, 0x0a, 0x0c, 0x56, 0x69, 0x65, 0x77, 0x43, 0x61, 0x72,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x61, 0x72, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x61, 0x72, 0x74, 0x49, 0x64, 0x12, 0x38,
	0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e,
	0x63, 0x61, 0x72, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x56, 0x69, 0x65, 0x77, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x2e, 0x49, 0x74, 0x65,
	0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x1a, 0x5c, 0x0a, 0x04, 0x49, 0x74, 0x65, 0x6d,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75,
	0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x71, 0x75,
	0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x32, 0xed, 0x01, 0x0a, 0x04, 0x43, 0x61, 0x72, 0x74, 0x12,
	0x46, 0x0a, 0x07, 0x4e, 0x65, 0x77, 0x43, 0x61, 0x72, 0x74, 0x12, 0x1b, 0x2e, 0x63, 0x61, 0x72,
	0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x65, 0x77,
	0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1c, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x65, 0x77, 0x43, 0x61, 0x72,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x58, 0x0a, 0x0d, 0x41, 0x64, 0x64, 0x49, 0x74,
	0x65, 0x6d, 0x54, 0x6f, 0x43, 0x61, 0x72, 0x74, 0x12, 0x21, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x49, 0x74,
	0x65, 0x6d, 0x54, 0x6f, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x22, 0x2e, 0x63, 0x61,
	0x72, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64,
	0x64, 0x49, 0x74, 0x65, 0x6d, 0x54, 0x6f, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22,
	0x00, 0x12, 0x43, 0x0a, 0x08, 0x56, 0x69, 0x65, 0x77, 0x43, 0x61, 0x72, 0x74, 0x12, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1d, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x69, 0x65, 0x77, 0x43, 0x61, 0x72, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x78, 0x2d, 0x62, 0x6f, 0x75, 0x74, 0x69, 0x71, 0x75, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x61, 0x72, 0x74, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_cart_proto_rawDescOnce sync.Once
	file_v1_cart_proto_rawDescData = file_v1_cart_proto_rawDesc
)

func file_v1_cart_proto_rawDescGZIP() []byte {
	file_v1_cart_proto_rawDescOnce.Do(func() {
		file_v1_cart_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_cart_proto_rawDescData)
	})
	return file_v1_cart_proto_rawDescData
}

var file_v1_cart_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_v1_cart_proto_goTypes = []interface{}{
	(*NewCartReq)(nil),        // 0: cart.service.v1.NewCartReq
	(*NewCartResp)(nil),       // 1: cart.service.v1.NewCartResp
	(*AddItemToCartReq)(nil),  // 2: cart.service.v1.AddItemToCartReq
	(*AddItemToCartResp)(nil), // 3: cart.service.v1.AddItemToCartResp
	(*ViewCartResp)(nil),      // 4: cart.service.v1.ViewCartResp
	(*ViewCartResp_Item)(nil), // 5: cart.service.v1.ViewCartResp.Item
	(*emptypb.Empty)(nil),     // 6: google.protobuf.Empty
}
var file_v1_cart_proto_depIdxs = []int32{
	5, // 0: cart.service.v1.ViewCartResp.items:type_name -> cart.service.v1.ViewCartResp.Item
	0, // 1: cart.service.v1.Cart.NewCart:input_type -> cart.service.v1.NewCartReq
	2, // 2: cart.service.v1.Cart.AddItemToCart:input_type -> cart.service.v1.AddItemToCartReq
	6, // 3: cart.service.v1.Cart.ViewCart:input_type -> google.protobuf.Empty
	1, // 4: cart.service.v1.Cart.NewCart:output_type -> cart.service.v1.NewCartResp
	3, // 5: cart.service.v1.Cart.AddItemToCart:output_type -> cart.service.v1.AddItemToCartResp
	4, // 6: cart.service.v1.Cart.ViewCart:output_type -> cart.service.v1.ViewCartResp
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_v1_cart_proto_init() }
func file_v1_cart_proto_init() {
	if File_v1_cart_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_cart_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewCartReq); i {
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
		file_v1_cart_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewCartResp); i {
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
		file_v1_cart_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddItemToCartReq); i {
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
		file_v1_cart_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddItemToCartResp); i {
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
		file_v1_cart_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ViewCartResp); i {
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
		file_v1_cart_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ViewCartResp_Item); i {
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
			RawDescriptor: file_v1_cart_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_cart_proto_goTypes,
		DependencyIndexes: file_v1_cart_proto_depIdxs,
		MessageInfos:      file_v1_cart_proto_msgTypes,
	}.Build()
	File_v1_cart_proto = out.File
	file_v1_cart_proto_rawDesc = nil
	file_v1_cart_proto_goTypes = nil
	file_v1_cart_proto_depIdxs = nil
}
