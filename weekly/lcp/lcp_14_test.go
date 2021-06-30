package lcp

import (
	"reflect"
	"testing"
)

func TestPrimeNums(t *testing.T) {
	type args struct {
		max int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{17}, []int{2, 3, 5, 7, 11, 13, 17}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PrimeNums(tt.args.max); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PrimeNums() = %v, want %v", got, tt.want)
			}
		})
	}
}
