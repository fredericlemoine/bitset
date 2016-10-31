package bitset

import (
	"fmt"
	"testing"
)

type IntValue struct {
	val int
}

func addCount(b *BitSet, bi *BitSetIndex) {
	val, ok := bi.Value(b)
	if !ok {
		bi.PutValue(b, &IntValue{1})
	} else {
		(val.(*IntValue)).val += 1
	}
}

func TestBitSetIndex(t *testing.T) {
	index := NewBitSetIndex(128, .75)
	sets := make([]*BitSet, 0, 100)
	for i := 0; i < 100; i++ {
		b := New(100)
		b.Set(uint(i))
		sets = append(sets, b)
		addCount(b, index)
		val, ok := index.Value(b)
		if val.(*IntValue).val != 1 || !ok {
			t.Error(fmt.Sprintf("BitSet value must be == 1 and is %d", val))
		}
	}

	for i := 2; i < 10; i++ {
		for _, b := range sets {
			addCount(b, index)
			val, ok := index.Value(b)
			if val.(*IntValue).val != i || !ok {
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

		addCount(b, index)

		val, ok := index.Value(b)
		if val.(*IntValue).val != 1 || !ok {
			t.Error(fmt.Sprintf("BitSet value must be == 1 and is %d", val))
		}
	}

	keys := index.Keys()

	if len(keys) != 100 {
		t.Error(fmt.Sprintf("BitSet index should have %d keys but have %d", 100, len(keys)))
	}

	for i := 2; i < 10; i++ {
		for _, b := range keys {
			addCount(b, index)

			val, ok := index.Value(b)
			if val.(*IntValue).val != i || !ok {
				t.Error(fmt.Sprintf("BitSet value must be == %d and is %d", i, val))
			}
		}
	}

}
