package sid

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type ID string

func (i ID) String() string {
	return string(i)
}

func (i ID) Value() (driver.Value, error) {
	r, err := strconv.ParseInt(string(i), 10, 64)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (i *ID) Scan(src interface{}) error {
	switch v := src.(type) {
	case nil:
		return nil
	case int64:
		*i = ID(fmt.Sprint(v))
		return nil
	}
	return errors.New("not a valid base62")
}

func New() ID {
	return NewLength(10)
}
func NewLength(len int) ID {
	return ID(randomString(len))
}

func randomString(n int) string {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	var out string
	for len(out) < n {
		out += fmt.Sprint(r.Int())
	}

	return out[:n]
}
