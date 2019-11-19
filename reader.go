package packet

import (
	"math"
)

// A Reader is a deserializer
type Reader struct {
	bufBytes         []byte
	currentReadIndex int
}

// NewReader will create Reader
func NewReader(in []byte) Reader {
	return Reader{
		bufBytes: in,
	}
}

// IsComplete check is server read all data
func (r Reader) IsComplete() bool {
	return len(r.bufBytes) == r.currentReadIndex
}

// ReadString will read boolean at index
func (r *Reader) ReadString() string {
	l := r.ReadUInt64()

	idx := r.currentReadIndex
	r.currentReadIndex += int(l)

	return string(r.bufBytes[idx : idx+int(l)])
}

// ReadBoolean will read boolean at index
func (r *Reader) ReadBoolean() bool {
	idx := r.currentReadIndex
	r.currentReadIndex++
	return r.bufBytes[idx] == 1
}

// ReadFloat32 will read float32 at index
func (r *Reader) ReadFloat32() float32 {
	return math.Float32frombits(r.ReadUInt32())
}

// ReadFloat64 will read float64 at index
func (r *Reader) ReadFloat64() float64 {
	return math.Float64frombits(r.ReadUInt64())
}

// ReadInt8 will read int8 at index
func (r *Reader) ReadInt8() int8 {
	idx := r.currentReadIndex
	r.currentReadIndex++
	return int8(r.bufBytes[idx])
}

// ReadInt16 will read int16 at index
func (r *Reader) ReadInt16() int16 {
	idx := r.currentReadIndex
	r.currentReadIndex += 2
	return int16(r.bufBytes[idx]) | int16(r.bufBytes[idx+1])<<8
}

// ReadInt32 will read int32 at index
func (r *Reader) ReadInt32() int32 {
	idx := r.currentReadIndex
	r.currentReadIndex += 4
	return int32(r.bufBytes[idx]) | int32(r.bufBytes[idx+1])<<8 | int32(r.bufBytes[idx+2])<<16 | int32(r.bufBytes[idx+3])<<24
}

// ReadInt64 will read int64 at index
func (r *Reader) ReadInt64() int64 {
	idx := r.currentReadIndex
	r.currentReadIndex += 8
	return int64(r.bufBytes[idx]) | int64(r.bufBytes[idx+1])<<8 | int64(r.bufBytes[idx+2])<<16 | int64(r.bufBytes[idx+3])<<24 |
		int64(r.bufBytes[idx+4])<<32 | int64(r.bufBytes[idx+5])<<40 | int64(r.bufBytes[idx+6])<<48 | int64(r.bufBytes[idx+7])<<56
}

// ReadUInt8 will read uint8 at index
func (r *Reader) ReadUInt8() uint8 {
	idx := r.currentReadIndex
	r.currentReadIndex++
	return uint8(r.bufBytes[idx])
}

// ReadUInt16 will read uint16 at index
func (r *Reader) ReadUInt16() uint16 {
	idx := r.currentReadIndex
	r.currentReadIndex += 2
	return uint16(r.bufBytes[idx]) | uint16(r.bufBytes[idx+1])<<8
}

// ReadUInt32 will read uint32 at index
func (r *Reader) ReadUInt32() uint32 {
	idx := r.currentReadIndex
	r.currentReadIndex += 4
	return uint32(r.bufBytes[idx]) | uint32(r.bufBytes[idx+1])<<8 | uint32(r.bufBytes[idx+2])<<16 | uint32(r.bufBytes[idx+3])<<24
}

// ReadUInt64 will read uint64 at index
func (r *Reader) ReadUInt64() uint64 {
	idx := r.currentReadIndex
	r.currentReadIndex += 8
	return uint64(r.bufBytes[idx]) | uint64(r.bufBytes[idx+1])<<8 | uint64(r.bufBytes[idx+2])<<16 | uint64(r.bufBytes[idx+3])<<24 |
		uint64(r.bufBytes[idx+4])<<32 | uint64(r.bufBytes[idx+5])<<40 | uint64(r.bufBytes[idx+6])<<48 | uint64(r.bufBytes[idx+7])<<56
}
