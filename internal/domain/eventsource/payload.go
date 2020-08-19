package eventsource

import (
	"encoding/json"
	"time"
)

type PayloadCreate struct {
	Category  string          `json:"category" db:"category"`
	EntityId  string          `json:"entityId" db:"entity_id"`
	EntitySeq int32           `json:"entitySeq" db:"entity_seq"`
	Type      string          `json:"type" db:"type"`
	Data      json.RawMessage `json:"data" db:"data"`
	CreatedAt *time.Time      `json:"createdAt" db:"created_at"`
}

type PayloadUpdate struct {
	PayloadCreate
	Id string `json:"id"`
}

type PayloadFilter struct {
}
