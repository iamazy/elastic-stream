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

public class SealRangeRequestT {
  private int timeoutMs;
  private byte kind;
  private com.automq.elasticstream.client.flatc.header.RangeT range;

  public int getTimeoutMs() { return timeoutMs; }

  public void setTimeoutMs(int timeoutMs) { this.timeoutMs = timeoutMs; }

  public byte getKind() { return kind; }

  public void setKind(byte kind) { this.kind = kind; }

  public com.automq.elasticstream.client.flatc.header.RangeT getRange() { return range; }

  public void setRange(com.automq.elasticstream.client.flatc.header.RangeT range) { this.range = range; }


  public SealRangeRequestT() {
    this.timeoutMs = 0;
    this.kind = 0;
    this.range = null;
  }
}

