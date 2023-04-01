package decimal

import (
	"fmt"
	"testing"
)

func TestRound(t *testing.T) {
	tests := []struct {
		val     interface{}
		places  int32
		want    float64
		wantErr bool
	}{
		// Test cases for float32 input
		{val: float32(1.234), places: 2, want: 1.23, wantErr: false},
		{val: float32(3.14159), places: 2, want: 3.14, wantErr: false},
		{val: float32(3.1459), places: 1, want: 3.1, wantErr: false},
		{val: float32(-3.1459), places: 2, want: -3.15, wantErr: false},
		{val: float32(3.1459), places: -1, want: 3, wantErr: false},

		// Test cases for float64 input
		{val: float64(1.234), places: 2, want: 1.23, wantErr: false},
		{val: 3.14159, places: 2, want: 3.14, wantErr: false},
		{val: 3.1459, places: 1, want: 3.1, wantErr: false},
		{val: -3.1459, places: 2, want: -3.15, wantErr: false},
		{val: 3.1459, places: -1, want: 3, wantErr: false},

		// Test cases for string input
		{val: "3.14159", places: 2, want: 3.14, wantErr: false},
		{val: "3.1459", places: 1, want: 3.1, wantErr: false},
		{val: "-3.1459", places: 2, want: -3.15, wantErr: false},
		{val: "3.1459", places: -1, want: 3, wantErr: false},
		{val: "1.234", places: 2, want: 1.23, wantErr: false},
		{val: "1.234", places: -2, want: 1, wantErr: false},
		{val: "invalid", places: 2, want: 0, wantErr: true},
		{val: "not a float", places: 2, wantErr: true},

		// Nil input
		{val: nil, places: 2, want: 0, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Round(%v, %d)", tt.val, tt.places), func(t *testing.T) {
			got, err := Round(tt.val, tt.places)

			if err != nil && !tt.wantErr {
				t.Errorf("Round(%v, %d) got error = %v, wantErr %v", tt.val, tt.places, err, tt.wantErr)
				return
			}

			if tt.want != got {
				t.Errorf("Round(%v, %d) = %v, want %v", tt.val, tt.places, got, tt.want)
			}
		})
	}
}

func TestCeil(t *testing.T) {
	tests := []struct {
		val     interface{}
		places  int32
		want    float64
		wantErr bool
	}{
		// Test cases for float32 input
		{val: float32(1.234), places: 2, want: 1.24, wantErr: false},
		{val: float32(3.14159), places: 2, want: 3.15, wantErr: false},
		{val: float32(3.1459), places: 1, want: 3.2, wantErr: false},
		{val: float32(-3.1459), places: 2, want: -3.14, wantErr: false},
		{val: float32(3.1459), places: -1, want: 4, wantErr: false},

		// Test cases for float64 input
		{val: float64(1.234), places: 2, want: 1.24, wantErr: false},
		{val: 3.14159, places: 2, want: 3.15, wantErr: false},
		{val: 3.1459, places: 1, want: 3.2, wantErr: false},
		{val: -3.1459, places: 2, want: -3.14, wantErr: false},
		{val: 3.1459, places: -1, want: 4, wantErr: false},

		// Test cases for string input
		{val: "3.14159", places: 2, want: 3.15, wantErr: false},
		{val: "3.1459", places: 1, want: 3.2, wantErr: false},
		{val: "-3.1459", places: 2, want: -3.14, wantErr: false},
		{val: "3.1459", places: -1, want: 4, wantErr: false},
		{val: "1.234", places: 2, want: 1.24, wantErr: false},
		{val: "1.234", places: -2, want: 2, wantErr: false},
		{val: "invalid", places: 2, want: 0, wantErr: true},
		{val: "not a float", places: 2, wantErr: true},

		// Nil input
		{val: nil, places: 2, want: 0, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("TestCeil(%v, %d)", tt.val, tt.places), func(t *testing.T) {
			got, err := Ceil(tt.val, tt.places)

			if err != nil && !tt.wantErr {
				t.Errorf("TestCeil(%v, %d) got error = %v, wantErr %v", tt.val, tt.places, err, tt.wantErr)
				return
			}

			if tt.want != got {
				t.Errorf("TestCeil(%v, %d) = %v, want %v", tt.val, tt.places, got, tt.want)
			}
		})
	}
}

func TestFloor(t *testing.T) {
	tests := []struct {
		val     interface{}
		places  int32
		want    float64
		wantErr bool
	}{
		// Test cases for float32 input
		{val: float32(1.234), places: 2, want: 1.23, wantErr: false},
		{val: float32(3.14159), places: 2, want: 3.14, wantErr: false},
		{val: float32(3.1459), places: 1, want: 3.1, wantErr: false},
		{val: float32(-3.1459), places: 2, want: -3.15, wantErr: false},
		{val: float32(3.1459), places: -1, want: 3, wantErr: false},

		// Test cases for float64 input
		{val: float64(1.234), places: 2, want: 1.23, wantErr: false},
		{val: 3.14159, places: 2, want: 3.14, wantErr: false},
		{val: 3.1459, places: 1, want: 3.1, wantErr: false},
		{val: -3.1459, places: 2, want: -3.15, wantErr: false},
		{val: 3.1459, places: -1, want: 3, wantErr: false},

		// Test cases for string input
		{val: "3.14159", places: 2, want: 3.14, wantErr: false},
		{val: "3.1459", places: 1, want: 3.1, wantErr: false},
		{val: "-3.1459", places: 2, want: -3.15, wantErr: false},
		{val: "3.1459", places: -1, want: 3, wantErr: false},
		{val: "1.234", places: 2, want: 1.23, wantErr: false},
		{val: "1.234", places: -2, want: 1, wantErr: false},
		{val: "invalid", places: 2, want: 0, wantErr: true},
		{val: "not a float", places: 2, wantErr: true},

		// Nil input
		{val: nil, places: 2, want: 0, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("TestFloor(%v, %d)", tt.val, tt.places), func(t *testing.T) {
			got, err := Floor(tt.val, tt.places)

			if err != nil && !tt.wantErr {
				t.Errorf("TestFloor(%v, %d) got error = %v, wantErr %v", tt.val, tt.places, err, tt.wantErr)
				return
			}

			if tt.want != got {
				t.Errorf("TestFloor(%v, %d) = %v, want %v", tt.val, tt.places, got, tt.want)
			}
		})
	}
}
