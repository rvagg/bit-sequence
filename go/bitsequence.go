package bitsequence

import "fmt"

// BitSequence turns an arbitrary sequence of bits from a byte array into an integer.
//
// - `bytes` is a simple byte array of arbitrary length
//
// - `bitStart` is a _bit_ index to start extraction (not a _byte_ index).
//
// - `bitLength` is the number of bits to extract from the `bytes` array. This value can only be a maximum of `32`, higher values will be adjusted down.
//
// Returns an unsigned integer version of the bit sequence. The most significant bit is not interpreted for a two's compliment representation.
// Returns an error if `bitLength` exceeds 32 or the requested sequence overflows the bytes boundary.
func BitSequence(bytes []byte, bitStart uint32, bitLength uint32) (uint32, error) {
	if bitLength > 32 {
		return 0, fmt.Errorf("maximum bits that can be read is 32")
	}
	startOffset := bitStart % 8
	byteCount := (7 + startOffset + bitLength) / 8
	byteStart := bitStart / 8
	endOffset := byteCount*8 - bitLength - startOffset

	if int(byteStart+byteCount) > len(bytes) {
		return 0, fmt.Errorf("cannot read past end of bytes array")
	}

	var result uint32

	var i uint32
	for i = 0; i < byteCount; i++ {
		local := bytes[byteStart+i]
		var shift uint32
		var localBitLength uint32 = 8 // how many bits of this byte we are extracting

		if i == 0 {
			localBitLength -= startOffset
		}

		if i == byteCount-1 {
			localBitLength -= endOffset
			shift = endOffset
			local = local >> shift // take off the trailing bits
		}

		if localBitLength < 8 {
			mask := uint8((1 << localBitLength) - 1)
			local = local & mask // mask off the leading bits
		}

		if shift < 8 {
			result = result << (8 - shift)
		}
		result = result | uint32(local)
	}

	return result, nil
}
