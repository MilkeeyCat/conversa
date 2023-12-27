package database

import "database/sql"

type Room struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Token string `json:"token"`
}

type RoomNotFound struct{}

func (*RoomNotFound) Error() string {
	return "room not found"
}

func CreateRoom(userId int, name string, token string) error {
	res, err := Db.Exec("INSERT INTO rooms (name, token) VALUES(?, ?)", name, token)
	if err != nil {
		return err
	}

	roomId, err := res.LastInsertId()
	if err != nil {
		return err
	}

	_, err = Db.Exec("INSERT INTO rooms_users (user_id, room_id) VALUES(?, ?)", userId, roomId)
	if err != nil {
		return err
	}

	return nil
}

func FindRoomByToken(token string) (Room, error) {
	var room Room

	err := Db.QueryRow("SELECT * FROM rooms WHERE token = ?", token).Scan(&room.Id, &room.Name, &room.Token)
	if err != nil {
		if err == sql.ErrNoRows {
			return room, &RoomNotFound{}
		}

		return room, err
	}

	return room, nil
}

func AddUserInRoom(token string, userId int) error {
	room, err := FindRoomByToken(token)
	if err != nil {
		return err
	}

	_, err = Db.Exec("INSERT INTO rooms_users (user_id, room_id) VALUES(?, ?)", userId, room.Id)
	if err != nil {
		return err
	}

	return nil
}
