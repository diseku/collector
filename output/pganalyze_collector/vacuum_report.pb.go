// Code generated by protoc-gen-go. DO NOT EDIT.
// source: vacuum_report.proto

package pganalyze_collector

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type VacuumReportData struct {
	DatabaseReferences       []*DatabaseReference `protobuf:"bytes,10,rep,name=database_references,json=databaseReferences,proto3" json:"database_references,omitempty"`
	RelationReferences       []*RelationReference `protobuf:"bytes,11,rep,name=relation_references,json=relationReferences,proto3" json:"relation_references,omitempty"`
	VacuumStatistics         []*VacuumStatistic   `protobuf:"bytes,20,rep,name=vacuum_statistics,json=vacuumStatistics,proto3" json:"vacuum_statistics,omitempty"`
	AutovacuumMaxWorkers     int32                `protobuf:"varint,30,opt,name=autovacuum_max_workers,json=autovacuumMaxWorkers,proto3" json:"autovacuum_max_workers,omitempty"`
	AutovacuumNaptimeSeconds int32                `protobuf:"varint,31,opt,name=autovacuum_naptime_seconds,json=autovacuumNaptimeSeconds,proto3" json:"autovacuum_naptime_seconds,omitempty"`
	XXX_NoUnkeyedLiteral     struct{}             `json:"-"`
	XXX_unrecognized         []byte               `json:"-"`
	XXX_sizecache            int32                `json:"-"`
}

func (m *VacuumReportData) Reset()         { *m = VacuumReportData{} }
func (m *VacuumReportData) String() string { return proto.CompactTextString(m) }
func (*VacuumReportData) ProtoMessage()    {}
func (*VacuumReportData) Descriptor() ([]byte, []int) {
	return fileDescriptor_vacuum_report_476c7d75cf9da597, []int{0}
}
func (m *VacuumReportData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VacuumReportData.Unmarshal(m, b)
}
func (m *VacuumReportData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VacuumReportData.Marshal(b, m, deterministic)
}
func (dst *VacuumReportData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VacuumReportData.Merge(dst, src)
}
func (m *VacuumReportData) XXX_Size() int {
	return xxx_messageInfo_VacuumReportData.Size(m)
}
func (m *VacuumReportData) XXX_DiscardUnknown() {
	xxx_messageInfo_VacuumReportData.DiscardUnknown(m)
}

var xxx_messageInfo_VacuumReportData proto.InternalMessageInfo

func (m *VacuumReportData) GetDatabaseReferences() []*DatabaseReference {
	if m != nil {
		return m.DatabaseReferences
	}
	return nil
}

func (m *VacuumReportData) GetRelationReferences() []*RelationReference {
	if m != nil {
		return m.RelationReferences
	}
	return nil
}

func (m *VacuumReportData) GetVacuumStatistics() []*VacuumStatistic {
	if m != nil {
		return m.VacuumStatistics
	}
	return nil
}

func (m *VacuumReportData) GetAutovacuumMaxWorkers() int32 {
	if m != nil {
		return m.AutovacuumMaxWorkers
	}
	return 0
}

func (m *VacuumReportData) GetAutovacuumNaptimeSeconds() int32 {
	if m != nil {
		return m.AutovacuumNaptimeSeconds
	}
	return 0
}

