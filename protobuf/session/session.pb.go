// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: session.proto

package session

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

type Session struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ttl        uint64 `protobuf:"varint,1,opt,name=ttl,proto3" json:"ttl,omitempty"`                                 // time to live of session
	SessionId  string `protobuf:"bytes,2,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`     // base64 encoded UUID of session
	CreatedAt  uint64 `protobuf:"varint,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`    // unix timestamp on when session has been created
	UpdatedAt  uint64 `protobuf:"varint,4,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`    // unix timestamp on when session has been updated
	CompiledAt uint64 `protobuf:"varint,5,opt,name=compiled_at,json=compiledAt,proto3" json:"compiled_at,omitempty"` // unix timestamp on when session code has been compiled last time
}

func (x *Session) Reset() {
	*x = Session{}
	if protoimpl.UnsafeEnabled {
		mi := &file_session_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Session) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Session) ProtoMessage() {}

func (x *Session) ProtoReflect() protoreflect.Message {
	mi := &file_session_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Session.ProtoReflect.Descriptor instead.
func (*Session) Descriptor() ([]byte, []int) {
	return file_session_proto_rawDescGZIP(), []int{0}
}

func (x *Session) GetTtl() uint64 {
	if x != nil {
		return x.Ttl
	}
	return 0
}

func (x *Session) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *Session) GetCreatedAt() uint64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Session) GetUpdatedAt() uint64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *Session) GetCompiledAt() uint64 {
	if x != nil {
		return x.CompiledAt
	}
	return 0
}

type CompilationInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version      string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	CompileError bool   `protobuf:"varint,2,opt,name=compile_error,json=compileError,proto3" json:"compile_error,omitempty"`
	Message      string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *CompilationInfo) Reset() {
	*x = CompilationInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_session_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompilationInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompilationInfo) ProtoMessage() {}

func (x *CompilationInfo) ProtoReflect() protoreflect.Message {
	mi := &file_session_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompilationInfo.ProtoReflect.Descriptor instead.
func (*CompilationInfo) Descriptor() ([]byte, []int) {
	return file_session_proto_rawDescGZIP(), []int{1}
}

func (x *CompilationInfo) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *CompilationInfo) GetCompileError() bool {
	if x != nil {
		return x.CompileError
	}
	return false
}

func (x *CompilationInfo) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type SessionLegacyNode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	IsFile   bool                 `protobuf:"varint,2,opt,name=is_file,json=isFile,proto3" json:"is_file,omitempty"`
	Children []*SessionLegacyNode `protobuf:"bytes,3,rep,name=children,proto3" json:"children,omitempty"`
}

func (x *SessionLegacyNode) Reset() {
	*x = SessionLegacyNode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_session_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SessionLegacyNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SessionLegacyNode) ProtoMessage() {}

func (x *SessionLegacyNode) ProtoReflect() protoreflect.Message {
	mi := &file_session_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SessionLegacyNode.ProtoReflect.Descriptor instead.
func (*SessionLegacyNode) Descriptor() ([]byte, []int) {
	return file_session_proto_rawDescGZIP(), []int{2}
}

func (x *SessionLegacyNode) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SessionLegacyNode) GetIsFile() bool {
	if x != nil {
		return x.IsFile
	}
	return false
}

func (x *SessionLegacyNode) GetChildren() []*SessionLegacyNode {
	if x != nil {
		return x.Children
	}
	return nil
}

type SessionTree struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FilePaths  map[string]string `protobuf:"bytes,1,rep,name=file_paths,json=filePaths,proto3" json:"file_paths,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	FilesOrder []string          `protobuf:"bytes,2,rep,name=files_order,json=filesOrder,proto3" json:"files_order,omitempty"`
}

func (x *SessionTree) Reset() {
	*x = SessionTree{}
	if protoimpl.UnsafeEnabled {
		mi := &file_session_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SessionTree) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SessionTree) ProtoMessage() {}

func (x *SessionTree) ProtoReflect() protoreflect.Message {
	mi := &file_session_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SessionTree.ProtoReflect.Descriptor instead.
func (*SessionTree) Descriptor() ([]byte, []int) {
	return file_session_proto_rawDescGZIP(), []int{3}
}

func (x *SessionTree) GetFilePaths() map[string]string {
	if x != nil {
		return x.FilePaths
	}
	return nil
}

func (x *SessionTree) GetFilesOrder() []string {
	if x != nil {
		return x.FilesOrder
	}
	return nil
}

var File_session_proto protoreflect.FileDescriptor

var file_session_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x99, 0x01, 0x0a, 0x07, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x74,
	0x74, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x74, 0x74, 0x6c, 0x12, 0x1d, 0x0a,
	0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f,
	0x6d, 0x70, 0x69, 0x6c, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x64, 0x41, 0x74, 0x22, 0x6a, 0x0a, 0x0f, 0x43,
	0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18,
	0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x70,
	0x69, 0x6c, 0x65, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0c, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x70, 0x0a, 0x11, 0x53, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x4c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x17, 0x0a, 0x07, 0x69, 0x73, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x06, 0x69, 0x73, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x2e, 0x0a, 0x08, 0x63, 0x68, 0x69,
	0x6c, 0x64, 0x72, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x53, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x4e, 0x6f, 0x64, 0x65, 0x52,
	0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x22, 0xa8, 0x01, 0x0a, 0x0b, 0x53, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x72, 0x65, 0x65, 0x12, 0x3a, 0x0a, 0x0a, 0x66, 0x69, 0x6c,
	0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x72, 0x65, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x65,
	0x50, 0x61, 0x74, 0x68, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x66, 0x69, 0x6c, 0x65,
	0x50, 0x61, 0x74, 0x68, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x5f, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x69, 0x6c, 0x65,
	0x73, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x1a, 0x3c, 0x0a, 0x0e, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x61,
	0x74, 0x68, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x42, 0x16, 0x5a, 0x14, 0x73, 0x77, 0x69, 0x2d, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_session_proto_rawDescOnce sync.Once
	file_session_proto_rawDescData = file_session_proto_rawDesc
)

func file_session_proto_rawDescGZIP() []byte {
	file_session_proto_rawDescOnce.Do(func() {
		file_session_proto_rawDescData = protoimpl.X.CompressGZIP(file_session_proto_rawDescData)
	})
	return file_session_proto_rawDescData
}

var file_session_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_session_proto_goTypes = []interface{}{
	(*Session)(nil),           // 0: Session
	(*CompilationInfo)(nil),   // 1: CompilationInfo
	(*SessionLegacyNode)(nil), // 2: SessionLegacyNode
	(*SessionTree)(nil),       // 3: SessionTree
	nil,                       // 4: SessionTree.FilePathsEntry
}
var file_session_proto_depIdxs = []int32{
	2, // 0: SessionLegacyNode.children:type_name -> SessionLegacyNode
	4, // 1: SessionTree.file_paths:type_name -> SessionTree.FilePathsEntry
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_session_proto_init() }
func file_session_proto_init() {
	if File_session_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_session_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Session); i {
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
		file_session_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompilationInfo); i {
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
		file_session_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SessionLegacyNode); i {
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
		file_session_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SessionTree); i {
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
			RawDescriptor: file_session_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_session_proto_goTypes,
		DependencyIndexes: file_session_proto_depIdxs,
		MessageInfos:      file_session_proto_msgTypes,
	}.Build()
	File_session_proto = out.File
	file_session_proto_rawDesc = nil
	file_session_proto_goTypes = nil
	file_session_proto_depIdxs = nil
}
