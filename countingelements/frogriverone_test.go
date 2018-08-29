package countingelements

import "testing"

func TestFrogRiverOne(t *testing.T) {
	type args struct {
		X int
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
				A: []int{1, 3, 1, 4, 2, 3, 5, 4},
				X: 4,
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FrogRiverOne(tt.args.X, tt.args.A); got != tt.want {
				t.Errorf("FrogRiverOne() = %v, want %v", got, tt.want)
			}
		})
	}
}
