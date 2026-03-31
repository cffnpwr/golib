package path_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/cffnpwr/golib/path"
)

func Test_IsInsideBase(t *testing.T) {
	t.Parallel()

	// Set up temp directory tree:
	//   base/
	//     child/
	//       file.txt
	//   outside/
	base := t.TempDir()
	outside := t.TempDir()

	child := filepath.Join(base, "child")
	if err := os.MkdirAll(child, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(child, "file.txt"), []byte(""), 0o644); err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name   string
		base   string
		target string
		want   bool
	}{
		{
			name:   "[positive] it is inside base when target is a direct child of base",
			base:   base,
			target: child,
			want:   true,
		},
		{
			name:   "[positive] it is inside base when target is base itself",
			base:   base,
			target: base,
			want:   true,
		},
		{
			name:   "[positive] it is inside base when target is a file inside base",
			base:   base,
			target: filepath.Join(child, "file.txt"),
			want:   true,
		},
		{
			name:   "[negative] it is not inside base when target is outside base",
			base:   base,
			target: outside,
			want:   false,
		},
		{
			name:   "[negative] it is not inside base when path traversal via /../",
			base:   base,
			target: filepath.Join(base, "child", "..", "..", filepath.Base(outside)),
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
