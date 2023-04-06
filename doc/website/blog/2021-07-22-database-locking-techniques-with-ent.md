---
title: Database Locking Techniques with Ent
author: Rotem Tamir
authorURL: "https://github.com/rotemtam"
authorImageURL: "https://s.gravatar.com/avatar/36b3739951a27d2e37251867b7d44b1a?s=80"
authorTwitter: _rtam
---

Locks are one of the fundamental building blocks of any concurrent 
computer program. When many things are happening simultaneously, 
programmers reach out to locks to guarantee the mutual exclusion of 
concurrent access to a resource. Locks (and other mutual exclusion 
primitives) exist in many different layers of the stack from low-level
CPU instructions to application-level APIs (such as `sync.Mutex` in Go).

When working with relational databases, one of the common needs of 
application developers is the ability to acquire a lock on records.
Imagine an `inventory` table, listing items available for sale on 
an e-commerce website. This table might have a column named `state`
that could either be set to `available` or `purchased`.  avoid the
scenario where two users think they have successfully purchased the
same inventory item, the application must prevent two operations
from mutating the item from an available to a purchased state.

How can the application guarantee this? Having the server check 
if the desired item is `available` before setting it to `purchased` 
would not be good enough. Imagine a scenario where two users 
simultaneously try to purchase the same item. Two requests would
travel from their browsers to the application server and arrive 
roughly at the same time. Both would query the database for the 
item's state, and see the item is `available`. Seeing this, both
request handlers would issue an `UPDATE` query setting the state
to `purchased` and the `buyer_id` to the id of the requesting user.
Both queries will succeed, but the final state of the record will 
be that the user who issued the `UPDATE` query last will be
considered the buyer of the item.

Over the years, different techniques have evolved to allow developers
to write applications that provide these guarantees to users. Some 
of them involve explicit locking mechanisms provided by databases, 
while others rely on more general ACID properties of databases to
achieve mutual exclusion. In this post we will explore the 
implementation of two of these techniques using Ent.

### Optimistic Locking

Optimistic locking (sometimes also called Optimistic Concurrency 
Control) is a technique that can be used to achieve locking
behavior without explicitly acquiring a lock on any record.

On a high-level, this is how optimistic locking works:

- Each record is assigned a numeric version number. This value 
  must be monotonically increasing. Often Unix timestamps of the latest row update are used.
- A transaction reads a record, noting its version number from the 
  database.
- An `UPDATE` statement is issued to modify the record:
    - The statement must include a predicate requiring that the 
      version number has not changed from its previous value. For example: `WHERE id=<id> AND version=<previous version>`.
    - The statement must increase the version. Some applications
      will increase the current value by 1, and some will set it
      to the current timestamp.
- The database returns the amount of rows modified by
  the `UPDATE` statement. If the number is 0, this means someone
  else has modified the record between the time we read it, and 
  the time we wanted to update it. The transaction is considered 
  failed, rolled back and can be retried.

Optimistic locking is commonly used in "low contention" 
environments (situations where the likelihood of two transactions 
interfering with one another is relatively low) and where the 
locking logic can be trusted to happen in the application layer.
If there are writers to the database that we cannot ensure to
obey the required logic, this technique is rendered useless.

Let’s see how this technique can be employed using Ent. 

We start by defining our `ent.Schema` for a `User`. The user has an 
`online` boolean field to specify whether they are currently
online and an `int64` field for the current version number.

```go
// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Bool("online"),
		field.Int64("version").
			DefaultFunc(func() int64 {
				return time.Now().UnixNano()
			}).
			Comment("Unix time of when the latest update occurred")
	}
}
```

Next, let's implement a simple optimistically locked update to our
`online` field:

```go
func optimisticUpdate(tx *ent.Tx, prev *ent.User, online bool) error {
	// The next version number for the record must monotonically increase
	// using the current timestamp is a common technique to achieve this. 
	nextVer := time.Now().UnixNano()

	// We begin the update operation:
	n := tx.User.Update().

		// We limit our update to only work on the correct record and version:
		Where(user.ID(prev.ID), user.Version(prev.Version)).

		// We set the next version:
		SetVersion(nextVer).

		// We set the value we were passed by the user:
		SetOnline(online).
		SaveX(context.Background())

	// SaveX returns the number of affected records. If this value is 
	// different from 1 the record must have been changed by another
	// process.
	if n != 1 {
		return fmt.Errorf("update failed: user id=%d updated by another process", prev.ID)
	}
	return nil
}
```

Next, let's write a test to verify that if two processes try to
edit the same record, only one will succeed:

```go
func TestOCC(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	ctx := context.Background()

	// Create the user for the first time.
	orig := client.User.Create().SetOnline(true).SaveX(ctx)

	// Read another copy of the same user.
	userCopy := client.User.GetX(ctx, orig.ID)

	// Open a new transaction:
	tx, err := client.Tx(ctx)
	if err != nil {
		log.Fatalf("failed creating transaction: %v", err)
	}

	// Try to update the record once. This should succeed.
	if err := optimisticUpdate(tx, userCopy, false); err != nil {
		tx.Rollback()
		log.Fatal("unexpected failure:", err)
	}

	// Try to update the record a second time. This should fail.
	err = optimisticUpdate(tx, orig, false)
	if err == nil {
		log.Fatal("expected second update to fail")
	}
	fmt.Println(err)
}
```

Running our test:

```go
=== RUN   TestOCC
update failed: user id=1 updated by another process
--- PASS: Test (0.00s)
```

