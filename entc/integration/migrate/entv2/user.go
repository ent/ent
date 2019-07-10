// Code generated (@generated) by entc, DO NOT EDIT.

package entv2

import (
	"bytes"
	"fmt"
	"strconv"

	"fbc/ent/dialect/sql"

	"fbc/lib/go/gremlin"
)

// User is the model entity for the User schema.
type User struct {
	config
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Age holds the value of the "age" field.
	Age int `json:"age,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Phone holds the value of the "phone" field.
	Phone string `json:"phone,omitempty"`
}

// FromResponse scans the gremlin response data into User.
func (u *User) FromResponse(res *gremlin.Response) error {
	vmap, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	var vu struct {
		ID    string `json:"id,omitempty"`
		Age   int    `json:"age,omitempty"`
		Name  string `json:"name,omitempty"`
		Phone string `json:"phone,omitempty"`
	}
	if err := vmap.Decode(&vu); err != nil {
		return err
	}
	u.ID = vu.ID
	u.Age = vu.Age
	u.Name = vu.Name
	u.Phone = vu.Phone
	return nil
}

// FromRows scans the sql response data into User.
func (u *User) FromRows(rows *sql.Rows) error {
	var vu struct {
		ID    int
		Age   int
		Name  string
		Phone string
	}
	// the order here should be the same as in the `user.Columns`.
	if err := rows.Scan(
		&vu.ID,
		&vu.Age,
		&vu.Name,
		&vu.Phone,
	); err != nil {
		return err
	}
	u.ID = strconv.Itoa(vu.ID)
	u.Age = vu.Age
	u.Name = vu.Name
	u.Phone = vu.Phone
	return nil
}

// Update returns a builder for updating this User.
// Note that, you need to call User.Unwrap() before calling this method, if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return (&UserClient{u.config}).UpdateOne(u)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("entv2: User is not a transactional entity")
	}
	u.config.driver = tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	buf := bytes.NewBuffer(nil)
	buf.WriteString("User(")
	buf.WriteString(fmt.Sprintf("id=%v,", u.ID))
	buf.WriteString(fmt.Sprintf("age=%v", u.Age))
	buf.WriteString(fmt.Sprintf(", name=%v", u.Name))
	buf.WriteString(fmt.Sprintf(", phone=%v", u.Phone))
	buf.WriteString(")")
	return buf.String()
}

// id returns the int representation of the ID field.
func (u *User) id() int {
	id, _ := strconv.Atoi(u.ID)
	return id
}

// Users is a parsable slice of User.
type Users []*User

// FromResponse scans the gremlin response data into Users.
func (u *Users) FromResponse(res *gremlin.Response) error {
	vmap, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	var vu []struct {
		ID    string `json:"id,omitempty"`
		Age   int    `json:"age,omitempty"`
		Name  string `json:"name,omitempty"`
		Phone string `json:"phone,omitempty"`
	}
	if err := vmap.Decode(&vu); err != nil {
		return err
	}
	for _, v := range vu {
		*u = append(*u, &User{
			ID:    v.ID,
			Age:   v.Age,
			Name:  v.Name,
			Phone: v.Phone,
		})
	}
	return nil
}

// FromRows scans the sql response data into Users.
func (u *Users) FromRows(rows *sql.Rows) error {
	for rows.Next() {
		vu := &User{}
		if err := vu.FromRows(rows); err != nil {
			return err
		}
		*u = append(*u, vu)
	}
	return nil
}

func (u Users) config(cfg config) {
	for i := range u {
		u[i].config = cfg
	}
}
