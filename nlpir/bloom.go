/*
* Copyright ©2013 The nlpt Authors. All rights reserved.
* Use of this source code is governed by a BSD-style
* license that can be found in the LICENSE file.

* Bloom provides a simple bloom filter for doing
* information retrieval
*
 */

package nlpir

// Constants for the Hashing
const (
	FNV_BASIS_64 = uint64(14695981039346656031)
	FNV_PRIME_64 = uint64((1 << 40) + 435)
	FNV_MASK_64  = uint64(^uint64(0) >> 1)
	NUM_BITS     = 64

	FNV_BASIS_32 = uint32(0x811c9dc5)
	FNV_PRIME_32 = uint32((1 << 24) + 403)
	FNV_MASK_32  = uint32(^uint32(0) >> 1)
)

// Iterates through all of the 8 bytes (64 bits) and tests
//    each bit that is set to 1 in the query's filter against
//    the bit in the comparison's filter.  If the bit is not
//    also 1, you do not have a match.
func TestBytesFromQuery(bf int, qBloom int) bool {
	for i := uint(0); i < NUM_BITS; i++ {
		//a & (1 << idx) == b & (1 << idx)
		if (bf&(1<<i) != (1 << i)) && qBloom&(1<<i) == (1<<i) {
			return false
		}
	}
	return true
}

func computeBloomFilter(s string) int {
	cnt := len(s)
	if cnt <= 0 {
		return 0
	}

	var filter int
	hash := uint64(0)

	for i := 0; i < cnt; i++ {
		c := s[i]

		//first hash
		hash ^= uint64(0xFF & c)
		hash *= FNV_PRIME_64

		//second hash (reduces collissions for bloom)
		hash ^= uint64(0xFF & (c >> 16))
		hash *= FNV_PRIME_64

		//position of bit mod number of bits (8 bytes = 64 bits)
		bitpos := hash % NUM_BITS
		if bitpos < 0 {
			bitpos += NUM_BITS
		}
		filter = filter | (1 << bitpos)
	}

	return filter
}
