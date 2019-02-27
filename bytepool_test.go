package bpool

import "testing"

func TestBytePool(t *testing.T) {
	var size int = 4
	var width int = 10

	bufPool := NewBytePool(size, width)

	// Check the width
	if bufPool.Width() != width {
		t.Fatalf("bytepool width invalid: got %v want %v", bufPool.Width(), width)
	}

	// Check that retrieved buffer are of the expected width
	b := bufPool.Get()
	if len(b) != width {
		t.Fatalf("bytepool length invalid: got %v want %v", len(b), width)
	}

	// Try putting some invalid buffers into pool
	bufPool.Put(make([]byte, width-1))
	bufPool.Put(make([]byte, width)[2:])
	if len(bufPool.c) > 0 {
		t.Fatal("bytepool should have rejected invalid packets")
	}

	// Try putting a short slice into pool
	bufPool.Put(make([]byte, width)[:2])
	if len(bufPool.c) != 1 {
		t.Fatal("bytepool should have accepted short slice with sufficient capacity")
	}

	b = bufPool.Get()
	if len(b) != width {
		t.Fatalf("bytepool length invalid: got %v want %v", len(b), width)
	}

	// Fill the pool beyond the capped pool size.
	for i := 0; i < size*2; i++ {
		bufPool.Put(make([]byte, bufPool.w))
	}

	// Close the channel so we can iterate over it.
	close(bufPool.c)

	// Check the size of the pool.
	if len(bufPool.c) != size {
		t.Fatalf("bytepool size invalid: got %v want %v", len(bufPool.c), size)
	}

}
