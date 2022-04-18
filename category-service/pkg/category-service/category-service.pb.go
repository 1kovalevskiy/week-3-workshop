// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: api/category-service/category-service.proto

package category_service

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type Category struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Category) Reset() {
	*x = Category{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_category_service_category_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Category) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Category) ProtoMessage() {}

func (x *Category) ProtoReflect() protoreflect.Message {
	mi := &file_api_category_service_category_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Category.ProtoReflect.Descriptor instead.
func (*Category) Descriptor() ([]byte, []int) {
	return file_api_category_service_category_service_proto_rawDescGZIP(), []int{0}
}

func (x *Category) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Category) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetCategoryByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetCategoryByIdRequest) Reset() {
	*x = GetCategoryByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_category_service_category_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCategoryByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCategoryByIdRequest) ProtoMessage() {}

func (x *GetCategoryByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_category_service_category_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCategoryByIdRequest.ProtoReflect.Descriptor instead.
func (*GetCategoryByIdRequest) Descriptor() ([]byte, []int) {
	return file_api_category_service_category_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetCategoryByIdRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetCategoryByIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Категория
	Category *Category `protobuf:"bytes,1,opt,name=category,proto3" json:"category,omitempty"`
}

func (x *GetCategoryByIdResponse) Reset() {
	*x = GetCategoryByIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_category_service_category_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCategoryByIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCategoryByIdResponse) ProtoMessage() {}

func (x *GetCategoryByIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_category_service_category_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCategoryByIdResponse.ProtoReflect.Descriptor instead.
func (*GetCategoryByIdResponse) Descriptor() ([]byte, []int) {
	return file_api_category_service_category_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetCategoryByIdResponse) GetCategory() *Category {
	if x != nil {
		return x.Category
	}
	return nil
}

var File_api_category_service_category_service_proto protoreflect.FileDescriptor

var file_api_category_service_category_service_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2d, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2d,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x3b, 0x6f,
	0x7a, 0x6f, 0x6e, 0x6d, 0x70, 0x2e, 0x77, 0x65, 0x65, 0x6b, 0x5f, 0x33, 0x5f, 0x77, 0x6f, 0x72,
	0x6b, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x2e, 0x0a, 0x08, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x31, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x20, 0x00,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x7c, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x61, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x45, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x6d, 0x70, 0x2e, 0x77, 0x65, 0x65, 0x6b, 0x5f,
	0x33, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x32, 0xd0, 0x01, 0x0a, 0x0f, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xbc, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x42, 0x79, 0x49, 0x64, 0x12, 0x53, 0x2e, 0x6f, 0x7a, 0x6f,
	0x6e, 0x6d, 0x70, 0x2e, 0x77, 0x65, 0x65, 0x6b, 0x5f, 0x33, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x73,
	0x68, 0x6f, 0x70, 0x2e, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x54, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x6d, 0x70, 0x2e, 0x77, 0x65, 0x65, 0x6b, 0x5f, 0x33, 0x5f,
	0x77, 0x6f, 0x72, 0x6b, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x27, 0x5a, 0x25, 0x70, 0x6b, 0x67, 0x2f, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x3b, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_category_service_category_service_proto_rawDescOnce sync.Once
	file_api_category_service_category_service_proto_rawDescData = file_api_category_service_category_service_proto_rawDesc
)

func file_api_category_service_category_service_proto_rawDescGZIP() []byte {
	file_api_category_service_category_service_proto_rawDescOnce.Do(func() {
		file_api_category_service_category_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_category_service_category_service_proto_rawDescData)
	})
	return file_api_category_service_category_service_proto_rawDescData
}

var file_api_category_service_category_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_category_service_category_service_proto_goTypes = []interface{}{
	(*Category)(nil),                // 0: ozon-edu-go-2021.week_3_workshop.category_service.category_service.v1.Category
	(*GetCategoryByIdRequest)(nil),  // 1: ozon-edu-go-2021.week_3_workshop.category_service.category_service.v1.GetCategoryByIdRequest
	(*GetCategoryByIdResponse)(nil), // 2: ozon-edu-go-2021.week_3_workshop.category_service.category_service.v1.GetCategoryByIdResponse
}
var file_api_category_service_category_service_proto_depIdxs = []int32{
	0, // 0: ozon-edu-go-2021.week_3_workshop.category_service.category_service.v1.GetCategoryByIdResponse.category:type_name -> ozon-edu-go-2021.week_3_workshop.category_service.category_service.v1.Category
	1, // 1: ozon-edu-go-2021.week_3_workshop.category_service.category_service.v1.CategoryService.GetCategoryById:input_type -> ozon-edu-go-2021.week_3_workshop.category_service.category_service.v1.GetCategoryByIdRequest
	2, // 2: ozon-edu-go-2021.week_3_workshop.category_service.category_service.v1.CategoryService.GetCategoryById:output_type -> ozon-edu-go-2021.week_3_workshop.category_service.category_service.v1.GetCategoryByIdResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_category_service_category_service_proto_init() }
func file_api_category_service_category_service_proto_init() {
	if File_api_category_service_category_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_category_service_category_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Category); i {
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
		file_api_category_service_category_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCategoryByIdRequest); i {
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
		file_api_category_service_category_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCategoryByIdResponse); i {
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
			RawDescriptor: file_api_category_service_category_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_category_service_category_service_proto_goTypes,
		DependencyIndexes: file_api_category_service_category_service_proto_depIdxs,
		MessageInfos:      file_api_category_service_category_service_proto_msgTypes,
	}.Build()
	File_api_category_service_category_service_proto = out.File
	file_api_category_service_category_service_proto_rawDesc = nil
	file_api_category_service_category_service_proto_goTypes = nil
	file_api_category_service_category_service_proto_depIdxs = nil
}
