// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.2
// source: stage_service.proto

package proto

import (
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

type GetListStagesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Booking *Booking `protobuf:"bytes,1,opt,name=booking,proto3" json:"booking,omitempty"`
}

func (x *GetListStagesRequest) Reset() {
	*x = GetListStagesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stage_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListStagesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListStagesRequest) ProtoMessage() {}

func (x *GetListStagesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stage_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListStagesRequest.ProtoReflect.Descriptor instead.
func (*GetListStagesRequest) Descriptor() ([]byte, []int) {
	return file_stage_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetListStagesRequest) GetBooking() *Booking {
	if x != nil {
		return x.Booking
	}
	return nil
}

type GetListStagesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stages []*Stage `protobuf:"bytes,1,rep,name=stages,proto3" json:"stages,omitempty"`
}

func (x *GetListStagesResponse) Reset() {
	*x = GetListStagesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stage_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListStagesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListStagesResponse) ProtoMessage() {}

func (x *GetListStagesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stage_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListStagesResponse.ProtoReflect.Descriptor instead.
func (*GetListStagesResponse) Descriptor() ([]byte, []int) {
	return file_stage_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetListStagesResponse) GetStages() []*Stage {
	if x != nil {
		return x.Stages
	}
	return nil
}

type CreateStageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BookingId string `protobuf:"bytes,1,opt,name=booking_id,json=bookingId,proto3" json:"booking_id,omitempty"`
}

func (x *CreateStageRequest) Reset() {
	*x = CreateStageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stage_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateStageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateStageRequest) ProtoMessage() {}

func (x *CreateStageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stage_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateStageRequest.ProtoReflect.Descriptor instead.
func (*CreateStageRequest) Descriptor() ([]byte, []int) {
	return file_stage_service_proto_rawDescGZIP(), []int{2}
}

func (x *CreateStageRequest) GetBookingId() string {
	if x != nil {
		return x.BookingId
	}
	return ""
}

type CreateStageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StageId string `protobuf:"bytes,1,opt,name=stage_id,json=stageId,proto3" json:"stage_id,omitempty"`
}

func (x *CreateStageResponse) Reset() {
	*x = CreateStageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stage_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateStageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateStageResponse) ProtoMessage() {}

func (x *CreateStageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stage_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateStageResponse.ProtoReflect.Descriptor instead.
func (*CreateStageResponse) Descriptor() ([]byte, []int) {
	return file_stage_service_proto_rawDescGZIP(), []int{3}
}

func (x *CreateStageResponse) GetStageId() string {
	if x != nil {
		return x.StageId
	}
	return ""
}

type UpdateStageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stage *Stage `protobuf:"bytes,1,opt,name=stage,proto3" json:"stage,omitempty"`
}

func (x *UpdateStageRequest) Reset() {
	*x = UpdateStageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stage_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateStageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateStageRequest) ProtoMessage() {}

func (x *UpdateStageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stage_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateStageRequest.ProtoReflect.Descriptor instead.
func (*UpdateStageRequest) Descriptor() ([]byte, []int) {
	return file_stage_service_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateStageRequest) GetStage() *Stage {
	if x != nil {
		return x.Stage
	}
	return nil
}

type UpdateStageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stage *Stage `protobuf:"bytes,1,opt,name=stage,proto3" json:"stage,omitempty"`
}

func (x *UpdateStageResponse) Reset() {
	*x = UpdateStageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stage_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateStageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateStageResponse) ProtoMessage() {}

func (x *UpdateStageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stage_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateStageResponse.ProtoReflect.Descriptor instead.
func (*UpdateStageResponse) Descriptor() ([]byte, []int) {
	return file_stage_service_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateStageResponse) GetStage() *Stage {
	if x != nil {
		return x.Stage
	}
	return nil
}

type RemoveStageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stage *Stage `protobuf:"bytes,1,opt,name=stage,proto3" json:"stage,omitempty"`
}

func (x *RemoveStageRequest) Reset() {
	*x = RemoveStageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stage_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveStageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveStageRequest) ProtoMessage() {}

