package eventsource

import (
	"context"
	"errors"
)

type ServiceInterface interface {
	PostEvent(ctx context.Context, payload *PayloadCreate) (*Event, error)
}

type service struct {
}

func (s *service) PostEvent(ctx context.Context, payload *PayloadCreate) (*Event, error) {
	return nil, errors.New("unimplemented")
}

func NewService() ServiceInterface {
	return &service{}
}
