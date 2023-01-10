// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.2
// source: room.proto

package proto

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

type Room struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid           string           `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Author         *User            `protobuf:"bytes,2,opt,name=author,proto3" json:"author,omitempty"`
	Title          string           `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Description    string           `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	CreatedAt      string           `protobuf:"bytes,5,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt      string           `protobuf:"bytes,6,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	TargetAudience string           `protobuf:"bytes,7,opt,name=targetAudience,proto3" json:"targetAudience,omitempty"`
	Status         AddressStatus    `protobuf:"varint,8,opt,name=status,proto3,enum=booking.AddressStatus" json:"status,omitempty"`
	Assignee       *User            `protobuf:"bytes,9,opt,name=assignee,proto3" json:"assignee,omitempty"` //Согласующий (employee)
	Outmembers     []*Outmember     `protobuf:"bytes,10,rep,name=outmembers,proto3" json:"outmembers,omitempty"`
	Links          []*Link          `protobuf:"bytes,11,rep,name=links,proto3" json:"links,omitempty"`
	Attachments    []*Attachment    `protobuf:"bytes,12,rep,name=attachments,proto3" json:"attachments,omitempty"` //(files)
	Number         uint64           `protobuf:"varint,13,opt,name=number,proto3" json:"number,omitempty"`
	CreationDate   string           `protobuf:"bytes,14,opt,name=creationDate,proto3" json:"creationDate,omitempty"`
	Owner          *User            `protobuf:"bytes,15,opt,name=owner,proto3" json:"owner,omitempty"`
	Releases       []*Release       `protobuf:"bytes,16,rep,name=releases,proto3" json:"releases,omitempty"`
	Slots          []*RoomSlot      `protobuf:"bytes,17,rep,name=slots,proto3" json:"slots,omitempty"`
	Equipments     []*RoomEquipment `protobuf:"bytes,18,rep,name=equipments,proto3" json:"equipments,omitempty"`
	Bookings       []*Booking       `protobuf:"bytes,19,rep,name=bookings,proto3" json:"bookings,omitempty"`
	Reports        []*RoomReport    `protobuf:"bytes,20,rep,name=reports,proto3" json:"reports,omitempty"`
	Participants   []*Participant   `protobuf:"bytes,21,rep,name=participants,proto3" json:"participants,omitempty"`
	State          AddressState     `protobuf:"varint,22,opt,name=state,proto3,enum=booking.AddressState" json:"state,omitempty"`
}

func (x *Room) Reset() {
	*x = Room{}
	if protoimpl.UnsafeEnabled {
		mi := &file_room_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Room) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Room) ProtoMessage() {}

func (x *Room) ProtoReflect() protoreflect.Message {
	mi := &file_room_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Room.ProtoReflect.Descriptor instead.
func (*Room) Descriptor() ([]byte, []int) {
	return file_room_proto_rawDescGZIP(), []int{0}
}

func (x *Room) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *Room) GetAuthor() *User {
	if x != nil {
		return x.Author
	}
	return nil
}

func (x *Room) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Room) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Room) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Room) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Room) GetTargetAudience() string {
	if x != nil {
		return x.TargetAudience
	}
	return ""
}

func (x *Room) GetStatus() AddressStatus {
	if x != nil {
		return x.Status
	}
	return AddressStatus_ADDRESS_STATUS
}

func (x *Room) GetAssignee() *User {
	if x != nil {
		return x.Assignee
	}
	return nil
}

func (x *Room) GetOutmembers() []*Outmember {
	if x != nil {
		return x.Outmembers
	}
	return nil
}

func (x *Room) GetLinks() []*Link {
	if x != nil {
		return x.Links
	}
	return nil
}

func (x *Room) GetAttachments() []*Attachment {
	if x != nil {
		return x.Attachments
	}
	return nil
}

func (x *Room) GetNumber() uint64 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *Room) GetCreationDate() string {
	if x != nil {
		return x.CreationDate
	}
	return ""
}

func (x *Room) GetOwner() *User {
	if x != nil {
		return x.Owner
	}
	return nil
}

