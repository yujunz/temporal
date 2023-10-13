package api

import (
	"context"

	"go.temporal.io/server/api/matchingservice/v1"
	"go.temporal.io/server/common/log"
)

func (h *Handler) ApplyTaskQueueUserDataReplicationEvent(
	ctx context.Context,
	request *matchingservice.ApplyTaskQueueUserDataReplicationEventRequest,
) (_ *matchingservice.ApplyTaskQueueUserDataReplicationEventResponse, retError error) {
	defer log.CapturePanic(h.logger, &retError)
	return h.engine.ApplyTaskQueueUserDataReplicationEvent(ctx, request)
}
