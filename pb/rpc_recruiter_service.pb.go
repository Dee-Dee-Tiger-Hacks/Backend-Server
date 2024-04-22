// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v4.25.2
// source: rpc_recruiter_service.proto

package pb

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

type Recruiter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId          string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	LinkedinUrl     string `protobuf:"bytes,3,opt,name=linkedin_url,json=linkedinUrl,proto3" json:"linkedin_url,omitempty"`
	Name            string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Email           string `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	Company         string `protobuf:"bytes,6,opt,name=company,proto3" json:"company,omitempty"`
	Overview        string `protobuf:"bytes,7,opt,name=overview,proto3" json:"overview,omitempty"`
	SuggestedEmail  string `protobuf:"bytes,8,opt,name=suggested_email,json=suggestedEmail,proto3" json:"suggested_email,omitempty"`
	PotentialTopics string `protobuf:"bytes,9,opt,name=potential_topics,json=potentialTopics,proto3" json:"potential_topics,omitempty"`
}

func (x *Recruiter) Reset() {
	*x = Recruiter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_recruiter_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Recruiter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Recruiter) ProtoMessage() {}

func (x *Recruiter) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_recruiter_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Recruiter.ProtoReflect.Descriptor instead.
func (*Recruiter) Descriptor() ([]byte, []int) {
	return file_rpc_recruiter_service_proto_rawDescGZIP(), []int{0}
}

func (x *Recruiter) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Recruiter) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Recruiter) GetLinkedinUrl() string {
	if x != nil {
		return x.LinkedinUrl
	}
	return ""
}

func (x *Recruiter) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Recruiter) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Recruiter) GetCompany() string {
	if x != nil {
		return x.Company
	}
	return ""
}

func (x *Recruiter) GetOverview() string {
	if x != nil {
		return x.Overview
	}
	return ""
}

func (x *Recruiter) GetSuggestedEmail() string {
	if x != nil {
		return x.SuggestedEmail
	}
	return ""
}

func (x *Recruiter) GetPotentialTopics() string {
	if x != nil {
		return x.PotentialTopics
	}
	return ""
}

// Create Recruiter
type CreateRecruiterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId          string  `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	LinkedinUrl     *string `protobuf:"bytes,3,opt,name=linkedin_url,json=linkedinUrl,proto3,oneof" json:"linkedin_url,omitempty"`
	Name            *string `protobuf:"bytes,4,opt,name=name,proto3,oneof" json:"name,omitempty"`
	Email           *string `protobuf:"bytes,5,opt,name=email,proto3,oneof" json:"email,omitempty"`
	Company         *string `protobuf:"bytes,6,opt,name=company,proto3,oneof" json:"company,omitempty"`
	Overview        *string `protobuf:"bytes,7,opt,name=overview,proto3,oneof" json:"overview,omitempty"`
	SuggestedEmail  *string `protobuf:"bytes,8,opt,name=suggested_email,json=suggestedEmail,proto3,oneof" json:"suggested_email,omitempty"`
	PotentialTopics *string `protobuf:"bytes,9,opt,name=potential_topics,json=potentialTopics,proto3,oneof" json:"potential_topics,omitempty"`
}

func (x *CreateRecruiterRequest) Reset() {
	*x = CreateRecruiterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_recruiter_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRecruiterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRecruiterRequest) ProtoMessage() {}

func (x *CreateRecruiterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_recruiter_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRecruiterRequest.ProtoReflect.Descriptor instead.
func (*CreateRecruiterRequest) Descriptor() ([]byte, []int) {
	return file_rpc_recruiter_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateRecruiterRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateRecruiterRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CreateRecruiterRequest) GetLinkedinUrl() string {
	if x != nil && x.LinkedinUrl != nil {
		return *x.LinkedinUrl
	}
	return ""
}

func (x *CreateRecruiterRequest) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *CreateRecruiterRequest) GetEmail() string {
	if x != nil && x.Email != nil {
		return *x.Email
	}
	return ""
}

func (x *CreateRecruiterRequest) GetCompany() string {
	if x != nil && x.Company != nil {
		return *x.Company
	}
	return ""
}

func (x *CreateRecruiterRequest) GetOverview() string {
	if x != nil && x.Overview != nil {
		return *x.Overview
	}
	return ""
}

func (x *CreateRecruiterRequest) GetSuggestedEmail() string {
	if x != nil && x.SuggestedEmail != nil {
		return *x.SuggestedEmail
	}
	return ""
}

func (x *CreateRecruiterRequest) GetPotentialTopics() string {
	if x != nil && x.PotentialTopics != nil {
		return *x.PotentialTopics
	}
	return ""
}

type CreateRecruiterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Recruiter *Recruiter `protobuf:"bytes,1,opt,name=recruiter,proto3" json:"recruiter,omitempty"`
}

func (x *CreateRecruiterResponse) Reset() {
	*x = CreateRecruiterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_recruiter_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRecruiterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRecruiterResponse) ProtoMessage() {}

