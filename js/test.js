const assert = require('assert')
const bitSequence = require('./bit-sequence.js')

function binaryStringToBytes (s) {
  const byteLength = Math.ceil(s.length / 8)
  s = s.padStart(byteLength * 8, '0')
  const bytes = new Uint8Array(byteLength)
  for (let i = 0; i < byteLength; i++) {
    bytes[i] = parseInt(s.substring(i * 8, (i + 1) * 8), 2)
  }
  return bytes
}

// sanity check binaryStringToBytes
;[
  '00000001:01',
  '10000000:80',
  '11111111:ff',
  '11000000:c0',
  '11110000:f0',
  '1111111111111111:ffff',
  '0000000000000001:0001',
  '000000000000000000000001:000001',
  '111111111111111111111111:ffffff',
  '100000001000000010000000:808080',
  '10000000100000001000000010000000:80808080',
  '10000000111111111000000011111111:80ff80ff',
  '001:01',
  '111:07',
  '1111:0f',
  '01111:0f',
  '001111:0f',
  '0000001111:000f',
  '10000000000001111:01000f'
].forEach((s) => {
  const [ binary, hex ] = s.split(':')
  const bytesHex = Array.prototype.map.call(
    binaryStringToBytes(binary),
    (b) => b.toString(16).padStart(2, '0')
  ).join('')
  // console.log(binary, '=>', binaryStringToBytes(binary).toString('hex'), '<>', hex, '?', parseInt(binary, 2).toString(16))
  assert.strictEqual(bytesHex, hex)
  // sanity check the sanity check by reading the binary
  assert.strictEqual(bytesHex.replace(/^0+/, ''), parseInt(binary, 2).toString(16))
})

let asserts = 0
;[
  '00000001',
  '11111111',
  '01010101',
  '10001000',
  '0000000000000001',
  '0000000100000001',
  '1111111111111111',
  '1010101010101010',
  '0101010101010101',
  '1001001001001001',
  '0100100100100100',
  '1000100010001000',
  '0100010001000100',
  '1111111100000000',
  '0000000011111111',
  '0000111111110000',
  '000000000000000000000001',
  '111111111111111111111111',
  '001001100100110010011111',
  '00000000000000000000000000000001',
  '11111111111111111111111111111111',
  '10000000000000000000000000000001',
  '10001100010110011110001010111010',
  '1111111111111111111111111111111111111111111111111111111111111111',
  '0000000000000000000000000000000000000000000000000000000000000001',
  '1000000000000000000000000000000000000000000000000000000000000001',
  '1010110100111110101110101001010001000000001101011101110010011111',
  '11111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111'
].forEach((s) => {
  const bytes = binaryStringToBytes(s)

  for (let start = 0; start < bytes.length * 8; start++) {
    for (let length = 1; length <= bytes.length * 8 - start; length++) {
      const expected = parseInt(s.substring(start, start + length), 2)
      const actual = bitSequence(bytes, start, length)
      // console.log(`[${s}] start=${start} length=${length} ${actual} <> ${expected}`)
      asserts++
      assert.strictEqual(actual, expected, `[${s}] start=${start} length=${length} ${actual} <> ${expected}`)
      assert.ok(actual >= 0, 'sanity check that we\'re only dealing with unsigned integers')
    }
  }
})

assert.ok(asserts > 10000, 'did a lot of asserts')
