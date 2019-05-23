package packet

import (
	"bytes"
	"testing"
)

func TestReadUInt8(t *testing.T) {
	t.Parallel()
	expectedResult := uint8(20)

	var data = []byte{20}

	packetReader := Reader{
		value: bytes.NewReader(data),
	}

	actualResult := packetReader.ReadUInt8()

	if expectedResult != actualResult && packetReader.value.Len() != 0 {
		t.Errorf("Expected %d, actual %d", expectedResult, actualResult)
	}
}

func TestReadUInt16(t *testing.T) {
	t.Parallel()
	expectedResult := uint16(10001)

	var data = []byte{17, 39}

	packetReader := Reader{
		value: bytes.NewReader(data),
	}

	actualResult := packetReader.ReadUInt16()

	if expectedResult != actualResult && packetReader.value.Len() != 0 {
		t.Errorf("Expected %d, actual %d", expectedResult, actualResult)
	}
}

func TestReadUInt32(t *testing.T) {
	t.Parallel()
	expectedResult := uint32(20)

	var data = []byte{20, 0, 0, 0}

	packetReader := Reader{
		value: bytes.NewReader(data),
	}

	actualResult := packetReader.ReadUInt32()

	if expectedResult != actualResult {
		t.Errorf("Expected %d, actual %d", expectedResult, actualResult)
	}
}

func TestReadUInt64(t *testing.T) {
	t.Parallel()
	expectedResult := uint64(20)

	var data = []byte{20, 0, 0, 0, 0, 0, 0, 0}

	packetReader := Reader{
		value: bytes.NewReader(data),
	}

	actualResult := packetReader.ReadUInt64()

	if expectedResult != actualResult {
		t.Errorf("Expected %d, actual %d", expectedResult, actualResult)
	}
}

func TestReadInt8(t *testing.T) {
	t.Parallel()
	expectedResult := int8(-20)

	var data = []byte{236}

	packetReader := Reader{
		value: bytes.NewReader(data),
	}

	actualResult := packetReader.ReadInt8()

	if expectedResult != actualResult && packetReader.value.Len() != 0 {
		t.Errorf("Expected %d, actual %d", expectedResult, actualResult)
	}
}

func TestReadInt16(t *testing.T) {
	t.Parallel()
	expectedResult := int16(-10001)

	var data = []byte{239, 216}

	packetReader := Reader{
		value: bytes.NewReader(data),
	}

	actualResult := packetReader.ReadInt16()

	if expectedResult != actualResult && packetReader.value.Len() != 0 {
		t.Errorf("Expected %d, actual %d", expectedResult, actualResult)
	}
}

func TestReadInt32(t *testing.T) {
	t.Parallel()
	expectedResult := int32(-20)

	var data = []byte{236, 255, 255, 255}

	packetReader := Reader{
		value: bytes.NewReader(data),
	}

	actualResult := packetReader.ReadInt32()

	if expectedResult != actualResult {
		t.Errorf("Expected %d, actual %d", expectedResult, actualResult)
	}
}

func TestReadInt64(t *testing.T) {
	t.Parallel()
	expectedResult := int64(-20)

	var data = []byte{236, 255, 255, 255, 255, 255, 255, 255}

	packetReader := Reader{
		value: bytes.NewReader(data),
	}

	actualResult := packetReader.ReadInt64()

	if expectedResult != actualResult {
		t.Errorf("Expected %d, actual %d", expectedResult, actualResult)
	}
}

func TestReadFloat32(t *testing.T) {
	t.Parallel()
	expectedResult := float32(-20.2)

	var data = []byte{154, 153, 161, 193}

	packetReader := Reader{
		value: bytes.NewReader(data),
	}

	actualResult := packetReader.ReadFloat32()

	if expectedResult != actualResult {
		t.Errorf("Expected %f, actual %f", expectedResult, actualResult)
	}
}

func TestReadFloat64(t *testing.T) {
	t.Parallel()
	expectedResult := float64(-20.2)

	var data = []byte{51, 51, 51, 51, 51, 51, 52, 192}

	packetReader := Reader{
		value: bytes.NewReader(data),
	}

	actualResult := packetReader.ReadFloat64()

	if expectedResult != actualResult {
		t.Errorf("Expected %f, actual %f", expectedResult, actualResult)
	}
}

func TestReadString(t *testing.T) {
	t.Parallel()
	expectedResult := "test"

	var data = []byte{4, 0, 116, 101, 115, 116}

	packetReader := Reader{
		value: bytes.NewReader(data),
	}

	actualResult := packetReader.ReadString()

	if expectedResult != actualResult {
		t.Errorf("Expected %s, actual %s", expectedResult, actualResult)
	}
}

func TestBooleanWhenOneMustBeTrue(t *testing.T) {
	t.Parallel()
	expectedResult := true

	var data = []byte{1}

	packetReader := Reader{
		value: bytes.NewReader(data),
	}

	actualResult := packetReader.ReadBoolean()

	if expectedResult != actualResult {
		t.Errorf("Expected %t, actual %t", expectedResult, actualResult)
	}
}

func TestBooleanWhenZeroMustBeFalse(t *testing.T) {
	t.Parallel()
	expectedResult := false

	var data = []byte{0}

	packetReader := Reader{
		value: bytes.NewReader(data),
	}

	actualResult := packetReader.ReadBoolean()

	if expectedResult != actualResult {
		t.Errorf("Expected %t, actual %t", expectedResult, actualResult)
	}
}
