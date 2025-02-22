// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.2
// source: city.proto

package arb_protos

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

type City struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Lat  float32 `protobuf:"fixed32,3,opt,name=lat,proto3" json:"lat,omitempty"`
	Lng  float32 `protobuf:"fixed32,4,opt,name=lng,proto3" json:"lng,omitempty"`
}

func (x *City) Reset() {
	*x = City{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *City) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*City) ProtoMessage() {}

func (x *City) ProtoReflect() protoreflect.Message {
	mi := &file_city_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use City.ProtoReflect.Descriptor instead.
func (*City) Descriptor() ([]byte, []int) {
	return file_city_proto_rawDescGZIP(), []int{0}
}

func (x *City) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *City) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *City) GetLat() float32 {
	if x != nil {
		return x.Lat
	}
	return 0
}

func (x *City) GetLng() float32 {
	if x != nil {
		return x.Lng
	}
	return 0
}

type CityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Lat  float32 `protobuf:"fixed32,2,opt,name=lat,proto3" json:"lat,omitempty"`
	Lng  float32 `protobuf:"fixed32,3,opt,name=lng,proto3" json:"lng,omitempty"`
}

func (x *CityRequest) Reset() {
	*x = CityRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CityRequest) ProtoMessage() {}

func (x *CityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_city_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CityRequest.ProtoReflect.Descriptor instead.
func (*CityRequest) Descriptor() ([]byte, []int) {
	return file_city_proto_rawDescGZIP(), []int{1}
}

func (x *CityRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CityRequest) GetLat() float32 {
	if x != nil {
		return x.Lat
	}
	return 0
}

func (x *CityRequest) GetLng() float32 {
	if x != nil {
		return x.Lng
	}
	return 0
}

type CityResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool  `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Data    *City `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *CityResponse) Reset() {
	*x = CityResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CityResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CityResponse) ProtoMessage() {}

func (x *CityResponse) ProtoReflect() protoreflect.Message {
	mi := &file_city_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CityResponse.ProtoReflect.Descriptor instead.
func (*CityResponse) Descriptor() ([]byte, []int) {
	return file_city_proto_rawDescGZIP(), []int{2}
}

func (x *CityResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *CityResponse) GetData() *City {
	if x != nil {
		return x.Data
	}
	return nil
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_city_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_city_proto_rawDescGZIP(), []int{3}
}

type DeleteCityResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *DeleteCityResponse) Reset() {
	*x = DeleteCityResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCityResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCityResponse) ProtoMessage() {}

