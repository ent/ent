package user

import (
	"database/sql/driver"
)

type User struct {
	Name string
}

func (u *User) Value() (driver.Value, error) {
	return u.Name, nil
}

func (u *User) Scan(val any) error {
	switch v := val.(type) {
	case nil:
		return nil
	case string:
		u.Name = v
	case []uint8:
		u.Name = string(v)
	}
	return nil
}