type VacuumStatistic struct {
	RelationIdx                     int32          `protobuf:"varint,1,opt,name=relation_idx,json=relationIdx,proto3" json:"relation_idx,omitempty"`
	LiveRowCount                    int32          `protobuf:"varint,10,opt,name=live_row_count,json=liveRowCount,proto3" json:"live_row_count,omitempty"`
	DeadRowCount                    int32          `protobuf:"varint,11,opt,name=dead_row_count,json=deadRowCount,proto3" json:"dead_row_count,omitempty"`
	Relfrozenxid                    int32          `protobuf:"varint,12,opt,name=relfrozenxid,proto3" json:"relfrozenxid,omitempty"`
	Relminmxid                      int32          `protobuf:"varint,13,opt,name=relminmxid,proto3" json:"relminmxid,omitempty"`
	LastManualVacuumRun             *NullTimestamp `protobuf:"bytes,14,opt,name=last_manual_vacuum_run,json=lastManualVacuumRun,proto3" json:"last_manual_vacuum_run,omitempty"`
	LastAutoVacuumRun               *NullTimestamp `protobuf:"bytes,15,opt,name=last_auto_vacuum_run,json=lastAutoVacuumRun,proto3" json:"last_auto_vacuum_run,omitempty"`
	LastManualAnalyzeRun            *NullTimestamp `protobuf:"bytes,16,opt,name=last_manual_analyze_run,json=lastManualAnalyzeRun,proto3" json:"last_manual_analyze_run,omitempty"`
	LastAutoAnalyzeRun              *NullTimestamp `protobuf:"bytes,17,opt,name=last_auto_analyze_run,json=lastAutoAnalyzeRun,proto3" json:"last_auto_analyze_run,omitempty"`
	Fillfactor                      int32          `protobuf:"varint,18,opt,name=fillfactor,proto3" json:"fillfactor,omitempty"`
	AutovacuumEnabled               bool           `protobuf:"varint,20,opt,name=autovacuum_enabled,json=autovacuumEnabled,proto3" json:"autovacuum_enabled,omitempty"`
	AutovacuumVacuumThreshold       int32          `protobuf:"varint,21,opt,name=autovacuum_vacuum_threshold,json=autovacuumVacuumThreshold,proto3" json:"autovacuum_vacuum_threshold,omitempty"`
	AutovacuumAnalyzeThreshold      int32          `protobuf:"varint,22,opt,name=autovacuum_analyze_threshold,json=autovacuumAnalyzeThreshold,proto3" json:"autovacuum_analyze_threshold,omitempty"`
	AutovacuumVacuumScaleFactor     float64        `protobuf:"fixed64,23,opt,name=autovacuum_vacuum_scale_factor,json=autovacuumVacuumScaleFactor,proto3" json:"autovacuum_vacuum_scale_factor,omitempty"`
	AutovacuumAnalyzeScaleFactor    float64        `protobuf:"fixed64,24,opt,name=autovacuum_analyze_scale_factor,json=autovacuumAnalyzeScaleFactor,proto3" json:"autovacuum_analyze_scale_factor,omitempty"`
	AutovacuumFreezeMaxAge          int32          `protobuf:"varint,25,opt,name=autovacuum_freeze_max_age,json=autovacuumFreezeMaxAge,proto3" json:"autovacuum_freeze_max_age,omitempty"`
	AutovacuumMultixactFreezeMaxAge int32          `protobuf:"varint,26,opt,name=autovacuum_multixact_freeze_max_age,json=autovacuumMultixactFreezeMaxAge,proto3" json:"autovacuum_multixact_freeze_max_age,omitempty"`
	AutovacuumVacuumCostDelay       int32          `protobuf:"varint,27,opt,name=autovacuum_vacuum_cost_delay,json=autovacuumVacuumCostDelay,proto3" json:"autovacuum_vacuum_cost_delay,omitempty"`
	AutovacuumVacuumCostLimit       int32          `protobuf:"varint,28,opt,name=autovacuum_vacuum_cost_limit,json=autovacuumVacuumCostLimit,proto3" json:"autovacuum_vacuum_cost_limit,omitempty"`
	XXX_NoUnkeyedLiteral            struct{}       `json:"-"`
	XXX_unrecognized                []byte         `json:"-"`
	XXX_sizecache                   int32          `json:"-"`
}

func (m *VacuumStatistic) Reset()         { *m = VacuumStatistic{} }
func (m *VacuumStatistic) String() string { return proto.CompactTextString(m) }
func (*VacuumStatistic) ProtoMessage()    {}
func (*VacuumStatistic) Descriptor() ([]byte, []int) {
	return fileDescriptor_vacuum_report_476c7d75cf9da597, []int{1}
}
func (m *VacuumStatistic) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VacuumStatistic.Unmarshal(m, b)
}
func (m *VacuumStatistic) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VacuumStatistic.Marshal(b, m, deterministic)
}
func (dst *VacuumStatistic) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VacuumStatistic.Merge(dst, src)
}
func (m *VacuumStatistic) XXX_Size() int {
	return xxx_messageInfo_VacuumStatistic.Size(m)
}
func (m *VacuumStatistic) XXX_DiscardUnknown() {
	xxx_messageInfo_VacuumStatistic.DiscardUnknown(m)
}

var xxx_messageInfo_VacuumStatistic proto.InternalMessageInfo

