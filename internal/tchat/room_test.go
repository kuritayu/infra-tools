package tchat

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewRoom(t *testing.T) {
	actual := NewRoom()
	assert.Equal(t, &room{}, actual)
}

func TestRoom_Add(t *testing.T) {
	client := CreateClient(new(MockConn), "TEST-USER")
	room := NewRoom()
	room.Add(client)
	actual := len(room.clients)
	expected := 1
	assert.Equal(t, expected, actual)
}

func TestRoom_Send(t *testing.T) {
	client := CreateClient(new(MockConn), "TEST-USER")
	room := NewRoom()
	room.Add(client)
	ch := make(chan []byte)
	go room.Send(ch)
	ch <- MakeMsg("test", client.Name, YELLOW)
}
