package arrays

import "testing"

func TestOddOccurence(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				A: []int{9, 3, 9, 3, 9, 7, 9},
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		if got := OddOccurence(tt.args.A); got != tt.want {
			t.Errorf("%q. OddOccurence() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
