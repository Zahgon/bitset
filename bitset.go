/*
Package bitset implements bitsets, a mapping
between non-negative integers and boolean values. It should be more
efficient than map[uint] bool.

It provides methods for setting, clearing, flipping, and testing
individual integers.

But it also provides set intersection, union, difference,
complement, and symmetric operations, as well as tests to
check whether any, all, or no bits are set, and querying a
bitset's current length and number of positive bits.

BitSets are expanded to the size of the largest set bit; the
memory allocation is approximately Max bits, where Max is
the largest set bit. BitSets are never shrunk. On creation,
a hint can be given for the number of bits that will be used.

Many of the methods, including Set,Clear, and Flip, return
a BitSet pointer, which allows for chaining.

Example use:

	import "bitset"
	var b BitSet
	b.Set(10).Set(11)
	if b.Test(1000) {
		b.Clear(1000)
	}
	if B.Intersection(bitset.New(100).Set(10)).Count() > 1 {
		fmt.Println("Intersection works.")
	}

As an alternative to BitSets, one should check out the 'big' package,
which provides a (less set-theoretical) view of bitsets.
*/
package bitset

import (
	"encoding/base64"
	"encoding/binary"
	"io"
)

// the wordSize of a bit set
const wordSize = 64

// the wordSize of a bit set in bytes
const wordBytes = wordSize / 8

// wordMask is wordSize-1, used for bit indexing in a word
const wordMask = wordSize - 1

// log2WordSize is lg(wordSize)
const log2WordSize = 6

// allBits has every bit set
const allBits uint64 = 0xffffffffffffffff

// default binary BigEndian
var binaryOrder binary.ByteOrder = binary.BigEndian

// default json encoding base64.URLEncoding
var base64Encoding = base64.URLEncoding

// Base64StdEncoding Marshal/Unmarshal BitSet with base64.StdEncoding(Default: base64.URLEncoding)
func Base64StdEncoding() { _ = "STUB: not implemented"; return }

// LittleEndian sets Marshal/Unmarshal Binary as Little Endian (Default: binary.BigEndian)
func LittleEndian() { _ = "STUB: not implemented"; return }

// BigEndian sets Marshal/Unmarshal Binary as Big Endian (Default: binary.BigEndian)
func BigEndian() { _ = "STUB: not implemented"; return }

// BinaryOrder returns the current binary order, see also LittleEndian()
// and BigEndian() to change the order.
func BinaryOrder() binary.ByteOrder {
	_ = "STUB: not implemented"

	// A BitSet is a set of bits. The zero value of a BitSet is an empty set of length 0.
	return *new(binary.ByteOrder)
}

type BitSet struct {
	length uint
	set    []uint64
}

// Error is used to distinguish errors (panics) generated in this package.
type Error string

// safeSet will fixup b.set to be non-nil and return the field value
func (b *BitSet) safeSet() []uint64 { _ = "STUB: not implemented"; return nil }

// SetBitsetFrom fills the bitset with an array of integers without creating a new BitSet instance
func (b *BitSet) SetBitsetFrom(buf []uint64) { _ = "STUB: not implemented"; return }

// From is a constructor used to create a BitSet from an array of words
func From(buf []uint64) *BitSet { _ = "STUB: not implemented"; return nil }

// FromWithLength constructs from an array of words and length in bits.
// This function is for advanced users, most users should prefer
// the From function.
// As a user of FromWithLength, you are responsible for ensuring
// that the length is correct: your slice should have length at
// least (length+63)/64 in 64-bit words.
func FromWithLength(length uint, set []uint64) *BitSet { _ = "STUB: not implemented"; return nil }

// Bytes returns the bitset as array of 64-bit words, giving direct access to the internal representation.
// It is not a copy, so changes to the returned slice will affect the bitset.
// It is meant for advanced users.
//
// Deprecated: Bytes is deprecated. Use [BitSet.Words] instead.
func (b *BitSet) Bytes() []uint64 {
	_ = "STUB: not implemented"

	// Words returns the bitset as array of 64-bit words, giving direct access to the internal representation.
	// It is not a copy, so changes to the returned slice will affect the bitset.
	// It is meant for advanced users.
	return nil
}

func (b *BitSet) Words() []uint64 {
	_ = "STUB: not implemented"

	// wordsNeeded calculates the number of words needed for i bits
	return nil
}

