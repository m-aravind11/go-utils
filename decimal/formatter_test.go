package decimal

import "testing"

func TestFormatWithRound(t *testing.T) {
	type args struct {
		val    float64
		places int32
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add testcases
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatWithRound(tt.args.val, tt.args.places); got != tt.want {
				t.Errorf("FormatWithRound() = %v, want %v", got, tt.want)
			}
		})
	}
}
