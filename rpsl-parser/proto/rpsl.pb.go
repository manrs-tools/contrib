// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpsl.proto

package rpsl_proto

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

// Type is the ENUM of various record types that are possible
// in standard RPSL.
type Type int32

const (
	Type_UNKNOWN     Type = 0
	Type_EOF         Type = 1
	Type_ADDRESS     Type = 16
	Type_ADMINC      Type = 17
	Type_AGGRBNDRY   Type = 18
	Type_AGGRMTD     Type = 19
	Type_ALIAS       Type = 20
	Type_ASNAME      Type = 21
	Type_ASSET       Type = 22
	Type_AUTH        Type = 23
	Type_AUTNUM      Type = 24
	Type_CERTIF      Type = 25
	Type_CHANGED     Type = 26
	Type_COMPONENTS  Type = 27
	Type_COUNTRY     Type = 28
	Type_DEFAULT     Type = 29
	Type_DESCR       Type = 30
	Type_EMAIL       Type = 31
	Type_EXPORT      Type = 32
	Type_EXPORTCOMPS Type = 33
	Type_EXPORTVIA   Type = 34
	Type_FAXNO       Type = 35
	Type_FILTER      Type = 36
	Type_FILTERSET   Type = 37
	Type_FINGERPR    Type = 39
	Type_GEOIDX      Type = 40
	Type_HOLES       Type = 41
	Type_IFADDR      Type = 42
	Type_IMPORT      Type = 43
	Type_IMPORTVIA   Type = 44
	Type_INET6NUM    Type = 45
	Type_INETNUM     Type = 46
	Type_INETRTR     Type = 47
	Type_INTERFACE   Type = 48
	Type_KEYCERT     Type = 49
	Type_LOCALAS     Type = 50
	Type_MBRSBYREF   Type = 51
	Type_MEMBEROF    Type = 52
	Type_MEMBERS     Type = 53
	Type_METHOD      Type = 54
	Type_MNTBY       Type = 55
	Type_MNTNER      Type = 56
	Type_MNTNFY      Type = 57
	Type_MPEXPORT    Type = 58
	Type_MPFILTER    Type = 59
	Type_MPIMPORT    Type = 60
	Type_MPMEMBERS   Type = 61
	Type_MPPEER      Type = 62
	Type_MPPEERING   Type = 63
	Type_NETNAME     Type = 64
	Type_NICHDL      Type = 65
	Type_NOTIFY      Type = 66
	Type_ORIGIN      Type = 67
	Type_OWNER       Type = 68
	Type_PEER        Type = 69
	Type_PEERING     Type = 70
	Type_PEERINGSET  Type = 71
	Type_PERSON      Type = 72
	Type_PHONE       Type = 73
	Type_REMARKS     Type = 74
	Type_ROAURI      Type = 75
	Type_ROLE        Type = 76
	Type_ROUTE       Type = 77
	Type_ROUTE6      Type = 78
	Type_ROUTESET    Type = 79
	Type_RSIN        Type = 80
	Type_RSOUT       Type = 81
	Type_RTRSET      Type = 82
	Type_SOURCE      Type = 83
	Type_STATUS      Type = 84
	Type_TECHC       Type = 85
	Type_TROUBLE     Type = 86
	Type_UPDTO       Type = 87
	Type_XXE         Type = 88
	Type_XXNER       Type = 89
	Type_XXNUM       Type = 90
	Type_XXRINGSET   Type = 91
	Type_XXSET       Type = 92
	Type_XXSON       Type = 93
	Type_XXTE        Type = 94
	Type_XXTE6       Type = 95
	Type_XXTESET     Type = 96
)

