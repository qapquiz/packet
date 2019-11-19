package packet

import (
	"math"
)

// DefaultBufferCap is a default size of the []byte in Writer
const defaultBufferCap int = 64

// A Writer is a serializer
type Writer struct {
	bufBytes   []byte
	idx        int
	currentCap int
}

// NewWriter will create new Writer
func NewWriter() Writer {
	return Writer{
		bufBytes:   make([]byte, defaultBufferCap),
		currentCap: defaultBufferCap,
	}
}

// Bytes will return []byte
func (w Writer) Bytes() []byte {
	return w.bufBytes[:w.currentCap]
}

// WriteString will write string to []byte at index
func (w *Writer) WriteString(s string) {
	l := len(s)
	if w.idx+8+l > w.currentCap {
		w.growBufferCap(w.idx + 8 + l)
	}

	w.WriteUInt64(uint64(l))

	copy(w.bufBytes[w.idx:w.idx+l], s)

	w.idx += l
}

// WriteBoolean will write bool to the []byte at index
func (w *Writer) WriteBoolean(b bool) {
	if w.idx+1 > w.currentCap {
		w.growBufferCap(w.idx + 1)
	}

	if b {
		w.WriteUInt8(1)
	} else {
		w.WriteUInt8(0)
	}

	w.idx++
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

	w.bufBytes[w.idx] = byte(n)

	w.idx++
}

// WriteInt16 will write uint16 to the []byte at index
func (w *Writer) WriteInt16(n int16) {
	if w.idx+2 > w.currentCap {
		w.growBufferCap(w.idx + 2)
	}

	for i := 0; i < 2; i++ {
		w.bufBytes[w.idx+i] = byte(n >> (i*8))
	}

	w.idx += 2
}

// WriteInt32 will write uint32 to the []byte at index
func (w *Writer) WriteInt32(n int32) {
	if w.idx+4 > w.currentCap {
		w.growBufferCap(w.idx + 4)
	}

	for i := 0; i < 4; i++ {
		w.bufBytes[w.idx+i] = byte(n >> (i*8))
	}

	w.idx += 4
}

// WriteInt64 will write uint64 to the []byte at index
func (w *Writer) WriteInt64(n int64) {
	if w.idx+8 > w.currentCap {
		w.growBufferCap(w.idx + 8)
	}

	for i := 0; i < 8; i++ {
		w.bufBytes[w.idx+i] = byte(n >> (i*8))
	}

	w.idx += 8
}

// WriteUInt8 will write uint8 to the []byte at index
func (w *Writer) WriteUInt8(n uint8) {
	if w.idx+1 > w.currentCap {
		w.growBufferCap(w.idx + 1)
	}

	w.bufBytes[w.idx] = byte(n)

	w.idx++
}

// WriteUInt16 will write uint16 to the []byte at index
func (w *Writer) WriteUInt16(n uint16) {
	if w.idx+2 > w.currentCap {
		w.growBufferCap(w.idx + 2)
	}

	for i := 0; i < 2; i++ {
		w.bufBytes[w.idx+i] = byte(n >> (i*8))
	}

	w.idx += 2
}

// WriteUInt32 will write uint32 to the []byte at index
func (w *Writer) WriteUInt32(n uint32) {
	if w.idx+4 > w.currentCap {
		w.growBufferCap(w.idx + 4)
	}

	for i := 0; i < 4; i++ {
		w.bufBytes[w.idx+i] = byte(n >> (i*8))
	}

	w.idx += 4
}

// WriteUInt64 will write uint64 to the []byte at index
func (w *Writer) WriteUInt64(n uint64) {
	if w.idx+8 > w.currentCap {
		w.growBufferCap(w.idx + 8)
	}

	for i := 0; i < 8; i++ {
		w.bufBytes[w.idx+i] = byte(n >> (i * 8))
	}

	w.idx += 8
}

func (w *Writer) growBufferCap(lengthNeed int) {
	for cap(w.bufBytes) < lengthNeed {
		w.bufBytes = append(w.bufBytes, make([]byte, w.currentCap)...)
		w.currentCap = w.currentCap * 2
	}
}
