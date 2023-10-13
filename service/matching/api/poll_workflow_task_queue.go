package api

import (
	"context"

	"go.temporal.io/server/api/matchingservice/v1"
	"go.temporal.io/server/common"
	"go.temporal.io/server/common/log"
	"go.temporal.io/server/common/metrics"
	"go.temporal.io/server/common/namespace"
)

// PollWorkflowTaskQueue - long poll for a workflow task.
func (h *Handler) PollWorkflowTaskQueue(
	ctx context.Context,
	request *matchingservice.PollWorkflowTaskQueueRequest,
) (_ *matchingservice.PollWorkflowTaskQueueResponse, retError error) {
	defer log.CapturePanic(h.logger, &retError)
	opMetrics := h.opMetricsHandler(
		namespace.ID(request.GetNamespaceId()),
		request.GetPollRequest().GetTaskQueue(),
		metrics.MatchingPollWorkflowTaskQueueScope,
	)

	if request.GetForwardedSource() != "" {
		h.reportForwardedPerTaskQueueCounter(opMetrics, namespace.ID(request.GetNamespaceId()))
	}

	if _, err := common.ValidateLongPollContextTimeoutIsSet(
		ctx,
		"PollWorkflowTaskQueue",
		h.throttledLogger,
	); err != nil {
		return nil, err
	}

	return h.engine.PollWorkflowTaskQueue(ctx, request, opMetrics)
}