var Type_name = map[int32]string{
	0:  "UNKNOWN",
	1:  "EOF",
	16: "ADDRESS",
	17: "ADMINC",
	18: "AGGRBNDRY",
	19: "AGGRMTD",
	20: "ALIAS",
	21: "ASNAME",
	22: "ASSET",
	23: "AUTH",
	24: "AUTNUM",
	25: "CERTIF",
	26: "CHANGED",
	27: "COMPONENTS",
	28: "COUNTRY",
	29: "DEFAULT",
	30: "DESCR",
	31: "EMAIL",
	32: "EXPORT",
	33: "EXPORTCOMPS",
	34: "EXPORTVIA",
	35: "FAXNO",
	36: "FILTER",
	37: "FILTERSET",
	39: "FINGERPR",
	40: "GEOIDX",
	41: "HOLES",
	42: "IFADDR",
	43: "IMPORT",
	44: "IMPORTVIA",
	45: "INET6NUM",
	46: "INETNUM",
	47: "INETRTR",
	48: "INTERFACE",
	49: "KEYCERT",
	50: "LOCALAS",
	51: "MBRSBYREF",
	52: "MEMBEROF",
	53: "MEMBERS",
	54: "METHOD",
	55: "MNTBY",
	56: "MNTNER",
	57: "MNTNFY",
	58: "MPEXPORT",
	59: "MPFILTER",
	60: "MPIMPORT",
	61: "MPMEMBERS",
	62: "MPPEER",
	63: "MPPEERING",
	64: "NETNAME",
	65: "NICHDL",
	66: "NOTIFY",
	67: "ORIGIN",
	68: "OWNER",
	69: "PEER",
	70: "PEERING",
	71: "PEERINGSET",
	72: "PERSON",
	73: "PHONE",
	74: "REMARKS",
	75: "ROAURI",
	76: "ROLE",
	77: "ROUTE",
	78: "ROUTE6",
	79: "ROUTESET",
	80: "RSIN",
	81: "RSOUT",
	82: "RTRSET",
	83: "SOURCE",
	84: "STATUS",
	85: "TECHC",
	86: "TROUBLE",
	87: "UPDTO",
	88: "XXE",
	89: "XXNER",
	90: "XXNUM",
	91: "XXRINGSET",
	92: "XXSET",
	93: "XXSON",
	94: "XXTE",
	95: "XXTE6",
	96: "XXTESET",
}

var Type_value = map[string]int32{
	"UNKNOWN":     0,
	"EOF":         1,
	"ADDRESS":     16,
	"ADMINC":      17,
	"AGGRBNDRY":   18,
	"AGGRMTD":     19,
	"ALIAS":       20,
	"ASNAME":      21,
	"ASSET":       22,
	"AUTH":        23,
	"AUTNUM":      24,
	"CERTIF":      25,
	"CHANGED":     26,
	"COMPONENTS":  27,
	"COUNTRY":     28,
	"DEFAULT":     29,
	"DESCR":       30,
	"EMAIL":       31,
	"EXPORT":      32,
	"EXPORTCOMPS": 33,
	"EXPORTVIA":   34,
	"FAXNO":       35,
	"FILTER":      36,
	"FILTERSET":   37,
	"FINGERPR":    39,
	"GEOIDX":      40,
	"HOLES":       41,
	"IFADDR":      42,
	"IMPORT":      43,
	"IMPORTVIA":   44,
	"INET6NUM":    45,
	"INETNUM":     46,
	"INETRTR":     47,
	"INTERFACE":   48,
	"KEYCERT":     49,
	"LOCALAS":     50,
	"MBRSBYREF":   51,
	"MEMBEROF":    52,
	"MEMBERS":     53,
	"METHOD":      54,
	"MNTBY":       55,
	"MNTNER":      56,
	"MNTNFY":      57,
	"MPEXPORT":    58,
	"MPFILTER":    59,
	"MPIMPORT":    60,
	"MPMEMBERS":   61,
	"MPPEER":      62,
	"MPPEERING":   63,
	"NETNAME":     64,
	"NICHDL":      65,
	"NOTIFY":      66,
	"ORIGIN":      67,
	"OWNER":       68,
	"PEER":        69,
	"PEERING":     70,
	"PEERINGSET":  71,
	"PERSON":      72,
	"PHONE":       73,
	"REMARKS":     74,
	"ROAURI":      75,
	"ROLE":        76,
	"ROUTE":       77,
	"ROUTE6":      78,
	"ROUTESET":    79,
	"RSIN":        80,
	"RSOUT":       81,
	"RTRSET":      82,
	"SOURCE":      83,
	"STATUS":      84,
	"TECHC":       85,
	"TROUBLE":     86,
	"UPDTO":       87,
	"XXE":         88,
	"XXNER":       89,
	"XXNUM":       90,
	"XXRINGSET":   91,
	"XXSET":       92,
	"XXSON":       93,
	"XXTE":        94,
	"XXTE6":       95,
	"XXTESET":     96,
}

func (x Type) String() string {
	return proto.EnumName(Type_name, int32(x))
}

