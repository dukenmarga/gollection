package deque

import (
	"testing"
)

type testCase[T any] struct {
	name         string
	input        []T
	wantDequeVal []T
}

func TestNewDequeueFromString(t *testing.T) {
	tests := []testCase[string]{
		{
			name: "Test new dequeue list from string values",
			input: []string{
				"10",
				"20",
			},
			wantDequeVal: []string{
				"10",
				"20",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := NewDequeue[string](tt.input)
			for _, wantVal := range tt.wantDequeVal {
				pop := got.PopLeft()
				if wantVal != pop {
					t.Errorf("NewDequeueList() = %v, want %v", pop, wantVal)
				}
			}
		})
	}
}
