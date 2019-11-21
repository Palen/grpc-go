// Code generated by protoc-gen-go. DO NOT EDIT.
// source: metrics.proto

package grpc_testing

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

// Response message containing the gauge name and value
type GaugeResponse struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Types that are valid to be assigned to Value:
	//	*GaugeResponse_LongValue
	//	*GaugeResponse_DoubleValue
	//	*GaugeResponse_StringValue
	Value                isGaugeResponse_Value `protobuf_oneof:"value"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *GaugeResponse) Reset()         { *m = GaugeResponse{} }
func (m *GaugeResponse) String() string { return proto.CompactTextString(m) }
func (*GaugeResponse) ProtoMessage()    {}
func (*GaugeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6039342a2ba47b72, []int{0}
}

func (m *GaugeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GaugeResponse.Unmarshal(m, b)
}
func (m *GaugeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GaugeResponse.Marshal(b, m, deterministic)
}
func (m *GaugeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GaugeResponse.Merge(m, src)
}
func (m *GaugeResponse) XXX_Size() int {
	return xxx_messageInfo_GaugeResponse.Size(m)
}
func (m *GaugeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GaugeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GaugeResponse proto.InternalMessageInfo

func (m *GaugeResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type isGaugeResponse_Value interface {
	isGaugeResponse_Value()
}

type GaugeResponse_LongValue struct {
	LongValue int64 `protobuf:"varint,2,opt,name=long_value,json=longValue,proto3,oneof"`
}

type GaugeResponse_DoubleValue struct {
	DoubleValue float64 `protobuf:"fixed64,3,opt,name=double_value,json=doubleValue,proto3,oneof"`
}

type GaugeResponse_StringValue struct {
	StringValue string `protobuf:"bytes,4,opt,name=string_value,json=stringValue,proto3,oneof"`
}

func (*GaugeResponse_LongValue) isGaugeResponse_Value() {}

func (*GaugeResponse_DoubleValue) isGaugeResponse_Value() {}

func (*GaugeResponse_StringValue) isGaugeResponse_Value() {}

func (m *GaugeResponse) GetValue() isGaugeResponse_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *GaugeResponse) GetLongValue() int64 {
	if x, ok := m.GetValue().(*GaugeResponse_LongValue); ok {
		return x.LongValue
	}
	return 0
}

func (m *GaugeResponse) GetDoubleValue() float64 {
	if x, ok := m.GetValue().(*GaugeResponse_DoubleValue); ok {
		return x.DoubleValue
	}
	return 0
}

func (m *GaugeResponse) GetStringValue() string {
	if x, ok := m.GetValue().(*GaugeResponse_StringValue); ok {
		return x.StringValue
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*GaugeResponse) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*GaugeResponse_LongValue)(nil),
		(*GaugeResponse_DoubleValue)(nil),
		(*GaugeResponse_StringValue)(nil),
	}
}

// Request message containing the gauge name
type GaugeRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GaugeRequest) Reset()         { *m = GaugeRequest{} }
func (m *GaugeRequest) String() string { return proto.CompactTextString(m) }
func (*GaugeRequest) ProtoMessage()    {}
func (*GaugeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6039342a2ba47b72, []int{1}
}

func (m *GaugeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GaugeRequest.Unmarshal(m, b)
}
func (m *GaugeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GaugeRequest.Marshal(b, m, deterministic)
}
func (m *GaugeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GaugeRequest.Merge(m, src)
}
func (m *GaugeRequest) XXX_Size() int {
	return xxx_messageInfo_GaugeRequest.Size(m)
}
func (m *GaugeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GaugeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GaugeRequest proto.InternalMessageInfo

func (m *GaugeRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type EmptyMessage struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmptyMessage) Reset()         { *m = EmptyMessage{} }
func (m *EmptyMessage) String() string { return proto.CompactTextString(m) }
func (*EmptyMessage) ProtoMessage()    {}
func (*EmptyMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_6039342a2ba47b72, []int{2}
}

func (m *EmptyMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmptyMessage.Unmarshal(m, b)
}
func (m *EmptyMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmptyMessage.Marshal(b, m, deterministic)
}
func (m *EmptyMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmptyMessage.Merge(m, src)
}
func (m *EmptyMessage) XXX_Size() int {
	return xxx_messageInfo_EmptyMessage.Size(m)
}
func (m *EmptyMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_EmptyMessage.DiscardUnknown(m)
}

var xxx_messageInfo_EmptyMessage proto.InternalMessageInfo

func init() {
	proto.RegisterType((*GaugeResponse)(nil), "grpc.testing.GaugeResponse")
	proto.RegisterType((*GaugeRequest)(nil), "grpc.testing.GaugeRequest")
	proto.RegisterType((*EmptyMessage)(nil), "grpc.testing.EmptyMessage")
}

func init() { proto.RegisterFile("metrics.proto", fileDescriptor_6039342a2ba47b72) }

var fileDescriptor_6039342a2ba47b72 = []byte{
	// 256 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x3f, 0x4f, 0xc3, 0x30,
	0x10, 0xc5, 0x6b, 0x5a, 0xfe, 0xf4, 0x70, 0x3b, 0x78, 0xaa, 0xca, 0x40, 0x14, 0x96, 0x4c, 0x11,
	0x82, 0x4f, 0x00, 0x08, 0xa5, 0x0c, 0x5d, 0x82, 0xc4, 0x8a, 0xd2, 0x70, 0xb2, 0x22, 0x39, 0x71,
	0xf0, 0x5d, 0x2a, 0xf1, 0x49, 0x58, 0xf9, 0xa8, 0xc8, 0x4e, 0x55, 0xa5, 0x08, 0x75, 0xb3, 0x7e,
	0xf7, 0xfc, 0xfc, 0x9e, 0x0f, 0x66, 0x35, 0xb2, 0xab, 0x4a, 0x4a, 0x5b, 0x67, 0xd9, 0x2a, 0xa9,
	0x5d, 0x5b, 0xa6, 0x8c, 0xc4, 0x55, 0xa3, 0xe3, 0x6f, 0x01, 0xb3, 0xac, 0xe8, 0x34, 0xe6, 0x48,
	0xad, 0x6d, 0x08, 0x95, 0x82, 0x49, 0x53, 0xd4, 0xb8, 0x10, 0x91, 0x48, 0xa6, 0x79, 0x38, 0xab,
	0x6b, 0x00, 0x63, 0x1b, 0xfd, 0xbe, 0x2d, 0x4c, 0x87, 0x8b, 0x93, 0x48, 0x24, 0xe3, 0xd5, 0x28,
	0x9f, 0x7a, 0xf6, 0xe6, 0x91, 0xba, 0x01, 0xf9, 0x61, 0xbb, 0x8d, 0xc1, 0x9d, 0x64, 0x1c, 0x89,
	0x44, 0xac, 0x46, 0xf9, 0x65, 0x4f, 0xf7, 0x22, 0x62, 0x57, 0xed, 0x7d, 0x26, 0xfe, 0x05, 0x2f,
	0xea, 0x69, 0x10, 0x3d, 0x9e, 0xc3, 0x69, 0x98, 0xc6, 0x31, 0xc8, 0x5d, 0xb0, 0xcf, 0x0e, 0x89,
	0xff, 0xcb, 0x15, 0xcf, 0x41, 0x3e, 0xd7, 0x2d, 0x7f, 0xad, 0x91, 0xa8, 0xd0, 0x78, 0xf7, 0x23,
	0x60, 0xbe, 0xee, 0xdb, 0xbe, 0xa2, 0xdb, 0x56, 0x25, 0xaa, 0x17, 0x90, 0x19, 0xf2, 0x83, 0x31,
	0xc1, 0x8c, 0xd4, 0x32, 0x1d, 0xf6, 0x4f, 0x87, 0xd7, 0x97, 0x57, 0x87, 0xb3, 0x83, 0x7f, 0xb9,
	0x15, 0xea, 0x09, 0x2e, 0x32, 0xe4, 0x40, 0xff, 0xda, 0x0c, 0x93, 0x1e, 0xb5, 0xd9, 0x9c, 0x85,
	0x2d, 0xdc, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0x5e, 0x7d, 0xb2, 0xc9, 0x96, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MetricsServiceClient is the client API for MetricsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/github.com/Palen/grpc-go#ClientConn.NewStream.
type MetricsServiceClient interface {
	// Returns the values of all the gauges that are currently being maintained by
	// the service
	GetAllGauges(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (MetricsService_GetAllGaugesClient, error)
	// Returns the value of one gauge
	GetGauge(ctx context.Context, in *GaugeRequest, opts ...grpc.CallOption) (*GaugeResponse, error)
}

type metricsServiceClient struct {
	cc *grpc.ClientConn
}

func NewMetricsServiceClient(cc *grpc.ClientConn) MetricsServiceClient {
	return &metricsServiceClient{cc}
}

func (c *metricsServiceClient) GetAllGauges(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (MetricsService_GetAllGaugesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_MetricsService_serviceDesc.Streams[0], "/grpc.testing.MetricsService/GetAllGauges", opts...)
	if err != nil {
		return nil, err
	}
	x := &metricsServiceGetAllGaugesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MetricsService_GetAllGaugesClient interface {
	Recv() (*GaugeResponse, error)
	grpc.ClientStream
}

type metricsServiceGetAllGaugesClient struct {
	grpc.ClientStream
}

func (x *metricsServiceGetAllGaugesClient) Recv() (*GaugeResponse, error) {
	m := new(GaugeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *metricsServiceClient) GetGauge(ctx context.Context, in *GaugeRequest, opts ...grpc.CallOption) (*GaugeResponse, error) {
	out := new(GaugeResponse)
	err := c.cc.Invoke(ctx, "/grpc.testing.MetricsService/GetGauge", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MetricsServiceServer is the server API for MetricsService service.
type MetricsServiceServer interface {
	// Returns the values of all the gauges that are currently being maintained by
	// the service
	GetAllGauges(*EmptyMessage, MetricsService_GetAllGaugesServer) error
	// Returns the value of one gauge
	GetGauge(context.Context, *GaugeRequest) (*GaugeResponse, error)
}

// UnimplementedMetricsServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMetricsServiceServer struct {
}

func (*UnimplementedMetricsServiceServer) GetAllGauges(req *EmptyMessage, srv MetricsService_GetAllGaugesServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAllGauges not implemented")
}
func (*UnimplementedMetricsServiceServer) GetGauge(ctx context.Context, req *GaugeRequest) (*GaugeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGauge not implemented")
}

func RegisterMetricsServiceServer(s *grpc.Server, srv MetricsServiceServer) {
	s.RegisterService(&_MetricsService_serviceDesc, srv)
}

func _MetricsService_GetAllGauges_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(EmptyMessage)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MetricsServiceServer).GetAllGauges(m, &metricsServiceGetAllGaugesServer{stream})
}

type MetricsService_GetAllGaugesServer interface {
	Send(*GaugeResponse) error
	grpc.ServerStream
}

type metricsServiceGetAllGaugesServer struct {
	grpc.ServerStream
}

func (x *metricsServiceGetAllGaugesServer) Send(m *GaugeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _MetricsService_GetGauge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GaugeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsServiceServer).GetGauge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.testing.MetricsService/GetGauge",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsServiceServer).GetGauge(ctx, req.(*GaugeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MetricsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.testing.MetricsService",
	HandlerType: (*MetricsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGauge",
			Handler:    _MetricsService_GetGauge_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetAllGauges",
			Handler:       _MetricsService_GetAllGauges_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "metrics.proto",
}
