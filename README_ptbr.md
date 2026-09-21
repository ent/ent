## ent - An Entity Framework For Go

[![Twitter](https://img.shields.io/twitter/url/https/twitter.com/entgo_io.svg?style=social&label=Follow%20%40entgo_io)](https://twitter.com/entgo_io)
[![Discord](https://img.shields.io/discord/885059418646003782?label=discord&logo=discord&style=flat-square&logoColor=white)](https://discord.gg/qZmPgTE6RX)

[English](README.md) | [中文](README_zh.md) | [日本語](README_jp.md) | [한국어](README_kr.md) | [Português Brasil](README_ptbr.md)

<img width="50%"
align="right"
style="display: block; margin:40px auto;"
src="https://s3.eu-central-1.amazonaws.com/entgo.io/assets/gopher_graph.png"/>

Estrutura de entidades simples, mas poderosa, para Go, que facilita a criação e a manutenção de aplicativos com grande modelos de dados.

- **Schema como código** - modele qualquer schema de banco de dados como objeto Go.
- **Percorra facilmente qualquer grafo** - execute consultas, agregações e percorra qualquer estrutura de grafo facilmente.
- **API estaticamente tipada e explícita** - API 100% estaticamente tipada e explícita usando geração de codigo.
- **Multiplos drivers de banco de dados** - suporta MySQL, MariaDB, TiDB, PostgreSQL, CockroachDB, SQLite and Gremlin.
- **Extensível** - simples de estender e customizar usando modelos em Go.

## Instalação rápida
```console
go install entgo.io/ent/cmd/ent@latest
```

Para uma instalação adequada usando [Go modules], visite [entgo.io website][entgo install].

## Documentação e suporte
A documentação para desenvolver e usar ent está disponível em: https://entgo.io

Para discussão e suporte, [open an issue](https://github.com/ent/ent/issues/new/choose) ou entre no nosso [channel](https://gophers.slack.com/archives/C01FMSQDT53) gophers Slack.

## Sobre o projeto
O projeto `ent` foi inspirado pelo Ent, um framework de entidades usado internamente pela Meta (Facebook). Ele foi criado por [a8m](https://github.com/a8m) e [alexsn](https://github.com/alexsn) da equipe de [Facebook Connectivity][fbc]. Atualmente, ele é desenvolvido e mantido pela equipe [Atlas](https://github.com/ariga/atlas) e a trilha para seu lançamento v1 é descrito [aqui](https://github.com/ent/ent/issues/46).

## Junte-se à comunidade ent
A construção do `ent` não teria sido possível sem o trabalho coletivo de toda a nossa comunidade. Mantemos uma [página de contribuidores](doc/md/contributors.md)
que lista os contribuidores que ajudou no `ent`. 

Para contribuir com `ent`, veja o arquivo de [CONTRIBUTING](CONTRIBUTING.md) para saber como começar.
Se sua empresa ou seu produto estiver usando `ent`, por favor nos avise adicionando-se à [página de usuários ent](https://github.com/ent/ent/wiki/ent-users).

Para atualizações, siga a gente no Twitter em https://twitter.com/entgo_io

## License
ent é licenciado sob Apache 2.0, conforme encontrado no [arquivo de Licença](LICENSE).

[entgo install]: https://entgo.io/docs/code-gen/#version-compatibility-between-entc-and-ent
[Go modules]: https://github.com/golang/go/wiki/Modules#quick-start
[fbc]: https://connectivity.fb.com
