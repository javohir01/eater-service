// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.3
// source: eater.proto

package eater

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

type EaterProfile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EaterId                string `protobuf:"bytes,1,opt,name=eater_id,proto3" json:"eater_id,omitempty"`
	PhoneNumber            string `protobuf:"bytes,2,opt,name=phone_number,proto3" json:"phone_number,omitempty"`
	Name                   string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	IsPhoneNumberConfirmed bool   `protobuf:"varint,4,opt,name=is_phone_number_confirmed,proto3" json:"is_phone_number_confirmed,omitempty"`
	CreatedAt              string `protobuf:"bytes,5,opt,name=created_at,proto3" json:"created_at,omitempty"`
	UpdatedAt              string `protobuf:"bytes,6,opt,name=updated_at,proto3" json:"updated_at,omitempty"`
}

func (x *EaterProfile) Reset() {
	*x = EaterProfile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eater_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EaterProfile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EaterProfile) ProtoMessage() {}

func (x *EaterProfile) ProtoReflect() protoreflect.Message {
	mi := &file_eater_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EaterProfile.ProtoReflect.Descriptor instead.
func (*EaterProfile) Descriptor() ([]byte, []int) {
	return file_eater_proto_rawDescGZIP(), []int{0}
}

func (x *EaterProfile) GetEaterId() string {
	if x != nil {
		return x.EaterId
	}
	return ""
}

func (x *EaterProfile) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *EaterProfile) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *EaterProfile) GetIsPhoneNumberConfirmed() bool {
	if x != nil {
		return x.IsPhoneNumberConfirmed
	}
	return false
}

func (x *EaterProfile) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *EaterProfile) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type SignupEaterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PhoneNumber string `protobuf:"bytes,1,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
}

func (x *SignupEaterRequest) Reset() {
	*x = SignupEaterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eater_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignupEaterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignupEaterRequest) ProtoMessage() {}

func (x *SignupEaterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_eater_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignupEaterRequest.ProtoReflect.Descriptor instead.
func (*SignupEaterRequest) Descriptor() ([]byte, []int) {
	return file_eater_proto_rawDescGZIP(), []int{1}
}

func (x *SignupEaterRequest) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

type SignupEaterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EaterId string `protobuf:"bytes,1,opt,name=eater_id,proto3" json:"eater_id,omitempty"`
}

func (x *SignupEaterResponse) Reset() {
	*x = SignupEaterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eater_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignupEaterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignupEaterResponse) ProtoMessage() {}

func (x *SignupEaterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_eater_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignupEaterResponse.ProtoReflect.Descriptor instead.
func (*SignupEaterResponse) Descriptor() ([]byte, []int) {
	return file_eater_proto_rawDescGZIP(), []int{2}
}

func (x *SignupEaterResponse) GetEaterId() string {
	if x != nil {
		return x.EaterId
	}
	return ""
}

type ConfirmSmsCodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EaterId string `protobuf:"bytes,1,opt,name=eater_id,json=eaterId,proto3" json:"eater_id,omitempty"`
	SmsCode string `protobuf:"bytes,2,opt,name=sms_code,json=smsCode,proto3" json:"sms_code,omitempty"`
}

func (x *ConfirmSmsCodeRequest) Reset() {
	*x = ConfirmSmsCodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eater_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfirmSmsCodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfirmSmsCodeRequest) ProtoMessage() {}

func (x *ConfirmSmsCodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_eater_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfirmSmsCodeRequest.ProtoReflect.Descriptor instead.
func (*ConfirmSmsCodeRequest) Descriptor() ([]byte, []int) {
	return file_eater_proto_rawDescGZIP(), []int{3}
}

func (x *ConfirmSmsCodeRequest) GetEaterId() string {
	if x != nil {
		return x.EaterId
	}
	return ""
}

func (x *ConfirmSmsCodeRequest) GetSmsCode() string {
	if x != nil {
		return x.SmsCode
	}
	return ""
}

type ConfirmSmsCodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Profile *EaterProfile `protobuf:"bytes,1,opt,name=profile,proto3" json:"profile,omitempty"`
	Token   string        `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *ConfirmSmsCodeResponse) Reset() {
	*x = ConfirmSmsCodeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eater_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfirmSmsCodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfirmSmsCodeResponse) ProtoMessage() {}

