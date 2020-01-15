package tchat

type room struct {
	clients []*Client
}

// NewRoomはルームを設定する。
func NewRoom() *room {
	return &room{}
}

// Addはルームにクライアントを追加する。
func (r *room) Add(c *Client) {
	r.clients = append(r.clients, c)
}

// Sendはルームにいるクライアント全員にメッセージを送信する。
func (r *room) Send(ch <-chan []byte) {
	msg := <-ch
	for _, cl := range r.clients {
		_, err := cl.Conn.Write(msg)
		if err != nil {
			continue
		}
	}
}
