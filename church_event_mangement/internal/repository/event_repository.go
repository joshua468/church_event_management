package repository

import (
	"github.com/joshua468/church_event_management/internal/model"
	"gorm.io/gorm"
)

type EventRepository interface {
	CreateEvent(event *model.Event) error
	GetEventByID(id uint) (*model.Event, error)
}

type eventRepository struct {
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) EventRepository {
	return &eventRepository{db}
}

func (r *eventRepository) CreateEvent(event *model.Event) error {
	return r.db.Create(event).Error
}

func (r *eventRepository) GetEventByID(id uint) (*model.Event, error) {
	var event model.Event
	err := r.db.Where("id = ?", id).First(&event).Error
	return &event, err
}
