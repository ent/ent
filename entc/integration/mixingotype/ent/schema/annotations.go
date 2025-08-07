package schema

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type Annotations = map[string]interface{}

// AnnotationsMixin adds annotations to the schema
type AnnotationsMixin struct {
	mixin.Schema
}

// Fields of the IDMixin.
func (AnnotationsMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("annotations").
			GoType(Annotations{}).
			ValueScanner(JSONStringToMapValueScanner()).
			Optional(),
	}
}

// JSONStringToMapValueScanner is a value scanner for the Annotations field, this is just one example of a custom value scanner.
func JSONStringToMapValueScanner() field.ValueScannerFunc[Annotations, *sql.NullString] {
	return field.ValueScannerFunc[Annotations, *sql.NullString]{
		V: func(t Annotations) (driver.Value, error) {
			if len(t) == 0 {
				return nil, nil
			}

			return json.Marshal(t)
		},
		S: func(ns *sql.NullString) (Annotations, error) {
			v := new(Annotations)
			if ns == nil || !ns.Valid {
				return *v, nil
			}

			b := []byte(ns.String)
			if err := json.Unmarshal(b, v); err != nil {
				return *v, err
			}

			return *v, nil
		},
	}
}
