package pin

import (
	"errors"
	"fmt"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsConsecutive(t *testing.T) {
	//	logger := logging.New("extractor", "DEBUG")
	tests := []struct {
		name     string
		pin      *[]int64
		pos      int
		expected bool
	}{
		{
			name:     "[1, 2, 3, 4] Number sequence is not consecutive",
			pin:      &[]int64{1, 2, 3, 4},
			pos:      4,
			expected: false,
		},
		{
			name:     "[2, 5, 8, 2] Number sequence is not consecutive",
			pin:      &[]int64{2, 5, 8, 2},
			pos:      4,
			expected: false,
		},
		{
			name:     "[1, 2, 3, 3] Number sequence is consecutive",
			pin:      &[]int64{1, 2, 3, 3},
			pos:      4,
			expected: true,
		},
		{
			name:     "[1, 2, 3, 3] Number sequence is not consecutive since position is only evaluating the first 2 digits",
			pin:      &[]int64{1, 2, 3, 3},
			pos:      1,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isConsecutive(tt.pin, tt.pos)
			fmt.Printf("got: %v", got)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestIsIncrementalSeq(t *testing.T) {
	//	logger := logging.New("extractor", "DEBUG")
	tests := []struct {
		name     string
		pin      *[]int64
		pos      int
		expected bool
	}{
		{
			name:     "[1, 2, 3, 4] Number sequence is incremental",
			pin:      &[]int64{1, 2, 3, 4},
			pos:      4,
			expected: true,
		},
		{
			name:     "[1, 5, 9, 2] Number sequence is not incremental",
			pin:      &[]int64{1, 5, 9, 2},
			pos:      4,
			expected: false,
		},
		{
			name:     "[1, 3, 4, 5] Number sequence is incremental",
			pin:      &[]int64{1, 3, 4, 5},
			pos:      4,
			expected: true,
		},
		{
			name:     "[8, 5, 9, 2] Number sequence is not incremental",
			pin:      &[]int64{8, 5, 9, 2},
			pos:      4,
			expected: false,
		},
		{
			name:     "[1, 3, 4, 5] Number sequence is not incremental since position is only evaluating the first 2 digits",
			pin:      &[]int64{1, 3, 4, 5},
			pos:      1,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isIncrementalSeq(tt.pin, tt.pos)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestIsUnique(t *testing.T) {
	//	logger := logging.New("extractor", "DEBUG")
	tests := []struct {
		name     string
		pins     *[]string
		pin      string
		expected bool
	}{
		{
			name:     "New pin is not unique",
			pins:     &[]string{"3456", "7385", "0987", "6542"},
			pin:      "3456",
			expected: false,
		},
		{
			name:     "New pin is unique",
			pins:     &[]string{"3456", "7385", "0987", "6542"},
			pin:      "2456",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isUnique(tt.pins, tt.pin)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestGenerateRandomNumber(t *testing.T) {
	//	logger := logging.New("extractor", "DEBUG")
	tests := []struct {
		name        string
		mockFunc    func()
		expected    int64
		errExpected bool
		errMsg      string
	}{
		{
			name: "Successfully generated random number",
			mockFunc: func() {
				randomNumberWrapper = func() (n *big.Int, err error) {
					return big.NewInt(1), nil
				}
			},
			expected:    1,
			errExpected: false,
		},
		{
			name: "Failed generate random number",
			mockFunc: func() {
				randomNumberWrapper = func() (n *big.Int, err error) {
					return nil, errors.New("can't generate random number")
				}
			},
			errExpected: true,
			errMsg:      "can't generate random number",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("\nrecovered in f:", r)
					assert.Equal(t, tt.errMsg, fmt.Sprint(r))
				}
			}()
			tt.mockFunc()
			got := generateRandomNumber()
			if tt.errExpected {
				t.Errorf("did not panic")
			} else {
				assert.Equal(t, tt.expected, got)
			}
		})
	}
}
