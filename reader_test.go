package packet

import "testing"

func TestReadString(t *testing.T) {
	want := "abc"

	w := NewWriter()
	w.WriteString(want, 0)

	r := NewReader(w.Bytes())
	got := r.ReadString()

	if got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestReadBooleanTrue(t *testing.T) {
	want := true

	w := NewWriter()
	w.WriteBoolean(want, 0)

	r := NewReader(w.Bytes())
	got := r.ReadBoolean()

	if got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestReadBooleanFalse(t *testing.T) {
	want := false

	w := NewWriter()
	w.WriteBoolean(want, 0)

	r := NewReader(w.Bytes())
	got := r.ReadBoolean()

	if got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestReadFloat32(t *testing.T) {
	want := float32(20.2)

	w := NewWriter()
	w.WriteFloat32(want, 0)

	r := NewReader(w.Bytes())
	got := r.ReadFloat32()

	if got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestReadFloat64(t *testing.T) {
	want := float64(20.2)

	w := NewWriter()
	w.WriteFloat64(want, 0)

	r := NewReader(w.Bytes())
	got := r.ReadFloat64()

	if got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestReadInt8(t *testing.T) {
	want := int8(20)

	w := NewWriter()
	w.WriteInt8(want, 0)

	r := NewReader(w.Bytes())
	got := r.ReadInt8()

	if got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestReadInt16(t *testing.T) {
	want := int16(20)

	w := NewWriter()
	w.WriteInt16(want, 0)

	r := NewReader(w.Bytes())
	got := r.ReadInt16()

	if got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestReadInt32(t *testing.T) {
	want := int32(20)

	w := NewWriter()
	w.WriteInt32(want, 0)

	r := NewReader(w.Bytes())
	got := r.ReadInt32()

	if got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestReadInt64(t *testing.T) {
	want := int64(20)

	w := NewWriter()
	w.WriteInt64(want, 0)

	r := NewReader(w.Bytes())
	got := r.ReadInt64()

	if got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestReadUInt8(t *testing.T) {
	want := uint8(20)

	w := NewWriter()
	w.WriteUInt8(want, 0)

	r := NewReader(w.Bytes())
	got := r.ReadUInt8()

	if got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestReadUInt16(t *testing.T) {
	want := uint16(20)

	w := NewWriter()
	w.WriteUInt16(want, 0)

	r := NewReader(w.Bytes())
	got := r.ReadUInt16()

	if got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestReadUInt32(t *testing.T) {
	want := uint32(20)

	w := NewWriter()
	w.WriteUInt32(want, 0)

	r := NewReader(w.Bytes())
	got := r.ReadUInt32()

	if got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestReadUInt64(t *testing.T) {
	want := uint64(20)

	w := NewWriter()
	w.WriteUInt64(want, 0)

	r := NewReader(w.Bytes())
	got := r.ReadUInt64()

	if got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}
