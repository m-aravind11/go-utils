package strings

import (
	"fmt"
	"testing"
)

func TestIsUpper(t *testing.T) {
	type args struct {
		ch rune
	}
	tests := []struct {
		args args
		want bool
	}{
		{args: args{ch: 'A'}, want: true},
		{args: args{ch: 'R'}, want: true},
		{args: args{ch: 'F'}, want: true},
		{args: args{ch: 'Q'}, want: true},
		{args: args{ch: 'Z'}, want: true},
		{args: args{ch: 'a'}, want: false},
		{args: args{ch: 'g'}, want: false},
		{args: args{ch: '0'}, want: false},
		{args: args{ch: '5'}, want: false},
		{args: args{ch: '9'}, want: false},
		{args: args{ch: 'z'}, want: false},
		{args: args{ch: '['}, want: false},
		{args: args{ch: '#'}, want: false},
		{args: args{ch: '='}, want: false},
		{args: args{ch: '\''}, want: false},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("TestIsUpper(%v)", tt.args.ch), func(t *testing.T) {
			if got := IsUpper(tt.args.ch); got != tt.want {
				t.Errorf("IsUpper(%v) = %v, want %v", tt.args.ch, got, tt.want)
			}
		})
	}
}

func TestIsLower(t *testing.T) {
	type args struct {
		ch rune
	}
	tests := []struct {
		args args
		want bool
	}{
		{args: args{ch: 'A'}, want: false},
		{args: args{ch: 'R'}, want: false},
		{args: args{ch: 'F'}, want: false},
		{args: args{ch: 'Q'}, want: false},
		{args: args{ch: 'Z'}, want: false},
		{args: args{ch: 'a'}, want: true},
		{args: args{ch: 'g'}, want: true},
		{args: args{ch: '0'}, want: false},
		{args: args{ch: '5'}, want: false},
		{args: args{ch: '9'}, want: false},
		{args: args{ch: 'z'}, want: true},
		{args: args{ch: '['}, want: false},
		{args: args{ch: '#'}, want: false},
		{args: args{ch: '='}, want: false},
		{args: args{ch: '\''}, want: false},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("TestIsLower(%v)", tt.args.ch), func(t *testing.T) {
			if got := IsLower(tt.args.ch); got != tt.want {
				t.Errorf("IsLower(%v) = %v, want %v", tt.args.ch, got, tt.want)
			}
		})
	}
}

func TestIsLetter(t *testing.T) {
	type args struct {
		ch rune
	}
	tests := []struct {
		args args
		want bool
	}{
		{args: args{ch: 'A'}, want: true},
		{args: args{ch: 'R'}, want: true},
		{args: args{ch: 'F'}, want: true},
		{args: args{ch: 'Q'}, want: true},
		{args: args{ch: 'Z'}, want: true},
		{args: args{ch: 'a'}, want: true},
		{args: args{ch: 'g'}, want: true},
		{args: args{ch: '0'}, want: false},
		{args: args{ch: '5'}, want: false},
		{args: args{ch: '9'}, want: false},
		{args: args{ch: 'z'}, want: true},
		{args: args{ch: '['}, want: false},
		{args: args{ch: '#'}, want: false},
		{args: args{ch: '='}, want: false},
		{args: args{ch: '\''}, want: false},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("TestIsLetter(%v)", tt.args.ch), func(t *testing.T) {
			if got := IsLetter(tt.args.ch); got != tt.want {
				t.Errorf("IsLetter(%v) = %v, want %v", tt.args.ch, got, tt.want)
			}
		})
	}
}

func TestIsDigit(t *testing.T) {
	type args struct {
		ch rune
	}
	tests := []struct {
		args args
		want bool
	}{
		{args: args{ch: 'A'}, want: false},
		{args: args{ch: 'R'}, want: false},
		{args: args{ch: 'F'}, want: false},
		{args: args{ch: 'Q'}, want: false},
		{args: args{ch: 'Z'}, want: false},
		{args: args{ch: 'a'}, want: false},
		{args: args{ch: 'g'}, want: false},
		{args: args{ch: '0'}, want: true},
		{args: args{ch: '5'}, want: true},
		{args: args{ch: '9'}, want: true},
		{args: args{ch: 'z'}, want: false},
		{args: args{ch: '['}, want: false},
		{args: args{ch: '#'}, want: false},
		{args: args{ch: '='}, want: false},
		{args: args{ch: '\''}, want: false},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("TestIsDigit(%v)", tt.args.ch), func(t *testing.T) {
			if got := IsDigit(tt.args.ch); got != tt.want {
				t.Errorf("TestIsDigit(%v) = %v, want %v", tt.args.ch, got, tt.want)
			}
		})
	}
}

func TestIsAlphaNumeric(t *testing.T) {
	type args struct {
		ch rune
	}
	tests := []struct {
		args args
		want bool
	}{
		{args: args{ch: 'A'}, want: true},
		{args: args{ch: 'R'}, want: true},
		{args: args{ch: 'F'}, want: true},
		{args: args{ch: 'Q'}, want: true},
		{args: args{ch: 'Z'}, want: true},
		{args: args{ch: 'a'}, want: true},
		{args: args{ch: 'g'}, want: true},
		{args: args{ch: '0'}, want: true},
		{args: args{ch: '5'}, want: true},
		{args: args{ch: '9'}, want: true},
		{args: args{ch: 'z'}, want: true},
		{args: args{ch: '['}, want: false},
		{args: args{ch: '#'}, want: false},
		{args: args{ch: '='}, want: false},
		{args: args{ch: '\''}, want: false},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("TestIsAlphaNumeric(%v)", tt.args.ch), func(t *testing.T) {
			if got := IsAlphaNumeric(tt.args.ch); got != tt.want {
				t.Errorf("TestIsAlphaNumeric(%v) = %v, want %v", tt.args.ch, got, tt.want)
			}
		})
	}
}

