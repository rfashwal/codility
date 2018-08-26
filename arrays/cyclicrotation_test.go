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
				K: 6,
			},
			want: []int{6, 3, 8, 9, 7},
		},
	}
	for _, tt := range tests {
		if got := CyclicRotation(tt.args.A, tt.args.K); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. CyclicRotation() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
