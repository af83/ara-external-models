// Code generated by protoc-gen-go. DO NOT EDIT.
// source: vehicle_journeys.proto

package external_models

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type ExternalVehicleJourney struct {
	Objectid             string   `protobuf:"bytes,1,opt,name=objectid,proto3" json:"objectid,omitempty"`
	LineRef              string   `protobuf:"bytes,2,opt,name=lineRef,proto3" json:"lineRef,omitempty"`
	OriginRef            string   `protobuf:"bytes,3,opt,name=originRef,proto3" json:"originRef,omitempty"`
	OriginName           string   `protobuf:"bytes,4,opt,name=originName,proto3" json:"originName,omitempty"`
	DestinationRef       string   `protobuf:"bytes,5,opt,name=destinationRef,proto3" json:"destinationRef,omitempty"`
	DestinationName      string   `protobuf:"bytes,6,opt,name=destinationName,proto3" json:"destinationName,omitempty"`
	Direction            string   `protobuf:"bytes,7,opt,name=direction,proto3" json:"direction,omitempty"`
	ViaName              string   `protobuf:"bytes,8,opt,name=viaName,proto3" json:"viaName,omitempty"`
	Advance              int32    `protobuf:"varint,9,opt,name=advance,proto3" json:"advance,omitempty"`
	Delay                int32    `protobuf:"varint,10,opt,name=delay,proto3" json:"delay,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExternalVehicleJourney) Reset()         { *m = ExternalVehicleJourney{} }
func (m *ExternalVehicleJourney) String() string { return proto.CompactTextString(m) }
func (*ExternalVehicleJourney) ProtoMessage()    {}
func (*ExternalVehicleJourney) Descriptor() ([]byte, []int) {
	return fileDescriptor_f2317bc119de9bd2, []int{0}
}

func (m *ExternalVehicleJourney) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExternalVehicleJourney.Unmarshal(m, b)
}
func (m *ExternalVehicleJourney) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExternalVehicleJourney.Marshal(b, m, deterministic)
}
func (m *ExternalVehicleJourney) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExternalVehicleJourney.Merge(m, src)
}
func (m *ExternalVehicleJourney) XXX_Size() int {
	return xxx_messageInfo_ExternalVehicleJourney.Size(m)
}
func (m *ExternalVehicleJourney) XXX_DiscardUnknown() {
	xxx_messageInfo_ExternalVehicleJourney.DiscardUnknown(m)
}

var xxx_messageInfo_ExternalVehicleJourney proto.InternalMessageInfo

func (m *ExternalVehicleJourney) GetObjectid() string {
	if m != nil {
		return m.Objectid
	}
	return ""
}

func (m *ExternalVehicleJourney) GetLineRef() string {
	if m != nil {
		return m.LineRef
	}
	return ""
}

func (m *ExternalVehicleJourney) GetOriginRef() string {
	if m != nil {
		return m.OriginRef
	}
	return ""
}

func (m *ExternalVehicleJourney) GetOriginName() string {
	if m != nil {
		return m.OriginName
	}
	return ""
}

func (m *ExternalVehicleJourney) GetDestinationRef() string {
	if m != nil {
		return m.DestinationRef
	}
	return ""
}

func (m *ExternalVehicleJourney) GetDestinationName() string {
	if m != nil {
		return m.DestinationName
	}
	return ""
}

func (m *ExternalVehicleJourney) GetDirection() string {
	if m != nil {
		return m.Direction
	}
	return ""
}

func (m *ExternalVehicleJourney) GetViaName() string {
	if m != nil {
		return m.ViaName
	}
	return ""
}

func (m *ExternalVehicleJourney) GetAdvance() int32 {
	if m != nil {
		return m.Advance
	}
	return 0
}

func (m *ExternalVehicleJourney) GetDelay() int32 {
	if m != nil {
		return m.Delay
	}
	return 0
}

type ExternalVehicleJourneys struct {
	VehicleJourneys      []*ExternalVehicleJourney `protobuf:"bytes,1,rep,name=vehicle_journeys,json=vehicleJourneys,proto3" json:"vehicle_journeys,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *ExternalVehicleJourneys) Reset()         { *m = ExternalVehicleJourneys{} }
func (m *ExternalVehicleJourneys) String() string { return proto.CompactTextString(m) }
func (*ExternalVehicleJourneys) ProtoMessage()    {}
func (*ExternalVehicleJourneys) Descriptor() ([]byte, []int) {
	return fileDescriptor_f2317bc119de9bd2, []int{1}
}

func (m *ExternalVehicleJourneys) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExternalVehicleJourneys.Unmarshal(m, b)
}
func (m *ExternalVehicleJourneys) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExternalVehicleJourneys.Marshal(b, m, deterministic)
}
func (m *ExternalVehicleJourneys) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExternalVehicleJourneys.Merge(m, src)
}
func (m *ExternalVehicleJourneys) XXX_Size() int {
	return xxx_messageInfo_ExternalVehicleJourneys.Size(m)
}
func (m *ExternalVehicleJourneys) XXX_DiscardUnknown() {
	xxx_messageInfo_ExternalVehicleJourneys.DiscardUnknown(m)
}

