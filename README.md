# bit-sequence

**Turn an arbitrary sequence of bits from a byte array and turn it into an integer**

* [JavaScript](#javascript)
  * [Example](#example)
  * [API](#api)
* [Go](#go)
* [License and Copyright](#license-and-copyright)

## JavaScript

[![NPM](https://nodei.co/npm/bit-sequence.svg)](https://nodei.co/npm/bit-sequence/)

Given an `Array`-like containing bytes (unsigned 8-bit integers), extract an arbitrary sequence of the underlying bits and convert them into an unsigned integer value.

Useful for cases where a sub-sequence of bits within a longer byte sequence is used to form an index, such as in a [hash array mapped trie](https://en.wikipedia.org/wiki/Hash_array_mapped_trie) where an index at each level of the tree structure is formed by incremental chunks of a key's hash.

### Example

```js
const bitSequence = require('bit-sequence')
const assert = require('assert')
const bytes = new Uint8Array([ 0b00010101, 0b10101000, 0b00000000, 0b00000000 ])
//               extract bits from here ^         to here ^
const int = bitSequence(bytes, 7, 11)
assert.strictEqual(int, 0b11010100000) // or `1696`
```

### API

**`bitSequence(bytes, start, length)`**

* `bytes` is an `Array`-like containing bytes. It's assumed that it can be accessed with an array indexing operator and that it will return 8-bit integers (`< 255`). A `Uint8Array` or Node.js `Buffer` fits into this category, as does a standard `Array`, however the 8-bit assumption is not enforced so an `Array` full of larger integers will yield undefined results.
* `start` is an integer _bit_ index to start extraction (not a _byte_ index).
* `length` is the number of bits to extract from the `bytes` array.

Returns an **unsigned integer** version of the bit sequence. The most significant bit is not interpreted for a two's compliment representation. You should only get positive numbers `>= 0`.

As per the example above, the assumption is that we are extracting bytes where the least significant bit is to the right, so we extract in the same order as presented by binary literals.

JavaScript's crazy numbers allows us to extract potentially very large bit sequences, but the usual caveats apply at 32-bits and beyond [`Number.MAX_SAFE_INTEGER`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Number/MAX_SAFE_INTEGER).

## Go

`github.com/rvagg/bit-sequence/go`

### API

Exports `BitSequence(bytes []byte, bitStart uint32, bitLength uint32) (uint32, error)` where:

* `bytes` is a simple byte array of arbitrary length
* `bitStart` is a _bit_ index to start extraction (not a _byte_ index).
* `bitLength` is the number of bits to extract from the `bytes` array. This value can only be a maximum of `32`, higher values will be adjusted down.

Returns an unsigned integer version of the bit sequence. The most significant bit is not interpreted for a two's compliment representation.
Returns an error if `bitLength` exceeds 32 or the requested sequence overflows the bytes boundary.

## License and Copyright

Copyright 2019 Rod Vagg

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License. You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.