func (Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_cf10760c6fc5b93a, []int{0}
}

type KeyValue struct {
	Key                  Type     `protobuf:"varint,1,opt,name=key,proto3,enum=rpsl.proto.Type" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeyValue) Reset()         { *m = KeyValue{} }
func (m *KeyValue) String() string { return proto.CompactTextString(m) }
func (*KeyValue) ProtoMessage()    {}
func (*KeyValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf10760c6fc5b93a, []int{0}
}

func (m *KeyValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyValue.Unmarshal(m, b)
}
func (m *KeyValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyValue.Marshal(b, m, deterministic)
}
func (m *KeyValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyValue.Merge(m, src)
}
func (m *KeyValue) XXX_Size() int {
	return xxx_messageInfo_KeyValue.Size(m)
}
func (m *KeyValue) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyValue.DiscardUnknown(m)
}

var xxx_messageInfo_KeyValue proto.InternalMessageInfo

func (m *KeyValue) GetKey() Type {
	if m != nil {
		return m.Key
	}
	return Type_UNKNOWN
}

func (m *KeyValue) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

//
// Record, a single rpsl record.
type Record struct {
	Type                 Type        `protobuf:"varint,1,opt,name=type,proto3,enum=rpsl.proto.Type" json:"type,omitempty"`
	Fields               []*KeyValue `protobuf:"bytes,2,rep,name=fields,proto3" json:"fields,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Record) Reset()         { *m = Record{} }
func (m *Record) String() string { return proto.CompactTextString(m) }
func (*Record) ProtoMessage()    {}
func (*Record) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf10760c6fc5b93a, []int{1}
}

func (m *Record) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Record.Unmarshal(m, b)
}
func (m *Record) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Record.Marshal(b, m, deterministic)
}
func (m *Record) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Record.Merge(m, src)
}
func (m *Record) XXX_Size() int {
	return xxx_messageInfo_Record.Size(m)
}
func (m *Record) XXX_DiscardUnknown() {
	xxx_messageInfo_Record.DiscardUnknown(m)
}

var xxx_messageInfo_Record proto.InternalMessageInfo

func (m *Record) GetType() Type {
	if m != nil {
		return m.Type
	}
	return Type_UNKNOWN
}

func (m *Record) GetFields() []*KeyValue {
	if m != nil {
		return m.Fields
	}
	return nil
}

func init() {
	proto.RegisterEnum("rpsl.proto.Type", Type_name, Type_value)
	proto.RegisterType((*KeyValue)(nil), "rpsl.proto.KeyValue")
	proto.RegisterType((*Record)(nil), "rpsl.proto.Record")
}

func init() { proto.RegisterFile("rpsl.proto", fileDescriptor_cf10760c6fc5b93a) }

