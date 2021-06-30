package e242

import "testing"

/*
示例 1：
输入：s = "1101"
输出：true

示例 2：
输入：s = "111000"
输出：false

示例 3：
输入：s = "110100010"
输出：false
*/
func Test_checkZeroOnes(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{"1101"}, true},
		{"2", args{"111000"}, false},
		{"3", args{"110100010"}, false},
		{"4", args{"1"}, true},
		{"5", args{"10"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkZeroOnes(tt.args.s); got != tt.want {
				t.Errorf("checkZeroOnes() = %v, want %v", got, tt.want)
			}
		})
	}
}
