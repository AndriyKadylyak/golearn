package main

import "testing"

func Test_sumMaxMin(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{
			name: "test case 1",
			args: args{
				nums: []int{1, 2, 3, 4, 5},
			},
			want:  14,
			want1: 10,
		},
		{
			name: "test case 2",
			args: args{
				nums: []int{0, 2, 3, 4, 5},
			},
			want:  14,
			want1: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := sumMaxMin(tt.args.nums)
			if got != tt.want {
				t.Errorf("sumMaxMin() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("sumMaxMin() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
