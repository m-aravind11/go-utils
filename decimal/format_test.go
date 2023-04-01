package decimal

import (
	"fmt"
	"testing"
)

func TestRoundAndFormat(t *testing.T) {
	tests := []struct {
		val     interface{}
		places  int32
		want    string
		wantErr bool
	}{
		// Basic test cases
		{1.234, 2, "1.23", false},
		{1.235, 2, "1.24", false},
		{1.236, 2, "1.24", false},
		{1.225, 2, "1.23", false},
		{1.215, 2, "1.22", false},
		{1.225, 1, "1.2", false},
		{1.234, 0, "1", false},
		{1234, 2, "1234.00", false},
		{float32(1.234), 2, "1.23", false},
		{"1.234", 2, "1.23", false},
		{"1", 2, "1.00", false},
		{"1.2345", 2, "1.23", false},
		{"-1.234", 2, "-1.23", false},
		{"1.239", 0, "1", false},
		{"-1.239", 0, "-1", false},
		{0, 2, "0.00", false},

		// Test cases with padding/trimming zeros
		{1.2, 4, "1.2000", false},
		{1.234, 4, "1.2340", false},
		{1.200, 4, "1.2000", false},
		{1.000, 4, "1.0000", false},
		{1, 4, "1.0000", false},
		{"1.2", 4, "1.2000", false},
		{"1.234", 4, "1.2340", false},
		{"1.200", 4, "1.2000", false},
		{"1.000", 4, "1.0000", false},
		{"1", 4, "1.0000", false},
		{"-1.2", 4, "-1.2000", false},
		{"-1.234", 4, "-1.2340", false},
		{"-1.200", 4, "-1.2000", false},
		{"-1.000", 4, "-1.0000", false},
		{"-1", 4, "-1.0000", false},
		{"0", 4, "0.0000", false},

		// Error test cases
		{"invalid", 2, "", true},
		{complex(1, 2), 2, "", true},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("TestRoundAndFormat(%v, %d)", tt.val, tt.places), func(t *testing.T) {
			got, err := RoundAndFormat(tt.val, tt.places)

			if err != nil && !tt.wantErr {
				t.Errorf("TestRoundAndFormat(%v, %d) got error = %v, wantErr %v", tt.val, tt.places, err, tt.wantErr)
				return
			}

			if tt.want != got {
				t.Errorf("TestRoundAndFormat(%v, %d) = %v, want %v", tt.val, tt.places, got, tt.want)
			}
		})
	}
}

func TestCeilAndFormat(t *testing.T) {
	tests := []struct {
		val     interface{}
		places  int32
		want    string
		wantErr bool
	}{
		// Basic test cases
		{1.234, 2, "1.24", false},
		{1.235, 2, "1.24", false},
		{1.236, 2, "1.24", false},
		{1.225, 2, "1.23", false},
		{1.215, 2, "1.22", false},
		{1.225, 1, "1.3", false},
		{1.234, 0, "2", false},
		{1234, 2, "1234.00", false},
		{float32(1.234), 2, "1.24", false},
		{"1.234", 2, "1.24", false},
		{"1", 2, "1.00", false},
		{"1.2345", 2, "1.24", false},
		{"-1.234", 2, "-1.23", false},
		{"1.239", 0, "2", false},
		{"-1.239", 0, "-1", false},
		{0, 2, "0.00", false},

		// Test cases with padding/trimming zeros
		{1.2, 4, "1.2000", false},
		{1.234, 4, "1.2340", false},
		{1.200, 4, "1.2000", false},
		{1.000, 4, "1.0000", false},
		{1, 4, "1.0000", false},
		{"1.2", 4, "1.2000", false},
		{"1.234", 4, "1.2340", false},
		{"1.200", 4, "1.2000", false},
		{"1.000", 4, "1.0000", false},
		{"1", 4, "1.0000", false},
		{"-1.2", 4, "-1.2000", false},
		{"-1.234", 4, "-1.2340", false},
		{"-1.200", 4, "-1.2000", false},
		{"-1.000", 4, "-1.0000", false},
		{"-1", 4, "-1.0000", false},
		{"0", 4, "0.0000", false},

		// Error test cases
		{"invalid", 2, "", true},
		{complex(1, 2), 2, "", true},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("TestCeilAndFormat(%v, %d)", tt.val, tt.places), func(t *testing.T) {
			got, err := CeilAndFormat(tt.val, tt.places)

			if err != nil && !tt.wantErr {
				t.Errorf("TestCeilAndFormat(%v, %d) got error = %v, wantErr %v", tt.val, tt.places, err, tt.wantErr)
				return
			}

			if tt.want != got {
				t.Errorf("TestCeilAndFormat(%v, %d) = %v, want %v", tt.val, tt.places, got, tt.want)
			}
		})
	}
}

func TestFloorAndFormat(t *testing.T) {
	tests := []struct {
		val     interface{}
		places  int32
		want    string
		wantErr bool
	}{
		// Basic test cases
		{1.234, 2, "1.23", false},
		{1.235, 2, "1.23", false},
		{1.236, 2, "1.23", false},
		{1.225, 2, "1.22", false},
		{1.215, 2, "1.21", false},
		{1.225, 1, "1.2", false},
		{1.234, 0, "1", false},
		{1234, 2, "1234.00", false},
		{float32(1.234), 2, "1.23", false},
		{"1.234", 2, "1.23", false},
		{"1", 2, "1.00", false},
		{"1.2345", 2, "1.23", false},
		{"-1.234", 2, "-1.24", false},
		{"1.239", 0, "1", false},
		{"-1.239", 0, "-2", false},
		{0, 2, "0.00", false},

		// Test cases with padding/trimming zeros
		{1.2, 4, "1.2000", false},
		{1.234, 4, "1.2340", false},
		{1.200, 4, "1.2000", false},
		{1.000, 4, "1.0000", false},
		{1, 4, "1.0000", false},
		{"1.2", 4, "1.2000", false},
		{"1.234", 4, "1.2340", false},
		{"1.200", 4, "1.2000", false},
		{"1.000", 4, "1.0000", false},
		{"1", 4, "1.0000", false},
		{"-1.2", 4, "-1.2000", false},
		{"-1.234", 4, "-1.2340", false},
		{"-1.200", 4, "-1.2000", false},
		{"-1.000", 4, "-1.0000", false},
		{"-1", 4, "-1.0000", false},
		{"0", 4, "0.0000", false},

		// Error test cases
		{"invalid", 2, "", true},
		{complex(1, 2), 2, "", true},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("TestFloorAndFormat(%v, %d)", tt.val, tt.places), func(t *testing.T) {
			got, err := FloorAndFormat(tt.val, tt.places)

			if err != nil && !tt.wantErr {
				t.Errorf("TestFloorAndFormat(%v, %d) got error = %v, wantErr %v", tt.val, tt.places, err, tt.wantErr)
				return
			}

			if tt.want != got {
				t.Errorf("TestFloorAndFormat(%v, %d) = %v, want %v", tt.val, tt.places, got, tt.want)
			}
		})
	}
}
