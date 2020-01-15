package tchat

type room struct {
	clients []*Client
}

func NewRoom() *room {
	return &room{}
}

func (r *room) Add(c *Client) {
	r.clients = append(r.clients, c)
}

func (r *room) Send(ch <-chan []byte) {
	msg := <-ch
	for _, cl := range r.clients {
		_, err := cl.Conn.Write(msg)
		if err != nil {
			continue
		}
	}
}
