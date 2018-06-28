package hashmap

import (
	"testing"
)

const n = 1000

func Test_SetAndGet(t *testing.T) {
	cases := []struct {
		key           Key
		value         interface{}
	}{
		{
			"k1",
			"v",
		},
		{
			"k2",
			12345,
		},
	}

	fn := func(blockSize int, key Key) int {
		return len(key) % blockSize
	}

	h := NewHashMap(16, fn)
	for _, c := range cases {
		err := h.Set(c.key, c.value)
		if err != nil {
			t.Fatalf("Unexpected error. %v", err)
		}

		value, err := h.Get(c.key)
		if c.value != value {
			t.Fatalf("Unexpected value. Expected %v, got %v", c.value, value)
		}
	}
}

func Test_SetAndGetALot(t *testing.T) {
	fn := func(blockSize int, key Key) int {
		return len(key) % blockSize
	}

	h := NewHashMap(16, fn)

	for i := 0; i < n; i++ {
		h.Set(Key(i), i)
	}

	if h.Count() != n {
		t.Fatalf("Unexpected count. Expected %d, got %d", n, h.Count())
	}

	for i := 0; i < n; i++ {
		v, err := h.Get(Key(i))
		if err != nil {
			t.Fatalf("Unexpected err %v", err)
		}

		if v == nil || v.(int) != i {
			t.Fatalf("Unexpected value. Expected %v, got %v", i, v)
		}
	}
}

func Test_Unset(t *testing.T) {
	fn := func(blockSize int, key Key) int {
		return len(key) % blockSize
	}

	h := NewHashMap(16, fn)

	for i := 0; i < 10; i++ {
		h.Set(Key(i), i)
	}

	for i := 1; i < 5; i++ {
		err := h.Unset(Key(i))
		if err != nil {
			t.Fatalf("Unexpected err %v", err)
		}
	}

	for i := 0; i < 10; i++ {
		v, _ := h.Get(Key(i))

		if (v == nil || v.(int) != i) && i >= 5 {
			t.Fatalf("Unexpected value. Expected %v, got %v", i, v)
		}

		if i < 5 && v != nil {
			t.Fatalf("Unexpected value. Expected %v, got %v", nil, v)
		}
	}
}

