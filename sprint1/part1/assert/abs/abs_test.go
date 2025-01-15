package abs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAbs(t *testing.T) {
	tests := []struct {
		name  string
		value float64
		want  float64
	}{
		{"negative value", -3, 3},
		{"positive value", 3, 3},
		{"negative float value", -2.000001, 2.000001},
		{"negative value", -0.000003, 0.000003},
		{"negative zero", -0, 0},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, Abs(test.value))
		})
	}
}
