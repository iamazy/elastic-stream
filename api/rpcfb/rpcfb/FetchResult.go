// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package rpcfb

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type FetchResultT struct {
	StreamId int64 `json:"stream_id"`
	RequestIndex int32 `json:"request_index"`
	BatchLength int32 `json:"batch_length"`
	Status *StatusT `json:"status"`
}

func (t *FetchResultT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	statusOffset := t.Status.Pack(builder)
	FetchResultStart(builder)
	FetchResultAddStreamId(builder, t.StreamId)
	FetchResultAddRequestIndex(builder, t.RequestIndex)
	FetchResultAddBatchLength(builder, t.BatchLength)
	FetchResultAddStatus(builder, statusOffset)
	return FetchResultEnd(builder)
}

func (rcv *FetchResult) UnPackTo(t *FetchResultT) {
	t.StreamId = rcv.StreamId()
	t.RequestIndex = rcv.RequestIndex()
	t.BatchLength = rcv.BatchLength()
	t.Status = rcv.Status(nil).UnPack()
}

func (rcv *FetchResult) UnPack() *FetchResultT {
	if rcv == nil { return nil }
	t := &FetchResultT{}
	rcv.UnPackTo(t)
	return t
}

type FetchResult struct {
	_tab flatbuffers.Table
}

func GetRootAsFetchResult(buf []byte, offset flatbuffers.UOffsetT) *FetchResult {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &FetchResult{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsFetchResult(buf []byte, offset flatbuffers.UOffsetT) *FetchResult {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &FetchResult{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *FetchResult) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *FetchResult) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *FetchResult) StreamId() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return -1
}

func (rcv *FetchResult) MutateStreamId(n int64) bool {
	return rcv._tab.MutateInt64Slot(4, n)
}

func (rcv *FetchResult) RequestIndex() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *FetchResult) MutateRequestIndex(n int32) bool {
	return rcv._tab.MutateInt32Slot(6, n)
}

func (rcv *FetchResult) BatchLength() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *FetchResult) MutateBatchLength(n int32) bool {
	return rcv._tab.MutateInt32Slot(8, n)
}

func (rcv *FetchResult) Status(obj *Status) *Status {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(Status)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func FetchResultStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func FetchResultAddStreamId(builder *flatbuffers.Builder, streamId int64) {
	builder.PrependInt64Slot(0, streamId, -1)
}
func FetchResultAddRequestIndex(builder *flatbuffers.Builder, requestIndex int32) {
	builder.PrependInt32Slot(1, requestIndex, 0)
}
func FetchResultAddBatchLength(builder *flatbuffers.Builder, batchLength int32) {
	builder.PrependInt32Slot(2, batchLength, 0)
}
func FetchResultAddStatus(builder *flatbuffers.Builder, status flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(status), 0)
}
func FetchResultEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
