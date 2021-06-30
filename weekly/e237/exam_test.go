package e237

import (
	"reflect"
	"testing"
)

/*
[[1,2],[2,4],[3,2],[4,1]]
[0,2,3,1]

[[7,10],[7,12],[7,5],[7,4],[7,2]]
[4,3,2,0,1]

[[19,13],[16,9],[21,10],[32,25],[37,4],[49,24],[2,15],[38,41],[37,34],[33,6],[45,4],[18,18],[46,39],[12,24]]
[6,1,2,9,4,10,0,11,5,13,3,8,12,7]

[[35,36],[11,7],[15,47],[34,2],[47,19],[16,14],[19,8],[7,34],[38,15],[16,18],[27,22],[7,15],[43,2],[10,5],[5,4],[3,11]]
[15,14,13,1,6,3,5,12,8,11,9,4,10,7,0,2]
*/
func Test_getOrder(t *testing.T) {
	type args struct {
		tasks [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{[][]int{{1, 2}, {2, 4}, {3, 2}, {4, 1}}}, []int{0, 2, 3, 1}},
		{"2", args{[][]int{{7, 10}, {7, 12}, {7, 5}, {7, 4}, {7, 2}}},
			[]int{4, 3, 2, 0, 1}},
		{"3", args{[][]int{{19, 13}, {16, 9}, {21, 10}, {32, 25}, {37, 4}, {49, 24}, {2, 15}, {38, 41}, {37, 34}, {33, 6}, {45, 4}, {18, 18}, {46, 39}, {12, 24}}},
			[]int{6, 1, 2, 9, 4, 10, 0, 11, 5, 13, 3, 8, 12, 7}},
		{"4", args{[][]int{{35, 36}, {11, 7}, {15, 47}, {34, 2}, {47, 19}, {16, 14}, {19, 8}, {7, 34}, {38, 15}, {16, 18}, {27, 22}, {7, 15}, {43, 2}, {10, 5}, {5, 4}, {3, 11}}},
			[]int{15, 14, 13, 1, 6, 3, 5, 12, 8, 11, 9, 4, 10, 7, 0, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getOrder(tt.args.tasks); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}