package path_test

import (
	"io/fs"
	"testing"
	"testing/fstest"

	"github.com/cffnpwr/golib/path"
)

func Test_IsInsideBaseFS(t *testing.T) {
	t.Parallel()

	// File system layout:
	//   base/
	//     child/
	//       file.txt
	//     symlink_inside  -> child        (resolves inside base)
	//     symlink_outside -> ../outside   (resolves outside base)
	//   outside/
	//     file.txt
	fsys := fstest.MapFS{
		"base/child/file.txt": &fstest.MapFile{
			Mode: 0o644,
		},
		"base/symlink_inside": &fstest.MapFile{
			Data: []byte("child"),
			Mode: fs.ModeSymlink | 0o644,
		},
		"base/symlink_outside": &fstest.MapFile{
			Data: []byte("../outside"),
			Mode: fs.ModeSymlink | 0o644,
		},
		"outside/file.txt": &fstest.MapFile{
			Mode: 0o644,
		},
	}

	tests := []struct {
		name   string
		base   string
		target string
		want   bool
	}{
		{
			name:   "[positive] it is inside base when target is a direct child of base",
			base:   "base",
			target: "base/child",
			want:   true,
		},
		{
			name:   "[positive] it is inside base when target is base itself",
			base:   "base",
			target: "base",
			want:   true,
		},
		{
			name:   "[positive] it is inside base when target is a file inside base",
			base:   "base",
			target: "base/child/file.txt",
			want:   true,
		},
		{
			name:   "[positive] it is inside base when symlink resolves to inside base",
			base:   "base",
			target: "base/symlink_inside",
			want:   true,
		},
		{
			name:   "[negative] it is not inside base when target is outside base",
			base:   "base",
			target: "outside",
			want:   false,
		},
		{
			name:   "[negative] it is not inside base when symlink resolves to outside base",
			base:   "base",
			target: "base/symlink_outside",
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := path.IsInsideBaseFS(fsys, tt.base, tt.target)
			if got != tt.want {
				t.Errorf("IsInsideBaseFS(%q, %q) = %v; want %v", tt.base, tt.target, got, tt.want)
			}
		})
	}
}
