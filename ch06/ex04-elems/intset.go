package intset

type IntSet struct {
	words []uint64
}

func NewIntSet(vals []int) *IntSet {
	intSet := &IntSet{}

	for _, v := range vals {
		intSet.Add(v)
	}

	return intSet
}

func (s *IntSet) Elems() []int {
	elems := make([]int, 0)
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				elems = append(elems, i*64+j)
			}
		}
	}

	return elems
}

func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}
