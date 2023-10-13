package api

import (
	"context"

	"go.temporal.io/server/api/matchingservice/v1"
	"go.temporal.io/server/common/log"
)

func (h *Handler) GetTaskQueueUserData(
	ctx context.Context,
	request *matchingservice.GetTaskQueueUserDataRequest,
) (_ *matchingservice.GetTaskQueueUserDataResponse, retError error) {
	defer log.CapturePanic(h.logger, &retError)
	return h.engine.GetTaskQueueUserData(ctx, request)
}
