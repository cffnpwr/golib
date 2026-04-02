//go:build windows

package path

import (
	"testing"
)

func Test_IsPath_Windows(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "[positive] it is a path string when drive letter with backslash",
			s:    `C:\Users\test`,
			want: true,
		},
		{
			name: "[positive] it is a path string when drive letter with slash",
			s:    "C:/Users/test",
			want: true,
		},
		{
			name: "[positive] it is a path string when drive-relative path",
			s:    `C:foo\bar`,
			want: true,
		},
		{
			name: "[positive] it is a path string when lowercase drive letter",
			s:    `c:\users`,
			want: true,
		},
		{
			name: "[positive] it is a path string when UNC path with backslash",
			s:    `\\server\share\path`,
			want: true,
		},
		{
			name: "[positive] it is a path string when UNC path with slash",
			s:    "//server/share/path",
			want: true,
		},
		{
			name: "[positive] it is a path string when device path (dot)",
			s:    `\\.\COM1`,
			want: true,
		},
		{
			name: "[positive] it is a path string when device path (question)",
			s:    `\\?\very\long\path`,
			want: true,
		},
		{
			name: "[positive] it is a path string when device path with drive letter",
			s:    `\\?\C:\very\long\path`,
			want: true,
		},
		{
			name: "[positive] it is a path string when reserved name CON",
			s:    "CON",
			want: true,
		},
		{
			name: "[positive] it is a path string when path contains reserved name",
			s:    `C:\folder\CON`,
			want: true,
		},
		{
			name: "[positive] it is a path string when relative path with backslash",
			s:    `.\foo\bar`,
			want: true,
		},
		{
			name: "[negative] it is not a path string when contains forbidden character <",
			s:    `C:\path<file`,
			want: false,
		},
		{
			name: "[negative] it is not a path string when contains forbidden character >",
			s:    `C:\path>file`,
			want: false,
		},
		{
			name: "[negative] it is not a path string when contains forbidden character \"",
			s:    `C:\path"file`,
			want: false,
		},
		{
			name: "[negative] it is not a path string when contains forbidden character |",
			s:    `C:\path|file`,
			want: false,
		},
		{
			name: "[negative] it is not a path string when contains forbidden character ? in path",
			s:    `C:\path?\file`,
			want: false,
		},
		{
			name: "[negative] it is not a path string when contains forbidden character *",
			s:    `C:\path*file`,
			want: false,
		},
		{
			name: "[negative] it is not a path string when contains control character (0x01)",
			s:    "C:\\path\x01file",
			want: false,
		},
		{
			name: "[negative] it is not a path string when contains control character (0x1F)",
			s:    "C:\\path\x1Ffile",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := IsPath(tt.s)
			if got != tt.want {
				t.Errorf("IsPath(%q) = %v; want %v", tt.s, got, tt.want)
			}
		})
	}
}

func Test_extractDriveLetter(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		s        string
		want     string
		wantBool bool
	}{
		{
			name:     "[positive] it returns a drive letter when valid drive letter",
			s:        `C:\path\to\file`,
			want:     "C:",
			wantBool: true,
		},
		{
			name:     "[positive] it returns a drive letter when lowercase drive letter",
			s:        "d:/another/path",
			want:     "d:",
			wantBool: true,
		},
		{
			name:     "[negative] it returns no drive letter when empty string",
			s:        "",
			want:     "",
			wantBool: false,
		},
		{
			name:     "[negative] it returns no drive letter when no drive letter",
			s:        `relative\path`,
			want:     "",
			wantBool: false,
		},
		{
			name:     "[negative] it returns no drive letter when invalid drive letter (digit)",
			s:        `1:\invalid\drive`,
			want:     "",
			wantBool: false,
		},
		{
			name:     "[negative] it returns no drive letter when invalid drive letter (symbol)",
			s:        `%:\invalid\drive`,
			want:     "",
			wantBool: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, gotBool := extractDriveLetter(tt.s)
			if got != tt.want || gotBool != tt.wantBool {
				t.Errorf(
					"ExtractDriveLetter(%q) = (%q, %v); want (%q, %v)",
					tt.s,
					got,
					gotBool,
					tt.want,
					tt.wantBool,
				)
			}
		})
	}
}

func Test_extractDevicePathPrefix(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		s        string
		want     string
		wantBool bool
	}{
		{
			name:     "[positive] it returns a device path prefix when device path (dot)",
			s:        `\\.\COM1`,
			want:     `\\.\`,
			wantBool: true,
		},
		{
			name:     "[positive] it returns a device path prefix when device path (question)",
			s:        `\\?\C:\very\long\path`,
			want:     `\\?\`,
			wantBool: true,
		},
		{
			name:     "[negative] it returns no device path prefix when empty string",
			s:        "",
			want:     "",
			wantBool: false,
		},
		{
			name:     "[negative] it returns no device path prefix when no device path prefix",
			s:        `C:\normal\path`,
			want:     "",
			wantBool: false,
		},
		{
			name:     "[negative] it returns no device path prefix when relative path",
			s:        `relative\path`,
			want:     "",
			wantBool: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, gotBool := extractDevicePathPrefix(tt.s)
			if got != tt.want || gotBool != tt.wantBool {
				t.Errorf(
					"ExtractDevicePathPrefix(%q) = (%q, %v); want (%q, %v)",
					tt.s,
					got,
					gotBool,
					tt.want,
					tt.wantBool,
				)
			}
		})
	}
}
