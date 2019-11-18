package packet

import (
	"testing"
	"time"
)

var a = A{
	Name:     "armariya",
	BirthDay: time.Unix(4309820, 0),
	Phone:    "0812345678",
	Siblings: 2,
	Spouse:   true,
	Money:    float64(2034.80),
}

var aBytes = a.Marshal()

type A struct {
	Name     string
	BirthDay time.Time
	Phone    string
	Siblings int
	Spouse   bool
	Money    float64
}

func (a A) Marshal() []byte {
	w := NewWriter()

	idx := 0
	idx += w.WriteString(a.Name, idx)
	idx += w.WriteInt64(a.BirthDay.Unix(), idx)
	idx += w.WriteString(a.Phone, idx)
	idx += w.WriteInt32(int32(a.Siblings), idx)
	idx += w.WriteBoolean(a.Spouse, idx)
	idx += w.WriteFloat64(a.Money, idx)

	return w.Bytes()
}

func (a *A) Unmarshal(in []byte) {
	r := NewReader(in)

	a.Name = r.ReadString()
	a.BirthDay = time.Unix(r.ReadInt64(), 0)
	a.Phone = r.ReadString()
	a.Siblings = int(r.ReadInt32())
	a.Spouse = r.ReadBoolean()
	a.Money = r.ReadFloat64()
}

func BenchmarkMarshalA(b *testing.B) {
	for n := 0; n < b.N; n++ {
		a.Marshal()
	}
}

func BenchmarkUnmarshalA(b *testing.B) {
	for n := 0; n < b.N; n++ {
		a.Unmarshal(aBytes)
	}
}
