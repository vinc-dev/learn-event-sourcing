package eventsource

import (
	"encoding/json"

	"github.com/jinzhu/gorm"
)

type Event struct {
	gorm.Model
	EntityCategory string          `json:"entityCategory" gorm:"entity_category"`
	EntityId       string          `json:"entityId" gorm:"entity_id"`
	EntitySeq      int32           `json:"entitySeq" gorm:"entity_seq"`
	EntityType     string          `json:"entityType" gorm:"entity_type"`
	EntityData     json.RawMessage `json:"entityData" gorm:"entity_data"`
}
