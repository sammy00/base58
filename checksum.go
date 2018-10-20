package base58

import "crypto/sha256"

// Checksum calculates a CRC for the input data as the first 4 bytes of sha256^2
func Checksum(input []byte) [4]byte {
	h := sha256.Sum256(input)
	h2 := sha256.Sum256(h[:])

	var cksum [4]byte

	copy(cksum[:], h2[:4])
	return cksum
}
