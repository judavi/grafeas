// Code generated by protoc-gen-go. DO NOT EDIT.
// source: iris.proto

package iris_go_proto

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

// A note that indicates the entitlements on UEM
type Iris struct {
	Details              *Details `protobuf:"bytes,1,opt,name=details,proto3" json:"details,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Iris) Reset()         { *m = Iris{} }
func (m *Iris) String() string { return proto.CompactTextString(m) }
func (*Iris) ProtoMessage()    {}
func (*Iris) Descriptor() ([]byte, []int) {
	return fileDescriptor_7457d0ebad74ef5a, []int{0}
}

func (m *Iris) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Iris.Unmarshal(m, b)
}
func (m *Iris) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Iris.Marshal(b, m, deterministic)
}
func (m *Iris) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Iris.Merge(m, src)
}
func (m *Iris) XXX_Size() int {
	return xxx_messageInfo_Iris.Size(m)
}
func (m *Iris) XXX_DiscardUnknown() {
	xxx_messageInfo_Iris.DiscardUnknown(m)
}

var xxx_messageInfo_Iris proto.InternalMessageInfo

func (m *Iris) GetDetails() *Details {
	if m != nil {
		return m.Details
	}
	return nil
}

// Details of a Citi occurrence.
type Details struct {
	// Required. Server name
	ServerName string `protobuf:"bytes,1,opt,name=server_name,json=serverName,proto3" json:"server_name,omitempty"`
	// Required. Instance name
	InstanceName string `protobuf:"bytes,2,opt,name=instance_name,json=instanceName,proto3" json:"instance_name,omitempty"`
	// Required.  Port
	Port int32 `protobuf:"varint,3,opt,name=port,proto3" json:"port,omitempty"`
	// Required. APP CSI ID
	Csiid string `protobuf:"bytes,4,opt,name=csiid,proto3" json:"csiid,omitempty"`
	// Required. Where the server was found
	ServerFoundIn string `protobuf:"bytes,5,opt,name=server_found_in,json=serverFoundIn,proto3" json:"server_found_in,omitempty"`
	// Required. Version extracted from
	VersionExtractedFrom string `protobuf:"bytes,6,opt,name=version_extracted_from,json=versionExtractedFrom,proto3" json:"version_extracted_from,omitempty"`
	// Required. Date scanned DD/MM/YYYY
	DateScanned string `protobuf:"bytes,7,opt,name=date_scanned,json=dateScanned,proto3" json:"date_scanned,omitempty"`
	// Required. CAS farm name
	CasFarmName string `protobuf:"bytes,8,opt,name=cas_farm_name,json=casFarmName,proto3" json:"cas_farm_name,omitempty"`
	// Required. Sector
	Sector string `protobuf:"bytes,9,opt,name=sector,proto3" json:"sector,omitempty"`
	// Required. Environment
	Env string `protobuf:"bytes,10,opt,name=env,proto3" json:"env,omitempty"`
	// Required. DMZ
	Dmz string `protobuf:"bytes,11,opt,name=dmz,proto3" json:"dmz,omitempty"`
	// Required. Support Distribution List
	MwSupportGroupDl     string   `protobuf:"bytes,12,opt,name=mw_support_group_dl,json=mwSupportGroupDl,proto3" json:"mw_support_group_dl,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Details) Reset()         { *m = Details{} }
func (m *Details) String() string { return proto.CompactTextString(m) }
func (*Details) ProtoMessage()    {}
func (*Details) Descriptor() ([]byte, []int) {
	return fileDescriptor_7457d0ebad74ef5a, []int{1}
}

func (m *Details) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Details.Unmarshal(m, b)
}
func (m *Details) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Details.Marshal(b, m, deterministic)
}
func (m *Details) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Details.Merge(m, src)
}
func (m *Details) XXX_Size() int {
	return xxx_messageInfo_Details.Size(m)
}
func (m *Details) XXX_DiscardUnknown() {
	xxx_messageInfo_Details.DiscardUnknown(m)
}

var xxx_messageInfo_Details proto.InternalMessageInfo

func (m *Details) GetServerName() string {
	if m != nil {
		return m.ServerName
	}
	return ""
}

func (m *Details) GetInstanceName() string {
	if m != nil {
		return m.InstanceName
	}
	return ""
}

func (m *Details) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *Details) GetCsiid() string {
	if m != nil {
		return m.Csiid
	}
	return ""
}

func (m *Details) GetServerFoundIn() string {
	if m != nil {
		return m.ServerFoundIn
	}
	return ""
}

func (m *Details) GetVersionExtractedFrom() string {
	if m != nil {
		return m.VersionExtractedFrom
	}
	return ""
}

func (m *Details) GetDateScanned() string {
	if m != nil {
		return m.DateScanned
	}
	return ""
}

func (m *Details) GetCasFarmName() string {
	if m != nil {
		return m.CasFarmName
	}
	return ""
}

func (m *Details) GetSector() string {
	if m != nil {
		return m.Sector
	}
	return ""
}

func (m *Details) GetEnv() string {
	if m != nil {
		return m.Env
	}
	return ""
}

func (m *Details) GetDmz() string {
	if m != nil {
		return m.Dmz
	}
	return ""
}

func (m *Details) GetMwSupportGroupDl() string {
	if m != nil {
		return m.MwSupportGroupDl
	}
	return ""
}

func init() {
	proto.RegisterType((*Iris)(nil), "grafeas.v1beta1.iris.Iris")
	proto.RegisterType((*Details)(nil), "grafeas.v1beta1.iris.Details")
}

