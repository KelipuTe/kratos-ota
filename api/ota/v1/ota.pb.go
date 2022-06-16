// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.1
// source: ota/v1/ota.proto

package v1

import (
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

type GetTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId     string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	ClientSecret string `protobuf:"bytes,2,opt,name=client_secret,json=clientSecret,proto3" json:"client_secret,omitempty"`
}

func (x *GetTokenRequest) Reset() {
	*x = GetTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ota_v1_ota_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTokenRequest) ProtoMessage() {}

func (x *GetTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ota_v1_ota_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTokenRequest.ProtoReflect.Descriptor instead.
func (*GetTokenRequest) Descriptor() ([]byte, []int) {
	return file_ota_v1_ota_proto_rawDescGZIP(), []int{0}
}

func (x *GetTokenRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *GetTokenRequest) GetClientSecret() string {
	if x != nil {
		return x.ClientSecret
	}
	return ""
}

type GetTokenReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	Token    string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *GetTokenReply) Reset() {
	*x = GetTokenReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ota_v1_ota_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTokenReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTokenReply) ProtoMessage() {}

func (x *GetTokenReply) ProtoReflect() protoreflect.Message {
	mi := &file_ota_v1_ota_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTokenReply.ProtoReflect.Descriptor instead.
func (*GetTokenReply) Descriptor() ([]byte, []int) {
	return file_ota_v1_ota_proto_rawDescGZIP(), []int{1}
}

func (x *GetTokenReply) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *GetTokenReply) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type RoomType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomTypeId   string `protobuf:"bytes,1,opt,name=room_type_id,json=roomTypeId,proto3" json:"room_type_id,omitempty"`
	RoomTypeName string `protobuf:"bytes,2,opt,name=room_type_name,json=roomTypeName,proto3" json:"room_type_name,omitempty"`
}

func (x *RoomType) Reset() {
	*x = RoomType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ota_v1_ota_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomType) ProtoMessage() {}

func (x *RoomType) ProtoReflect() protoreflect.Message {
	mi := &file_ota_v1_ota_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomType.ProtoReflect.Descriptor instead.
func (*RoomType) Descriptor() ([]byte, []int) {
	return file_ota_v1_ota_proto_rawDescGZIP(), []int{2}
}

func (x *RoomType) GetRoomTypeId() string {
	if x != nil {
		return x.RoomTypeId
	}
	return ""
}

func (x *RoomType) GetRoomTypeName() string {
	if x != nil {
		return x.RoomTypeName
	}
	return ""
}

type GetHotelRoomTypeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ota        string `protobuf:"bytes,1,opt,name=ota,proto3" json:"ota,omitempty"`
	HotelId    string `protobuf:"bytes,2,opt,name=hotel_id,json=hotelId,proto3" json:"hotel_id,omitempty"`
	RoomTypeId string `protobuf:"bytes,3,opt,name=room_type_id,json=roomTypeId,proto3" json:"room_type_id,omitempty"`
}

func (x *GetHotelRoomTypeRequest) Reset() {
	*x = GetHotelRoomTypeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ota_v1_ota_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHotelRoomTypeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHotelRoomTypeRequest) ProtoMessage() {}

func (x *GetHotelRoomTypeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ota_v1_ota_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHotelRoomTypeRequest.ProtoReflect.Descriptor instead.
func (*GetHotelRoomTypeRequest) Descriptor() ([]byte, []int) {
	return file_ota_v1_ota_proto_rawDescGZIP(), []int{3}
}

func (x *GetHotelRoomTypeRequest) GetOta() string {
	if x != nil {
		return x.Ota
	}
	return ""
}

func (x *GetHotelRoomTypeRequest) GetHotelId() string {
	if x != nil {
		return x.HotelId
	}
	return ""
}

func (x *GetHotelRoomTypeRequest) GetRoomTypeId() string {
	if x != nil {
		return x.RoomTypeId
	}
	return ""
}

type GetHotelRoomTypeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ota      string    `protobuf:"bytes,1,opt,name=ota,proto3" json:"ota,omitempty"`
	HotelId  string    `protobuf:"bytes,2,opt,name=hotel_id,json=hotelId,proto3" json:"hotel_id,omitempty"`
	RoomType *RoomType `protobuf:"bytes,3,opt,name=room_type,json=roomType,proto3" json:"room_type,omitempty"`
}

func (x *GetHotelRoomTypeReply) Reset() {
	*x = GetHotelRoomTypeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ota_v1_ota_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHotelRoomTypeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHotelRoomTypeReply) ProtoMessage() {}

func (x *GetHotelRoomTypeReply) ProtoReflect() protoreflect.Message {
	mi := &file_ota_v1_ota_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHotelRoomTypeReply.ProtoReflect.Descriptor instead.
func (*GetHotelRoomTypeReply) Descriptor() ([]byte, []int) {
	return file_ota_v1_ota_proto_rawDescGZIP(), []int{4}
}

func (x *GetHotelRoomTypeReply) GetOta() string {
	if x != nil {
		return x.Ota
	}
	return ""
}

func (x *GetHotelRoomTypeReply) GetHotelId() string {
	if x != nil {
		return x.HotelId
	}
	return ""
}

func (x *GetHotelRoomTypeReply) GetRoomType() *RoomType {
	if x != nil {
		return x.RoomType
	}
	return nil
}

type ListHotelRoomTypeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ota     string `protobuf:"bytes,1,opt,name=ota,proto3" json:"ota,omitempty"`
	HotelId string `protobuf:"bytes,2,opt,name=hotel_id,json=hotelId,proto3" json:"hotel_id,omitempty"`
}

