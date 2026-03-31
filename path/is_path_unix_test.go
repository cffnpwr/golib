//go:build unix

package path_test

import (
	"testing"

	"github.com/cffnpwr/golib/path"
)

func Test_IsPath_Unix(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "[positive] it is a path string when absolute path",
			s:    "/path/to/file",
			want: true,
		},
		{
			name: "[positive] it is a path string when root directory",
			s:    "/",
			want: true,
		},
		{
			name: "[positive] it is a path string when consecutive slashes",
			s:    "///foo/bar",
			want: true,
		},
		{
			name: "[positive] it is a path string when contains angle brackets",
			s:    "/path<with>file",
			want: true,
		},
		{
			name: "[positive] it is a path string when contains colon",
			s:    "/path:file",
			want: true,
		},
		{
			name: "[positive] it is a path string when contains pipe",
			s:    "/path|file",
			want: true,
		},
		{
			name: "[positive] it is a path string when contains question mark",
			s:    "/path?file",
			want: true,
		},
		{
			name: "[positive] it is a path string when contains asterisk",
			s:    "/path*file",
			want: true,
		},
		{
			name: "[positive] it is a path string when contains double quote",
			s:    "/path\"file",
			want: true,
		},
		{
			name: "[positive] it is a path string when contains backslash",
			s:    "/path\\file",
			want: true,
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
