package models

import (
	pbWellness "github.com/eyolas/wellness-ws/proto"
	"gopkg.in/gorp.v1"
)

type Lesson struct {
	ID           int32  `db:"id"`
	Name         string `db:"name"`
	DayOfWeek    int32  `db:"dayOfWeek"`
	BeginHours   int32  `db:"beginHours"`
	BeginMinutes int32  `db:"beginMinutes"`
	EndHours     int32  `db:"endHours"`
	EndMinutes   int32  `db:"endMinutes"`
	Duration     int32  `db:"duration"`
	NbPlace      int32  `db:"nbPlace"`
	Roomid       int32  `db:"roomid"`
}

func (r *Lesson) ForGRPC() *pbWellness.Lesson {
	return &pbWellness.Lesson{
		Id:           r.ID,
		Name:         r.Name,
		DayOfWeek:    r.DayOfWeek,
		BeginHours:   r.BeginHours,
		BeginMinutes: r.BeginMinutes,
		EndHours:     r.EndHours,
		EndMinutes:   r.EndMinutes,
		Duration:     r.Duration,
		NbPlace:      r.NbPlace,
	}
}

type LessonService struct {
	db *gorp.DbMap
}

func NewLessonService(db *gorp.DbMap) *LessonService {
	return &LessonService{
		db: db,
	}
}
