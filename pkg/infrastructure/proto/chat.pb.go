// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: proto/chat.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_chat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_chat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_proto_chat_proto_rawDescGZIP(), []int{0}
}

func (x *LoginRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *LoginRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *LoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token   string    `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Players []*Player `protobuf:"bytes,2,rep,name=players,proto3" json:"players,omitempty"`
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_chat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_chat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse.ProtoReflect.Descriptor instead.
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return file_proto_chat_proto_rawDescGZIP(), []int{1}
}

func (x *LoginResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *LoginResponse) GetPlayers() []*Player {
	if x != nil {
		return x.Players
	}
	return nil
}

type StreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestMessage *Message `protobuf:"bytes,1,opt,name=request_message,json=requestMessage,proto3" json:"request_message,omitempty"`
}

func (x *StreamRequest) Reset() {
	*x = StreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_chat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamRequest) ProtoMessage() {}

func (x *StreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_chat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamRequest.ProtoReflect.Descriptor instead.
func (*StreamRequest) Descriptor() ([]byte, []int) {
	return file_proto_chat_proto_rawDescGZIP(), []int{2}
}

func (x *StreamRequest) GetRequestMessage() *Message {
	if x != nil {
		return x.RequestMessage
	}
	return nil
}

type StreamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// Types that are assignable to Event:
	//
	//	*StreamResponse_AddPlayer
	//	*StreamResponse_UpdatePlayer
	//	*StreamResponse_RemovePlayer
	//	*StreamResponse_RoundOver
	//	*StreamResponse_RoundStart
	//	*StreamResponse_ResponseMessage
	Event isStreamResponse_Event `protobuf_oneof:"event"`
}

func (x *StreamResponse) Reset() {
	*x = StreamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_chat_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamResponse) ProtoMessage() {}

func (x *StreamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_chat_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamResponse.ProtoReflect.Descriptor instead.
func (*StreamResponse) Descriptor() ([]byte, []int) {
	return file_proto_chat_proto_rawDescGZIP(), []int{3}
}

func (x *StreamResponse) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (m *StreamResponse) GetEvent() isStreamResponse_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (x *StreamResponse) GetAddPlayer() *AddPlayer {
	if x, ok := x.GetEvent().(*StreamResponse_AddPlayer); ok {
		return x.AddPlayer
	}
	return nil
}

func (x *StreamResponse) GetUpdatePlayer() *UpdatePlayer {
	if x, ok := x.GetEvent().(*StreamResponse_UpdatePlayer); ok {
		return x.UpdatePlayer
	}
	return nil
}

func (x *StreamResponse) GetRemovePlayer() *RemovePlayer {
	if x, ok := x.GetEvent().(*StreamResponse_RemovePlayer); ok {
		return x.RemovePlayer
	}
	return nil
}

func (x *StreamResponse) GetRoundOver() *RoundOver {
	if x, ok := x.GetEvent().(*StreamResponse_RoundOver); ok {
		return x.RoundOver
	}
	return nil
}

func (x *StreamResponse) GetRoundStart() *RoundStart {
	if x, ok := x.GetEvent().(*StreamResponse_RoundStart); ok {
		return x.RoundStart
	}
	return nil
}

func (x *StreamResponse) GetResponseMessage() *Message {
	if x, ok := x.GetEvent().(*StreamResponse_ResponseMessage); ok {
		return x.ResponseMessage
	}
	return nil
}

type isStreamResponse_Event interface {
	isStreamResponse_Event()
}

type StreamResponse_AddPlayer struct {
	AddPlayer *AddPlayer `protobuf:"bytes,2,opt,name=addPlayer,proto3,oneof"`
}

type StreamResponse_UpdatePlayer struct {
	UpdatePlayer *UpdatePlayer `protobuf:"bytes,3,opt,name=updatePlayer,proto3,oneof"`
}

type StreamResponse_RemovePlayer struct {
	RemovePlayer *RemovePlayer `protobuf:"bytes,4,opt,name=removePlayer,proto3,oneof"`
}

type StreamResponse_RoundOver struct {
	RoundOver *RoundOver `protobuf:"bytes,5,opt,name=roundOver,proto3,oneof"`
}

type StreamResponse_RoundStart struct {
	RoundStart *RoundStart `protobuf:"bytes,6,opt,name=roundStart,proto3,oneof"`
}

type StreamResponse_ResponseMessage struct {
	ResponseMessage *Message `protobuf:"bytes,7,opt,name=response_message,json=responseMessage,proto3,oneof"`
}

func (*StreamResponse_AddPlayer) isStreamResponse_Event() {}

func (*StreamResponse_UpdatePlayer) isStreamResponse_Event() {}

func (*StreamResponse_RemovePlayer) isStreamResponse_Event() {}

func (*StreamResponse_RoundOver) isStreamResponse_Event() {}

func (*StreamResponse_RoundStart) isStreamResponse_Event() {}

func (*StreamResponse_ResponseMessage) isStreamResponse_Event() {}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From    string `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To      string `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_chat_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_proto_chat_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_proto_chat_proto_rawDescGZIP(), []int{4}
}

func (x *Message) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *Message) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *Message) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Player struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Player) Reset() {
	*x = Player{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_chat_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Player) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Player) ProtoMessage() {}

func (x *Player) ProtoReflect() protoreflect.Message {
	mi := &file_proto_chat_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Player.ProtoReflect.Descriptor instead.
func (*Player) Descriptor() ([]byte, []int) {
	return file_proto_chat_proto_rawDescGZIP(), []int{5}
}

func (x *Player) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Player) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type AddPlayer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Player *Player `protobuf:"bytes,1,opt,name=player,proto3" json:"player,omitempty"`
}

func (x *AddPlayer) Reset() {
	*x = AddPlayer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_chat_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddPlayer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddPlayer) ProtoMessage() {}

func (x *AddPlayer) ProtoReflect() protoreflect.Message {
	mi := &file_proto_chat_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddPlayer.ProtoReflect.Descriptor instead.
func (*AddPlayer) Descriptor() ([]byte, []int) {
	return file_proto_chat_proto_rawDescGZIP(), []int{6}
}

func (x *AddPlayer) GetPlayer() *Player {
	if x != nil {
		return x.Player
	}
	return nil
}

type UpdatePlayer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Player *Player `protobuf:"bytes,1,opt,name=player,proto3" json:"player,omitempty"`
}

func (x *UpdatePlayer) Reset() {
	*x = UpdatePlayer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_chat_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePlayer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePlayer) ProtoMessage() {}

func (x *UpdatePlayer) ProtoReflect() protoreflect.Message {
	mi := &file_proto_chat_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePlayer.ProtoReflect.Descriptor instead.
func (*UpdatePlayer) Descriptor() ([]byte, []int) {
	return file_proto_chat_proto_rawDescGZIP(), []int{7}
}

func (x *UpdatePlayer) GetPlayer() *Player {
	if x != nil {
		return x.Player
	}
	return nil
}

type RemovePlayer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RemovePlayer) Reset() {
	*x = RemovePlayer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_chat_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemovePlayer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemovePlayer) ProtoMessage() {}

func (x *RemovePlayer) ProtoReflect() protoreflect.Message {
	mi := &file_proto_chat_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemovePlayer.ProtoReflect.Descriptor instead.
func (*RemovePlayer) Descriptor() ([]byte, []int) {
	return file_proto_chat_proto_rawDescGZIP(), []int{8}
}

func (x *RemovePlayer) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type RoundOver struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoundWinnerId string                 `protobuf:"bytes,1,opt,name=roundWinnerId,proto3" json:"roundWinnerId,omitempty"`
	NewRoundAt    *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=newRoundAt,proto3" json:"newRoundAt,omitempty"`
}