func (x *DeleteCityResponse) ProtoReflect() protoreflect.Message {
	mi := &file_city_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCityResponse.ProtoReflect.Descriptor instead.
func (*DeleteCityResponse) Descriptor() ([]byte, []int) {
	return file_city_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteCityResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type GetCitiesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool    `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Data    []*City `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *GetCitiesResponse) Reset() {
	*x = GetCitiesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCitiesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCitiesResponse) ProtoMessage() {}

func (x *GetCitiesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_city_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCitiesResponse.ProtoReflect.Descriptor instead.
func (*GetCitiesResponse) Descriptor() ([]byte, []int) {
	return file_city_proto_rawDescGZIP(), []int{5}
}

func (x *GetCitiesResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *GetCitiesResponse) GetData() []*City {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetCityByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetCityByIdRequest) Reset() {
	*x = GetCityByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCityByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCityByIdRequest) ProtoMessage() {}

func (x *GetCityByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_city_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCityByIdRequest.ProtoReflect.Descriptor instead.
func (*GetCityByIdRequest) Descriptor() ([]byte, []int) {
	return file_city_proto_rawDescGZIP(), []int{6}
}

func (x *GetCityByIdRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetCityByIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool  `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Data    *City `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetCityByIdResponse) Reset() {
	*x = GetCityByIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCityByIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCityByIdResponse) ProtoMessage() {}

func (x *GetCityByIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_city_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCityByIdResponse.ProtoReflect.Descriptor instead.
func (*GetCityByIdResponse) Descriptor() ([]byte, []int) {
	return file_city_proto_rawDescGZIP(), []int{7}
}

func (x *GetCityByIdResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *GetCityByIdResponse) GetData() *City {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetCityByNameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetCityByNameRequest) Reset() {
	*x = GetCityByNameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCityByNameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCityByNameRequest) ProtoMessage() {}

func (x *GetCityByNameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_city_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCityByNameRequest.ProtoReflect.Descriptor instead.
func (*GetCityByNameRequest) Descriptor() ([]byte, []int) {
	return file_city_proto_rawDescGZIP(), []int{8}
}

func (x *GetCityByNameRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetCityByNameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool  `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Data    *City `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetCityByNameResponse) Reset() {
	*x = GetCityByNameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCityByNameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCityByNameResponse) ProtoMessage() {}

func (x *GetCityByNameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_city_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCityByNameResponse.ProtoReflect.Descriptor instead.
func (*GetCityByNameResponse) Descriptor() ([]byte, []int) {
	return file_city_proto_rawDescGZIP(), []int{9}
}

func (x *GetCityByNameResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *GetCityByNameResponse) GetData() *City {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_city_proto protoreflect.FileDescriptor

var file_city_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x63, 0x69,
	0x74, 0x79, 0x22, 0x4e, 0x0a, 0x04, 0x43, 0x69, 0x74, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x6c, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x6c, 0x61, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x6c,
	0x6e, 0x67, 0x22, 0x45, 0x0a, 0x0b, 0x43, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x03, 0x6c, 0x61, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6e, 0x67, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x6c, 0x6e, 0x67, 0x22, 0x48, 0x0a, 0x0c, 0x43, 0x69, 0x74,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x12, 0x1e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0a, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x43, 0x69, 0x74, 0x79, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x2e, 0x0a, 0x12,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x4d, 0x0a, 0x11,
	0x47, 0x65, 0x74, 0x43, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x1e, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x63, 0x69, 0x74, 0x79,
	0x2e, 0x43, 0x69, 0x74, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x24, 0x0a, 0x12, 0x47,
	0x65, 0x74, 0x43, 0x69, 0x74, 0x79, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x4f, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x43, 0x69, 0x74, 0x79, 0x42, 0x79, 0x49, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x12, 0x1e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0a, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x43, 0x69, 0x74, 0x79, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x22, 0x2a, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x43, 0x69, 0x74, 0x79, 0x42, 0x79, 0x4e,
	0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x51,
	0x0a, 0x15, 0x47, 0x65, 0x74, 0x43, 0x69, 0x74, 0x79, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x12, 0x1e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0a, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x43, 0x69, 0x74, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x32, 0xa5, 0x02, 0x0a, 0x0b, 0x43, 0x69, 0x74, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x33, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x69, 0x74, 0x79, 0x12,
	0x11, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x43, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x12, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x43, 0x69, 0x74, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x43, 0x69, 0x74,
	0x79, 0x12, 0x18, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x69, 0x74, 0x79,
	0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x63, 0x69,
	0x74, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x69, 0x74, 0x79, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x43, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x12, 0x0b, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x17, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x69, 0x74, 0x69, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x0a, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x43, 0x69, 0x74, 0x79, 0x12, 0x0a, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x43,
	0x69, 0x74, 0x79, 0x1a, 0x12, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x43, 0x69, 0x74, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x43, 0x69, 0x74, 0x79, 0x12, 0x18, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x47, 0x65, 0x74,
	0x43, 0x69, 0x74, 0x79, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x18, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x69, 0x74,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x47, 0x6f, 0x42, 0x6f, 0x6f, 0x74, 0x43, 0x61,
	0x6d, 0x70, 0x2d, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x31, 0x2f, 0x61, 0x72, 0x62, 0x2d, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_city_proto_rawDescOnce sync.Once
	file_city_proto_rawDescData = file_city_proto_rawDesc
)

func file_city_proto_rawDescGZIP() []byte {
	file_city_proto_rawDescOnce.Do(func() {
		file_city_proto_rawDescData = protoimpl.X.CompressGZIP(file_city_proto_rawDescData)
	})
	return file_city_proto_rawDescData
}

var file_city_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_city_proto_goTypes = []any{
	(*City)(nil),                  // 0: city.City
	(*CityRequest)(nil),           // 1: city.CityRequest
	(*CityResponse)(nil),          // 2: city.CityResponse
	(*Empty)(nil),                 // 3: city.Empty
	(*DeleteCityResponse)(nil),    // 4: city.DeleteCityResponse
	(*GetCitiesResponse)(nil),     // 5: city.GetCitiesResponse
	(*GetCityByIdRequest)(nil),    // 6: city.GetCityByIdRequest
	(*GetCityByIdResponse)(nil),   // 7: city.GetCityByIdResponse
	(*GetCityByNameRequest)(nil),  // 8: city.GetCityByNameRequest
	(*GetCityByNameResponse)(nil), // 9: city.GetCityByNameResponse
}
var file_city_proto_depIdxs = []int32{
	0, // 0: city.CityResponse.data:type_name -> city.City
	0, // 1: city.GetCitiesResponse.data:type_name -> city.City
	0, // 2: city.GetCityByIdResponse.data:type_name -> city.City
	0, // 3: city.GetCityByNameResponse.data:type_name -> city.City
	1, // 4: city.CityService.CreateCity:input_type -> city.CityRequest
	6, // 5: city.CityService.GetCity:input_type -> city.GetCityByIdRequest
	3, // 6: city.CityService.GetCities:input_type -> city.Empty
	0, // 7: city.CityService.UpdateCity:input_type -> city.City
	6, // 8: city.CityService.DeleteCity:input_type -> city.GetCityByIdRequest
	2, // 9: city.CityService.CreateCity:output_type -> city.CityResponse
	7, // 10: city.CityService.GetCity:output_type -> city.GetCityByIdResponse
	5, // 11: city.CityService.GetCities:output_type -> city.GetCitiesResponse
	2, // 12: city.CityService.UpdateCity:output_type -> city.CityResponse
	4, // 13: city.CityService.DeleteCity:output_type -> city.DeleteCityResponse
	9, // [9:14] is the sub-list for method output_type
	4, // [4:9] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_city_proto_init() }
func file_city_proto_init() {
	if File_city_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_city_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*City); i {
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
		file_city_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CityRequest); i {
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
		file_city_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*CityResponse); i {
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
		file_city_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*Empty); i {
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
		file_city_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteCityResponse); i {
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
		file_city_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*GetCitiesResponse); i {
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
		file_city_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*GetCityByIdRequest); i {
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
		file_city_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*GetCityByIdResponse); i {
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
		file_city_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*GetCityByNameRequest); i {
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
		file_city_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*GetCityByNameResponse); i {
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
			RawDescriptor: file_city_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_city_proto_goTypes,
		DependencyIndexes: file_city_proto_depIdxs,
		MessageInfos:      file_city_proto_msgTypes,
	}.Build()
	File_city_proto = out.File
	file_city_proto_rawDesc = nil
	file_city_proto_goTypes = nil
	file_city_proto_depIdxs = nil
}
