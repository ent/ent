package sql

import (
	"entgo.io/ent/dialect"
)

func MaxParams(dialectName string) int {
	switch dialectName {
	case dialect.Postgres:
		return 65535
	default:
		return -1
	}
}

func Chunked(fn func(start, end int) error, n int, maxPerChunk int) error {
	if maxPerChunk == -1 {
		return fn(0, n)
	}

	for start := 0; start < n; start += maxPerChunk {
		end := start + maxPerChunk
		if end > n {
			end = n
		}

		err := fn(start, end)
		if err != nil {
			return err
		}
	}

	return nil
}
