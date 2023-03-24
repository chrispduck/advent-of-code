package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemove(t *testing.T) {
	s := []string{"A", "B", "C", "D", "E"}
	s = IdempotentRemove(s, "C")
	if len(s) != 4 {
		t.Errorf("Expected len(s) to be 4, got %d", len(s))
	}
	assert.Equal(t, []string{"A", "B", "D", "E"}, s)
}

func TestAdd(t *testing.T) {
	s := []string{"A", "B", "C", "D", "E"}
	s = IdempotentAdd(s, "C")
	if len(s) != 5 {
		t.Errorf("Expected len(s) to be 5, got %d", len(s))
	}
	assert.Equal(t, []string{"A", "B", "C", "D", "E"}, s)
}
