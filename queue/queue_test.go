package queue_test

import (
	"testing"

	"github.com/wxdis-public/goalgos/queue"
)

func TestEnqueue(t *testing.T) {
	tt := map[string]struct {
		values        []int
		queueLength   int
		dequeueCount  int
		expectedBack  int
		expectedFront int
		expectedEmpty bool
	}{
		"Enqueue no values": {
			values:        nil,
			queueLength:   0,
			expectedBack:  0,
			expectedFront: 0,
			expectedEmpty: true,
		},
		"Enqueue single value": {
			values:        []int{5},
			queueLength:   1,
			expectedBack:  5,
			expectedFront: 5,
		},
		"Enqueue two values": {
			values:        []int{5, 7},
			queueLength:   2,
			expectedBack:  7,
			expectedFront: 5,
		},
		"Enqueue three values": {
			values:        []int{5, 6, 7},
			queueLength:   3,
			expectedBack:  7,
			expectedFront: 5,
		},
		"Enqueue three values with a single dequeue": {
			values:        []int{5, 6, 7},
			queueLength:   3,
			dequeueCount:  1,
			expectedBack:  7,
			expectedFront: 6,
		},
		"Enqueue a single value with a single dequeue": {
			values:        []int{5},
			queueLength:   1,
			dequeueCount:  1,
			expectedBack:  0,
			expectedFront: 0,
			expectedEmpty: true,
		},
		"Enqueue two values with two dequeues": {
			values:        []int{5, 6},
			queueLength:   2,
			dequeueCount:  2,
			expectedBack:  0,
			expectedFront: 0,
			expectedEmpty: true,
		},
		"Enqueue three values with two dequeues": {
			values:        []int{5, 6, 7},
			queueLength:   3,
			dequeueCount:  2,
			expectedBack:  7,
			expectedFront: 7,
		},
		"Enqueue three values to a 1 capacity queue": {
			values:        []int{5, 6, 7},
			queueLength:   1,
			expectedBack:  5,
			expectedFront: 5,
		},
		"Enqueue zero values to a 1 capacity queue": {
			values:        nil,
			queueLength:   1,
			expectedBack:  0,
			expectedFront: 0,
			expectedEmpty: true,
		},
		"Dequeue from single capacity empty queue": {
			values:        nil,
			queueLength:   1,
			dequeueCount:  1,
			expectedBack:  0,
			expectedFront: 0,
			expectedEmpty: true,
		},
	}

	for tn, tc := range tt {
		q := queue.New[int](tc.queueLength)
		q.Enqueue(tc.values...)
		for i := 0; i < tc.dequeueCount; i++ {
			q.Dequeue()
		}
		frontVal, frontPass := q.Front()
		if frontPass == tc.expectedEmpty {
			t.Errorf("%s failed: expected to be able to fetch a value from the front of the queue, but failed", tn)
		} else if frontVal != tc.expectedFront {
			t.Errorf("%s failed: expected '%v' from the front of the queue, but got '%v'", tn, tc.expectedFront, frontVal)
		}
		backVal, backPass := q.Back()
		if backPass == tc.expectedEmpty {
			t.Errorf("%s failed: expected to be able to fetch a value from the back of the queue, but failed", tn)
		} else if backVal != tc.expectedBack {
			t.Errorf("%s failed: expected '%v' from the back of the queue, but got '%v'", tn, tc.expectedBack, backVal)
		}
	}
}
