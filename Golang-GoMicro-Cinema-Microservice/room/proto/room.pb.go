// Code generated by protoc-gen-go. DO NOT EDIT.
// source: room/proto/room.proto

package room

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type AddRoomRequest struct {
	RoomName             string   `protobuf:"bytes,1,opt,name=roomName,proto3" json:"roomName,omitempty"`
	IsRoomAvailable      bool     `protobuf:"varint,2,opt,name=isRoomAvailable,proto3" json:"isRoomAvailable,omitempty"`
	AvailableSeats       int32    `protobuf:"varint,3,opt,name=availableSeats,proto3" json:"availableSeats,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddRoomRequest) Reset()         { *m = AddRoomRequest{} }
func (m *AddRoomRequest) String() string { return proto.CompactTextString(m) }
func (*AddRoomRequest) ProtoMessage()    {}
func (*AddRoomRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3858eaf69540bf01, []int{0}
}

func (m *AddRoomRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddRoomRequest.Unmarshal(m, b)
}
func (m *AddRoomRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddRoomRequest.Marshal(b, m, deterministic)
}
func (m *AddRoomRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddRoomRequest.Merge(m, src)
}
func (m *AddRoomRequest) XXX_Size() int {
	return xxx_messageInfo_AddRoomRequest.Size(m)
}
func (m *AddRoomRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddRoomRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddRoomRequest proto.InternalMessageInfo

func (m *AddRoomRequest) GetRoomName() string {
	if m != nil {
		return m.RoomName
	}
	return ""
}

func (m *AddRoomRequest) GetIsRoomAvailable() bool {
	if m != nil {
		return m.IsRoomAvailable
	}
	return false
}

func (m *AddRoomRequest) GetAvailableSeats() int32 {
	if m != nil {
		return m.AvailableSeats
	}
	return 0
}

type AddRoomResponse struct {
	Message              string      `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Room                 *RoomStruct `protobuf:"bytes,2,opt,name=room,proto3" json:"room,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *AddRoomResponse) Reset()         { *m = AddRoomResponse{} }
func (m *AddRoomResponse) String() string { return proto.CompactTextString(m) }
func (*AddRoomResponse) ProtoMessage()    {}
func (*AddRoomResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3858eaf69540bf01, []int{1}
}

func (m *AddRoomResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddRoomResponse.Unmarshal(m, b)
}
func (m *AddRoomResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddRoomResponse.Marshal(b, m, deterministic)
}
func (m *AddRoomResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddRoomResponse.Merge(m, src)
}
func (m *AddRoomResponse) XXX_Size() int {
	return xxx_messageInfo_AddRoomResponse.Size(m)
}
func (m *AddRoomResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddRoomResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddRoomResponse proto.InternalMessageInfo

func (m *AddRoomResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *AddRoomResponse) GetRoom() *RoomStruct {
	if m != nil {
		return m.Room
	}
	return nil
}

type RemoveRoomByIDRequest struct {
	RoomName             string   `protobuf:"bytes,1,opt,name=roomName,proto3" json:"roomName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveRoomByIDRequest) Reset()         { *m = RemoveRoomByIDRequest{} }
func (m *RemoveRoomByIDRequest) String() string { return proto.CompactTextString(m) }
func (*RemoveRoomByIDRequest) ProtoMessage()    {}
func (*RemoveRoomByIDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3858eaf69540bf01, []int{2}
}

func (m *RemoveRoomByIDRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveRoomByIDRequest.Unmarshal(m, b)
}
func (m *RemoveRoomByIDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveRoomByIDRequest.Marshal(b, m, deterministic)
}
func (m *RemoveRoomByIDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveRoomByIDRequest.Merge(m, src)
}
func (m *RemoveRoomByIDRequest) XXX_Size() int {
	return xxx_messageInfo_RemoveRoomByIDRequest.Size(m)
}
func (m *RemoveRoomByIDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveRoomByIDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveRoomByIDRequest proto.InternalMessageInfo

func (m *RemoveRoomByIDRequest) GetRoomName() string {
	if m != nil {
		return m.RoomName
	}
	return ""
}

type RemoveRoomByIDResponse struct {
	Message              string      `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Room                 *RoomStruct `protobuf:"bytes,2,opt,name=room,proto3" json:"room,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *RemoveRoomByIDResponse) Reset()         { *m = RemoveRoomByIDResponse{} }
func (m *RemoveRoomByIDResponse) String() string { return proto.CompactTextString(m) }
func (*RemoveRoomByIDResponse) ProtoMessage()    {}
func (*RemoveRoomByIDResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3858eaf69540bf01, []int{3}
}

func (m *RemoveRoomByIDResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveRoomByIDResponse.Unmarshal(m, b)
}
func (m *RemoveRoomByIDResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveRoomByIDResponse.Marshal(b, m, deterministic)
}
func (m *RemoveRoomByIDResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveRoomByIDResponse.Merge(m, src)
}
func (m *RemoveRoomByIDResponse) XXX_Size() int {
	return xxx_messageInfo_RemoveRoomByIDResponse.Size(m)
}
func (m *RemoveRoomByIDResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveRoomByIDResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveRoomByIDResponse proto.InternalMessageInfo

func (m *RemoveRoomByIDResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *RemoveRoomByIDResponse) GetRoom() *RoomStruct {
	if m != nil {
		return m.Room
	}
	return nil
}

type ShowAllRoomsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShowAllRoomsRequest) Reset()         { *m = ShowAllRoomsRequest{} }
func (m *ShowAllRoomsRequest) String() string { return proto.CompactTextString(m) }
func (*ShowAllRoomsRequest) ProtoMessage()    {}
func (*ShowAllRoomsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3858eaf69540bf01, []int{4}
}

func (m *ShowAllRoomsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShowAllRoomsRequest.Unmarshal(m, b)
}
func (m *ShowAllRoomsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShowAllRoomsRequest.Marshal(b, m, deterministic)
}
func (m *ShowAllRoomsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShowAllRoomsRequest.Merge(m, src)
}
func (m *ShowAllRoomsRequest) XXX_Size() int {
	return xxx_messageInfo_ShowAllRoomsRequest.Size(m)
}
func (m *ShowAllRoomsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ShowAllRoomsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ShowAllRoomsRequest proto.InternalMessageInfo

type ShowAllRoomsResponse struct {
	Rooms                []*RoomStruct `protobuf:"bytes,1,rep,name=rooms,proto3" json:"rooms,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ShowAllRoomsResponse) Reset()         { *m = ShowAllRoomsResponse{} }
func (m *ShowAllRoomsResponse) String() string { return proto.CompactTextString(m) }
func (*ShowAllRoomsResponse) ProtoMessage()    {}
func (*ShowAllRoomsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3858eaf69540bf01, []int{5}
}

func (m *ShowAllRoomsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShowAllRoomsResponse.Unmarshal(m, b)
}
func (m *ShowAllRoomsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShowAllRoomsResponse.Marshal(b, m, deterministic)
}
func (m *ShowAllRoomsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShowAllRoomsResponse.Merge(m, src)
}
func (m *ShowAllRoomsResponse) XXX_Size() int {
	return xxx_messageInfo_ShowAllRoomsResponse.Size(m)
}
func (m *ShowAllRoomsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ShowAllRoomsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ShowAllRoomsResponse proto.InternalMessageInfo

func (m *ShowAllRoomsResponse) GetRooms() []*RoomStruct {
	if m != nil {
		return m.Rooms
	}
	return nil
}

type FindRoomByIDRequest struct {
	RoomName             string   `protobuf:"bytes,1,opt,name=roomName,proto3" json:"roomName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindRoomByIDRequest) Reset()         { *m = FindRoomByIDRequest{} }
func (m *FindRoomByIDRequest) String() string { return proto.CompactTextString(m) }
func (*FindRoomByIDRequest) ProtoMessage()    {}
func (*FindRoomByIDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3858eaf69540bf01, []int{6}
}

func (m *FindRoomByIDRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindRoomByIDRequest.Unmarshal(m, b)
}
func (m *FindRoomByIDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindRoomByIDRequest.Marshal(b, m, deterministic)
}
func (m *FindRoomByIDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindRoomByIDRequest.Merge(m, src)
}
func (m *FindRoomByIDRequest) XXX_Size() int {
	return xxx_messageInfo_FindRoomByIDRequest.Size(m)
}
func (m *FindRoomByIDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindRoomByIDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindRoomByIDRequest proto.InternalMessageInfo

func (m *FindRoomByIDRequest) GetRoomName() string {
	if m != nil {
		return m.RoomName
	}
	return ""
}

type FindRoomByIDResponse struct {
	Message              string      `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Room                 *RoomStruct `protobuf:"bytes,2,opt,name=room,proto3" json:"room,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *FindRoomByIDResponse) Reset()         { *m = FindRoomByIDResponse{} }
func (m *FindRoomByIDResponse) String() string { return proto.CompactTextString(m) }
func (*FindRoomByIDResponse) ProtoMessage()    {}
func (*FindRoomByIDResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3858eaf69540bf01, []int{7}
}

func (m *FindRoomByIDResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindRoomByIDResponse.Unmarshal(m, b)
}
func (m *FindRoomByIDResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindRoomByIDResponse.Marshal(b, m, deterministic)
}
func (m *FindRoomByIDResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindRoomByIDResponse.Merge(m, src)
}
func (m *FindRoomByIDResponse) XXX_Size() int {
	return xxx_messageInfo_FindRoomByIDResponse.Size(m)
}
func (m *FindRoomByIDResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FindRoomByIDResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FindRoomByIDResponse proto.InternalMessageInfo

func (m *FindRoomByIDResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *FindRoomByIDResponse) GetRoom() *RoomStruct {
	if m != nil {
		return m.Room
	}
	return nil
}

type ShowAllShowingsInRoomRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShowAllShowingsInRoomRequest) Reset()         { *m = ShowAllShowingsInRoomRequest{} }
func (m *ShowAllShowingsInRoomRequest) String() string { return proto.CompactTextString(m) }
func (*ShowAllShowingsInRoomRequest) ProtoMessage()    {}
func (*ShowAllShowingsInRoomRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3858eaf69540bf01, []int{8}
}

func (m *ShowAllShowingsInRoomRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShowAllShowingsInRoomRequest.Unmarshal(m, b)
}
func (m *ShowAllShowingsInRoomRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShowAllShowingsInRoomRequest.Marshal(b, m, deterministic)
}
func (m *ShowAllShowingsInRoomRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShowAllShowingsInRoomRequest.Merge(m, src)
}
func (m *ShowAllShowingsInRoomRequest) XXX_Size() int {
	return xxx_messageInfo_ShowAllShowingsInRoomRequest.Size(m)
}
func (m *ShowAllShowingsInRoomRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ShowAllShowingsInRoomRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ShowAllShowingsInRoomRequest proto.InternalMessageInfo

type ShowAllShowingsInRoomResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShowAllShowingsInRoomResponse) Reset()         { *m = ShowAllShowingsInRoomResponse{} }
func (m *ShowAllShowingsInRoomResponse) String() string { return proto.CompactTextString(m) }
func (*ShowAllShowingsInRoomResponse) ProtoMessage()    {}
func (*ShowAllShowingsInRoomResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3858eaf69540bf01, []int{9}
}

func (m *ShowAllShowingsInRoomResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShowAllShowingsInRoomResponse.Unmarshal(m, b)
}
func (m *ShowAllShowingsInRoomResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShowAllShowingsInRoomResponse.Marshal(b, m, deterministic)
}
func (m *ShowAllShowingsInRoomResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShowAllShowingsInRoomResponse.Merge(m, src)
}
func (m *ShowAllShowingsInRoomResponse) XXX_Size() int {
	return xxx_messageInfo_ShowAllShowingsInRoomResponse.Size(m)
}
func (m *ShowAllShowingsInRoomResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ShowAllShowingsInRoomResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ShowAllShowingsInRoomResponse proto.InternalMessageInfo

type SetRoomAvailabilityRequest struct {
	RoomName             string   `protobuf:"bytes,1,opt,name=roomName,proto3" json:"roomName,omitempty"`
	IsRoomAvailable      bool     `protobuf:"varint,2,opt,name=isRoomAvailable,proto3" json:"isRoomAvailable,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetRoomAvailabilityRequest) Reset()         { *m = SetRoomAvailabilityRequest{} }
func (m *SetRoomAvailabilityRequest) String() string { return proto.CompactTextString(m) }
func (*SetRoomAvailabilityRequest) ProtoMessage()    {}
func (*SetRoomAvailabilityRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3858eaf69540bf01, []int{10}
}

func (m *SetRoomAvailabilityRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetRoomAvailabilityRequest.Unmarshal(m, b)
}
func (m *SetRoomAvailabilityRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetRoomAvailabilityRequest.Marshal(b, m, deterministic)
}
func (m *SetRoomAvailabilityRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetRoomAvailabilityRequest.Merge(m, src)
}
func (m *SetRoomAvailabilityRequest) XXX_Size() int {
	return xxx_messageInfo_SetRoomAvailabilityRequest.Size(m)
}
func (m *SetRoomAvailabilityRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetRoomAvailabilityRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetRoomAvailabilityRequest proto.InternalMessageInfo

func (m *SetRoomAvailabilityRequest) GetRoomName() string {
	if m != nil {
		return m.RoomName
	}
	return ""
}

func (m *SetRoomAvailabilityRequest) GetIsRoomAvailable() bool {
	if m != nil {
		return m.IsRoomAvailable
	}
	return false
}

type SetRoomAvailabilityResponse struct {
	Message              string      `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Room                 *RoomStruct `protobuf:"bytes,2,opt,name=room,proto3" json:"room,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *SetRoomAvailabilityResponse) Reset()         { *m = SetRoomAvailabilityResponse{} }
func (m *SetRoomAvailabilityResponse) String() string { return proto.CompactTextString(m) }
func (*SetRoomAvailabilityResponse) ProtoMessage()    {}
func (*SetRoomAvailabilityResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3858eaf69540bf01, []int{11}
}

func (m *SetRoomAvailabilityResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetRoomAvailabilityResponse.Unmarshal(m, b)
}
func (m *SetRoomAvailabilityResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetRoomAvailabilityResponse.Marshal(b, m, deterministic)
}
func (m *SetRoomAvailabilityResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetRoomAvailabilityResponse.Merge(m, src)
}
func (m *SetRoomAvailabilityResponse) XXX_Size() int {
	return xxx_messageInfo_SetRoomAvailabilityResponse.Size(m)
}
func (m *SetRoomAvailabilityResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SetRoomAvailabilityResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SetRoomAvailabilityResponse proto.InternalMessageInfo

func (m *SetRoomAvailabilityResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *SetRoomAvailabilityResponse) GetRoom() *RoomStruct {
	if m != nil {
		return m.Room
	}
	return nil
}

type RoomStruct struct {
	RoomName             string   `protobuf:"bytes,1,opt,name=roomName,proto3" json:"roomName,omitempty"`
	IsRoomAvailable      bool     `protobuf:"varint,2,opt,name=isRoomAvailable,proto3" json:"isRoomAvailable,omitempty"`
	AvailableSeats       int32    `protobuf:"varint,3,opt,name=availableSeats,proto3" json:"availableSeats,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoomStruct) Reset()         { *m = RoomStruct{} }
func (m *RoomStruct) String() string { return proto.CompactTextString(m) }
func (*RoomStruct) ProtoMessage()    {}
func (*RoomStruct) Descriptor() ([]byte, []int) {
	return fileDescriptor_3858eaf69540bf01, []int{12}
}

func (m *RoomStruct) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoomStruct.Unmarshal(m, b)
}
func (m *RoomStruct) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoomStruct.Marshal(b, m, deterministic)
}
func (m *RoomStruct) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoomStruct.Merge(m, src)
}
func (m *RoomStruct) XXX_Size() int {
	return xxx_messageInfo_RoomStruct.Size(m)
}
func (m *RoomStruct) XXX_DiscardUnknown() {
	xxx_messageInfo_RoomStruct.DiscardUnknown(m)
}

var xxx_messageInfo_RoomStruct proto.InternalMessageInfo

func (m *RoomStruct) GetRoomName() string {
	if m != nil {
		return m.RoomName
	}
	return ""
}

func (m *RoomStruct) GetIsRoomAvailable() bool {
	if m != nil {
		return m.IsRoomAvailable
	}
	return false
}

func (m *RoomStruct) GetAvailableSeats() int32 {
	if m != nil {
		return m.AvailableSeats
	}
	return 0
}

func init() {
	proto.RegisterType((*AddRoomRequest)(nil), "AddRoomRequest")
	proto.RegisterType((*AddRoomResponse)(nil), "AddRoomResponse")
	proto.RegisterType((*RemoveRoomByIDRequest)(nil), "RemoveRoomByIDRequest")
	proto.RegisterType((*RemoveRoomByIDResponse)(nil), "RemoveRoomByIDResponse")
	proto.RegisterType((*ShowAllRoomsRequest)(nil), "ShowAllRoomsRequest")
	proto.RegisterType((*ShowAllRoomsResponse)(nil), "ShowAllRoomsResponse")
	proto.RegisterType((*FindRoomByIDRequest)(nil), "FindRoomByIDRequest")
	proto.RegisterType((*FindRoomByIDResponse)(nil), "FindRoomByIDResponse")
	proto.RegisterType((*ShowAllShowingsInRoomRequest)(nil), "ShowAllShowingsInRoomRequest")
	proto.RegisterType((*ShowAllShowingsInRoomResponse)(nil), "ShowAllShowingsInRoomResponse")
	proto.RegisterType((*SetRoomAvailabilityRequest)(nil), "SetRoomAvailabilityRequest")
	proto.RegisterType((*SetRoomAvailabilityResponse)(nil), "SetRoomAvailabilityResponse")
	proto.RegisterType((*RoomStruct)(nil), "RoomStruct")
}

func init() { proto.RegisterFile("room/proto/room.proto", fileDescriptor_3858eaf69540bf01) }

var fileDescriptor_3858eaf69540bf01 = []byte{
	// 410 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x94, 0x4f, 0x6f, 0xda, 0x40,
	0x10, 0xc5, 0x6d, 0xfe, 0x14, 0x3a, 0x54, 0x50, 0x2d, 0x36, 0xb5, 0x0c, 0x2d, 0xee, 0x1e, 0x2a,
	0x9f, 0x16, 0x15, 0x4e, 0x3d, 0xf4, 0x40, 0x5b, 0x55, 0x42, 0xaa, 0x2a, 0x75, 0x7d, 0xe9, 0xd5,
	0x84, 0x15, 0xb1, 0x64, 0x7b, 0x09, 0x6b, 0x88, 0xc8, 0x21, 0xdf, 0x30, 0xdf, 0x29, 0x5a, 0x63,
	0x88, 0x6d, 0x2d, 0x51, 0x22, 0xa2, 0x9c, 0xd8, 0x9d, 0xe1, 0xf9, 0xcd, 0xc8, 0xbf, 0x67, 0x30,
	0xd7, 0x9c, 0x47, 0xa3, 0xd5, 0x9a, 0x27, 0x7c, 0x24, 0x8f, 0x24, 0x3d, 0xe2, 0x5b, 0x68, 0x4f,
	0x17, 0x0b, 0xca, 0x79, 0x44, 0xd9, 0xd5, 0x86, 0x89, 0x04, 0xd9, 0xd0, 0x94, 0xfd, 0xbf, 0x7e,
	0xc4, 0x2c, 0xdd, 0xd1, 0xdd, 0xb7, 0xf4, 0x78, 0x47, 0x2e, 0x74, 0x02, 0x21, 0xff, 0x3c, 0xdd,
	0xfa, 0x41, 0xe8, 0xcf, 0x43, 0x66, 0x55, 0x1c, 0xdd, 0x6d, 0xd2, 0x72, 0x19, 0x7d, 0x81, 0xb6,
	0x7f, 0xb8, 0x78, 0xcc, 0x4f, 0x84, 0x55, 0x75, 0x74, 0xb7, 0x4e, 0x4b, 0x55, 0xfc, 0x07, 0x3a,
	0x47, 0x7f, 0xb1, 0xe2, 0xb1, 0x60, 0xc8, 0x82, 0x46, 0xc4, 0x84, 0xf0, 0x97, 0x07, 0xff, 0xc3,
	0x15, 0x0d, 0xa1, 0x26, 0x47, 0x49, 0x3d, 0x5b, 0xe3, 0x16, 0x91, 0x32, 0x2f, 0x59, 0x6f, 0x2e,
	0x12, 0x9a, 0x36, 0xf0, 0x04, 0x4c, 0xca, 0x22, 0xbe, 0x65, 0xb2, 0xf3, 0x63, 0x37, 0xfb, 0xf5,
	0x84, 0xa5, 0xb0, 0x07, 0xbd, 0xb2, 0xe8, 0xfc, 0x49, 0x4c, 0xe8, 0x7a, 0x97, 0xfc, 0x7a, 0x1a,
	0x86, 0xb2, 0x25, 0xb2, 0x39, 0xf0, 0x37, 0x30, 0x8a, 0xe5, 0xcc, 0xe9, 0x33, 0xd4, 0xa5, 0x4c,
	0x58, 0xba, 0x53, 0x2d, 0x3f, 0x70, 0xdf, 0xc1, 0x5f, 0xa1, 0xfb, 0x3b, 0x88, 0x17, 0xcf, 0xd9,
	0xec, 0x1f, 0x18, 0x45, 0xc9, 0xf9, 0x7b, 0x7d, 0x82, 0x41, 0xb6, 0x80, 0xfc, 0x09, 0xe2, 0xa5,
	0x98, 0xc5, 0x39, 0x7a, 0xf0, 0x10, 0x3e, 0x9e, 0xe8, 0xef, 0xbd, 0xf1, 0x1c, 0x6c, 0x8f, 0x25,
	0x39, 0x58, 0x82, 0x30, 0x48, 0x76, 0x2f, 0x0a, 0x1f, 0xfe, 0x0f, 0x7d, 0xa5, 0xc7, 0xf9, 0xeb,
	0xdf, 0x00, 0x3c, 0xd4, 0x5e, 0x37, 0x2a, 0xe3, 0xbb, 0x0a, 0xd4, 0xa4, 0x12, 0x11, 0x68, 0x64,
	0x99, 0x41, 0x1d, 0x52, 0x4c, 0xaf, 0xfd, 0x9e, 0x94, 0xe2, 0x84, 0x35, 0xf4, 0x13, 0xda, 0x45,
	0xc0, 0x51, 0x8f, 0x28, 0x63, 0x62, 0x7f, 0x20, 0xea, 0x24, 0x60, 0x0d, 0x7d, 0x87, 0x77, 0x79,
	0x72, 0x91, 0x41, 0x14, 0x7c, 0xdb, 0x26, 0x51, 0xe1, 0xbd, 0x97, 0xe7, 0x51, 0x44, 0x06, 0x51,
	0xc0, 0x6c, 0x9b, 0x44, 0xc5, 0x2b, 0xd6, 0x10, 0x85, 0xae, 0xe2, 0x8d, 0xa2, 0x3e, 0x39, 0xcd,
	0x92, 0x3d, 0x20, 0x8f, 0x40, 0x80, 0xb5, 0xf9, 0x9b, 0xf4, 0x0b, 0x38, 0xb9, 0x0f, 0x00, 0x00,
	0xff, 0xff, 0x0f, 0xaf, 0x5b, 0xf4, 0x1a, 0x05, 0x00, 0x00,
}
