package tchat

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewRoom(t *testing.T) {
	actual := NewRoom("PUBLIC")
	expected := &room{
		Name:    "PUBLIC",
		clients: make(map[*Client]bool),
	}
	assert.Equal(t, expected, actual)
}

func TestRoom_Add(t *testing.T) {
	client := NewClient(new(MockConn), "TEST-USER")
	room := NewRoom("PUBLIC")
	room.Add(client)
	actual := len(room.clients)
	expected := 1
	assert.Equal(t, expected, actual)
}

func TestRoom_Send(t *testing.T) {
	client := NewClient(new(MockConn), "TEST-USER")
	room := NewRoom("PUBLIC")
	room.Add(client)
	ch := make(chan []byte)
	go room.Send(ch)
	ch <- MakeMsg("test", client.Name, YELLOW)
}

func TestRoom_Delete(t *testing.T) {
	client := NewClient(new(MockConn), "TEST-USER")
	room := NewRoom("PUBLIC")
	room.Add(client)
	room.Delete(client)
	actual := len(room.clients)
	expected := 0
	assert.Equal(t, expected, actual)
}

func TestRoom_Show(t *testing.T) {
	room := NewRoom("PUBLIC")
	client1 := NewClient(new(MockConn), "TEST-USER1")
	client2 := NewClient(new(MockConn), "TEST-USER2")
	room.Add(client1)
	room.Add(client2)
	actual := len(room.Show())
	expected := 2
	assert.Equal(t, expected, actual)

}