func (m *VacuumStatistic) GetRelationIdx() int32 {
	if m != nil {
		return m.RelationIdx
	}
	return 0
}

func (m *VacuumStatistic) GetLiveRowCount() int32 {
	if m != nil {
		return m.LiveRowCount
	}
	return 0
}

func (m *VacuumStatistic) GetDeadRowCount() int32 {
	if m != nil {
		return m.DeadRowCount
	}
	return 0
}

func (m *VacuumStatistic) GetRelfrozenxid() int32 {
	if m != nil {
		return m.Relfrozenxid
	}
	return 0
}

func (m *VacuumStatistic) GetRelminmxid() int32 {
	if m != nil {
		return m.Relminmxid
	}
	return 0
}

func (m *VacuumStatistic) GetLastManualVacuumRun() *NullTimestamp {
	if m != nil {
		return m.LastManualVacuumRun
	}
	return nil
}

func (m *VacuumStatistic) GetLastAutoVacuumRun() *NullTimestamp {
	if m != nil {
		return m.LastAutoVacuumRun
	}
	return nil
}

func (m *VacuumStatistic) GetLastManualAnalyzeRun() *NullTimestamp {
	if m != nil {
		return m.LastManualAnalyzeRun
	}
	return nil
}

func (m *VacuumStatistic) GetLastAutoAnalyzeRun() *NullTimestamp {
	if m != nil {
		return m.LastAutoAnalyzeRun
	}
	return nil
}

func (m *VacuumStatistic) GetFillfactor() int32 {
	if m != nil {
		return m.Fillfactor
	}
	return 0
}

func (m *VacuumStatistic) GetAutovacuumEnabled() bool {
	if m != nil {
		return m.AutovacuumEnabled
	}
	return false
}

func (m *VacuumStatistic) GetAutovacuumVacuumThreshold() int32 {
	if m != nil {
		return m.AutovacuumVacuumThreshold
	}
	return 0
}

func (m *VacuumStatistic) GetAutovacuumAnalyzeThreshold() int32 {
	if m != nil {
		return m.AutovacuumAnalyzeThreshold
	}
	return 0
}

func (m *VacuumStatistic) GetAutovacuumVacuumScaleFactor() float64 {
	if m != nil {
		return m.AutovacuumVacuumScaleFactor
	}
	return 0
}

func (m *VacuumStatistic) GetAutovacuumAnalyzeScaleFactor() float64 {
	if m != nil {
		return m.AutovacuumAnalyzeScaleFactor
	}
	return 0
}

func (m *VacuumStatistic) GetAutovacuumFreezeMaxAge() int32 {
	if m != nil {
		return m.AutovacuumFreezeMaxAge
	}
	return 0
}

func (m *VacuumStatistic) GetAutovacuumMultixactFreezeMaxAge() int32 {
	if m != nil {
		return m.AutovacuumMultixactFreezeMaxAge
	}
	return 0
}

func (m *VacuumStatistic) GetAutovacuumVacuumCostDelay() int32 {
	if m != nil {
		return m.AutovacuumVacuumCostDelay
	}
	return 0
}

func (m *VacuumStatistic) GetAutovacuumVacuumCostLimit() int32 {
	if m != nil {
		return m.AutovacuumVacuumCostLimit
	}
	return 0
}

func init() {
	proto.RegisterType((*VacuumReportData)(nil), "pganalyze.collector.VacuumReportData")
	proto.RegisterType((*VacuumStatistic)(nil), "pganalyze.collector.VacuumStatistic")
}

func init() { proto.RegisterFile("vacuum_report.proto", fileDescriptor_vacuum_report_476c7d75cf9da597) }

