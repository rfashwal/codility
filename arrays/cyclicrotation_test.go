package arrays

import (
	"reflect"
	"testing"
)

func TestCyclicRotation(t *testing.T) {
	type args struct {
		A []int
		K int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test 1",
			args: args{
				A: []int{3, 8, 9, 7, 6},
				K: 3,
			},
			want: []int{9, 7, 6, 3, 8},
		},
		{
			name: "Test 2",
			args: args{
				A: []int{1, 2, 3, 4},
				K: 4,
			},
			want: []int{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CyclicRotation(tt.args.A, tt.args.K); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CyclicRotation() = %v, want %v", got, tt.want)
			}
		})
	}
}