func (x *Room) GetReleases() []*Release {
	if x != nil {
		return x.Releases
	}
	return nil
}

func (x *Room) GetSlots() []*RoomSlot {
	if x != nil {
		return x.Slots
	}
	return nil
}

func (x *Room) GetEquipments() []*RoomEquipment {
	if x != nil {
		return x.Equipments
	}
	return nil
}

func (x *Room) GetBookings() []*Booking {
	if x != nil {
		return x.Bookings
	}
	return nil
}

func (x *Room) GetReports() []*RoomReport {
	if x != nil {
		return x.Reports
	}
	return nil
}

func (x *Room) GetParticipants() []*Participant {
	if x != nil {
		return x.Participants
	}
	return nil
}

func (x *Room) GetState() AddressState {
	if x != nil {
		return x.State
	}
	return AddressState_UNKNOWN_STATE
}

type Release struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid        string        `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Title       string        `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description string        `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Date        string        `protobuf:"bytes,4,opt,name=date,proto3" json:"date,omitempty"`
	FactSlot    *Notification `protobuf:"bytes,5,opt,name=factSlot,proto3" json:"factSlot,omitempty"`
	RoomID      string        `protobuf:"bytes,6,opt,name=roomID,proto3" json:"roomID,omitempty"`
}

func (x *Release) Reset() {
	*x = Release{}
	if protoimpl.UnsafeEnabled {
		mi := &file_room_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Release) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Release) ProtoMessage() {}

func (x *Release) ProtoReflect() protoreflect.Message {
	mi := &file_room_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Release.ProtoReflect.Descriptor instead.
func (*Release) Descriptor() ([]byte, []int) {
	return file_room_proto_rawDescGZIP(), []int{1}
}

func (x *Release) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *Release) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Release) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Release) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *Release) GetFactSlot() *Notification {
	if x != nil {
		return x.FactSlot
	}
	return nil
}

func (x *Release) GetRoomID() string {
	if x != nil {
		return x.RoomID
	}
	return ""
}

type RoomSlot struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid      string        `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Timeline  *Timeline     `protobuf:"bytes,2,opt,name=timeline,proto3" json:"timeline,omitempty"`
	PlanSlot  *Notification `protobuf:"bytes,3,opt,name=planSlot,proto3" json:"planSlot,omitempty"`
	FactSlot  *Notification `protobuf:"bytes,4,opt,name=factSlot,proto3" json:"factSlot,omitempty"`
	CreatedAt string        `protobuf:"bytes,5,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
}

func (x *RoomSlot) Reset() {
	*x = RoomSlot{}
	if protoimpl.UnsafeEnabled {
		mi := &file_room_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomSlot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomSlot) ProtoMessage() {}

func (x *RoomSlot) ProtoReflect() protoreflect.Message {
	mi := &file_room_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomSlot.ProtoReflect.Descriptor instead.
func (*RoomSlot) Descriptor() ([]byte, []int) {
	return file_room_proto_rawDescGZIP(), []int{2}
}

func (x *RoomSlot) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *RoomSlot) GetTimeline() *Timeline {
	if x != nil {
		return x.Timeline
	}
	return nil
}

func (x *RoomSlot) GetPlanSlot() *Notification {
	if x != nil {
		return x.PlanSlot
	}
	return nil
}

func (x *RoomSlot) GetFactSlot() *Notification {
	if x != nil {
		return x.FactSlot
	}
	return nil
}

func (x *RoomSlot) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type RoomEquipment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid        string    `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Title       string    `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Timeline    *Timeline `protobuf:"bytes,3,opt,name=timeline,proto3" json:"timeline,omitempty"`
	Description string    `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	PlanValue   float32   `protobuf:"fixed32,5,opt,name=planValue,proto3" json:"planValue,omitempty"`
	FactValue   float32   `protobuf:"fixed32,6,opt,name=factValue,proto3" json:"factValue,omitempty"`
	CreatedAt   string    `protobuf:"bytes,7,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
}

func (x *RoomEquipment) Reset() {
	*x = RoomEquipment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_room_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomEquipment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomEquipment) ProtoMessage() {}

func (x *RoomEquipment) ProtoReflect() protoreflect.Message {
	mi := &file_room_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomEquipment.ProtoReflect.Descriptor instead.
func (*RoomEquipment) Descriptor() ([]byte, []int) {
	return file_room_proto_rawDescGZIP(), []int{3}
}

func (x *RoomEquipment) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *RoomEquipment) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *RoomEquipment) GetTimeline() *Timeline {
	if x != nil {
		return x.Timeline
	}
	return nil
}

func (x *RoomEquipment) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *RoomEquipment) GetPlanValue() float32 {
	if x != nil {
		return x.PlanValue
	}
	return 0
}

func (x *RoomEquipment) GetFactValue() float32 {
	if x != nil {
		return x.FactValue
	}
	return 0
}

func (x *RoomEquipment) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type Slot struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timeline *Timeline     `protobuf:"bytes,1,opt,name=timeline,proto3" json:"timeline,omitempty"`
	PlanSlot *Notification `protobuf:"bytes,2,opt,name=planSlot,proto3" json:"planSlot,omitempty"`
	FactSlot *Notification `protobuf:"bytes,3,opt,name=factSlot,proto3" json:"factSlot,omitempty"`
}

func (x *Slot) Reset() {
	*x = Slot{}
	if protoimpl.UnsafeEnabled {
		mi := &file_room_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Slot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Slot) ProtoMessage() {}

func (x *Slot) ProtoReflect() protoreflect.Message {
	mi := &file_room_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Slot.ProtoReflect.Descriptor instead.
func (*Slot) Descriptor() ([]byte, []int) {
	return file_room_proto_rawDescGZIP(), []int{4}
}

func (x *Slot) GetTimeline() *Timeline {
	if x != nil {
		return x.Timeline
	}
	return nil
}

func (x *Slot) GetPlanSlot() *Notification {
	if x != nil {
		return x.PlanSlot
	}
	return nil
}

func (x *Slot) GetFactSlot() *Notification {
	if x != nil {
		return x.FactSlot
	}
	return nil
}

type Equipment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       string    `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Timeline    *Timeline `protobuf:"bytes,2,opt,name=timeline,proto3" json:"timeline,omitempty"`
	Description string    `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	PlanValue   float32   `protobuf:"fixed32,4,opt,name=planValue,proto3" json:"planValue,omitempty"`
	FactValue   float32   `protobuf:"fixed32,5,opt,name=factValue,proto3" json:"factValue,omitempty"`
}

func (x *Equipment) Reset() {
	*x = Equipment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_room_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Equipment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Equipment) ProtoMessage() {}

func (x *Equipment) ProtoReflect() protoreflect.Message {
	mi := &file_room_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Equipment.ProtoReflect.Descriptor instead.
func (*Equipment) Descriptor() ([]byte, []int) {
	return file_room_proto_rawDescGZIP(), []int{5}
}

func (x *Equipment) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Equipment) GetTimeline() *Timeline {
	if x != nil {
		return x.Timeline
	}
	return nil
}

func (x *Equipment) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Equipment) GetPlanValue() float32 {
	if x != nil {
		return x.PlanValue
	}
	return 0
}

func (x *Equipment) GetFactValue() float32 {
	if x != nil {
		return x.FactValue
	}
	return 0
}

type RoomReport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timeline   *Timeline    `protobuf:"bytes,1,opt,name=timeline,proto3" json:"timeline,omitempty"`
	Comment    string       `protobuf:"bytes,2,opt,name=comment,proto3" json:"comment,omitempty"`
	Uuid       string       `protobuf:"bytes,3,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Releases   []*Release   `protobuf:"bytes,4,rep,name=releases,proto3" json:"releases,omitempty"`
	Slots      []*Slot      `protobuf:"bytes,5,rep,name=slots,proto3" json:"slots,omitempty"`
	Equipments []*Equipment `protobuf:"bytes,6,rep,name=equipments,proto3" json:"equipments,omitempty"`
}

