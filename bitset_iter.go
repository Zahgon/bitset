//go:build go1.23
// +build go1.23

package bitset

import (
	"iter"
)

func (b *BitSet) EachSet() iter.Seq[uint] { _ = "STUB: not implemented"; return nil }
