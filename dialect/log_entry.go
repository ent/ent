// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package dialect

import (
	"database/sql"
	"fmt"
)

type DriverAction string

const (
	DriverActionTx      DriverAction = "Tx"
	DriverActionBeginTx DriverAction = "BeginTx"
	DriverActionExec    DriverAction = "Exec"
	DriverActionQuery   DriverAction = "Query"
)

type TxAction string

const (
	TxActionExec     TxAction = "Exec"
	TxActionQuery    TxAction = "Query"
	TxActionCommit   TxAction = "Commit"
	TxActionRollback TxAction = "Rollback"
)

type LogEntry struct {
	Action   DriverAction
	TxAction TxAction
	TxID     string
	TxOpt    *sql.TxOptions
	Query    string
	Args     interface{}
}

func (l LogEntry) String() string {
	switch l.TxAction {
	case TxActionExec:
		fallthrough
	case TxActionQuery:
		return fmt.Sprintf("%s(%s).%s: query=%v args=%v", l.Action, l.TxID, l.TxAction, l.Query, l.Args)
	case TxActionCommit:
		return fmt.Sprintf("%s(%s): committed", l.Action, l.TxID)
	case TxActionRollback:
		return fmt.Sprintf("%s(%s): rollbacked", l.Action, l.TxID)
	}

	switch l.Action {
	case DriverActionTx:
		fallthrough
	case DriverActionBeginTx:
		return fmt.Sprintf("driver.%s(%s): started", l.Action, l.TxID)
	case DriverActionExec:
		fallthrough
	case DriverActionQuery:
		return fmt.Sprintf("driver.%s: query=%v args=%v", l.Action, l.Query, l.Args)
	}

	panic(fmt.Errorf("no log action was specified"))
}
