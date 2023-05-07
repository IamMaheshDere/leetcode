package main

import (
	"reflect"
	"testing"
)

func Test_duplicateZeros(t *testing.T) {
	type args struct {
		inputArr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test case 1",
			args: args{
				inputArr: []int{1, 2, 3, 4},
			},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "test case 2",
			args: args{
				inputArr: []int{1, 2, 0, 3, 4},
			},
			want: []int{1, 2, 0, 0, 3},
		},
		{
			name: "test case 3",
			args: args{
				inputArr: []int{1, 2, 0, 0, 0},
			},
			want: []int{1, 2, 0, 0, 0},
		},
		{
			name: "test case 4",
			args: args{
				inputArr: []int{0, 0, 0, 0, 0},
			},
			want: []int{0, 0, 0, 0, 0},
		},
		{
			name: "test case 5",
			args: args{
				inputArr: []int{1, 2, 3, 4, 5},
			},
			want: []int{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := duplicateZeros(tt.args.inputArr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("duplicateZeros() = %v, want %v", got, tt.want)
			}
		})
	}
}
