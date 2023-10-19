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

package xdc

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	"go.temporal.io/api/workflowservice/v1"
	sdkclient "go.temporal.io/sdk/client"
	sdkworker "go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"
	"go.temporal.io/server/common"
	"go.temporal.io/server/common/primitives/timestamp"
)

type replicationFailureSuite struct {
	xdcBaseSuite
}

func TestReplicationFailureSuite(t *testing.T) {
	suite.Run(t, new(replicationFailureSuite))
}

func (s *replicationFailureSuite) SetupSuite() {
	s.setupSuite([]string{"replication_failure_test_active_cluster", "replication_failure_test_standby_cluster"})
}

func (s *replicationFailureSuite) TearDownSuite() {
	s.tearDownSuite()
}

func (s *replicationFailureSuite) SetupTest() {
	s.setupTest()
}

func (s *replicationFailureSuite) TestReplicationFailure() {
	for _, tc := range []struct {
		name       string
		shouldWait bool
	}{
		{
			name:       "no_wait",
			shouldWait: false,
		},
		{
			name:       "wait",
			shouldWait: true,
		},
	} {
		s.Run(tc.name, func() {
			ns := s.T().Name() + "-" + common.GenerateRandomString(8)
			activeClient, err := sdkclient.Dial(sdkclient.Options{
				HostPort:  s.cluster1.GetHost().FrontendGRPCAddress(),
				Namespace: ns,
			})
			s.NoError(err)

			ctx := context.Background()

			myWorkflow := func(ctx workflow.Context) (string, error) {
				return "hello", nil
			}
			tq := "test-task-queue-" + s.T().Name()
			worker := sdkworker.New(activeClient, tq, sdkworker.Options{})
			worker.RegisterWorkflow(myWorkflow)

			_, err = s.cluster1.GetFrontendClient().RegisterNamespace(ctx, &workflowservice.RegisterNamespaceRequest{
				Namespace:                        ns,
				Clusters:                         s.clusterReplicationConfig(),
				ActiveClusterName:                s.clusterNames[0],
				IsGlobalNamespace:                true,
				WorkflowExecutionRetentionPeriod: timestamp.DurationPtr(1 * time.Hour * 24),
			})
			s.NoError(err)
			s.NoError(worker.Start())
			defer worker.Stop()

			if tc.shouldWait {
				time.Sleep(cacheRefreshInterval)
			}

			run, err := activeClient.ExecuteWorkflow(ctx, sdkclient.StartWorkflowOptions{
				TaskQueue: tq,
			}, myWorkflow)
			s.NoError(err)

			var result string
			err = run.Get(ctx, &result)
			s.NoError(err)
			s.Equal("hello", result)
			standbyClient, err := sdkclient.Dial(sdkclient.Options{
				HostPort:  s.cluster2.GetHost().FrontendGRPCAddress(),
				Namespace: ns,
			})
			s.NoError(err)

			t := time.Now()
			for time.Now().Sub(t) < 10*time.Second {
				run = standbyClient.GetWorkflow(ctx, run.GetID(), run.GetRunID())
				result = ""
				err = run.Get(ctx, &result)
				s.NoError(ctx.Err(), "Timed out waiting to get workflow from standby cluster")

				if err == nil {
					break
				}

				time.Sleep(100 * time.Millisecond)
			}
			if tc.shouldWait {
				s.NoError(err, "If we wait, the workflow should be replicated")
				s.Equal("hello", result)
			} else {
				s.Error(err, "If we don't wait, the workflow should not be replicated")
			}
		})
	}
}
