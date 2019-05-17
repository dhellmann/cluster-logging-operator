// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/services/hotel_group_view_service.proto

package services // import "google.golang.org/genproto/googleapis/ads/googleads/v1/services"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import resources "google.golang.org/genproto/googleapis/ads/googleads/v1/resources"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Request message for [HotelGroupViewService.GetHotelGroupView][google.ads.googleads.v1.services.HotelGroupViewService.GetHotelGroupView].
type GetHotelGroupViewRequest struct {
	// Resource name of the Hotel Group View to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetHotelGroupViewRequest) Reset()         { *m = GetHotelGroupViewRequest{} }
func (m *GetHotelGroupViewRequest) String() string { return proto.CompactTextString(m) }
func (*GetHotelGroupViewRequest) ProtoMessage()    {}
func (*GetHotelGroupViewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_hotel_group_view_service_1b795ba7f86fa296, []int{0}
}
func (m *GetHotelGroupViewRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetHotelGroupViewRequest.Unmarshal(m, b)
}
func (m *GetHotelGroupViewRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetHotelGroupViewRequest.Marshal(b, m, deterministic)
}
func (dst *GetHotelGroupViewRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetHotelGroupViewRequest.Merge(dst, src)
}
func (m *GetHotelGroupViewRequest) XXX_Size() int {
	return xxx_messageInfo_GetHotelGroupViewRequest.Size(m)
}
func (m *GetHotelGroupViewRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetHotelGroupViewRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetHotelGroupViewRequest proto.InternalMessageInfo

func (m *GetHotelGroupViewRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetHotelGroupViewRequest)(nil), "google.ads.googleads.v1.services.GetHotelGroupViewRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HotelGroupViewServiceClient is the client API for HotelGroupViewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HotelGroupViewServiceClient interface {
	// Returns the requested Hotel Group View in full detail.
	GetHotelGroupView(ctx context.Context, in *GetHotelGroupViewRequest, opts ...grpc.CallOption) (*resources.HotelGroupView, error)
}

type hotelGroupViewServiceClient struct {
	cc *grpc.ClientConn
}

func NewHotelGroupViewServiceClient(cc *grpc.ClientConn) HotelGroupViewServiceClient {
	return &hotelGroupViewServiceClient{cc}
}

func (c *hotelGroupViewServiceClient) GetHotelGroupView(ctx context.Context, in *GetHotelGroupViewRequest, opts ...grpc.CallOption) (*resources.HotelGroupView, error) {
	out := new(resources.HotelGroupView)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.HotelGroupViewService/GetHotelGroupView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HotelGroupViewServiceServer is the server API for HotelGroupViewService service.
type HotelGroupViewServiceServer interface {
	// Returns the requested Hotel Group View in full detail.
	GetHotelGroupView(context.Context, *GetHotelGroupViewRequest) (*resources.HotelGroupView, error)
}

func RegisterHotelGroupViewServiceServer(s *grpc.Server, srv HotelGroupViewServiceServer) {
	s.RegisterService(&_HotelGroupViewService_serviceDesc, srv)
}

func _HotelGroupViewService_GetHotelGroupView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHotelGroupViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HotelGroupViewServiceServer).GetHotelGroupView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.HotelGroupViewService/GetHotelGroupView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HotelGroupViewServiceServer).GetHotelGroupView(ctx, req.(*GetHotelGroupViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HotelGroupViewService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v1.services.HotelGroupViewService",
	HandlerType: (*HotelGroupViewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetHotelGroupView",
			Handler:    _HotelGroupViewService_GetHotelGroupView_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v1/services/hotel_group_view_service.proto",
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/services/hotel_group_view_service.proto", fileDescriptor_hotel_group_view_service_1b795ba7f86fa296)
}

var fileDescriptor_hotel_group_view_service_1b795ba7f86fa296 = []byte{
	// 371 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xc1, 0x4a, 0xeb, 0x40,
	0x14, 0x86, 0x49, 0x2e, 0x5c, 0xb8, 0xe1, 0xba, 0x30, 0x20, 0x94, 0xe0, 0xa2, 0xd4, 0x2e, 0xa4,
	0x8b, 0x19, 0x62, 0x37, 0x3a, 0x22, 0x25, 0xdd, 0xa4, 0x2b, 0x29, 0x15, 0xb2, 0x90, 0x40, 0x88,
	0xcd, 0x21, 0x06, 0x9a, 0x4c, 0xcc, 0x99, 0xa4, 0x0b, 0x71, 0xa1, 0xaf, 0xe0, 0x1b, 0xb8, 0xf4,
	0x1d, 0x7c, 0x01, 0xb7, 0x2e, 0x7c, 0x01, 0x57, 0x3e, 0x85, 0xa4, 0xd3, 0x09, 0x14, 0x1b, 0xba,
	0xfb, 0x99, 0xf9, 0xbf, 0x7f, 0xce, 0xf9, 0x13, 0x63, 0x14, 0x73, 0x1e, 0x2f, 0x80, 0x86, 0x11,
	0x52, 0x29, 0x6b, 0x55, 0xd9, 0x14, 0xa1, 0xa8, 0x92, 0x39, 0x20, 0xbd, 0xe5, 0x02, 0x16, 0x41,
	0x5c, 0xf0, 0x32, 0x0f, 0xaa, 0x04, 0x96, 0xc1, 0xfa, 0x86, 0xe4, 0x05, 0x17, 0xdc, 0xec, 0x4a,
	0x8a, 0x84, 0x11, 0x92, 0x26, 0x80, 0x54, 0x36, 0x51, 0x01, 0xd6, 0x69, 0xdb, 0x13, 0x05, 0x20,
	0x2f, 0x8b, 0x6d, 0x6f, 0xc8, 0x6c, 0xeb, 0x50, 0x91, 0x79, 0x42, 0xc3, 0x2c, 0xe3, 0x22, 0x14,
	0x09, 0xcf, 0x50, 0xde, 0xf6, 0x46, 0x46, 0xc7, 0x05, 0x31, 0xa9, 0x51, 0xb7, 0x26, 0xbd, 0x04,
	0x96, 0x33, 0xb8, 0x2b, 0x01, 0x85, 0x79, 0x64, 0xec, 0xa9, 0xf4, 0x20, 0x0b, 0x53, 0xe8, 0x68,
	0x5d, 0xed, 0xf8, 0xdf, 0xec, 0xbf, 0x3a, 0xbc, 0x0c, 0x53, 0x38, 0xf9, 0xd4, 0x8c, 0x83, 0x4d,
	0xfc, 0x4a, 0xce, 0x6c, 0xbe, 0x69, 0xc6, 0xfe, 0xaf, 0x6c, 0x93, 0x91, 0x5d, 0xbb, 0x92, 0xb6,
	0x81, 0x2c, 0xbb, 0x95, 0x6d, 0x5a, 0x20, 0x9b, 0x64, 0xef, 0xec, 0xe9, 0xe3, 0xeb, 0x59, 0x1f,
	0x9a, 0x76, 0xdd, 0xd5, 0xfd, 0xc6, 0x3a, 0x17, 0xf3, 0x12, 0x05, 0x4f, 0xa1, 0x40, 0x3a, 0x90,
	0xe5, 0x35, 0x18, 0xd2, 0xc1, 0xc3, 0xf8, 0x51, 0x37, 0xfa, 0x73, 0x9e, 0xee, 0x9c, 0x77, 0x6c,
	0x6d, 0xdd, 0x7f, 0x5a, 0xf7, 0x3b, 0xd5, 0xae, 0x27, 0x6b, 0x3e, 0xe6, 0x8b, 0x30, 0x8b, 0x09,
	0x2f, 0x62, 0x1a, 0x43, 0xb6, 0x6a, 0x5f, 0x7d, 0xc9, 0x3c, 0xc1, 0xf6, 0x7f, 0xe7, 0x5c, 0x89,
	0x17, 0xfd, 0x8f, 0xeb, 0x38, 0xaf, 0x7a, 0xd7, 0x95, 0x81, 0x4e, 0x84, 0x44, 0xca, 0x5a, 0x79,
	0x36, 0x59, 0x3f, 0x8c, 0xef, 0xca, 0xe2, 0x3b, 0x11, 0xfa, 0x8d, 0xc5, 0xf7, 0x6c, 0x5f, 0x59,
	0xbe, 0xf5, 0xbe, 0x3c, 0x67, 0xcc, 0x89, 0x90, 0xb1, 0xc6, 0xc4, 0x98, 0x67, 0x33, 0xa6, 0x6c,
	0x37, 0x7f, 0x57, 0x73, 0x0e, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x34, 0xd6, 0x3e, 0x66, 0xe2,
	0x02, 0x00, 0x00,
}
