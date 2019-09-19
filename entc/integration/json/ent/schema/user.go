package schema

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.JSON("url", &url.URL{}).
			Optional(),
		field.JSON("raw", json.RawMessage{}).
			Optional(),
		field.JSON("dirs", []http.Dir{}).
			Optional(),
		field.Ints("ints").
			Optional(),
		field.Floats("floats").
			Optional(),
		field.Strings("strings").
			Optional(),
	}
}
