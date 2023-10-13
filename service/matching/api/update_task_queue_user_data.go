package api

import (
	"context"

	"go.temporal.io/server/api/matchingservice/v1"
	"go.temporal.io/server/common/log"
)

func (h *Handler) UpdateTaskQueueUserData(
	ctx context.Context,
	request *matchingservice.UpdateTaskQueueUserDataRequest,
) (_ *matchingservice.UpdateTaskQueueUserDataResponse, retError error) {
	defer log.CapturePanic(h.logger, &retError)
	return h.engine.UpdateTaskQueueUserData(ctx, request)
}
