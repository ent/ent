package schema

import (
	"context"
	"regexp"
	"strings"
	"testing"

	"fbc/ent/dialect/sql"
	"fbc/ent/field"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func TestMySQL_Create(t *testing.T) {
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
			name: "no tables",
			before: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectQuery(escape("SHOW VARIABLES LIKE 'version'")).
					WillReturnRows(sqlmock.NewRows([]string{"Variable_name", "Value"}).AddRow("version", "5.7.23"))
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
						{Name: "name", Type: field.TypeString, Nullable: &null, Charset: "utf8"},
						{Name: "age", Type: field.TypeInt},
					},
				},
			},
			before: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectQuery(escape("SHOW VARIABLES LIKE 'version'")).
					WillReturnRows(sqlmock.NewRows([]string{"Variable_name", "Value"}).AddRow("version", "5.7.23"))
				mock.ExpectQuery(escape("SELECT COUNT(*) FROM INFORMATION_SCHEMA.TABLES WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ?")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(0))
				mock.ExpectExec(escape("CREATE TABLE IF NOT EXISTS `users`(`id` int AUTO_INCREMENT, `name` varchar(255) CHARSET utf8 NULL, `age` int, PRIMARY KEY(`id`)) CHARACTER SET utf8mb4")).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectCommit()
			},
		},
		{
			name: "create new table 5.6",
			tables: []*Table{
				{
					Name: "users",
					PrimaryKey: []*Column{
						{Name: "id", Type: field.TypeInt, Increment: true},
					},
					Columns: []*Column{
						{Name: "id", Type: field.TypeInt, Increment: true},
						{Name: "age", Type: field.TypeInt},
						{Name: "name", Type: field.TypeString, Unique: true},
					},
				},
			},
			before: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectQuery(escape("SHOW VARIABLES LIKE 'version'")).
					WillReturnRows(sqlmock.NewRows([]string{"Variable_name", "Value"}).AddRow("version", "5.6.35"))
				mock.ExpectQuery(escape("SELECT COUNT(*) FROM INFORMATION_SCHEMA.TABLES WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ?")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(0))
				mock.ExpectExec(escape("CREATE TABLE IF NOT EXISTS `users`(`id` int AUTO_INCREMENT, `age` int, `name` varchar(191) UNIQUE, PRIMARY KEY(`id`)) CHARACTER SET utf8mb4")).
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
				mock.ExpectQuery(escape("SHOW VARIABLES LIKE 'version'")).
					WillReturnRows(sqlmock.NewRows([]string{"Variable_name", "Value"}).AddRow("version", "5.7.23"))
				mock.ExpectQuery(escape("SELECT COUNT(*) FROM INFORMATION_SCHEMA.TABLES WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ?")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(0))
				mock.ExpectExec(escape("CREATE TABLE IF NOT EXISTS `users`(`id` int AUTO_INCREMENT, `name` varchar(255) NULL, `created_at` timestamp NULL, PRIMARY KEY(`id`)) CHARACTER SET utf8mb4")).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectQuery(escape("SELECT COUNT(*) FROM INFORMATION_SCHEMA.TABLES WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ?")).
					WithArgs("pets").
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(0))
				mock.ExpectExec(escape("CREATE TABLE IF NOT EXISTS `pets`(`id` int AUTO_INCREMENT, `name` varchar(255), `owner_id` int, PRIMARY KEY(`id`)) CHARACTER SET utf8mb4")).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectQuery(escape("SELECT COUNT(*) FROM INFORMATION_SCHEMA.TABLE_CONSTRAINTS WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `CONSTRAINT_TYPE` = ? AND `CONSTRAINT_NAME` = ?")).
					WithArgs("FOREIGN KEY", "pets_owner").
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(0))
				mock.ExpectExec(escape("ALTER TABLE `pets` ADD CONSTRAINT `pets_owner` FOREIGN KEY(`owner_id`) REFERENCES `users`(`id`) ON DELETE CASCADE")).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectCommit()
			},
		},
		{
			name: "add column to table",
			tables: []*Table{
				{
					Name: "users",
					Columns: []*Column{
						{Name: "id", Type: field.TypeInt, Increment: true},
						{Name: "name", Type: field.TypeString, Nullable: &null},
						{Name: "age", Type: field.TypeInt},
					},
					PrimaryKey: []*Column{
						{Name: "id", Type: field.TypeInt, Increment: true},
					},
				},
			},
			before: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectQuery(escape("SHOW VARIABLES LIKE 'version'")).
					WillReturnRows(sqlmock.NewRows([]string{"Variable_name", "Value"}).AddRow("version", "5.7.23"))
				mock.ExpectQuery(escape("SELECT COUNT(*) FROM INFORMATION_SCHEMA.TABLES WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ?")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))
				mock.ExpectQuery(escape("SELECT `column_name`, `column_type`, `is_nullable`, `column_key`, `column_default`, `extra`, `character_set_name`, `collation_name` FROM INFORMATION_SCHEMA.COLUMNS WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ?")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"column_name", "column_type", "is_nullable", "column_key", "column_default", "extra", "character_set_name", "collation_name"}).
						AddRow("id", "int(11)", "NO", "PRI", "NULL", "auto_increment", "", "").
						AddRow("name", "varchar(255)", "NO", "YES", "NULL", "", "", ""))
				mock.ExpectExec(escape("ALTER TABLE `users` ADD COLUMN `age` int")).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectCommit()
			},
		},
		{
			name: "modify column",
			tables: []*Table{
				{
					Name: "users",
					Columns: []*Column{
						{Name: "id", Type: field.TypeInt, Increment: true},
						{Name: "age", Type: field.TypeInt},
						{Name: "name", Type: field.TypeString, Nullable: &null, Charset: "utf8"},
					},
					PrimaryKey: []*Column{
						{Name: "id", Type: field.TypeInt, Increment: true},
					},
				},
			},
			before: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectQuery(escape("SHOW VARIABLES LIKE 'version'")).
					WillReturnRows(sqlmock.NewRows([]string{"Variable_name", "Value"}).AddRow("version", "5.7.23"))
				mock.ExpectQuery(escape("SELECT COUNT(*) FROM INFORMATION_SCHEMA.TABLES WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ?")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))
				mock.ExpectQuery(escape("SELECT `column_name`, `column_type`, `is_nullable`, `column_key`, `column_default`, `extra`, `character_set_name`, `collation_name` FROM INFORMATION_SCHEMA.COLUMNS WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ?")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"column_name", "column_type", "is_nullable", "column_key", "column_default", "extra", "character_set_name", "collation_name"}).
						AddRow("id", "int(11)", "NO", "PRI", "NULL", "auto_increment", "", "").
						AddRow("name", "varchar(255)", "NO", "YES", "NULL", "", "", "").
						AddRow("age", "int(11)", "NO", "NO", "NULL", "", "", ""))
				mock.ExpectExec(escape("ALTER TABLE `users` MODIFY COLUMN `name` varchar(255) CHARSET utf8 NULL")).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectCommit()
			},
		},
		{
			name: "add edge to table",
			tables: func() []*Table {
				var (
					c1 = []*Column{
						{Name: "id", Type: field.TypeInt, Increment: true},
						{Name: "name", Type: field.TypeString, Nullable: &null},
						{Name: "spouse_id", Type: field.TypeInt},
					}
					t1 = &Table{
						Name:       "users",
						Columns:    c1,
						PrimaryKey: c1[0:1],
						ForeignKeys: []*ForeignKey{
							{
								Symbol:     "user_spouse" + strings.Repeat("_", 64), // super long fk.
								Columns:    c1[2:],
								RefColumns: c1[0:1],
								OnDelete:   Cascade,
							},
						},
					}
				)
				t1.ForeignKeys[0].RefTable = t1
				return []*Table{t1}
			}(),
			before: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectQuery(escape("SHOW VARIABLES LIKE 'version'")).
					WillReturnRows(sqlmock.NewRows([]string{"Variable_name", "Value"}).AddRow("version", "5.7.23"))
				mock.ExpectQuery(escape("SELECT COUNT(*) FROM INFORMATION_SCHEMA.TABLES WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ?")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))
				mock.ExpectQuery(escape("SELECT `column_name`, `column_type`, `is_nullable`, `column_key`, `column_default`, `extra`, `character_set_name`, `collation_name` FROM INFORMATION_SCHEMA.COLUMNS WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ?")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"column_name", "column_type", "is_nullable", "column_key", "column_default", "extra", "character_set_name", "collation_name"}).
						AddRow("id", "int(11)", "NO", "PRI", "NULL", "auto_increment", "", "").
						AddRow("name", "varchar(255)", "NO", "YES", "NULL", "", "", ""))
				mock.ExpectExec(escape("ALTER TABLE `users` ADD COLUMN `spouse_id` int")).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectQuery(escape("SELECT COUNT(*) FROM INFORMATION_SCHEMA.TABLE_CONSTRAINTS WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `CONSTRAINT_TYPE` = ? AND `CONSTRAINT_NAME` = ?")).
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(0))
				mock.ExpectExec("ALTER TABLE `users` ADD CONSTRAINT `.{64}` FOREIGN KEY\\(`spouse_id`\\) REFERENCES `users`\\(`id`\\) ON DELETE CASCADE").
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
			mysql := &MySQL{sql.OpenDB("mysql", db)}
			err = mysql.Create(context.Background(), tt.tables...)
			require.Equal(t, tt.wantErr, err != nil, err)
		})
	}
}

func escape(query string) string {
	rows := strings.Split(query, "\n")
	for i := range rows {
		rows[i] = strings.TrimPrefix(rows[i], " ")
	}
	query = strings.Join(rows, " ")
	return regexp.QuoteMeta(query)
}
