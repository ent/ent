package dialect

import (
	"testing"
)

func Test_logEntry_String(t *testing.T) {
	tests := []struct {
		entry logEntry
		want  string
	}{
		{
			entry: logEntry{
				Action: DriverActionExec,
				Query:  "SELECT * FROM nothing WHERE id = ?",
				Args:   []interface{}{1},
			},
			want: "driver.Exec: query=SELECT * FROM nothing WHERE id = ? args=[1]",
		},
		{
			entry: logEntry{
				Action: DriverActionQuery,
				Query:  "SELECT * FROM nothing WHERE id = ?",
				Args:   []interface{}{1},
			},
			want: "driver.Query: query=SELECT * FROM nothing WHERE id = ? args=[1]",
		},
		{
			entry: logEntry{
				Action: DriverActionTx,
				TxId:   "tx-id",
			},
			want: "driver.Tx(tx-id): started",
		},
		{
			entry: logEntry{
				Action: DriverActionBeginTx,
				TxId:   "tx-id",
			},
			want: "driver.BeginTx(tx-id): started",
		},
		{
			entry: logEntry{
				Action:   DriverActionTx,
				TxAction: TxActionExec,
				TxId:     "tx-id",
				Query:    "SELECT * FROM nothing WHERE id = ?",
				Args:     []interface{}{1},
			},
			want: "Tx(tx-id).Exec: query=SELECT * FROM nothing WHERE id = ? args=[1]",
		},
		{
			entry: logEntry{
				Action:   DriverActionTx,
				TxAction: TxActionQuery,
				TxId:     "tx-id",
				Query:    "SELECT * FROM nothing WHERE id = ?",
				Args:     []interface{}{1},
			},
			want: "Tx(tx-id).Query: query=SELECT * FROM nothing WHERE id = ? args=[1]",
		},
		{
			entry: logEntry{
				Action:   DriverActionTx,
				TxAction: TxActionCommit,
				TxId:     "tx-id",
			},
			want: "Tx(tx-id): committed",
		},
		{
			entry: logEntry{
				Action:   DriverActionTx,
				TxAction: TxActionRollback,
				TxId:     "tx-id",
			},
			want: "Tx(tx-id): rollbacked",
		},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := tt.entry.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
