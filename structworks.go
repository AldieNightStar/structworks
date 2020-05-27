package structworks

import (
	"bytes"
	"encoding/binary"
)

// Converts struct object into []byte
func StructToBytes(object interface{}) (dat []byte, err error) {
	// Convert Structure to bytes
	buf := &bytes.Buffer{}
	err = binary.Write(buf, binary.BigEndian, object)
	dat = buf.Bytes()
	return
}

// Converts []byte back to struct
func BytesToStruct(data []byte, outObject interface{}) (err error) {
	// Read bytes into an object
	buf := bytes.NewReader(data)
	size1 := len(data)
	size2 := binary.Size(outObject)
	if size2 > size1 {
		newData := make([]byte, size2)
		for i := 0; i < len(data); i++ {
			newData[i] = data[i]
		}

		buf = bytes.NewReader(newData)
	}
	err = binary.Read(buf, binary.BigEndian, outObject)
	return
}

// Unsafe convert one structure to another, using []byte under the hood
func StructToStruct(inObject interface{}, outObject interface{}) (err error) {
	data, err := StructToBytes(inObject)
	if err != nil {
		return
	}
	err = BytesToStruct(data, outObject)
	if err != nil {
		print(err.Error())
	}
	return
}