func (x *RoomReport) Reset() {
	*x = RoomReport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_room_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomReport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomReport) ProtoMessage() {}

func (x *RoomReport) ProtoReflect() protoreflect.Message {
	mi := &file_room_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomReport.ProtoReflect.Descriptor instead.
func (*RoomReport) Descriptor() ([]byte, []int) {
	return file_room_proto_rawDescGZIP(), []int{6}
}

func (x *RoomReport) GetTimeline() *Timeline {
	if x != nil {
		return x.Timeline
	}
	return nil
}

func (x *RoomReport) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

func (x *RoomReport) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *RoomReport) GetReleases() []*Release {
	if x != nil {
		return x.Releases
	}
	return nil
}

func (x *RoomReport) GetSlots() []*Slot {
	if x != nil {
		return x.Slots
	}
	return nil
}

func (x *RoomReport) GetEquipments() []*Equipment {
	if x != nil {
		return x.Equipments
	}
	return nil
}

var File_room_proto protoreflect.FileDescriptor

var file_room_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x62, 0x6f,
	0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x1a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x18, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x6f, 0x75, 0x74, 0x6d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x62, 0x6f, 0x6f,
	0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfc, 0x06, 0x0a, 0x04, 0x52,
	0x6f, 0x6f, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x25, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e,
	0x67, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x41, 0x75, 0x64, 0x69,
	0x65, 0x6e, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x41, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x2e, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x62, 0x6f, 0x6f,
	0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x29, 0x0a, 0x08, 0x61, 0x73,
	0x73, 0x69, 0x67, 0x6e, 0x65, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x62,
	0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x08, 0x61, 0x73, 0x73,
	0x69, 0x67, 0x6e, 0x65, 0x65, 0x12, 0x32, 0x0a, 0x0a, 0x6f, 0x75, 0x74, 0x6d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x62, 0x6f, 0x6f, 0x6b,
	0x69, 0x6e, 0x67, 0x2e, 0x4f, 0x75, 0x74, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x0a, 0x6f,
	0x75, 0x74, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x12, 0x23, 0x0a, 0x05, 0x6c, 0x69, 0x6e,
	0x6b, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69,
	0x6e, 0x67, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x05, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x12, 0x35,
	0x0a, 0x0b, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x0c, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x41, 0x74,
	0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0b, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x22, 0x0a,
	0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74,
	0x65, 0x12, 0x23, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x2c, 0x0a, 0x08, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73,
	0x65, 0x73, 0x18, 0x10, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69,
	0x6e, 0x67, 0x2e, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x52, 0x08, 0x72, 0x65, 0x6c, 0x65,
	0x61, 0x73, 0x65, 0x73, 0x12, 0x27, 0x0a, 0x05, 0x73, 0x6c, 0x6f, 0x74, 0x73, 0x18, 0x11, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x6f,
	0x6f, 0x6d, 0x53, 0x6c, 0x6f, 0x74, 0x52, 0x05, 0x73, 0x6c, 0x6f, 0x74, 0x73, 0x12, 0x36, 0x0a,
	0x0a, 0x65, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x12, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x6f, 0x6f, 0x6d,
	0x45, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0a, 0x65, 0x71, 0x75, 0x69, 0x70,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x2c, 0x0a, 0x08, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67,
	0x73, 0x18, 0x13, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e,
	0x67, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x08, 0x62, 0x6f, 0x6f, 0x6b, 0x69,
	0x6e, 0x67, 0x73, 0x12, 0x2d, 0x0a, 0x07, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x18, 0x14,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x52,
	0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x07, 0x72, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x73, 0x12, 0x38, 0x0a, 0x0c, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e,
	0x74, 0x73, 0x18, 0x15, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69,
	0x6e, 0x67, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x52, 0x0c,
	0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x73, 0x12, 0x2b, 0x0a, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x16, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x62, 0x6f,
	0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0xb4, 0x01, 0x0a, 0x07, 0x52, 0x65,
	0x6c, 0x65, 0x61, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x31, 0x0a, 0x08, 0x66, 0x61, 0x63, 0x74, 0x53, 0x6c, 0x6f,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e,
	0x67, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08,
	0x66, 0x61, 0x63, 0x74, 0x53, 0x6c, 0x6f, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x6f, 0x6f, 0x6d,
	0x49, 0x44, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x44,
	0x22, 0xd1, 0x01, 0x0a, 0x08, 0x52, 0x6f, 0x6f, 0x6d, 0x53, 0x6c, 0x6f, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69,
	0x64, 0x12, 0x2d, 0x0a, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65,
	0x12, 0x31, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x6e, 0x53, 0x6c, 0x6f, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x6e, 0x53,
	0x6c, 0x6f, 0x74, 0x12, 0x31, 0x0a, 0x08, 0x66, 0x61, 0x63, 0x74, 0x53, 0x6c, 0x6f, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x66, 0x61,
	0x63, 0x74, 0x53, 0x6c, 0x6f, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x22, 0xe4, 0x01, 0x0a, 0x0d, 0x52, 0x6f, 0x6f, 0x6d, 0x45, 0x71, 0x75,
	0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x2d, 0x0a, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x70, 0x6c, 0x61, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x66, 0x61, 0x63, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x09, 0x66, 0x61, 0x63, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x9b, 0x01, 0x0a, 0x04,
	0x53, 0x6c, 0x6f, 0x74, 0x12, 0x2d, 0x0a, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x6c,
	0x69, 0x6e, 0x65, 0x12, 0x31, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x6e, 0x53, 0x6c, 0x6f, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x70, 0x6c,
	0x61, 0x6e, 0x53, 0x6c, 0x6f, 0x74, 0x12, 0x31, 0x0a, 0x08, 0x66, 0x61, 0x63, 0x74, 0x53, 0x6c,
	0x6f, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69,
	0x6e, 0x67, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x08, 0x66, 0x61, 0x63, 0x74, 0x53, 0x6c, 0x6f, 0x74, 0x22, 0xae, 0x01, 0x0a, 0x09, 0x45, 0x71,
	0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x2d, 0x0a,
	0x08, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x6c, 0x69,
	0x6e, 0x65, 0x52, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c,
	0x0a, 0x09, 0x70, 0x6c, 0x61, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x09, 0x70, 0x6c, 0x61, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x66, 0x61, 0x63, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x09, 0x66, 0x61, 0x63, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xf0, 0x01, 0x0a, 0x0a, 0x52,
	0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x2d, 0x0a, 0x08, 0x74, 0x69, 0x6d,
	0x65, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x62, 0x6f,
	0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x08,
	0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x2c, 0x0a, 0x08, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73,
	0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69,
	0x6e, 0x67, 0x2e, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x52, 0x08, 0x72, 0x65, 0x6c, 0x65,
	0x61, 0x73, 0x65, 0x73, 0x12, 0x23, 0x0a, 0x05, 0x73, 0x6c, 0x6f, 0x74, 0x73, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x6c,
	0x6f, 0x74, 0x52, 0x05, 0x73, 0x6c, 0x6f, 0x74, 0x73, 0x12, 0x32, 0x0a, 0x0a, 0x65, 0x71, 0x75,
	0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x45, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x0a, 0x65, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x0a, 0x5a,
	0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_room_proto_rawDescOnce sync.Once
	file_room_proto_rawDescData = file_room_proto_rawDesc
)

