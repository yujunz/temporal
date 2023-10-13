package api

import (
	"context"

	"go.temporal.io/server/api/matchingservice/v1"
	"go.temporal.io/server/common/log"
)

// ListTaskQueuePartitions returns information about partitions for a taskQueue
func (h *Handler) ListTaskQueuePartitions(
	ctx context.Context,
	request *matchingservice.ListTaskQueuePartitionsRequest,
) (_ *matchingservice.ListTaskQueuePartitionsResponse, retError error) {
	defer log.CapturePanic(h.logger, &retError)
	return h.engine.ListTaskQueuePartitions(ctx, request)
}
