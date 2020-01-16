package tchat

type room struct {
	Name    string
	clients map[*Client]bool
}

// NewRoomはルームを設定する。
func NewRoom(name string) *room {
	return &room{
		Name:    name,
		clients: make(map[*Client]bool),
	}
}

// Addはルームにクライアントを追加する。
func (r *room) Add(c *Client) {
	r.clients[c] = true
}

// Deleteはルームからクライアントを削除する。
func (r *room) Delete(c *Client) {
	delete(r.clients, c)
}

// Sendはルームにいるクライアント全員にメッセージを送信する。
func (r *room) Send(ch <-chan []byte) {
	msg := <-ch
	for cl, status := range r.clients {
		if status {
			_, err := cl.Conn.Write(msg)
			if err != nil {
				continue
			}
		}
	}
	_ = saveMessage(msg)
}

// Showはルームに在席中のクライアントリストを返却する。
func (r *room) Show() []string {
	var clientList []string
	for cl, status := range r.clients {
		if status {
			clientList = append(clientList, cl.Name)
		}
	}
	return clientList
}
