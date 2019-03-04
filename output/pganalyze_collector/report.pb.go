// Code generated by protoc-gen-go. DO NOT EDIT.
// source: report.proto

package pganalyze_collector

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Report struct {
	ReportRunId string               `protobuf:"bytes,1,opt,name=report_run_id,json=reportRunId,proto3" json:"report_run_id,omitempty"`
	ReportType  string               `protobuf:"bytes,2,opt,name=report_type,json=reportType,proto3" json:"report_type,omitempty"`
	CollectedAt *timestamp.Timestamp `protobuf:"bytes,3,opt,name=collected_at,json=collectedAt,proto3" json:"collected_at,omitempty"`
	// Types that are valid to be assigned to Data:
	//	*Report_BloatReportData
	//	*Report_BuffercacheReportData
	//	*Report_VacuumReportData
	//	*Report_SequenceReportData
	Data                 isReport_Data `protobuf_oneof:"data"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Report) Reset()         { *m = Report{} }
func (m *Report) String() string { return proto.CompactTextString(m) }
func (*Report) ProtoMessage()    {}
func (*Report) Descriptor() ([]byte, []int) {
	return fileDescriptor_report_6f005bc968a2c30e, []int{0}
}
func (m *Report) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Report.Unmarshal(m, b)
}
func (m *Report) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Report.Marshal(b, m, deterministic)
}
func (dst *Report) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Report.Merge(dst, src)
}
func (m *Report) XXX_Size() int {
	return xxx_messageInfo_Report.Size(m)
}
func (m *Report) XXX_DiscardUnknown() {
	xxx_messageInfo_Report.DiscardUnknown(m)
}

var xxx_messageInfo_Report proto.InternalMessageInfo

type isReport_Data interface {
	isReport_Data()
}

type Report_BloatReportData struct {
	BloatReportData *BloatReportData `protobuf:"bytes,10,opt,name=bloat_report_data,json=bloatReportData,proto3,oneof"`
}
type Report_BuffercacheReportData struct {
	BuffercacheReportData *BuffercacheReportData `protobuf:"bytes,11,opt,name=buffercache_report_data,json=buffercacheReportData,proto3,oneof"`
}
type Report_VacuumReportData struct {
	VacuumReportData *VacuumReportData `protobuf:"bytes,12,opt,name=vacuum_report_data,json=vacuumReportData,proto3,oneof"`
}
type Report_SequenceReportData struct {
	SequenceReportData *SequenceReportData `protobuf:"bytes,13,opt,name=sequence_report_data,json=sequenceReportData,proto3,oneof"`
}

func (*Report_BloatReportData) isReport_Data()       {}
func (*Report_BuffercacheReportData) isReport_Data() {}
func (*Report_VacuumReportData) isReport_Data()      {}
func (*Report_SequenceReportData) isReport_Data()    {}

func (m *Report) GetData() isReport_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Report) GetReportRunId() string {
	if m != nil {
		return m.ReportRunId
	}
	return ""
}

func (m *Report) GetReportType() string {
	if m != nil {
		return m.ReportType
	}
	return ""
}

func (m *Report) GetCollectedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CollectedAt
	}
	return nil
}

func (m *Report) GetBloatReportData() *BloatReportData {
	if x, ok := m.GetData().(*Report_BloatReportData); ok {
		return x.BloatReportData
	}
	return nil
}

func (m *Report) GetBuffercacheReportData() *BuffercacheReportData {
	if x, ok := m.GetData().(*Report_BuffercacheReportData); ok {
		return x.BuffercacheReportData
	}
	return nil
}

func (m *Report) GetVacuumReportData() *VacuumReportData {
	if x, ok := m.GetData().(*Report_VacuumReportData); ok {
		return x.VacuumReportData
	}
	return nil
}

func (m *Report) GetSequenceReportData() *SequenceReportData {
	if x, ok := m.GetData().(*Report_SequenceReportData); ok {
		return x.SequenceReportData
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Report) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Report_OneofMarshaler, _Report_OneofUnmarshaler, _Report_OneofSizer, []interface{}{
		(*Report_BloatReportData)(nil),
		(*Report_BuffercacheReportData)(nil),
		(*Report_VacuumReportData)(nil),
		(*Report_SequenceReportData)(nil),
	}
}

func _Report_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Report)
	// data
	switch x := m.Data.(type) {
	case *Report_BloatReportData:
		b.EncodeVarint(10<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.BloatReportData); err != nil {
			return err
		}
	case *Report_BuffercacheReportData:
		b.EncodeVarint(11<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.BuffercacheReportData); err != nil {
			return err
		}
	case *Report_VacuumReportData:
		b.EncodeVarint(12<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.VacuumReportData); err != nil {
			return err
		}
	case *Report_SequenceReportData:
		b.EncodeVarint(13<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.SequenceReportData); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Report.Data has unexpected type %T", x)
	}
	return nil
}

func _Report_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Report)
	switch tag {
	case 10: // data.bloat_report_data
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(BloatReportData)
		err := b.DecodeMessage(msg)
		m.Data = &Report_BloatReportData{msg}
		return true, err
	case 11: // data.buffercache_report_data
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(BuffercacheReportData)
		err := b.DecodeMessage(msg)
		m.Data = &Report_BuffercacheReportData{msg}
		return true, err
	case 12: // data.vacuum_report_data
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(VacuumReportData)
		err := b.DecodeMessage(msg)
		m.Data = &Report_VacuumReportData{msg}
		return true, err
	case 13: // data.sequence_report_data
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(SequenceReportData)
		err := b.DecodeMessage(msg)
		m.Data = &Report_SequenceReportData{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Report_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Report)
	// data
	switch x := m.Data.(type) {
	case *Report_BloatReportData:
		s := proto.Size(x.BloatReportData)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Report_BuffercacheReportData:
		s := proto.Size(x.BuffercacheReportData)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Report_VacuumReportData:
		s := proto.Size(x.VacuumReportData)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Report_SequenceReportData:
		s := proto.Size(x.SequenceReportData)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*Report)(nil), "pganalyze.collector.Report")
}

func init() { proto.RegisterFile("report.proto", fileDescriptor_report_6f005bc968a2c30e) }

var fileDescriptor_report_6f005bc968a2c30e = []byte{
	// 325 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xc1, 0x4f, 0x32, 0x31,
	0x10, 0xc5, 0x3f, 0x3e, 0x09, 0x89, 0xb3, 0x10, 0xb5, 0x48, 0xdc, 0xec, 0x05, 0x42, 0x34, 0x12,
	0x0f, 0x25, 0xd1, 0xb3, 0x07, 0x89, 0x07, 0xbd, 0xae, 0xe8, 0xc5, 0xc3, 0xa6, 0xed, 0x0e, 0x48,
	0xb2, 0x6c, 0x6b, 0x69, 0x49, 0xf0, 0xe8, 0x5f, 0x6e, 0x68, 0x41, 0x29, 0xec, 0xb1, 0x6f, 0xde,
	0xfc, 0x26, 0xef, 0x15, 0x9a, 0x1a, 0x95, 0xd4, 0x86, 0x2a, 0x2d, 0x8d, 0x24, 0x6d, 0x35, 0x65,
	0x25, 0x2b, 0x56, 0x5f, 0x48, 0x85, 0x2c, 0x0a, 0x14, 0x46, 0xea, 0xa4, 0x3b, 0x95, 0x72, 0x5a,
	0xe0, 0xd0, 0x59, 0xb8, 0x9d, 0x0c, 0xcd, 0x6c, 0x8e, 0x0b, 0xc3, 0xe6, 0xca, 0x6f, 0x25, 0x84,
	0x17, 0x92, 0x99, 0x6c, 0x97, 0x94, 0xc4, 0xdc, 0x4e, 0x26, 0xa8, 0x05, 0x13, 0x1f, 0x18, 0x4e,
	0xda, 0x4b, 0x26, 0xac, 0x9d, 0x87, 0x62, 0x67, 0x81, 0x9f, 0x16, 0x4b, 0x11, 0x7a, 0xfb, 0xdf,
	0x75, 0x68, 0xa4, 0x4e, 0x20, 0x7d, 0x68, 0xf9, 0x51, 0xa6, 0x6d, 0x99, 0xcd, 0xf2, 0xb8, 0xd6,
	0xab, 0x0d, 0x8e, 0xd3, 0xc8, 0x8b, 0xa9, 0x2d, 0x9f, 0x73, 0xd2, 0x85, 0xcd, 0x33, 0x33, 0x2b,
	0x85, 0xf1, 0x7f, 0xe7, 0x00, 0x2f, 0x8d, 0x57, 0x0a, 0xc9, 0x3d, 0x34, 0x37, 0xb9, 0x30, 0xcf,
	0x98, 0x89, 0x8f, 0x7a, 0xb5, 0x41, 0x74, 0x9b, 0x50, 0x9f, 0x90, 0x6e, 0x13, 0xd2, 0xf1, 0x36,
	0x61, 0x1a, 0xfd, 0xfa, 0x1f, 0x0c, 0x49, 0xe1, 0x6c, 0x37, 0x6a, 0x96, 0x33, 0xc3, 0x62, 0x70,
	0x8c, 0x4b, 0x5a, 0x51, 0x1d, 0x1d, 0xad, 0xdd, 0x3e, 0xc0, 0x23, 0x33, 0xec, 0xe9, 0x5f, 0x7a,
	0xc2, 0x43, 0x89, 0xe4, 0x70, 0x71, 0x58, 0x95, 0x27, 0x47, 0x8e, 0x7c, 0x53, 0x4d, 0xfe, 0xdb,
	0x09, 0xf8, 0x1d, 0x5e, 0x35, 0x20, 0xaf, 0x40, 0x82, 0xda, 0xfd, 0x81, 0xa6, 0x3b, 0x70, 0x55,
	0x79, 0xe0, 0xcd, 0xd9, 0x03, 0xf6, 0xe9, 0x72, 0x4f, 0x23, 0xef, 0x70, 0xbe, 0xf7, 0x71, 0x1e,
	0xdc, 0x72, 0xe0, 0xeb, 0x4a, 0xf0, 0xcb, 0x66, 0x21, 0x40, 0x93, 0xc5, 0x81, 0x3a, 0x6a, 0x40,
	0x7d, 0x0d, 0xe3, 0x0d, 0xf7, 0x2d, 0x77, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xfb, 0x45, 0x0d,
	0xb4, 0xab, 0x02, 0x00, 0x00,
}