func file_room_proto_rawDescGZIP() []byte {
	file_room_proto_rawDescOnce.Do(func() {
		file_room_proto_rawDescData = protoimpl.X.CompressGZIP(file_room_proto_rawDescData)
	})
	return file_room_proto_rawDescData
}

var file_room_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_room_proto_goTypes = []interface{}{
	(*Room)(nil),          // 0: booking.Room
	(*Release)(nil),       // 1: booking.Release
	(*RoomSlot)(nil),      // 2: booking.RoomSlot
	(*RoomEquipment)(nil), // 3: booking.RoomEquipment
	(*Slot)(nil),          // 4: booking.Slot
	(*Equipment)(nil),     // 5: booking.Equipment
	(*RoomReport)(nil),    // 6: booking.RoomReport
	(*User)(nil),          // 7: booking.User
	(AddressStatus)(0),    // 8: booking.AddressStatus
	(*Outmember)(nil),     // 9: booking.Outmember
	(*Link)(nil),          // 10: booking.Link
	(*Attachment)(nil),    // 11: booking.Attachment
	(*Booking)(nil),       // 12: booking.Booking
	(*Participant)(nil),   // 13: booking.Participant
	(AddressState)(0),     // 14: booking.AddressState
	(*Notification)(nil),  // 15: booking.Notification
	(*Timeline)(nil),      // 16: booking.Timeline
}
var file_room_proto_depIdxs = []int32{
	7,  // 0: booking.Room.author:type_name -> booking.User
	8,  // 1: booking.Room.status:type_name -> booking.AddressStatus
	7,  // 2: booking.Room.assignee:type_name -> booking.User
	9,  // 3: booking.Room.outmembers:type_name -> booking.Outmember
	10, // 4: booking.Room.links:type_name -> booking.Link
	11, // 5: booking.Room.attachments:type_name -> booking.Attachment
	7,  // 6: booking.Room.owner:type_name -> booking.User
	1,  // 7: booking.Room.releases:type_name -> booking.Release
	2,  // 8: booking.Room.slots:type_name -> booking.RoomSlot
	3,  // 9: booking.Room.equipments:type_name -> booking.RoomEquipment
	12, // 10: booking.Room.bookings:type_name -> booking.Booking
	6,  // 11: booking.Room.reports:type_name -> booking.RoomReport
	13, // 12: booking.Room.participants:type_name -> booking.Participant
	14, // 13: booking.Room.state:type_name -> booking.AddressState
	15, // 14: booking.Release.factSlot:type_name -> booking.Notification
	16, // 15: booking.RoomSlot.timeline:type_name -> booking.Timeline
	15, // 16: booking.RoomSlot.planSlot:type_name -> booking.Notification
	15, // 17: booking.RoomSlot.factSlot:type_name -> booking.Notification
	16, // 18: booking.RoomEquipment.timeline:type_name -> booking.Timeline
	16, // 19: booking.Slot.timeline:type_name -> booking.Timeline
	15, // 20: booking.Slot.planSlot:type_name -> booking.Notification
	15, // 21: booking.Slot.factSlot:type_name -> booking.Notification
	16, // 22: booking.Equipment.timeline:type_name -> booking.Timeline
	16, // 23: booking.RoomReport.timeline:type_name -> booking.Timeline
	1,  // 24: booking.RoomReport.releases:type_name -> booking.Release
	4,  // 25: booking.RoomReport.slots:type_name -> booking.Slot
	5,  // 26: booking.RoomReport.equipments:type_name -> booking.Equipment
	27, // [27:27] is the sub-list for method output_type
	27, // [27:27] is the sub-list for method input_type
	27, // [27:27] is the sub-list for extension type_name
	27, // [27:27] is the sub-list for extension extendee
	0,  // [0:27] is the sub-list for field type_name
}

func init() { file_room_proto_init() }
func file_room_proto_init() {
	if File_room_proto != nil {
		return
	}
	file_user_proto_init()
	file_dictionary_service_proto_init()
	file_address_proto_init()
	file_outmember_proto_init()
	file_booking_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_room_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Room); i {
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
		file_room_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Release); i {
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
		file_room_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomSlot); i {
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
		file_room_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomEquipment); i {
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
		file_room_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Slot); i {
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
		file_room_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Equipment); i {
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
		file_room_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomReport); i {
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
			RawDescriptor: file_room_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_room_proto_goTypes,
		DependencyIndexes: file_room_proto_depIdxs,
		MessageInfos:      file_room_proto_msgTypes,
	}.Build()
	File_room_proto = out.File
	file_room_proto_rawDesc = nil
	file_room_proto_goTypes = nil
	file_room_proto_depIdxs = nil
}
