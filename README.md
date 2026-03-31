# golib

[![GitHub License](https://img.shields.io/github/license/cffnpwr/golib?style=flat)](./LICENSE)

Common Go libraries for cffnpwr.

[日本語版のREADMEはこちら](./README-ja.md)

## Installation

```bash
go get github.com/cffnpwr/golib@latest
```

## Packages

### [path](./path)

Provides utilities for validating and inspecting file system paths and names
in a cross-platform manner.

| Function                                         | Description                                       |
| ------------------------------------------------ | ------------------------------------------------- |
| `IsPath(s string) bool`                          | Reports whether `s` is a valid file system path   |
| `IsValidName(s string) bool`                     | Reports whether `s` is a valid filename component |
| `IsInsideBase(base, target string) bool`         | Reports whether `target` is inside `base`         |
| `IsInsideBaseFS(fsys, base, target string) bool` | Same as above on `fs.ReadLinkFS`                  |

See the [path package README](./path/README.md) for details.

## License

[MIT License](./LICENSE)
