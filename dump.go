// Package utils
package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
)

// DumpToString interface to string
func DumpToString(v interface{}) string {

	str, ok := v.(string)
	if !ok {
		buff := &bytes.Buffer{}
		json.NewEncoder(buff).Encode(v)
		return buff.String()
	}

	return str
}

// DebugPrint for debug print on terminal
func DebugPrint(v ...interface{}) {
	for _, x := range v {
		fmt.Println(DumpToString(x))
	}
}

// Data To json Bytes
func ToJSONByte(v interface{}) []byte {

	switch v.(type) {
	case []byte:
		return v.([]byte)
	case string:
		return []byte(v.(string))
	}

	buff := &bytes.Buffer{}
	json.NewEncoder(buff).Encode(v)
	return buff.Bytes()
}

// StreamToByte cast
func StreamToByte(stream io.Reader) []byte {
	buf := new(bytes.Buffer)
	buf.ReadFrom(stream)
	return buf.Bytes()
}

// StreamToString cast
func StreamToString(stream io.Reader) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(stream)
	return buf.String()
}
