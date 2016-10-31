package bitset

import (
	"fmt"
	"testing"
)

func TestBitSetIndex(t *testing.T) {
	index := NewBitSetIndex(128, .75)
	sets := make([]*BitSet, 0, 100)
	for i := 0; i < 100; i++ {
		b := New(100)
		b.Set(uint(i))
		sets = append(sets, b)
		index.AddCount(b)

		val, ok := index.Value(b)
		if val != 1 || !ok {
			t.Error(fmt.Sprintf("BitSet value must be == 1 and is %d", val))
		}
	}

	for i := 2; i < 10; i++ {
		for _, b := range sets {
			index.AddCount(b)

			val, ok := index.Value(b)
			if val != i || !ok {
				t.Error(fmt.Sprintf("BitSet value must be == %d and is %d", i, val))
			}
		}
	}
}

func TestKeys(t *testing.T) {
	index := NewBitSetIndex(128, .75)
	for i := 0; i < 100; i++ {
		b := New(100)
		b.Set(uint(i))
		index.AddCount(b)

		val, ok := index.Value(b)
		if val != 1 || !ok {
			t.Error(fmt.Sprintf("BitSet value must be == 1 and is %d", val))
		}
	}

	keys := index.Keys()

	if len(keys) != 100 {
		t.Error(fmt.Sprintf("BitSet index should have %d keys but have %d", 100, len(keys)))
	}

	for i := 2; i < 10; i++ {
		for _, b := range keys {
			index.AddCount(b)

			val, ok := index.Value(b)
			if val != i || !ok {
				t.Error(fmt.Sprintf("BitSet value must be == %d and is %d", i, val))
			}
		}
	}

}
