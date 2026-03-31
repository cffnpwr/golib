package path

import (
	"io/fs"
	"path"
	"path/filepath"
	"strings"
)

// isInsideClean reports whether cleaned target path is inside cleaned base path.
// Both paths must already be cleaned (no symlinks, no ..).
func isInsideClean(base, target string) bool {
	if base == target {
		return true
	}
	return strings.HasPrefix(target, strings.TrimSuffix(base, "/")+"/")
}

// IsInsideBase reports whether target is inside base (including base itself),
// resolving symlinks before comparison.
func IsInsideBase(base, target string) bool {
	resolvedBase, err := filepath.EvalSymlinks(base)
	if err != nil {
		return false
	}
	resolvedTarget, err := filepath.EvalSymlinks(target)
	if err != nil {
		return false
	}

	rel, err := filepath.Rel(resolvedBase, resolvedTarget)
	if err != nil {
		return false
	}
	return !strings.HasPrefix(rel, "..")
}

// IsInsideBaseFS reports whether target is inside base (including base itself)
// within the given file system, resolving symlinks before comparison.
func IsInsideBaseFS(fsys fs.ReadLinkFS, base, target string) bool {
	resolvedBase, err := evalSymlinksFS(fsys, base)
	if err != nil {
		return false
	}
	resolvedTarget, err := evalSymlinksFS(fsys, target)
	if err != nil {
		return false
	}

	return isInsideClean(resolvedBase, resolvedTarget)
}

// evalSymlinksFS resolves symlinks in the given path within fsys.
func evalSymlinksFS(fsys fs.ReadLinkFS, p string) (string, error) {
	const maxLinks = 255
	p = path.Clean(p)
	for range maxLinks {
		info, err := fsys.Lstat(p)
		if err != nil {
			return "", err
		}
		if info.Mode()&fs.ModeSymlink == 0 {
			return p, nil
		}
		link, err := fsys.ReadLink(p)
		if err != nil {
			return "", err
		}
		if path.IsAbs(link) {
			p = link
		} else {
			p = path.Join(path.Dir(p), link)
		}
		p = path.Clean(p)
	}
	return "", &fs.PathError{Op: "evalSymlinks", Path: p, Err: fs.ErrInvalid}
}