func (x *ListHotelRoomTypeRequest) Reset() {
	*x = ListHotelRoomTypeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ota_v1_ota_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListHotelRoomTypeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListHotelRoomTypeRequest) ProtoMessage() {}

func (x *ListHotelRoomTypeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ota_v1_ota_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListHotelRoomTypeRequest.ProtoReflect.Descriptor instead.
func (*ListHotelRoomTypeRequest) Descriptor() ([]byte, []int) {
	return file_ota_v1_ota_proto_rawDescGZIP(), []int{5}
}

func (x *ListHotelRoomTypeRequest) GetOta() string {
	if x != nil {
		return x.Ota
	}
	return ""
}

func (x *ListHotelRoomTypeRequest) GetHotelId() string {
	if x != nil {
		return x.HotelId
	}
	return ""
}

type ListHotelRoomTypeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ota          string      `protobuf:"bytes,1,opt,name=ota,proto3" json:"ota,omitempty"`
	HotelId      string      `protobuf:"bytes,2,opt,name=hotel_id,json=hotelId,proto3" json:"hotel_id,omitempty"`
	RoomTypeList []*RoomType `protobuf:"bytes,3,rep,name=room_type_list,json=roomTypeList,proto3" json:"room_type_list,omitempty"`
}

func (x *ListHotelRoomTypeReply) Reset() {
	*x = ListHotelRoomTypeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ota_v1_ota_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListHotelRoomTypeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListHotelRoomTypeReply) ProtoMessage() {}