func wordsNeeded(i uint) int { _ = "STUB: not implemented"; return 0 }

// wordsNeededUnbound calculates the number of words needed for i bits, possibly exceeding the capacity.
// This function is useful if you know that the capacity cannot be exceeded (e.g., you have an existing BitSet).
func wordsNeededUnbound(i uint) int { _ = "STUB: not implemented"; return 0 }

// wordsIndex calculates the index of words in a `uint64`
func wordsIndex(i uint) uint { _ = "STUB: not implemented"; return 0 }

// New creates a new BitSet with a hint that length bits will be required.
// The memory usage is at least length/8 bytes.
// In case of allocation failure, the function will return a BitSet with zero
// capacity.
func New(length uint) (bset *BitSet) { _ = "STUB: not implemented"; return nil }

// MustNew creates a new BitSet with the given length bits.
// It panics if length exceeds the possible capacity or by a lack of memory.
func MustNew(length uint) (bset *BitSet) { _ = "STUB: not implemented"; return nil }

// may panic on lack of memory

// Cap returns the total possible capacity, or number of bits
// that can be stored in the BitSet theoretically. Under 32-bit system,
// it is 4294967295 and under 64-bit system, it is 18446744073709551615.
// Note that this is further limited by the maximum allocation size in Go,
// and your available memory, as any Go data structure.
func Cap() uint {
	_ = "STUB: not implemented"

	// Len returns the number of bits in the BitSet.
	// Note that it differ from Count function.
	return 0
}

func (b *BitSet) Len() uint {
	_ = "STUB: not implemented"

	// extendSet adds additional words to incorporate new bits if needed
	return 0
}

func (b *BitSet) extendSet(i uint) { _ = "STUB: not implemented"; return }

// fast resize

// increase capacity 2x

// Test whether bit i is set.
func (b *BitSet) Test(i uint) bool { _ = "STUB: not implemented"; return false }

// GetWord64AtBit retrieves bits i through i+63 as a single uint64 value
func (b *BitSet) GetWord64AtBit(i uint) uint64 { _ = "STUB: not implemented"; return 0 }

// The word that the index falls within, shifted so the index is at bit 0

// The next word, masked to only include the necessary bits and shifted to cover the
// top of the word

// Set bit i to 1, the capacity of the bitset is automatically
// increased accordingly.
// Warning: using a very large value for 'i'
// may lead to a memory shortage and a panic: the caller is responsible
// for providing sensible parameters in line with their memory capacity.
// The memory usage is at least slightly over i/8 bytes.
func (b *BitSet) Set(i uint) *BitSet {
	_ = "STUB: not implemented"
	// if we need more bits, make 'em
	return nil
}

// Clear bit i to 0. This never cause a memory allocation. It is always safe.
func (b *BitSet) Clear(i uint) *BitSet { _ = "STUB: not implemented"; return nil }

// SetTo sets bit i to value.
// Warning: using a very large value for 'i'
// may lead to a memory shortage and a panic: the caller is responsible
// for providing sensible parameters in line with their memory capacity.
func (b *BitSet) SetTo(i uint, value bool) *BitSet { _ = "STUB: not implemented"; return nil }

// Flip bit at i.
// Warning: using a very large value for 'i'
// may lead to a memory shortage and a panic: the caller is responsible
// for providing sensible parameters in line with their memory capacity.
func (b *BitSet) Flip(i uint) *BitSet { _ = "STUB: not implemented"; return nil }

// FlipRange bit in [start, end).
// Warning: using a very large value for 'end'
// may lead to a memory shortage and a panic: the caller is responsible
// for providing sensible parameters in line with their memory capacity.
func (b *BitSet) FlipRange(start, end uint) *BitSet { _ = "STUB: not implemented"; return nil }

// if we need more bits, make 'em

// b.set[startWord] ^= ^(^uint64(0) << wordsIndex(start))
//  e.g:
//  start = 71,
//  startWord = 1
//  wordsIndex(start) = 71 % 64 = 7
//   (^uint64(0) << 7) = 0b111111....11110000000
//
//  mask = ^(^uint64(0) << 7) = 0b000000....00001111111
//
// flips the first 7 bits in b.set[1] and
// in the range loop, the b.set[1] gets again flipped
// so the two expressions flip results in a flip
// in b.set[1] from [7,63]
//
// handle startWord special, get's reflipped in range loop

