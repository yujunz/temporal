package api

import (
	"context"

	"go.temporal.io/server/api/matchingservice/v1"
	"go.temporal.io/server/common/log"
)

func (h *Handler) GetBuildIdTaskQueueMapping(
	ctx context.Context,
	request *matchingservice.GetBuildIdTaskQueueMappingRequest,
) (_ *matchingservice.GetBuildIdTaskQueueMappingResponse, retError error) {
	defer log.CapturePanic(h.logger, &retError)
	return h.engine.GetBuildIdTaskQueueMapping(ctx, request)
}