func (x *RemoveStageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stage_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveStageRequest.ProtoReflect.Descriptor instead.
func (*RemoveStageRequest) Descriptor() ([]byte, []int) {
	return file_stage_service_proto_rawDescGZIP(), []int{6}
}

func (x *RemoveStageRequest) GetStage() *Stage {
	if x != nil {
		return x.Stage
	}
	return nil
}

type RemoveStageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StageId string `protobuf:"bytes,1,opt,name=stage_id,json=stageId,proto3" json:"stage_id,omitempty"`
}

func (x *RemoveStageResponse) Reset() {
	*x = RemoveStageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stage_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveStageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveStageResponse) ProtoMessage() {}

func (x *RemoveStageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stage_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveStageResponse.ProtoReflect.Descriptor instead.
func (*RemoveStageResponse) Descriptor() ([]byte, []int) {
	return file_stage_service_proto_rawDescGZIP(), []int{7}
}

func (x *RemoveStageResponse) GetStageId() string {
	if x != nil {
		return x.StageId
	}
	return ""
}

type IssueIDMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IssueId string `protobuf:"bytes,1,opt,name=issue_id,json=issueId,proto3" json:"issue_id,omitempty"`
}

func (x *IssueIDMessage) Reset() {
	*x = IssueIDMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stage_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IssueIDMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IssueIDMessage) ProtoMessage() {}

func (x *IssueIDMessage) ProtoReflect() protoreflect.Message {
	mi := &file_stage_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IssueIDMessage.ProtoReflect.Descriptor instead.
func (*IssueIDMessage) Descriptor() ([]byte, []int) {
	return file_stage_service_proto_rawDescGZIP(), []int{8}
}

func (x *IssueIDMessage) GetIssueId() string {
	if x != nil {
		return x.IssueId
	}
	return ""
}

type CreateIssueRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StageId string `protobuf:"bytes,1,opt,name=stage_id,json=stageId,proto3" json:"stage_id,omitempty"`
}

func (x *CreateIssueRequest) Reset() {
	*x = CreateIssueRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stage_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateIssueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateIssueRequest) ProtoMessage() {}

func (x *CreateIssueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stage_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateIssueRequest.ProtoReflect.Descriptor instead.
func (*CreateIssueRequest) Descriptor() ([]byte, []int) {
	return file_stage_service_proto_rawDescGZIP(), []int{9}
}

func (x *CreateIssueRequest) GetStageId() string {
	if x != nil {
		return x.StageId
	}
	return ""
}

type UpdateIssueRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                     string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                  string    `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description            string    `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Timeline               *Timeline `protobuf:"bytes,4,opt,name=timeline,proto3" json:"timeline,omitempty"`
	ParticipantPortalCodes []uint32  `protobuf:"varint,5,rep,packed,name=participant_portal_codes,json=participantPortalCodes,proto3" json:"participant_portal_codes,omitempty"`
	AttachmentId           []string  `protobuf:"bytes,6,rep,name=attachment_id,json=attachmentId,proto3" json:"attachment_id,omitempty"`
}

func (x *UpdateIssueRequest) Reset() {
	*x = UpdateIssueRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stage_service_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateIssueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateIssueRequest) ProtoMessage() {}

