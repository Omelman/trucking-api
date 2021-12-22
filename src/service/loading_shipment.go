package service

import (
	"context"

	log "github.com/sirupsen/logrus"
)

func (s Service) ScheduledLoadingShipment() {
	ctx := context.Background()

	log.WithContext(ctx).Info("start loading ..")
}
