package gobtools

import (
	"bytes"
	"encoding/gob"
)

func Serialize(value interface{}) ([]byte, error) {
	buf := bytes.Buffer{}
	enc := gob.NewEncoder(&buf)
	// gob.Register(value)

	err := enc.Encode(&value)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func Deserialize(valueBytes []byte) (interface{}, error) {
	var value interface{}
	buf := bytes.NewBuffer(valueBytes)
	dec := gob.NewDecoder(buf)

	err := dec.Decode(&value)
	if err != nil {
		return nil, err
	}

	return value, nil
}
