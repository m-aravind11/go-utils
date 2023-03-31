package decimal

import (
	"testing"
)

func TestRound(t *testing.T) {
	type args struct {
		val    float64
		places int32
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "2.345 to 2 decimal places",
			args: args{
				val:    2.345,
				places: 2,
			},
			want: 2.35,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Round(tt.args.val, tt.args.places); got != tt.want {
				t.Errorf("Round() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRoundUp(t *testing.T) {
	type args struct {
		val    float64
		places int32
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "2.382 to 2 decimal places",
			args: args{
				val:    2.382,
				places: 2,
			},
			want: 2.39,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RoundUp(tt.args.val, tt.args.places); got != tt.want {
				t.Errorf("RoundUp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRoundDown(t *testing.T) {
	type args struct {
		val    float64
		places int32
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "2.389 to 2 decimal places",
			args: args{
				val:    2.382,
				places: 2,
			},
			want: 2.38,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RoundDown(tt.args.val, tt.args.places); got != tt.want {
				t.Errorf("RoundDown() = %v, want %v", got, tt.want)
			}
		})
	}
}
