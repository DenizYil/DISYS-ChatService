// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: api/chat.proto

package chat

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

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User    string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_chat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_api_chat_proto_msgTypes[0]
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
	return file_api_chat_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *Message) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type JoinMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *JoinMessage) Reset() {
	*x = JoinMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_chat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinMessage) ProtoMessage() {}

func (x *JoinMessage) ProtoReflect() protoreflect.Message {
	mi := &file_api_chat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinMessage.ProtoReflect.Descriptor instead.
func (*JoinMessage) Descriptor() ([]byte, []int) {
	return file_api_chat_proto_rawDescGZIP(), []int{1}
}

func (x *JoinMessage) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

type RetrieveMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RetrieveMessage) Reset() {
	*x = RetrieveMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_chat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetrieveMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetrieveMessage) ProtoMessage() {}

func (x *RetrieveMessage) ProtoReflect() protoreflect.Message {
	mi := &file_api_chat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetrieveMessage.ProtoReflect.Descriptor instead.
func (*RetrieveMessage) Descriptor() ([]byte, []int) {
	return file_api_chat_proto_rawDescGZIP(), []int{2}
}

type RetrieveReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *RetrieveReply) Reset() {
	*x = RetrieveReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_chat_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetrieveReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetrieveReply) ProtoMessage() {}

func (x *RetrieveReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_chat_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetrieveReply.ProtoReflect.Descriptor instead.
func (*RetrieveReply) Descriptor() ([]byte, []int) {
	return file_api_chat_proto_rawDescGZIP(), []int{3}
}

func (x *RetrieveReply) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_chat_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_api_chat_proto_msgTypes[4]
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
	return file_api_chat_proto_rawDescGZIP(), []int{4}
}

var File_api_chat_proto protoreflect.FileDescriptor

var file_api_chat_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x04, 0x63, 0x68, 0x61, 0x74, 0x22, 0x37, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22,
	0x21, 0x0a, 0x0b, 0x4a, 0x6f, 0x69, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x22, 0x11, 0x0a, 0x0f, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x29, 0x0a, 0x0d, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76,
	0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x32, 0xe3, 0x01, 0x0a, 0x04, 0x50, 0x65,
	0x65, 0x72, 0x12, 0x27, 0x0a, 0x09, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x12,
	0x0d, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x0b,
	0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x2a, 0x0a, 0x04, 0x4a,
	0x6f, 0x69, 0x6e, 0x12, 0x11, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x0d, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x30, 0x01, 0x12, 0x25, 0x0a, 0x07, 0x50, 0x75, 0x62, 0x6c, 0x69,
	0x73, 0x68, 0x12, 0x0d, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x1a, 0x0b, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x38,
	0x0a, 0x08, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x12, 0x15, 0x2e, 0x63, 0x68, 0x61,
	0x74, 0x2e, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x1a, 0x13, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76,
	0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x30, 0x01, 0x12, 0x25, 0x0a, 0x07, 0x52, 0x65, 0x6c, 0x65,
	0x61, 0x73, 0x65, 0x12, 0x0b, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x0b, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x30, 0x01, 0x42,
	0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_api_chat_proto_rawDescOnce sync.Once
	file_api_chat_proto_rawDescData = file_api_chat_proto_rawDesc
)

func file_api_chat_proto_rawDescGZIP() []byte {
	file_api_chat_proto_rawDescOnce.Do(func() {
		file_api_chat_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_chat_proto_rawDescData)
	})
	return file_api_chat_proto_rawDescData
}

var file_api_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_api_chat_proto_goTypes = []interface{}{
	(*Message)(nil),         // 0: chat.Message
	(*JoinMessage)(nil),     // 1: chat.JoinMessage
	(*RetrieveMessage)(nil), // 2: chat.RetrieveMessage
	(*RetrieveReply)(nil),   // 3: chat.RetrieveReply
	(*Empty)(nil),           // 4: chat.Empty
}
var file_api_chat_proto_depIdxs = []int32{
	0, // 0: chat.Peer.Broadcast:input_type -> chat.Message
	1, // 1: chat.Peer.Join:input_type -> chat.JoinMessage
	0, // 2: chat.Peer.Publish:input_type -> chat.Message
	2, // 3: chat.Peer.Retrieve:input_type -> chat.RetrieveMessage
	4, // 4: chat.Peer.Release:input_type -> chat.Empty
	4, // 5: chat.Peer.Broadcast:output_type -> chat.Empty
	0, // 6: chat.Peer.Join:output_type -> chat.Message
	4, // 7: chat.Peer.Publish:output_type -> chat.Empty
	3, // 8: chat.Peer.Retrieve:output_type -> chat.RetrieveReply
	4, // 9: chat.Peer.Release:output_type -> chat.Empty
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_chat_proto_init() }
func file_api_chat_proto_init() {
	if File_api_chat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_chat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_api_chat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinMessage); i {
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
		file_api_chat_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetrieveMessage); i {
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
		file_api_chat_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetrieveReply); i {
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
		file_api_chat_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_chat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_chat_proto_goTypes,
		DependencyIndexes: file_api_chat_proto_depIdxs,
		MessageInfos:      file_api_chat_proto_msgTypes,
	}.Build()
	File_api_chat_proto = out.File
	file_api_chat_proto_rawDesc = nil
	file_api_chat_proto_goTypes = nil
	file_api_chat_proto_depIdxs = nil
}
