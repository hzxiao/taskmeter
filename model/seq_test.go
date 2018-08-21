package model

import (
	"testing"
	"github.com/hzxiao/goutil/assert"
)

func TestNextSeq(t *testing.T) {
	removeAll()

	value1, err := NextSeq("user", 1)
	assert.NoError(t, err)
	assert.Equal(t, uint64(1), value1)

	value2, err := NextSeq("user", 1)
	assert.NoError(t, err)
	assert.Equal(t, uint64(2), value2)
}
