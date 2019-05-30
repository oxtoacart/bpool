package bpool

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestByteSlice(t *testing.T) {
	full := []byte("abcde")
	bs := WrapByteSlice(full, 1)
	assert.EqualValues(t, full[1:], bs.Bytes())
	assert.EqualValues(t, full, bs.BytesWithHeader())
	assert.EqualValues(t, full, bs.Full())
	bs = bs.ResliceTo(2)
	assert.EqualValues(t, full[1:3], bs.Bytes())
	assert.EqualValues(t, full[:3], bs.BytesWithHeader())
	assert.EqualValues(t, full, bs.Full())
	bs = bs.ResliceTo(1)
	assert.EqualValues(t, full[1:2], bs.Bytes())
	assert.EqualValues(t, full[:2], bs.BytesWithHeader())
	assert.EqualValues(t, full, bs.Full())

}

func TestHeaderPreservingByteSlicePool(t *testing.T) {
	full := []byte{0, 0, 'a', 'b', 'c'}
	data := full[2:]
	pool := NewHeaderPreservingByteSlicePool(1, 3, 2)
	b := pool.GetSlice()
	copy(b.Bytes(), data)
	assert.Equal(t, data, b.Bytes())
	assert.Equal(t, full, b.Full())
	pool.PutSlice(b)
	assert.Equal(t, 1, pool.NumPooled())
	pool.PutSlice(WrapByteSlice(full, 2))
	assert.Equal(t, 1, pool.NumPooled(), "Pool should not grow beyond its size limit")
}
