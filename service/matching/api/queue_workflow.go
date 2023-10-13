package api

import (
	"context"

	"go.temporal.io/server/api/matchingservice/v1"
	"go.temporal.io/server/common/log"
	"go.temporal.io/server/common/metrics"
	"go.temporal.io/server/common/namespace"
)

// QueryWorkflow queries a given workflow synchronously and return the query result.
func (h *Handler) QueryWorkflow(
	ctx context.Context,
	request *matchingservice.QueryWorkflowRequest,
) (_ *matchingservice.QueryWorkflowResponse, retError error) {
	defer log.CapturePanic(h.logger, &retError)
	opMetrics := h.opMetricsHandler(
		namespace.ID(request.GetNamespaceId()),
		request.GetTaskQueue(),
		metrics.MatchingQueryWorkflowScope,
	)

	if request.GetForwardedSource() != "" {
		h.reportForwardedPerTaskQueueCounter(opMetrics, namespace.ID(request.GetNamespaceId()))
	}

	return h.engine.QueryWorkflow(ctx, request)
}
