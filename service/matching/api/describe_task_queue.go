package api

import (
	"context"

	"go.temporal.io/server/api/matchingservice/v1"
	"go.temporal.io/server/common/log"
)

// DescribeTaskQueue returns information about the target task queue, right now this API returns the
// pollers which polled this task queue in last few minutes. If includeTaskQueueStatus field is true,
// it will also return status of task queue's ackManager (readLevel, ackLevel, backlogCountHint and taskIDBlock).
func (h *Handler) DescribeTaskQueue(
	ctx context.Context,
	request *matchingservice.DescribeTaskQueueRequest,
) (_ *matchingservice.DescribeTaskQueueResponse, retError error) {
	defer log.CapturePanic(h.logger, &retError)
	return h.engine.DescribeTaskQueue(ctx, request)
}
