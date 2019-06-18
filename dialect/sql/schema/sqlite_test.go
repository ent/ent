package schema

import (
	"context"
	"testing"

	"fbc/ent/dialect/sql"
	"fbc/ent/field"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func TestSQLite_Create(t *testing.T) {
	null := true
	tests := []struct {
		name    string
		tables  []*Table
		before  func(sqlmock.Sqlmock)
		wantErr bool
	}{
		{
			name: "tx failed",
			before: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin().WillReturnError(sqlmock.ErrCancelled)
			},
			wantErr: true,
		},
		{
			name: "fk disabled",
			before: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectQuery("PRAGMA foreign_keys").
					WillReturnRows(sqlmock.NewRows([]string{"foreign_keys"}).AddRow(0))
				mock.ExpectRollback()
			},
			wantErr: true,
		},
		{
			name: "no tables",
			before: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectQuery("PRAGMA foreign_keys").
					WillReturnRows(sqlmock.NewRows([]string{"foreign_keys"}).AddRow(1))
				mock.ExpectCommit()
			},
		},
		{
			name: "create new table",
			tables: []*Table{
				{
					Name: "users",
					PrimaryKey: []*Column{
						{Name: "id", Type: field.TypeInt, Increment: true},
					},
					Columns: []*Column{
						{Name: "id", Type: field.TypeInt, Increment: true},
						{Name: "name", Type: field.TypeString, Nullable: &null},
						{Name: "age", Type: field.TypeInt},
					},
				},
			},
			before: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectQuery("PRAGMA foreign_keys").
					WillReturnRows(sqlmock.NewRows([]string{"foreign_keys"}).AddRow(1))
				mock.ExpectQuery(escape("SELECT COUNT(*) FROM `sqlite_master` WHERE `type` = ? AND `name` = ?")).
					WithArgs("table", "users").
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(0))
				mock.ExpectExec(escape("CREATE TABLE `users`(`id` integer PRIMARY KEY AUTOINCREMENT, `name` varchar(255) NULL, `age` integer)")).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectCommit()
			},
		},
		{
			name: "create new table with foreign key",
			tables: func() []*Table {
				var (
					c1 = []*Column{
						{Name: "id", Type: field.TypeInt, Increment: true},
						{Name: "name", Type: field.TypeString, Nullable: &null},
						{Name: "created_at", Type: field.TypeTime},
					}
					c2 = []*Column{
						{Name: "id", Type: field.TypeInt, Increment: true},
						{Name: "name", Type: field.TypeString},
						{Name: "owner_id", Type: field.TypeInt},
					}
					t1 = &Table{
						Name:       "users",
						Columns:    c1,
						PrimaryKey: c1[0:1],
					}
					t2 = &Table{
						Name:       "pets",
						Columns:    c2,
						PrimaryKey: c2[0:1],
						ForeignKeys: []*ForeignKey{
							{
								Symbol:     "pets_owner",
								Columns:    c2[2:],
								RefTable:   t1,
								RefColumns: c1[0:1],
								OnDelete:   Cascade,
							},
						},
					}
				)
				return []*Table{t1, t2}
			}(),
			before: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectQuery("PRAGMA foreign_keys").
					WillReturnRows(sqlmock.NewRows([]string{"foreign_keys"}).AddRow(1))
				mock.ExpectQuery(escape("SELECT COUNT(*) FROM `sqlite_master` WHERE `type` = ? AND `name` = ?")).
					WithArgs("table", "users").
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(0))
				mock.ExpectExec(escape("CREATE TABLE `users`(`id` integer PRIMARY KEY AUTOINCREMENT, `name` varchar(255) NULL, `created_at` datetime)")).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectQuery(escape("SELECT COUNT(*) FROM `sqlite_master` WHERE `type` = ? AND `name` = ?")).
					WithArgs("table", "pets").
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(0))
				mock.ExpectExec(escape("CREATE TABLE `pets`(`id` integer PRIMARY KEY AUTOINCREMENT, `name` varchar(255), `owner_id` integer, FOREIGN KEY(`owner_id`) REFERENCES `users`(`id`) ON DELETE CASCADE)")).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectCommit()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			require.NoError(t, err)
			tt.before(mock)
			sqlite := &SQLite{sql.OpenDB("sqlite", db)}
			err = sqlite.Create(context.Background(), tt.tables...)
			require.Equal(t, tt.wantErr, err != nil, err)
		})
	}
}
