package database

type Message struct {
	Id        int    `json:"id"`
	UserId    int    `json:"user_id"`
	Message   string `json:"message"`
	CreatedAt string `json:"created_at"`
}

type MessageWithUser struct {
	Id        int    `json:"id"`
	Username  string `json:"username"`
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

func GetRoomMessagesInRoom(roomId int) ([]MessageWithUser, error) {
	messages := []MessageWithUser{}

	rows, err := Db.Query("SELECT id, user_id, message, created_at FROM messages WHERE room_id = ?", roomId)
	if err != nil {
		return messages, err
	}
	defer rows.Close()

	for rows.Next() {
		msg := Message{}

		if err := rows.Scan(&msg.Id, &msg.UserId, &msg.Message, &msg.CreatedAt); err != nil {
			return messages, err
		}

		usr, err := FindUserById(msg.UserId)
		if err != nil {
			return messages, err
		}

		messages = append(messages, MessageWithUser{
			Id:        msg.Id,
			Username:  usr.Name,
			UserId:    msg.UserId,
			Message:   msg.Message,
			CreatedAt: msg.CreatedAt,
		})
	}

	return messages, nil
}
