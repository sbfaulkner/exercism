package variablelengthquantity

import (
	"errors"
)

const testVersion = 4

// DecodeVarint decodes a variable int stream of bytes into 32-bit integers
func DecodeVarint(encoded []byte) ([]uint32, error) {
	complete := true
	decoded := []uint32{}

	for _, b := range encoded {
		if complete {
			decoded = append(decoded, 0)
		}

		complete = (b & 0x80) == 0

		decoded[len(decoded)-1] <<= 7
		decoded[len(decoded)-1] |= uint32(b & 0x7f)
	}

	if !complete {
		return []uint32{}, errors.New("variablelengthquantity: incomplete sequence")
	}

	return decoded, nil
}

// EncodeVarint encodes 32-bit integer as a variable int stream of bytes
func EncodeVarint(decoded []uint32) []byte {
	encoded := []byte{}

	for _, u := range decoded {
		e := []byte{byte(u & 0x7f)}

		for {
			u >>= 7

			if u == 0 {
				break
			}

			e = append([]byte{byte(u&0x7f) | 0x80}, e...)
		}

		encoded = append(encoded, e...)
	}

	return encoded
}
