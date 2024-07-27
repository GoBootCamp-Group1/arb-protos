// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.2
// source: UserService/perm.proto

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

type CreatePermissionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Jwt  string `protobuf:"bytes,3,opt,name=jwt,proto3" json:"jwt,omitempty"`
}

func (x *CreatePermissionRequest) Reset() {
	*x = CreatePermissionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_UserService_perm_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePermissionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePermissionRequest) ProtoMessage() {}

func (x *CreatePermissionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_UserService_perm_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePermissionRequest.ProtoReflect.Descriptor instead.
func (*CreatePermissionRequest) Descriptor() ([]byte, []int) {
	return file_UserService_perm_proto_rawDescGZIP(), []int{0}
}

func (x *CreatePermissionRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreatePermissionRequest) GetJwt() string {
	if x != nil {
		return x.Jwt
	}
	return ""
}

type CreatePermissionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Id      string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Error   string `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *CreatePermissionResponse) Reset() {
	*x = CreatePermissionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_UserService_perm_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePermissionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePermissionResponse) ProtoMessage() {}

func (x *CreatePermissionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_UserService_perm_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePermissionResponse.ProtoReflect.Descriptor instead.
func (*CreatePermissionResponse) Descriptor() ([]byte, []int) {
	return file_UserService_perm_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePermissionResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *CreatePermissionResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreatePermissionResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *CreatePermissionResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type UpdatePermissionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Jwt  string `protobuf:"bytes,4,opt,name=jwt,proto3" json:"jwt,omitempty"`
}

func (x *UpdatePermissionRequest) Reset() {
	*x = UpdatePermissionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_UserService_perm_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePermissionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePermissionRequest) ProtoMessage() {}

func (x *UpdatePermissionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_UserService_perm_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePermissionRequest.ProtoReflect.Descriptor instead.
func (*UpdatePermissionRequest) Descriptor() ([]byte, []int) {
	return file_UserService_perm_proto_rawDescGZIP(), []int{2}
}

func (x *UpdatePermissionRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdatePermissionRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdatePermissionRequest) GetJwt() string {
	if x != nil {
		return x.Jwt
	}
	return ""
}

type UpdatePermissionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Error   string `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *UpdatePermissionResponse) Reset() {
	*x = UpdatePermissionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_UserService_perm_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePermissionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePermissionResponse) ProtoMessage() {}

