// The MIT License
//
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
//
// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: temporal/server/api/common/v1/dlq.proto

package commonspb

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"

	proto "github.com/gogo/protobuf/proto"
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

type HistoryDLQTaskMetadata struct {
	// message_id is the zero-indexed sequence number of the message in the queue that contains this history task.
	MessageId int64 `protobuf:"varint,1,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
}

func (m *HistoryDLQTaskMetadata) Reset()      { *m = HistoryDLQTaskMetadata{} }
func (*HistoryDLQTaskMetadata) ProtoMessage() {}
func (*HistoryDLQTaskMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_d169aecba75076d1, []int{0}
}
func (m *HistoryDLQTaskMetadata) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HistoryDLQTaskMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HistoryDLQTaskMetadata.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HistoryDLQTaskMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HistoryDLQTaskMetadata.Merge(m, src)
}
func (m *HistoryDLQTaskMetadata) XXX_Size() int {
	return m.Size()
}
func (m *HistoryDLQTaskMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_HistoryDLQTaskMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_HistoryDLQTaskMetadata proto.InternalMessageInfo

func (m *HistoryDLQTaskMetadata) GetMessageId() int64 {
	if m != nil {
		return m.MessageId
	}
	return 0
}

// HistoryDLQTask is a history task that has been moved to the DLQ, so it also has a message ID (index within that
// queue).
type HistoryDLQTask struct {
	Metadata *HistoryDLQTaskMetadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// This is named payload to prevent stuttering (e.g. task.Task).
	Payload *ShardedTask `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (m *HistoryDLQTask) Reset()      { *m = HistoryDLQTask{} }
func (*HistoryDLQTask) ProtoMessage() {}
func (*HistoryDLQTask) Descriptor() ([]byte, []int) {
	return fileDescriptor_d169aecba75076d1, []int{1}
}
func (m *HistoryDLQTask) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HistoryDLQTask) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HistoryDLQTask.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HistoryDLQTask) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HistoryDLQTask.Merge(m, src)
}
func (m *HistoryDLQTask) XXX_Size() int {
	return m.Size()
}
func (m *HistoryDLQTask) XXX_DiscardUnknown() {
	xxx_messageInfo_HistoryDLQTask.DiscardUnknown(m)
}

var xxx_messageInfo_HistoryDLQTask proto.InternalMessageInfo

func (m *HistoryDLQTask) GetMetadata() *HistoryDLQTaskMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *HistoryDLQTask) GetPayload() *ShardedTask {
	if m != nil {
		return m.Payload
	}
	return nil
}

// HistoryDLQKey is a compound key that identifies a history DLQ.
type HistoryDLQKey struct {
	// task_category is the category of the task. The default values are defined in the TaskCategory enum. However, there
	// may also be other categories registered at runtime with the history/tasks package. As a result, the category here
	// is an integer instead of an enum to support both the default values and custom values.
	TaskCategory int32 `protobuf:"varint,1,opt,name=task_category,json=taskCategory,proto3" json:"task_category,omitempty"`
	// source_cluster and target_cluster must both be non-empty. For non-cross DC tasks, i.e. non-replication tasks,
	// they should be the same. The reason for this is that we may support wildcard clusters in the future, and we want
	// to differentiate between queues which go from one cluster to all other clusters, and queues which don't leave the
	// current cluster.
	SourceCluster string `protobuf:"bytes,2,opt,name=source_cluster,json=sourceCluster,proto3" json:"source_cluster,omitempty"`
	TargetCluster string `protobuf:"bytes,3,opt,name=target_cluster,json=targetCluster,proto3" json:"target_cluster,omitempty"`
}

func (m *HistoryDLQKey) Reset()      { *m = HistoryDLQKey{} }
func (*HistoryDLQKey) ProtoMessage() {}
func (*HistoryDLQKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_d169aecba75076d1, []int{2}
}
func (m *HistoryDLQKey) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HistoryDLQKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HistoryDLQKey.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HistoryDLQKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HistoryDLQKey.Merge(m, src)
}
func (m *HistoryDLQKey) XXX_Size() int {
	return m.Size()
}
func (m *HistoryDLQKey) XXX_DiscardUnknown() {
	xxx_messageInfo_HistoryDLQKey.DiscardUnknown(m)
}

