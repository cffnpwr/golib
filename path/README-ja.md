# path

[![GitHub License](https://img.shields.io/github/license/cffnpwr/golib?style=flat)](../LICENSE)

ファイルシステムのパスおよびファイル名をクロスプラットフォームで検証・検査する
ユーティリティパッケージです。

[README.md for English is available here](./README.md)

## Functions

### `IsPath(s string) bool`

`s`が現在のOS上で有効なファイルシステムパスかどうかを返します。

`s`が空文字またはヌルバイト (`\x00`) を含む場合は`false`を返します。
Windowsでは追加のチェック（制御文字、`<>:"|?*`などの予約文字）も行われます。

```go
path.IsPath("/usr/local/bin") // true
path.IsPath("")               // false
path.IsPath("foo\x00bar")     // false
```

### `IsValidName(s string) bool`

`s`が有効な単一ファイル名コンポーネント（パス区切り文字を含まない）かどうかを
返します。

`s`が空文字、255バイト超、ヌルバイトを含む、または`/`や`\`を含む場合は
`false`を返します。
Windowsでは制御文字、予約文字、予約デバイス名（`CON`、`NUL`、`COM1`など）も
拒否されます。

```go
path.IsValidName("file.txt") // true
path.IsValidName("a/b")      // false
path.IsValidName("")         // false
```

### `IsInsideBase(base, target string) bool`

シンボリックリンクを解決した上で、`target`が`base`の内側（`base`自身を含む）
にあるかどうかを返します。

どちらかのパスを解決できない場合は`false`を返します。

```go
path.IsInsideBase("/var/www", "/var/www/html/index.html") // true
path.IsInsideBase("/var/www", "/etc/passwd")              // false
```

### `IsInsideBaseFS(fsys fs.ReadLinkFS, base, target string) bool`

`IsInsideBase`と同様ですが、実際のファイルシステムの代わりに指定された
`fs.ReadLinkFS`上で動作します。

```go
path.IsInsideBaseFS(fsys, "static", "static/css/style.css") // true
path.IsInsideBaseFS(fsys, "static", "templates/index.html") // false
```

## License

[MIT License](../LICENSE)
