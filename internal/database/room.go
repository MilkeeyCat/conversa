package database

func CreateRoom(name string, userId int) error {
	res, err := Db.Exec("INSERT INTO rooms (name) VALUES(?)", name)
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
