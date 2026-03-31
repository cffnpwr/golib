package path_test

import (
	"testing"

	"github.com/cffnpwr/golib/path"
)

func Test_IsPath(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "[positive] it is a path string when current directory",
			s:    ".",
			want: true,
		},
		{
			name: "[positive] it is a path string when parent directory",
			s:    "..",
			want: true,
		},
		{
			name: "[positive] it is a path string when relative path (./)",
			s:    "./foo/bar",
			want: true,
		},
		{
			name: "[positive] it is a path string when relative path (../)",
			s:    "../foo/bar",
			want: true,
		},
		{
			name: "[positive] it is a path string when simple filename",
			s:    "file.txt",
			want: true,
		},
		{
			name: "[positive] it is a path string when path with spaces",
			s:    "path with spaces/file",
			want: true,
		},
		{
			name: "[negative] it is not a path string when empty string",
			s:    "",
			want: false,
		},
		{
			name: "[negative] it is not a path string when contains NUL character",
			s:    "/path\x00/file",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := path.IsPath(tt.s)
			if got != tt.want {
				t.Errorf("IsPath(%q) = %v; want %v", tt.s, got, tt.want)
			}
		})
	}
}
