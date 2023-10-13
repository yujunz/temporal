package api

import (
	"context"

	"go.temporal.io/server/api/matchingservice/v1"
	"go.temporal.io/server/common/log"
)

// CancelOutstandingPoll is used to cancel outstanding pollers
func (h *Handler) CancelOutstandingPoll(ctx context.Context,
	request *matchingservice.CancelOutstandingPollRequest) (_ *matchingservice.CancelOutstandingPollResponse, retError error) {
	defer log.CapturePanic(h.logger, &retError)
	err := h.engine.CancelOutstandingPoll(ctx, request)
	return &matchingservice.CancelOutstandingPollResponse{}, err
}
