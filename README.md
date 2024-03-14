# xlsxpict

[microsoft/pict](https://github.com/microsoft/pict)の結果をExcelに出力するツールです。

毎回、Excelに貼りつけるのが面倒なので作成しただけのツールです。

Windows以外では動きません。

## Build

```sh
$ task
task: [default] go build -o bin/xlsxpict.exe main.go
```

## Usage

実行ファイルは bin ディレクトリの下にあります。

予め model.txt に pict のモデルデータを入力しておいてください。

xlsxpict.exe を起動すれば同じフォルダに result.xlsx というファイルが出力されます。


## REFERENCES

- [microsoft/pict](https://github.com/microsoft/pict)
- [組み合わせテストの技法](https://gihyo.jp/dev/feature/01/sp-test/0001)
- [PICTの基本的な使い方](https://gihyo.jp/dev/feature/01/sp-test/0002)
- [PICTの機能説明と使用例 （前編）](https://gihyo.jp/dev/feature/01/sp-test/0003)
- [PICTの機能説明と使用例 （中編）](https://gihyo.jp/dev/feature/01/sp-test/0004)
- [PICTの機能説明と使用例 （後編）](https://gihyo.jp/dev/feature/01/sp-test/0005)
- [PICTをより使いやすくする](https://gihyo.jp/dev/feature/01/sp-test/0006)