var xxx_messageInfo_HistoryDLQKey proto.InternalMessageInfo

func (m *HistoryDLQKey) GetTaskCategory() int32 {
	if m != nil {
		return m.TaskCategory
	}
	return 0
}

func (m *HistoryDLQKey) GetSourceCluster() string {
	if m != nil {
		return m.SourceCluster
	}
	return ""
}

func (m *HistoryDLQKey) GetTargetCluster() string {
	if m != nil {
		return m.TargetCluster
	}
	return ""
}

func init() {
	proto.RegisterType((*HistoryDLQTaskMetadata)(nil), "temporal.server.api.common.v1.HistoryDLQTaskMetadata")
	proto.RegisterType((*HistoryDLQTask)(nil), "temporal.server.api.common.v1.HistoryDLQTask")
	proto.RegisterType((*HistoryDLQKey)(nil), "temporal.server.api.common.v1.HistoryDLQKey")
}

func init() {
	proto.RegisterFile("temporal/server/api/common/v1/dlq.proto", fileDescriptor_d169aecba75076d1)
}

var fileDescriptor_d169aecba75076d1 = []byte{
	// 359 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0x3f, 0x4b, 0xfb, 0x40,
	0x18, 0xc7, 0x73, 0xbf, 0xf2, 0x53, 0x7b, 0xda, 0x0e, 0x19, 0xa4, 0x08, 0x3d, 0xa4, 0x22, 0xfe,
	0x01, 0x13, 0xaa, 0x88, 0x83, 0x9b, 0xed, 0xa0, 0xa8, 0x43, 0xa3, 0x93, 0x4b, 0x79, 0x9a, 0x1c,
	0x31, 0x34, 0xf1, 0xe2, 0xdd, 0xb5, 0x90, 0x4d, 0x7c, 0x05, 0xbe, 0x0c, 0x7d, 0x27, 0x8e, 0x1d,
	0x3b, 0xda, 0xeb, 0xe2, 0xd8, 0x97, 0x20, 0xc9, 0xa5, 0x2d, 0x42, 0xa9, 0x5b, 0xf8, 0xf2, 0x79,
	0x3e, 0xcf, 0x37, 0xf7, 0xe0, 0x3d, 0x49, 0xa3, 0x98, 0x71, 0x08, 0x6d, 0x41, 0x79, 0x9f, 0x72,
	0x1b, 0xe2, 0xc0, 0x76, 0x59, 0x14, 0xb1, 0x27, 0xbb, 0x5f, 0xb7, 0xbd, 0xf0, 0xd9, 0x8a, 0x39,
	0x93, 0xcc, 0xac, 0x4e, 0x41, 0x4b, 0x83, 0x16, 0xc4, 0x81, 0xa5, 0x41, 0xab, 0x5f, 0xdf, 0x3a,
	0x58, 0xee, 0x91, 0x20, 0xba, 0x42, 0x9b, 0x6a, 0x67, 0x78, 0xf3, 0x32, 0x10, 0x92, 0xf1, 0xa4,
	0x79, 0xd3, 0xba, 0x07, 0xd1, 0xbd, 0xa5, 0x12, 0x3c, 0x90, 0x60, 0x56, 0x31, 0x8e, 0xa8, 0x10,
	0xe0, 0xd3, 0x76, 0xe0, 0x55, 0xd0, 0x36, 0xda, 0x2f, 0x38, 0xc5, 0x3c, 0xb9, 0xf2, 0x6a, 0x1f,
	0x08, 0x97, 0x7f, 0x4f, 0x9a, 0x2d, 0xbc, 0x16, 0xe5, 0xd3, 0x19, 0xbf, 0x7e, 0x7c, 0x6a, 0x2d,
	0x2d, 0x6a, 0x2d, 0x5e, 0xed, 0xcc, 0x34, 0x66, 0x13, 0xaf, 0xc6, 0x90, 0x84, 0x0c, 0xbc, 0xca,
	0xbf, 0xcc, 0x78, 0xf8, 0x87, 0xf1, 0xee, 0x11, 0xb8, 0x47, 0xbd, 0x54, 0xe7, 0x4c, 0x47, 0x6b,
	0xaf, 0x08, 0x97, 0xe6, 0xab, 0xae, 0x69, 0x62, 0xee, 0xe0, 0x52, 0xfa, 0x0a, 0x6d, 0x17, 0x24,
	0xf5, 0x19, 0x4f, 0xb2, 0xbe, 0xff, 0x9d, 0x8d, 0x34, 0x6c, 0xe4, 0x99, 0xb9, 0x8b, 0xcb, 0x82,
	0xf5, 0xb8, 0x4b, 0xdb, 0x6e, 0xd8, 0x13, 0x92, 0xf2, 0xac, 0x43, 0xd1, 0x29, 0xe9, 0xb4, 0xa1,
	0xc3, 0x14, 0x93, 0xc0, 0x7d, 0x2a, 0x67, 0x58, 0x41, 0x63, 0x3a, 0xcd, 0xb1, 0x0b, 0x77, 0x30,
	0x22, 0xc6, 0x70, 0x44, 0x8c, 0xc9, 0x88, 0xa0, 0x17, 0x45, 0xd0, 0xbb, 0x22, 0xe8, 0x53, 0x11,
	0x34, 0x50, 0x04, 0x7d, 0x29, 0x82, 0xbe, 0x15, 0x31, 0x26, 0x8a, 0xa0, 0xb7, 0x31, 0x31, 0x06,
	0x63, 0x62, 0x0c, 0xc7, 0xc4, 0x78, 0x38, 0xf2, 0xd9, 0xfc, 0x8f, 0x03, 0xb6, 0xf0, 0xa0, 0xe7,
	0xfa, 0x4b, 0xc4, 0x9d, 0xce, 0x4a, 0x76, 0xd5, 0x93, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xac,
	0x8f, 0x87, 0x3d, 0x4a, 0x02, 0x00, 0x00,
}