// handle endWord special
//  e.g.
// end = 135
//  endWord = 2
//
//  wordsIndex(-7) = 57
//  see the golang spec:
//   "For unsigned integer values, the operations +, -, *, and << are computed
//   modulo 2n, where n is the bit width of the unsigned integer's type."
//
//   mask = ^uint64(0) >> 57 = 0b00000....0001111111
//
// flips in b.set[2] from [0,7]
//
// is end at word boundary?

// Shrink shrinks BitSet so that the provided value is the last possible
// set value. It clears all bits > the provided index and reduces the size
// and length of the set.
//
// Note that the parameter value is not the new length in bits: it is the
// maximal value that can be stored in the bitset after the function call.
// The new length in bits is the parameter value + 1. Thus it is not possible
// to use this function to set the length to 0, the minimal value of the length
// after this function call is 1.
//
// A new slice is allocated to store the new bits, so you may see an increase in
// memory usage until the GC runs. Normally this should not be a problem, but if you
// have an extremely large BitSet its important to understand that the old BitSet will
// remain in memory until the GC frees it.
// If you are memory constrained, this function may cause a panic.
func (b *BitSet) Shrink(lastbitindex uint) *BitSet { _ = "STUB: not implemented"; return nil }

// Compact shrinks BitSet to so that we preserve all set bits, while minimizing
// memory usage. Compact calls Shrink.
// A new slice is allocated to store the new bits, so you may see an increase in
// memory usage until the GC runs. Normally this should not be a problem, but if you
// have an extremely large BitSet its important to understand that the old BitSet will
// remain in memory until the GC frees it.
// If you are memory constrained, this function may cause a panic.
func (b *BitSet) Compact() *BitSet { _ = "STUB: not implemented"; return nil }

// nothing to do

// We preserve one word

// InsertAt takes an index which indicates where a bit should be
// inserted. Then it shifts all the bits in the set to the left by 1, starting
// from the given index position, and sets the index position to 0.
//
// Depending on the size of your BitSet, and where you are inserting the new entry,
// this method could be extremely slow and in some cases might cause the entire BitSet
// to be recopied.
func (b *BitSet) InsertAt(idx uint) *BitSet { _ = "STUB: not implemented"; return nil }

// if length of set is a multiple of wordSize we need to allocate more space first

// all elements above the position where we want to insert can simply by shifted

// we take the most significant bit of the previous element and set it as
// the least significant bit of the current element

// generate a mask to extract the data that we need to shift left
// within the element where we insert a bit

// extract that data that we'll shift

// set the positions of the data mask to 0 in the element where we insert

// shift data mask to the left and insert its data to the slice element

// add 1 to length of BitSet

// String creates a string representation of the BitSet. It is only intended for
// human-readable output and not for serialization.
func (b *BitSet) String() string {
	_ = "STUB: not implemented"
	// follows code from https://github.com/RoaringBitmap/roaring
	return ""
}

// to avoid exhausting the memory

// DeleteAt deletes the bit at the given index position from
// within the bitset
// All the bits residing on the left of the deleted bit get
// shifted right by 1
// The running time of this operation may potentially be
// relatively slow, O(length)
func (b *BitSet) DeleteAt(i uint) *BitSet {
	_ = "STUB: not implemented"
	// the index of the slice element where we'll delete a bit
	return nil
}

// generate a mask for the data that needs to be shifted right
// within that slice element that gets modified

// extract the data that we'll shift right from the slice element

// set the masked area to 0 while leaving the rest as it is

// shift the previously extracted data to the right and then
// set it in the previously masked area

// loop over all the consecutive slice elements to copy each
// lowest bit into the highest position of the previous element,
// then shift the entire content to the right by 1

// AppendTo appends all set bits to buf and returns the (maybe extended) buf.
// In case of allocation failure, the function will panic.
//
// See also [BitSet.AsSlice] and [BitSet.NextSetMany].
func (b *BitSet) AppendTo(buf []uint) []uint {
	_ = "STUB: not implemented"
	// In theory, we could overflow uint, but in practice, we will not.
	return nil
}

// In theory idx<<log2WordSize could overflow, but it will not overflow
// in practice.

// clear the rightmost set bit

// AsSlice returns all set bits as slice.
// It panics if the capacity of buf is < b.Count()
//
// See also [BitSet.AppendTo] and [BitSet.NextSetMany].
func (b *BitSet) AsSlice(buf []uint) []uint {
	_ = "STUB: not implemented"
	// len = cap
	return nil
}

