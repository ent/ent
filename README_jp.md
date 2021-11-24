## ent - Goのエンティティーフレームワーク

[![Twitter](https://img.shields.io/twitter/url/https/twitter.com/entgo_io.svg?style=social&label=Follow%20%40entgo_io)](https://twitter.com/entgo_io)

[English](README.md) | [中文](README_zh.md) | [日本語](README_jp.md)

<img width="50%"
align="right"
style="display: block; margin:40px auto;"
src="https://s3.eu-central-1.amazonaws.com/entgo.io/assets/gopher_graph.png"/>

シンプルながらもパワフルなGoのエンティティフレームワークであり、大規模なデータモデルを持つアプリケーションを容易に構築・保守できるようにします。

- **Schema As Code(コードとしてのスキーマ)** - あらゆるデータベーススキーマをGoオブジェクトとしてモデル化します。
- **任意のグラフを簡単にトラバースできます** - クエリや集約の実行、任意のグラフ構造の走査を容易に実行できます。
- **100%静的に型付けされた明示的なAPI** - コード生成により、100%静的に型付けされた曖昧さのないAPIを提供します。
- **マルチストレージドライバ** - MySQL、PostgreSQL、SQLite、Gremlinをサポートしています。
- **拡張性** - Goテンプレートを使用して簡単に拡張、カスタマイズできます。

## クイックインストール
```console
go get -d entgo.io/ent/cmd/ent
```

[Go modules]を使ったインストールについては、[entgo.io website][entgo instal]をご覧ください。

## ドキュメントとサポート
entを開発・使用するためのドキュメントは、こちら。: https://entgo.io

議論やサポートについては、[Issueを立てる](https://github.com/ent/ent/issues/new/choose)か、gophers Slackの[チャンネル](https://gophers.slack.com/archives/C01FMSQDT53)に参加してください。

## entコミュニティへの参加
`ent`の構築は、コミュニティ全体の協力なしには実現できませんでした。 私たちは、この`ent`の貢献者をリストアップした[contributorsページ](doc/md/contributors.md)を管理しています。

`ent`に貢献するときは、まず[CONTRIBUTING](CONTRIBUTING.md)を参照してください。
もし、あなたの会社や製品で`ent`を利用している場合は、[ent usersページ](https://github.com/ent/ent/wiki/ent-users)に追記する形で、そのことをぜひ教えて下さい。

最新情報については、Twitter(<https://twitter.com/entgo_io>)をフォローしてください。



## プロジェクトについて
`ent`プロジェクトは、私たちが社内で使用しているエンティティフレームワークである`Ent`からインスピレーションを得ています。
entは、[Facebook Connectivity][fbc]チームの[a8m](https://github.com/a8m)と[alexsn](https://github.com/alexsn)が開発・保守しています。
本番環境では複数のチームやプロジェクトで使用されており、v1リリースまでのロードマップは[こちら](https://github.com/ent/ent/issues/46)に記載されています。
このプロジェクトの動機については[こちら](https://entgo.io/blog/2019/10/03/introducing-ent)をご覧ください。

## ライセンス
entは、[LICENSEファイル](LICENSE)にもある通り、Apache 2.0でライセンスされています。


[entgo instal]: https://entgo.io/docs/code-gen/#version-compatibility-between-entc-and-ent
[Go modules]: https://github.com/golang/go/wiki/Modules#quick-start
[fbc]: https://connectivity.fb.com
