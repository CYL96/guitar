package common

import "testing"

func TestGetMyIpAddr(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetMyIpAddr()
		})
	}
}