func TestIsVowel(t *testing.T) {
	type args struct {
		ch rune
	}
	tests := []struct {
		args args
		want bool
	}{
		{args: args{ch: 'A'}, want: true},
		{args: args{ch: 'E'}, want: true},
		{args: args{ch: 'R'}, want: false},
		{args: args{ch: 'F'}, want: false},
		{args: args{ch: 'Q'}, want: false},
		{args: args{ch: 'Z'}, want: false},
		{args: args{ch: 'a'}, want: true},
		{args: args{ch: 'u'}, want: true},
		{args: args{ch: 'g'}, want: false},
		{args: args{ch: '0'}, want: false},
		{args: args{ch: '5'}, want: false},
		{args: args{ch: '9'}, want: false},
		{args: args{ch: 'z'}, want: false},
		{args: args{ch: '['}, want: false},
		{args: args{ch: '#'}, want: false},
		{args: args{ch: '='}, want: false},
		{args: args{ch: '\''}, want: false},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("TestIsVowel(%v)", tt.args.ch), func(t *testing.T) {
			if got := IsVowel(tt.args.ch); got != tt.want {
				t.Errorf("TestIsVowel(%v) = %v, want %v", tt.args.ch, got, tt.want)
			}
		})
	}
}

func TestIsConsonant(t *testing.T) {
	type args struct {
		ch rune
	}
	tests := []struct {
		args args
		want bool
	}{
		{args: args{ch: 'A'}, want: false},
		{args: args{ch: 'E'}, want: false},
		{args: args{ch: 'R'}, want: true},
		{args: args{ch: 'F'}, want: true},
		{args: args{ch: 'Q'}, want: true},
		{args: args{ch: 'Z'}, want: true},
		{args: args{ch: 'a'}, want: false},
		{args: args{ch: 'u'}, want: false},
		{args: args{ch: 'g'}, want: true},
		{args: args{ch: '0'}, want: false},
		{args: args{ch: '5'}, want: false},
		{args: args{ch: '9'}, want: false},
		{args: args{ch: 'z'}, want: true},
		{args: args{ch: '['}, want: false},
		{args: args{ch: '#'}, want: false},
		{args: args{ch: '='}, want: false},
		{args: args{ch: '\''}, want: false},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("TestIsConsonant(%v)", tt.args.ch), func(t *testing.T) {
			if got := IsConsonant(tt.args.ch); got != tt.want {
				t.Errorf("TestIsConsonant(%v) = %v, want %v", tt.args.ch, got, tt.want)
			}
		})
	}
}

func TestToLower(t *testing.T) {
	type args struct {
		ch rune
	}
	tests := []struct {
		args args
		want rune
	}{
		{args: args{ch: 'A'}, want: 'a'},
		{args: args{ch: 'E'}, want: 'e'},
		{args: args{ch: 'R'}, want: 'r'},
		{args: args{ch: 'F'}, want: 'f'},
		{args: args{ch: 'Q'}, want: 'q'},
		{args: args{ch: 'Z'}, want: 'z'},
		{args: args{ch: 'a'}, want: 'a'},
		{args: args{ch: 'u'}, want: 'u'},
		{args: args{ch: 'g'}, want: 'g'},
		{args: args{ch: '0'}, want: '0'},
		{args: args{ch: '5'}, want: '5'},
		{args: args{ch: '9'}, want: '9'},
		{args: args{ch: 'z'}, want: 'z'},
		{args: args{ch: '['}, want: '['},
		{args: args{ch: '#'}, want: '#'},
		{args: args{ch: '='}, want: '='},
		{args: args{ch: '\''}, want: '\''},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("TestToLower(%v)", tt.args.ch), func(t *testing.T) {
			if got := ToLower(tt.args.ch); got != tt.want {
				t.Errorf("TestToLower(%v) = %v, want %v", tt.args.ch, got, tt.want)
			}
		})
	}
}

func TestToUpper(t *testing.T) {
	type args struct {
		ch rune
	}
	tests := []struct {
		args args
		want rune
	}{
		{args: args{ch: 'A'}, want: 'A'},
		{args: args{ch: 'E'}, want: 'E'},
		{args: args{ch: 'R'}, want: 'R'},
		{args: args{ch: 'F'}, want: 'F'},
		{args: args{ch: 'Q'}, want: 'Q'},
		{args: args{ch: 'Z'}, want: 'Z'},
		{args: args{ch: 'a'}, want: 'A'},
		{args: args{ch: 'u'}, want: 'U'},
		{args: args{ch: 'g'}, want: 'G'},
		{args: args{ch: '0'}, want: '0'},
		{args: args{ch: '5'}, want: '5'},
		{args: args{ch: '9'}, want: '9'},
		{args: args{ch: 'z'}, want: 'Z'},
		{args: args{ch: '['}, want: '['},
		{args: args{ch: '#'}, want: '#'},
		{args: args{ch: '='}, want: '='},
		{args: args{ch: '\''}, want: '\''},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("TestToUpper(%v)", tt.args.ch), func(t *testing.T) {
			if got := ToUpper(tt.args.ch); got != tt.want {
				t.Errorf("TestToUpper(%v) = %v, want %v", tt.args.ch, got, tt.want)
			}
		})
	}
}