func (x *UpdatePermissionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_UserService_perm_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePermissionResponse.ProtoReflect.Descriptor instead.
func (*UpdatePermissionResponse) Descriptor() ([]byte, []int) {
	return file_UserService_perm_proto_rawDescGZIP(), []int{3}
}

func (x *UpdatePermissionResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *UpdatePermissionResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *UpdatePermissionResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type GetPermissionByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id  string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Jwt string `protobuf:"bytes,2,opt,name=jwt,proto3" json:"jwt,omitempty"`
}

func (x *GetPermissionByIdRequest) Reset() {
	*x = GetPermissionByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_UserService_perm_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPermissionByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPermissionByIdRequest) ProtoMessage() {}

func (x *GetPermissionByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_UserService_perm_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPermissionByIdRequest.ProtoReflect.Descriptor instead.
func (*GetPermissionByIdRequest) Descriptor() ([]byte, []int) {
	return file_UserService_perm_proto_rawDescGZIP(), []int{4}
}

func (x *GetPermissionByIdRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetPermissionByIdRequest) GetJwt() string {
	if x != nil {
		return x.Jwt
	}
	return ""
}

type GetPermissionByIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool        `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Perm    *Permission `protobuf:"bytes,2,opt,name=perm,proto3" json:"perm,omitempty"`
	Message string      `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Error   string      `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *GetPermissionByIdResponse) Reset() {
	*x = GetPermissionByIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_UserService_perm_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPermissionByIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPermissionByIdResponse) ProtoMessage() {}

func (x *GetPermissionByIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_UserService_perm_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPermissionByIdResponse.ProtoReflect.Descriptor instead.
func (*GetPermissionByIdResponse) Descriptor() ([]byte, []int) {
	return file_UserService_perm_proto_rawDescGZIP(), []int{5}
}

func (x *GetPermissionByIdResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *GetPermissionByIdResponse) GetPerm() *Permission {
	if x != nil {
		return x.Perm
	}
	return nil
}

func (x *GetPermissionByIdResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetPermissionByIdResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type GetPermissionByNameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Jwt  string `protobuf:"bytes,2,opt,name=jwt,proto3" json:"jwt,omitempty"`
}

func (x *GetPermissionByNameRequest) Reset() {
	*x = GetPermissionByNameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_UserService_perm_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPermissionByNameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPermissionByNameRequest) ProtoMessage() {}

func (x *GetPermissionByNameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_UserService_perm_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPermissionByNameRequest.ProtoReflect.Descriptor instead.
func (*GetPermissionByNameRequest) Descriptor() ([]byte, []int) {
	return file_UserService_perm_proto_rawDescGZIP(), []int{6}
}

func (x *GetPermissionByNameRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetPermissionByNameRequest) GetJwt() string {
	if x != nil {
		return x.Jwt
	}
	return ""
}

type GetPermissionByNameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool        `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Perm    *Permission `protobuf:"bytes,2,opt,name=perm,proto3" json:"perm,omitempty"`
	Message string      `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Error   string      `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *GetPermissionByNameResponse) Reset() {
	*x = GetPermissionByNameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_UserService_perm_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPermissionByNameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPermissionByNameResponse) ProtoMessage() {}

func (x *GetPermissionByNameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_UserService_perm_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPermissionByNameResponse.ProtoReflect.Descriptor instead.
func (*GetPermissionByNameResponse) Descriptor() ([]byte, []int) {
	return file_UserService_perm_proto_rawDescGZIP(), []int{7}
}

func (x *GetPermissionByNameResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *GetPermissionByNameResponse) GetPerm() *Permission {
	if x != nil {
		return x.Perm
	}
	return nil
}

func (x *GetPermissionByNameResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetPermissionByNameResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type DeletePermissionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id  string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Jwt string `protobuf:"bytes,2,opt,name=jwt,proto3" json:"jwt,omitempty"`
}

func (x *DeletePermissionRequest) Reset() {
	*x = DeletePermissionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_UserService_perm_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePermissionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePermissionRequest) ProtoMessage() {}

func (x *DeletePermissionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_UserService_perm_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePermissionRequest.ProtoReflect.Descriptor instead.
func (*DeletePermissionRequest) Descriptor() ([]byte, []int) {
	return file_UserService_perm_proto_rawDescGZIP(), []int{8}
}

func (x *DeletePermissionRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DeletePermissionRequest) GetJwt() string {
	if x != nil {
		return x.Jwt
	}
	return ""
}

type DeletePermissionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Error   string `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *DeletePermissionResponse) Reset() {
	*x = DeletePermissionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_UserService_perm_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePermissionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePermissionResponse) ProtoMessage() {}

func (x *DeletePermissionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_UserService_perm_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePermissionResponse.ProtoReflect.Descriptor instead.
func (*DeletePermissionResponse) Descriptor() ([]byte, []int) {
	return file_UserService_perm_proto_rawDescGZIP(), []int{9}
}

func (x *DeletePermissionResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *DeletePermissionResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *DeletePermissionResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type GetAllPermissionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Jwt string `protobuf:"bytes,1,opt,name=jwt,proto3" json:"jwt,omitempty"`
}

func (x *GetAllPermissionsRequest) Reset() {
	*x = GetAllPermissionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_UserService_perm_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllPermissionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllPermissionsRequest) ProtoMessage() {}

func (x *GetAllPermissionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_UserService_perm_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllPermissionsRequest.ProtoReflect.Descriptor instead.
func (*GetAllPermissionsRequest) Descriptor() ([]byte, []int) {
	return file_UserService_perm_proto_rawDescGZIP(), []int{10}
}

func (x *GetAllPermissionsRequest) GetJwt() string {
	if x != nil {
		return x.Jwt
	}
	return ""
}

type GetAllPermissionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool          `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Perms   []*Permission `protobuf:"bytes,2,rep,name=perms,proto3" json:"perms,omitempty"`
	Message string        `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Error   string        `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *GetAllPermissionsResponse) Reset() {
	*x = GetAllPermissionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_UserService_perm_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllPermissionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllPermissionsResponse) ProtoMessage() {}

func (x *GetAllPermissionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_UserService_perm_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllPermissionsResponse.ProtoReflect.Descriptor instead.
func (*GetAllPermissionsResponse) Descriptor() ([]byte, []int) {
	return file_UserService_perm_proto_rawDescGZIP(), []int{11}
}

func (x *GetAllPermissionsResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *GetAllPermissionsResponse) GetPerms() []*Permission {
	if x != nil {
		return x.Perms
	}
	return nil
}

func (x *GetAllPermissionsResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetAllPermissionsResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_UserService_perm_proto protoreflect.FileDescriptor

var file_UserService_perm_proto_rawDesc = []byte{
	0x0a, 0x16, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x65,
	0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x1a, 0x16,
	0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3f, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6a, 0x77, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6a, 0x77, 0x74, 0x22, 0x74, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x4f, 0x0a,
	0x17, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x6a, 0x77, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6a, 0x77, 0x74, 0x22, 0x64,
	0x0a, 0x18, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x22, 0x3c, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x10, 0x0a, 0x03, 0x6a, 0x77, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6a,
	0x77, 0x74, 0x22, 0x8b, 0x01, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x24, 0x0a, 0x04, 0x70, 0x65,
	0x72, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x04, 0x70, 0x65, 0x72, 0x6d,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x22, 0x42, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6a, 0x77, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6a, 0x77, 0x74, 0x22, 0x8d, 0x01, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x24,
	0x0a, 0x04, 0x70, 0x65, 0x72, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x04,
	0x70, 0x65, 0x72, 0x6d, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x22, 0x3b, 0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x10, 0x0a, 0x03, 0x6a, 0x77, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6a, 0x77,
	0x74, 0x22, 0x64, 0x0a, 0x18, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x2c, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6a, 0x77, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6a, 0x77, 0x74, 0x22, 0x8d, 0x01, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x26, 0x0a,
	0x05, 0x70, 0x65, 0x72, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x05,
	0x70, 0x65, 0x72, 0x6d, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0xa0, 0x04, 0x0a, 0x11, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x53, 0x0a, 0x10, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x1d, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e,
	0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x53, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x56, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x49, 0x64, 0x12, 0x1e, 0x2e, 0x72, 0x6f, 0x6c,
	0x65, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x42,
	0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x72, 0x6f, 0x6c,
	0x65, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x42,
	0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5c, 0x0a,
	0x13, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x79,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x50,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x47, 0x65,
	0x74, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x4e, 0x61, 0x6d,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x53, 0x0a, 0x10, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x1d, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e,
	0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x56, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1e, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x47, 0x6f, 0x42, 0x6f, 0x6f, 0x74, 0x43, 0x61, 0x6d,
	0x70, 0x2d, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x31, 0x2f, 0x61, 0x72, 0x62, 0x2d, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_UserService_perm_proto_rawDescOnce sync.Once
	file_UserService_perm_proto_rawDescData = file_UserService_perm_proto_rawDesc
)

func file_UserService_perm_proto_rawDescGZIP() []byte {
	file_UserService_perm_proto_rawDescOnce.Do(func() {
		file_UserService_perm_proto_rawDescData = protoimpl.X.CompressGZIP(file_UserService_perm_proto_rawDescData)
	})
	return file_UserService_perm_proto_rawDescData
}

var file_UserService_perm_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_UserService_perm_proto_goTypes = []any{
	(*CreatePermissionRequest)(nil),     // 0: role.CreatePermissionRequest
	(*CreatePermissionResponse)(nil),    // 1: role.CreatePermissionResponse
	(*UpdatePermissionRequest)(nil),     // 2: role.UpdatePermissionRequest
	(*UpdatePermissionResponse)(nil),    // 3: role.UpdatePermissionResponse
	(*GetPermissionByIdRequest)(nil),    // 4: role.GetPermissionByIdRequest
	(*GetPermissionByIdResponse)(nil),   // 5: role.GetPermissionByIdResponse
	(*GetPermissionByNameRequest)(nil),  // 6: role.GetPermissionByNameRequest
	(*GetPermissionByNameResponse)(nil), // 7: role.GetPermissionByNameResponse
	(*DeletePermissionRequest)(nil),     // 8: role.DeletePermissionRequest
	(*DeletePermissionResponse)(nil),    // 9: role.DeletePermissionResponse
	(*GetAllPermissionsRequest)(nil),    // 10: role.GetAllPermissionsRequest
	(*GetAllPermissionsResponse)(nil),   // 11: role.GetAllPermissionsResponse
	(*Permission)(nil),                  // 12: user.Permission
}
var file_UserService_perm_proto_depIdxs = []int32{
	12, // 0: role.GetPermissionByIdResponse.perm:type_name -> user.Permission
	12, // 1: role.GetPermissionByNameResponse.perm:type_name -> user.Permission
	12, // 2: role.GetAllPermissionsResponse.perms:type_name -> user.Permission
	0,  // 3: role.PermissionService.CreatePermission:input_type -> role.CreatePermissionRequest
	2,  // 4: role.PermissionService.UpdatePermission:input_type -> role.UpdatePermissionRequest
	4,  // 5: role.PermissionService.GetPermissionById:input_type -> role.GetPermissionByIdRequest
	6,  // 6: role.PermissionService.GetPermissionByName:input_type -> role.GetPermissionByNameRequest
	8,  // 7: role.PermissionService.DeletePermission:input_type -> role.DeletePermissionRequest
	10, // 8: role.PermissionService.GetAllPermissions:input_type -> role.GetAllPermissionsRequest
	1,  // 9: role.PermissionService.CreatePermission:output_type -> role.CreatePermissionResponse
	3,  // 10: role.PermissionService.UpdatePermission:output_type -> role.UpdatePermissionResponse
	5,  // 11: role.PermissionService.GetPermissionById:output_type -> role.GetPermissionByIdResponse
	7,  // 12: role.PermissionService.GetPermissionByName:output_type -> role.GetPermissionByNameResponse
	9,  // 13: role.PermissionService.DeletePermission:output_type -> role.DeletePermissionResponse
	11, // 14: role.PermissionService.GetAllPermissions:output_type -> role.GetAllPermissionsResponse
	9,  // [9:15] is the sub-list for method output_type
	3,  // [3:9] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_UserService_perm_proto_init() }
func file_UserService_perm_proto_init() {
	if File_UserService_perm_proto != nil {
		return
	}
	file_UserService_user_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_UserService_perm_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CreatePermissionRequest); i {
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
		file_UserService_perm_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CreatePermissionResponse); i {
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
		file_UserService_perm_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*UpdatePermissionRequest); i {
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
		file_UserService_perm_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*UpdatePermissionResponse); i {
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
		file_UserService_perm_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetPermissionByIdRequest); i {
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
		file_UserService_perm_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*GetPermissionByIdResponse); i {
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
		file_UserService_perm_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*GetPermissionByNameRequest); i {
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
		file_UserService_perm_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*GetPermissionByNameResponse); i {
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
		file_UserService_perm_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*DeletePermissionRequest); i {
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
		file_UserService_perm_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*DeletePermissionResponse); i {
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
		file_UserService_perm_proto_msgTypes[10].Exporter = func(v any, i int) any {
			switch v := v.(*GetAllPermissionsRequest); i {
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
		file_UserService_perm_proto_msgTypes[11].Exporter = func(v any, i int) any {
			switch v := v.(*GetAllPermissionsResponse); i {
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
			RawDescriptor: file_UserService_perm_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_UserService_perm_proto_goTypes,
		DependencyIndexes: file_UserService_perm_proto_depIdxs,
		MessageInfos:      file_UserService_perm_proto_msgTypes,
	}.Build()
	File_UserService_perm_proto = out.File
	file_UserService_perm_proto_rawDesc = nil
	file_UserService_perm_proto_goTypes = nil
	file_UserService_perm_proto_depIdxs = nil
}
