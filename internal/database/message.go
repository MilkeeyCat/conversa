package database

type Message struct {
	Id        int    `json:"id"`
	UserId    int    `json:"user_id"`
	Message   string `json:"message"`
	CreatedAt string `json:"created_at"`
}

// NOTE: -1 means public room
func CreateMessage(userId int, message string, roomId int) error {
	_, err := Db.Exec("INSERT INTO messages (user_id, message, room_id) VALUES (?, ?, ?)", userId, message, roomId)
	if err != nil {
		return err
	}

	return nil
}
