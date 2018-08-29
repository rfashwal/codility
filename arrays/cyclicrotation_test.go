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
<<<<<<< HEAD
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
=======
				K: 6,
			},
			want: []int{6, 3, 8, 9, 7},
		},
	}
	for _, tt := range tests {
		if got := CyclicRotation(tt.args.A, tt.args.K); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. CyclicRotation() = %v, want %v", tt.name, got, tt.want)
		}
>>>>>>> 11cc22cec302a686b6a64e41583deb7ddce7e95c
	}
}
