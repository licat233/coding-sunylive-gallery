package utils

import (
	"bytes"
	"encoding/binary"
)

func IntToBytes(n int) ([]byte, error) {
	data := int64(n)
	bytebuf := bytes.NewBuffer([]byte{})
	err := binary.Write(bytebuf, binary.BigEndian, data)
	return bytebuf.Bytes(), err
}

func BytesToInt(bys []byte) (int, error) {
	bytebuff := bytes.NewBuffer(bys)
	var data int64
	err := binary.Read(bytebuff, binary.BigEndian, &data)
	return int(data), err
}

func BytesToInt64(bys []byte) (int64, error) {
	bytebuff := bytes.NewBuffer(bys)
	var data int64
	err := binary.Read(bytebuff, binary.BigEndian, &data)
	return data, err
}
