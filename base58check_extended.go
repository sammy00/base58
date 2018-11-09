package base58

import (
	"bytes"
)

// CheckDecodeX extends CheckDecode by allowing more version bytes
func CheckDecodeX(input string, versionLen int) ([]byte, []byte, error) {
	decoded := Decode(input)

	if len(decoded) < ChecksumLen+versionLen {
		return nil, nil, ErrInvalidFormat
	}

	if cksum := Checksum(decoded[:len(decoded)-4]); !bytes.Equal(
		cksum[:], decoded[len(decoded)-4:]) {
		return nil, nil, ErrChecksum
	}

	version := decoded[:versionLen]
	payload := decoded[versionLen : len(decoded)-4]

	return payload, version, nil
}

// CheckEncodeX extends CheckEncode by allowing more version bytes
func CheckEncodeX(version, input []byte) string {
	b := make([]byte, 0, len(version)+len(input)+4)
	b = append(b, version...)
	b = append(b, input...)

	cksum := Checksum(b)
	b = append(b, cksum[:]...)

	return Encode(b)
}
