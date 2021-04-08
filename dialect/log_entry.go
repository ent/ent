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

type logEntry struct {
	Action   DriverAction
	TxAction TxAction
	TxId     string
	TxOpt    *sql.TxOptions
	Query    string
	Args     interface{}
}

func (l logEntry) String() string {
	switch l.TxAction {
	case TxActionExec:
		fallthrough
	case TxActionQuery:
		return fmt.Sprintf("%s(%s).%s: query=%v args=%v", l.Action, l.TxId, l.TxAction, l.Query, l.Args)
	case TxActionCommit:
		return fmt.Sprintf("%s(%s): committed", l.Action, l.TxId)
	case TxActionRollback:
		return fmt.Sprintf("%s(%s): rollbacked", l.Action, l.TxId)
	}

	switch l.Action {
	case DriverActionTx:
		fallthrough
	case DriverActionBeginTx:
		return fmt.Sprintf("driver.%s(%s): started", l.Action, l.TxId)
	case DriverActionExec:
		fallthrough
	case DriverActionQuery:
		return fmt.Sprintf("driver.%s: query=%v args=%v", l.Action, l.Query, l.Args)
	}

	panic(fmt.Errorf("no log action was specified"))
}
