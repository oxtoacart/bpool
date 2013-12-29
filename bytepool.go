package bpool

/*
BytePool implements a leaky pool of []byte in the form of a bounded
channel.
*/
type BytePool struct {
	c chan []byte
	w int
}

/*
NewBytePool creates a new BytePool bounded to the given size, with new byte
arrays sized based on maxWidth.
*/
func NewBytePool(size int, maxWidth int) (bp *BytePool) {
	return &BytePool{
		c: make(chan []byte, size),
		w: maxWidth,
	}
}

/*
Get gets a []byte from the BytePool, or creates a new one if none are available
in the pool.
*/
func (bp *BytePool) Get() (b []byte) {
	select {
	case b = <-bp.c:
	// reuse existing buffer
	default:
		// create new buffer
		b = make([]byte, bp.w)
	}
	return
}

/*
Put returns the given Buffer to the BytePool.
*/
func (bp *BytePool) Put(b []byte) {
	bp.c <- b
}
