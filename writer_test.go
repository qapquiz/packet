package packet

import (
	"bytes"
	"encoding/binary"
	"testing"
)

func TestWriteString(t *testing.T) {
	w := NewWriter()

	w.WriteString("abc")

	got := w.Bytes()
	want := []byte{3, 0, 0, 0, 0, 0, 0, 0, 97, 98, 99}

	if got[0] != want[0] && got[8] != want[8] && got[9] != want[9] && got[10] != want[10] {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestWriteStringSuffixZero(t *testing.T) {
	w := NewWriter()

	w.WriteString("abc")

	got := w.Bytes()
	want := []byte{97, 98, 99, 0}

	if got[0] != want[0] && got[1] != want[1] && got[2] != want[2] && got[3] != want[3] {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestWriteTrueBoolean(t *testing.T) {
	w := NewWriter()

	w.WriteBoolean(true)

	buf := bytes.NewBuffer([]byte{})
	err := binary.Write(buf, binary.LittleEndian, true)
	if err != nil {
		t.Error(err)
	}

	got := w.Bytes()
	want := buf.Bytes()

	if got[0] != want[0] {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestWriteFalseBoolean(t *testing.T) {
	w := NewWriter()

	w.WriteBoolean(false)

	buf := bytes.NewBuffer([]byte{})
	err := binary.Write(buf, binary.LittleEndian, false)
	if err != nil {
		t.Error(err)
	}

	got := w.Bytes()
	want := buf.Bytes()

	if got[0] != want[0] {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestWriteFloat32(t *testing.T) {
	w := NewWriter()

	w.WriteFloat32(float32(20))

	buf := bytes.NewBuffer([]byte{})
	err := binary.Write(buf, binary.LittleEndian, float32(20))
	if err != nil {
		t.Error(err)
	}

	got := w.Bytes()
	want := buf.Bytes()

	if got[0] != want[0] && got[1] != want[1] && got[2] != want[2] && got[3] != want[3] {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestWriteFloat64(t *testing.T) {
	w := NewWriter()

	w.WriteFloat64(float64(20))

	buf := bytes.NewBuffer([]byte{})
	err := binary.Write(buf, binary.LittleEndian, float64(20))
	if err != nil {
		t.Error(err)
	}

	got := w.Bytes()
	want := buf.Bytes()

	if got[0] != want[0] && got[1] != want[1] && got[2] != want[2] && got[3] != want[3] && got[4] != want[4] && got[5] != want[5] && got[6] != want[6] && got[7] != want[7] {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestWriteInt8(t *testing.T) {
	w := NewWriter()

	w.WriteInt8(int8(20))

	buf := bytes.NewBuffer([]byte{})
	err := binary.Write(buf, binary.LittleEndian, int8(20))
	if err != nil {
		t.Error(err)
	}

	got := w.Bytes()
	want := buf.Bytes()

	if got[0] != want[0] {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestWriteInt16(t *testing.T) {
	w := NewWriter()

	w.WriteInt16(int16(20))

	buf := bytes.NewBuffer([]byte{})
	err := binary.Write(buf, binary.LittleEndian, int16(20))
	if err != nil {
		t.Error(err)
	}

	got := w.Bytes()
	want := buf.Bytes()

	if got[0] != want[0] && got[1] != want[1] {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestWriteInt32(t *testing.T) {
	w := NewWriter()

	w.WriteInt32(int32(20))

	buf := bytes.NewBuffer([]byte{})
	err := binary.Write(buf, binary.LittleEndian, int32(20))
	if err != nil {
		t.Error(err)
	}

	got := w.Bytes()
	want := buf.Bytes()

	if got[0] != want[0] && got[1] != want[1] && got[2] != want[2] && got[3] != want[3] {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestWriteInt64(t *testing.T) {
	w := NewWriter()

	w.WriteInt64(int64(20))

	buf := bytes.NewBuffer([]byte{})
	err := binary.Write(buf, binary.LittleEndian, int64(20))
	if err != nil {
		t.Error(err)
	}

	got := w.Bytes()
	want := buf.Bytes()

	if got[0] != want[0] && got[1] != want[1] && got[2] != want[2] && got[3] != want[3] && got[4] != want[4] && got[5] != want[5] && got[6] != want[6] && got[7] != want[7] {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestWriteUInt8(t *testing.T) {
	w := NewWriter()

	w.WriteUInt8(uint8(20))

	buf := bytes.NewBuffer([]byte{})
	err := binary.Write(buf, binary.LittleEndian, uint8(20))
	if err != nil {
		t.Error(err)
	}

	got := w.Bytes()
	want := buf.Bytes()

	if got[0] != want[0] {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestWriteUInt16(t *testing.T) {
	w := NewWriter()

	w.WriteUInt16(uint16(20))

	buf := bytes.NewBuffer([]byte{})
	err := binary.Write(buf, binary.LittleEndian, uint16(20))
	if err != nil {
		t.Error(err)
	}

	got := w.Bytes()
	want := buf.Bytes()

	if got[0] != want[0] && got[1] != want[1] {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestWriteUInt32(t *testing.T) {
	w := NewWriter()

	w.WriteUInt32(uint32(20))

	buf := bytes.NewBuffer([]byte{})
	err := binary.Write(buf, binary.LittleEndian, uint32(20))
	if err != nil {
		t.Error(err)
	}

	got := w.Bytes()
	want := buf.Bytes()

	if got[0] != want[0] && got[1] != want[1] && got[2] != want[2] && got[3] != want[3] {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestWriteUInt64(t *testing.T) {
	w := NewWriter()

	w.WriteUInt64(uint64(20))

	buf := bytes.NewBuffer([]byte{})
	err := binary.Write(buf, binary.LittleEndian, uint64(20))
	if err != nil {
		t.Error(err)
	}

	got := w.Bytes()
	want := buf.Bytes()

	if got[0] != want[0] && got[1] != want[1] && got[2] != want[2] && got[3] != want[3] && got[4] != want[4] && got[5] != want[5] && got[6] != want[6] && got[7] != want[7] {
		t.Errorf("got: %v, want: %v", got, want)
	}
}