// panics if capacity of buf is exceeded.
// In theory idx<<log2WordSize could overflow, but it will not overflow
// in practice.

// clear the rightmost set bit

// NextSet returns the next bit set from the specified index,
// including possibly the current index
// along with an error code (true = valid, false = no set bit found)
// for i,e := v.NextSet(0); e; i,e = v.NextSet(i + 1) {...}
//
// Users concerned with performance may want to use NextSetMany to
// retrieve several values at once.
func (b *BitSet) NextSet(i uint) (uint, bool) { _ = "STUB: not implemented"; return 0, false }

// process first (partial) word

// process the following full words until next bit is set
// x < len(b.set), no out-of-bounds panic in following slice expression

// NextSetMany returns many next bit sets from the specified index,
// including possibly the current index and up to cap(buffer).
// If the returned slice has len zero, then no more set bits were found
//
//	buffer := make([]uint, 256) // this should be reused
//	j := uint(0)
//	j, buffer = bitmap.NextSetMany(j, buffer)
//	for ; len(buffer) > 0; j, buffer = bitmap.NextSetMany(j,buffer) {
//	 for k := range buffer {
//	  do something with buffer[k]
//	 }
//	 j += 1
//	}
//
// It is possible to retrieve all set bits as follow:
//
//	indices := make([]uint, bitmap.Count())
//	bitmap.NextSetMany(0, indices)
//
// It is also possible to retrieve all set bits with [BitSet.AppendTo]
// or [BitSet.AsSlice].
//
// However if Count() is large, it might be preferable to
// use several calls to NextSetMany for memory reasons.
func (b *BitSet) NextSetMany(i uint, buffer []uint) (uint, []uint) {
	_ = "STUB: not implemented"
	// In theory, we could overflow uint, but in practice, we will not.
	return 0, nil
}

// process first (partial) word

// clear the rightmost set bit

// process the following full words
// x < len(b.set), no out-of-bounds panic in following slice expression

// clear the rightmost set bit

// NextClear returns the next clear bit from the specified index,
// including possibly the current index
// along with an error code (true = valid, false = no bit found i.e. all bits are set)
func (b *BitSet) NextClear(i uint) (uint, bool) { _ = "STUB: not implemented"; return 0, false }

// process first (maybe partial) word

// process the following full words until next bit is cleared
// x < len(b.set), no out-of-bounds panic in following slice expression

// PreviousSet returns the previous set bit from the specified index,
// including possibly the current index
// along with an error code (true = valid, false = no bit found i.e. all bits are clear)
func (b *BitSet) PreviousSet(i uint) (uint, bool) { _ = "STUB: not implemented"; return 0, false }

// Clear the bits above the index

// PreviousClear returns the previous clear bit from the specified index,
// including possibly the current index
// along with an error code (true = valid, false = no clear bit found i.e. all bits are set)
func (b *BitSet) PreviousClear(i uint) (uint, bool) { _ = "STUB: not implemented"; return 0, false }

// Flip all bits and find the highest one bit

// Clear the bits above the index

// ClearAll clears the entire BitSet.
// It does not free the memory.
func (b *BitSet) ClearAll() *BitSet { _ = "STUB: not implemented"; return nil }

// SetAll sets the entire BitSet
func (b *BitSet) SetAll() *BitSet { _ = "STUB: not implemented"; return nil }

// wordCount returns the number of words used in a bit set
func (b *BitSet) wordCount() int { _ = "STUB: not implemented"; return 0 }

// Clone this BitSet, returning a new BitSet that has the same bits set.
// In case of allocation failure, the function will return an empty BitSet.
func (b *BitSet) Clone() *BitSet { _ = "STUB: not implemented"; return nil }

// Clone should not modify current object

// Copy into a destination BitSet using the Go array copy semantics:
// the number of bits copied is the minimum of the number of bits in the current
// BitSet (Len()) and the destination Bitset.
// We return the number of bits copied in the destination BitSet.
func (b *BitSet) Copy(c *BitSet) (count uint) { _ = "STUB: not implemented"; return 0 }

// Copy should not modify current object

// Cleaning the last word is needed to keep the invariant that other functions, such as Count, require
// that any bits in the last word that would exceed the length of the bitmask are set to 0.

