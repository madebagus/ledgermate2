// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.9
// source: google/rpc/context/audit_context.proto

package context

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// `AuditContext` provides information that is needed for audit logging.
type AuditContext struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Serialized audit log.
	AuditLog []byte `protobuf:"bytes,1,opt,name=audit_log,json=auditLog,proto3" json:"audit_log,omitempty"`
	// An API request message that is scrubbed based on the method annotation.
	// This field should only be filled if audit_log field is present.
	// Service Control will use this to assemble a complete log for Cloud Audit
	// Logs and Google internal audit logs.
	ScrubbedRequest *structpb.Struct `protobuf:"bytes,2,opt,name=scrubbed_request,json=scrubbedRequest,proto3" json:"scrubbed_request,omitempty"`
	// An API response message that is scrubbed based on the method annotation.
	// This field should only be filled if audit_log field is present.
	// Service Control will use this to assemble a complete log for Cloud Audit
	// Logs and Google internal audit logs.
	ScrubbedResponse *structpb.Struct `protobuf:"bytes,3,opt,name=scrubbed_response,json=scrubbedResponse,proto3" json:"scrubbed_response,omitempty"`
	// Number of scrubbed response items.
	ScrubbedResponseItemCount int32 `protobuf:"varint,4,opt,name=scrubbed_response_item_count,json=scrubbedResponseItemCount,proto3" json:"scrubbed_response_item_count,omitempty"`
	// Audit resource name which is scrubbed.
	TargetResource string `protobuf:"bytes,5,opt,name=target_resource,json=targetResource,proto3" json:"target_resource,omitempty"`
}

func (x *AuditContext) Reset() {
	*x = AuditContext{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_rpc_context_audit_context_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuditContext) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuditContext) ProtoMessage() {}

func (x *AuditContext) ProtoReflect() protoreflect.Message {
	mi := &file_google_rpc_context_audit_context_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuditContext.ProtoReflect.Descriptor instead.
func (*AuditContext) Descriptor() ([]byte, []int) {
	return file_google_rpc_context_audit_context_proto_rawDescGZIP(), []int{0}
}

func (x *AuditContext) GetAuditLog() []byte {
	if x != nil {
		return x.AuditLog
	}
	return nil
}

func (x *AuditContext) GetScrubbedRequest() *structpb.Struct {
	if x != nil {
		return x.ScrubbedRequest
	}
	return nil
}

func (x *AuditContext) GetScrubbedResponse() *structpb.Struct {
	if x != nil {
		return x.ScrubbedResponse
	}
	return nil
}

func (x *AuditContext) GetScrubbedResponseItemCount() int32 {
	if x != nil {
		return x.ScrubbedResponseItemCount
	}
	return 0
}

func (x *AuditContext) GetTargetResource() string {
	if x != nil {
		return x.TargetResource
	}
	return ""
}

var File_google_rpc_context_audit_context_proto protoreflect.FileDescriptor

var file_google_rpc_context_audit_context_proto_rawDesc = []byte{
	0x0a, 0x26, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x78, 0x74, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9f, 0x02, 0x0a, 0x0c, 0x41,
	0x75, 0x64, 0x69, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x61,
	0x75, 0x64, 0x69, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08,
	0x61, 0x75, 0x64, 0x69, 0x74, 0x4c, 0x6f, 0x67, 0x12, 0x42, 0x0a, 0x10, 0x73, 0x63, 0x72, 0x75,
	0x62, 0x62, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x0f, 0x73, 0x63, 0x72,
	0x75, 0x62, 0x62, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x44, 0x0a, 0x11,
	0x73, 0x63, 0x72, 0x75, 0x62, 0x62, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x52, 0x10, 0x73, 0x63, 0x72, 0x75, 0x62, 0x62, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x3f, 0x0a, 0x1c, 0x73, 0x63, 0x72, 0x75, 0x62, 0x62, 0x65, 0x64, 0x5f, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x19, 0x73, 0x63, 0x72, 0x75, 0x62, 0x62,
	0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x6b, 0x0a, 0x16,
	0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x42, 0x11, 0x41, 0x75, 0x64, 0x69, 0x74, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x78, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x39, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x3b, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0xf8, 0x01, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_google_rpc_context_audit_context_proto_rawDescOnce sync.Once
	file_google_rpc_context_audit_context_proto_rawDescData = file_google_rpc_context_audit_context_proto_rawDesc
)

func file_google_rpc_context_audit_context_proto_rawDescGZIP() []byte {
	file_google_rpc_context_audit_context_proto_rawDescOnce.Do(func() {
		file_google_rpc_context_audit_context_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_rpc_context_audit_context_proto_rawDescData)
	})
	return file_google_rpc_context_audit_context_proto_rawDescData
}

var file_google_rpc_context_audit_context_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_rpc_context_audit_context_proto_goTypes = []interface{}{
	(*AuditContext)(nil),    // 0: google.rpc.context.AuditContext
	(*structpb.Struct)(nil), // 1: google.protobuf.Struct
}
var file_google_rpc_context_audit_context_proto_depIdxs = []int32{
	1, // 0: google.rpc.context.AuditContext.scrubbed_request:type_name -> google.protobuf.Struct
	1, // 1: google.rpc.context.AuditContext.scrubbed_response:type_name -> google.protobuf.Struct
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_google_rpc_context_audit_context_proto_init() }
func file_google_rpc_context_audit_context_proto_init() {
	if File_google_rpc_context_audit_context_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_rpc_context_audit_context_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuditContext); i {
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
			RawDescriptor: file_google_rpc_context_audit_context_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_rpc_context_audit_context_proto_goTypes,
		DependencyIndexes: file_google_rpc_context_audit_context_proto_depIdxs,
		MessageInfos:      file_google_rpc_context_audit_context_proto_msgTypes,
	}.Build()
	File_google_rpc_context_audit_context_proto = out.File
	file_google_rpc_context_audit_context_proto_rawDesc = nil
	file_google_rpc_context_audit_context_proto_goTypes = nil
	file_google_rpc_context_audit_context_proto_depIdxs = nil
}
