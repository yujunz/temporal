package api

import (
	"context"

	"go.temporal.io/server/api/matchingservice/v1"
	"go.temporal.io/server/common/log"
)

// UpdateWorkerBuildIdCompatibility allows changing the worker versioning graph for a task queue
func (h *Handler) UpdateWorkerBuildIdCompatibility(
	ctx context.Context,
	request *matchingservice.UpdateWorkerBuildIdCompatibilityRequest,
) (_ *matchingservice.UpdateWorkerBuildIdCompatibilityResponse, retError error) {
	defer log.CapturePanic(h.logger, &retError)
	return h.engine.UpdateWorkerBuildIdCompatibility(ctx, request)
}