func (x *UpdateIssueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stage_service_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateIssueRequest.ProtoReflect.Descriptor instead.
func (*UpdateIssueRequest) Descriptor() ([]byte, []int) {
	return file_stage_service_proto_rawDescGZIP(), []int{10}
}

func (x *UpdateIssueRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateIssueRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UpdateIssueRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateIssueRequest) GetTimeline() *Timeline {
	if x != nil {
		return x.Timeline
	}
	return nil
}

func (x *UpdateIssueRequest) GetParticipantPortalCodes() []uint32 {
	if x != nil {
		return x.ParticipantPortalCodes
	}
	return nil
}

func (x *UpdateIssueRequest) GetAttachmentId() []string {
	if x != nil {
		return x.AttachmentId
	}
	return nil
}

var File_stage_service_proto protoreflect.FileDescriptor

var file_stage_service_proto_rawDesc = []byte{
	0x0a, 0x13, 0x73, 0x74, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x1a, 0x0d,
	0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x73,
	0x74, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x42, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x53, 0x74, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a,
	0x0a, 0x07, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e,
	0x67, 0x52, 0x07, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x22, 0x3f, 0x0a, 0x15, 0x47, 0x65,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x74,
	0x61, 0x67, 0x65, 0x52, 0x06, 0x73, 0x74, 0x61, 0x67, 0x65, 0x73, 0x22, 0x33, 0x0a, 0x12, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x64,
	0x22, 0x30, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x67, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x74, 0x61, 0x67, 0x65,
	0x49, 0x64, 0x22, 0x3a, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e,
	0x67, 0x2e, 0x53, 0x74, 0x61, 0x67, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x67, 0x65, 0x22, 0x3b,
	0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x53,
	0x74, 0x61, 0x67, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x67, 0x65, 0x22, 0x3a, 0x0a, 0x12, 0x52,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x53, 0x74, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x24, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x74, 0x61, 0x67, 0x65,
	0x52, 0x05, 0x73, 0x74, 0x61, 0x67, 0x65, 0x22, 0x30, 0x0a, 0x13, 0x52, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x53, 0x74, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19,
	0x0a, 0x08, 0x73, 0x74, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x73, 0x74, 0x61, 0x67, 0x65, 0x49, 0x64, 0x22, 0x2b, 0x0a, 0x0e, 0x49, 0x73, 0x73,
	0x75, 0x65, 0x49, 0x44, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x69,
	0x73, 0x73, 0x75, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x69,
	0x73, 0x73, 0x75, 0x65, 0x49, 0x64, 0x22, 0x2f, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x49, 0x73, 0x73, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08,
	0x73, 0x74, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x73, 0x74, 0x61, 0x67, 0x65, 0x49, 0x64, 0x22, 0xea, 0x01, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x49, 0x73, 0x73, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2d, 0x0a, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69,
	0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69,
	0x6e, 0x67, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x08, 0x74, 0x69, 0x6d,
	0x65, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x38, 0x0a, 0x18, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69,
	0x70, 0x61, 0x6e, 0x74, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x16, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69,
	0x70, 0x61, 0x6e, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x64, 0x65, 0x73, 0x12,
	0x23, 0x0a, 0x0d, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x32, 0x8b, 0x04, 0x0a, 0x0c, 0x53, 0x74, 0x61, 0x67, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x50, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74,
	0x53, 0x74, 0x61, 0x67, 0x65, 0x73, 0x12, 0x1d, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67,
	0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e,
	0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x53, 0x74, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61,
	0x67, 0x65, 0x12, 0x1b, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1c, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x53, 0x74, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x4a, 0x0a, 0x0b, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x53, 0x74, 0x61, 0x67, 0x65, 0x12, 0x1b,
	0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x53,
	0x74, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x62, 0x6f,
	0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x53, 0x74, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x0b, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x73, 0x73, 0x75, 0x65, 0x12, 0x1b, 0x2e, 0x62, 0x6f, 0x6f,
	0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x73, 0x73, 0x75, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e,
	0x67, 0x2e, 0x49, 0x73, 0x73, 0x75, 0x65, 0x49, 0x44, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x00, 0x12, 0x3c, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x73, 0x73, 0x75,
	0x65, 0x12, 0x1b, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x49, 0x73, 0x73, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e,
	0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x49, 0x73, 0x73, 0x75, 0x65, 0x22, 0x00,
	0x12, 0x40, 0x0a, 0x0b, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x49, 0x73, 0x73, 0x75, 0x65, 0x12,
	0x17, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x49, 0x73, 0x73, 0x75, 0x65, 0x49,
	0x44, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x00, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_stage_service_proto_rawDescOnce sync.Once
	file_stage_service_proto_rawDescData = file_stage_service_proto_rawDesc
)

func file_stage_service_proto_rawDescGZIP() []byte {
	file_stage_service_proto_rawDescOnce.Do(func() {
		file_stage_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_stage_service_proto_rawDescData)
	})
	return file_stage_service_proto_rawDescData
}

