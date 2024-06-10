package service

import (
	"github.com/joshua468/church_event_management/internal/model"
	"github.com/joshua468/church_event_management/internal/repository"
)

type EventService interface {
	CreateEvent(event *model.Event) error
	GetEventByID(id uint) (*model.Event, error)
}

type eventService struct {
	repo repository.EventRepository
}

func NewEventService(repo repository.EventRepository) EventService {
	return &eventService{repo}
}

func (s *eventService) CreateEvent(event *model.Event) error {
	return s.repo.CreateEvent(event)
}

func (s *eventService) GetEventByID(id uint) (*model.Event, error) {
	return s.repo.GetEventByID(id)
}
