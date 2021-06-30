package pattern123

import "testing"

/*
[1,2,3,4]
[3,1,4,2]
[-1,3,2,0]
[2,4,5,1,3]
[1,3,2,4,5,6,7,8,9,10]
[1,4,0,-1,-2,-3,-1,-2]
[1,0,1,-4,-3]
[1,2,3,4,-4,-3,-5,-1]

false
true
true
true
true
true
true
false
*/
func Test_find132pattern(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{nums: []int{1, 2, 3, 4}}, false},
		{"2", args{nums: []int{3, 1, 4, 2}}, true},
		{"3", args{nums: []int{-1, 3, 2, 0}}, true},
		{"4", args{nums: []int{2, 4, 5, 1, 3}}, true},
		{"5", args{nums: []int{1, 3, 2, 4, 5, 6, 7, 8, 9, 10}}, true},
		{"6", args{nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 10, 9}}, true},
		{"7", args{nums: []int{1, 4, 0, -1, -2, -3, -1, -2}}, true},
		{"8", args{nums: []int{1, 0, 1, -4, -3}}, false},
		{"9", args{nums: []int{1, 2, 3, 4, -4, -3, -5, -1}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := find132pattern(tt.args.nums); got != tt.want {
				t.Errorf("find132pattern() = %v, want %v", got, tt.want)
			}
		})
	}
}