func (x *CreateRecruiterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_recruiter_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRecruiterResponse.ProtoReflect.Descriptor instead.
func (*CreateRecruiterResponse) Descriptor() ([]byte, []int) {
	return file_rpc_recruiter_service_proto_rawDescGZIP(), []int{2}
}

func (x *CreateRecruiterResponse) GetRecruiter() *Recruiter {
	if x != nil {
		return x.Recruiter
	}
	return nil
}

// Get Recruiter
type GetRecruiterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RecruiterId string `protobuf:"bytes,1,opt,name=recruiter_id,json=recruiterId,proto3" json:"recruiter_id,omitempty"`
}

func (x *GetRecruiterRequest) Reset() {
	*x = GetRecruiterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_recruiter_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRecruiterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRecruiterRequest) ProtoMessage() {}

func (x *GetRecruiterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_recruiter_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRecruiterRequest.ProtoReflect.Descriptor instead.
func (*GetRecruiterRequest) Descriptor() ([]byte, []int) {
	return file_rpc_recruiter_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetRecruiterRequest) GetRecruiterId() string {
	if x != nil {
		return x.RecruiterId
	}
	return ""
}

type GetRecruiterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Recruiter *Recruiter `protobuf:"bytes,1,opt,name=recruiter,proto3" json:"recruiter,omitempty"`
}

func (x *GetRecruiterResponse) Reset() {
	*x = GetRecruiterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_recruiter_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRecruiterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRecruiterResponse) ProtoMessage() {}

func (x *GetRecruiterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_recruiter_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRecruiterResponse.ProtoReflect.Descriptor instead.
func (*GetRecruiterResponse) Descriptor() ([]byte, []int) {
	return file_rpc_recruiter_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetRecruiterResponse) GetRecruiter() *Recruiter {
	if x != nil {
		return x.Recruiter
	}
	return nil
}

// Get Recruiters
type GetRecruitersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetRecruitersRequest) Reset() {
	*x = GetRecruitersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_recruiter_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRecruitersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRecruitersRequest) ProtoMessage() {}

func (x *GetRecruitersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_recruiter_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRecruitersRequest.ProtoReflect.Descriptor instead.
func (*GetRecruitersRequest) Descriptor() ([]byte, []int) {
	return file_rpc_recruiter_service_proto_rawDescGZIP(), []int{5}
}

func (x *GetRecruitersRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type GetRecruitersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Recruiters []*Recruiter `protobuf:"bytes,1,rep,name=recruiters,proto3" json:"recruiters,omitempty"`
}

func (x *GetRecruitersResponse) Reset() {
	*x = GetRecruitersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_recruiter_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRecruitersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRecruitersResponse) ProtoMessage() {}

func (x *GetRecruitersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_recruiter_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRecruitersResponse.ProtoReflect.Descriptor instead.
func (*GetRecruitersResponse) Descriptor() ([]byte, []int) {
	return file_rpc_recruiter_service_proto_rawDescGZIP(), []int{6}
}

func (x *GetRecruitersResponse) GetRecruiters() []*Recruiter {
	if x != nil {
		return x.Recruiters
	}
	return nil
}

var File_rpc_recruiter_service_proto protoreflect.FileDescriptor

var file_rpc_recruiter_service_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x72, 0x70, 0x63, 0x5f, 0x72, 0x65, 0x63, 0x72, 0x75, 0x69, 0x74, 0x65, 0x72, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70,
	0x62, 0x22, 0x8b, 0x02, 0x0a, 0x09, 0x52, 0x65, 0x63, 0x72, 0x75, 0x69, 0x74, 0x65, 0x72, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x6c, 0x69, 0x6e, 0x6b,
	0x65, 0x64, 0x69, 0x6e, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x6c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x69, 0x6e, 0x55, 0x72, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12,
	0x1a, 0x0a, 0x08, 0x6f, 0x76, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6f, 0x76, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x12, 0x27, 0x0a, 0x0f, 0x73,
	0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x65, 0x64, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x29, 0x0a, 0x10, 0x70, 0x6f, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x61,
	0x6c, 0x5f, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f,
	0x70, 0x6f, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x22,
	0xa1, 0x03, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x63, 0x72, 0x75, 0x69,
	0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0c, 0x6c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x69, 0x6e, 0x5f,
	0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x6c, 0x69, 0x6e,
	0x6b, 0x65, 0x64, 0x69, 0x6e, 0x55, 0x72, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x88, 0x01, 0x01, 0x12,
	0x1d, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x03, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x88, 0x01, 0x01, 0x12, 0x1f,
	0x0a, 0x08, 0x6f, 0x76, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x04, 0x52, 0x08, 0x6f, 0x76, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x88, 0x01, 0x01, 0x12,
	0x2c, 0x0a, 0x0f, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x48, 0x05, 0x52, 0x0e, 0x73, 0x75, 0x67, 0x67,
	0x65, 0x73, 0x74, 0x65, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x2e, 0x0a,
	0x10, 0x70, 0x6f, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x74, 0x6f, 0x70, 0x69, 0x63,
	0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x48, 0x06, 0x52, 0x0f, 0x70, 0x6f, 0x74, 0x65, 0x6e,
	0x74, 0x69, 0x61, 0x6c, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x88, 0x01, 0x01, 0x42, 0x0f, 0x0a,
	0x0d, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x69, 0x6e, 0x5f, 0x75, 0x72, 0x6c, 0x42, 0x07,
	0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x42, 0x0b, 0x0a,
	0x09, 0x5f, 0x6f, 0x76, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x73,
	0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x42, 0x13,
	0x0a, 0x11, 0x5f, 0x70, 0x6f, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x74, 0x6f, 0x70,
	0x69, 0x63, 0x73, 0x22, 0x46, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x63,
	0x72, 0x75, 0x69, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b,
	0x0a, 0x09, 0x72, 0x65, 0x63, 0x72, 0x75, 0x69, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x63, 0x72, 0x75, 0x69, 0x74, 0x65, 0x72,
	0x52, 0x09, 0x72, 0x65, 0x63, 0x72, 0x75, 0x69, 0x74, 0x65, 0x72, 0x22, 0x38, 0x0a, 0x13, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x63, 0x72, 0x75, 0x69, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x63, 0x72, 0x75, 0x69, 0x74, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x63, 0x72, 0x75, 0x69,
	0x74, 0x65, 0x72, 0x49, 0x64, 0x22, 0x43, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x72,
	0x75, 0x69, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a,
	0x09, 0x72, 0x65, 0x63, 0x72, 0x75, 0x69, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x63, 0x72, 0x75, 0x69, 0x74, 0x65, 0x72, 0x52,
	0x09, 0x72, 0x65, 0x63, 0x72, 0x75, 0x69, 0x74, 0x65, 0x72, 0x22, 0x2f, 0x0a, 0x14, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x63, 0x72, 0x75, 0x69, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x46, 0x0a, 0x15, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x63, 0x72, 0x75, 0x69, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x0a, 0x72, 0x65, 0x63, 0x72, 0x75, 0x69, 0x74, 0x65,
	0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65,
	0x63, 0x72, 0x75, 0x69, 0x74, 0x65, 0x72, 0x52, 0x0a, 0x72, 0x65, 0x63, 0x72, 0x75, 0x69, 0x74,
	0x65, 0x72, 0x73, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x44, 0x65, 0x65, 0x2d, 0x44, 0x65, 0x65, 0x2d, 0x54, 0x69, 0x67, 0x65, 0x72, 0x2d,
	0x48, 0x61, 0x63, 0x6b, 0x73, 0x2f, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_recruiter_service_proto_rawDescOnce sync.Once
	file_rpc_recruiter_service_proto_rawDescData = file_rpc_recruiter_service_proto_rawDesc
)

func file_rpc_recruiter_service_proto_rawDescGZIP() []byte {
	file_rpc_recruiter_service_proto_rawDescOnce.Do(func() {
		file_rpc_recruiter_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_recruiter_service_proto_rawDescData)
	})
	return file_rpc_recruiter_service_proto_rawDescData
}

