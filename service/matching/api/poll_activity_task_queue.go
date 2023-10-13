package api

import (
	"context"

	"go.temporal.io/server/api/matchingservice/v1"
	"go.temporal.io/server/common"
	"go.temporal.io/server/common/log"
	"go.temporal.io/server/common/metrics"
	"go.temporal.io/server/common/namespace"
)

// PollActivityTaskQueue - long poll for an activity task.
func (h *Handler) PollActivityTaskQueue(
	ctx context.Context,
	request *matchingservice.PollActivityTaskQueueRequest,
) (_ *matchingservice.PollActivityTaskQueueResponse, retError error) {
	defer log.CapturePanic(h.logger, &retError)
	opMetrics := h.opMetricsHandler(
		namespace.ID(request.GetNamespaceId()),
		request.GetPollRequest().GetTaskQueue(),
		metrics.MatchingPollActivityTaskQueueScope,
	)

	if request.GetForwardedSource() != "" {
		h.reportForwardedPerTaskQueueCounter(opMetrics, namespace.ID(request.GetNamespaceId()))
	}

	if _, err := common.ValidateLongPollContextTimeoutIsSet(
		ctx,
		"PollActivityTaskQueue",
		h.throttledLogger,
	); err != nil {
		return nil, err
	}

	return h.engine.PollActivityTaskQueue(ctx, request, opMetrics)
}
