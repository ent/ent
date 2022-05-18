package neo4j

import (
	"context"
	"entgo.io/ent/dialect"
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type neo4jDialectDriver struct {
	driver         neo4j.Driver
	currentSession neo4j.Session
	currentTx      *neo4jTransaction
}

func NewNeo4jDriver(driver neo4j.Driver) dialect.Driver {
	return &neo4jDialectDriver{
		driver:         driver,
		currentSession: driver.NewSession(neo4j.SessionConfig{}),
	}
}

func (driver *neo4jDialectDriver) Exec(ctx context.Context, query string, args, v interface{}) error {
	driver.maybeInit()
	return driver.currentTx.Exec(ctx, query, args, v)
}

func (driver *neo4jDialectDriver) Query(ctx context.Context, query string, args, v interface{}) error {
	driver.maybeInit()
	return driver.currentTx.Query(ctx, query, args, v)
}

func (driver *neo4jDialectDriver) Tx(ctx context.Context) (dialect.Tx, error) {
	driver.maybeInit()
	return driver.currentTx, nil
}

func (driver *neo4jDialectDriver) Close() error {
	err1 := driver.currentSession.Close()
	err2 := driver.driver.Close()
	switch {
	case err1 != nil && err2 != nil:
		return fmt.Errorf("dialect/neo4j: could not close either session: %v, or driver: %w", err1, err2)
	case err1 != nil:
		return err1
	default:
		return err2
	}
}

func (driver *neo4jDialectDriver) Dialect() string {
	return dialect.Neo4j
}

func (driver *neo4jDialectDriver) maybeInit() {
	if driver.currentTx == nil {
		driver.currentTx = &neo4jTransaction{session: driver.currentSession}
	}
}

type neo4jTransaction struct {
	session  neo4j.Session
	tx       neo4j.Transaction
	executor *cypherExecutor
}

func (transaction *neo4jTransaction) Exec(ctx context.Context, query string, args, v interface{}) error {
	if err := transaction.maybeInit(); err != nil {
		return err
	}
	return transaction.executor.Exec(ctx, query, args, v)
}

func (transaction *neo4jTransaction) Query(ctx context.Context, query string, args, v interface{}) error {
	if err := transaction.maybeInit(); err != nil {
		return err
	}
	return transaction.executor.Query(ctx, query, args, v)
}

func (transaction *neo4jTransaction) Commit() error {
	return transaction.tx.Commit()
}

func (transaction *neo4jTransaction) Rollback() error {
	return transaction.tx.Rollback()
}

func (transaction *neo4jTransaction) maybeInit() error {
	if transaction.tx == nil {
		tx, err := transaction.session.BeginTransaction()
		if err != nil {
			return err
		}
		transaction.tx = tx
		transaction.executor = &cypherExecutor{tx: tx}
	}
	return nil
}

type cypherExecutor struct {
	// TODO: autocommit queries can only be run with session.Run
	// explicit tx like this would not work in those cases
	tx neo4j.Transaction
}

func (executor *cypherExecutor) Exec(ctx context.Context, query string, args, v interface{}) error {
	params, err := executor.validateParams(args)
	if err != nil {
		return err
	}
	result, ok := v.(*neo4j.ResultSummary)
	if !ok {
		return fmt.Errorf("dialect/neo4j: invalid type %T for result receiver, expected *neo4j.ResultSummary", v)
	}
	summary, err := executor.executeQuery(query, params)
	if err != nil {
		return err
	}
	*result = summary.(neo4j.ResultSummary)
	return nil
}

func (executor *cypherExecutor) Query(ctx context.Context, query string, args, v interface{}) error {
	params, err := executor.validateParams(args)
	if err != nil {
		return err
	}
	result, ok := v.(*[]map[string]interface{})
	if !ok {
		return fmt.Errorf("dialect/neo4j: invalid type %T for result receiver, expected *[]map[string]interface{}", v)
	}
	summary, err := executor.runQuery(query, params)
	if err != nil {
		return nil
	}
	*result = summary
	return nil
}

func (executor *cypherExecutor) executeQuery(query string, params map[string]interface{}) (neo4j.ResultSummary, error) {
	result, err := executor.tx.Run(query, params)
	if err != nil {
		return nil, err
	}
	summary, err := result.Consume()
	if err != nil {
		return nil, err
	}
	return summary, nil
}

func (executor *cypherExecutor) runQuery(query string, params map[string]interface{}) ([]map[string]interface{}, error) {
	results, err := executor.tx.Run(query, params)
	if err != nil {
		return nil, err
	}
	var result []map[string]interface{}
	for results.Next() {
		record := results.Record()
		result = append(result, recordAsMap(record))
	}
	if err = results.Err(); err != nil {
		return nil, err
	}
	return result, nil
}

func (executor *cypherExecutor) validateParams(args interface{}) (map[string]interface{}, error) {
	params, ok := args.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("dialect/neo4j: invalid type %T for parameters, expected map[string]interface{}", args)
	}
	return params, nil
}

func recordAsMap(record *neo4j.Record) map[string]interface{} {
	result := make(map[string]interface{}, len(record.Keys))
	for _, key := range record.Keys {
		result[key], _ = record.Get(key)
	}
	return result
}
