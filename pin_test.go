package pin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateBatchPIN(t *testing.T) {
	//	logger := logging.New("extractor", "DEBUG")
	tests := []struct {
		name     string
		expected int
	}{
		{
			name:     "generate a random batch of 1000 pin codes",
			expected: 1000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenerateBatchPIN()
			assert.Equal(t, tt.expected, len(got))
		})
	}
}

func TestGeneratePIN(t *testing.T) {
	//	logger := logging.New("extractor", "DEBUG")
	tests := []struct {
		name     string
		expected int
	}{
		{
			name:     "generate a random pin of 4 digits",
			expected: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generatePIN()
			assert.Equal(t, tt.expected, len(got))
		})
	}
}
