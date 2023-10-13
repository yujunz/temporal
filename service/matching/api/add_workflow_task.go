package api

import (
	"context"
	"time"

	"go.temporal.io/server/api/matchingservice/v1"
	"go.temporal.io/server/common/log"
	"go.temporal.io/server/common/metrics"
	"go.temporal.io/server/common/namespace"
)

// AddWorkflowTask - adds a workflow task.
func (h *Handler) AddWorkflowTask(
	ctx context.Context,
	request *matchingservice.AddWorkflowTaskRequest,
) (_ *matchingservice.AddWorkflowTaskResponse, retError error) {
	defer log.CapturePanic(h.logger, &retError)
	startT := time.Now().UTC()
	opMetrics := h.opMetricsHandler(
		namespace.ID(request.GetNamespaceId()),
		request.GetTaskQueue(),
		metrics.MatchingAddWorkflowTaskScope,
	)

	if request.GetForwardedSource() != "" {
		h.reportForwardedPerTaskQueueCounter(opMetrics, namespace.ID(request.GetNamespaceId()))
	}

	syncMatch, err := h.engine.AddWorkflowTask(ctx, request)
	if syncMatch {
		opMetrics.Timer(metrics.SyncMatchLatencyPerTaskQueue.GetMetricName()).Record(time.Since(startT))
	}
	return &matchingservice.AddWorkflowTaskResponse{}, err
}
