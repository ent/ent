## ent - An Entity Framework For Go

[![Twitter](https://img.shields.io/twitter/url/https/twitter.com/entgo_io.svg?style=social&label=Follow%20%40entgo_io)](https://twitter.com/entgo_io)

[English](README.md) | [中文](README_zh.md) | [日本語](README_jp.md) | [한국어](README_kr.md)

<img width="50%"
align="right"
style="display: block; margin:40px auto;"
src="https://s3.eu-central-1.amazonaws.com/entgo.io/assets/gopher_graph.png"/>

단순하지만 강력한 Go용 엔티티 프레임워크이며, 대규모 데이터 모델을 가진 어플리케이션을 쉽게
빌드 및 유지보수 할 수 있습니다.

- **Schema As Code(스키마를 코드로)** - 모든 데이터베이스 스키마를 Go 객체로 모델링 합니다.
- **Easily Traverse Any Graph(모든 그래프를 쉽게 순회)** - 쿼리와 집계를 실행하고 모든 그래프 모형을 쉽게 순회합니다.
- **Statically Typed And Explicit API(정적 타입 및 명시적 API)** - 코드 생성으로 100% 정적 타입 및 명시적 API을 지원합니다.
- **Multi Storage Driver(멀티 스토리지 드라이버)** - MySQL, PostgreSQL, SQLite, Gremlin을 지원합니다.
- **Extendable(확장성)** - Go 템플릿을 사용하여 쉽게 확장하고 사용자 정의할 수 있습니다.

## 빠른 설치
```console
go get -d entgo.io/ent/cmd/ent
```

[Go modules]를 이용한 설치는 [entgo.io website][entgo instal]를 참조하세요.

## 문서 및 지원
ent를 사용 및 개발하기 위한 문서는 이쪽: https://entgo.io

토론 또는 지원을 위해서는 [open an issue](https://github.com/ent/ent/issues/new/choose) 또는 [channel](https://gophers.slack.com/archives/C01FMSQDT53) 슬랙 채널에 참여하세요.

## ent 커뮤니티에 참여하세요
`ent`의 구축은 커뮤니티 전체의 협력없이는 실현할 수 없었습니다. 
우리는 `ent`의 모든 기여자들을 [contributors page](doc/md/contributors.md)에서 관리하고있습니다.

`ent`에 기여하기 전, 먼저 [CONTRIBUTING](CONTRIBUTING.md)을 읽어주시길 바랍니다.
만약 당신의 회사 또는 당신의 제품으로 `ent`를 이용할 경우, [ent users page](https://github.com/ent/ent/wiki/ent-users) 에 우리가 알 수 있도록 추가해 주세요.

업데이트를 확인하려면 Twitter에서 https://twitter.com/entgo_io 를 Follow 하세요.


## 프로젝트에 관하여
`ent` 프로젝트는 우리가 내부적으로 사용하는 entity 프레임워크인 Ent에서 영감을 받았습니다.
`ent` 프로젝트는 [Facebook Connectivity][fbc] 의 팀인 [a8m](https://github.com/a8m) 과 [alexsn](https://github.com/alexsn) 에 의해서 개발 및 유지보수 되고 있습니다.
`ent` 프로젝트는 다양한 팀과 프로덕션 프로젝트에서 사용되고 있으며, v1 릴리즈 까지의 로드맵은 [여기](https://github.com/ent/ent/issues/46) 에 있습니다.
이 프로젝트의 대해서 동기에 대해서 더 [자세히](https://entgo.io/blog/2019/10/03/introducing-ent) 알아보세요.

## License
ent is licensed under Apache 2.0 as found in the [LICENSE file](LICENSE).


[entgo instal]: https://entgo.io/docs/code-gen/#version-compatibility-between-entc-and-ent
[Go modules]: https://github.com/golang/go/wiki/Modules#quick-start
[fbc]: https://connectivity.fb.com
