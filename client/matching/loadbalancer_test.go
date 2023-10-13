package matching

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	enumspb "go.temporal.io/api/enums/v1"
	"go.temporal.io/server/common/tqname"
)

func TestTQLoadBalancer(t *testing.T) {
	partitionCount := 4
	key := taskQueueKey{
		NamespaceID: "fake-namespace-id",
		Name:        mustParseTQName("fake-taskqueue"),
		Type:        enumspb.TASK_QUEUE_TYPE_ACTIVITY,
	}
	tqlb := newTaskQueueLoadBalancer(key)

	// pick 4 times, each partition picked would have one poller
	tqlb.pickReadPartition(partitionCount, -1)
	assert.Equal(t, 1, maxPollerCount(tqlb))
	tqlb.pickReadPartition(partitionCount, -1)
	assert.Equal(t, 1, maxPollerCount(tqlb))
	tqlb.pickReadPartition(partitionCount, -1)
	assert.Equal(t, 1, maxPollerCount(tqlb))
	p3 := tqlb.pickReadPartition(partitionCount, -1)
	assert.Equal(t, 1, maxPollerCount(tqlb))

	// release one, and pick one, the newly picked one should have one poller
	p3.Release()
	tqlb.pickReadPartition(partitionCount, -1)
	assert.Equal(t, 1, maxPollerCount(tqlb))

	// pick one again, this time it should have 2 pollers
	tqlb.pickReadPartition(partitionCount, -1)
	assert.Equal(t, 2, maxPollerCount(tqlb))
}

func TestTQLoadBalancerForce(t *testing.T) {
	partitionCount := 4
	key := taskQueueKey{
		NamespaceID: "fake-namespace-id",
		Name:        mustParseTQName("fake-taskqueue"),
		Type:        enumspb.TASK_QUEUE_TYPE_ACTIVITY,
	}
	tqlb := newTaskQueueLoadBalancer(key)

	// pick 4 times, each partition picked would have one poller
	p1 := tqlb.pickReadPartition(partitionCount, 1)
	assert.Equal(t, 1, p1.partitionID)
	assert.Equal(t, 1, maxPollerCount(tqlb))
	tqlb.pickReadPartition(partitionCount, 1)
	assert.Equal(t, 2, maxPollerCount(tqlb))

	// when we don't force it should balance out
	tqlb.pickReadPartition(partitionCount, -1)
	tqlb.pickReadPartition(partitionCount, -1)
	tqlb.pickReadPartition(partitionCount, -1)
	tqlb.pickReadPartition(partitionCount, -1)
	tqlb.pickReadPartition(partitionCount, -1)
	tqlb.pickReadPartition(partitionCount, -1)
	assert.Equal(t, 2, maxPollerCount(tqlb))

	// releasing the forced one and adding another should still be balanced
	p1.Release()
	tqlb.pickReadPartition(partitionCount, -1)
	assert.Equal(t, 2, maxPollerCount(tqlb))

	tqlb.pickReadPartition(partitionCount, -1)
	assert.Equal(t, 3, maxPollerCount(tqlb))
}

func TestLoadBalancerConcurrent(t *testing.T) {
	wg := &sync.WaitGroup{}
	partitionCount := 4
	key := taskQueueKey{
		NamespaceID: "fake-namespace-id",
		Name:        mustParseTQName("fake-taskqueue"),
		Type:        enumspb.TASK_QUEUE_TYPE_ACTIVITY,
	}
	tqlb := newTaskQueueLoadBalancer(key)

	concurrentCount := 10 * partitionCount
	wg.Add(concurrentCount)
	for i := 0; i < concurrentCount; i++ {
		go func() {
			defer wg.Done()
			tqlb.pickReadPartition(partitionCount, -1)
		}()
	}
	wg.Wait()

	// verify all partition have same pollers.
	assert.Equal(t, 10, tqlb.pollerCounts[0])
	assert.Equal(t, 10, tqlb.pollerCounts[1])
	assert.Equal(t, 10, tqlb.pollerCounts[2])
	assert.Equal(t, 10, tqlb.pollerCounts[3])
}

func TestLoadBalancer_ReducedPartitionCount(t *testing.T) {
	partitionCount := 2
	key := taskQueueKey{
		NamespaceID: "fake-namespace-id",
		Name:        mustParseTQName("fake-taskqueue"),
		Type:        enumspb.TASK_QUEUE_TYPE_ACTIVITY,
	}
	tqlb := newTaskQueueLoadBalancer(key)
	p1 := tqlb.pickReadPartition(partitionCount, -1)
	p2 := tqlb.pickReadPartition(partitionCount, -1)
	assert.Equal(t, 1, maxPollerCount(tqlb))
	assert.Equal(t, 1, maxPollerCount(tqlb))

	partitionCount += 2 // increase partition count
	p3 := tqlb.pickReadPartition(partitionCount, -1)
	p4 := tqlb.pickReadPartition(partitionCount, -1)
	assert.Equal(t, 1, maxPollerCount(tqlb))
	assert.Equal(t, 1, maxPollerCount(tqlb))

	partitionCount -= 2 // reduce partition count
	p5 := tqlb.pickReadPartition(partitionCount, -1)
	p6 := tqlb.pickReadPartition(partitionCount, -1)
	assert.Equal(t, 2, maxPollerCount(tqlb))
	assert.Equal(t, 2, maxPollerCount(tqlb))
	p7 := tqlb.pickReadPartition(partitionCount, -1)
	assert.Equal(t, 3, maxPollerCount(tqlb))

	// release all of them and it should be ok.
	p1.Release()
	p2.Release()
	p3.Release()
	p4.Release()
	p5.Release()
	p6.Release()
	p7.Release()

	tqlb.pickReadPartition(partitionCount, -1)
	tqlb.pickReadPartition(partitionCount, -1)
	assert.Equal(t, 1, maxPollerCount(tqlb))
	assert.Equal(t, 1, maxPollerCount(tqlb))
	tqlb.pickReadPartition(partitionCount, -1)
	assert.Equal(t, 2, maxPollerCount(tqlb))
}

func mustParseTQName(baseName string) tqname.Name {
	n, err := tqname.Parse(baseName)
	if err != nil {
		panic(err)
	}
	return n
}

func maxPollerCount(tqlb *tqLoadBalancer) int {
	res := -1
	for _, c := range tqlb.pollerCounts {
		if c > res {
			res = c
		}
	}
	return res
}