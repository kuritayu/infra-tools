package lstar

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	actual := New("../../test/test.tar")
	var expected *Tar
	assert.IsType(t, expected, actual)
}
