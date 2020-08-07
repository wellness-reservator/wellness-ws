package server

import (
	"context"

	"github.com/eyolas/wellness-ws/config"
	pbWellness "github.com/eyolas/wellness-ws/proto"
	"github.com/eyolas/wellness-ws/server/models"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Backend implements the protobuf interface
type Backend struct {
	roomService *models.RoomService
}

// New initializes a new Backend struct.
func New() *Backend {
	return &Backend{
		roomService: models.NewRoomService(config.DbConnect),
	}
}

func (b *Backend) ListRooms(_ context.Context, _ *pbWellness.ListRoomsRequest) (*pbWellness.RoomsArray, error) {
	roomsList, err := b.roomService.GetAll()
	if err != nil {
		return nil, err
	}

	var rooms []*pbWellness.Room
	if roomsList != nil && len(*roomsList) > 0 {

		for _, room := range *roomsList {
			rooms = append(rooms, room.ForGRPC())
		}
	}

	return &pbWellness.RoomsArray{
		Rooms: rooms,
	}, nil
}

// GetConsumers lists all consumers
func (b *Backend) GetRoom(_ context.Context, req *pbWellness.GetRoomRequest) (*pbWellness.Room, error) {
	return b.roomService.GetOne(req.GetRoomId())
}

// GetConsumers lists all consumers
func (b *Backend) GetLessonsOfRoom(_ context.Context, req *pbWellness.GetRoomRequest) (*pbWellness.LessonsArray, error) {
	lessonsList, err := b.roomService.GetLessonsOfRoom(req.GetRoomId())

	if err != nil {
		return nil, err
	}

	var lessons []*pbWellness.Lesson
	if lessonsList != nil && len(*lessonsList) > 0 {

		for _, room := range *lessonsList {
			lessons = append(lessons, room.ForGRPC())
		}
	}
	return &pbWellness.LessonsArray{
		Lessons: lessons,
	}, nil
}

func (b *Backend) ListLessons(_ context.Context, req *pbWellness.ListLessonsRequest) (*pbWellness.LessonsArray, error) {
	var lessons []*pbWellness.Lesson
	if err := req.Validate(); err != nil {
		logrus.Infof("Error %q", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	logrus.Info("GetDayOfWeek", req.GetDayOfWeek())
	if req.GetDayOfWeek() != nil {
		logrus.Info("GetDayOfWeek value", req.GetDayOfWeek().Value)
	}

	return &pbWellness.LessonsArray{
		Lessons: lessons,
	}, nil
}
