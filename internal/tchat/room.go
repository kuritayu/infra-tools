package tchat

type room struct {
	clients []*Client
}

func newRoom() *room {
	return &room{}
}

func (r *room) add(c *Client) {
	r.clients = append(r.clients, c)
}

func (r *room) send(ch <-chan []byte) {
	msg := <-ch
	for _, cl := range r.clients {
		_, err := cl.conn.Write(msg)
		if err != nil {
			continue
		}
	}
}
