package api

import (
	"context"

	"go.temporal.io/server/api/matchingservice/v1"
	"go.temporal.io/server/common/log"
)

// GetWorkerBuildIdCompatibility fetches the worker versioning data for a task queue
func (h *Handler) GetWorkerBuildIdCompatibility(
	ctx context.Context,
	request *matchingservice.GetWorkerBuildIdCompatibilityRequest,
) (_ *matchingservice.GetWorkerBuildIdCompatibilityResponse, retError error) {
	defer log.CapturePanic(h.logger, &retError)
	return h.engine.GetWorkerBuildIdCompatibility(ctx, request)
}
