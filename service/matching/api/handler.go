// The MIT License
//
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
//
// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package api

import (
	"sync"

	"go.temporal.io/server/service/matching"

	taskqueuepb "go.temporal.io/api/taskqueue/v1"

	"go.temporal.io/server/api/matchingservice/v1"
	"go.temporal.io/server/common/cluster"
	"go.temporal.io/server/common/log"
	"go.temporal.io/server/common/membership"
	"go.temporal.io/server/common/metrics"
	"go.temporal.io/server/common/namespace"
	"go.temporal.io/server/common/persistence"
	"go.temporal.io/server/common/persistence/visibility/manager"
	"go.temporal.io/server/common/resource"
)

type (
	// Handler - gRPC handler interface for matchingservice
	Handler struct {
		engine            matching.Engine
		config            *matching.Config
		metricsHandler    metrics.Handler
		logger            log.Logger
		startWG           sync.WaitGroup
		throttledLogger   log.Logger
		namespaceRegistry namespace.Registry
	}
)

var (
	_ matchingservice.MatchingServiceServer = (*Handler)(nil)
)

// NewHandler creates a gRPC handler for the matchingservice
func NewHandler(
	config *matching.Config,
	logger log.Logger,
	throttledLogger log.Logger,
	taskManager persistence.TaskManager,
	historyClient resource.HistoryClient,
	matchingRawClient resource.MatchingRawClient,
	matchingServiceResolver membership.ServiceResolver,
	metricsHandler metrics.Handler,
	namespaceRegistry namespace.Registry,
	clusterMetadata cluster.Metadata,
	namespaceReplicationQueue persistence.NamespaceReplicationQueue,
	visibilityManager manager.VisibilityManager,
) *Handler {
	handler := &Handler{
		config:          config,
		metricsHandler:  metricsHandler,
		logger:          logger,
		throttledLogger: throttledLogger,
		engine: matching.NewEngine(
			taskManager,
			historyClient,
			matchingRawClient, // Use non retry client inside matching
			config,
			logger,
			throttledLogger,
			metricsHandler,
			namespaceRegistry,
			matchingServiceResolver,
			clusterMetadata,
			namespaceReplicationQueue,
			visibilityManager,
		),
		namespaceRegistry: namespaceRegistry,
	}

	// prevent from serving requests before matching engine is started and ready
	handler.startWG.Add(1)

	return handler
}

// Start starts the handler
func (h *Handler) Start() {
	h.engine.Start()
	h.startWG.Done()
}

// Stop stops the handler
func (h *Handler) Stop() {
	h.engine.Stop()
}

func (h *Handler) opMetricsHandler(
	namespaceID namespace.ID,
	taskQueue *taskqueuepb.TaskQueue,
	operation string,
) metrics.Handler {
	return metrics.GetPerTaskQueueScope(
		h.metricsHandler.WithTags(metrics.OperationTag(operation)),
		h.namespaceName(namespaceID).String(),
		taskQueue.GetName(),
		taskQueue.GetKind())
}

func (h *Handler) namespaceName(id namespace.ID) namespace.Name {
	entry, err := h.namespaceRegistry.GetNamespaceByID(id)
	if err != nil {
		return ""
	}
	return entry.Name()
}

func (h *Handler) reportForwardedPerTaskQueueCounter(opMetrics metrics.Handler, namespaceId namespace.ID) {
	opMetrics.Counter(metrics.ForwardedPerTaskQueueCounter.GetMetricName()).Record(1)
	h.metricsHandler.Counter(metrics.MatchingClientForwardedCounter.GetMetricName()).
		Record(
			1,
			metrics.OperationTag(metrics.MatchingAddWorkflowTaskScope),
			metrics.NamespaceTag(h.namespaceName(namespaceId).String()),
			metrics.ServiceRoleTag(metrics.MatchingRoleTagValue))
}
