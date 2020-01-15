package tchat

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewRoom(t *testing.T) {
	actual := NewRoom()
	expected := &room{clients: make(map[*Client]bool)}
	assert.Equal(t, expected, actual)
}

func TestRoom_Add(t *testing.T) {
	client := NewClient(new(MockConn), "TEST-USER")
	room := NewRoom()
	room.Add(client)
	actual := len(room.clients)
	expected := 1
	assert.Equal(t, expected, actual)
}

func TestRoom_Send(t *testing.T) {
	client := NewClient(new(MockConn), "TEST-USER")
	room := NewRoom()
	room.Add(client)
	ch := make(chan []byte)
	go room.Send(ch)
	ch <- MakeMsg("test", client.Name, YELLOW)
}

func TestRoom_Delete(t *testing.T) {
	client := NewClient(new(MockConn), "TEST-USER")
	room := NewRoom()
	room.Add(client)
	room.Delete(client)
	actual := len(room.clients)
	expected := 0
	assert.Equal(t, expected, actual)
}
