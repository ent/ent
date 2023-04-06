---
title: "Announcing v0.10: Ent gets a brand-new migration engine"
author: Ariel Mashraki
authorURL: https://github.com/a8m
authorImageURL: https://avatars0.githubusercontent.com/u/7413593
authorTwitter: arielmashraki
---
Dear community, 

I'm very happy to announce the release of the next version of Ent: v0.10. It has been 
almost six months since v0.9.1, so naturally there's a ton of new stuff in this release.
Still, I wanted to take the time to discuss one major improvement we have been working
on for the past few months: a brand-new migration engine. 

### Enter: [Atlas](https://github.com/ariga/atlas)

<img src="https://atlasgo.io/uploads/gopher.svg" width="200" />

Ent's current migration engine is great, and it does some pretty neat stuff which our
community has been using in production for years now, but as time went on issues
which we could not resolve with the existing architecture started piling up. In addition,
we feel that existing database migration frameworks leave much to be desired. We have 
learned so much as an industry about safely managing changes to production systems in 
the past decade with principles such as Infrastructure-as-Code and declarative configuration
management, that simply did not exist when most of these projects were conceived. 

Seeing that these problems were fairly generic and relevant to application regardless of the framework
or programming language it was written in, we saw the opportunity to fix them as common
infrastructure that any project could use. For this reason, instead of just rewriting
Ent's migration engine, we decided to extract the solution to a new open-source project, 
[Atlas](https://atlasgo.io) ([GitHub](https://ariga.io/atlas)).

Atlas is distributed as a CLI tool that uses a new [DDL](https://atlasgo.io/ddl/intro) based
on HCL (similar to Terraform), but can also be used as a [Go package](https://pkg.go.dev/ariga.io/atlas).
Just as Ent, Atlas is licensed under the [Apache License 2.0](https://github.com/ariga/atlas/blob/master/LICENSE).

Finally, after much work and testing, the Atlas integration for Ent is finally ready to use.  This is
great news to many of our users who opened issues (such as [#1652](https://github.com/ent/ent/issues/1652), 
[#1631](https://github.com/ent/ent/issues/1631), [#1625](https://github.com/ent/ent/issues/1625), 
[#1546](https://github.com/ent/ent/issues/1546) and [#1845](https://github.com/ent/ent/issues/1845)) 
that could not be well addressed  using the existing migration system, but are now resolved using the Atlas engine.

As with any substantial change, using Atlas as the migration engine for your project is currently opt-in.
In the near future, we will switch to an opt-out mode, and finally deprecate the existing engine. 
Naturally, this transition will be made slowly, and we will progress as we get positive indications 
from the community.

### Getting started with Atlas migrations for Ent

First, upgrade to the latest version of Ent:

```shell
go get entgo.io/ent@v0.10.0
```

Next, in order to execute a migration with the Atlas engine, use the `WithAtlas(true)` option.

```go {17}
package main
import (
    "context"
    "log"
    "<project>/ent"
    "<project>/ent/migrate"
    "entgo.io/ent/dialect/sql/schema"
)
func main() {
    client, err := ent.Open("mysql", "root:pass@tcp(localhost:3306)/test")
    if err != nil {
        log.Fatalf("failed connecting to mysql: %v", err)
    }
    defer client.Close()
    ctx := context.Background()
    // Run migration.
    err = client.Schema.Create(ctx, schema.WithAtlas(true))
    if err != nil {
        log.Fatalf("failed creating schema resources: %v", err)
    }
}
```
And that's it! 

One of the great improvements of the Atlas engine over the existing Ent code,
is it's layered structure, that cleanly separates between ***inspection*** (understanding 
the current state of a database), ***diffing*** (calculating the difference between the 
current and desired state), ***planning*** (calculating a concrete plan for remediating
the diff), and ***applying***. This diagram demonstrates the way Ent uses Atlas:

![atlas-migration-process](https://entgo.io/images/assets/migrate-atlas-process.png)

In addition to the standard options (e.g. `WithDropColumn`, 
`WithGlobalUniqueID`), the Atlas integration provides additional options for 
hooking into schema migration steps.

Here are two examples that show how to hook into the Atlas `Diff` and `Apply` steps.

```go
package main
import (
    "context"
    "log"
    "<project>/ent"
    "<project>/ent/migrate"
	"ariga.io/atlas/sql/migrate"
	atlas "ariga.io/atlas/sql/schema"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
)
func main() {
    client, err := ent.Open("mysql", "root:pass@tcp(localhost:3306)/test")
    if err != nil {
        log.Fatalf("failed connecting to mysql: %v", err)
    }
    defer client.Close()
    ctx := context.Background()
    // Run migration.
    err := 	client.Schema.Create(
		ctx,
		// Hook into Atlas Diff process.
		schema.WithDiffHook(func(next schema.Differ) schema.Differ {
			return schema.DiffFunc(func(current, desired *atlas.Schema) ([]atlas.Change, error) {
				// Before calculating changes.
				changes, err := next.Diff(current, desired)
				if err != nil {
					return nil, err
				}
				// After diff, you can filter
				// changes or return new ones.
				return changes, nil
			})
		}),
		// Hook into Atlas Apply process.
		schema.WithApplyHook(func(next schema.Applier) schema.Applier {
			return schema.ApplyFunc(func(ctx context.Context, conn dialect.ExecQuerier, plan *migrate.Plan) error {
				// Example to hook into the apply process, or implement
				// a custom applier. For example, write to a file.
				//
				//	for _, c := range plan.Changes {
				//		fmt.Printf("%s: %s", c.Comment, c.Cmd)
				//		if err := conn.Exec(ctx, c.Cmd, c.Args, nil); err != nil {
				//			return err
				//		}
				//	}
				//
				return next.Apply(ctx, conn, plan)
			})
		}),
	)
    if err != nil {
        log.Fatalf("failed creating schema resources: %v", err)
    }
}
```

### What's next: v0.11

I know we took a while to get this release out the door, but the next one is right around
the corner. Here's what's in store for v0.11:

* [Add support for edge/relation schemas](https://github.com/ent/ent/issues/1949) - supporting attaching metadata fields to relations. 
* Reimplementing the GraphQL integration to be fully compatible with the Relay spec. 
  Supporting generating GraphQL assets (schemas or full servers) from Ent schemas.
* Adding support for "Migration Authoring": the Atlas libraries have infrastructure for creating "versioned" 
  migration directories, as is commonly used in many migration frameworks (such as Flyway, Liquibase, go-migrate, etc.).
  Many users have built solutions for integrating with these kinds of systems, and we plan to use Atlas to provide solid 
  infrastructure for these flows. 
* Query hooks (interceptors) - currently hooks are only supported for [Mutations](https://entgo.io/docs/hooks/#hooks).
  Many users have requested adding support for read operations as well.  
* Polymorphic edges - The issue about adding support for polymorphism has been [open for over a year](https://github.com/ent/ent/issues/1048).
  With Go Generic Types support landing in 1.18, we want to re-open the discussion about a possible implementation using
  them.

### Wrapping up

Aside from the exciting announcement about the new migration engine, this release is huge
in size and contents, featuring [199 commits from 42 unique contributors](https://github.com/ent/ent/releases/tag/v0.10.0). Ent is a community
effort and keeps getting better every day thanks to all of you. So here's huge thanks and infinite
kudos to everyone who took part in this release (alphabetically sorted):

[attackordie](https://github.com/attackordie),
[bbkane](https://github.com/bbkane),
[bodokaiser](https://github.com/bodokaiser),
[cjraa](https://github.com/cjraa),
[dakimura](https://github.com/dakimura),
[dependabot](https://github.com/dependabot),
[EndlessIdea](https://github.com/EndlessIdea),
[ernado](https://github.com/ernado),
[evanlurvey](https://github.com/evanlurvey),
[freb](https://github.com/freb),
[genevieve](https://github.com/genevieve),
[giautm](https://github.com/giautm),
[grevych](https://github.com/grevych),
[hedwigz](https://github.com/hedwigz),
[heliumbrain](https://github.com/heliumbrain),
[hilakashai](https://github.com/hilakashai),
[HurSungYun](https://github.com/HurSungYun),
[idc77](https://github.com/idc77),
[isoppp](https://github.com/isoppp),
[JeremyV2014](https://github.com/JeremyV2014),
[Laconty](https://github.com/Laconty),
[lenuse](https://github.com/lenuse),
[masseelch](https://github.com/masseelch),
[mattn](https://github.com/mattn),
[mookjp](https://github.com/mookjp),
[msal4](https://github.com/msal4),
[naormatania](https://github.com/naormatania),
[odeke-em](https://github.com/odeke-em),
[peanut-cc](https://github.com/peanut-cc),
[posener](https://github.com/posener),
[RiskyFeryansyahP](https://github.com/RiskyFeryansyahP),
[rotemtam](https://github.com/rotemtam),
[s-takehana](https://github.com/s-takehana),
[sadmansakib](https://github.com/sadmansakib),
[sashamelentyev](https://github.com/sashamelentyev),
[seiichi1101](https://github.com/seiichi1101),
[sivchari](https://github.com/sivchari),
[storyicon](https://github.com/storyicon),
[tarrencev](https://github.com/tarrencev),
[ThinkontrolSY](https://github.com/ThinkontrolSY),
[timoha](https://github.com/timoha),
[vecpeng](https://github.com/vecpeng),
[yonidavidson](https://github.com/yonidavidson), and
[zeevmoney](https://github.com/zeevmoney).

Best,
Ariel

:::note For more Ent news and updates:

- Subscribe to our [Newsletter](https://entgo.substack.com/)
- Follow us on [Twitter](https://twitter.com/entgo_io)
- Join us on #ent on the [Gophers Slack](https://entgo.io/docs/slack)
- Join us on the [Ent Discord Server](https://discord.gg/qZmPgTE6RX)

:::