package main

import "testing"

func TestUnpack(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			"letters and digits",
			args{text: "a4bc2d5e"},
			"aaaabccddddde",
		},
		{
			"only letters",
			args{text: "abcd"},
			"abcd",
		},
		{
			"digit first",
			args{text: "3abc"},
			"",
		},
		{
			"only digits",
			args{text: "45"},
			"",
		},
		{
			"two digits together",
			args{text: "aaa10b"},
			"",
		},
		{
			"zero",
			args{text: "aaa0b"},
			"aab",
		},
		{
			"empty string",
			args{text: ""},
			"",
		},
		{
			"ext1",
			args{text: `qwe\4\5`},
			"qwe45",
		},
		{
			"ext2",
			args{text: `qwe\45`},
			"qwe44444",
		},
		{
			"ext3",
			args{text: `qwe\\5`},
			`qwe\\\\\`,
		},
		{
			"ext4",
			args{text: `qw\ne`},
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Unpack(tt.args.text); got != tt.want {
				t.Errorf("Unpack() = %v, want %v", got, tt.want)
			}
		})
	}
}
