package graphson

import (
	"time"
	"unsafe"

	"github.com/json-iterator/go"
)

func init() {
	RegisterTypeEncoder("time.Time", typeEncoder{timeCodec{}, Timestamp})
	RegisterTypeDecoder("time.Time", typeDecoder{timeCodec{}, Types{Timestamp, Date}})
}

type timeCodec struct{}

func (timeCodec) IsEmpty(ptr unsafe.Pointer) bool {
	ts := *((*time.Time)(ptr))
	return ts.IsZero()
}

func (timeCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	ts := *((*time.Time)(ptr))
	stream.WriteInt64(ts.UnixNano() / time.Millisecond.Nanoseconds())
}

func (timeCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	ns := iter.ReadInt64() * time.Millisecond.Nanoseconds()
	*((*time.Time)(ptr)) = time.Unix(0, ns)
}
