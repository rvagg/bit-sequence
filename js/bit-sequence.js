/**
 * @param {Uint8Array} bytes
 * @param {number} bitStart
 * @param {number} bitLength
 * @returns {number}
 */
export default function bitSequence (bytes, bitStart, bitLength) {
  // making an assumption that bytes is an Array-like that will give us one
  // byte per element, so either an Array full of 8-bit integers or a
  // Uint8Array or a Node.js Buffer, or something like that

  const startOffset = bitStart % 8
  const byteCount = Math.ceil((startOffset + bitLength) / 8)
  const byteStart = bitStart >> 3
  const endOffset = byteCount * 8 - bitLength - startOffset

  let result = 0

  for (let i = 0; i < byteCount; i++) {
    let local = bytes[byteStart + i]
    let shift = 0
    let localBitLength = 8 // how many bits of this byte we are extracting

    if (i === 0) {
      localBitLength -= startOffset
    }

    if (i === byteCount - 1) {
      localBitLength -= endOffset
      shift = endOffset
      local >>= shift // take off the trailing bits
    }

    if (localBitLength < 8) {
      const mask = (1 << localBitLength) - 1
      local &= mask // mask off the leading bits
    }

    if (i < 3) {
      if (shift < 8) {
        result = result << (8 - shift)
      }
      result |= local
    } else {
      // once we start shifting beyond the 24-bit range we get to signed integers
      // and our bitwise operations break down, because JavaScript. But if we do
      // it without bitwise operations then we can cheat into very large numbers
      if (shift < 8) {
        result = result * Math.pow(2, (8 - shift))
      }
      result += local
    }
  }

  return result
}
