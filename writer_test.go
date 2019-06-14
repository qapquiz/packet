package packet

import (
	"reflect"
	"testing"
)

func TestWriteUInt8(t *testing.T) {
	t.Parallel()
	var expectedResult = []byte{17, 39, 20}

	packetWriter := NewWriter(10001)
	packetWriter.WriteUInt8(uint8(20))
	actualResult := packetWriter.GetDataWithoutRemoveHeader()

	if !compareBytes(expectedResult, actualResult) {
		t.Errorf("Expected %v, actual %v", expectedResult, actualResult)
	}
}

func TestWriteUInt16(t *testing.T) {
	t.Parallel()
	var expectedResult = []byte{17, 39, 20, 0}

	packetWriter := NewWriter(10001)
	packetWriter.WriteUInt16(uint16(20))
	actualResult := packetWriter.GetDataWithoutRemoveHeader()

	if !compareBytes(expectedResult, actualResult) {
		t.Errorf("Expected %v, actual %v", expectedResult, actualResult)
	}
}

func TestWriteUInt32(t *testing.T) {
	t.Parallel()
	var expectedResult = []byte{17, 39, 20, 0, 0, 0}

	packetWriter := NewWriter(10001)
	packetWriter.WriteUInt32(uint32(20))
	actualResult := packetWriter.GetDataWithoutRemoveHeader()

	if !compareBytes(expectedResult, actualResult) {
		t.Errorf("Expected %v, actual %v", expectedResult, actualResult)
	}
}

func TestWriteUInt64(t *testing.T) {
	t.Parallel()
	var expectedResult = []byte{17, 39, 20, 0, 0, 0, 0, 0, 0, 0}

	packetWriter := NewWriter(10001)
	packetWriter.WriteUInt64(uint64(20))
	actualResult := packetWriter.GetDataWithoutRemoveHeader()

	if !compareBytes(expectedResult, actualResult) {
		t.Errorf("Expected %v, actual %v", expectedResult, actualResult)
	}
}

func TestWriteInt8(t *testing.T) {
	t.Parallel()
	var expectedResult = []byte{17, 39, 20}

	packetWriter := NewWriter(10001)
	packetWriter.WriteInt8(int8(20))
	actualResult := packetWriter.GetDataWithoutRemoveHeader()

	if !compareBytes(expectedResult, actualResult) {
		t.Errorf("Expected %v, actual %v", expectedResult, actualResult)
	}
}

func TestWriteInt16(t *testing.T) {
	t.Parallel()
	var expectedResult = []byte{17, 39, 20, 0}

	packetWriter := NewWriter(10001)
	packetWriter.WriteInt16(int16(20))
	actualResult := packetWriter.GetDataWithoutRemoveHeader()

	if !compareBytes(expectedResult, actualResult) {
		t.Errorf("Expected %v, actual %v", expectedResult, actualResult)
	}
}

func TestWriteInt32(t *testing.T) {
	t.Parallel()
	var expectedResult = []byte{17, 39, 20, 0, 0, 0}

	packetWriter := NewWriter(10001)
	packetWriter.WriteInt32(int32(20))
	actualResult := packetWriter.GetDataWithoutRemoveHeader()

	if !compareBytes(expectedResult, actualResult) {
		t.Errorf("Expected %v, actual %v", expectedResult, actualResult)
	}
}

func TestWriteInt64(t *testing.T) {
	t.Parallel()
	var expectedResult = []byte{17, 39, 20, 0, 0, 0, 0, 0, 0, 0}

	packetWriter := NewWriter(10001)
	packetWriter.WriteInt64(int64(20))
	actualResult := packetWriter.GetDataWithoutRemoveHeader()

	if !compareBytes(expectedResult, actualResult) {
		t.Errorf("Expected %v, actual %v", expectedResult, actualResult)
	}
}

func TestWriteFloat32(t *testing.T) {
	t.Parallel()
	var expectedResult = []byte{17, 39, 205, 204, 204, 61}

	packetWriter := NewWriter(10001)
	packetWriter.WriteFloat32(0.1)
	actualResult := packetWriter.GetDataWithoutRemoveHeader()

	if !compareBytes(expectedResult, actualResult) {
		t.Errorf("Expected %v, actual %v", expectedResult, actualResult)
	}
}

func TestWriteFloat64(t *testing.T) {
	t.Parallel()
	var expectedResult = []byte{17, 39, 154, 153, 153, 153, 153, 153, 185, 63}

	packetWriter := NewWriter(10001)
	packetWriter.WriteFloat64(0.1)
	actualResult := packetWriter.GetDataWithoutRemoveHeader()

	if !compareBytes(expectedResult, actualResult) {
		t.Errorf("Expected %v, actual %v", expectedResult, actualResult)
	}
}

func TestWriteBoolean(t *testing.T) {
	t.Parallel()
	var expectedResult = []byte{17, 39, 1}

	packetWriter := NewWriter(10001)
	packetWriter.WriteBoolean(true)
	actualResult := packetWriter.GetDataWithoutRemoveHeader()

	if !compareBytes(expectedResult, actualResult) {
		t.Errorf("Expected %v, actual %v", expectedResult, actualResult)
	}
}

func TestWriteString(t *testing.T) {
	t.Parallel()
	var expectedResult = []byte{17, 39, 4, 0, 84, 101, 115, 116}

	packetWriter := NewWriter(10001)
	packetWriter.WriteString("Test")
	actualResult := packetWriter.GetDataWithoutRemoveHeader()

	if !compareBytes(expectedResult, actualResult) {
		t.Errorf("Expected %v, actual %v", expectedResult, actualResult)
	}
}

func compareBytes(firstBytes []byte, secondBytes []byte) bool {
	return reflect.DeepEqual(firstBytes, secondBytes)
}
