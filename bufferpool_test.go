package bpool

import (
	"bytes"
	"testing"
)

func TestBufferPool(t *testing.T) {
	maxSize := 4

	bufPool := NewBufferPool(maxSize)

	// Test Get/Put
	b := bufPool.Get()
	if b.Cap() != 0 {
		t.Errorf("bufferpool width invalid: got %v want %v", b.Cap(), 0)
	}
	bufPool.Put(b)

	// Add some additional buffers beyond the pool size.
	for i := 0; i < maxSize*2; i++ {
		bufPool.Put(bytes.NewBuffer([]byte{}))
	}

	// Close the channel so we can iterate over it.
	close(bufPool.c)

	// Check the size of the pool.
	if len(bufPool.c) != maxSize {
		t.Fatalf("bufferpool size invalid: got %v want %v", len(bufPool.c), maxSize)
	}
}

func TestBufferPoolWidth(t *testing.T) {
	maxSize := 4
	width := 10

	bufPool := NewBufferPoolWidth(maxSize, width)

	// Test Get/Put
	b := bufPool.Get()
	if b.Cap() != width {
		t.Errorf("bufferpool width invalid: got %v want %v", b.Cap(), width)
	}
	bufPool.Put(b)

	// Add some additional buffers beyond the pool size.
	for i := 0; i < maxSize*2; i++ {
		bufPool.Put(bytes.NewBuffer([]byte{}))
	}

	// Close the channel so we can iterate over it.
	close(bufPool.c)

	// Check the size of the pool.
	if len(bufPool.c) != maxSize {
		t.Fatalf("bufferpool size invalid: got %v want %v", len(bufPool.c), maxSize)
	}

}
