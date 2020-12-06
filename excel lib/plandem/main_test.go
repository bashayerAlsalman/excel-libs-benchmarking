package main

import "testing"

func BenchmarkMain(t *testing.B) {
	tests := []struct {
		name string
	}{
		{"test main function"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.B) {
			main()
		})
	}
}
