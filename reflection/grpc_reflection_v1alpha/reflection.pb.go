// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc_reflection_v1alpha/reflection.proto

package grpc_reflection_v1alpha

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "github.com/Palen/grpc-go"
	codes "github.com/Palen/grpc-go/codes"
	status "github.com/Palen/grpc-go/status"
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

// The message sent by the client when calling ServerReflectionInfo method.
type ServerReflectionRequest struct {
	Host string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	// To use reflection service, the client should set one of the following
	// fields in message_request. The server distinguishes requests by their
	// defined field and then handles them using corresponding methods.
	//
	// Types that are valid to be assigned to MessageRequest:
	//	*ServerReflectionRequest_FileByFilename
	//	*ServerReflectionRequest_FileContainingSymbol
	//	*ServerReflectionRequest_FileContainingExtension
	//	*ServerReflectionRequest_AllExtensionNumbersOfType
	//	*ServerReflectionRequest_ListServices
	MessageRequest       isServerReflectionRequest_MessageRequest `protobuf_oneof:"message_request"`
	XXX_NoUnkeyedLiteral struct{}                                 `json:"-"`
	XXX_unrecognized     []byte                                   `json:"-"`
	XXX_sizecache        int32                                    `json:"-"`
}

func (m *ServerReflectionRequest) Reset()         { *m = ServerReflectionRequest{} }
func (m *ServerReflectionRequest) String() string { return proto.CompactTextString(m) }
func (*ServerReflectionRequest) ProtoMessage()    {}
func (*ServerReflectionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_42a8ac412db3cb03, []int{0}
}

func (m *ServerReflectionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerReflectionRequest.Unmarshal(m, b)
}
func (m *ServerReflectionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerReflectionRequest.Marshal(b, m, deterministic)
}
func (m *ServerReflectionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerReflectionRequest.Merge(m, src)
}
func (m *ServerReflectionRequest) XXX_Size() int {
	return xxx_messageInfo_ServerReflectionRequest.Size(m)
}
func (m *ServerReflectionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerReflectionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ServerReflectionRequest proto.InternalMessageInfo

func (m *ServerReflectionRequest) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

type isServerReflectionRequest_MessageRequest interface {
	isServerReflectionRequest_MessageRequest()
}

type ServerReflectionRequest_FileByFilename struct {
	FileByFilename string `protobuf:"bytes,3,opt,name=file_by_filename,json=fileByFilename,proto3,oneof"`
}

type ServerReflectionRequest_FileContainingSymbol struct {
	FileContainingSymbol string `protobuf:"bytes,4,opt,name=file_containing_symbol,json=fileContainingSymbol,proto3,oneof"`
}

type ServerReflectionRequest_FileContainingExtension struct {
	FileContainingExtension *ExtensionRequest `protobuf:"bytes,5,opt,name=file_containing_extension,json=fileContainingExtension,proto3,oneof"`
}

type ServerReflectionRequest_AllExtensionNumbersOfType struct {
	AllExtensionNumbersOfType string `protobuf:"bytes,6,opt,name=all_extension_numbers_of_type,json=allExtensionNumbersOfType,proto3,oneof"`
}

type ServerReflectionRequest_ListServices struct {
	ListServices string `protobuf:"bytes,7,opt,name=list_services,json=listServices,proto3,oneof"`
}

func (*ServerReflectionRequest_FileByFilename) isServerReflectionRequest_MessageRequest() {}

func (*ServerReflectionRequest_FileContainingSymbol) isServerReflectionRequest_MessageRequest() {}

func (*ServerReflectionRequest_FileContainingExtension) isServerReflectionRequest_MessageRequest() {}

func (*ServerReflectionRequest_AllExtensionNumbersOfType) isServerReflectionRequest_MessageRequest() {}

func (*ServerReflectionRequest_ListServices) isServerReflectionRequest_MessageRequest() {}

func (m *ServerReflectionRequest) GetMessageRequest() isServerReflectionRequest_MessageRequest {
	if m != nil {
		return m.MessageRequest
	}
	return nil
}

func (m *ServerReflectionRequest) GetFileByFilename() string {
	if x, ok := m.GetMessageRequest().(*ServerReflectionRequest_FileByFilename); ok {
		return x.FileByFilename
	}
	return ""
}

func (m *ServerReflectionRequest) GetFileContainingSymbol() string {
	if x, ok := m.GetMessageRequest().(*ServerReflectionRequest_FileContainingSymbol); ok {
		return x.FileContainingSymbol
	}
	return ""
}

func (m *ServerReflectionRequest) GetFileContainingExtension() *ExtensionRequest {
	if x, ok := m.GetMessageRequest().(*ServerReflectionRequest_FileContainingExtension); ok {
		return x.FileContainingExtension
	}
	return nil
}

func (m *ServerReflectionRequest) GetAllExtensionNumbersOfType() string {
	if x, ok := m.GetMessageRequest().(*ServerReflectionRequest_AllExtensionNumbersOfType); ok {
		return x.AllExtensionNumbersOfType
	}
	return ""
}

func (m *ServerReflectionRequest) GetListServices() string {
	if x, ok := m.GetMessageRequest().(*ServerReflectionRequest_ListServices); ok {
		return x.ListServices
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ServerReflectionRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ServerReflectionRequest_FileByFilename)(nil),
		(*ServerReflectionRequest_FileContainingSymbol)(nil),
		(*ServerReflectionRequest_FileContainingExtension)(nil),
		(*ServerReflectionRequest_AllExtensionNumbersOfType)(nil),
		(*ServerReflectionRequest_ListServices)(nil),
	}
}

