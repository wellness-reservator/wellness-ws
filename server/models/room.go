package models

import (
	pbWellness "github.com/eyolas/wellness-ws/proto"
	"github.com/sirupsen/logrus"
	"gopkg.in/gorp.v1"
)

type Room struct {
	ID   int32  `db:"id"`
	Name string `db:"name"`
}

func (r *Room) ForGRPC() *pbWellness.Room {
	return &pbWellness.Room{
		Id:   r.ID,
		Name: r.Name,
	}
}

type RoomService struct {
	db *gorp.DbMap
}

func NewRoomService(db *gorp.DbMap) *RoomService {
	return &RoomService{
		db: db,
	}
}

func (c *RoomService) GetAll() (*[]Room, error) {
	var rooms []Room
	_, err := c.db.Select(&rooms, "SELECT id, name FROM room")
	if err != nil {
		logrus.Error("Error on call database", err)
		return nil, err
	}

	if len(rooms) == 0 {
		logrus.Debug("No consumers found")
		return nil, nil
	}

	return &rooms, nil
}

func (c *RoomService) GetOne(id int32) (*pbWellness.Room, error) {
	var room *Room
	err := c.db.SelectOne(&room, "SELECT id, name FROM room WHERE id=$1", id)
	if err != nil {
		logrus.Error("Error on call database", err)
		return nil, err
	}

	if room == nil {
		return nil, nil
	}

	return room.ForGRPC(), nil
}

func (c *RoomService) GetLessonsOfRoom(id int32) (*[]Lesson, error) {
	var lessons []Lesson
	_, err := c.db.Select(&lessons, "SELECT * FROM lesson where roomID=$1", id)
	if err != nil {
		logrus.Error("Error on call database", err)
		return nil, err
	}

	if len(lessons) == 0 {
		logrus.Debug("No consumers found")
		return nil, nil
	}

	return &lessons, nil
}
