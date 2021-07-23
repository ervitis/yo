// Code generated by yo. DO NOT EDIT.

// Package models contains the types.
package models

import (
	"context"
	"fmt"

	"cloud.google.com/go/spanner"
)

// OutOfOrderPrimaryKey represents a row from 'OutOfOrderPrimaryKeys'.
type OutOfOrderPrimaryKey struct {
	PKey1 string `spanner:"PKey1" json:"PKey1"` // PKey1
	PKey2 string `spanner:"PKey2" json:"PKey2"` // PKey2
	PKey3 string `spanner:"PKey3" json:"PKey3"` // PKey3
}

func OutOfOrderPrimaryKeyPrimaryKeys() []string {
	return []string{
		"PKey2",
		"PKey1",
		"PKey3",
	}
}

func OutOfOrderPrimaryKeyColumns() []string {
	return []string{
		"PKey1",
		"PKey2",
		"PKey3",
	}
}

func OutOfOrderPrimaryKeyWritableColumns() []string {
	return []string{
		"PKey1",
		"PKey2",
		"PKey3",
	}
}

func (ooopk *OutOfOrderPrimaryKey) columnsToPtrs(cols []string) ([]interface{}, error) {
	ret := make([]interface{}, 0, len(cols))
	for _, col := range cols {
		switch col {
		case "PKey1":
			ret = append(ret, yoDecode(&ooopk.PKey1))
		case "PKey2":
			ret = append(ret, yoDecode(&ooopk.PKey2))
		case "PKey3":
			ret = append(ret, yoDecode(&ooopk.PKey3))
		default:
			return nil, fmt.Errorf("unknown column: %s", col)
		}
	}
	return ret, nil
}

func (ooopk *OutOfOrderPrimaryKey) columnsToValues(cols []string) ([]interface{}, error) {
	ret := make([]interface{}, 0, len(cols))
	for _, col := range cols {
		switch col {
		case "PKey1":
			ret = append(ret, yoEncode(ooopk.PKey1))
		case "PKey2":
			ret = append(ret, yoEncode(ooopk.PKey2))
		case "PKey3":
			ret = append(ret, yoEncode(ooopk.PKey3))
		default:
			return nil, fmt.Errorf("unknown column: %s", col)
		}
	}

	return ret, nil
}

// newOutOfOrderPrimaryKey_Decoder returns a decoder which reads a row from *spanner.Row
// into OutOfOrderPrimaryKey. The decoder is not goroutine-safe. Don't use it concurrently.
func newOutOfOrderPrimaryKey_Decoder(cols []string) func(*spanner.Row) (*OutOfOrderPrimaryKey, error) {
	return func(row *spanner.Row) (*OutOfOrderPrimaryKey, error) {
		var ooopk OutOfOrderPrimaryKey
		ptrs, err := ooopk.columnsToPtrs(cols)
		if err != nil {
			return nil, err
		}

		if err := row.Columns(ptrs...); err != nil {
			return nil, err
		}

		return &ooopk, nil
	}
}

// Insert returns a Mutation to insert a row into a table. If the row already
// exists, the write or transaction fails.
func (ooopk *OutOfOrderPrimaryKey) Insert(ctx context.Context) *spanner.Mutation {
	values, _ := ooopk.columnsToValues(OutOfOrderPrimaryKeyWritableColumns())
	return spanner.Insert("OutOfOrderPrimaryKeys", OutOfOrderPrimaryKeyWritableColumns(), values)
}

// Delete deletes the OutOfOrderPrimaryKey from the database.
func (ooopk *OutOfOrderPrimaryKey) Delete(ctx context.Context) *spanner.Mutation {
	values, _ := ooopk.columnsToValues(OutOfOrderPrimaryKeyPrimaryKeys())
	return spanner.Delete("OutOfOrderPrimaryKeys", spanner.Key(values))
}
