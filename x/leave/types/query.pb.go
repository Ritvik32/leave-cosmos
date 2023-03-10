// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/leave/v1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type GetStudentRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *GetStudentRequest) Reset()         { *m = GetStudentRequest{} }
func (m *GetStudentRequest) String() string { return proto.CompactTextString(m) }
func (*GetStudentRequest) ProtoMessage()    {}
func (*GetStudentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f325f862ed9d7a2f, []int{0}
}
func (m *GetStudentRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetStudentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetStudentRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetStudentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStudentRequest.Merge(m, src)
}
func (m *GetStudentRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetStudentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStudentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetStudentRequest proto.InternalMessageInfo

func (m *GetStudentRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetStudentResponse struct {
	Student *Student `protobuf:"bytes,1,opt,name=student,proto3" json:"student,omitempty"`
}

func (m *GetStudentResponse) Reset()         { *m = GetStudentResponse{} }
func (m *GetStudentResponse) String() string { return proto.CompactTextString(m) }
func (*GetStudentResponse) ProtoMessage()    {}
func (*GetStudentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f325f862ed9d7a2f, []int{1}
}
func (m *GetStudentResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetStudentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetStudentResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetStudentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStudentResponse.Merge(m, src)
}
func (m *GetStudentResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetStudentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStudentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetStudentResponse proto.InternalMessageInfo

func (m *GetStudentResponse) GetStudent() *Student {
	if m != nil {
		return m.Student
	}
	return nil
}

type ListAllLeavesRequest struct {
}

func (m *ListAllLeavesRequest) Reset()         { *m = ListAllLeavesRequest{} }
func (m *ListAllLeavesRequest) String() string { return proto.CompactTextString(m) }
func (*ListAllLeavesRequest) ProtoMessage()    {}
func (*ListAllLeavesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f325f862ed9d7a2f, []int{2}
}
func (m *ListAllLeavesRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListAllLeavesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListAllLeavesRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ListAllLeavesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListAllLeavesRequest.Merge(m, src)
}
func (m *ListAllLeavesRequest) XXX_Size() int {
	return m.Size()
}
func (m *ListAllLeavesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListAllLeavesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListAllLeavesRequest proto.InternalMessageInfo

type ListAllLeavesResponse struct {
	Leave []*Leave `protobuf:"bytes,1,rep,name=leave,proto3" json:"leave,omitempty"`
}

func (m *ListAllLeavesResponse) Reset()         { *m = ListAllLeavesResponse{} }
func (m *ListAllLeavesResponse) String() string { return proto.CompactTextString(m) }
func (*ListAllLeavesResponse) ProtoMessage()    {}
func (*ListAllLeavesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f325f862ed9d7a2f, []int{3}
}
func (m *ListAllLeavesResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListAllLeavesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListAllLeavesResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ListAllLeavesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListAllLeavesResponse.Merge(m, src)
}
func (m *ListAllLeavesResponse) XXX_Size() int {
	return m.Size()
}
func (m *ListAllLeavesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListAllLeavesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListAllLeavesResponse proto.InternalMessageInfo

func (m *ListAllLeavesResponse) GetLeave() []*Leave {
	if m != nil {
		return m.Leave
	}
	return nil
}

type ListStudentLeaveRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *ListStudentLeaveRequest) Reset()         { *m = ListStudentLeaveRequest{} }
func (m *ListStudentLeaveRequest) String() string { return proto.CompactTextString(m) }
func (*ListStudentLeaveRequest) ProtoMessage()    {}
func (*ListStudentLeaveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f325f862ed9d7a2f, []int{4}
}
func (m *ListStudentLeaveRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListStudentLeaveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListStudentLeaveRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ListStudentLeaveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListStudentLeaveRequest.Merge(m, src)
}
func (m *ListStudentLeaveRequest) XXX_Size() int {
	return m.Size()
}
func (m *ListStudentLeaveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListStudentLeaveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListStudentLeaveRequest proto.InternalMessageInfo

func (m *ListStudentLeaveRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ListStudentLeaveResponse struct {
	Leave []*Leave `protobuf:"bytes,1,rep,name=leave,proto3" json:"leave,omitempty"`
}

func (m *ListStudentLeaveResponse) Reset()         { *m = ListStudentLeaveResponse{} }
func (m *ListStudentLeaveResponse) String() string { return proto.CompactTextString(m) }
func (*ListStudentLeaveResponse) ProtoMessage()    {}
func (*ListStudentLeaveResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f325f862ed9d7a2f, []int{5}
}
func (m *ListStudentLeaveResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListStudentLeaveResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListStudentLeaveResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ListStudentLeaveResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListStudentLeaveResponse.Merge(m, src)
}
func (m *ListStudentLeaveResponse) XXX_Size() int {
	return m.Size()
}
func (m *ListStudentLeaveResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListStudentLeaveResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListStudentLeaveResponse proto.InternalMessageInfo

func (m *ListStudentLeaveResponse) GetLeave() []*Leave {
	if m != nil {
		return m.Leave
	}
	return nil
}

func init() {
	proto.RegisterType((*GetStudentRequest)(nil), "cosmos.leave.v1.GetStudentRequest")
	proto.RegisterType((*GetStudentResponse)(nil), "cosmos.leave.v1.GetStudentResponse")
	proto.RegisterType((*ListAllLeavesRequest)(nil), "cosmos.leave.v1.ListAllLeavesRequest")
	proto.RegisterType((*ListAllLeavesResponse)(nil), "cosmos.leave.v1.ListAllLeavesResponse")
	proto.RegisterType((*ListStudentLeaveRequest)(nil), "cosmos.leave.v1.ListStudentLeaveRequest")
	proto.RegisterType((*ListStudentLeaveResponse)(nil), "cosmos.leave.v1.ListStudentLeaveResponse")
}

func init() { proto.RegisterFile("cosmos/leave/v1/query.proto", fileDescriptor_f325f862ed9d7a2f) }

var fileDescriptor_f325f862ed9d7a2f = []byte{
	// 318 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4e, 0xce, 0x2f, 0xce,
	0xcd, 0x2f, 0xd6, 0xcf, 0x49, 0x4d, 0x2c, 0x4b, 0xd5, 0x2f, 0x33, 0xd4, 0x2f, 0x2c, 0x4d, 0x2d,
	0xaa, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x87, 0x48, 0xea, 0x81, 0x25, 0xf5, 0xca,
	0x0c, 0xa5, 0x30, 0x54, 0x17, 0x97, 0x24, 0x96, 0xa4, 0x42, 0x54, 0x2b, 0x29, 0x73, 0x09, 0xba,
	0xa7, 0x96, 0x04, 0x97, 0x94, 0xa6, 0xa4, 0xe6, 0x95, 0x04, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97,
	0x08, 0xf1, 0x71, 0x31, 0x65, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x31, 0x65, 0xa6,
	0x28, 0x79, 0x70, 0x09, 0x21, 0x2b, 0x2a, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x15, 0x32, 0xe2, 0x62,
	0x2f, 0x86, 0x08, 0x81, 0x95, 0x72, 0x1b, 0x49, 0xe8, 0xa1, 0x59, 0xad, 0x07, 0xd3, 0x02, 0x53,
	0xa8, 0x24, 0xc6, 0x25, 0xe2, 0x93, 0x59, 0x5c, 0xe2, 0x98, 0x93, 0xe3, 0x03, 0x52, 0x53, 0x0c,
	0xb5, 0x51, 0xc9, 0x95, 0x4b, 0x14, 0x4d, 0x1c, 0x6a, 0x89, 0x0e, 0x17, 0x2b, 0xd8, 0x34, 0x09,
	0x46, 0x05, 0x66, 0x0d, 0x6e, 0x23, 0x31, 0x0c, 0x2b, 0xc0, 0xea, 0x83, 0x20, 0x8a, 0x94, 0x34,
	0xb9, 0xc4, 0x41, 0xc6, 0x40, 0xad, 0x85, 0x48, 0xe1, 0xf4, 0x93, 0x04, 0xa6, 0x52, 0x72, 0x2c,
	0x35, 0x5a, 0xc9, 0xc4, 0xc5, 0x11, 0x08, 0x8a, 0x00, 0xdf, 0xe2, 0x74, 0xa1, 0x50, 0x2e, 0x2e,
	0x44, 0x50, 0x09, 0x29, 0x61, 0xe8, 0xc4, 0x08, 0x6c, 0x29, 0x65, 0xbc, 0x6a, 0xa0, 0x2e, 0x8a,
	0xe5, 0xe2, 0x41, 0x0e, 0x1f, 0x21, 0x55, 0x4c, 0x27, 0x61, 0x09, 0x56, 0x29, 0x35, 0x42, 0xca,
	0xa0, 0xc6, 0xa7, 0x73, 0x09, 0xa0, 0x07, 0x86, 0x90, 0x06, 0x56, 0xbd, 0x58, 0x82, 0x56, 0x4a,
	0x93, 0x08, 0x95, 0x10, 0x8b, 0x9c, 0xd4, 0x4f, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1,
	0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e,
	0x21, 0x8a, 0xb7, 0x02, 0x9a, 0x40, 0x4b, 0x2a, 0x0b, 0x52, 0x8b, 0x93, 0xd8, 0xc0, 0xc9, 0xd3,
	0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xfb, 0xc6, 0xa2, 0x4f, 0xeb, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryMsgClient is the client API for QueryMsg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryMsgClient interface {
	GetStudent(ctx context.Context, in *GetStudentRequest, opts ...grpc.CallOption) (*GetStudentResponse, error)
	ListAllLeave(ctx context.Context, in *ListAllLeavesRequest, opts ...grpc.CallOption) (*ListAllLeavesResponse, error)
	ListStudentLeave(ctx context.Context, in *ListStudentLeaveRequest, opts ...grpc.CallOption) (*ListStudentLeaveResponse, error)
}

type queryMsgClient struct {
	cc grpc1.ClientConn
}

func NewQueryMsgClient(cc grpc1.ClientConn) QueryMsgClient {
	return &queryMsgClient{cc}
}

func (c *queryMsgClient) GetStudent(ctx context.Context, in *GetStudentRequest, opts ...grpc.CallOption) (*GetStudentResponse, error) {
	out := new(GetStudentResponse)
	err := c.cc.Invoke(ctx, "/cosmos.leave.v1.QueryMsg/GetStudent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryMsgClient) ListAllLeave(ctx context.Context, in *ListAllLeavesRequest, opts ...grpc.CallOption) (*ListAllLeavesResponse, error) {
	out := new(ListAllLeavesResponse)
	err := c.cc.Invoke(ctx, "/cosmos.leave.v1.QueryMsg/ListAllLeave", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryMsgClient) ListStudentLeave(ctx context.Context, in *ListStudentLeaveRequest, opts ...grpc.CallOption) (*ListStudentLeaveResponse, error) {
	out := new(ListStudentLeaveResponse)
	err := c.cc.Invoke(ctx, "/cosmos.leave.v1.QueryMsg/ListStudentLeave", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryMsgServer is the server API for QueryMsg service.
type QueryMsgServer interface {
	GetStudent(context.Context, *GetStudentRequest) (*GetStudentResponse, error)
	ListAllLeave(context.Context, *ListAllLeavesRequest) (*ListAllLeavesResponse, error)
	ListStudentLeave(context.Context, *ListStudentLeaveRequest) (*ListStudentLeaveResponse, error)
}

// UnimplementedQueryMsgServer can be embedded to have forward compatible implementations.
type UnimplementedQueryMsgServer struct {
}

func (*UnimplementedQueryMsgServer) GetStudent(ctx context.Context, req *GetStudentRequest) (*GetStudentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStudent not implemented")
}
func (*UnimplementedQueryMsgServer) ListAllLeave(ctx context.Context, req *ListAllLeavesRequest) (*ListAllLeavesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAllLeave not implemented")
}
func (*UnimplementedQueryMsgServer) ListStudentLeave(ctx context.Context, req *ListStudentLeaveRequest) (*ListStudentLeaveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListStudentLeave not implemented")
}

func RegisterQueryMsgServer(s grpc1.Server, srv QueryMsgServer) {
	s.RegisterService(&_QueryMsg_serviceDesc, srv)
}

func _QueryMsg_GetStudent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStudentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryMsgServer).GetStudent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.leave.v1.QueryMsg/GetStudent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryMsgServer).GetStudent(ctx, req.(*GetStudentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueryMsg_ListAllLeave_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAllLeavesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryMsgServer).ListAllLeave(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.leave.v1.QueryMsg/ListAllLeave",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryMsgServer).ListAllLeave(ctx, req.(*ListAllLeavesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueryMsg_ListStudentLeave_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListStudentLeaveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryMsgServer).ListStudentLeave(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.leave.v1.QueryMsg/ListStudentLeave",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryMsgServer).ListStudentLeave(ctx, req.(*ListStudentLeaveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _QueryMsg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.leave.v1.QueryMsg",
	HandlerType: (*QueryMsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStudent",
			Handler:    _QueryMsg_GetStudent_Handler,
		},
		{
			MethodName: "ListAllLeave",
			Handler:    _QueryMsg_ListAllLeave_Handler,
		},
		{
			MethodName: "ListStudentLeave",
			Handler:    _QueryMsg_ListStudentLeave_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmos/leave/v1/query.proto",
}

func (m *GetStudentRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetStudentRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetStudentRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GetStudentResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetStudentResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetStudentResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Student != nil {
		{
			size, err := m.Student.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ListAllLeavesRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListAllLeavesRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ListAllLeavesRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *ListAllLeavesResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListAllLeavesResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ListAllLeavesResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Leave) > 0 {
		for iNdEx := len(m.Leave) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Leave[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *ListStudentLeaveRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListStudentLeaveRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ListStudentLeaveRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ListStudentLeaveResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListStudentLeaveResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ListStudentLeaveResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Leave) > 0 {
		for iNdEx := len(m.Leave) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Leave[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GetStudentRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *GetStudentResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Student != nil {
		l = m.Student.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *ListAllLeavesRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *ListAllLeavesResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Leave) > 0 {
		for _, e := range m.Leave {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func (m *ListStudentLeaveRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *ListStudentLeaveResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Leave) > 0 {
		for _, e := range m.Leave {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GetStudentRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetStudentRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetStudentRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GetStudentResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetStudentResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetStudentResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Student", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Student == nil {
				m.Student = &Student{}
			}
			if err := m.Student.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ListAllLeavesRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ListAllLeavesRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListAllLeavesRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ListAllLeavesResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ListAllLeavesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListAllLeavesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Leave", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Leave = append(m.Leave, &Leave{})
			if err := m.Leave[len(m.Leave)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ListStudentLeaveRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ListStudentLeaveRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListStudentLeaveRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ListStudentLeaveResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ListStudentLeaveResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListStudentLeaveResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Leave", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Leave = append(m.Leave, &Leave{})
			if err := m.Leave[len(m.Leave)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
