package main

import (
	info "Terraria/info"
	"testing"
)

func TestGetToolInfo(t *testing.T) {
	t.Run("Returning tool info", func(t *testing.T) {
		// Test code here
		got, err := info.GetTestInfo()
		if err != nil {
			t.Errorf("Error: %v", err)
		}
		if got == nil {
			t.Errorf("Got nil, expected tool info")
		}

	})
}
