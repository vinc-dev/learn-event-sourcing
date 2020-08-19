package event

import (
	"context"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/vinc-dev/learn-event-sourcing/config"
	"github.com/vinc-dev/learn-event-sourcing/internal/domain/eventsource"
	"github.com/vinc-dev/learn-event-sourcing/lib/http/response"
)

func Create(eventService eventsource.ServiceInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		ctx = context.WithValue(ctx, config.GormConfigKey, config.GetDB())

		now := time.Now()

		result, err := eventService.PostEvent(ctx, &eventsource.PayloadCreate{
			Category:  "event",
			EntityId:  uuid.New().String(),
			EntitySeq: 1,
			Type:      "EVENT_CREATED",
			Data:      nil,
			CreatedAt: &now,
		})
		if nil != err {
			response.Text(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response.Json(w, result, http.StatusOK)
		return
	}
}
