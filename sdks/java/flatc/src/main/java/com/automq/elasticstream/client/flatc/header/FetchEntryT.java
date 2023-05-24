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

public class FetchEntryT {
  private com.automq.elasticstream.client.flatc.header.RangeT range;
  private long fetchOffset;
  private long endOffset;
  private int batchMaxBytes;

  public com.automq.elasticstream.client.flatc.header.RangeT getRange() { return range; }

  public void setRange(com.automq.elasticstream.client.flatc.header.RangeT range) { this.range = range; }

  public long getFetchOffset() { return fetchOffset; }

  public void setFetchOffset(long fetchOffset) { this.fetchOffset = fetchOffset; }

  public long getEndOffset() { return endOffset; }

  public void setEndOffset(long endOffset) { this.endOffset = endOffset; }

  public int getBatchMaxBytes() { return batchMaxBytes; }

  public void setBatchMaxBytes(int batchMaxBytes) { this.batchMaxBytes = batchMaxBytes; }


  public FetchEntryT() {
    this.range = null;
    this.fetchOffset = 0L;
    this.endOffset = 0L;
    this.batchMaxBytes = 0;
  }
}

