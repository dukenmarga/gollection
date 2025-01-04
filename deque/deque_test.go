package deque

import (
	"fmt"
	"testing"
)

type testCasePush[T any] struct {
	name         string
	input        []T
	wantDequeVal []T
}

func TestNewDequeueFromString(t *testing.T) {
	tests := []testCasePush[string]{
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
				pop, _ := got.PopLeft()
				if wantVal != pop {
					t.Errorf("expected = %v, want %v", pop, wantVal)
				}
			}
		})
	}
}

func TestPushRight(t *testing.T) {
	tests := []testCasePush[string]{
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

			currentNode := list.head
			counter := 0
			for currentNode != nil {
				if currentNode.value != tt.wantDequeVal[counter] {
					t.Errorf("expected = %v, want %v", currentNode.value, tt.wantDequeVal[counter])
				}
				currentNode = currentNode.next
				counter++
			}
		})
	}
}

func TestPushLeft(t *testing.T) {
	tests := []testCasePush[string]{
		{
			name: "Test push left using string values",
			input: []string{
				"10",
				"20",
			},
			wantDequeVal: []string{
				"20",
				"10",
			},
		},
		{
			name: "Test push left from several empty string",
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
			name:         "Test no push left from empty list",
			input:        []string{},
			wantDequeVal: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := NewDequeue[string]([]string{})
			for _, input := range tt.input {
				list.PushLeft(input)
			}

			currentNode := list.head
			counter := 0
			for currentNode != nil {
				if currentNode.value != tt.wantDequeVal[counter] {
					t.Errorf("expected = %v, want %v", currentNode.value, tt.wantDequeVal[counter])
				}
				currentNode = currentNode.next
				counter++
			}
		})
	}
}

type testCasePop[T any] struct {
	name         string
	input        []T
	wantDequeVal []T
	wantPop      T
	wantErr      error
}

func TestPopLeft(t *testing.T) {
	tests := []testCasePop[string]{
		{
			name: "Test pop left using string values",
			input: []string{
				"10",
				"20",
				"30",
			},
			wantDequeVal: []string{
				"20",
				"30",
			},
			wantPop: "10",
			wantErr: nil,
		},
		{
			name: "Test pop left from several empty string",
			input: []string{
				"",
				"",
			},
			wantDequeVal: []string{
				"",
			},
			wantPop: "",
			wantErr: nil,
		},
		{
			name:         "Test no pop left from empty list",
			input:        []string{},
			wantDequeVal: []string{},
			wantPop:      "",
			wantErr:      fmt.Errorf("list is empty"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := NewDequeue[string](tt.input)

			pop, err := list.PopLeft()
			if pop != tt.wantPop {
				t.Errorf("expected = %v, want %v", pop, tt.wantPop)
			}
			if err != nil && err.Error() != tt.wantErr.Error() {
				t.Errorf("expected = %v, want %v", err, tt.wantErr)
			}

			currentNode := list.head
			counter := 0
			for currentNode != nil {
				if currentNode.value != tt.wantDequeVal[counter] {
					t.Errorf("expected = %v, want %v", currentNode.value, tt.wantDequeVal[counter])
				}
				currentNode = currentNode.next
				counter++
			}
		})
	}
}

func TestPopRight(t *testing.T) {
	tests := []testCasePop[string]{
		{
			name: "Test pop right using string values",
			input: []string{
				"10",
				"20",
				"30",
			},
			wantDequeVal: []string{
				"10",
				"20",
			},
			wantPop: "30",
			wantErr: nil,
		},
		{
			name: "Test pop right from several empty string",
			input: []string{
				"",
				"",
			},
			wantDequeVal: []string{
				"",
			},
			wantPop: "",
			wantErr: nil,
		},
		{
			name:         "Test no pop right from empty list",
			input:        []string{},
			wantDequeVal: []string{},
			wantPop:      "",
			wantErr:      fmt.Errorf("list is empty"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := NewDequeue[string](tt.input)

			pop, err := list.PopRight()
			if pop != tt.wantPop {
				t.Errorf("expected = %v, want %v", pop, tt.wantPop)
			}
			if err != nil && err.Error() != tt.wantErr.Error() {
				t.Errorf("expected = %v, want %v", err, tt.wantErr)
			}

			currentNode := list.head
			counter := 0
			for currentNode != nil {
				if currentNode.value != tt.wantDequeVal[counter] {
					t.Errorf("expected = %v, want %v", currentNode.value, tt.wantDequeVal[counter])
				}
				currentNode = currentNode.next
				counter++
			}
		})
	}
}

type testCaseLength[T any] struct {
	name       string
	input      []T
	wantLength uint
}

func TestLengthFromSlice(t *testing.T) {
	tests := []testCaseLength[string]{
		{
			name:       "Test length using string values",
			input:      []string{"10", "20", "30"},
			wantLength: 3,
		},
		{
			name:       "Test length from several empty string",
			input:      []string{"", "", ""},
			wantLength: 3,
		},
		{
			name:       "Test no length from empty list",
			input:      []string{},
			wantLength: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := NewDequeue[string](tt.input)
			if list.length != tt.wantLength {
				t.Errorf("expected = %v, want %v", list.length, tt.wantLength)
			}
		})
	}
}

type testCaseLengthPushPop[T any] struct {
	name       string
	inputFunc  func(*DequeueList[T])
	wantLength uint
}

func TestLengthFromPushPop(t *testing.T) {
	tests := []testCaseLengthPushPop[string]{
		{
			name: "Test: push operation is same as pop",
			inputFunc: func(list *DequeueList[string]) {
				list.PushLeft("10")
				list.PushRight("20")
				list.PopLeft()
				list.PopRight()

			},
			wantLength: 0,
		},
		{
			name: "Test: push is more than pop",
			inputFunc: func(list *DequeueList[string]) {
				list.PushLeft("10")
				list.PushRight("20")
				list.PopLeft()
			},
			wantLength: 1,
		},
		{
			name: "Test: pop is more than push",
			inputFunc: func(list *DequeueList[string]) {
				list.PushRight("20")
				list.PushLeft("10")
				list.PopRight()
				list.PopRight()
				list.PopRight()
			},
			wantLength: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := &DequeueList[string]{}
			tt.inputFunc(list)
			if list.length != tt.wantLength {
				t.Errorf("expected = %v, want %v", list.length, tt.wantLength)
			}
		})
	}
}

type testCaseClear[T any] struct {
	name         string
	inputFunc    func(*DequeueList[T])
	wantLength   uint
	wantDequeVal []T
}

func TestClear(t *testing.T) {
	tests := []testCaseClear[string]{
		{
			name: "Test clear",
			inputFunc: func(list *DequeueList[string]) {
				list.PushLeft("10")
				list.PushRight("20")
				list.Clear()
			},
			wantLength:   0,
			wantDequeVal: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := &DequeueList[string]{}
			tt.inputFunc(list)
			if list.length != tt.wantLength {
				t.Errorf("expected = %v, want %v", list.length, tt.wantLength)
			}

			currentNode := list.head
			counter := 0
			for currentNode != nil {
				if currentNode.value != tt.wantDequeVal[counter] {
					t.Errorf("expected = %v, want %v", currentNode.value, tt.wantDequeVal[counter])
				}
				currentNode = currentNode.next
				counter++
			}
		})
	}
}
