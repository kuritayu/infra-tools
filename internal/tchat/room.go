package tchat

//TODO ルーム名を設定する
//TODO ルームに入っているクライアントの一覧を表示する
type room struct {
	clients map[*Client]bool
}

// NewRoomはルームを設定する。
func NewRoom() *room {
	return &room{
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
}
