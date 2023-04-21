## ent - An Entity Framework For Go

[English](README.md) | [中文](README_zh.md) | [日本語](README_jp.md) | [한국어](README_kr.md)

<img width="50%"
align="right"
style="display: block; margin:40px auto;"
src="https://s3.eu-central-1.amazonaws.com/entgo.io/assets/gopher_graph.png"/>

간단하지만 강력한 Go용 엔터티 프레임워크로, 대규모 데이터 모델이 포함된 애플리케이션을 쉽게 만들고 유지할 수 있습니다.

-   **스키마를 코드로 관리** - 모든 데이터베이스 스키마와 모델을 Go Object로 구현 가능.
-   **어떤 그래프든 쉽게 탐색가능** - 쿼리실행, 집계, 그래프구조를 쉽게 탐색 가능.
-   **정적 타입 그리고 명시적인 API** - 100% 생성된 코드로, 정적타입과 명시적인 API를 제공.
-   **다양한 스토리지 드라이버** - MySQL, MariaDB, TiDB, PostgreSQL, CockroachDB, SQLite and Gremlin 를 지원
-   **확장성** - Go 템플릿을 이용하여 간단하게 확장, 커스터마이징 가능.

## 빠른 설치

```console
go install entgo.io/ent/cmd/ent@latest
```

[Go modules]을 사용하여 바르게 설치하려면, [entgo.io 웹페이지][entgo install]를 방문해주시길 바랍니다.

## 문서 및 지원

Ent 개발 및 사용에 관한 문서는 여기서 확인할 수 있습니다. : https://entgo.io

토론, 지원을 위해서 [open an issue](https://github.com/ent/ent/issues/new/choose)깃허브 이슈 또는 gophers Slack [채널](https://gophers.slack.com/archives/C01FMSQDT53)에 가입해주세요.

## ent 커뮤니티 가입

ent 커뮤니티의 공동작업이 없었다면, ent를 만들 수 없었을 것입니다. 우리는 기여한 사람들을 [contributors 페이지](doc/md/contributors.md)에 올리고 유지합니다.

ent에 기여하려면 [CONTRIBUTING](CONTRIBUTING.md)에서 시작 방법을 확인해보세요.
프로젝트나 회사에서 ent를 사용중이면, [ent 유저 페이지](https://github.com/ent/ent/wiki/ent-users)에 추가하여 알려주세요.

트위터계정을 팔로우하여 업데이트 소식을 확인하세요. https://twitter.com/entgo_io

## 프로젝트에 관하여

ent프로젝트는 내부적으로 사용하는 엔터티 프레임워크 "Ent"에서 영감을 받았습니다. 개발 및 유지보수는 [a8m](https://github.com/a8m) 및 [alexsn](https://github.com/alexsn)[Facebook Connectivity][fbc] 팀에서 담당합니다. 여러 팀이 프로덕션 환경에서 사용하고 있습니다. v1 릴리즈 로드맵에 대한 설명은 [여기](https://github.com/ent/ent/issues/46)를 클릭해주세요.
프로젝트 동기에 대해 더 궁금하시다면 [여기](https://entgo.io/blog/2019/10/03/introducing-ent)를 클릭해주세요.

## 라이센스

ent 라이센스는 Apache 2.0입니다. [LICENSE file](LICENSE)파일에서도 확인 가능합니다.

[entgo install]: https://entgo.io/docs/code-gen/#version-compatibility-between-entc-and-ent
[go modules]: https://github.com/golang/go/wiki/Modules#quick-start
[fbc]: https://connectivity.fb.com
