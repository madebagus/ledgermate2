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
// 	protoc        v3.12.2
// source: google/cloud/networksecurity/v1/tls.proto

package networksecurity

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Specification of the GRPC Endpoint.
type GrpcEndpoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The target URI of the gRPC endpoint. Only UDS path is supported, and
	// should start with "unix:".
	TargetUri string `protobuf:"bytes,1,opt,name=target_uri,json=targetUri,proto3" json:"target_uri,omitempty"`
}

func (x *GrpcEndpoint) Reset() {
	*x = GrpcEndpoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_networksecurity_v1_tls_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GrpcEndpoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GrpcEndpoint) ProtoMessage() {}

func (x *GrpcEndpoint) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_networksecurity_v1_tls_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GrpcEndpoint.ProtoReflect.Descriptor instead.
func (*GrpcEndpoint) Descriptor() ([]byte, []int) {
	return file_google_cloud_networksecurity_v1_tls_proto_rawDescGZIP(), []int{0}
}

func (x *GrpcEndpoint) GetTargetUri() string {
	if x != nil {
		return x.TargetUri
	}
	return ""
}

// Specification of ValidationCA. Defines the mechanism to obtain the
// Certificate Authority certificate to validate the peer certificate.
type ValidationCA struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The type of certificate provider which provides the CA certificate.
	//
	// Types that are assignable to Type:
	//	*ValidationCA_GrpcEndpoint
	//	*ValidationCA_CertificateProviderInstance
	Type isValidationCA_Type `protobuf_oneof:"type"`
}

func (x *ValidationCA) Reset() {
	*x = ValidationCA{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_networksecurity_v1_tls_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidationCA) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidationCA) ProtoMessage() {}

func (x *ValidationCA) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_networksecurity_v1_tls_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidationCA.ProtoReflect.Descriptor instead.
func (*ValidationCA) Descriptor() ([]byte, []int) {
	return file_google_cloud_networksecurity_v1_tls_proto_rawDescGZIP(), []int{1}
}

func (m *ValidationCA) GetType() isValidationCA_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *ValidationCA) GetGrpcEndpoint() *GrpcEndpoint {
	if x, ok := x.GetType().(*ValidationCA_GrpcEndpoint); ok {
		return x.GrpcEndpoint
	}
	return nil
}

func (x *ValidationCA) GetCertificateProviderInstance() *CertificateProviderInstance {
	if x, ok := x.GetType().(*ValidationCA_CertificateProviderInstance); ok {
		return x.CertificateProviderInstance
	}
	return nil
}

type isValidationCA_Type interface {
	isValidationCA_Type()
}

type ValidationCA_GrpcEndpoint struct {
	// gRPC specific configuration to access the gRPC server to
	// obtain the CA certificate.
	GrpcEndpoint *GrpcEndpoint `protobuf:"bytes,2,opt,name=grpc_endpoint,json=grpcEndpoint,proto3,oneof"`
}

type ValidationCA_CertificateProviderInstance struct {
	// The certificate provider instance specification that will be passed to
	// the data plane, which will be used to load necessary credential
	// information.
	CertificateProviderInstance *CertificateProviderInstance `protobuf:"bytes,3,opt,name=certificate_provider_instance,json=certificateProviderInstance,proto3,oneof"`
}

func (*ValidationCA_GrpcEndpoint) isValidationCA_Type() {}

func (*ValidationCA_CertificateProviderInstance) isValidationCA_Type() {}

// Specification of a TLS certificate provider instance. Workloads may have one
// or more CertificateProvider instances (plugins) and one of them is enabled
// and configured by specifying this message. Workloads use the values from this
// message to locate and load the CertificateProvider instance configuration.
type CertificateProviderInstance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Plugin instance name, used to locate and load CertificateProvider instance
	// configuration. Set to "google_cloud_private_spiffe" to use Certificate
	// Authority Service certificate provider instance.
	PluginInstance string `protobuf:"bytes,1,opt,name=plugin_instance,json=pluginInstance,proto3" json:"plugin_instance,omitempty"`
}

func (x *CertificateProviderInstance) Reset() {
	*x = CertificateProviderInstance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_networksecurity_v1_tls_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CertificateProviderInstance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CertificateProviderInstance) ProtoMessage() {}

