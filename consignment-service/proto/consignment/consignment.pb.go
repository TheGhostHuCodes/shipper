// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/consignment/consignment.proto

/*
Package go_micro_svc_consignment is a generated protocol buffer package.

It is generated from these files:
	proto/consignment/consignment.proto

It has these top-level messages:
	Consignment
	Container
	GetRequest
	Response
*/
package go_micro_svc_consignment

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
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

type Consignment struct {
	Id         string       `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Descrption string       `protobuf:"bytes,2,opt,name=descrption" json:"descrption,omitempty"`
	Weight     int32        `protobuf:"varint,3,opt,name=weight" json:"weight,omitempty"`
	Containers []*Container `protobuf:"bytes,4,rep,name=containers" json:"containers,omitempty"`
	VesselId   string       `protobuf:"bytes,5,opt,name=vessel_id,json=vesselId" json:"vessel_id,omitempty"`
}

func (m *Consignment) Reset()                    { *m = Consignment{} }
func (m *Consignment) String() string            { return proto.CompactTextString(m) }
func (*Consignment) ProtoMessage()               {}
func (*Consignment) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Consignment) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Consignment) GetDescrption() string {
	if m != nil {
		return m.Descrption
	}
	return ""
}

func (m *Consignment) GetWeight() int32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *Consignment) GetContainers() []*Container {
	if m != nil {
		return m.Containers
	}
	return nil
}

func (m *Consignment) GetVesselId() string {
	if m != nil {
		return m.VesselId
	}
	return ""
}

type Container struct {
	Id         string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	CustomerId string `protobuf:"bytes,2,opt,name=customer_id,json=customerId" json:"customer_id,omitempty"`
	Origin     string `protobuf:"bytes,3,opt,name=origin" json:"origin,omitempty"`
	UserId     string `protobuf:"bytes,4,opt,name=user_id,json=userId" json:"user_id,omitempty"`
}

func (m *Container) Reset()                    { *m = Container{} }
func (m *Container) String() string            { return proto.CompactTextString(m) }
func (*Container) ProtoMessage()               {}
func (*Container) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Container) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Container) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *Container) GetOrigin() string {
	if m != nil {
		return m.Origin
	}
	return ""
}

func (m *Container) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type GetRequest struct {
}

func (m *GetRequest) Reset()                    { *m = GetRequest{} }
func (m *GetRequest) String() string            { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()               {}
func (*GetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type Response struct {
	Created      bool           `protobuf:"varint,1,opt,name=created" json:"created,omitempty"`
	Consignment  *Consignment   `protobuf:"bytes,2,opt,name=consignment" json:"consignment,omitempty"`
	Consignments []*Consignment `protobuf:"bytes,3,rep,name=consignments" json:"consignments,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Response) GetCreated() bool {
	if m != nil {
		return m.Created
	}
	return false
}

func (m *Response) GetConsignment() *Consignment {
	if m != nil {
		return m.Consignment
	}
	return nil
}

func (m *Response) GetConsignments() []*Consignment {
	if m != nil {
		return m.Consignments
	}
	return nil
}

func init() {
	proto.RegisterType((*Consignment)(nil), "go.micro.svc.consignment.Consignment")
	proto.RegisterType((*Container)(nil), "go.micro.svc.consignment.Container")
	proto.RegisterType((*GetRequest)(nil), "go.micro.svc.consignment.GetRequest")
	proto.RegisterType((*Response)(nil), "go.micro.svc.consignment.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for ShippingService service

type ShippingServiceClient interface {
	CreateConsignment(ctx context.Context, in *Consignment, opts ...client.CallOption) (*Response, error)
	GetConsignments(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*Response, error)
}

type shippingServiceClient struct {
	c           client.Client
	serviceName string
}

func NewShippingServiceClient(serviceName string, c client.Client) ShippingServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "go.micro.svc.consignment"
	}
	return &shippingServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *shippingServiceClient) CreateConsignment(ctx context.Context, in *Consignment, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "ShippingService.CreateConsignment", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shippingServiceClient) GetConsignments(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "ShippingService.GetConsignments", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ShippingService service

type ShippingServiceHandler interface {
	CreateConsignment(context.Context, *Consignment, *Response) error
	GetConsignments(context.Context, *GetRequest, *Response) error
}

func RegisterShippingServiceHandler(s server.Server, hdlr ShippingServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&ShippingService{hdlr}, opts...))
}

type ShippingService struct {
	ShippingServiceHandler
}

func (h *ShippingService) CreateConsignment(ctx context.Context, in *Consignment, out *Response) error {
	return h.ShippingServiceHandler.CreateConsignment(ctx, in, out)
}

func (h *ShippingService) GetConsignments(ctx context.Context, in *GetRequest, out *Response) error {
	return h.ShippingServiceHandler.GetConsignments(ctx, in, out)
}

