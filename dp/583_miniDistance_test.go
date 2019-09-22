package dp

import "testing"

func Test_minDistance1(t *testing.T) {
	type args struct {
		word1 string
		word2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{word1: "abcde", word2: "ace"},
			want: 2,
		},
		{
			name: "case2",
			args: args{word1: "abc", word2: "abc"},
			want: 0,
		},
		{
			name: "case3",
			args: args{word1: "abc", word2: "def"},
			want: 6,
		},
		{
			name: "case4",
			args: args{word1: "sea", word2: "eat"},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDistance1(tt.args.word1, tt.args.word2); got != tt.want {
				t.Errorf("minDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
