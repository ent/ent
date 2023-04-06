## ent - 一个强大的Go语言实体框架

[English](README.md) | [中文](README_zh.md) | [日本語](README_jp.md)

<img width="50%"
align="right"
style="display: block; margin:40px auto;"
src="https://s3.eu-central-1.amazonaws.com/entgo.io/assets/gopher_graph.png"/>

ent是一个简单而又功能强大的Go语言实体框架，ent易于构建和维护应用程序与大数据模型。

- **图就是代码** - 将任何数据库表建模为Go对象。
- **轻松地遍历任何图形** - 可以轻松地运行查询、聚合和遍历任何图形结构。
- **静态类型和显式API** - 使用代码生成静态类型和显式API，查询数据更加便捷。
- **多存储驱动程序** - 支持MySQL, PostgreSQL, SQLite 和 Gremlin。
- **可扩展** - 简单地扩展和使用Go模板自定义。

## 快速安装
```console
go install entgo.io/ent/cmd/ent@latest
```

请访问[entgo.io website][entgo instal]以使用[Go modules]进行正确安装。

## 文档和支持
开发和使用ent的文档请参照： https://entgo.io

如要讨论问题和支持, [创建一个issue](https://github.com/ent/ent/issues/new/choose) 或者加入我们的Gopher Slack(Slack软件,类似于论坛)[讨论组](https://gophers.slack.com/archives/C01FMSQDT53)

## 加入 ent 社区
如果你想为`ent`做出贡献, [贡献代码](CONTRIBUTING.md) 中写了如何做出自己的贡献
如果你的公司或者产品在使用`ent`，请让我们知道你已经加入 [ent 用户](https://github.com/ent/ent/wiki/ent-users)

## 关于项目
`ent` 项目灵感来自于Ent，Ent是一个facebook内部使用的一个实体框架项目。 它由 [Facebook Connectivity][fbc] 团队通过 [a8m](https://github.com/a8m) 和 [alexsn](https://github.com/alexsn) 开发和维护
, 它被生产中的多个团队和项目使用。它的v1版本的路线图为 [版本的路线图](https://github.com/ent/ent/issues/46).
关于项目更多的信息 [ent介绍](https://entgo.io/blog/2019/10/03/introducing-ent)。

## 声明
ent使用Apache 2.0协议授权，可以在[LICENSE文件](LICENSE)中找到。

[entgo instal]: https://entgo.io/docs/code-gen/#version-compatibility-between-entc-and-ent
[Go modules]: https://github.com/golang/go/wiki/Modules#quick-start
[fbc]: https://connectivity.fb.com