Great! Using optimistic locking we can prevent two processes from
stepping on each other's toes!

### Pessimistic Locking

As we've mentioned above, optimistic locking isn't always 
appropriate. For use cases where we prefer to delegate the 
responsibility for maintaining the integrity of the lock to
the databases, some database engines (such as MySQL, Postgres,
and MariaDB, but not SQLite) offer pessimistic locking
capabilities.  These databases support a modifier to `SELECT` 
statements that is called `SELECT ... FOR UPDATE`. The MySQL 
documentation [explains](https://dev.mysql.com/doc/refman/8.0/en/innodb-locking-reads.html):

> A SELECT ... FOR UPDATE reads the latest available data, setting
> exclusive locks on each row it reads. Thus, it sets the same locks 
> a searched SQL UPDATE would set on the rows.

Alternatively, users can use `SELECT ... FOR SHARE` statements, as 
explained by the docs, `SELECT ... FOR SHARE`:

> Sets a shared mode lock on any rows that are read. Other sessions
> can read the rows, but cannot modify them until your transaction
> commits. If any of these rows were changed by another transaction
> that has not yet committed, your query waits until that 
> transaction ends and then uses the latest values.

Ent has recently added support for `FOR SHARE`/ `FOR UPDATE` 
statements via a feature-flag called `sql/lock`. To use it,
modify your `generate.go` file to include `--feature sql/lock`:

```go
//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate --feature sql/lock ./schema 
```

Next, let's implement a function that will use pessimistic
locking to make sure only a single process can update our `User` 
object's `online` field:

```go
func pessimisticUpdate(tx *ent.Tx, id int, online bool) (*ent.User, error) {
	ctx := context.Background()

	// On our active transaction, we begin a query against the user table
	u, err := tx.User.Query().

		// We add a predicate limiting the lock to the user we want to update.
		Where(user.ID(id)).

		// We use the ForUpdate method to tell ent to ask our DB to lock
		// the returned records for update.
		ForUpdate(
			// We specify that the query should not wait for the lock to be
			// released and instead fail immediately if the record is locked.
			sql.WithLockAction(sql.NoWait),
		).
		Only(ctx)
	
	// If we failed to acquire the lock we do not proceed to update the record.
	if err != nil {
		return nil, err
	}
	
	// Finally, we set the online field to the desired value. 
	return u.Update().SetOnline(online).Save(ctx)
}
```

Now, let's write a test that verifies that if two processes try to 
edit the same record, only one will succeed:

```go
func TestPessimistic(t *testing.T) {
	ctx := context.Background()
	client := enttest.Open(t, dialect.MySQL, "root:pass@tcp(localhost:3306)/test?parseTime=True")

	// Create the user for the first time.
	orig := client.User.Create().SetOnline(true).SaveX(ctx)

	// Open a new transaction. This transaction will acquire the lock on our user record.
	tx, err := client.Tx(ctx)
	if err != nil {
		log.Fatalf("failed creating transaction: %v", err)
	}
	defer tx.Commit()
	
	// Open a second transaction. This transaction is expected to fail at 
	// acquiring the lock on our user record. 
	tx2, err := client.Tx(ctx)
	if err != nil {
		log.Fatalf("failed creating transaction: %v", err)
	}
	defer tx.Commit()
	
	// The first update is expected to succeed.
	if _, err := pessimisticUpdate(tx, orig.ID, true); err != nil {
		log.Fatalf("unexpected error: %s", err)
	}
	
	// Because we did not run tx.Commit yet, the row is still locked when
	// we try to update it a second time. This operation is expected to 
	// fail. 
	_, err = pessimisticUpdate(tx2, orig.ID, true)
	if err == nil {
		log.Fatal("expected second update to fail")
	}
	fmt.Println(err)
}
```

A few things are worth mentioning in this example:

- Notice that we use a real MySQL instance to run this test 
  against, as SQLite does not support `SELECT .. FOR UPDATE`.
- For the simplicity of the example, we used the `sql.NoWait` 
  option to tell the database to return an error if the lock cannot be acquired. This means that the calling application needs to retry the write after receiving the error. If we don't specify this option, we can create flows where our application blocks until the lock is released and then proceeds without retrying. This is not always desirable but it opens up some interesting design options.
- We must always commit our transaction. Forgetting to do so can
  result in some serious issues. Remember that while the lock
  is maintained, no one can read or update this record.

Running our test:

```go
=== RUN   TestPessimistic
Error 3572: Statement aborted because lock(s) could not be acquired immediately and NOWAIT is set.
--- PASS: TestPessimistic (0.08s)
```

Great! We have used MySQL's "locking reads" capabilities and Ent's
new support for it to implement a locking mechanism that provides 
real mutual exclusion guarantees.

### Conclusion

We began this post by presenting the type of business requirements 
that lead application developers to reach out for locking techniques when working with databases. We continued by presenting two different approaches to achieving mutual exclusion when updating database records and demonstrated how to employ these techniques using Ent.

Have questions? Need help with getting started? Feel free to join 
our [Discord server](https://discord.gg/qZmPgTE6RX) or [Slack channel](https://entgo.io/docs/slack).

:::note For more Ent news and updates:

- Subscribe to our [Newsletter](https://entgo.substack.com/)
- Follow us on [Twitter](https://twitter.com/entgo_io)
- Join us on #ent on the [Gophers Slack](https://entgo.io/docs/slack)
- Join us on the [Ent Discord Server](https://discord.gg/qZmPgTE6RX)

:::