var file_rpc_recruiter_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_rpc_recruiter_service_proto_goTypes = []interface{}{
	(*Recruiter)(nil),               // 0: pb.Recruiter
	(*CreateRecruiterRequest)(nil),  // 1: pb.CreateRecruiterRequest
	(*CreateRecruiterResponse)(nil), // 2: pb.CreateRecruiterResponse
	(*GetRecruiterRequest)(nil),     // 3: pb.GetRecruiterRequest
	(*GetRecruiterResponse)(nil),    // 4: pb.GetRecruiterResponse
	(*GetRecruitersRequest)(nil),    // 5: pb.GetRecruitersRequest
	(*GetRecruitersResponse)(nil),   // 6: pb.GetRecruitersResponse
}
var file_rpc_recruiter_service_proto_depIdxs = []int32{
	0, // 0: pb.CreateRecruiterResponse.recruiter:type_name -> pb.Recruiter
	0, // 1: pb.GetRecruiterResponse.recruiter:type_name -> pb.Recruiter
	0, // 2: pb.GetRecruitersResponse.recruiters:type_name -> pb.Recruiter
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_rpc_recruiter_service_proto_init() }
func file_rpc_recruiter_service_proto_init() {
	if File_rpc_recruiter_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_recruiter_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Recruiter); i {
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
		file_rpc_recruiter_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRecruiterRequest); i {
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
		file_rpc_recruiter_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRecruiterResponse); i {
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
		file_rpc_recruiter_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRecruiterRequest); i {
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
		file_rpc_recruiter_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRecruiterResponse); i {
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
		file_rpc_recruiter_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRecruitersRequest); i {
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
		file_rpc_recruiter_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRecruitersResponse); i {
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
	file_rpc_recruiter_service_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_rpc_recruiter_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rpc_recruiter_service_proto_goTypes,
		DependencyIndexes: file_rpc_recruiter_service_proto_depIdxs,
		MessageInfos:      file_rpc_recruiter_service_proto_msgTypes,
	}.Build()
	File_rpc_recruiter_service_proto = out.File
	file_rpc_recruiter_service_proto_rawDesc = nil
	file_rpc_recruiter_service_proto_goTypes = nil
	file_rpc_recruiter_service_proto_depIdxs = nil
}
