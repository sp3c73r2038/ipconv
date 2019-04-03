package ipconv

import (
	"testing"

	"gotest.tools/assert"
)

func Test_Conv(t *testing.T) {
	var n uint32
	var err error
	n, err = Conv("1.1.1.1")
	assert.Assert(t, err == nil)
	assert.Equal(t, n, uint32(16843009))

	n, err = Conv("192.168.1.1")
	assert.Assert(t, err == nil)
	assert.Equal(t, n, uint32(3232235777))

	n, err = Conv("223.5.5.5")
	assert.Assert(t, err == nil)
	assert.Equal(t, n, uint32(3741648133))

	n, err = Conv("172.16.254.1")
	assert.Assert(t, err == nil)
	assert.Equal(t, n, uint32(2886794753))

	/* invalid input */
	// less than 4-part
	n, err = Conv("1.1.1")
	assert.Assert(t, err != nil)
	assert.Equal(t, n, uint32(0))

	// more than 4-part
	n, err = Conv("1.1.1.1.1")
	assert.Assert(t, err != nil)
	assert.Equal(t, n, uint32(0))

	// less than 1
	n, err = Conv("0.1.1.1")
	assert.Assert(t, err != nil)
	assert.Equal(t, n, uint32(0))

	// greater than 255
	n, err = Conv("1.1.256.1")
	assert.Assert(t, err != nil)
	assert.Equal(t, n, uint32(0))

	// greater than 255
	n, err = Conv("1.1.2345.1")
	assert.Assert(t, err != nil)
	assert.Equal(t, n, uint32(0))

	// not even a number...
	n, err = Conv("1.1.a.1")
	assert.Assert(t, err != nil)
	assert.Equal(t, n, uint32(0))
}
