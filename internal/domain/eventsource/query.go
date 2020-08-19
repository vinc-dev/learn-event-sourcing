package eventsource

import (
	"context"
)

type QueryInterface interface {
	FindAll(ctx context.Context) ([]*Event, error)
	Find(ctx context.Context, id string) (*Event, error)
}

type query struct {
	repository RepositoryInterface
}

func (q *query) FindAll(ctx context.Context) ([]*Event, error) {
	var eventModel []*Event
	_ = q.repository.GetDB(ctx).Find(&eventModel)
	return eventModel, nil
}

func (q *query) Find(ctx context.Context, id string) (*Event, error) {
	panic("implement me")
}

func NewQuery() QueryInterface {
	return &query{
		repository: NewRepository(),
	}
}