var file_stage_service_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_stage_service_proto_goTypes = []interface{}{
	(*GetListStagesRequest)(nil),  // 0: booking.GetListStagesRequest
	(*GetListStagesResponse)(nil), // 1: booking.GetListStagesResponse
	(*CreateStageRequest)(nil),    // 2: booking.CreateStageRequest
	(*CreateStageResponse)(nil),   // 3: booking.CreateStageResponse
	(*UpdateStageRequest)(nil),    // 4: booking.UpdateStageRequest
	(*UpdateStageResponse)(nil),   // 5: booking.UpdateStageResponse
	(*RemoveStageRequest)(nil),    // 6: booking.RemoveStageRequest
	(*RemoveStageResponse)(nil),   // 7: booking.RemoveStageResponse
	(*IssueIDMessage)(nil),        // 8: booking.IssueIDMessage
	(*CreateIssueRequest)(nil),    // 9: booking.CreateIssueRequest
	(*UpdateIssueRequest)(nil),    // 10: booking.UpdateIssueRequest
	(*Booking)(nil),               // 11: booking.Booking
	(*Stage)(nil),                 // 12: booking.Stage
	(*Timeline)(nil),              // 13: booking.Timeline
	(*Issue)(nil),                 // 14: booking.Issue
	(*emptypb.Empty)(nil),         // 15: google.protobuf.Empty
}
var file_stage_service_proto_depIdxs = []int32{
	11, // 0: booking.GetListStagesRequest.booking:type_name -> booking.Booking
	12, // 1: booking.GetListStagesResponse.stages:type_name -> booking.Stage
	12, // 2: booking.UpdateStageRequest.stage:type_name -> booking.Stage
	12, // 3: booking.UpdateStageResponse.stage:type_name -> booking.Stage
	12, // 4: booking.RemoveStageRequest.stage:type_name -> booking.Stage
	13, // 5: booking.UpdateIssueRequest.timeline:type_name -> booking.Timeline
	0,  // 6: booking.StageService.GetListStages:input_type -> booking.GetListStagesRequest
	2,  // 7: booking.StageService.CreateStage:input_type -> booking.CreateStageRequest
	4,  // 8: booking.StageService.UpdateStage:input_type -> booking.UpdateStageRequest
	6,  // 9: booking.StageService.RemoveStage:input_type -> booking.RemoveStageRequest
	9,  // 10: booking.StageService.CreateIssue:input_type -> booking.CreateIssueRequest
	10, // 11: booking.StageService.UpdateIssue:input_type -> booking.UpdateIssueRequest
	8,  // 12: booking.StageService.RemoveIssue:input_type -> booking.IssueIDMessage
	1,  // 13: booking.StageService.GetListStages:output_type -> booking.GetListStagesResponse
	3,  // 14: booking.StageService.CreateStage:output_type -> booking.CreateStageResponse
	5,  // 15: booking.StageService.UpdateStage:output_type -> booking.UpdateStageResponse
	7,  // 16: booking.StageService.RemoveStage:output_type -> booking.RemoveStageResponse
	8,  // 17: booking.StageService.CreateIssue:output_type -> booking.IssueIDMessage
	14, // 18: booking.StageService.UpdateIssue:output_type -> booking.Issue
	15, // 19: booking.StageService.RemoveIssue:output_type -> google.protobuf.Empty
	13, // [13:20] is the sub-list for method output_type
	6,  // [6:13] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_stage_service_proto_init() }
func file_stage_service_proto_init() {
	if File_stage_service_proto != nil {
		return
	}
	file_booking_proto_init()
	file_stage_proto_init()
	file_address_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_stage_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListStagesRequest); i {
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
		file_stage_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListStagesResponse); i {
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
		file_stage_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateStageRequest); i {
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
		file_stage_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateStageResponse); i {
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
		file_stage_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateStageRequest); i {
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
		file_stage_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateStageResponse); i {
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
		file_stage_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveStageRequest); i {
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
		file_stage_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveStageResponse); i {
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
		file_stage_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IssueIDMessage); i {
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
		file_stage_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateIssueRequest); i {
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
		file_stage_service_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateIssueRequest); i {
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
			RawDescriptor: file_stage_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_stage_service_proto_goTypes,
		DependencyIndexes: file_stage_service_proto_depIdxs,
		MessageInfos:      file_stage_service_proto_msgTypes,
	}.Build()
	File_stage_service_proto = out.File
	file_stage_service_proto_rawDesc = nil
	file_stage_service_proto_goTypes = nil
	file_stage_service_proto_depIdxs = nil
}
