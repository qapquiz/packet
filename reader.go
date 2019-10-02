package packet

import (
	"bytes"
	"encoding/binary"
	"log"
)

// Reader will contain bytes.Reader
type Reader struct {
	value *bytes.Reader
}

// NewPlainReader will create Writer without cut header
// NewPlainReader :: []byte -> *Reader
func NewPlainReader(in []byte) *Reader {
	return &Reader{
		value: bytes.NewReader(in),
	}
}

// NewReader will create Reader with data
func NewReader(in []byte) *Reader {
	return &Reader{
		value: bytes.NewReader(in[2:]),
	}
}

// ReadString will read string
func (pr *Reader) ReadString() string {
	message := make([]byte, 0)

	for {
		charCode, err := pr.value.ReadByte()
		if err != nil {
			log.Fatal("cannot read bytes in packet.ReadString()")
		}

		if uint8(charCode) == 0 {
			break
		}

		message = append(message, charCode)
	}

	return string(message[:])
}

// ReadUInt8 will read uint8
func (pr *Reader) ReadUInt8() uint8 {
	number, err := pr.value.ReadByte()
	if err != nil {
		log.Println("ReadUInt8 Error!")
	}

	return uint8(number)
}

// ReadUInt16 will read uint16
func (pr *Reader) ReadUInt16() uint16 {
	var number uint16
	err := binary.Read(pr.value, binary.LittleEndian, &number)
	if err != nil {
		log.Println("binary.Read uint16 failed: ", err)
	}

	return number
}

// ReadUInt32 will read uint32
func (pr *Reader) ReadUInt32() uint32 {
	var number uint32
	err := binary.Read(pr.value, binary.LittleEndian, &number)
	if err != nil {
		log.Println("binary.Read uint32 failed: ", err)
	}

	return number
}

// ReadUInt64 will read uint64
func (pr *Reader) ReadUInt64() uint64 {
	var number uint64
	err := binary.Read(pr.value, binary.LittleEndian, &number)
	if err != nil {
		log.Println("binary.Read uint64 failed: ", err)
	}

	return number
}

// ReadInt8 will read int8
func (pr *Reader) ReadInt8() int8 {
	number, err := pr.value.ReadByte()
	if err != nil {
		log.Println("ReadInt8 Error!")
	}

	return int8(number)
}

// ReadInt16 will read int16
func (pr *Reader) ReadInt16() int16 {
	var number int16
	err := binary.Read(pr.value, binary.LittleEndian, &number)
	if err != nil {
		log.Println("binary.Read int16 failed: ", err)
	}

	return number
}

// ReadInt32 will read int32
func (pr *Reader) ReadInt32() int32 {
	var number int32
	err := binary.Read(pr.value, binary.LittleEndian, &number)
	if err != nil {
		log.Println("binary.Read int32 failed: ", err)
	}

	return number
}

// ReadInt64 will read int64
func (pr *Reader) ReadInt64() int64 {
	var number int64
	err := binary.Read(pr.value, binary.LittleEndian, &number)
	if err != nil {
		log.Println("binary.Read int64 failed: ", err)
	}

	return number
}

// ReadFloat32 will read float32
func (pr *Reader) ReadFloat32() float32 {
	var number float32
	err := binary.Read(pr.value, binary.LittleEndian, &number)
	if err != nil {
		log.Println("binary.Read float32 failed: ", err)
	}

	return number
}

// ReadFloat64 will read float64
func (pr *Reader) ReadFloat64() float64 {
	var number float64
	err := binary.Read(pr.value, binary.LittleEndian, &number)
	if err != nil {
		log.Println("binary.Read float64 failed: ", err)
	}

	return number
}

// ReadBoolean will read bool
func (pr *Reader) ReadBoolean() bool {
	number, err := pr.value.ReadByte()
	if err != nil {
		log.Println("ReadBoolean Error!")
	}

	return uint8(number) == 1
}

// Complete will return true if data reach zero
func (pr *Reader) Complete() bool {
	if pr.value.Len() == 0 {
		return true
	}

	return false
}
