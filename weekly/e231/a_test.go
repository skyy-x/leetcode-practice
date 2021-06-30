package e231

import "testing"

func Test_checkOnesSegment(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			"1",
			args{
				"1001",
			},
			false,
		},
		{
			"11",
			args{
				"11",
			},
			true,
		},
		{
			"11011",
			args{
				"11011",
			},
			false,
		},
		{
			"",
			args{
				"",
			},
			false,
		},
		{
			"0",
			args{
				"0",
			},
			false,
		},
		{
			"1",
			args{
				"1",
			},
			true,
		},
		{
			"00100",
			args{
				"00100",
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkOnesSegment(tt.args.s); got != tt.want {
				t.Errorf("checkOnesSegment() = %v, want %v", got, tt.want)
			}
		})
	}
}
