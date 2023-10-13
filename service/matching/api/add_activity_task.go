package api

import (
	"context"
	"time"

	"go.temporal.io/server/api/matchingservice/v1"
	"go.temporal.io/server/common/log"
	"go.temporal.io/server/common/metrics"
	"go.temporal.io/server/common/namespace"
)

// AddActivityTask - adds an activity task.
func (h *Handler) AddActivityTask(
	ctx context.Context,
	request *matchingservice.AddActivityTaskRequest,
) (_ *matchingservice.AddActivityTaskResponse, retError error) {
	defer log.CapturePanic(h.logger, &retError)
	startT := time.Now().UTC()
	opMetrics := h.opMetricsHandler(
		namespace.ID(request.GetNamespaceId()),
		request.GetTaskQueue(),
		metrics.MatchingAddActivityTaskScope,
	)

	if request.GetForwardedSource() != "" {
		h.reportForwardedPerTaskQueueCounter(opMetrics, namespace.ID(request.GetNamespaceId()))
	}

	syncMatch, err := h.engine.AddActivityTask(ctx, request)
	if syncMatch {
		opMetrics.Timer(metrics.SyncMatchLatencyPerTaskQueue.GetMetricName()).Record(time.Since(startT))
	}

	return &matchingservice.AddActivityTaskResponse{}, err
}
