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
		{
			name: "Test new dequeue list from several empty string",
			input: []string{
				"",
				"",
			},
			wantDequeVal: []string{
				"",
				"",
			},
		},
		{
			name:         "Test new dequeue list from empty list",
			input:        []string{},
			wantDequeVal: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := NewDequeue[string](tt.input)
			for _, wantVal := range tt.wantDequeVal {
				pop := got.PopLeft()
				if wantVal != pop {
					t.Errorf("expected = %v, want %v", pop, wantVal)
				}
			}
		})
	}
}

func TestPushRight(t *testing.T) {
	tests := []testCase[string]{
		{
			name: "Test push right using string values",
			input: []string{
				"10",
				"20",
			},
			wantDequeVal: []string{
				"10",
				"20",
			},
		},
		{
			name: "Test push right from several empty string",
			input: []string{
				"",
				"",
			},
			wantDequeVal: []string{
				"",
				"",
			},
		},
		{
			name:         "Test no push right from empty list",
			input:        []string{},
			wantDequeVal: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := NewDequeue[string]([]string{})
			for _, input := range tt.input {
				list.PushRight(input)
			}
			for _, wantVal := range tt.wantDequeVal {
				pop := list.PopLeft()
				if wantVal != pop {
					t.Errorf("expected = %v, want %v", pop, wantVal)
				}
			}
		})
	}
}