func (this *HistoryDLQTaskMetadata) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*HistoryDLQTaskMetadata)
	if !ok {
		that2, ok := that.(HistoryDLQTaskMetadata)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.MessageId != that1.MessageId {
		return false
	}
	return true
}
func (this *HistoryDLQTask) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*HistoryDLQTask)
	if !ok {
		that2, ok := that.(HistoryDLQTask)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Metadata.Equal(that1.Metadata) {
		return false
	}
	if !this.Payload.Equal(that1.Payload) {
		return false
	}
	return true
}
func (this *HistoryDLQKey) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*HistoryDLQKey)
	if !ok {
		that2, ok := that.(HistoryDLQKey)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.TaskCategory != that1.TaskCategory {
		return false
	}
	if this.SourceCluster != that1.SourceCluster {
		return false
	}
	if this.TargetCluster != that1.TargetCluster {
		return false
	}
	return true
}
func (this *HistoryDLQTaskMetadata) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&commonspb.HistoryDLQTaskMetadata{")
	s = append(s, "MessageId: "+fmt.Sprintf("%#v", this.MessageId)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *HistoryDLQTask) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&commonspb.HistoryDLQTask{")
	if this.Metadata != nil {
		s = append(s, "Metadata: "+fmt.Sprintf("%#v", this.Metadata)+",\n")
	}
	if this.Payload != nil {
		s = append(s, "Payload: "+fmt.Sprintf("%#v", this.Payload)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *HistoryDLQKey) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&commonspb.HistoryDLQKey{")
	s = append(s, "TaskCategory: "+fmt.Sprintf("%#v", this.TaskCategory)+",\n")
	s = append(s, "SourceCluster: "+fmt.Sprintf("%#v", this.SourceCluster)+",\n")
	s = append(s, "TargetCluster: "+fmt.Sprintf("%#v", this.TargetCluster)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringDlq(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *HistoryDLQTaskMetadata) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HistoryDLQTaskMetadata) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HistoryDLQTaskMetadata) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MessageId != 0 {
		i = encodeVarintDlq(dAtA, i, uint64(m.MessageId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *HistoryDLQTask) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HistoryDLQTask) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HistoryDLQTask) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Payload != nil {
		{
			size, err := m.Payload.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintDlq(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Metadata != nil {
		{
			size, err := m.Metadata.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintDlq(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *HistoryDLQKey) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HistoryDLQKey) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HistoryDLQKey) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TargetCluster) > 0 {
		i -= len(m.TargetCluster)
		copy(dAtA[i:], m.TargetCluster)
		i = encodeVarintDlq(dAtA, i, uint64(len(m.TargetCluster)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.SourceCluster) > 0 {
		i -= len(m.SourceCluster)
		copy(dAtA[i:], m.SourceCluster)
		i = encodeVarintDlq(dAtA, i, uint64(len(m.SourceCluster)))
		i--
		dAtA[i] = 0x12
	}
	if m.TaskCategory != 0 {
		i = encodeVarintDlq(dAtA, i, uint64(m.TaskCategory))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintDlq(dAtA []byte, offset int, v uint64) int {
	offset -= sovDlq(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *HistoryDLQTaskMetadata) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MessageId != 0 {
		n += 1 + sovDlq(uint64(m.MessageId))
	}
	return n
}

func (m *HistoryDLQTask) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Metadata != nil {
		l = m.Metadata.Size()
		n += 1 + l + sovDlq(uint64(l))
	}
	if m.Payload != nil {
		l = m.Payload.Size()
		n += 1 + l + sovDlq(uint64(l))
	}
	return n
}

func (m *HistoryDLQKey) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TaskCategory != 0 {
		n += 1 + sovDlq(uint64(m.TaskCategory))
	}
	l = len(m.SourceCluster)
	if l > 0 {
		n += 1 + l + sovDlq(uint64(l))
	}
	l = len(m.TargetCluster)
	if l > 0 {
		n += 1 + l + sovDlq(uint64(l))
	}
	return n
}

func sovDlq(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDlq(x uint64) (n int) {
	return sovDlq(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *HistoryDLQTaskMetadata) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&HistoryDLQTaskMetadata{`,
		`MessageId:` + fmt.Sprintf("%v", this.MessageId) + `,`,
		`}`,
	}, "")
	return s
}
func (this *HistoryDLQTask) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&HistoryDLQTask{`,
		`Metadata:` + strings.Replace(this.Metadata.String(), "HistoryDLQTaskMetadata", "HistoryDLQTaskMetadata", 1) + `,`,
		`Payload:` + strings.Replace(fmt.Sprintf("%v", this.Payload), "ShardedTask", "ShardedTask", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *HistoryDLQKey) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&HistoryDLQKey{`,
		`TaskCategory:` + fmt.Sprintf("%v", this.TaskCategory) + `,`,
		`SourceCluster:` + fmt.Sprintf("%v", this.SourceCluster) + `,`,
		`TargetCluster:` + fmt.Sprintf("%v", this.TargetCluster) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringDlq(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *HistoryDLQTaskMetadata) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDlq
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
			return fmt.Errorf("proto: HistoryDLQTaskMetadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HistoryDLQTaskMetadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MessageId", wireType)
			}
			m.MessageId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDlq
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MessageId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipDlq(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDlq
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthDlq
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
func (m *HistoryDLQTask) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDlq
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
			return fmt.Errorf("proto: HistoryDLQTask: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HistoryDLQTask: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDlq
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
				return ErrInvalidLengthDlq
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDlq
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Metadata == nil {
				m.Metadata = &HistoryDLQTaskMetadata{}
			}
			if err := m.Metadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payload", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDlq
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
				return ErrInvalidLengthDlq
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDlq
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Payload == nil {
				m.Payload = &ShardedTask{}
			}
			if err := m.Payload.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDlq(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDlq
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthDlq
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
func (m *HistoryDLQKey) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDlq
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
			return fmt.Errorf("proto: HistoryDLQKey: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HistoryDLQKey: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TaskCategory", wireType)
			}
			m.TaskCategory = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDlq
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TaskCategory |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceCluster", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDlq
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
				return ErrInvalidLengthDlq
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDlq
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SourceCluster = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TargetCluster", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDlq
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
				return ErrInvalidLengthDlq
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDlq
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TargetCluster = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDlq(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDlq
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthDlq
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
func skipDlq(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDlq
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
					return 0, ErrIntOverflowDlq
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
					return 0, ErrIntOverflowDlq
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
				return 0, ErrInvalidLengthDlq
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDlq
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDlq
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDlq        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDlq          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDlq = fmt.Errorf("proto: unexpected end of group")
)