package main

import (
	"github.com/vinc-dev/learn-event-sourcing/config"
	"github.com/vinc-dev/learn-event-sourcing/internal/domain/eventsource"
)

func main() {
	db := config.GetDB()

	db.DropTableIfExists(&eventsource.Event{})
	db.AutoMigrate(&eventsource.Event{})
}