func (x *ListHotelRoomTypeReply) ProtoReflect() protoreflect.Message {
	mi := &file_ota_v1_ota_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListHotelRoomTypeReply.ProtoReflect.Descriptor instead.
func (*ListHotelRoomTypeReply) Descriptor() ([]byte, []int) {
	return file_ota_v1_ota_proto_rawDescGZIP(), []int{6}
}

func (x *ListHotelRoomTypeReply) GetOta() string {
	if x != nil {
		return x.Ota
	}
	return ""
}

func (x *ListHotelRoomTypeReply) GetHotelId() string {
	if x != nil {
		return x.HotelId
	}
	return ""
}

func (x *ListHotelRoomTypeReply) GetRoomTypeList() []*RoomType {
	if x != nil {
		return x.RoomTypeList
	}
	return nil
}

var File_ota_v1_ota_proto protoreflect.FileDescriptor

var file_ota_v1_ota_proto_rawDesc = []byte{
	0x0a, 0x10, 0x6f, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x06, 0x6f, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x53, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x22, 0x42, 0x0a,
	0x0d, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1b,
	0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x22, 0x52, 0x0a, 0x08, 0x52, 0x6f, 0x6f, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a,
	0x0c, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x6f, 0x6f, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x49, 0x64, 0x12,
	0x24, 0x0a, 0x0e, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x6f, 0x6f, 0x6d, 0x54, 0x79, 0x70,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x68, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x74, 0x65,
	0x6c, 0x52, 0x6f, 0x6f, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x6f, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6f,
	0x74, 0x61, 0x12, 0x19, 0x0a, 0x08, 0x68, 0x6f, 0x74, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x68, 0x6f, 0x74, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x20, 0x0a,
	0x0c, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x6f, 0x6f, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x49, 0x64, 0x22,
	0x73, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x52, 0x6f, 0x6f, 0x6d, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6f, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6f, 0x74, 0x61, 0x12, 0x19, 0x0a, 0x08, 0x68, 0x6f,
	0x74, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x68, 0x6f,
	0x74, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x2d, 0x0a, 0x09, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6f, 0x74, 0x61, 0x2e, 0x76,
	0x31, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x72, 0x6f, 0x6f, 0x6d,
	0x54, 0x79, 0x70, 0x65, 0x22, 0x47, 0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x6f, 0x74, 0x65,
	0x6c, 0x52, 0x6f, 0x6f, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x6f, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6f,
	0x74, 0x61, 0x12, 0x19, 0x0a, 0x08, 0x68, 0x6f, 0x74, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x68, 0x6f, 0x74, 0x65, 0x6c, 0x49, 0x64, 0x22, 0x7d, 0x0a,
	0x16, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x52, 0x6f, 0x6f, 0x6d, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6f, 0x74, 0x61, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6f, 0x74, 0x61, 0x12, 0x19, 0x0a, 0x08, 0x68, 0x6f, 0x74,
	0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x68, 0x6f, 0x74,
	0x65, 0x6c, 0x49, 0x64, 0x12, 0x36, 0x0a, 0x0e, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6f,
	0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0c,
	0x72, 0x6f, 0x6f, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x32, 0xdd, 0x02, 0x0a,
	0x03, 0x4f, 0x74, 0x61, 0x12, 0x58, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x17, 0x2e, 0x6f, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x6f, 0x74, 0x61, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x22, 0x11, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x67, 0x65, 0x74, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x3a, 0x01, 0x2a, 0x12, 0x7b,
	0x0a, 0x10, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x52, 0x6f, 0x6f, 0x6d, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x1f, 0x2e, 0x6f, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x48,
	0x6f, 0x74, 0x65, 0x6c, 0x52, 0x6f, 0x6f, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x6f, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x48, 0x6f, 0x74, 0x65, 0x6c, 0x52, 0x6f, 0x6f, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x27, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21, 0x12, 0x1f, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x74, 0x61, 0x2f, 0x67, 0x65, 0x74, 0x5f, 0x68, 0x6f, 0x74, 0x65,
	0x6c, 0x5f, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x12, 0x7f, 0x0a, 0x11, 0x4c,
	0x69, 0x73, 0x74, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x52, 0x6f, 0x6f, 0x6d, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x20, 0x2e, 0x6f, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x6f,
	0x74, 0x65, 0x6c, 0x52, 0x6f, 0x6f, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x6f, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x48, 0x6f, 0x74, 0x65, 0x6c, 0x52, 0x6f, 0x6f, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x28, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22, 0x12, 0x20, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x74, 0x61, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x68, 0x6f, 0x74,
	0x65, 0x6c, 0x5f, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x42, 0x13, 0x5a, 0x11,
	0x6f, 0x74, 0x61, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x3b, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ota_v1_ota_proto_rawDescOnce sync.Once
	file_ota_v1_ota_proto_rawDescData = file_ota_v1_ota_proto_rawDesc
)

func file_ota_v1_ota_proto_rawDescGZIP() []byte {
	file_ota_v1_ota_proto_rawDescOnce.Do(func() {
		file_ota_v1_ota_proto_rawDescData = protoimpl.X.CompressGZIP(file_ota_v1_ota_proto_rawDescData)
	})
	return file_ota_v1_ota_proto_rawDescData
}

var file_ota_v1_ota_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_ota_v1_ota_proto_goTypes = []interface{}{
	(*GetTokenRequest)(nil),          // 0: ota.v1.GetTokenRequest
	(*GetTokenReply)(nil),            // 1: ota.v1.GetTokenReply
	(*RoomType)(nil),                 // 2: ota.v1.RoomType
	(*GetHotelRoomTypeRequest)(nil),  // 3: ota.v1.GetHotelRoomTypeRequest
	(*GetHotelRoomTypeReply)(nil),    // 4: ota.v1.GetHotelRoomTypeReply
	(*ListHotelRoomTypeRequest)(nil), // 5: ota.v1.ListHotelRoomTypeRequest
	(*ListHotelRoomTypeReply)(nil),   // 6: ota.v1.ListHotelRoomTypeReply
}
var file_ota_v1_ota_proto_depIdxs = []int32{
	2, // 0: ota.v1.GetHotelRoomTypeReply.room_type:type_name -> ota.v1.RoomType
	2, // 1: ota.v1.ListHotelRoomTypeReply.room_type_list:type_name -> ota.v1.RoomType
	0, // 2: ota.v1.Ota.GetToken:input_type -> ota.v1.GetTokenRequest
	3, // 3: ota.v1.Ota.GetHotelRoomType:input_type -> ota.v1.GetHotelRoomTypeRequest
	5, // 4: ota.v1.Ota.ListHotelRoomType:input_type -> ota.v1.ListHotelRoomTypeRequest
	1, // 5: ota.v1.Ota.GetToken:output_type -> ota.v1.GetTokenReply
	4, // 6: ota.v1.Ota.GetHotelRoomType:output_type -> ota.v1.GetHotelRoomTypeReply
	6, // 7: ota.v1.Ota.ListHotelRoomType:output_type -> ota.v1.ListHotelRoomTypeReply
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_ota_v1_ota_proto_init() }
func file_ota_v1_ota_proto_init() {
	if File_ota_v1_ota_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ota_v1_ota_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTokenRequest); i {
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
		file_ota_v1_ota_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTokenReply); i {
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
		file_ota_v1_ota_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomType); i {
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
		file_ota_v1_ota_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHotelRoomTypeRequest); i {
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
		file_ota_v1_ota_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHotelRoomTypeReply); i {
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
		file_ota_v1_ota_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListHotelRoomTypeRequest); i {
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
		file_ota_v1_ota_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListHotelRoomTypeReply); i {
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
			RawDescriptor: file_ota_v1_ota_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ota_v1_ota_proto_goTypes,
		DependencyIndexes: file_ota_v1_ota_proto_depIdxs,
		MessageInfos:      file_ota_v1_ota_proto_msgTypes,
	}.Build()
	File_ota_v1_ota_proto = out.File
	file_ota_v1_ota_proto_rawDesc = nil
	file_ota_v1_ota_proto_goTypes = nil
	file_ota_v1_ota_proto_depIdxs = nil
}