func (x *CertificateProviderInstance) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_networksecurity_v1_tls_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CertificateProviderInstance.ProtoReflect.Descriptor instead.
func (*CertificateProviderInstance) Descriptor() ([]byte, []int) {
	return file_google_cloud_networksecurity_v1_tls_proto_rawDescGZIP(), []int{2}
}

func (x *CertificateProviderInstance) GetPluginInstance() string {
	if x != nil {
		return x.PluginInstance
	}
	return ""
}

// Specification of certificate provider. Defines the mechanism to obtain the
// certificate and private key for peer to peer authentication.
type CertificateProvider struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The type of certificate provider which provides the certificates and
	// private keys.
	//
	// Types that are assignable to Type:
	//	*CertificateProvider_GrpcEndpoint
	//	*CertificateProvider_CertificateProviderInstance
	Type isCertificateProvider_Type `protobuf_oneof:"type"`
}

func (x *CertificateProvider) Reset() {
	*x = CertificateProvider{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_networksecurity_v1_tls_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CertificateProvider) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CertificateProvider) ProtoMessage() {}

func (x *CertificateProvider) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_networksecurity_v1_tls_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CertificateProvider.ProtoReflect.Descriptor instead.
func (*CertificateProvider) Descriptor() ([]byte, []int) {
	return file_google_cloud_networksecurity_v1_tls_proto_rawDescGZIP(), []int{3}
}

func (m *CertificateProvider) GetType() isCertificateProvider_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *CertificateProvider) GetGrpcEndpoint() *GrpcEndpoint {
	if x, ok := x.GetType().(*CertificateProvider_GrpcEndpoint); ok {
		return x.GrpcEndpoint
	}
	return nil
}

func (x *CertificateProvider) GetCertificateProviderInstance() *CertificateProviderInstance {
	if x, ok := x.GetType().(*CertificateProvider_CertificateProviderInstance); ok {
		return x.CertificateProviderInstance
	}
	return nil
}

type isCertificateProvider_Type interface {
	isCertificateProvider_Type()
}

type CertificateProvider_GrpcEndpoint struct {
	// gRPC specific configuration to access the gRPC server to
	// obtain the cert and private key.
	GrpcEndpoint *GrpcEndpoint `protobuf:"bytes,2,opt,name=grpc_endpoint,json=grpcEndpoint,proto3,oneof"`
}

type CertificateProvider_CertificateProviderInstance struct {
	// The certificate provider instance specification that will be passed to
	// the data plane, which will be used to load necessary credential
	// information.
	CertificateProviderInstance *CertificateProviderInstance `protobuf:"bytes,3,opt,name=certificate_provider_instance,json=certificateProviderInstance,proto3,oneof"`
}

func (*CertificateProvider_GrpcEndpoint) isCertificateProvider_Type() {}

func (*CertificateProvider_CertificateProviderInstance) isCertificateProvider_Type() {}

var File_google_cloud_networksecurity_v1_tls_proto protoreflect.FileDescriptor

var file_google_cloud_networksecurity_v1_tls_proto_rawDesc = []byte{
	0x0a, 0x29, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2f, 0x76,
	0x31, 0x2f, 0x74, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62,
	0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x32, 0x0a,
	0x0c, 0x47, 0x72, 0x70, 0x63, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x22, 0x0a,
	0x0a, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x09, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x55, 0x72,
	0x69, 0x22, 0xf1, 0x01, 0x0a, 0x0c, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x43, 0x41, 0x12, 0x54, 0x0a, 0x0d, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72, 0x70, 0x63,
	0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x0c, 0x67, 0x72, 0x70, 0x63,
	0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x82, 0x01, 0x0a, 0x1d, 0x63, 0x65, 0x72,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x3c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x50, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x48, 0x00,
	0x52, 0x1b, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x42, 0x06, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x4b, 0x0a, 0x1b, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x12, 0x2c, 0x0a, 0x0f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x5f, 0x69,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0,
	0x41, 0x02, 0x52, 0x0e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x22, 0xf8, 0x01, 0x0a, 0x13, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x54, 0x0a, 0x0d, 0x67, 0x72,
	0x70, 0x63, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72, 0x70, 0x63, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x48, 0x00, 0x52, 0x0c, 0x67, 0x72, 0x70, 0x63, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x12, 0x82, 0x01, 0x0a, 0x1d, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x65, 0x72, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x49, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x48, 0x00, 0x52, 0x1b, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x42, 0x06, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x42, 0xea, 0x01,
	0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69,
	0x74, 0x79, 0x2e, 0x76, 0x31, 0x42, 0x08, 0x54, 0x6c, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x4e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67,
	0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2f, 0x76,
	0x31, 0x3b, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74,
	0x79, 0xaa, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79,
	0x2e, 0x56, 0x31, 0xca, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x5c, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69,
	0x74, 0x79, 0x5c, 0x56, 0x31, 0xea, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x53, 0x65,
	0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_google_cloud_networksecurity_v1_tls_proto_rawDescOnce sync.Once
	file_google_cloud_networksecurity_v1_tls_proto_rawDescData = file_google_cloud_networksecurity_v1_tls_proto_rawDesc
)

