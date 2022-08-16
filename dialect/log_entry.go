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
	DriverActionTx         DriverAction = "Tx"
	DriverActionBeginTx    DriverAction = "BeginTx"
	DriverActionTxCommit   DriverAction = "Tx.Commit"
	DriverActionTxRollback DriverAction = "Tx.Rollback"
	DriverActionExec       DriverAction = "Exec"
	DriverActionQuery      DriverAction = "Query"
)

type LogEntry struct {
	Action DriverAction
	TxID   string
	TxOpt  *sql.TxOptions
	Query  string
	Args   interface{}
}

func (l LogEntry) String() string {
	switch l.Action {
	case DriverActionTxCommit:
		return fmt.Sprintf("Tx(%s): committed", l.TxID)
	case DriverActionTxRollback:
		return fmt.Sprintf("Tx(%s): rollbacked", l.TxID)
	case DriverActionTx:
		fallthrough
	case DriverActionBeginTx:
		return fmt.Sprintf("driver.%s(%s): started", l.Action, l.TxID)
	case DriverActionExec:
		fallthrough
	case DriverActionQuery:
		if l.TxID != "" {
			return fmt.Sprintf("Tx(%s).%s: query=%v args=%v", l.TxID, l.Action, l.Query, l.Args)
		} else {
			return fmt.Sprintf("driver.%s: query=%v args=%v", l.Action, l.Query, l.Args)
		}
	}

	panic(fmt.Errorf("no log action was specified"))
}
