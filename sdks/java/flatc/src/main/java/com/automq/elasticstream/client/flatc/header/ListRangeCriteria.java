// automatically generated by the FlatBuffers compiler, do not modify

package com.automq.elasticstream.client.flatc.header;

import com.google.flatbuffers.BaseVector;
import com.google.flatbuffers.BooleanVector;
import com.google.flatbuffers.ByteVector;
import com.google.flatbuffers.Constants;
import com.google.flatbuffers.DoubleVector;
import com.google.flatbuffers.FlatBufferBuilder;
import com.google.flatbuffers.FloatVector;
import com.google.flatbuffers.IntVector;
import com.google.flatbuffers.LongVector;
import com.google.flatbuffers.ShortVector;
import com.google.flatbuffers.StringVector;
import com.google.flatbuffers.Struct;
import com.google.flatbuffers.Table;
import com.google.flatbuffers.UnionVector;
import java.nio.ByteBuffer;
import java.nio.ByteOrder;

@SuppressWarnings("unused")
public final class ListRangeCriteria extends Table {
  public static void ValidateVersion() { Constants.FLATBUFFERS_23_3_3(); }
  public static ListRangeCriteria getRootAsListRangeCriteria(ByteBuffer _bb) { return getRootAsListRangeCriteria(_bb, new ListRangeCriteria()); }
  public static ListRangeCriteria getRootAsListRangeCriteria(ByteBuffer _bb, ListRangeCriteria obj) { _bb.order(ByteOrder.LITTLE_ENDIAN); return (obj.__assign(_bb.getInt(_bb.position()) + _bb.position(), _bb)); }
  public void __init(int _i, ByteBuffer _bb) { __reset(_i, _bb); }
  public ListRangeCriteria __assign(int _i, ByteBuffer _bb) { __init(_i, _bb); return this; }

  public int nodeId() { int o = __offset(4); return o != 0 ? bb.getInt(o + bb_pos) : -1; }
  public long streamId() { int o = __offset(6); return o != 0 ? bb.getLong(o + bb_pos) : -1L; }

  public static int createListRangeCriteria(FlatBufferBuilder builder,
      int nodeId,
      long streamId) {
    builder.startTable(2);
    ListRangeCriteria.addStreamId(builder, streamId);
    ListRangeCriteria.addNodeId(builder, nodeId);
    return ListRangeCriteria.endListRangeCriteria(builder);
  }

  public static void startListRangeCriteria(FlatBufferBuilder builder) { builder.startTable(2); }
  public static void addNodeId(FlatBufferBuilder builder, int nodeId) { builder.addInt(0, nodeId, -1); }
  public static void addStreamId(FlatBufferBuilder builder, long streamId) { builder.addLong(1, streamId, -1L); }
  public static int endListRangeCriteria(FlatBufferBuilder builder) {
    int o = builder.endTable();
    return o;
  }

  public static final class Vector extends BaseVector {
    public Vector __assign(int _vector, int _element_size, ByteBuffer _bb) { __reset(_vector, _element_size, _bb); return this; }

    public ListRangeCriteria get(int j) { return get(new ListRangeCriteria(), j); }
    public ListRangeCriteria get(ListRangeCriteria obj, int j) {  return obj.__assign(__indirect(__element(j), bb), bb); }
  }
  public ListRangeCriteriaT unpack() {
    ListRangeCriteriaT _o = new ListRangeCriteriaT();
    unpackTo(_o);
    return _o;
  }
  public void unpackTo(ListRangeCriteriaT _o) {
    int _oNodeId = nodeId();
    _o.setNodeId(_oNodeId);
    long _oStreamId = streamId();
    _o.setStreamId(_oStreamId);
  }
  public static int pack(FlatBufferBuilder builder, ListRangeCriteriaT _o) {
    if (_o == null) return 0;
    return createListRangeCriteria(
      builder,
      _o.getNodeId(),
      _o.getStreamId());
  }
}

