package tchat

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateClient(t *testing.T) {
	actual := CreateClient(new(MockConn), "test")
	expected := "test"
	assert.Equal(t, expected, actual.Name)

}
