package types

import (
	"database/sql/driver"
	"fmt"
	"strings"
)

type Strings []string

func (s *Strings) Scan(v interface{}) error {
	switch v := v.(type) {
	case nil:
	case []byte:
		return s.scan(string(v))
	case string:
		return s.scan(v)
	default:
		return fmt.Errorf("unexpected type %T", v)
	}

	return nil
}

func (s *Strings) scan(v string) error {
	if v == "" {
		return nil
	}

	if l := len(v); l < 2 || v[0] != '{' && v[l-1] != '}' {
		return fmt.Errorf("unexpcted array format %q", v)
	}

	*s = strings.Split(v[1:len(v)-1], ",")

	return nil
}

func (s Strings) Value() (driver.Value, error) {
	return "{" + strings.Join(s, ",") + "}", nil
}
