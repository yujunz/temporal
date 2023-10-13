package api

import (
	"context"

	"go.temporal.io/server/api/matchingservice/v1"
	"go.temporal.io/server/common/log"
)

func (h *Handler) ReplicateTaskQueueUserData(
	ctx context.Context,
	request *matchingservice.ReplicateTaskQueueUserDataRequest,
) (_ *matchingservice.ReplicateTaskQueueUserDataResponse, retError error) {
	defer log.CapturePanic(h.logger, &retError)
	return h.engine.ReplicateTaskQueueUserData(ctx, request)
}
