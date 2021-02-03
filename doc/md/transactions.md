---
id: transactions
title: Transactions
---

## Starting A Transaction

```go
// GenTx generates group of entities in a transaction.
func GenTx(ctx context.Context, client *ent.Client) error {
	tx, err := client.Tx(ctx)
	if err != nil {
		return fmt.Errorf("starting a transaction: %v", err)
	}
	hub, err := tx.Group.
		Create().
		SetName("Github").
		Save(ctx)
	if err != nil {
		return rollback(tx, fmt.Errorf("failed creating the group: %v", err))
	}
	// Create the admin of the group.
	dan, err := tx.User.
		Create().
		SetAge(29).
		SetName("Dan").
		AddManage(hub).
		Save(ctx)
	if err != nil {
		return rollback(tx, err)
	}
	// Create user "Ariel".
	a8m, err := tx.User.
		Create().
		SetAge(30).
		SetName("Ariel").
		AddGroups(hub).
		AddFriends(dan).
		Save(ctx)
	if err != nil {
		return rollback(tx, err)
	}
	fmt.Println(a8m)
	// Output:
	// User(id=2, age=30, name=Ariel)
	
	// Commit the transaction.
	return tx.Commit()
}

// rollback calls to tx.Rollback and wraps the given error
// with the rollback error if occurred.
func rollback(tx *ent.Tx, err error) error {
	if rerr := tx.Rollback(); rerr != nil {
		err = fmt.Errorf("%v: %v", err, rerr)
	}
	return err
}
```

The full example exists in [GitHub](https://github.com/facebook/ent/tree/master/examples/traversal).

## Transactional Client

Sometimes, you have an existing code that already works with `*ent.Client`, and you want to change it (or wrap it)
to interact with transactions. For these use cases, you have a transactional client. An `*ent.Client` that you can
get from an existing transaction.

```go
// WrapGen wraps the existing "Gen" function in a transaction.
func WrapGen(ctx context.Context, client *ent.Client) error {
	tx, err := client.Tx(ctx)
	if err != nil {
		return err
	}
	txClient := tx.Client()
	// Use the "Gen" below, but give it the transactional client; no code changes to "Gen".
	if err := Gen(ctx, txClient); err != nil {
		return rollback(tx, err)
	}
	return tx.Commit()
}

// Gen generates a group of entities.
func Gen(ctx context.Context, client *ent.Client) error {
	// ...
	return nil
}
```

The full example exists in [GitHub](https://github.com/facebook/ent/tree/master/examples/traversal).

## Best Practices

Reusable function that runs callbacks in a transaction:

```go
func WithTx(ctx context.Context, client *ent.Client, fn func(tx *ent.Tx) error) error {
	tx, err := client.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if v := recover(); v != nil {
			tx.Rollback()
			panic(v)
		}
	}()
	if err := fn(tx); err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = errors.Wrapf(err, "rolling back transaction: %v", rerr)
		}
		return err
	}
	if err := tx.Commit(); err != nil {
		return errors.Wrapf(err, "committing transaction: %v", err)
	}
	return nil
}
```

Its usage:

```go
func Do(ctx context.Context, client *ent.Client) {
	// WithTx helper.
	if err := WithTx(ctx, client, func(tx *ent.Tx) error {
		return Gen(ctx, tx.Client())
	}); err != nil {
		log.Fatal(err)
	}
}
```

## Hooks

Same as [schema hooks](hooks.md#schema-hooks) and [runtime hooks](hooks.md#runtime-hooks), hooks can be registered on
active transactions, and will be executed on `Tx.Commit` or `Tx.Rollback`:

```go
func Do(ctx context.Context, client *ent.Client) error {
    tx, err := client.Tx(ctx)
    if err != nil {
        return err
    }
    // Add a hook on Tx.Commit.
    tx.OnCommit(func(next ent.Committer) ent.Committer {
        return ent.CommitFunc(func(ctx context.Context, tx *ent.Tx) error {
            // Code before the actual commit.
            err := next.Commit(ctx, tx)
            // Code after the transaction was committed.
            return err
        })
    })
    // Add a hook on Tx.Rollback.
    tx.OnRollback(func(next ent.Rollbacker) ent.Rollbacker {
        return ent.RollbackFunc(func(ctx context.Context, tx *ent.Tx) error {
            // Code before the actual rollback.
            err := next.Rollback(ctx, tx)
            // Code after the transaction was rolled back.
            return err
        })
    })
    //
    // <Code goes here>
    //
    return err
}
```
