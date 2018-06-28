package hashmap

import (
	"testing"
)

var fn = func(blockSize int, key Key) int {
	return len(key) % blockSize
}

func BenchmarkSet16(b *testing.B) {
	h := NewHashMap(16, fn)
	for i := 0; i < b.N; i++ {
		h.Set(Key(i), i)
	}
}

func BenchmarkSet64(b *testing.B) {
	h := NewHashMap(64, fn)
	for i := 0; i < b.N; i++ {
		h.Set(Key(i), i)
	}
}

func BenchmarkSet128(b *testing.B) {
	h := NewHashMap(128, fn)
	for i := 0; i < b.N; i++ {
		h.Set(Key(i), i)
	}
}

func BenchmarkSet1024(b *testing.B) {
	h := NewHashMap(1024, fn)
	for i := 0; i < b.N; i++ {
		h.Set(Key(i), i)
	}
}

func BenchmarkGet16(b *testing.B) {
	h := NewHashMap(16, fn)
	for i := 0; i < b.N; i++ {
		_, _ = h.Get(Key(i))
	}
}


func BenchmarkGet64(b *testing.B) {
	h := NewHashMap(64, fn)
	for i := 0; i < b.N; i++ {
		_, _ = h.Get(Key(i))
	}
}


func BenchmarkGet128(b *testing.B) {
	h := NewHashMap(128, fn)
	for i := 0; i < b.N; i++ {
		_, _ = h.Get(Key(i))
	}
}

func BenchmarkGet1024(b *testing.B) {
	h := NewHashMap(1024, fn)
	for i := 0; i < b.N; i++ {
		_, _ = h.Get(Key(i))
	}
}

func BenchmarkUnset16(b *testing.B) {
	h := NewHashMap(16, fn)
	for i := 0; i < b.N; i++ {
		h.Unset(Key(i))
	}
}

func BenchmarkUnset64(b *testing.B) {
	h := NewHashMap(64, fn)
	for i := 0; i < b.N; i++ {
		h.Unset(Key(i))
	}
}


func BenchmarkUnset128(b *testing.B) {
	h := NewHashMap(128, fn)
	for i := 0; i < b.N; i++ {
		h.Unset(Key(i))
	}
}
func BenchmarkUnset1024(b *testing.B) {
	h := NewHashMap(1024, fn)
	for i := 0; i < b.N; i++ {
		h.Unset(Key(i))
	}
}

func BenchmarkSetMap(b *testing.B) {
	m := make(map[int]interface{})
	for i := 0; i < b.N; i++ {
		m[i] = i
	}
}

func BenchmarkGetMap(b *testing.B) {
	m := make(map[int]interface{})
	for i := 0; i < b.N; i++ {
		_ = m[i]
	}
}

func BenchmarkUnsetMap(b *testing.B) {
	m := make(map[int]interface{})
	for i := 0; i < b.N; i++ {
		delete(m, i)
	}
}