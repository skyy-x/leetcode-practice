package double_week_50

import (
	"reflect"
	"testing"
)

func Test_minOperations(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1, 1, 1}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minOperations(tt.args.nums); got != tt.want {
				t.Errorf("minOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}

/*
输入：points = [[1,3],[3,3],[5,3],[2,2]], queries = [[2,3,1],[4,3,1],[1,1,2]]
输出：[3,2,2]

[[1,1],[2,2],[3,3],[4,4],[5,5]]
[[1,2,2],[2,2,2],[4,3,2],[4,3,3]]
[2,3,2,4]
*/
func Test_countPoints(t *testing.T) {
	type args struct {
		points  [][]int
		queries [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{[][]int{{1, 3}, {3, 3}, {5, 3}, {2, 2}}, [][]int{{2, 3, 1}, {4, 3, 1}, {1, 1, 2}}}, []int{3, 2, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countPoints(tt.args.points, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}

/*
输入：nums = [0,1,1,3], maximumBit = 2
输出：[0,3,2,3]
*/
func Test_getMaximumXor(t *testing.T) {
	type args struct {
		nums       []int
		maximumBit int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{[]int{0, 1, 1, 3}, 2}, []int{0, 3, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMaximumXor(tt.args.nums, tt.args.maximumBit); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getMaximumXor() = %v, want %v", got, tt.want)
			}
		})
	}
}
