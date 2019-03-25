package bpool

import (
	"bytes"
)

// BufferPool implements a pool of bytes.Buffers in the form of a bounded
// channel.
type BufferPool struct {
	c chan *bytes.Buffer
	w int
}

// NewBufferPool creates a new BufferPool bounded to the given maxSize
// with a buffer width initialized to 0.
func NewBufferPool(maxSize int) (bp *BufferPool) {
	return NewBufferPoolWidth(maxSize, 0)
}

// NewBufferPoolWidth creates a new BufferPool bounded to the given maxSize
// with new buffer's backend byte slices sized based on width.
func NewBufferPoolWidth(maxSize, width int) (bp *BufferPool) {
	return &BufferPool{
		c: make(chan *bytes.Buffer, maxSize),
		w: width,
	}
}

// Get gets a Buffer from the BufferPool, or creates a new one if none are
// available in the pool.
func (bp *BufferPool) Get() (b *bytes.Buffer) {
	select {
	case b = <-bp.c:
	// reuse existing buffer
	default:
		// create new buffer
		b = bytes.NewBuffer(make([]byte, 0, bp.w))
	}
	return
}

// Put returns the given Buffer to the BufferPool.
func (bp *BufferPool) Put(b *bytes.Buffer) {
	b.Reset()
	select {
	case bp.c <- b:
	default: // Discard the buffer if the pool is full.
	}
}
