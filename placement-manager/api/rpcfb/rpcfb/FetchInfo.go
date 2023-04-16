// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package rpcfb

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type FetchInfoT struct {
	StreamId int64 `json:"stream_id"`
	RequestIndex int32 `json:"request_index"`
	FetchOffset int64 `json:"fetch_offset"`
	BatchMaxBytes int32 `json:"batch_max_bytes"`
}

func (t *FetchInfoT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	FetchInfoStart(builder)
	FetchInfoAddStreamId(builder, t.StreamId)
	FetchInfoAddRequestIndex(builder, t.RequestIndex)
	FetchInfoAddFetchOffset(builder, t.FetchOffset)
	FetchInfoAddBatchMaxBytes(builder, t.BatchMaxBytes)
	return FetchInfoEnd(builder)
}

func (rcv *FetchInfo) UnPackTo(t *FetchInfoT) {
	t.StreamId = rcv.StreamId()
	t.RequestIndex = rcv.RequestIndex()
	t.FetchOffset = rcv.FetchOffset()
	t.BatchMaxBytes = rcv.BatchMaxBytes()
}

func (rcv *FetchInfo) UnPack() *FetchInfoT {
	if rcv == nil { return nil }
	t := &FetchInfoT{}
	rcv.UnPackTo(t)
	return t
}

type FetchInfo struct {
	_tab flatbuffers.Table
}

func GetRootAsFetchInfo(buf []byte, offset flatbuffers.UOffsetT) *FetchInfo {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &FetchInfo{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsFetchInfo(buf []byte, offset flatbuffers.UOffsetT) *FetchInfo {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &FetchInfo{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *FetchInfo) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *FetchInfo) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *FetchInfo) StreamId() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return -1
}

func (rcv *FetchInfo) MutateStreamId(n int64) bool {
	return rcv._tab.MutateInt64Slot(4, n)
}

func (rcv *FetchInfo) RequestIndex() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *FetchInfo) MutateRequestIndex(n int32) bool {
	return rcv._tab.MutateInt32Slot(6, n)
}

func (rcv *FetchInfo) FetchOffset() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *FetchInfo) MutateFetchOffset(n int64) bool {
	return rcv._tab.MutateInt64Slot(8, n)
}

func (rcv *FetchInfo) BatchMaxBytes() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *FetchInfo) MutateBatchMaxBytes(n int32) bool {
	return rcv._tab.MutateInt32Slot(10, n)
}

func FetchInfoStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func FetchInfoAddStreamId(builder *flatbuffers.Builder, streamId int64) {
	builder.PrependInt64Slot(0, streamId, -1)
}
func FetchInfoAddRequestIndex(builder *flatbuffers.Builder, requestIndex int32) {
	builder.PrependInt32Slot(1, requestIndex, 0)
}
func FetchInfoAddFetchOffset(builder *flatbuffers.Builder, fetchOffset int64) {
	builder.PrependInt64Slot(2, fetchOffset, 0)
}
func FetchInfoAddBatchMaxBytes(builder *flatbuffers.Builder, batchMaxBytes int32) {
	builder.PrependInt32Slot(3, batchMaxBytes, 0)
}
func FetchInfoEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
