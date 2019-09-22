package main

import "testing"

func Test_reverseParentheses(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case1",
			args: args{s: "(abcd)"},
			want: "dcba",
		},
		{
			name: "case2",
			args: args{s: "(ed(et(oc))el)"},
			want: "leetcode",
		},
		{
			name: "case3",
			args: args{s: "a(bcdefghijkl(mno)p)q"},
			want: "apmnolkjihgfedcbq",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseParentheses(tt.args.s); got != tt.want {
				t.Errorf("reverseParentheses() = %v, want %v", got, tt.want)
			}
		})
	}
}
