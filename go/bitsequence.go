package bitsequence

import (
	"math"
)

func BitSequence(bytes []byte, bitStart uint32, bitLength uint32) uint32 {
	if bitLength > 32 {
		// does this constraint deserve an Error?
		bitLength = 32
	}
	startOffset := bitStart % 8
	byteCount := uint32(math.Ceil((float64(startOffset) + float64(bitLength)) / 8))
	byteStart := bitStart >> 3
	endOffset := byteCount*8 - bitLength - startOffset

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

	return result
}
