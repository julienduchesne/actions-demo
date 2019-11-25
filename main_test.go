package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetGreeting(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name     string
		greeting string
	}{
		{
			name:     "Coveo",
			greeting: "Hello, Coveo!",
		},
		{
			name:     "People",
			greeting: "Hello, People!",
		},
		{
			name:     "Montreal",
			greeting: "Hello, Montreal!",
		},
	}
	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.greeting, GetGreeting(tt.name))
		})
	}
}
