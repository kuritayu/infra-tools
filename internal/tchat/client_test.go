package tchat

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewConnection(t *testing.T) {
	actual := NewConnection(new(MockConn))
	assert.Equal(t, true, actual.Status)
}

func TestConnection_SendToServer(t *testing.T) {
	connection := NewConnection(new(MockConn))
	actual := connection.SendToServer([]byte("test"))
	expected := errors.New("dummy")
	assert.Equal(t, expected, actual)
}