// CopyFull copies into a destination BitSet such that the destination is
// identical to the source after the operation, allocating memory if necessary.
func (b *BitSet) CopyFull(c *BitSet) { _ = "STUB: not implemented"; return }

// Count (number of set bits).
// Also known as "popcount" or "population count".
func (b *BitSet) Count() uint { _ = "STUB: not implemented"; return 0 }

// Equal tests the equivalence of two BitSets.
// False if they are of different sizes, otherwise true
// only if all the same bits are set
func (b *BitSet) Equal(c *BitSet) bool { _ = "STUB: not implemented"; return false }

// if they have both length == 0, then could have nil set

// bounds check elimination

func panicIfNull(b *BitSet) { _ = "STUB: not implemented"; return }

// Difference of base set and other set
// This is the BitSet equivalent of &^ (and not)
func (b *BitSet) Difference(compare *BitSet) (result *BitSet) {
	_ = "STUB: not implemented"
	return nil
}

// clone b (in case b is bigger than compare)

// DifferenceCardinality computes the cardinality of the difference
func (b *BitSet) DifferenceCardinality(compare *BitSet) uint { _ = "STUB: not implemented"; return 0 }

// InPlaceDifference computes the difference of base set and other set
// This is the BitSet equivalent of &^ (and not)
func (b *BitSet) InPlaceDifference(compare *BitSet) { _ = "STUB: not implemented"; return }

// bounds check elimination

// Convenience function: return two bitsets ordered by
// increasing length. Note: neither can be nil
func sortByLength(a *BitSet, b *BitSet) (ap *BitSet, bp *BitSet) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Intersection of base set and other set
// This is the BitSet equivalent of & (and)
// In case of allocation failure, the function will return an empty BitSet.
func (b *BitSet) Intersection(compare *BitSet) (result *BitSet) {
	_ = "STUB: not implemented"
	return nil
}

// IntersectionCardinality computes the cardinality of the intersection
func (b *BitSet) IntersectionCardinality(compare *BitSet) uint { _ = "STUB: not implemented"; return 0 }

// InPlaceIntersection destructively computes the intersection of
// base set and the compare set.
// This is the BitSet equivalent of & (and)
func (b *BitSet) InPlaceIntersection(compare *BitSet) { _ = "STUB: not implemented"; return }

// bounds check elimination

// Union of base set and other set
// This is the BitSet equivalent of | (or)
func (b *BitSet) Union(compare *BitSet) (result *BitSet) { _ = "STUB: not implemented"; return nil }

// UnionCardinality computes the cardinality of the uniton of the base set
// and the compare set.
func (b *BitSet) UnionCardinality(compare *BitSet) uint { _ = "STUB: not implemented"; return 0 }

// InPlaceUnion creates the destructive union of base set and compare set.
// This is the BitSet equivalent of | (or).
func (b *BitSet) InPlaceUnion(compare *BitSet) { _ = "STUB: not implemented"; return }

// bounds check elimination

// SymmetricDifference of base set and other set
// This is the BitSet equivalent of ^ (xor)
func (b *BitSet) SymmetricDifference(compare *BitSet) (result *BitSet) {
	_ = "STUB: not implemented"
	return nil
}

// compare is bigger, so clone it

// SymmetricDifferenceCardinality computes the cardinality of the symmetric difference
func (b *BitSet) SymmetricDifferenceCardinality(compare *BitSet) uint {
	_ = "STUB: not implemented"
	return 0
}

// InPlaceSymmetricDifference creates the destructive SymmetricDifference of base set and other set
// This is the BitSet equivalent of ^ (xor)
func (b *BitSet) InPlaceSymmetricDifference(compare *BitSet) { _ = "STUB: not implemented"; return }

// bounds check elimination

// Is the length an exact multiple of word sizes?
func (b *BitSet) isLenExactMultiple() bool { _ = "STUB: not implemented"; return false }

// Clean last word by setting unused bits to 0
func (b *BitSet) cleanLastWord() { _ = "STUB: not implemented"; return }

// Complement computes the (local) complement of a bitset (up to length bits)
// In case of allocation failure, the function will return an empty BitSet.
func (b *BitSet) Complement() (result *BitSet) { _ = "STUB: not implemented"; return nil }

// All returns true if all bits are set, false otherwise. Returns true for
// empty sets.
func (b *BitSet) All() bool { _ = "STUB: not implemented"; return false }