func file_google_cloud_networksecurity_v1_tls_proto_rawDescGZIP() []byte {
	file_google_cloud_networksecurity_v1_tls_proto_rawDescOnce.Do(func() {
		file_google_cloud_networksecurity_v1_tls_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_networksecurity_v1_tls_proto_rawDescData)
	})
	return file_google_cloud_networksecurity_v1_tls_proto_rawDescData
}

var file_google_cloud_networksecurity_v1_tls_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_google_cloud_networksecurity_v1_tls_proto_goTypes = []interface{}{
	(*GrpcEndpoint)(nil),                // 0: google.cloud.networksecurity.v1.GrpcEndpoint
	(*ValidationCA)(nil),                // 1: google.cloud.networksecurity.v1.ValidationCA
	(*CertificateProviderInstance)(nil), // 2: google.cloud.networksecurity.v1.CertificateProviderInstance
	(*CertificateProvider)(nil),         // 3: google.cloud.networksecurity.v1.CertificateProvider
}
var file_google_cloud_networksecurity_v1_tls_proto_depIdxs = []int32{
	0, // 0: google.cloud.networksecurity.v1.ValidationCA.grpc_endpoint:type_name -> google.cloud.networksecurity.v1.GrpcEndpoint
	2, // 1: google.cloud.networksecurity.v1.ValidationCA.certificate_provider_instance:type_name -> google.cloud.networksecurity.v1.CertificateProviderInstance
	0, // 2: google.cloud.networksecurity.v1.CertificateProvider.grpc_endpoint:type_name -> google.cloud.networksecurity.v1.GrpcEndpoint
	2, // 3: google.cloud.networksecurity.v1.CertificateProvider.certificate_provider_instance:type_name -> google.cloud.networksecurity.v1.CertificateProviderInstance
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_google_cloud_networksecurity_v1_tls_proto_init() }
func file_google_cloud_networksecurity_v1_tls_proto_init() {
	if File_google_cloud_networksecurity_v1_tls_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_networksecurity_v1_tls_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GrpcEndpoint); i {
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
		file_google_cloud_networksecurity_v1_tls_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidationCA); i {
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
		file_google_cloud_networksecurity_v1_tls_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CertificateProviderInstance); i {
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
		file_google_cloud_networksecurity_v1_tls_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CertificateProvider); i {
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
	file_google_cloud_networksecurity_v1_tls_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*ValidationCA_GrpcEndpoint)(nil),
		(*ValidationCA_CertificateProviderInstance)(nil),
	}
	file_google_cloud_networksecurity_v1_tls_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*CertificateProvider_GrpcEndpoint)(nil),
		(*CertificateProvider_CertificateProviderInstance)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_networksecurity_v1_tls_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_networksecurity_v1_tls_proto_goTypes,
		DependencyIndexes: file_google_cloud_networksecurity_v1_tls_proto_depIdxs,
		MessageInfos:      file_google_cloud_networksecurity_v1_tls_proto_msgTypes,
	}.Build()
	File_google_cloud_networksecurity_v1_tls_proto = out.File
	file_google_cloud_networksecurity_v1_tls_proto_rawDesc = nil
	file_google_cloud_networksecurity_v1_tls_proto_goTypes = nil
	file_google_cloud_networksecurity_v1_tls_proto_depIdxs = nil
}
