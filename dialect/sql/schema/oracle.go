package schema

import (
	"context"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/schema/field"
	"fmt"
	"strings"
)

// Oracle is a oracle migration driver.
type Oracle struct {
	dialect.Driver
	schema  string
	version string
}

func (o Oracle) init(ctx context.Context, tx dialect.Tx) error {
	rows := &sql.Rows{}
	if err := tx.Query(ctx, "select * from v$version", []interface{}{}, rows); err != nil {
		return fmt.Errorf("oracle: querying oracle version %w", err)
	}
	defer rows.Close()
	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return err
		}
		return fmt.Errorf("oracle: version variable was not found")
	}
	version := make([]string, 2)
	if err := rows.Scan(&version[0]); err != nil {
		return fmt.Errorf("oracle: scanning oracle version: %w", err)
	}
	o.version = version[1]
	return nil
}

func (o Oracle) table(ctx context.Context, tx dialect.Tx, name string) (*Table, error) {
	rows := &sql.Rows{}
	query, args := sql.Select("column_name", "column_type", "nullable", "data_default", "data_length", "data_precision").
		From(sql.Table("user_tab_columns")).
		Where(
			sql.EQ("table_name", name),
		).Query()
	if err := tx.Query(ctx, query, args, rows); err != nil {
		return nil, fmt.Errorf("oracle: reading table description %w", err)
	}
	// Call Close in cases of failures (Close is idempotent).
	defer rows.Close()
	t := NewTable(name)
	for rows.Next() {
		c := &Column{}
		if err := o.scanColumn(c, rows); err != nil {
			return nil, fmt.Errorf("oracle: %w", err)
		}
		if c.PrimaryKey() {
			t.PrimaryKey = append(t.PrimaryKey, c)
		}
		t.AddColumn(c)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	if err := rows.Close(); err != nil {
		return nil, fmt.Errorf("oracle: closing rows %w", err)
	}
	idxs, err := o.indexes(ctx, tx, name)
	if err != nil {
		return nil, err
	}
	for _, idx := range idxs {
		switch {
		case idx.primary:
			for _, name := range idx.columns {
				c, ok := t.column(name)
				if !ok {
					return nil, fmt.Errorf("index %q column %q was not found in table %q", idx.Name, name, t.Name)
				}
				c.Key = PrimaryKey
				t.PrimaryKey = append(t.PrimaryKey, c)
			}
		case idx.Unique && len(idx.columns) == 1:
			name := idx.columns[0]
			c, ok := t.column(name)
			if !ok {
				return nil, fmt.Errorf("index %q column %q was not found in table %q", idx.Name, name, t.Name)
			}
			c.Key = UniqueKey
			c.Unique = true
			fallthrough
		default:
			t.addIndex(idx)
		}
	}
	return t, nil
}

// scanColumn scans the information a column from column description.
func (o *Oracle) scanColumn(c *Column, rows *sql.Rows) error {
	var (
		nullable      sql.NullString
		defaults      sql.NullString
		udt           sql.NullString
		dataLen       sql.NullInt64
		dataPrecision sql.NullInt64
	)
	if err := rows.Scan(&c.Name, &c.typ, &nullable, &defaults, &dataLen, &dataPrecision); err != nil {
		return fmt.Errorf("scanning column description: %w", err)
	}
	if nullable.Valid {
		c.Nullable = nullable.String == "YES"
	}
	switch c.typ {
	case "number", "decimal":
		if v, err := dataPrecision.Value(); err != nil {
			if v != 0 {
				c.Type = field.TypeFloat64
			} else {
				c.Type = field.TypeInt64
			}
		}

	case "varchar2":
		c.Type = field.TypeString
		c.Size = maxCharSize + 1
	case "date", "time":
		c.Type = field.TypeTime
	case "bytea":
		c.Type = field.TypeBytes
	case "jsonb":
		c.Type = field.TypeJSON
	case "uuid":
		c.Type = field.TypeUUID
	case "cidr", "inet", "macaddr", "macaddr8":
		c.Type = field.TypeOther
	case "blob", "clob":
		c.Type = field.TypeOther
		if !udt.Valid {
			return fmt.Errorf("missing array type for column %q", c.Name)
		}
		// Note that for ARRAY types, the 'udt_name' column holds the array type
		// prefixed with '_'. For example, for 'integer[]' the result is '_int',
		// and for 'text[N][M]' the result is also '_text'. That's because, the
		// database ignores any size or multi-dimensions constraints.
		c.SchemaType = map[string]string{dialect.Oracle: "ARRAY"}
		c.typ = udt.String
	}
	switch {
	case !defaults.Valid || c.Type == field.TypeTime || seqfunc(defaults.String):
		return nil
	case strings.Contains(defaults.String, "::"):
		parts := strings.Split(defaults.String, "::")
		defaults.String = strings.Trim(parts[0], "'")
		fallthrough
	default:
		return c.ScanDefault(defaults.String)
	}
}

func (o *Oracle) indexes(ctx context.Context, tx dialect.Tx, table string) (Indexes, error) {
	rows := &sql.Rows{}
	query := "select t.index_name,i.index_type,i.uniqueness,t.column_name from user_ind_columns t,user_indexes i where t.index_name = i.index_name and t.table_name==upper('%s')"
	if err := tx.Query(ctx, fmt.Sprintf(query, table), nil, rows); err != nil {
		return nil, fmt.Errorf("querying indexes for table %s: %w", table, err)
	}
	defer rows.Close()
	var (
		idxs  Indexes
		names = make(map[string]*Index)
	)
	for rows.Next() {
		var (
			name, idxType, column, uniqueness string
			unique                            bool = false
		)
		if err := rows.Scan(&name, &idxType, &uniqueness, &column); err != nil {
			return nil, fmt.Errorf("scanning index description: %w", err)
		}
		if uniqueness == "UNIQUE" {
			unique = true

		}
		short := strings.TrimPrefix(name, table+"_")
		idx, ok := names[short]
		if !ok {
			idx = &Index{Name: short, Unique: unique, primary: unique, realname: name}
			idxs = append(idxs, idx)
			names[short] = idx
		}
		idx.columns = append(idx.columns, column)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return idxs, nil
}

func (o Oracle) tableExist(ctx context.Context, tx dialect.Tx, name string) (bool, error) {
	query, args := sql.Dialect(dialect.Oracle).
		Select(sql.Count("*")).From(sql.Table("user_tables")).
		Where(
			sql.EQ("table_name", strings.ToUpper(name)),
		).Query()
	return exist(ctx, tx, query, args...)
}

func (o Oracle) fkExist(ctx context.Context, tx dialect.Tx, name string) (bool, error) {
	return false, nil
}

func (o Oracle) setRange(ctx context.Context, tx dialect.Tx, table *Table, value int) error {
	panic("unsupported sequence")
}

func (o Oracle) dropIndex(ctx context.Context, tx dialect.Tx, idx *Index, table string) error {
	query, args := idx.DropBuilder(table).Query()
	return tx.Exec(ctx, query, args, nil)
}

func (o Oracle) cType(c *Column) (t string) {
	if c.SchemaType != nil && c.SchemaType[dialect.Postgres] != "" {
		return c.SchemaType[dialect.Postgres]
	}
	switch c.Type {
	case field.TypeBool:
		t = "number(1,0)"
	case field.TypeUint8, field.TypeInt8, field.TypeInt16, field.TypeUint16:
		t = "number(8,0)"
	case field.TypeInt32, field.TypeUint32:
		t = "number(10,0)"
	case field.TypeInt, field.TypeUint, field.TypeInt64, field.TypeUint64:
		t = "number(20,0)"
	case field.TypeFloat32:
		t = "number(20,2)"
	case field.TypeFloat64:
		t = "number(20,2)"
	case field.TypeBytes:
		t = "number(1)"
	case field.TypeJSON:
		t = "clob"
	case field.TypeUUID:
		t = "varchar2(50)"
	case field.TypeString:
		t = "varchar2(100)"
		if c.Size > maxCharSize {
			t = "clob"
		}
	case field.TypeTime:
		t = c.scanTypeOr("DATE")
	case field.TypeEnum:
		// Currently, the support for enums is weak (application level only.
		// like SQLite). Dialect needs to create and maintain its enum type.
		t = "varchar2(100)"
	case field.TypeOther:
		t = c.typ
	default:
		panic(fmt.Sprintf("unsupported type %q for column %q", c.Type.String(), c.Name))
	}
	return t
}

func (o Oracle) tBuilder(t *Table) *sql.TableBuilder {
	b := sql.Dialect(dialect.Oracle).
		CreateTable(t.Name)
	for _, c := range t.Columns {
		b.Column(o.addColumn(c))
	}
	for _, pk := range t.PrimaryKey {
		b.PrimaryKey(pk.Name)
	}
	return b
}

// addColumn returns the ColumnBuilder for adding the given column to a table.
func (o *Oracle) addColumn(c *Column) *sql.ColumnBuilder {
	b := sql.Dialect(dialect.Oracle).
		Column(c.Name).Type(o.cType(c)).Attr(c.Attr)
	c.unique(b)
	if c.Increment {
		//panic("unsupported increment!")
	}
	c.nullable(b)
	c.defaultValue(b)
	return b
}

func (o Oracle) addIndex(i *Index, table string) *sql.IndexBuilder {
	return i.Builder(table)
}

func (o Oracle) alterColumns(table string, add, modify, drop []*Column) sql.Queries {
	panic("implement me")
}

func (o Oracle) needsConversion(column *Column, column2 *Column) bool {
	panic("implement me")
}