// None returns true if no bit is set, false otherwise. Returns true for
// empty sets.
func (b *BitSet) None() bool { _ = "STUB: not implemented"; return false }

// Any returns true if any bit is set, false otherwise
func (b *BitSet) Any() bool { _ = "STUB: not implemented"; return false }

// IsSuperSet returns true if this is a superset of the other set
func (b *BitSet) IsSuperSet(other *BitSet) bool { _ = "STUB: not implemented"; return false }

// IsStrictSuperSet returns true if this is a strict superset of the other set
func (b *BitSet) IsStrictSuperSet(other *BitSet) bool { _ = "STUB: not implemented"; return false }

// DumpAsBits dumps a bit set as a string of bits. Following the usual convention in Go,
// the least significant bits are printed last (index 0 is at the end of the string).
// This is useful for debugging and testing. It is not suitable for serialization.
func (b *BitSet) DumpAsBits() string { _ = "STUB: not implemented"; return "" }

// BinaryStorageSize returns the binary storage requirements (see WriteTo) in bytes.
func (b *BitSet) BinaryStorageSize() int { _ = "STUB: not implemented"; return 0 }

func readUint64Array(reader io.Reader, data []uint64) error { _ = "STUB: not implemented"; return nil }

func writeUint64Array(writer io.Writer, data []uint64) error { _ = "STUB: not implemented"; return nil }

// WriteTo writes a BitSet to a stream. The format is:
// 1. uint64 length
// 2. []uint64 set
// The length is the number of bits in the BitSet.
//
// The set is a slice of uint64s containing between length and length + 63 bits.
// It is interpreted as a big-endian array of uint64s by default (see BinaryOrder())
// meaning that the first 8 bits are stored at byte index 7, the next 8 bits are stored
// at byte index 6... the bits 64 to 71 are stored at byte index 8, etc.
// If you change the binary order, you need to do so for both reading and writing.
// We recommend using the default binary order.
//
// Upon success, the number of bytes written is returned.
//
// Performance: if this function is used to write to a disk or network
// connection, it might be beneficial to wrap the stream in a bufio.Writer.
// E.g.,
//
//	      f, err := os.Create("myfile")
//		       w := bufio.NewWriter(f)
func (b *BitSet) WriteTo(stream io.Writer) (int64, error) {
	_ = "STUB: not implemented"
	return 0,

		// Write length
		nil
}

// Upon failure, we do not guarantee that we
// return the number of bytes written.

// Upon failure, we do not guarantee that we
// return the number of bytes written.

// ReadFrom reads a BitSet from a stream written using WriteTo
// The format is:
// 1. uint64 length
// 2. []uint64 set
// See WriteTo for details.
// Upon success, the number of bytes read is returned.
// If the current BitSet is not large enough to hold the data,
// it is extended. In case of error, the BitSet is either
// left unchanged or made empty if the error occurs too late
// to preserve the content.
//
// Performance: if this function is used to read from a disk or network
// connection, it might be beneficial to wrap the stream in a bufio.Reader.
// E.g.,
//
//	f, err := os.Open("myfile")
//	r := bufio.NewReader(f)
func (b *BitSet) ReadFrom(stream io.Reader) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// We do not want to leave the BitSet partially filled as
// it is error prone.

// MarshalBinary encodes a BitSet into a binary form and returns the result.
// Please see WriteTo for details.
func (b *BitSet) MarshalBinary() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalBinary decodes the binary form generated by MarshalBinary.
// Please see WriteTo for details.
func (b *BitSet) UnmarshalBinary(data []byte) error { _ = "STUB: not implemented"; return nil }

// MarshalJSON marshals a BitSet as a JSON structure
func (b BitSet) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// URLEncode all bytes

// UnmarshalJSON unmarshals a BitSet from JSON created using MarshalJSON
func (b *BitSet) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	// Unmarshal as string
	return nil
}

// URLDecode string

// Rank returns the number of set bits up to and including the index
// that are set in the bitset.
// See https://en.wikipedia.org/wiki/Ranking#Ranking_in_statistics
func (b *BitSet) Rank(index uint) (rank uint) {
	_ = "STUB: not implemented"
	// Rank is up to and including
	return 0
}

// needed more than once

// TODO: built-in min requires go1.21 or later
// idx := min(int(index>>6), len(b.set))