// The type name and extension number sent by the client when requesting
// file_containing_extension.
type ExtensionRequest struct {
	// Fully-qualified type name. The format should be <package>.<type>
	ContainingType       string   `protobuf:"bytes,1,opt,name=containing_type,json=containingType,proto3" json:"containing_type,omitempty"`
	ExtensionNumber      int32    `protobuf:"varint,2,opt,name=extension_number,json=extensionNumber,proto3" json:"extension_number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExtensionRequest) Reset()         { *m = ExtensionRequest{} }
func (m *ExtensionRequest) String() string { return proto.CompactTextString(m) }
func (*ExtensionRequest) ProtoMessage()    {}
func (*ExtensionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_42a8ac412db3cb03, []int{1}
}

func (m *ExtensionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExtensionRequest.Unmarshal(m, b)
}
func (m *ExtensionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExtensionRequest.Marshal(b, m, deterministic)
}
func (m *ExtensionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtensionRequest.Merge(m, src)
}
func (m *ExtensionRequest) XXX_Size() int {
	return xxx_messageInfo_ExtensionRequest.Size(m)
}
func (m *ExtensionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtensionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ExtensionRequest proto.InternalMessageInfo

func (m *ExtensionRequest) GetContainingType() string {
	if m != nil {
		return m.ContainingType
	}
	return ""
}

func (m *ExtensionRequest) GetExtensionNumber() int32 {
	if m != nil {
		return m.ExtensionNumber
	}
	return 0
}

// The message sent by the server to answer ServerReflectionInfo method.
type ServerReflectionResponse struct {
	ValidHost       string                   `protobuf:"bytes,1,opt,name=valid_host,json=validHost,proto3" json:"valid_host,omitempty"`
	OriginalRequest *ServerReflectionRequest `protobuf:"bytes,2,opt,name=original_request,json=originalRequest,proto3" json:"original_request,omitempty"`
	// The server sets one of the following fields according to the
	// message_request in the request.
	//
	// Types that are valid to be assigned to MessageResponse:
	//	*ServerReflectionResponse_FileDescriptorResponse
	//	*ServerReflectionResponse_AllExtensionNumbersResponse
	//	*ServerReflectionResponse_ListServicesResponse
	//	*ServerReflectionResponse_ErrorResponse
	MessageResponse      isServerReflectionResponse_MessageResponse `protobuf_oneof:"message_response"`
	XXX_NoUnkeyedLiteral struct{}                                   `json:"-"`
	XXX_unrecognized     []byte                                     `json:"-"`
	XXX_sizecache        int32                                      `json:"-"`
}

func (m *ServerReflectionResponse) Reset()         { *m = ServerReflectionResponse{} }
func (m *ServerReflectionResponse) String() string { return proto.CompactTextString(m) }
func (*ServerReflectionResponse) ProtoMessage()    {}
func (*ServerReflectionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_42a8ac412db3cb03, []int{2}
}

func (m *ServerReflectionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerReflectionResponse.Unmarshal(m, b)
}
func (m *ServerReflectionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerReflectionResponse.Marshal(b, m, deterministic)
}
func (m *ServerReflectionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerReflectionResponse.Merge(m, src)
}
func (m *ServerReflectionResponse) XXX_Size() int {
	return xxx_messageInfo_ServerReflectionResponse.Size(m)
}
func (m *ServerReflectionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerReflectionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ServerReflectionResponse proto.InternalMessageInfo

func (m *ServerReflectionResponse) GetValidHost() string {
	if m != nil {
		return m.ValidHost
	}
	return ""
}

func (m *ServerReflectionResponse) GetOriginalRequest() *ServerReflectionRequest {
	if m != nil {
		return m.OriginalRequest
	}
	return nil
}

type isServerReflectionResponse_MessageResponse interface {
	isServerReflectionResponse_MessageResponse()
}

type ServerReflectionResponse_FileDescriptorResponse struct {
	FileDescriptorResponse *FileDescriptorResponse `protobuf:"bytes,4,opt,name=file_descriptor_response,json=fileDescriptorResponse,proto3,oneof"`
}

type ServerReflectionResponse_AllExtensionNumbersResponse struct {
	AllExtensionNumbersResponse *ExtensionNumberResponse `protobuf:"bytes,5,opt,name=all_extension_numbers_response,json=allExtensionNumbersResponse,proto3,oneof"`
}

type ServerReflectionResponse_ListServicesResponse struct {
	ListServicesResponse *ListServiceResponse `protobuf:"bytes,6,opt,name=list_services_response,json=listServicesResponse,proto3,oneof"`
}

type ServerReflectionResponse_ErrorResponse struct {
	ErrorResponse *ErrorResponse `protobuf:"bytes,7,opt,name=error_response,json=errorResponse,proto3,oneof"`
}

func (*ServerReflectionResponse_FileDescriptorResponse) isServerReflectionResponse_MessageResponse() {}

func (*ServerReflectionResponse_AllExtensionNumbersResponse) isServerReflectionResponse_MessageResponse() {
}

func (*ServerReflectionResponse_ListServicesResponse) isServerReflectionResponse_MessageResponse() {}

func (*ServerReflectionResponse_ErrorResponse) isServerReflectionResponse_MessageResponse() {}

func (m *ServerReflectionResponse) GetMessageResponse() isServerReflectionResponse_MessageResponse {
	if m != nil {
		return m.MessageResponse
	}
	return nil
}

func (m *ServerReflectionResponse) GetFileDescriptorResponse() *FileDescriptorResponse {
	if x, ok := m.GetMessageResponse().(*ServerReflectionResponse_FileDescriptorResponse); ok {
		return x.FileDescriptorResponse
	}
	return nil
}

func (m *ServerReflectionResponse) GetAllExtensionNumbersResponse() *ExtensionNumberResponse {
	if x, ok := m.GetMessageResponse().(*ServerReflectionResponse_AllExtensionNumbersResponse); ok {
		return x.AllExtensionNumbersResponse
	}
	return nil
}

func (m *ServerReflectionResponse) GetListServicesResponse() *ListServiceResponse {
	if x, ok := m.GetMessageResponse().(*ServerReflectionResponse_ListServicesResponse); ok {
		return x.ListServicesResponse
	}
	return nil
}

func (m *ServerReflectionResponse) GetErrorResponse() *ErrorResponse {
	if x, ok := m.GetMessageResponse().(*ServerReflectionResponse_ErrorResponse); ok {
		return x.ErrorResponse
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ServerReflectionResponse) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ServerReflectionResponse_FileDescriptorResponse)(nil),
		(*ServerReflectionResponse_AllExtensionNumbersResponse)(nil),
		(*ServerReflectionResponse_ListServicesResponse)(nil),
		(*ServerReflectionResponse_ErrorResponse)(nil),
	}
}

// Serialized FileDescriptorProto messages sent by the server answering
// a file_by_filename, file_containing_symbol, or file_containing_extension
// request.
type FileDescriptorResponse struct {
	// Serialized FileDescriptorProto messages. We avoid taking a dependency on
	// descriptor.proto, which uses proto2 only features, by making them opaque
	// bytes instead.
	FileDescriptorProto  [][]byte `protobuf:"bytes,1,rep,name=file_descriptor_proto,json=fileDescriptorProto,proto3" json:"file_descriptor_proto,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileDescriptorResponse) Reset()         { *m = FileDescriptorResponse{} }
func (m *FileDescriptorResponse) String() string { return proto.CompactTextString(m) }
func (*FileDescriptorResponse) ProtoMessage()    {}
func (*FileDescriptorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_42a8ac412db3cb03, []int{3}
}

func (m *FileDescriptorResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileDescriptorResponse.Unmarshal(m, b)
}
func (m *FileDescriptorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileDescriptorResponse.Marshal(b, m, deterministic)
}
func (m *FileDescriptorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileDescriptorResponse.Merge(m, src)
}
func (m *FileDescriptorResponse) XXX_Size() int {
	return xxx_messageInfo_FileDescriptorResponse.Size(m)
}
func (m *FileDescriptorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FileDescriptorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FileDescriptorResponse proto.InternalMessageInfo

func (m *FileDescriptorResponse) GetFileDescriptorProto() [][]byte {
	if m != nil {
		return m.FileDescriptorProto
	}
	return nil
}

// A list of extension numbers sent by the server answering
// all_extension_numbers_of_type request.
type ExtensionNumberResponse struct {
	// Full name of the base type, including the package name. The format
	// is <package>.<type>
	BaseTypeName         string   `protobuf:"bytes,1,opt,name=base_type_name,json=baseTypeName,proto3" json:"base_type_name,omitempty"`
	ExtensionNumber      []int32  `protobuf:"varint,2,rep,packed,name=extension_number,json=extensionNumber,proto3" json:"extension_number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExtensionNumberResponse) Reset()         { *m = ExtensionNumberResponse{} }
func (m *ExtensionNumberResponse) String() string { return proto.CompactTextString(m) }
func (*ExtensionNumberResponse) ProtoMessage()    {}
func (*ExtensionNumberResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_42a8ac412db3cb03, []int{4}
}

func (m *ExtensionNumberResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExtensionNumberResponse.Unmarshal(m, b)
}
func (m *ExtensionNumberResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExtensionNumberResponse.Marshal(b, m, deterministic)
}
func (m *ExtensionNumberResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtensionNumberResponse.Merge(m, src)
}
func (m *ExtensionNumberResponse) XXX_Size() int {
	return xxx_messageInfo_ExtensionNumberResponse.Size(m)
}
func (m *ExtensionNumberResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtensionNumberResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ExtensionNumberResponse proto.InternalMessageInfo

func (m *ExtensionNumberResponse) GetBaseTypeName() string {
	if m != nil {
		return m.BaseTypeName
	}
	return ""
}

func (m *ExtensionNumberResponse) GetExtensionNumber() []int32 {
	if m != nil {
		return m.ExtensionNumber
	}
	return nil
}

// A list of ServiceResponse sent by the server answering list_services request.
type ListServiceResponse struct {
	// The information of each service may be expanded in the future, so we use
	// ServiceResponse message to encapsulate it.
	Service              []*ServiceResponse `protobuf:"bytes,1,rep,name=service,proto3" json:"service,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ListServiceResponse) Reset()         { *m = ListServiceResponse{} }
func (m *ListServiceResponse) String() string { return proto.CompactTextString(m) }
func (*ListServiceResponse) ProtoMessage()    {}
func (*ListServiceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_42a8ac412db3cb03, []int{5}
}

func (m *ListServiceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListServiceResponse.Unmarshal(m, b)
}
func (m *ListServiceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListServiceResponse.Marshal(b, m, deterministic)
}
func (m *ListServiceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListServiceResponse.Merge(m, src)
}
func (m *ListServiceResponse) XXX_Size() int {
	return xxx_messageInfo_ListServiceResponse.Size(m)
}
func (m *ListServiceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListServiceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListServiceResponse proto.InternalMessageInfo

func (m *ListServiceResponse) GetService() []*ServiceResponse {
	if m != nil {
		return m.Service
	}
	return nil
}

// The information of a single service used by ListServiceResponse to answer
// list_services request.
type ServiceResponse struct {
	// Full name of a registered service, including its package name. The format
	// is <package>.<service>
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServiceResponse) Reset()         { *m = ServiceResponse{} }
func (m *ServiceResponse) String() string { return proto.CompactTextString(m) }
func (*ServiceResponse) ProtoMessage()    {}
func (*ServiceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_42a8ac412db3cb03, []int{6}
}

func (m *ServiceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceResponse.Unmarshal(m, b)
}
func (m *ServiceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceResponse.Marshal(b, m, deterministic)
}
func (m *ServiceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceResponse.Merge(m, src)
}
func (m *ServiceResponse) XXX_Size() int {
	return xxx_messageInfo_ServiceResponse.Size(m)
}
func (m *ServiceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceResponse proto.InternalMessageInfo

func (m *ServiceResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The error code and error message sent by the server when an error occurs.
type ErrorResponse struct {
	// This field uses the error codes defined in grpc::StatusCode.
	ErrorCode            int32    `protobuf:"varint,1,opt,name=error_code,json=errorCode,proto3" json:"error_code,omitempty"`
	ErrorMessage         string   `protobuf:"bytes,2,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ErrorResponse) Reset()         { *m = ErrorResponse{} }
func (m *ErrorResponse) String() string { return proto.CompactTextString(m) }
func (*ErrorResponse) ProtoMessage()    {}
func (*ErrorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_42a8ac412db3cb03, []int{7}
}

func (m *ErrorResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ErrorResponse.Unmarshal(m, b)
}
func (m *ErrorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ErrorResponse.Marshal(b, m, deterministic)
}
func (m *ErrorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ErrorResponse.Merge(m, src)
}
func (m *ErrorResponse) XXX_Size() int {
	return xxx_messageInfo_ErrorResponse.Size(m)
}
func (m *ErrorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ErrorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ErrorResponse proto.InternalMessageInfo

func (m *ErrorResponse) GetErrorCode() int32 {
	if m != nil {
		return m.ErrorCode
	}
	return 0
}

func (m *ErrorResponse) GetErrorMessage() string {
	if m != nil {
		return m.ErrorMessage
	}
	return ""
}

func init() {
	proto.RegisterType((*ServerReflectionRequest)(nil), "grpc.reflection.v1alpha.ServerReflectionRequest")
	proto.RegisterType((*ExtensionRequest)(nil), "grpc.reflection.v1alpha.ExtensionRequest")
	proto.RegisterType((*ServerReflectionResponse)(nil), "grpc.reflection.v1alpha.ServerReflectionResponse")
	proto.RegisterType((*FileDescriptorResponse)(nil), "grpc.reflection.v1alpha.FileDescriptorResponse")
	proto.RegisterType((*ExtensionNumberResponse)(nil), "grpc.reflection.v1alpha.ExtensionNumberResponse")
	proto.RegisterType((*ListServiceResponse)(nil), "grpc.reflection.v1alpha.ListServiceResponse")
	proto.RegisterType((*ServiceResponse)(nil), "grpc.reflection.v1alpha.ServiceResponse")
	proto.RegisterType((*ErrorResponse)(nil), "grpc.reflection.v1alpha.ErrorResponse")
}

func init() {
	proto.RegisterFile("grpc_reflection_v1alpha/reflection.proto", fileDescriptor_42a8ac412db3cb03)
}

var fileDescriptor_42a8ac412db3cb03 = []byte{
	// 656 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x51, 0x73, 0xd2, 0x40,
	0x10, 0x6e, 0x5a, 0x68, 0x87, 0x85, 0x02, 0x5e, 0x2b, 0xa4, 0x3a, 0x75, 0x98, 0x68, 0x35, 0x75,
	0x1c, 0xda, 0xe2, 0x8c, 0x3f, 0x80, 0xaa, 0x83, 0x33, 0xb5, 0x75, 0x0e, 0x5f, 0x1c, 0x1f, 0x6e,
	0x02, 0x2c, 0x34, 0x1a, 0x72, 0xf1, 0x2e, 0x45, 0x79, 0xf2, 0x47, 0xf8, 0xa3, 0xfc, 0x4b, 0x3e,
	0x3a, 0x77, 0x09, 0x21, 0xa4, 0x44, 0xa7, 0x4f, 0x30, 0xdf, 0xee, 0xde, 0xb7, 0xbb, 0xdf, 0xb7,
	0x01, 0x7b, 0x22, 0x82, 0x21, 0x13, 0x38, 0xf6, 0x70, 0x18, 0xba, 0xdc, 0x67, 0xb3, 0x33, 0xc7,
	0x0b, 0xae, 0x9d, 0x93, 0x25, 0xd4, 0x0e, 0x04, 0x0f, 0x39, 0x69, 0xaa, 0xcc, 0x76, 0x0a, 0x8e,
	0x33, 0xad, 0x3f, 0x9b, 0xd0, 0xec, 0xa3, 0x98, 0xa1, 0xa0, 0x49, 0x90, 0xe2, 0xb7, 0x1b, 0x94,
	0x21, 0x21, 0x50, 0xb8, 0xe6, 0x32, 0x34, 0x8d, 0x96, 0x61, 0x97, 0xa8, 0xfe, 0x4f, 0x9e, 0x43,
	0x7d, 0xec, 0x7a, 0xc8, 0x06, 0x73, 0xa6, 0x7e, 0x7d, 0x67, 0x8a, 0xe6, 0x96, 0x8a, 0xf7, 0x36,
	0x68, 0x55, 0x21, 0xdd, 0xf9, 0xdb, 0x18, 0x27, 0xaf, 0xa0, 0xa1, 0x73, 0x87, 0xdc, 0x0f, 0x1d,
	0xd7, 0x77, 0xfd, 0x09, 0x93, 0xf3, 0xe9, 0x80, 0x7b, 0x66, 0x21, 0xae, 0xd8, 0x57, 0xf1, 0xf3,
	0x24, 0xdc, 0xd7, 0x51, 0x32, 0x81, 0x83, 0x6c, 0x1d, 0xfe, 0x08, 0xd1, 0x97, 0x2e, 0xf7, 0xcd,
	0x62, 0xcb, 0xb0, 0xcb, 0x9d, 0xe3, 0x76, 0xce, 0x40, 0xed, 0x37, 0x8b, 0xcc, 0x78, 0x8a, 0xde,
	0x06, 0x6d, 0xae, 0xb2, 0x24, 0x19, 0xa4, 0x0b, 0x87, 0x8e, 0xe7, 0x2d, 0x1f, 0x67, 0xfe, 0xcd,
	0x74, 0x80, 0x42, 0x32, 0x3e, 0x66, 0xe1, 0x3c, 0x40, 0x73, 0x3b, 0xee, 0xf3, 0xc0, 0xf1, 0xbc,
	0xa4, 0xec, 0x32, 0x4a, 0xba, 0x1a, 0x7f, 0x9c, 0x07, 0x48, 0x8e, 0x60, 0xd7, 0x73, 0x65, 0xc8,
	0x24, 0x8a, 0x99, 0x3b, 0x44, 0x69, 0xee, 0xc4, 0x35, 0x15, 0x05, 0xf7, 0x63, 0xb4, 0x7b, 0x0f,
	0x6a, 0x53, 0x94, 0xd2, 0x99, 0x20, 0x13, 0x51, 0x63, 0xd6, 0x18, 0xea, 0xd9, 0x66, 0xc9, 0x33,
	0xa8, 0xa5, 0xa6, 0xd6, 0x3d, 0x44, 0xdb, 0xaf, 0x2e, 0x61, 0x4d, 0x7b, 0x0c, 0xf5, 0x6c, 0xdb,
	0xe6, 0x66, 0xcb, 0xb0, 0x8b, 0xb4, 0x86, 0xab, 0x8d, 0x5a, 0xbf, 0x0b, 0x60, 0xde, 0x96, 0x58,
	0x06, 0xdc, 0x97, 0x48, 0x0e, 0x01, 0x66, 0x8e, 0xe7, 0x8e, 0x58, 0x4a, 0xe9, 0x92, 0x46, 0x7a,
	0x4a, 0xee, 0xcf, 0x50, 0xe7, 0xc2, 0x9d, 0xb8, 0xbe, 0xe3, 0x2d, 0xfa, 0xd6, 0x34, 0xe5, 0xce,
	0x69, 0xae, 0x02, 0x39, 0x76, 0xa2, 0xb5, 0xc5, 0x4b, 0x8b, 0x61, 0xbf, 0x82, 0xa9, 0x75, 0x1e,
	0xa1, 0x1c, 0x0a, 0x37, 0x08, 0xb9, 0x60, 0x22, 0xee, 0x4b, 0x3b, 0xa4, 0xdc, 0x39, 0xc9, 0x25,
	0x51, 0x26, 0x7b, 0x9d, 0xd4, 0x2d, 0xc6, 0xe9, 0x6d, 0x50, 0x6d, 0xb9, 0xdb, 0x11, 0xf2, 0x1d,
	0x1e, 0xad, 0xd7, 0x3a, 0xa1, 0x2c, 0xfe, 0x67, 0xae, 0x8c, 0x01, 0x52, 0x9c, 0x0f, 0xd7, 0xd8,
	0x23, 0x21, 0x1e, 0x41, 0x63, 0xc5, 0x20, 0x4b, 0xc2, 0x6d, 0x4d, 0xf8, 0x22, 0x97, 0xf0, 0x62,
	0x69, 0xa0, 0x14, 0xd9, 0x7e, 0xda, 0x57, 0x09, 0xcb, 0x15, 0x54, 0x51, 0x88, 0xf4, 0x06, 0x77,
	0xf4, 0xeb, 0x4f, 0xf3, 0xc7, 0x51, 0xe9, 0xa9, 0x77, 0x77, 0x31, 0x0d, 0x74, 0x09, 0xd4, 0x97,
	0x86, 0x8d, 0x30, 0xeb, 0x02, 0x1a, 0xeb, 0xf7, 0x4e, 0x3a, 0x70, 0x3f, 0x2b, 0xa5, 0xfe, 0xf0,
	0x98, 0x46, 0x6b, 0xcb, 0xae, 0xd0, 0xbd, 0x55, 0x51, 0x3e, 0xa8, 0x90, 0xf5, 0x05, 0x9a, 0x39,
	0x2b, 0x25, 0x4f, 0xa0, 0x3a, 0x70, 0x24, 0xea, 0x03, 0x60, 0xfa, 0x1b, 0x13, 0x39, 0xb3, 0xa2,
	0x50, 0xe5, 0xff, 0x4b, 0xf5, 0x7d, 0x59, 0x7f, 0x03, 0x5b, 0xeb, 0x6e, 0xe0, 0x13, 0xec, 0xad,
	0xd9, 0x26, 0xe9, 0xc2, 0x4e, 0x2c, 0x8b, 0x6e, 0xb4, 0xdc, 0xb1, 0xff, 0xe9, 0xea, 0x54, 0x29,
	0x5d, 0x14, 0x5a, 0x47, 0x50, 0xcb, 0x3e, 0x4b, 0xa0, 0x90, 0x6a, 0x5a, 0xff, 0xb7, 0xfa, 0xb0,
	0xbb, 0xb2, 0x71, 0x75, 0x79, 0x91, 0x62, 0x43, 0x3e, 0x8a, 0x52, 0x8b, 0xb4, 0xa4, 0x91, 0x73,
	0x3e, 0x42, 0xf2, 0x18, 0x22, 0x41, 0x58, 0xac, 0x82, 0x3e, 0xbb, 0x12, 0xad, 0x68, 0xf0, 0x7d,
	0x84, 0x75, 0x7e, 0x19, 0x50, 0xcf, 0x9e, 0x1b, 0xf9, 0x09, 0xfb, 0x59, 0xec, 0x9d, 0x3f, 0xe6,
	0xe4, 0xce, 0x17, 0xfb, 0xe0, 0xec, 0x0e, 0x15, 0xd1, 0x54, 0xb6, 0x71, 0x6a, 0x0c, 0xb6, 0xb5,
	0xf4, 0x2f, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0x85, 0x02, 0x09, 0x9d, 0x9f, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ServerReflectionClient is the client API for ServerReflection service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/github.com/Palen/grpc-go#ClientConn.NewStream.
type ServerReflectionClient interface {
	// The reflection service is structured as a bidirectional stream, ensuring
	// all related requests go to a single server.
	ServerReflectionInfo(ctx context.Context, opts ...grpc.CallOption) (ServerReflection_ServerReflectionInfoClient, error)
}

type serverReflectionClient struct {
	cc *grpc.ClientConn
}

func NewServerReflectionClient(cc *grpc.ClientConn) ServerReflectionClient {
	return &serverReflectionClient{cc}
}

func (c *serverReflectionClient) ServerReflectionInfo(ctx context.Context, opts ...grpc.CallOption) (ServerReflection_ServerReflectionInfoClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ServerReflection_serviceDesc.Streams[0], "/grpc.reflection.v1alpha.ServerReflection/ServerReflectionInfo", opts...)
	if err != nil {
		return nil, err
	}
	x := &serverReflectionServerReflectionInfoClient{stream}
	return x, nil
}

type ServerReflection_ServerReflectionInfoClient interface {
	Send(*ServerReflectionRequest) error
	Recv() (*ServerReflectionResponse, error)
	grpc.ClientStream
}

type serverReflectionServerReflectionInfoClient struct {
	grpc.ClientStream
}

func (x *serverReflectionServerReflectionInfoClient) Send(m *ServerReflectionRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *serverReflectionServerReflectionInfoClient) Recv() (*ServerReflectionResponse, error) {
	m := new(ServerReflectionResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ServerReflectionServer is the server API for ServerReflection service.
type ServerReflectionServer interface {
	// The reflection service is structured as a bidirectional stream, ensuring
	// all related requests go to a single server.
	ServerReflectionInfo(ServerReflection_ServerReflectionInfoServer) error
}

// UnimplementedServerReflectionServer can be embedded to have forward compatible implementations.
type UnimplementedServerReflectionServer struct {
}

func (*UnimplementedServerReflectionServer) ServerReflectionInfo(srv ServerReflection_ServerReflectionInfoServer) error {
	return status.Errorf(codes.Unimplemented, "method ServerReflectionInfo not implemented")
}

func RegisterServerReflectionServer(s *grpc.Server, srv ServerReflectionServer) {
	s.RegisterService(&_ServerReflection_serviceDesc, srv)
}

func _ServerReflection_ServerReflectionInfo_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ServerReflectionServer).ServerReflectionInfo(&serverReflectionServerReflectionInfoServer{stream})
}

type ServerReflection_ServerReflectionInfoServer interface {
	Send(*ServerReflectionResponse) error
	Recv() (*ServerReflectionRequest, error)
	grpc.ServerStream
}

type serverReflectionServerReflectionInfoServer struct {
	grpc.ServerStream
}

func (x *serverReflectionServerReflectionInfoServer) Send(m *ServerReflectionResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *serverReflectionServerReflectionInfoServer) Recv() (*ServerReflectionRequest, error) {
	m := new(ServerReflectionRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _ServerReflection_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.reflection.v1alpha.ServerReflection",
	HandlerType: (*ServerReflectionServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ServerReflectionInfo",
			Handler:       _ServerReflection_ServerReflectionInfo_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "grpc_reflection_v1alpha/reflection.proto",
}