func (x *RoundOver) Reset() {
	*x = RoundOver{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_chat_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoundOver) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoundOver) ProtoMessage() {}

func (x *RoundOver) ProtoReflect() protoreflect.Message {
	mi := &file_proto_chat_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoundOver.ProtoReflect.Descriptor instead.
func (*RoundOver) Descriptor() ([]byte, []int) {
	return file_proto_chat_proto_rawDescGZIP(), []int{9}
}

func (x *RoundOver) GetRoundWinnerId() string {
	if x != nil {
		return x.RoundWinnerId
	}
	return ""
}

func (x *RoundOver) GetNewRoundAt() *timestamppb.Timestamp {
	if x != nil {
		return x.NewRoundAt
	}
	return nil
}

type RoundStart struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Players []*Player `protobuf:"bytes,1,rep,name=players,proto3" json:"players,omitempty"`
}

func (x *RoundStart) Reset() {
	*x = RoundStart{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_chat_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoundStart) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoundStart) ProtoMessage() {}

func (x *RoundStart) ProtoReflect() protoreflect.Message {
	mi := &file_proto_chat_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoundStart.ProtoReflect.Descriptor instead.
func (*RoundStart) Descriptor() ([]byte, []int) {
	return file_proto_chat_proto_rawDescGZIP(), []int{10}
}

func (x *RoundStart) GetPlayers() []*Player {
	if x != nil {
		return x.Players
	}
	return nil
}

var File_proto_chat_proto protoreflect.FileDescriptor

var file_proto_chat_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x04, 0x63, 0x68, 0x61, 0x74, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4e, 0x0a, 0x0c, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x4d, 0x0a, 0x0d, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x26, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52,
	0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x22, 0x47, 0x0a, 0x0d, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x0f, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x0e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0x99, 0x03, 0x0a, 0x0e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x2f,
	0x0a, 0x09, 0x61, 0x64, 0x64, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x48, 0x00, 0x52, 0x09, 0x61, 0x64, 0x64, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12,
	0x38, 0x0a, 0x0c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x48, 0x00, 0x52, 0x0c, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x38, 0x0a, 0x0c, 0x72, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x48, 0x00, 0x52, 0x0c, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x12, 0x2f, 0x0a, 0x09, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x4f, 0x76, 0x65, 0x72,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x52, 0x6f,
	0x75, 0x6e, 0x64, 0x4f, 0x76, 0x65, 0x72, 0x48, 0x00, 0x52, 0x09, 0x72, 0x6f, 0x75, 0x6e, 0x64,
	0x4f, 0x76, 0x65, 0x72, 0x12, 0x32, 0x0a, 0x0a, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x53, 0x74, 0x61,
	0x72, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e,
	0x52, 0x6f, 0x75, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x72, 0x74, 0x48, 0x00, 0x52, 0x0a, 0x72, 0x6f,
	0x75, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x3a, 0x0a, 0x10, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x48, 0x00, 0x52, 0x0f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x42, 0x07, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x47, 0x0a,
	0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02,
	0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x2c, 0x0a, 0x06, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x31, 0x0a, 0x09, 0x41, 0x64, 0x64, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x12, 0x24, 0x0a, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52,
	0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x22, 0x34, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x22, 0x1e, 0x0a,
	0x0c, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x6d, 0x0a,
	0x09, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x4f, 0x76, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x0d, 0x72, 0x6f,
	0x75, 0x6e, 0x64, 0x57, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x57, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x3a, 0x0a, 0x0a, 0x6e, 0x65, 0x77, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x41, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x0a, 0x6e, 0x65, 0x77, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x41, 0x74, 0x22, 0x34, 0x0a, 0x0a,
	0x52, 0x6f, 0x75, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x26, 0x0a, 0x07, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x68,
	0x61, 0x74, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x73, 0x32, 0x75, 0x0a, 0x04, 0x47, 0x61, 0x6d, 0x65, 0x12, 0x32, 0x0a, 0x05, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x12, 0x12, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x39,
	0x0a, 0x06, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x13, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e,
	0x63, 0x68, 0x61, 0x74, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x3b, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_chat_proto_rawDescOnce sync.Once
	file_proto_chat_proto_rawDescData = file_proto_chat_proto_rawDesc
)

