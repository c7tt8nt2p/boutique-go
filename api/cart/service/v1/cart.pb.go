// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.3
// source: v1/cart.proto

package v1

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

type AddItemReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *AddItemReq) Reset() {
	*x = AddItemReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_cart_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddItemReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddItemReq) ProtoMessage() {}

func (x *AddItemReq) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use AddItemReq.ProtoReflect.Descriptor instead.
func (*AddItemReq) Descriptor() ([]byte, []int) {
	return file_v1_cart_proto_rawDescGZIP(), []int{0}
}

func (x *AddItemReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type AddItemResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *AddItemResp) Reset() {
	*x = AddItemResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_cart_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddItemResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddItemResp) ProtoMessage() {}

func (x *AddItemResp) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use AddItemResp.ProtoReflect.Descriptor instead.
func (*AddItemResp) Descriptor() ([]byte, []int) {
	return file_v1_cart_proto_rawDescGZIP(), []int{1}
}

func (x *AddItemResp) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type ViewCartReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *ViewCartReq) Reset() {
	*x = ViewCartReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_cart_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ViewCartReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ViewCartReq) ProtoMessage() {}

func (x *ViewCartReq) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ViewCartReq.ProtoReflect.Descriptor instead.
func (*ViewCartReq) Descriptor() ([]byte, []int) {
	return file_v1_cart_proto_rawDescGZIP(), []int{2}
}

func (x *ViewCartReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type ViewCartResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*ViewCartResp_Item `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ViewCartResp) Reset() {
	*x = ViewCartResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_cart_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ViewCartResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ViewCartResp) ProtoMessage() {}

func (x *ViewCartResp) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ViewCartResp.ProtoReflect.Descriptor instead.
func (*ViewCartResp) Descriptor() ([]byte, []int) {
	return file_v1_cart_proto_rawDescGZIP(), []int{3}
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

	ItemId   int64 `protobuf:"varint,1,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	Quantity int64 `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *ViewCartResp_Item) Reset() {
	*x = ViewCartResp_Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_cart_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ViewCartResp_Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ViewCartResp_Item) ProtoMessage() {}

func (x *ViewCartResp_Item) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ViewCartResp_Item.ProtoReflect.Descriptor instead.
func (*ViewCartResp_Item) Descriptor() ([]byte, []int) {
	return file_v1_cart_proto_rawDescGZIP(), []int{3, 0}
}

func (x *ViewCartResp_Item) GetItemId() int64 {
	if x != nil {
		return x.ItemId
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
	0x22, 0x25, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x26, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x49, 0x74,
	0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22,
	0x26, 0x0a, 0x0b, 0x56, 0x69, 0x65, 0x77, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x85, 0x01, 0x0a, 0x0c, 0x56, 0x69, 0x65, 0x77,
	0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x38, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x69, 0x65, 0x77, 0x43, 0x61,
	0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x1a, 0x3b, 0x0a, 0x04, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x74,
	0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x69, 0x74, 0x65,
	0x6d, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x32,
	0x99, 0x01, 0x0a, 0x04, 0x43, 0x61, 0x72, 0x74, 0x12, 0x46, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x49,
	0x74, 0x65, 0x6d, 0x12, 0x1b, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71,
	0x1a, 0x1c, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00,
	0x12, 0x49, 0x0a, 0x08, 0x56, 0x69, 0x65, 0x77, 0x43, 0x61, 0x72, 0x74, 0x12, 0x1c, 0x2e, 0x63,
	0x61, 0x72, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x56,
	0x69, 0x65, 0x77, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1d, 0x2e, 0x63, 0x61, 0x72,
	0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x69, 0x65,
	0x77, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x42, 0x2f, 0x5a, 0x2d, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x78, 0x2d, 0x62, 0x6f, 0x75,
	0x74, 0x69, 0x71, 0x75, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x61, 0x72, 0x74, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
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

var file_v1_cart_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_v1_cart_proto_goTypes = []interface{}{
	(*AddItemReq)(nil),        // 0: cart.service.v1.AddItemReq
	(*AddItemResp)(nil),       // 1: cart.service.v1.AddItemResp
	(*ViewCartReq)(nil),       // 2: cart.service.v1.ViewCartReq
	(*ViewCartResp)(nil),      // 3: cart.service.v1.ViewCartResp
	(*ViewCartResp_Item)(nil), // 4: cart.service.v1.ViewCartResp.Item
}
var file_v1_cart_proto_depIdxs = []int32{
	4, // 0: cart.service.v1.ViewCartResp.items:type_name -> cart.service.v1.ViewCartResp.Item
	0, // 1: cart.service.v1.Cart.AddItem:input_type -> cart.service.v1.AddItemReq
	2, // 2: cart.service.v1.Cart.ViewCart:input_type -> cart.service.v1.ViewCartReq
	1, // 3: cart.service.v1.Cart.AddItem:output_type -> cart.service.v1.AddItemResp
	3, // 4: cart.service.v1.Cart.ViewCart:output_type -> cart.service.v1.ViewCartResp
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
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
			switch v := v.(*AddItemReq); i {
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
			switch v := v.(*AddItemResp); i {
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
			switch v := v.(*ViewCartReq); i {
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
		file_v1_cart_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
			NumMessages:   5,
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