var xxx_messageInfo_ExternalVehicleJourneys proto.InternalMessageInfo

func (m *ExternalVehicleJourneys) GetVehicleJourneys() []*ExternalVehicleJourney {
	if m != nil {
		return m.VehicleJourneys
	}
	return nil
}

func init() {
	proto.RegisterType((*ExternalVehicleJourney)(nil), "external_models.ExternalVehicleJourney")
	proto.RegisterType((*ExternalVehicleJourneys)(nil), "external_models.ExternalVehicleJourneys")
}

func init() { proto.RegisterFile("vehicle_journeys.proto", fileDescriptor_f2317bc119de9bd2) }

var fileDescriptor_f2317bc119de9bd2 = []byte{
	// 271 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x31, 0x4f, 0xc3, 0x30,
	0x10, 0x85, 0x95, 0x96, 0xb4, 0xcd, 0x21, 0x11, 0x64, 0xa1, 0x62, 0x21, 0x84, 0xa2, 0x0e, 0x90,
	0x29, 0x03, 0xfc, 0x06, 0x16, 0x06, 0x86, 0x0c, 0xac, 0x95, 0x1b, 0x1f, 0xe0, 0xca, 0xb1, 0x91,
	0x63, 0x22, 0xfa, 0x07, 0xf9, 0x5d, 0xc8, 0xe7, 0xb6, 0xb4, 0x51, 0xc7, 0xf7, 0x3d, 0xdf, 0xd3,
	0x9d, 0x1f, 0xcc, 0x7b, 0xfc, 0x54, 0x8d, 0xc6, 0xe5, 0xda, 0x7e, 0x3b, 0x83, 0x9b, 0xae, 0xfa,
	0x72, 0xd6, 0x5b, 0x96, 0xe3, 0x8f, 0x47, 0x67, 0x84, 0x5e, 0xb6, 0x56, 0xa2, 0xee, 0x16, 0xbf,
	0x23, 0x98, 0x3f, 0x6f, 0xd9, 0x5b, 0x9c, 0x79, 0x89, 0x23, 0xec, 0x06, 0x66, 0x76, 0xb5, 0xc6,
	0xc6, 0x2b, 0xc9, 0x93, 0x22, 0x29, 0xb3, 0x7a, 0xaf, 0x19, 0x87, 0xa9, 0x56, 0x06, 0x6b, 0x7c,
	0xe7, 0x23, 0xb2, 0x76, 0x92, 0xdd, 0x42, 0x66, 0x9d, 0xfa, 0x50, 0x26, 0x78, 0x63, 0xf2, 0xfe,
	0x01, 0xbb, 0x03, 0x88, 0xe2, 0x55, 0xb4, 0xc8, 0xcf, 0xc8, 0x3e, 0x20, 0xec, 0x1e, 0x2e, 0x24,
	0x76, 0x5e, 0x19, 0xe1, 0x95, 0xa5, 0x88, 0x94, 0xde, 0x0c, 0x28, 0x2b, 0x21, 0x3f, 0x20, 0x14,
	0x36, 0xa1, 0x87, 0x43, 0x1c, 0xf6, 0x91, 0xca, 0x85, 0xad, 0xad, 0xe1, 0xd3, 0xb8, 0xcf, 0x1e,
	0x84, 0x3b, 0x7a, 0x25, 0x68, 0x7e, 0x16, 0xef, 0xd8, 0xca, 0xe0, 0x08, 0xd9, 0x0b, 0xd3, 0x20,
	0xcf, 0x8a, 0xa4, 0x4c, 0xeb, 0x9d, 0x64, 0x57, 0x90, 0x4a, 0xd4, 0x62, 0xc3, 0x81, 0x78, 0x14,
	0x8b, 0x16, 0xae, 0x4f, 0xff, 0x63, 0xc7, 0x6a, 0xb8, 0x1c, 0xd6, 0xc1, 0x93, 0x62, 0x5c, 0x9e,
	0x3f, 0x3e, 0x54, 0x83, 0x3e, 0xaa, 0xd3, 0x19, 0x75, 0xde, 0x1f, 0x67, 0xae, 0x26, 0xd4, 0xe7,
	0xd3, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x17, 0x9c, 0x56, 0xb6, 0xe9, 0x01, 0x00, 0x00,
}