func file_proto_chat_proto_rawDescGZIP() []byte {
	file_proto_chat_proto_rawDescOnce.Do(func() {
		file_proto_chat_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_chat_proto_rawDescData)
	})
	return file_proto_chat_proto_rawDescData
}

var file_proto_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_proto_chat_proto_goTypes = []interface{}{
	(*LoginRequest)(nil),          // 0: chat.LoginRequest
	(*LoginResponse)(nil),         // 1: chat.LoginResponse
	(*StreamRequest)(nil),         // 2: chat.StreamRequest
	(*StreamResponse)(nil),        // 3: chat.StreamResponse
	(*Message)(nil),               // 4: chat.Message
	(*Player)(nil),                // 5: chat.Player
	(*AddPlayer)(nil),             // 6: chat.AddPlayer
	(*UpdatePlayer)(nil),          // 7: chat.UpdatePlayer
	(*RemovePlayer)(nil),          // 8: chat.RemovePlayer
	(*RoundOver)(nil),             // 9: chat.RoundOver
	(*RoundStart)(nil),            // 10: chat.RoundStart
	(*timestamppb.Timestamp)(nil), // 11: google.protobuf.Timestamp
}
var file_proto_chat_proto_depIdxs = []int32{
	5,  // 0: chat.LoginResponse.players:type_name -> chat.Player
	4,  // 1: chat.StreamRequest.request_message:type_name -> chat.Message
	11, // 2: chat.StreamResponse.timestamp:type_name -> google.protobuf.Timestamp
	6,  // 3: chat.StreamResponse.addPlayer:type_name -> chat.AddPlayer
	7,  // 4: chat.StreamResponse.updatePlayer:type_name -> chat.UpdatePlayer
	8,  // 5: chat.StreamResponse.removePlayer:type_name -> chat.RemovePlayer
	9,  // 6: chat.StreamResponse.roundOver:type_name -> chat.RoundOver
	10, // 7: chat.StreamResponse.roundStart:type_name -> chat.RoundStart
	4,  // 8: chat.StreamResponse.response_message:type_name -> chat.Message
	5,  // 9: chat.AddPlayer.player:type_name -> chat.Player
	5,  // 10: chat.UpdatePlayer.player:type_name -> chat.Player
	11, // 11: chat.RoundOver.newRoundAt:type_name -> google.protobuf.Timestamp
	5,  // 12: chat.RoundStart.players:type_name -> chat.Player
	0,  // 13: chat.Game.Login:input_type -> chat.LoginRequest
	2,  // 14: chat.Game.Stream:input_type -> chat.StreamRequest
	1,  // 15: chat.Game.Login:output_type -> chat.LoginResponse
	3,  // 16: chat.Game.Stream:output_type -> chat.StreamResponse
	15, // [15:17] is the sub-list for method output_type
	13, // [13:15] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_proto_chat_proto_init() }
func file_proto_chat_proto_init() {
	if File_proto_chat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_chat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRequest); i {
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
		file_proto_chat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginResponse); i {
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
		file_proto_chat_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamRequest); i {
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
		file_proto_chat_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamResponse); i {
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
		file_proto_chat_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_proto_chat_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Player); i {
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
		file_proto_chat_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddPlayer); i {
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
		file_proto_chat_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePlayer); i {
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
		file_proto_chat_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemovePlayer); i {
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
		file_proto_chat_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoundOver); i {
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
		file_proto_chat_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoundStart); i {
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
	file_proto_chat_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*StreamResponse_AddPlayer)(nil),
		(*StreamResponse_UpdatePlayer)(nil),
		(*StreamResponse_RemovePlayer)(nil),
		(*StreamResponse_RoundOver)(nil),
		(*StreamResponse_RoundStart)(nil),
		(*StreamResponse_ResponseMessage)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_chat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_chat_proto_goTypes,
		DependencyIndexes: file_proto_chat_proto_depIdxs,
		MessageInfos:      file_proto_chat_proto_msgTypes,
	}.Build()
	File_proto_chat_proto = out.File
	file_proto_chat_proto_rawDesc = nil
	file_proto_chat_proto_goTypes = nil
	file_proto_chat_proto_depIdxs = nil
}