func init() { proto.RegisterFile("iris.proto", fileDescriptor_7457d0ebad74ef5a) }

var fileDescriptor_7457d0ebad74ef5a = []byte{
	// 386 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xcf, 0x8e, 0xd3, 0x30,
	0x10, 0x87, 0x95, 0x4d, 0xff, 0xb0, 0x93, 0x56, 0xac, 0x4c, 0xb5, 0xf8, 0x82, 0x58, 0x8a, 0x84,
	0xf6, 0x42, 0xaa, 0x05, 0x04, 0x47, 0x04, 0x5a, 0xba, 0xea, 0x05, 0xa1, 0xf4, 0x04, 0x17, 0xcb,
	0xb5, 0x9d, 0x60, 0xa9, 0xb6, 0x23, 0xdb, 0x49, 0x51, 0xdf, 0x80, 0xd7, 0xe0, 0x49, 0x91, 0xed,
	0x74, 0x4f, 0x3d, 0x65, 0xe6, 0xfb, 0x7d, 0x93, 0x68, 0x62, 0x03, 0x48, 0x2b, 0x5d, 0xd9, 0x5a,
	0xe3, 0x0d, 0x5a, 0x34, 0x96, 0xd6, 0x82, 0xba, 0xb2, 0xbf, 0xdb, 0x09, 0x4f, 0xef, 0xca, 0x90,
	0x2d, 0x3f, 0xc3, 0x68, 0x63, 0xa5, 0x43, 0x9f, 0x60, 0xca, 0x85, 0xa7, 0x72, 0xef, 0x70, 0x76,
	0x93, 0xdd, 0x16, 0xef, 0x5e, 0x94, 0xe7, 0xfc, 0xf2, 0x3e, 0x49, 0xd5, 0xc9, 0x5e, 0xfe, 0xcd,
	0x61, 0x3a, 0x40, 0xf4, 0x12, 0x0a, 0x27, 0x6c, 0x2f, 0x2c, 0xd1, 0x54, 0x89, 0xf8, 0xa2, 0xcb,
	0x0a, 0x12, 0xfa, 0x4e, 0x95, 0x40, 0xaf, 0x61, 0x2e, 0xb5, 0xf3, 0x54, 0x33, 0x91, 0x94, 0x8b,
	0xa8, 0xcc, 0x4e, 0x30, 0x4a, 0x08, 0x46, 0xad, 0xb1, 0x1e, 0xe7, 0x37, 0xd9, 0xed, 0xb8, 0x8a,
	0x35, 0x5a, 0xc0, 0x98, 0x39, 0x29, 0x39, 0x1e, 0xc5, 0x81, 0xd4, 0xa0, 0x37, 0xf0, 0x74, 0xf8,
	0x5e, 0x6d, 0x3a, 0xcd, 0x89, 0xd4, 0x78, 0x1c, 0xf3, 0x79, 0xc2, 0xeb, 0x40, 0x37, 0x1a, 0x7d,
	0x80, 0xeb, 0x5e, 0x58, 0x27, 0x8d, 0x26, 0xe2, 0x8f, 0xb7, 0x94, 0x79, 0xc1, 0x49, 0x6d, 0x8d,
	0xc2, 0x93, 0xa8, 0x2f, 0x86, 0xf4, 0xdb, 0x29, 0x5c, 0x5b, 0xa3, 0xd0, 0x2b, 0x98, 0x71, 0xea,
	0x05, 0x71, 0x8c, 0x6a, 0x2d, 0x38, 0x9e, 0x46, 0xb7, 0x08, 0x6c, 0x9b, 0x10, 0x5a, 0xc2, 0x9c,
	0x51, 0x47, 0x6a, 0x6a, 0x55, 0xda, 0xe7, 0x49, 0x72, 0x18, 0x75, 0x6b, 0x6a, 0x55, 0x5c, 0xe7,
	0x1a, 0x26, 0x4e, 0x30, 0x6f, 0x2c, 0xbe, 0x8c, 0xe1, 0xd0, 0xa1, 0x2b, 0xc8, 0x85, 0xee, 0x31,
	0x44, 0x18, 0xca, 0x40, 0xb8, 0x3a, 0xe2, 0x22, 0x11, 0xae, 0x8e, 0xe8, 0x2d, 0x3c, 0x53, 0x07,
	0xe2, 0xba, 0x36, 0xfc, 0x04, 0xd2, 0x58, 0xd3, 0xb5, 0x84, 0xef, 0xf1, 0x2c, 0x1a, 0x57, 0xea,
	0xb0, 0x4d, 0xc9, 0x43, 0x08, 0xee, 0xf7, 0x5f, 0x7f, 0xc2, 0x73, 0x69, 0xce, 0x9e, 0xdb, 0x8f,
	0xec, 0xd7, 0xc7, 0x46, 0xfa, 0xdf, 0xdd, 0xae, 0x64, 0x46, 0xad, 0x06, 0xe5, 0xf1, 0x19, 0x6f,
	0xc8, 0x6a, 0x18, 0x58, 0x85, 0x01, 0xd2, 0x18, 0x12, 0xe9, 0xbf, 0x8b, 0xfc, 0xa1, 0xfa, 0xb2,
	0x9b, 0xc4, 0xe6, 0xfd, 0xff, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6d, 0xdd, 0xd1, 0x94, 0x52, 0x02,
	0x00, 0x00,
}
