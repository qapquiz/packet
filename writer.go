package packet

import (
	"math"
	"sync"
)

// DefaultBufferCap is a default size of the []byte in Writer
const defaultBufferCap int = 64

var byteSlicePool = sync.Pool{
	New: func() interface{} {
		return make([]byte, defaultBufferCap)
	},
}

// A Writer is a serializer
type Writer struct {
	byteSlice  []byte
	idx        int
	currentCap int
}

// NewWriter will create new Writer
func NewWriter() Writer {
	byteSlice := byteSlicePool.Get().([]byte)

	return Writer{
		byteSlice:  byteSlice,
		currentCap: cap(byteSlice),
	}
}

// NewWriterWithHeader will create new Writer with header(uint16) for contain Content-Length
func NewWriterWithHeader() Writer {
	byteSlice := byteSlicePool.Get().([]byte)

	w := Writer{
		byteSlice:  byteSlice,
		currentCap: cap(byteSlice),
	}

	w.WriteUInt16(0)

	return w
}

// GetByteSliceFromPool will return []byte from pool
func GetByteSliceFromPool() []byte {
	return byteSlicePool.Get().([]byte)
}

// PutByteSliceToPool will return byteSlice to pool
func (w Writer) PutByteSliceToPool() {
	byteSlicePool.Put(w.byteSlice)
}

// Bytes will return []byte
func (w Writer) Bytes() []byte {
	return w.byteSlice[:w.idx]
}

// BytesWithHeader will return []byte with header
func (w Writer) BytesWithHeader() []byte {
	contentLength := uint16(len(w.byteSlice))

	w.byteSlice[0] = byte(contentLength >> 0)
	w.byteSlice[1] = byte(contentLength >> 8)

	return w.byteSlice[:w.idx]
}

// WriteBytes will write []byte at index
func (w *Writer) WriteBytes(b []byte) {
	l := len(b)
	if w.idx+l > w.currentCap {
		w.growBufferCap(w.idx + l)
	}

	copy(w.byteSlice[w.idx:w.idx+l], b)

	w.idx += l
}

// WriteString will write string to []byte at index
func (w *Writer) WriteString(s string) {
	l := len(s)
	if w.idx+8+l > w.currentCap {
		w.growBufferCap(w.idx + 8 + l)
	}

	w.WriteUInt64(uint64(l))

	copy(w.byteSlice[w.idx:w.idx+l], s)

	w.idx += l
}

// WriteStringSuffixZero will write string and end with 0
func (w *Writer) WriteStringSuffixZero(s string) {
	l := len(s)
	if w.idx+1+l > w.currentCap {
		w.growBufferCap(w.idx + 1 + l)
	}

	copy(w.byteSlice[w.idx:w.idx+l], s)

	w.idx += l

	w.WriteUInt8(0)
}

// WriteBoolean will write bool to the []byte at index
func (w *Writer) WriteBoolean(b bool) {
	if b {
		w.WriteUInt8(1)
	} else {
		w.WriteUInt8(0)
	}
}

// WriteFloat32 will write float32 to the []byte at index
func (w *Writer) WriteFloat32(f float32) {
	if w.idx+4 > w.currentCap {
		w.growBufferCap(w.idx + 4)
	}

	n := math.Float32bits(f)

	w.WriteUInt32(n)

	w.idx += 4
}

// WriteFloat64 will write float64 to the []byte at index
func (w *Writer) WriteFloat64(f float64) {
	if w.idx+8 > w.currentCap {
		w.growBufferCap(w.idx + 8)
	}

	n := math.Float64bits(f)

	w.WriteUInt64(n)

	w.idx += 8
}

// WriteInt8 will write uint8 to the []byte at index
func (w *Writer) WriteInt8(n int8) {
	if w.idx+1 > w.currentCap {
		w.growBufferCap(w.idx + 1)
	}

	w.byteSlice[w.idx] = byte(n)

	w.idx++
}

// WriteInt16 will write uint16 to the []byte at index
func (w *Writer) WriteInt16(n int16) {
	if w.idx+2 > w.currentCap {
		w.growBufferCap(w.idx + 2)
	}

	for i := 0; i < 2; i++ {
		w.byteSlice[w.idx+i] = byte(n >> (i * 8))
	}

	w.idx += 2
}

// WriteInt32 will write uint32 to the []byte at index
func (w *Writer) WriteInt32(n int32) {
	if w.idx+4 > w.currentCap {
		w.growBufferCap(w.idx + 4)
	}

	for i := 0; i < 4; i++ {
		w.byteSlice[w.idx+i] = byte(n >> (i * 8))
	}

	w.idx += 4
}

// WriteInt64 will write uint64 to the []byte at index
func (w *Writer) WriteInt64(n int64) {
	if w.idx+8 > w.currentCap {
		w.growBufferCap(w.idx + 8)
	}

	for i := 0; i < 8; i++ {
		w.byteSlice[w.idx+i] = byte(n >> (i * 8))
	}

	w.idx += 8
}

// WriteUInt8 will write uint8 to the []byte at index
func (w *Writer) WriteUInt8(n uint8) {
	if w.idx+1 > w.currentCap {
		w.growBufferCap(w.idx + 1)
	}

	w.byteSlice[w.idx] = byte(n)

	w.idx++
}

// WriteUInt16 will write uint16 to the []byte at index
func (w *Writer) WriteUInt16(n uint16) {
	if w.idx+2 > w.currentCap {
		w.growBufferCap(w.idx + 2)
	}

	for i := 0; i < 2; i++ {
		w.byteSlice[w.idx+i] = byte(n >> (i * 8))
	}

	w.idx += 2
}

// WriteUInt32 will write uint32 to the []byte at index
func (w *Writer) WriteUInt32(n uint32) {
	if w.idx+4 > w.currentCap {
		w.growBufferCap(w.idx + 4)
	}

	for i := 0; i < 4; i++ {
		w.byteSlice[w.idx+i] = byte(n >> (i * 8))
	}

	w.idx += 4
}

// WriteUInt64 will write uint64 to the []byte at index
func (w *Writer) WriteUInt64(n uint64) {
	if w.idx+8 > w.currentCap {
		w.growBufferCap(w.idx + 8)
	}

	for i := 0; i < 8; i++ {
		w.byteSlice[w.idx+i] = byte(n >> (i * 8))
	}

	w.idx += 8
}

func (w *Writer) growBufferCap(lengthNeed int) {
	for cap(w.byteSlice) < lengthNeed {
		w.byteSlice = append(w.byteSlice, make([]byte, w.currentCap)...)
		w.currentCap = w.currentCap * 2
	}
}