// sum up the popcounts until idx ...
// TODO: cannot range over idx (...): requires go1.22 or later
// for j := range idx {

// ... plus partial word at idx,
// make Rank inlineable and faster in the end
// don't test index&63 != 0, just add, less branching

// Select returns the index of the jth set bit, where j is the argument.
// The caller is responsible to ensure that 0 <= j < Count(): when j is
// out of range, the function returns the length of the bitset (b.length).
//
// Note that this function differs in convention from the Rank function which
// returns 1 when ranking the smallest value. We follow the conventional
// textbook definition of Select and Rank.
func (b *BitSet) Select(index uint) uint { _ = "STUB: not implemented"; return 0 }

// top detects the top bit set
func (b *BitSet) top() (uint, bool) { _ = "STUB: not implemented"; return 0, false }

// ShiftLeft shifts the bitset like << operation would do.
//
// Left shift may require bitset size extension. We try to avoid the
// unnecessary memory operations by detecting the leftmost set bit.
// The function will panic if shift causes excess of capacity.
func (b *BitSet) ShiftLeft(bits uint) { _ = "STUB: not implemented"; return }

// capacity check

// destination set

// not using extendSet() to avoid unneeded data copying

// happy case: just add pages

// zeroing extra pages

// ShiftRight shifts the bitset like >> operation would do.
func (b *BitSet) ShiftRight(bits uint) { _ = "STUB: not implemented"; return }

// happy case: just clear pages

// OnesBetween returns the number of set bits in the range [from, to).
// The range is inclusive of 'from' and exclusive of 'to'.
// Returns 0 if from >= to.
func (b *BitSet) OnesBetween(from, to uint) uint { _ = "STUB: not implemented"; return 0 }

// Calculate indices and masks for the starting and ending words
// Divide by wordSize

// Mod wordSize

// Case 1: Bits lie within a single word

// Create mask for bits between from and to

// Case 2: Bits span multiple words
// 2a: Count bits in first word (from startOffset to end of word)
// Mask for bits >= startOffset

// 2b: Count all bits in complete words between start and end

// 2c: Count bits in last word (from start of word to endOffset)

// Mask for bits < endOffset

// Extract extracts bits according to a mask and returns the result
// in a new BitSet. See ExtractTo for details.
func (b *BitSet) Extract(mask *BitSet) *BitSet { _ = "STUB: not implemented"; return nil }

// ExtractTo copies bits from the BitSet using positions specified in mask
// into a compacted form in dst. The number of set bits in mask determines
// the number of bits that will be extracted.
//
// For example, if mask has bits set at positions 1,4,5, then ExtractTo will
// take bits at those positions from the source BitSet and pack them into
// consecutive positions 0,1,2 in the destination BitSet.
func (b *BitSet) ExtractTo(mask *BitSet, dst *BitSet) { _ = "STUB: not implemented"; return }

// Ensure destination has enough space for extracted bits

// Process each word

// Skip words with no bits to extract

// Extract and compact bits according to mask

// Calculate destination position

// Write extracted bits, handling word boundary crossing

// Deposit creates a new BitSet and deposits bits according to a mask.
// See DepositTo for details.
func (b *BitSet) Deposit(mask *BitSet) *BitSet { _ = "STUB: not implemented"; return nil }

// DepositTo spreads bits from a compacted form in the BitSet into positions
// specified by mask in dst. This is the inverse operation of Extract.
//
// For example, if mask has bits set at positions 1,4,5, then DepositTo will
// take consecutive bits 0,1,2 from the source BitSet and place them into
// positions 1,4,5 in the destination BitSet.
func (b *BitSet) DepositTo(mask *BitSet, dst *BitSet) { _ = "STUB: not implemented"; return }

// Process each word

// Skip words with no bits to deposit

// Calculate source word index

// No more source bits available

// Get source bits, handling word boundary crossing

// Combine bits from current and next word

// Deposit bits according to mask

//go:generate go run cmd/pextgen/main.go -pkg=bitset

func pext(w, m uint64) (result uint64) {
	_ = "STUB: not implemented"

	// Process byte by byte
	return 0
}

// i * 8 using bit shift

func pdep(w, m uint64) (result uint64) {
	_ = "STUB: not implemented"

	// Process byte by byte
	return 0
}

// i * 8 using bit shift

// Get the bits we'll deposit from the source

// Deposit them according to the mask for this byte

// Add to result
