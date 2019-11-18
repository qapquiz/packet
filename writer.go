package packet

import (
	"math"
)

// DefaultBufferCap is a default size of the []byte in Writer
const defaultBufferCap int = 16

// A Writer is a serializer
type Writer struct {
	bufBytes   []byte
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
	return w.bufBytes
}

// WriteString will write string to []byte at index
func (w *Writer) WriteString(s string, idx int) int {
	l := len(s)
	if idx+8+l > w.currentCap {
		w.growBufferCap(idx + 8 + l)
	}

	n := w.WriteUInt64(uint64(l), idx)

	copy(w.bufBytes[idx+n:idx+n+l], s)

	return n + l
}

// WriteBoolean will write bool to the []byte at index
func (w *Writer) WriteBoolean(b bool, idx int) int {
	if idx+1 > w.currentCap {
		w.growBufferCap(idx + 1)
	}

	if b {
		w.WriteUInt8(1, idx)
	} else {
		w.WriteUInt8(0, idx)
	}

	return 1
}

// WriteFloat32 will write float32 to the []byte at index
func (w *Writer) WriteFloat32(f float32, idx int) int {
	if idx+4 > w.currentCap {
		w.growBufferCap(idx + 4)
	}

	n := math.Float32bits(f)

	w.WriteUInt32(n, idx)

	return 4
}

// WriteFloat64 will write float64 to the []byte at index
func (w *Writer) WriteFloat64(f float64, idx int) int {
	if idx+8 > w.currentCap {
		w.growBufferCap(idx + 8)
	}

	n := math.Float64bits(f)

	w.WriteUInt64(n, idx)

	return 8
}

// WriteInt8 will write uint8 to the []byte at index
func (w *Writer) WriteInt8(n int8, idx int) int {
	if idx+1 > w.currentCap {
		w.growBufferCap(idx + 1)
	}

	w.bufBytes[idx] = byte(n)

	return 1
}

// WriteInt16 will write uint16 to the []byte at index
func (w *Writer) WriteInt16(n int16, idx int) int {
	if idx+2 > w.currentCap {
		w.growBufferCap(idx + 2)
	}

	for i := 0; i < 2; i++ {
		w.bufBytes[idx+i] = byte(n >> (i * 8))
	}

	return 2
}

// WriteInt32 will write uint32 to the []byte at index
func (w *Writer) WriteInt32(n int32, idx int) int {
	if idx+4 > w.currentCap {
		w.growBufferCap(idx + 4)
	}

	for i := 0; i < 4; i++ {
		w.bufBytes[idx+i] = byte(n >> (i * 8))
	}

	return 4
}

// WriteInt64 will write uint64 to the []byte at index
func (w *Writer) WriteInt64(n int64, idx int) int {
	if idx+8 > w.currentCap {
		w.growBufferCap(idx + 8)
	}

	for i := 0; i < 8; i++ {
		w.bufBytes[idx+i] = byte(n >> (i * 8))
	}

	return 8
}

// WriteUInt8 will write uint8 to the []byte at index
func (w *Writer) WriteUInt8(n uint8, idx int) int {
	if idx+1 > w.currentCap {
		w.growBufferCap(idx + 1)
	}

	w.bufBytes[idx] = byte(n)

	return 1
}

// WriteUInt16 will write uint16 to the []byte at index
func (w *Writer) WriteUInt16(n uint16, idx int) int {
	if idx+2 > w.currentCap {
		w.growBufferCap(idx + 2)
	}

	for i := 0; i < 2; i++ {
		w.bufBytes[idx+i] = byte(n >> (i * 8))
	}

	return 2
}

// WriteUInt32 will write uint32 to the []byte at index
func (w *Writer) WriteUInt32(n uint32, idx int) int {
	if idx+4 > w.currentCap {
		w.growBufferCap(idx + 4)
	}

	for i := 0; i < 4; i++ {
		w.bufBytes[idx+i] = byte(n >> (i * 8))
	}

	return 4
}

// WriteUInt64 will write uint64 to the []byte at index
func (w *Writer) WriteUInt64(n uint64, idx int) int {
	if idx+8 > w.currentCap {
		w.growBufferCap(idx + 8)
	}

	for i := 0; i < 8; i++ {
		w.bufBytes[idx+i] = byte(n >> (i * 8))
	}

	return 8
}

func (w *Writer) growBufferCap(lengthNeed int) {
	for cap(w.bufBytes) < lengthNeed {
		w.bufBytes = append(w.bufBytes, make([]byte, w.currentCap)...)
		w.currentCap = w.currentCap * 2
	}
}
