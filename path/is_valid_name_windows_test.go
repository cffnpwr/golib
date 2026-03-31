//go:build windows

package path_test

import (
	"testing"

	"github.com/cffnpwr/golib/path"
)

func Test_IsValidName_Windows(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "[negative] it is not a valid name when contains <",
			s:    "file<name",
			want: false,
		},
		{
			name: "[negative] it is not a valid name when contains >",
			s:    "file>name",
			want: false,
		},
		{
			name: "[negative] it is not a valid name when contains :",
			s:    "file:name",
			want: false,
		},
		{
			name: "[negative] it is not a valid name when contains \"",
			s:    `file"name`,
			want: false,
		},
		{
			name: "[negative] it is not a valid name when contains |",
			s:    "file|name",
			want: false,
		},
		{
			name: "[negative] it is not a valid name when contains ?",
			s:    "file?name",
			want: false,
		},
		{
			name: "[negative] it is not a valid name when contains *",
			s:    "file*name",
			want: false,
		},
		{
			name: "[negative] it is not a valid name when contains control character 0x01",
			s:    "file\x01name",
			want: false,
		},
		{
			name: "[negative] it is not a valid name when contains control character 0x1F",
			s:    "file\x1fname",
			want: false,
		},
		{
			name: "[negative] it is not a valid name when reserved name CON",
			s:    "CON",
			want: false,
		},
		{
			name: "[negative] it is not a valid name when reserved name PRN",
			s:    "PRN",
			want: false,
		},
		{
			name: "[negative] it is not a valid name when reserved name AUX",
			s:    "AUX",
			want: false,
		},
		{
			name: "[negative] it is not a valid name when reserved name NUL",
			s:    "NUL",
			want: false,
		},
		{
			name: "[negative] it is not a valid name when reserved name COM1",
			s:    "COM1",
			want: false,
		},
		{
			name: "[negative] it is not a valid name when reserved name COM9",
			s:    "COM9",
			want: false,
		},
		{
			name: "[negative] it is not a valid name when reserved name LPT1",
			s:    "LPT1",
			want: false,
		},
		{
			name: "[negative] it is not a valid name when reserved name LPT9",
			s:    "LPT9",
			want: false,
		},
		{
			name: "[negative] it is not a valid name when reserved name with extension CON.txt",
			s:    "CON.txt",
			want: false,
		},
		{
			name: "[negative] it is not a valid name when reserved name lowercase con",
			s:    "con",
			want: false,
		},
		{
			name: "[negative] it is not a valid name when reserved name mixed case Con.txt",
			s:    "Con.txt",
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