func (x *ConfirmSmsCodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_eater_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfirmSmsCodeResponse.ProtoReflect.Descriptor instead.
func (*ConfirmSmsCodeResponse) Descriptor() ([]byte, []int) {
	return file_eater_proto_rawDescGZIP(), []int{4}
}

func (x *ConfirmSmsCodeResponse) GetProfile() *EaterProfile {
	if x != nil {
		return x.Profile
	}
	return nil
}

func (x *ConfirmSmsCodeResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type UpdateEaterProfileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EaterId  string `protobuf:"bytes,1,opt,name=eater_id,proto3" json:"eater_id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ImageUrl string `protobuf:"bytes,3,opt,name=image_url,proto3" json:"image_url,omitempty"`
}

func (x *UpdateEaterProfileRequest) Reset() {
	*x = UpdateEaterProfileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eater_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateEaterProfileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateEaterProfileRequest) ProtoMessage() {}

func (x *UpdateEaterProfileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_eater_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateEaterProfileRequest.ProtoReflect.Descriptor instead.
func (*UpdateEaterProfileRequest) Descriptor() ([]byte, []int) {
	return file_eater_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateEaterProfileRequest) GetEaterId() string {
	if x != nil {
		return x.EaterId
	}
	return ""
}

func (x *UpdateEaterProfileRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateEaterProfileRequest) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

type UpdateEaterProfileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Profile *EaterProfile `protobuf:"bytes,1,opt,name=profile,proto3" json:"profile,omitempty"`
}

func (x *UpdateEaterProfileResponse) Reset() {
	*x = UpdateEaterProfileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eater_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateEaterProfileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateEaterProfileResponse) ProtoMessage() {}

func (x *UpdateEaterProfileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_eater_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateEaterProfileResponse.ProtoReflect.Descriptor instead.
func (*UpdateEaterProfileResponse) Descriptor() ([]byte, []int) {
	return file_eater_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateEaterProfileResponse) GetProfile() *EaterProfile {
	if x != nil {
		return x.Profile
	}
	return nil
}

type GetEaterProfileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EaterId string `protobuf:"bytes,1,opt,name=eater_id,json=eaterId,proto3" json:"eater_id,omitempty"`
}

func (x *GetEaterProfileRequest) Reset() {
	*x = GetEaterProfileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eater_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEaterProfileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEaterProfileRequest) ProtoMessage() {}

func (x *GetEaterProfileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_eater_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEaterProfileRequest.ProtoReflect.Descriptor instead.
func (*GetEaterProfileRequest) Descriptor() ([]byte, []int) {
	return file_eater_proto_rawDescGZIP(), []int{7}
}

func (x *GetEaterProfileRequest) GetEaterId() string {
	if x != nil {
		return x.EaterId
	}
	return ""
}

type GetEaterProfileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Profile *EaterProfile `protobuf:"bytes,1,opt,name=profile,proto3" json:"profile,omitempty"`
}

func (x *GetEaterProfileResponse) Reset() {
	*x = GetEaterProfileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eater_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEaterProfileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEaterProfileResponse) ProtoMessage() {}

func (x *GetEaterProfileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_eater_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEaterProfileResponse.ProtoReflect.Descriptor instead.
func (*GetEaterProfileResponse) Descriptor() ([]byte, []int) {
	return file_eater_proto_rawDescGZIP(), []int{8}
}

func (x *GetEaterProfileResponse) GetProfile() *EaterProfile {
	if x != nil {
		return x.Profile
	}
	return nil
}

var File_eater_proto protoreflect.FileDescriptor

var file_eater_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x65, 0x61, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe0, 0x01,
	0x0a, 0x0c, 0x45, 0x61, 0x74, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x65, 0x61, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x65, 0x61, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x3c, 0x0a, 0x19, 0x69, 0x73, 0x5f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x19, 0x69, 0x73, 0x5f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64,
	0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x12, 0x1e, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x22, 0x37, 0x0a, 0x12, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x45, 0x61, 0x74, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x31, 0x0a, 0x13, 0x53, 0x69, 0x67,
	0x6e, 0x75, 0x70, 0x45, 0x61, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x65, 0x61, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x65, 0x61, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x22, 0x4d, 0x0a, 0x15,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x53, 0x6d, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x61, 0x74, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x61, 0x74, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x19, 0x0a, 0x08, 0x73, 0x6d, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x73, 0x6d, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x57, 0x0a, 0x16, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x53, 0x6d, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x45, 0x61, 0x74, 0x65, 0x72, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x69, 0x0a, 0x19, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x61,
	0x74, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x61, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x61, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x22,
	0x45, 0x0a, 0x1a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x61, 0x74, 0x65, 0x72, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a,
	0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x45, 0x61, 0x74, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x07, 0x70,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x22, 0x33, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x45, 0x61, 0x74,
	0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x19, 0x0a, 0x08, 0x65, 0x61, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x65, 0x61, 0x74, 0x65, 0x72, 0x49, 0x64, 0x22, 0x42, 0x0a, 0x17, 0x47,
	0x65, 0x74, 0x45, 0x61, 0x74, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x45, 0x61, 0x74, 0x65, 0x72, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x42,
	0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x65, 0x61, 0x74, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_eater_proto_rawDescOnce sync.Once
	file_eater_proto_rawDescData = file_eater_proto_rawDesc
)

func file_eater_proto_rawDescGZIP() []byte {
	file_eater_proto_rawDescOnce.Do(func() {
		file_eater_proto_rawDescData = protoimpl.X.CompressGZIP(file_eater_proto_rawDescData)
	})
	return file_eater_proto_rawDescData
}

var file_eater_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_eater_proto_goTypes = []interface{}{
	(*EaterProfile)(nil),               // 0: EaterProfile
	(*SignupEaterRequest)(nil),         // 1: SignupEaterRequest
	(*SignupEaterResponse)(nil),        // 2: SignupEaterResponse
	(*ConfirmSmsCodeRequest)(nil),      // 3: ConfirmSmsCodeRequest
	(*ConfirmSmsCodeResponse)(nil),     // 4: ConfirmSmsCodeResponse
	(*UpdateEaterProfileRequest)(nil),  // 5: UpdateEaterProfileRequest
	(*UpdateEaterProfileResponse)(nil), // 6: UpdateEaterProfileResponse
	(*GetEaterProfileRequest)(nil),     // 7: GetEaterProfileRequest
	(*GetEaterProfileResponse)(nil),    // 8: GetEaterProfileResponse
}
var file_eater_proto_depIdxs = []int32{
	0, // 0: ConfirmSmsCodeResponse.profile:type_name -> EaterProfile
	0, // 1: UpdateEaterProfileResponse.profile:type_name -> EaterProfile
	0, // 2: GetEaterProfileResponse.profile:type_name -> EaterProfile
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_eater_proto_init() }
func file_eater_proto_init() {
	if File_eater_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_eater_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EaterProfile); i {
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
		file_eater_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignupEaterRequest); i {
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
		file_eater_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignupEaterResponse); i {
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
		file_eater_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfirmSmsCodeRequest); i {
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
		file_eater_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfirmSmsCodeResponse); i {
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
		file_eater_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateEaterProfileRequest); i {
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
		file_eater_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateEaterProfileResponse); i {
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
		file_eater_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEaterProfileRequest); i {
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
		file_eater_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEaterProfileResponse); i {
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
			RawDescriptor: file_eater_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eater_proto_goTypes,
		DependencyIndexes: file_eater_proto_depIdxs,
		MessageInfos:      file_eater_proto_msgTypes,
	}.Build()
	File_eater_proto = out.File
	file_eater_proto_rawDesc = nil
	file_eater_proto_goTypes = nil
	file_eater_proto_depIdxs = nil
}
