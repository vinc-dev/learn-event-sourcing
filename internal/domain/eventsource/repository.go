package eventsource

import (
	"context"

	"github.com/vinc-dev/learn-event-sourcing/config"

	"github.com/jinzhu/gorm"
)

type RepositoryInterface interface {
	GetDB(ctx context.Context) *gorm.DB
}

type repository struct {
}

func (r *repository) GetDB(ctx context.Context) *gorm.DB {
	return ctx.Value(config.GormConfigKey).(*gorm.DB)
}

func NewRepository() RepositoryInterface {
	return &repository{}
}
