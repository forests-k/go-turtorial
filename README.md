# Go Turtorial

## 参考資料
[プログラミング言語Go完全入門](https://docs.google.com/presentation/d/1RVx8oeIMAWxbB7ZP2IcgZXnbZokjCmTUca-AbIpORGk/edit#slide=id.g4f417182ce_0_0)

## 各種チュートリアル
* [Getting started](https://go.dev/doc/tutorial/getting-started)
* [Create a module](https://go.dev/doc/tutorial/create-module)
* [Accessing a relational database](https://go.dev/doc/tutorial/database-access)
* [Developing a RESTful API with Go and Gin](https://go.dev/doc/tutorial/web-service-gin)
* [Getting started with generics](https://go.dev/doc/tutorial/generics)
* [Getting started with fuzzing](https://go.dev/doc/tutorial/fuzz)

## 開発環境構築

### DockerCompose構築

Dockerfile作成

```docker
# goバージョン
FROM golang:1.20.1-alpine
# アップデートとgitのインストール
RUN apk update && apk add git
# boiler-plateディレクトリの作成
RUN mkdir /go/src/app
# ワーキングディレクトリの設定
WORKDIR /go/src/app
# ホストのファイルをコンテナの作業ディレクトリに移行
ADD . /go/src/app
```

docker-compose.yml作成

```yaml
version: '3' # composeファイルのバージョン
services:
  turtorial: # サービス名
    container_name: turtorial
    build: # ビルドに使うDockerファイルのパス
      context: .
      dockerfile: ./build/Dockerfile
    volumes: # マウントディレクトリ
      - ./cmd:/go/src/app
    tty: true # コンテナの永続化
    environment:
      - TZ=Asia/Tokyo
```

以下のコマンドを実行

docker-compose up -d —build

cmdディレクトリを作成し、以下にmain.goファイルを作成する

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello world")
}
```

以下のコマンドを実行

```go
$ docker exec -it turtorial sh
/go/src/app # go run main.go
Hello world
```

### VisualStudioCode構築
VisualStudioCodeの拡張機能で提供されている[Go](https://marketplace.visualstudio.com/items?itemName=golang.go)をインストールする

## モジュール管理
Goにおけるモジュール管理はgo mod tidyを使用する。

[Qiita - go mod tidyの役割](https://qiita.com/wangqijiangjun/items/28037d06efe86ec8dd0f)
[Qiita - go mod tidy で不要なpackageを削除する](https://qiita.com/k-kurikuri/items/609141727320eb1a6d2b)

```go
/go/src/app # go mod
Go mod provides access to operations on modules.

Note that support for modules is built into all the go commands,
not just 'go mod'. For example, day-to-day adding, removing, upgrading,
and downgrading of dependencies should be done using 'go get'.
See 'go help modules' for an overview of module functionality.

Usage:

        go mod <command> [arguments]

The commands are:

        download    download modules to local cache
        edit        edit go.mod from tools or scripts
        graph       print module requirement graph
        init        initialize new module in current directory
        tidy        add missing and remove unused modules
        vendor      make vendored copy of dependencies
        verify      verify dependencies have expected content
        why         explain why packages or modules are needed

Use "go help mod <command>" for more information about a command.
```

## 初期化

初期化としてgo.modを作成する

```
/go/src/app # go mod init forests-k/turtorial
go: creating new go.mod: module forests-k/turtorial
go: to add module requirements and sums:
        go mod tidy
```

パッケージをインストールすると、go.sumファイルが作成される
```
/go/src/app # go get golang.org/x/tools/cmd/goimports
go: downloading golang.org/x/tools v0.6.0
go: downloading golang.org/x/sys v0.5.0
go: downloading golang.org/x/mod v0.8.0
go: added golang.org/x/mod v0.8.0
go: added golang.org/x/sys v0.5.0
```

標準的なライブラリとしては、以下がある
* fmt
  * 書式に関する処理
* net/http
  * HTTPサーバ系
* archive,compress
  * zipやgzipなど
* encoding
  * JSON、XML、CSVなど
* html/template
  * HTMLテンプレート
* os,path/filepath
  * ファイル操作

[Zenn - dockerでgo開発環境構築](https://zenn.dev/akakuro/articles/2426098256785b)

## 各ツール
* go build
  * ビルドコマンド
* go test
  * テストコード実行
* go doc
  * ドキュメント生成
* gofmt
  * ソールコードフォーマッタ
* go vet
  * コードチェッカ
* gopls
  * Language Server Protocol実装