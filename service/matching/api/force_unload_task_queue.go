package api

import (
	"context"

	"go.temporal.io/server/api/matchingservice/v1"
	"go.temporal.io/server/common/log"
)

func (h *Handler) ForceUnloadTaskQueue(
	ctx context.Context,
	request *matchingservice.ForceUnloadTaskQueueRequest,
) (_ *matchingservice.ForceUnloadTaskQueueResponse, retError error) {
	defer log.CapturePanic(h.logger, &retError)
	return h.engine.ForceUnloadTaskQueue(ctx, request)
}
