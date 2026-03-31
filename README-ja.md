# golib

[![GitHub License](https://img.shields.io/github/license/cffnpwr/golib?style=flat)](./LICENSE)

Common Go libraries for cffnpwr.

[README.md for English is available here](./README.md)

## Installation

```bash
go get github.com/cffnpwr/golib@latest
```

## Packages

### [path](./path)

ファイルシステムのパスおよびファイル名をクロスプラットフォームで検証・検査する
ユーティリティを提供します。

| 関数                                             | 説明                                              |
| ------------------------------------------------ | ------------------------------------------------- |
| `IsPath(s string) bool`                          | `s`が有効なファイルシステムパスかどうかを返す     |
| `IsValidName(s string) bool`                     | `s`が有効なファイル名コンポーネントかどうかを返す |
| `IsInsideBase(base, target string) bool`         | `target`が`base`の内側にあるかどうかを返す        |
| `IsInsideBaseFS(fsys, base, target string) bool` | `fs.ReadLinkFS`上で同様の判定を行う               |

詳細は [path パッケージの README](./path/README-ja.md) を参照してください。

## License

[MIT License](./LICENSE)
