// Code generated by protoc-gen-go. DO NOT EDIT.
// source: complete_model.proto

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

type ExternalCompleteModel struct {
	StopAreas            []*ExternalStopArea       `protobuf:"bytes,1,rep,name=stop_areas,json=stopAreas,proto3" json:"stop_areas,omitempty"`
	Lines                []*ExternalLine           `protobuf:"bytes,2,rep,name=lines,proto3" json:"lines,omitempty"`
	VehicleJourneys      []*ExternalVehicleJourney `protobuf:"bytes,3,rep,name=vehicle_journeys,json=vehicleJourneys,proto3" json:"vehicle_journeys,omitempty"`
	StopVisits           []*ExternalStopVisit      `protobuf:"bytes,4,rep,name=stop_visits,json=stopVisits,proto3" json:"stop_visits,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *ExternalCompleteModel) Reset()         { *m = ExternalCompleteModel{} }
func (m *ExternalCompleteModel) String() string { return proto.CompactTextString(m) }
func (*ExternalCompleteModel) ProtoMessage()    {}
func (*ExternalCompleteModel) Descriptor() ([]byte, []int) {
	return fileDescriptor_a3658c1dab36914d, []int{0}
}

func (m *ExternalCompleteModel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExternalCompleteModel.Unmarshal(m, b)
}
func (m *ExternalCompleteModel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExternalCompleteModel.Marshal(b, m, deterministic)
}
func (m *ExternalCompleteModel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExternalCompleteModel.Merge(m, src)
}
func (m *ExternalCompleteModel) XXX_Size() int {
	return xxx_messageInfo_ExternalCompleteModel.Size(m)
}
func (m *ExternalCompleteModel) XXX_DiscardUnknown() {
	xxx_messageInfo_ExternalCompleteModel.DiscardUnknown(m)
}

var xxx_messageInfo_ExternalCompleteModel proto.InternalMessageInfo

func (m *ExternalCompleteModel) GetStopAreas() []*ExternalStopArea {
	if m != nil {
		return m.StopAreas
	}
	return nil
}

func (m *ExternalCompleteModel) GetLines() []*ExternalLine {
	if m != nil {
		return m.Lines
	}
	return nil
}

func (m *ExternalCompleteModel) GetVehicleJourneys() []*ExternalVehicleJourney {
	if m != nil {
		return m.VehicleJourneys
	}
	return nil
}

func (m *ExternalCompleteModel) GetStopVisits() []*ExternalStopVisit {
	if m != nil {
		return m.StopVisits
	}
	return nil
}

func init() {
	proto.RegisterType((*ExternalCompleteModel)(nil), "external_models.ExternalCompleteModel")
}

func init() { proto.RegisterFile("complete_model.proto", fileDescriptor_a3658c1dab36914d) }

var fileDescriptor_a3658c1dab36914d = []byte{
	// 232 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x49, 0xce, 0xcf, 0x2d,
	0xc8, 0x49, 0x2d, 0x49, 0x8d, 0xcf, 0xcd, 0x4f, 0x49, 0xcd, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0xe2, 0x4f, 0xad, 0x28, 0x49, 0x2d, 0xca, 0x4b, 0xcc, 0x81, 0x88, 0x16, 0x4b, 0x09, 0x14,
	0x97, 0xe4, 0x17, 0xc4, 0x27, 0x16, 0xa5, 0x26, 0x16, 0x43, 0x94, 0x48, 0x71, 0xe7, 0x64, 0xe6,
	0xa5, 0xc2, 0x38, 0x62, 0x65, 0xa9, 0x19, 0x99, 0xc9, 0x39, 0xa9, 0xf1, 0x59, 0xf9, 0xa5, 0x45,
	0x79, 0xa9, 0x95, 0x30, 0x71, 0x41, 0xb0, 0xb6, 0xb2, 0xcc, 0xe2, 0xcc, 0x12, 0xa8, 0x90, 0xd2,
	0x72, 0x26, 0x2e, 0x51, 0x57, 0xa8, 0xe9, 0xce, 0x50, 0xbb, 0x7d, 0x41, 0x96, 0x08, 0x39, 0x70,
	0x71, 0x21, 0x6c, 0x91, 0x60, 0x54, 0x60, 0xd6, 0xe0, 0x36, 0x52, 0xd4, 0x43, 0x73, 0x89, 0x1e,
	0x4c, 0x6f, 0x70, 0x49, 0x7e, 0x81, 0x63, 0x51, 0x6a, 0x62, 0x10, 0x67, 0x31, 0x94, 0x55, 0x2c,
	0x64, 0xcc, 0xc5, 0x0a, 0x76, 0x95, 0x04, 0x13, 0x58, 0xb3, 0x2c, 0x4e, 0xcd, 0x3e, 0x99, 0x79,
	0xa9, 0x41, 0x10, 0xb5, 0x42, 0x41, 0x5c, 0x02, 0xe8, 0xae, 0x97, 0x60, 0x06, 0xeb, 0x57, 0xc7,
	0xa9, 0x3f, 0x0c, 0xa2, 0xc1, 0x0b, 0xa2, 0x3e, 0x88, 0xbf, 0x0c, 0x85, 0x5f, 0x2c, 0xe4, 0xcc,
	0xc5, 0x8d, 0xe4, 0x73, 0x09, 0x16, 0xb0, 0x71, 0x4a, 0x78, 0xfd, 0x12, 0x06, 0x52, 0x1a, 0x04,
	0x0e, 0x01, 0x30, 0xb3, 0x38, 0x89, 0x0d, 0x1c, 0x60, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xea, 0x9f, 0x1c, 0x6f, 0xa3, 0x01, 0x00, 0x00,
}
