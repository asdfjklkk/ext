package ext

import "encoding/base64"

type Bytes []byte

func ToBytes(value []byte) Bytes {
	return Bytes(value)
}

func (data Bytes) Bytes() []byte {
	return data
}

func (data Bytes) ToString() StringExt {
	return ToString(string(data.Bytes()))
}

func (data Bytes) Length() int {
	return len(data)
}

func (data Bytes) EncodeBase64() StringExt {
	e64 := base64.StdEncoding.EncodeToString(data)
	return ToString(e64)
}
