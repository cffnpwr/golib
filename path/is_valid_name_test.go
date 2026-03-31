package path_test

import (
	"testing"

	"github.com/cffnpwr/golib/path"
)

func Test_IsValidName(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "[positive] it is a valid name when simple filename",
			s:    "file.txt",
			want: true,
		},
		{
			name: "[positive] it is a valid name when filename without extension",
			s:    "file",
			want: true,
		},
		{
			name: "[positive] it is a valid name when dot file",
			s:    ".gitignore",
			want: true,
		},
		{
			name: "[positive] it is a valid name when single dot",
			s:    ".",
			want: true,
		},
		{
			name: "[positive] it is a valid name when double dot",
			s:    "..",
			want: true,
		},
		{
			name: "[negative] it is not a valid name when empty string",
			s:    "",
			want: false,
		},
		{
			name: "[negative] it is not a valid name when contains NUL character",
			s:    "file\x00name",
			want: false,
		},
		{
			name: "[negative] it is not a valid name when contains forward slash",
			s:    "path/file",
			want: false,
		},
		{
			name: "[negative] it is not a valid name when contains backslash",
			s:    `path\file`,
			want: false,
		},
		{
			name: "[negative] it is not a valid name when exceeds 255 bytes",
			s:    string(make([]byte, 256)),
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := path.IsValidName(tt.s)
			if got != tt.want {
				t.Errorf("IsValidName(%q) = %v; want %v", tt.s, got, tt.want)
			}
		})
	}
}
