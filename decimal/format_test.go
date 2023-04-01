package decimal

import "testing"

func TestRoundAndFormat(t *testing.T) {
	type args struct {
		val    interface{}
		places int32
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "test1",
			args: args{
				val:    "21.4526",
				places: 2,
			},
			want:    "21.45",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RoundAndFormat(tt.args.val, tt.args.places)
			if (err != nil) != tt.wantErr {
				t.Errorf("RoundAndFormat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("RoundAndFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}
