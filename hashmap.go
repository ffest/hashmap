package hashmap

import (
	"errors"
)

type Key string

type HashMaper interface {
	Set(key Key, value interface{}) error
	Get(key Key) (value interface{}, err error)
	Unset(key Key) error
	Count() int
}

func NewHashMap(blockSize int, fn func(blockSize int, key Key) int) *HashMap {
	return &HashMap{
		BlockSize: blockSize,
		Fn:        fn,
		buckets:   make([]*mapEntry, blockSize),
	}
}

type HashMap struct {
	BlockSize int
	Fn        func(blockSize int, key Key) int
	buckets   []*mapEntry
	size      int
}

// Put an item into the map
func (h *HashMap) Set(key Key, value interface{}) error {
	bucketNum := h.Fn(h.BlockSize, key)
	newEntry := &mapEntry{key, value, nil}

	mapEntry := h.buckets[bucketNum]
	if mapEntry == nil {
		h.buckets[bucketNum] = newEntry
		h.size++
	} else {
		for mapEntry != nil {
			if (*mapEntry).key == newEntry.key {
				(*mapEntry).value = newEntry.value
				return nil
			}
			if (*mapEntry).next == nil {
				(*mapEntry).next = newEntry
				h.size++
				return nil
			}
			mapEntry = (*mapEntry).next
		}
	}

	return nil
}

func (h *HashMap) Get(key Key) (value interface{}, err error) {
	bucketNum := h.Fn(h.BlockSize, key)

	mapEntry := h.buckets[bucketNum]

	for mapEntry != nil {
		if (*mapEntry).key == key {
			return (*mapEntry).value, nil
		}
		mapEntry = (*mapEntry).next
	}

	return nil, errors.New("no data")
}

func (h *HashMap) Unset(key Key) error {
	bucketNum := h.Fn(h.BlockSize, key)

	mapEntry := h.buckets[bucketNum]
	var prev = mapEntry

	for mapEntry != nil {
		if mapEntry.key == key {
			prev.next = mapEntry.next
			h.size--
		}

		prev = mapEntry
		mapEntry = mapEntry.next
	}

	return nil
}

func (h *HashMap) Count() int {
	return h.size
}

type mapEntry struct {
	key   Key
	value interface{}
	next  *mapEntry
}
