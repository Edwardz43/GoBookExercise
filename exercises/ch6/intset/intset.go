package intset

import (
	"bytes"
	"fmt"
)

// An IntSet is a set of small non-negative integers.
// Its zero valuer represents the empty set.
type IntSet struct {
	words []uint64
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// Clear remove all elements from s.
func (s *IntSet) Clear() {
	s.words = s.words[:0]
}

// Remove removes the element x from s.
// return true if remove successfully and false if s does not contain x.
func (s *IntSet) Remove(x int) bool {
	if !s.Has(x) {
		return false
	}
	word, bit := x/64, uint(x%64)
	s.words[word] -= 1 << bit
	return true
}

// Copy return the instance of s.
func (s *IntSet) Copy() *IntSet {
	var s2 IntSet
	s2.words = make([]uint64, len(s.words))
	copy(s2.words, s.words)
	return &s2
}

// Elems return all the elements of s.
func (s *IntSet) Elems() []int {
	if s.Len() <= 0 {
		return nil
	}
	var result []int
	for bit, word := range s.words {
		for i := 0; i < 64; i++ {
			if word&(1<<uint(i)) != 0 {
				result = append(result, 64*bit+int(i))
			}
		}
	}
	return result
}

// Len return the the number of elements.
func (s *IntSet) Len() int {
	sum := 0
	for _, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				sum++
			}
		}
	}
	return sum
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// IntersectWith sets s to the intersection of s and t.
func (s *IntSet) IntersectWith(t *IntSet) {
	for i := range s.words {
		if i < len(t.words) {
			s.words[i] &= t.words[i]
		} else {
			s.words[i] = 0
		}
	}
}

// DifferenceWith set s to the difference of s and t.
func (s *IntSet) DifferenceWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &^= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// SymmetricDifference sets s to the symmetric difference between s and t.
func (s *IntSet) SymmetricDifference(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] ^= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

// TODO

// 6.2
//func (*IntSet).AddAll(...int)
