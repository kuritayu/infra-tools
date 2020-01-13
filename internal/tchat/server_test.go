package tchat

import (
	"github.com/stretchr/testify/assert"
	"net"
	"testing"
)

func TestCreateClient(t *testing.T) {
	var testConn net.Conn
	actual := CreateClient(testConn, "test")
	expected := &Client{
		Name:  "test",
		conn:  testConn,
		color: actual.color,
	}
	assert.Equal(t, expected, actual)
}