var fileDescriptor_vacuum_report_476c7d75cf9da597 = []byte{
	// 674 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x95, 0xdd, 0x4e, 0xdb, 0x4a,
	0x10, 0x80, 0x15, 0xa1, 0x73, 0x74, 0xce, 0x26, 0x05, 0xb2, 0x09, 0xb0, 0x04, 0x14, 0xd2, 0x14,
	0x55, 0xb9, 0x69, 0x90, 0x68, 0x6f, 0x2a, 0x55, 0x6d, 0x11, 0x3f, 0x52, 0x25, 0x82, 0x54, 0x87,
	0x16, 0xf5, 0xca, 0x9a, 0xd8, 0x93, 0x60, 0x75, 0xed, 0x8d, 0x76, 0xd7, 0x10, 0x78, 0x85, 0xf6,
	0xa1, 0xab, 0x5d, 0xdb, 0x78, 0x13, 0x52, 0xc4, 0x95, 0x95, 0x99, 0x6f, 0x3e, 0xcf, 0x8c, 0xd7,
	0x0e, 0x69, 0xdc, 0x40, 0x90, 0xa6, 0xb1, 0x2f, 0x71, 0x2a, 0xa4, 0xee, 0x4f, 0xa5, 0xd0, 0x82,
	0x36, 0xa6, 0x13, 0x48, 0x80, 0xdf, 0xdd, 0x63, 0x3f, 0x10, 0x9c, 0x63, 0xa0, 0x85, 0x6c, 0xed,
	0x4d, 0x84, 0x98, 0x70, 0x3c, 0xb0, 0xc8, 0x28, 0x1d, 0x1f, 0xe8, 0x28, 0x46, 0xa5, 0x21, 0x9e,
	0x66, 0x55, 0xad, 0x9a, 0xba, 0x06, 0x89, 0x61, 0xf6, 0xab, 0xfb, 0x6b, 0x85, 0xac, 0x7f, 0xb7,
	0x6e, 0xcf, 0xaa, 0x4f, 0x40, 0x03, 0xbd, 0x22, 0x8d, 0x10, 0x34, 0x8c, 0x40, 0xa1, 0x2f, 0x71,
	0x8c, 0x12, 0x93, 0x00, 0x15, 0x23, 0x9d, 0x95, 0x5e, 0xf5, 0xf0, 0x75, 0x7f, 0xc9, 0x6d, 0xfb,
	0x27, 0x39, 0xef, 0x15, 0xb8, 0x47, 0xc3, 0xc5, 0x90, 0x32, 0x62, 0x89, 0x1c, 0x74, 0x24, 0x12,
	0x57, 0x5c, 0x7d, 0x42, 0xec, 0xe5, 0xbc, 0x23, 0x96, 0x8b, 0x21, 0x45, 0xbf, 0x92, 0x7a, 0xbe,
	0x21, 0xa5, 0x41, 0x47, 0x4a, 0x47, 0x81, 0x62, 0x4d, 0xab, 0xdd, 0x5f, 0xaa, 0xcd, 0x66, 0x1e,
	0x16, 0xb0, 0xb7, 0x7e, 0x33, 0x1f, 0x50, 0xf4, 0x1d, 0xd9, 0x84, 0x54, 0x8b, 0x5c, 0x1b, 0xc3,
	0xcc, 0xbf, 0x15, 0xf2, 0x27, 0x4a, 0xc5, 0xda, 0x9d, 0x4a, 0xef, 0x1f, 0xaf, 0x59, 0x66, 0x07,
	0x30, 0xbb, 0xca, 0x72, 0xf4, 0x03, 0x69, 0x39, 0x55, 0x09, 0x4c, 0xcd, 0xfa, 0x7d, 0x85, 0x81,
	0x48, 0x42, 0xc5, 0xf6, 0x6c, 0x25, 0x2b, 0x89, 0x8b, 0x0c, 0x18, 0x66, 0xf9, 0xee, 0xef, 0xff,
	0xc9, 0xda, 0x42, 0x67, 0xf4, 0x25, 0xa9, 0x3d, 0xec, 0x2c, 0x0a, 0x67, 0xac, 0x62, 0x1d, 0xd5,
	0x22, 0xf6, 0x25, 0x9c, 0xd1, 0x7d, 0xb2, 0xca, 0xa3, 0x1b, 0xf4, 0xa5, 0xb8, 0xf5, 0x03, 0x91,
	0x26, 0x9a, 0x11, 0x0b, 0xd5, 0x4c, 0xd4, 0x13, 0xb7, 0xc7, 0x26, 0x66, 0xa8, 0x10, 0x21, 0x74,
	0xa8, 0x6a, 0x46, 0x99, 0xe8, 0x03, 0xd5, 0xb5, 0xb7, 0x1b, 0x4b, 0x71, 0x8f, 0xc9, 0x2c, 0x0a,
	0x59, 0x2d, 0x63, 0xdc, 0x18, 0x6d, 0x13, 0x22, 0x91, 0xc7, 0x51, 0x12, 0x1b, 0xe2, 0x85, 0x25,
	0x9c, 0x08, 0xbd, 0x22, 0x9b, 0x1c, 0x94, 0xf6, 0x63, 0x48, 0x52, 0xe0, 0x7e, 0x71, 0x76, 0xd3,
	0x84, 0xad, 0x76, 0x2a, 0xbd, 0xea, 0x61, 0x77, 0xe9, 0x23, 0xb9, 0x48, 0x39, 0xbf, 0x2c, 0x0e,
	0xab, 0xd7, 0x30, 0x86, 0x81, 0x15, 0xe4, 0xe7, 0x33, 0x4d, 0xe8, 0x90, 0x34, 0xad, 0xd8, 0x2c,
	0xd0, 0xd5, 0xae, 0x3d, 0x5b, 0x5b, 0x37, 0xf5, 0x47, 0xa9, 0x16, 0xa5, 0xf4, 0x07, 0xd9, 0x72,
	0xbb, 0xcd, 0x0d, 0xd6, 0xbb, 0xfe, 0x6c, 0x6f, 0xb3, 0x6c, 0xf7, 0x28, 0x43, 0x8d, 0xfa, 0x1b,
	0xd9, 0x28, 0xfb, 0x75, 0xc5, 0xf5, 0x67, 0x8b, 0x69, 0xd1, 0xb0, 0xa3, 0x6d, 0x13, 0x32, 0x8e,
	0x38, 0x1f, 0x83, 0xe1, 0x19, 0xcd, 0xf6, 0x5f, 0x46, 0xe8, 0x1b, 0x42, 0x9d, 0x43, 0x88, 0x09,
	0x8c, 0x38, 0x86, 0xac, 0xd9, 0xa9, 0xf4, 0xfe, 0xf3, 0xea, 0x65, 0xe6, 0x34, 0x4b, 0xd0, 0x8f,
	0x64, 0xc7, 0xc1, 0xf3, 0x8b, 0xbe, 0x96, 0xa8, 0xae, 0x05, 0x0f, 0xd9, 0x86, 0xf5, 0x6f, 0x97,
	0x48, 0xb6, 0xba, 0xcb, 0x02, 0xa0, 0x9f, 0xc9, 0xae, 0x53, 0x5f, 0x8c, 0x59, 0x0a, 0x36, 0xad,
	0xc0, 0x79, 0x2f, 0xf2, 0x51, 0x4a, 0xc3, 0x31, 0x69, 0x3f, 0xee, 0x40, 0x05, 0xc0, 0xd1, 0xcf,
	0x87, 0xdc, 0xea, 0x54, 0x7a, 0x15, 0x6f, 0x67, 0xb1, 0x89, 0xa1, 0x61, 0xce, 0xb2, 0xa9, 0x4f,
	0xc9, 0xde, 0x92, 0x36, 0xe6, 0x2c, 0xcc, 0x5a, 0x76, 0x1f, 0x75, 0xe2, 0x6a, 0xde, 0x13, 0x67,
	0x54, 0x7f, 0x2c, 0x11, 0xef, 0xd1, 0xbe, 0xfe, 0x30, 0x41, 0xb6, 0x6d, 0x47, 0x71, 0x3e, 0x0c,
	0x67, 0x36, 0x3f, 0x80, 0xd9, 0xd1, 0x04, 0xe9, 0x39, 0x79, 0xe5, 0x7e, 0x32, 0x52, 0xae, 0xa3,
	0x19, 0x04, 0x7a, 0x51, 0xd2, 0xb2, 0x12, 0xa7, 0xd9, 0x41, 0x41, 0xce, 0xd9, 0x3e, 0xcd, 0xad,
	0x35, 0xbf, 0x04, 0x42, 0x69, 0x3f, 0x44, 0x0e, 0x77, 0x6c, 0x67, 0xf9, 0x73, 0x39, 0x16, 0x4a,
	0x9f, 0x18, 0xe0, 0x09, 0x01, 0x8f, 0xe2, 0x48, 0xb3, 0xdd, 0xbf, 0x0b, 0xce, 0x0d, 0x30, 0xfa,
	0xd7, 0xfe, 0x47, 0xbc, 0xfd, 0x13, 0x00, 0x00, 0xff, 0xff, 0x8d, 0xad, 0xcb, 0x86, 0x7e, 0x06,
	0x00, 0x00,
}
