## ent - Go için Bir Varlık Çerçevesi

[![Twitter](https://img.shields.io/twitter/url/https/twitter.com/entgo_io.svg?style=social&label=Follow%20%40entgo_io)](https://twitter.com/entgo_io)
[![Discord](https://img.shields.io/discord/885059418646003782?label=discord&logo=discord&style=flat-square&logoColor=white)](https://discord.gg/qZmPgTE6RX)

[English](README.md) | [中文](README_zh.md) | [日本語](README_jp.md)

<img width="50%"
align="right"
style="display: block; margin:40px auto;"
src="https://s3.eu-central-1.amazonaws.com/entgo.io/assets/gopher_graph.png"/>

Go için basit ama güçlü varlık çerçevesi, bu da büyük veri modelleri ile çalışırken uygulamaların oluşturulmasını ve bakımını kolaylaştırır.

- **Kod ile şema oluşturma** - Herhangi bir veritabanı şemasını Go nesneleri kullanarak modelleme yapın.
- **Farklı Bir Grafiğe Kolayca Geçin** - Sorguları, toplamaları çalıştırın ve herhangi bir grafik yapısına kolayca geçin.
- **Statik Tip ve Açık API** - Kod oluşturma kullanılarak %100 statik olarak yazılan ve açık API.
- **Çoklu Depolama Sürücüsü** - MySQL, MariaDB, TiDB, PostgreSQL, CockroachDB, SQLite ve Gremlin desteklenir.
- **Genişletilebilir** - Go şablonlarını kullanarak genişletmek ve özelleştirmek kolaydır.

## Hızlı Kurulum
```console
go install entgo.io/ent/cmd/ent@latest
```

[Go modules] kullanarak, doğru kurulumu yapmak için [entgo.io websiteni](https://entgo.io/docs/code-gen/#version-compatibility-between-entc-and-ent) ziyaret ediniz.

## Belgeler ve Destek
Ent geliştirme ve kullanma belgeleri şu adreste mevcuttur: https://entgo.io

Tartışma ve destek için [bir konu açın](https://github.com/ent/ent/issues/new/choose) veya gophers Slack'te [kanalımıza](https://gophers.slack.com/archives/C01FMSQDT53) katılın.

## Ent Topluluğuna Katılın
Tüm topluluğumuzun ortak çalışması olmadan inşa `ent'i` etmek mümkün olmazdı. `ent'e` katkıda bulunanları listeleyen bir [katkıda bulunanlar sayfası](doc/md/contributors.md) tutuyoruz.

`ent'e` katkıda bulunmak için ve nasıl başlayacağınızı öğrenmek için [CONTRIBUTING](CONTRIBUTING.md) dosyasına bakın. 
Firmanızda veya projenizde kullanıyorsanız lütfen kendinizi `ent` [kullanıcılar sayfasına](https://github.com/ent/ent/wiki/ent-users) ekleyerek bize bildirin

Güncellemeler ve bilgilendirmeler için bizi [Twitter](https://twitter.com/entgo_io) adresinden takip edin: https://twitter.com/entgo_io

## Proje hakkında
Proje `ent`, dahili olarak kullandığımız bir varlık çerçevesi olan Ent'ten ilham aldı.
Facebook Connectivity ekibinden [a8m](https://github.com/a8m) ve [alexsn](https://github.com/alexsn) tarafından geliştirilmiş ve bakımı yapılmaktadır. 
Üretimde birden fazla ekip ve proje tarafından kullanılır ve v1 sürümünün yol haritası [burada](https://github.com/ent/ent/issues/46) açıklanmıştır.
Projenin motivasyonu hakkında daha fazla bilgiyi [buradan](https://entgo.io/blog/2019/10/03/introducing-ent) okuyun.

## Lisans
`ent`, [Lisans dosyasında](LICENSE) bulunan Apache 2.0 altında lisanslanmıştır.

[entgo yükleme]: https://entgo.io/docs/code-gen/#version-compatibility-between-entc-and-ent
[Go modules]: https://github.com/golang/go/wiki/Modules#quick-start
[fbc]: https://connectivity.fb.com
