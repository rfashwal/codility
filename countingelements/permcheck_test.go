package countingelements

import "testing"

func TestPremCheck(t *testing.T) {
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
				A: []int{4, 1, 3, 2},
			},
			want: 1,
		},
		{
			name: "Test 2",
			args: args{
				A: []int{1, 3, 4},
			},
			want: 0,
		},
		{
			name: "Test 3",
			args: args{
				A: []int{1, 3, 4, 2, 1},
			},
			want: 0,
		},
		{
			name: "Test 4",
			args: args{
				A: []int{4, 1, 3, 2, 5, 7, 6, 9, 8, 11, 10},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PremCheck(tt.args.A); got != tt.want {
				t.Errorf("PremCheck() = %v, want %v", got, tt.want)
			}
		})
	}
}
