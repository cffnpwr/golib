//go:build unix

package path_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/cffnpwr/golib/path"
)

func Test_IsInsideBase_Symlink(t *testing.T) {
	t.Parallel()

	// Set up temp directory tree:
	//   base/
	//     child/
	//     symlink_inside  -> child
	//     symlink_outside -> ../outside
	//   outside/
	base := t.TempDir()
	outside := t.TempDir()

	child := filepath.Join(base, "child")
	if err := os.MkdirAll(child, 0o755); err != nil {
		t.Fatal(err)
	}

	symlinkInside := filepath.Join(base, "symlink_inside")
	if err := os.Symlink(child, symlinkInside); err != nil {
		t.Fatal(err)
	}

	symlinkOutside := filepath.Join(base, "symlink_outside")
	if err := os.Symlink(outside, symlinkOutside); err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name   string
		base   string
		target string
		want   bool
	}{
		{
			name:   "[positive] it is inside base when symlink resolves to inside base",
			base:   base,
			target: symlinkInside,
			want:   true,
		},
		{
			name:   "[negative] it is not inside base when symlink resolves to outside base",
			base:   base,
			target: symlinkOutside,
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := path.IsInsideBase(tt.base, tt.target)
			if got != tt.want {
				t.Errorf("IsInsideBase(%q, %q) = %v; want %v", tt.base, tt.target, got, tt.want)
			}
		})
	}
}
