# path

[![GitHub License](https://img.shields.io/github/license/cffnpwr/golib?style=flat)](../LICENSE)

Cross-platform utilities for validating and inspecting file system paths and names.

[日本語版のREADMEはこちら](./README-ja.md)

## Functions

### `IsPath(s string) bool`

Reports whether `s` is a valid file system path on the current OS.

Returns `false` if `s` is empty or contains a null byte (`\x00`).
On Windows, additional constraints are checked (control characters,
reserved characters such as `<>:"|?*`).

```go
path.IsPath("/usr/local/bin") // true
path.IsPath("")               // false
path.IsPath("foo\x00bar")     // false
```

### `IsValidName(s string) bool`

Reports whether `s` is a valid single filename component
(no path separators allowed).

Returns `false` if `s` is empty, longer than 255 bytes, contains a null
byte, or contains `/` or `\`.
On Windows, control characters, reserved characters, and reserved device
names (e.g. `CON`, `NUL`, `COM1`) are also rejected.

```go
path.IsValidName("file.txt") // true
path.IsValidName("a/b")      // false
path.IsValidName("")         // false
```

### `IsInsideBase(base, target string) bool`

Reports whether `target` is inside `base` (including `base` itself),
resolving symlinks before comparison.

Returns `false` if either path cannot be resolved.

```go
path.IsInsideBase("/var/www", "/var/www/html/index.html") // true
path.IsInsideBase("/var/www", "/etc/passwd")              // false
```

### `IsInsideBaseFS(fsys fs.ReadLinkFS, base, target string) bool`

Same as `IsInsideBase` but operates on a provided `fs.ReadLinkFS`
instead of the real file system.

```go
path.IsInsideBaseFS(fsys, "static", "static/css/style.css") // true
path.IsInsideBaseFS(fsys, "static", "templates/index.html") // false
```

## License

[MIT License](../LICENSE)