var fileDescriptor_cf10760c6fc5b93a = []byte{
	// 732 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x54, 0x6b, 0x73, 0xdc, 0x44,
	0x10, 0xc4, 0xb1, 0xe3, 0xd8, 0x6b, 0x08, 0x83, 0x30, 0x60, 0xde, 0xc6, 0x84, 0xc2, 0x84, 0x70,
	0x80, 0x03, 0xe6, 0xfd, 0xd8, 0x93, 0x46, 0xd2, 0x72, 0xd2, 0xae, 0x98, 0xdd, 0xb5, 0x25, 0x08,
	0xef, 0x1c, 0x55, 0x14, 0xae, 0xf2, 0x95, 0x13, 0xa8, 0xba, 0xaf, 0xfc, 0x72, 0xaa, 0x57, 0x77,
	0x45, 0xbe, 0xf0, 0xad, 0x7b, 0xb6, 0xa7, 0xb7, 0x67, 0x6e, 0x4f, 0x4a, 0x5d, 0x2d, 0x1e, 0x5c,
	0x4c, 0x16, 0x57, 0x97, 0x0f, 0x2f, 0xb3, 0x47, 0xf0, 0x51, 0xa1, 0x76, 0x66, 0xf3, 0xe5, 0xd9,
	0x2f, 0x17, 0x7f, 0xcd, 0xb3, 0x23, 0xb5, 0xf9, 0xe7, 0x7c, 0x79, 0xb0, 0x71, 0xb8, 0x71, 0x7c,
	0xf3, 0x84, 0x26, 0xff, 0xa9, 0x26, 0x61, 0xb9, 0x98, 0x0b, 0x0e, 0xb3, 0x7d, 0x75, 0xfd, 0x6f,
	0x88, 0x0f, 0xae, 0x1d, 0x6e, 0x1c, 0xef, 0xca, 0x48, 0x8e, 0xee, 0xa9, 0x6d, 0x99, 0xff, 0x76,
	0x79, 0x75, 0x3f, 0xbb, 0xa5, 0xb6, 0x1e, 0x2e, 0x17, 0xf3, 0xff, 0x35, 0x49, 0xa7, 0xd9, 0x1d,
	0xb5, 0xfd, 0xfb, 0x1f, 0xf3, 0x8b, 0xfb, 0x0f, 0x0e, 0xae, 0x1d, 0x6e, 0x1e, 0xef, 0x9d, 0xec,
	0x3f, 0xaa, 0x5b, 0xe7, 0x91, 0x95, 0xe6, 0xf6, 0x3f, 0x3b, 0x6a, 0x0b, 0xcd, 0xd9, 0x9e, 0xba,
	0x11, 0xed, 0xcc, 0xba, 0x73, 0x4b, 0x8f, 0x65, 0x37, 0xd4, 0x26, 0xbb, 0x92, 0x36, 0x50, 0xd5,
	0x45, 0x21, 0xec, 0x3d, 0x51, 0xa6, 0xd4, 0xb6, 0x2e, 0x5a, 0x63, 0x73, 0x7a, 0x2a, 0x7b, 0x42,
	0xed, 0xea, 0xaa, 0x92, 0xa9, 0x2d, 0x64, 0xa0, 0x2c, 0xe9, 0xaa, 0x4a, 0xda, 0x50, 0xd0, 0xd3,
	0xd9, 0xae, 0xba, 0xae, 0x1b, 0xa3, 0x3d, 0xed, 0xa7, 0x16, 0x6f, 0x75, 0xcb, 0xf4, 0x4c, 0x2a,
	0x7b, 0xcf, 0x81, 0x9e, 0xcd, 0x76, 0xd4, 0x96, 0x8e, 0xa1, 0xa6, 0xe7, 0x92, 0x20, 0x06, 0x1b,
	0x5b, 0x3a, 0x00, 0xce, 0x59, 0x82, 0x29, 0xe9, 0x79, 0x18, 0xe6, 0xb5, 0xb6, 0x15, 0x17, 0xf4,
	0x42, 0x76, 0x53, 0xa9, 0xdc, 0xb5, 0x9d, 0xb3, 0x6c, 0x83, 0xa7, 0x17, 0xd3, 0xa1, 0x8b, 0x36,
	0xc8, 0x40, 0x2f, 0x81, 0x14, 0x5c, 0xea, 0xd8, 0x04, 0x7a, 0x19, 0x77, 0x14, 0xec, 0x73, 0xa1,
	0x57, 0x00, 0xb9, 0xd5, 0xa6, 0xa1, 0x57, 0x61, 0xcc, 0x7d, 0xe7, 0x24, 0xd0, 0x61, 0xf6, 0xa4,
	0xda, 0x1b, 0x31, 0x1c, 0x3d, 0xbd, 0x86, 0x49, 0xc6, 0xc2, 0x99, 0xd1, 0x74, 0x84, 0xb6, 0x52,
	0xf7, 0xd6, 0xd1, 0xeb, 0x68, 0x2b, 0x4d, 0x13, 0x58, 0xe8, 0x16, 0x54, 0x23, 0xc6, 0x00, 0x6f,
	0x64, 0x8f, 0xab, 0x9d, 0xd2, 0xd8, 0x8a, 0xa5, 0x13, 0x7a, 0x13, 0xc2, 0x8a, 0x9d, 0x29, 0x7a,
	0x3a, 0x46, 0x7f, 0xed, 0x1a, 0xf6, 0xf4, 0x16, 0xca, 0xa6, 0xc4, 0xfa, 0xe8, 0x76, 0xc2, 0x6d,
	0x8a, 0xf0, 0x36, 0xbc, 0x46, 0x8c, 0x1b, 0xef, 0xc0, 0xcb, 0x58, 0x0e, 0xa7, 0x58, 0xc2, 0x3b,
	0x18, 0x07, 0x0c, 0x64, 0xb2, 0x26, 0x12, 0x84, 0xde, 0x4d, 0x6d, 0x36, 0xb0, 0x94, 0x3a, 0x67,
	0x7a, 0x0f, 0x67, 0x33, 0x1e, 0xb0, 0x30, 0x7a, 0x1f, 0xa4, 0x71, 0xb9, 0x6e, 0xb4, 0xa7, 0x13,
	0x08, 0xdb, 0xa9, 0xf8, 0xe9, 0x20, 0x5c, 0xd2, 0x5d, 0xf8, 0xb7, 0xdc, 0x4e, 0x59, 0x5c, 0x49,
	0x1f, 0x40, 0x39, 0x32, 0x4f, 0x1f, 0x22, 0x55, 0xcb, 0xa1, 0x76, 0x05, 0x9d, 0x22, 0x78, 0x6b,
	0xc3, 0x74, 0xa0, 0x8f, 0x52, 0xd9, 0x06, 0xcb, 0x42, 0x1f, 0xaf, 0x71, 0x39, 0xd0, 0x27, 0xc9,
	0xa9, 0x5b, 0x6d, 0xf2, 0xd3, 0x91, 0xad, 0x16, 0xf4, 0xd9, 0xc8, 0x56, 0x23, 0x7e, 0x9e, 0x22,
	0x74, 0xeb, 0x7b, 0xbe, 0x48, 0x26, 0x5d, 0xc7, 0x2c, 0xf4, 0xe5, 0x78, 0x04, 0x6c, 0x6c, 0x45,
	0x5f, 0x21, 0x0f, 0xc6, 0xc5, 0x13, 0xf9, 0x1a, 0x3a, 0x6b, 0xf2, 0xba, 0x68, 0x48, 0x27, 0xec,
	0x82, 0x29, 0x07, 0x9a, 0x02, 0x3b, 0x31, 0x95, 0xb1, 0x94, 0x23, 0xa7, 0x3b, 0x47, 0xb6, 0x02,
	0xcf, 0x28, 0x99, 0x32, 0x5c, 0xd6, 0x96, 0x25, 0x9e, 0xcb, 0x8a, 0xe0, 0xc7, 0xaa, 0xd0, 0xdd,
	0xb1, 0x78, 0x67, 0xa9, 0x46, 0x77, 0x57, 0x3b, 0xcb, 0x64, 0xd0, 0x23, 0xdc, 0x6a, 0x99, 0x79,
	0xfa, 0x06, 0x1a, 0x71, 0x3a, 0x8a, 0xa1, 0x19, 0x6c, 0xc5, 0x35, 0x4c, 0x0d, 0xd4, 0xe2, 0x62,
	0x60, 0x6a, 0x47, 0x41, 0x0c, 0x7c, 0x4a, 0x16, 0xb3, 0x26, 0x0c, 0x7b, 0x97, 0xe4, 0xde, 0x58,
	0xea, 0x92, 0xdc, 0xbb, 0x18, 0xe8, 0xdb, 0x24, 0x0f, 0xe9, 0xb1, 0x08, 0xb0, 0x77, 0x51, 0x72,
	0x26, 0x9f, 0x70, 0xd0, 0x21, 0x7a, 0x0a, 0x90, 0x07, 0xce, 0xeb, 0x9c, 0x22, 0xb2, 0x04, 0x71,
	0x71, 0xda, 0x30, 0x9d, 0xa1, 0x1e, 0xbb, 0x22, 0x38, 0x3a, 0xc7, 0x1f, 0xb1, 0xef, 0x99, 0x7a,
	0xd4, 0xfa, 0x1e, 0x53, 0x0f, 0x2b, 0x18, 0x5b, 0xfa, 0x0e, 0xbb, 0xec, 0xfb, 0xf5, 0xa0, 0xdf,
	0x8f, 0x27, 0x80, 0xf7, 0x56, 0xd0, 0x59, 0xfa, 0x01, 0xf9, 0xfa, 0x3e, 0x30, 0xfd, 0x38, 0x16,
	0x31, 0xc2, 0x4f, 0xb8, 0x10, 0x10, 0xe2, 0x9f, 0x7f, 0xdd, 0x4e, 0x1f, 0x87, 0xbb, 0xff, 0x06,
	0x00, 0x00, 0xff, 0xff, 0x53, 0x8d, 0x41, 0x44, 0xc9, 0x04, 0x00, 0x00,
}
