package double_week_48

import "testing"

func Test_secondHighest(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{"dfa12321afd"}, 2},
		{"2", args{"dfaafd"}, -1},
		{"3", args{"dfa111afd"}, -1},
		{"4", args{"dfa01afd"}, 0},
		{"4", args{"dfa011afd"}, 0},
		{"4", args{"dfa0123afd"}, 2},
		{"4", args{"dfa983210afd"}, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := secondHighest(tt.args.s); got != tt.want {
				t.Errorf("secondHighest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getMaximumConsecutive(t *testing.T) {
	type args struct {
		coins []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1, 1, 1, 4}}, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMaximumConsecutive(tt.args.coins); got != tt.want {
				t.Errorf("getMaximumConsecutive() = %v, want %v", got, tt.want)
			}
		})
	}
}
