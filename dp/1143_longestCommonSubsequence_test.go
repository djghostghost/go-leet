package dp

import "testing"

func Test_longestCommonSubsequenc(t *testing.T) {
	type args struct {
		text1 string
		text2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{

		{
			name: "case1",
			args: args{text1: "abcde", text2: "ace"},
			want: 3,
		},
		{
			name: "case2",
			args: args{text1: "abc", text2: "abc"},
			want: 3,
		},
		{
			name: "case3",
			args: args{text1: "abc", text2: "def"},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonSubsequence(tt.args.text1, tt.args.text2); got != tt.want {
				t.Errorf("longestCommonSubsequenc() = %v, want %v", got, tt.want)
			}
		})
	}
}
