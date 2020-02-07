// Code generated by protoc-gen-go. DO NOT EDIT.
// source: checkpointing.proto

package checkpointing

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	common "github.com/offchainlabs/arbitrum/packages/arb-util/common"
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

type BlockIdBufList struct {
	Bufs                 []*common.BlockIdBuf `protobuf:"bytes,1,rep,name=bufs,proto3" json:"bufs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *BlockIdBufList) Reset()         { *m = BlockIdBufList{} }
func (m *BlockIdBufList) String() string { return proto.CompactTextString(m) }
func (*BlockIdBufList) ProtoMessage()    {}
func (*BlockIdBufList) Descriptor() ([]byte, []int) {
	return fileDescriptor_3071aace9480105b, []int{0}
}

func (m *BlockIdBufList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockIdBufList.Unmarshal(m, b)
}
func (m *BlockIdBufList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockIdBufList.Marshal(b, m, deterministic)
}
func (m *BlockIdBufList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockIdBufList.Merge(m, src)
}
func (m *BlockIdBufList) XXX_Size() int {
	return xxx_messageInfo_BlockIdBufList.Size(m)
}
func (m *BlockIdBufList) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockIdBufList.DiscardUnknown(m)
}

var xxx_messageInfo_BlockIdBufList proto.InternalMessageInfo

func (m *BlockIdBufList) GetBufs() []*common.BlockIdBuf {
	if m != nil {
		return m.Bufs
	}
	return nil
}

type CheckpointMetadata struct {
	FormatVersion        uint64             `protobuf:"varint,1,opt,name=formatVersion,proto3" json:"formatVersion,omitempty"`
	Oldest               *common.BlockIdBuf `protobuf:"bytes,2,opt,name=oldest,proto3" json:"oldest,omitempty"`
	Newest               *common.BlockIdBuf `protobuf:"bytes,3,opt,name=newest,proto3" json:"newest,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *CheckpointMetadata) Reset()         { *m = CheckpointMetadata{} }
func (m *CheckpointMetadata) String() string { return proto.CompactTextString(m) }
func (*CheckpointMetadata) ProtoMessage()    {}
func (*CheckpointMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_3071aace9480105b, []int{1}
}

func (m *CheckpointMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckpointMetadata.Unmarshal(m, b)
}
func (m *CheckpointMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckpointMetadata.Marshal(b, m, deterministic)
}
func (m *CheckpointMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckpointMetadata.Merge(m, src)
}
func (m *CheckpointMetadata) XXX_Size() int {
	return xxx_messageInfo_CheckpointMetadata.Size(m)
}
func (m *CheckpointMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckpointMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_CheckpointMetadata proto.InternalMessageInfo

func (m *CheckpointMetadata) GetFormatVersion() uint64 {
	if m != nil {
		return m.FormatVersion
	}
	return 0
}

func (m *CheckpointMetadata) GetOldest() *common.BlockIdBuf {
	if m != nil {
		return m.Oldest
	}
	return nil
}

func (m *CheckpointMetadata) GetNewest() *common.BlockIdBuf {
	if m != nil {
		return m.Newest
	}
	return nil
}

type CheckpointLinks struct {
	Next                 *common.BlockIdBuf `protobuf:"bytes,1,opt,name=next,proto3" json:"next,omitempty"`
	Prev                 *common.BlockIdBuf `protobuf:"bytes,2,opt,name=prev,proto3" json:"prev,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *CheckpointLinks) Reset()         { *m = CheckpointLinks{} }
func (m *CheckpointLinks) String() string { return proto.CompactTextString(m) }
func (*CheckpointLinks) ProtoMessage()    {}
func (*CheckpointLinks) Descriptor() ([]byte, []int) {
	return fileDescriptor_3071aace9480105b, []int{2}
}

func (m *CheckpointLinks) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckpointLinks.Unmarshal(m, b)
}
func (m *CheckpointLinks) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckpointLinks.Marshal(b, m, deterministic)
}
func (m *CheckpointLinks) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckpointLinks.Merge(m, src)
}
func (m *CheckpointLinks) XXX_Size() int {
	return xxx_messageInfo_CheckpointLinks.Size(m)
}
func (m *CheckpointLinks) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckpointLinks.DiscardUnknown(m)
}

var xxx_messageInfo_CheckpointLinks proto.InternalMessageInfo

func (m *CheckpointLinks) GetNext() *common.BlockIdBuf {
	if m != nil {
		return m.Next
	}
	return nil
}

func (m *CheckpointLinks) GetPrev() *common.BlockIdBuf {
	if m != nil {
		return m.Prev
	}
	return nil
}

type CheckpointManifest struct {
	Values               []*common.HashBuf `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
	Machines             []*common.HashBuf `protobuf:"bytes,2,rep,name=machines,proto3" json:"machines,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *CheckpointManifest) Reset()         { *m = CheckpointManifest{} }
func (m *CheckpointManifest) String() string { return proto.CompactTextString(m) }
func (*CheckpointManifest) ProtoMessage()    {}
func (*CheckpointManifest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3071aace9480105b, []int{3}
}

func (m *CheckpointManifest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckpointManifest.Unmarshal(m, b)
}
func (m *CheckpointManifest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckpointManifest.Marshal(b, m, deterministic)
}
func (m *CheckpointManifest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckpointManifest.Merge(m, src)
}
func (m *CheckpointManifest) XXX_Size() int {
	return xxx_messageInfo_CheckpointManifest.Size(m)
}
func (m *CheckpointManifest) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckpointManifest.DiscardUnknown(m)
}

var xxx_messageInfo_CheckpointManifest proto.InternalMessageInfo

func (m *CheckpointManifest) GetValues() []*common.HashBuf {
	if m != nil {
		return m.Values
	}
	return nil
}

func (m *CheckpointManifest) GetMachines() []*common.HashBuf {
	if m != nil {
		return m.Machines
	}
	return nil
}

func init() {
	proto.RegisterType((*BlockIdBufList)(nil), "structures.BlockIdBufList")
	proto.RegisterType((*CheckpointMetadata)(nil), "structures.CheckpointMetadata")
	proto.RegisterType((*CheckpointLinks)(nil), "structures.CheckpointLinks")
	proto.RegisterType((*CheckpointManifest)(nil), "structures.CheckpointManifest")
}

func init() { proto.RegisterFile("checkpointing.proto", fileDescriptor_3071aace9480105b) }

var fileDescriptor_3071aace9480105b = []byte{
	// 311 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0xd2, 0x4f, 0x4b, 0xfb, 0x30,
	0x18, 0x07, 0x70, 0xba, 0x8d, 0xf1, 0x23, 0x3f, 0x74, 0x90, 0x5d, 0x8a, 0xa7, 0x51, 0x44, 0x87,
	0xe2, 0x0a, 0x7a, 0xf1, 0x3c, 0x11, 0x15, 0xe6, 0x65, 0x07, 0x0f, 0xde, 0x9e, 0xa6, 0xc9, 0x1a,
	0xdb, 0x26, 0x25, 0x79, 0x32, 0x7d, 0x1d, 0xbe, 0x62, 0x49, 0xba, 0x3f, 0x14, 0xac, 0xa7, 0x42,
	0xbe, 0x9f, 0xf2, 0x7c, 0x9f, 0x36, 0x64, 0xca, 0x0a, 0xce, 0xca, 0x46, 0x4b, 0x85, 0x52, 0x6d,
	0x16, 0x8d, 0xd1, 0xa8, 0x29, 0xb1, 0x68, 0x1c, 0x43, 0x67, 0xb8, 0x3d, 0x9b, 0x32, 0x5d, 0xd7,
	0x5a, 0xa5, 0xed, 0xa3, 0x05, 0xc9, 0x3d, 0x39, 0x5d, 0x56, 0x9a, 0x95, 0x2f, 0xf9, 0xd2, 0x89,
	0x95, 0xb4, 0x48, 0x2f, 0xc8, 0x28, 0x73, 0xc2, 0xc6, 0xd1, 0x6c, 0x38, 0xff, 0x7f, 0x4b, 0x17,
	0x3b, 0x7e, 0x54, 0xeb, 0x90, 0x27, 0xdf, 0x11, 0xa1, 0x0f, 0x87, 0x91, 0xaf, 0x1c, 0x21, 0x07,
	0x04, 0x7a, 0x4e, 0x4e, 0x84, 0x36, 0x35, 0xe0, 0x1b, 0x37, 0x56, 0x6a, 0x15, 0x47, 0xb3, 0x68,
	0x3e, 0x5a, 0x77, 0x0f, 0xe9, 0x15, 0x19, 0xeb, 0x2a, 0xe7, 0x16, 0xe3, 0xc1, 0x2c, 0xea, 0x19,
	0xb3, 0x13, 0xde, 0x2a, 0xfe, 0xe9, 0xed, 0xb0, 0xdf, 0xb6, 0x22, 0x01, 0x32, 0x39, 0x76, 0x5a,
	0x49, 0x55, 0x5a, 0xbf, 0x8f, 0xe2, 0x5f, 0x18, 0x7a, 0xf4, 0xec, 0xe3, 0x73, 0xef, 0x1a, 0xc3,
	0xb7, 0x7f, 0x14, 0x0a, 0x79, 0xf2, 0xd1, 0x59, 0x1b, 0x94, 0x14, 0xbe, 0xe4, 0x25, 0x19, 0x6f,
	0xa1, 0x72, 0x7c, 0xff, 0xdd, 0x26, 0xfb, 0xf7, 0x9f, 0xc1, 0x16, 0xa1, 0x61, 0x1b, 0xd3, 0x6b,
	0xf2, 0xaf, 0x06, 0x56, 0x48, 0xc5, 0x6d, 0x3c, 0xf8, 0x9d, 0x1e, 0xc0, 0xf2, 0xe9, 0xfd, 0x71,
	0x23, 0xb1, 0x70, 0x99, 0x27, 0xa9, 0x16, 0x82, 0x15, 0x20, 0x55, 0x05, 0x99, 0x4d, 0xc1, 0x64,
	0x12, 0x8d, 0xab, 0xd3, 0x06, 0x58, 0x09, 0x1b, 0x1e, 0x4e, 0x6e, 0xb6, 0x50, 0xc9, 0x1c, 0x50,
	0x9b, 0xb4, 0x73, 0x1b, 0xb2, 0x71, 0xf8, 0xdb, 0x77, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xba,
	0x2e, 0x92, 0x4c, 0x25, 0x02, 0x00, 0x00,
}