func init() { proto.RegisterFile("proto/consignment/consignment.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 349 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xdf, 0x4a, 0xc3, 0x30,
	0x14, 0x87, 0xed, 0xfe, 0xf7, 0x74, 0x38, 0xcc, 0x85, 0x06, 0x05, 0x1d, 0x9d, 0xc2, 0xae, 0x2a,
	0xcc, 0x47, 0xe8, 0xc5, 0xe8, 0x6d, 0x76, 0x2d, 0x3a, 0xd3, 0x43, 0x17, 0xb0, 0x49, 0x4d, 0xb2,
	0xf9, 0x68, 0xfa, 0x1c, 0x3e, 0x91, 0x2c, 0x5d, 0x5d, 0x44, 0x26, 0xbb, 0xcb, 0xef, 0x9c, 0x7c,
	0xa7, 0x5f, 0x0f, 0x81, 0x49, 0xa5, 0x95, 0x55, 0xf7, 0x5c, 0x49, 0x23, 0x0a, 0x59, 0xa2, 0xb4,
	0xfe, 0x39, 0x71, 0x5d, 0x42, 0x0b, 0x95, 0x94, 0x82, 0x6b, 0x95, 0x98, 0x0d, 0x4f, 0xbc, 0x7e,
	0xfc, 0x11, 0x40, 0x94, 0xee, 0x33, 0x39, 0x85, 0x96, 0xc8, 0x69, 0x30, 0x0e, 0xa6, 0x21, 0x6b,
	0x89, 0x9c, 0x5c, 0x03, 0xe4, 0x68, 0xb8, 0xae, 0xac, 0x50, 0x92, 0xb6, 0x5c, 0xdd, 0xab, 0x90,
	0x73, 0xe8, 0xbd, 0xa3, 0x28, 0x56, 0x96, 0xb6, 0xc7, 0xc1, 0xb4, 0xcb, 0x76, 0x89, 0xa4, 0x00,
	0x5c, 0x49, 0xbb, 0x14, 0x12, 0xb5, 0xa1, 0x9d, 0x71, 0x7b, 0x1a, 0xcd, 0x26, 0xc9, 0x21, 0x8d,
	0x24, 0x6d, 0xee, 0x32, 0x0f, 0x23, 0x57, 0x10, 0x6e, 0xd0, 0x18, 0x7c, 0x7d, 0x12, 0x39, 0xed,
	0xba, 0x6f, 0x0f, 0xea, 0x42, 0x96, 0xc7, 0x25, 0x84, 0x3f, 0xd4, 0x1f, 0xed, 0x1b, 0x88, 0xf8,
	0xda, 0x58, 0x55, 0xa2, 0xde, 0xb2, 0x3b, 0xef, 0xa6, 0x94, 0xe5, 0x5b, 0x6f, 0xa5, 0x45, 0x21,
	0xa4, 0xf3, 0x0e, 0xd9, 0x2e, 0x91, 0x0b, 0xe8, 0xaf, 0x4d, 0x0d, 0x75, 0xea, 0xc6, 0x36, 0x66,
	0x79, 0x3c, 0x04, 0x98, 0xa3, 0x65, 0xf8, 0xb6, 0x46, 0x63, 0xe3, 0xcf, 0x00, 0x06, 0x0c, 0x4d,
	0xa5, 0xa4, 0x41, 0x42, 0xa1, 0xcf, 0x35, 0x2e, 0x2d, 0xd6, 0x06, 0x03, 0xd6, 0x44, 0x32, 0x87,
	0xc8, 0xfb, 0x4b, 0xa7, 0x11, 0xcd, 0xee, 0xfe, 0x5d, 0x43, 0x73, 0x66, 0x3e, 0x49, 0x32, 0x18,
	0x7a, 0xd1, 0xd0, 0xb6, 0x5b, 0xe8, 0x91, 0x93, 0x7e, 0xa1, 0xb3, 0xaf, 0x00, 0x46, 0x8b, 0x95,
	0xa8, 0x2a, 0x21, 0x8b, 0x05, 0xea, 0x8d, 0xe0, 0x48, 0x9e, 0xe1, 0x2c, 0x75, 0xca, 0xfe, 0x53,
	0x38, 0x6e, 0xfa, 0x65, 0x7c, 0xf8, 0x5a, 0xb3, 0xa1, 0xf8, 0x84, 0x3c, 0xc2, 0x68, 0x8e, 0xd6,
	0xe3, 0x0c, 0xb9, 0x3d, 0x0c, 0xee, 0x37, 0x7d, 0xdc, 0xf8, 0x97, 0x9e, 0x7b, 0xe7, 0x0f, 0xdf,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x45, 0xf8, 0xff, 0xc6, 0x0e, 0x03, 0x00, 0x00,
}
