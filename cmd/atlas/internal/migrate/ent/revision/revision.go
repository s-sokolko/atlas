// Copyright 2021-present The Atlas Authors. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package revision

import (
	"github.com/s-sokolko/atlas/sql/migrate"
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the revision type in the database.
	Label = "revision"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "version"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldApplied holds the string denoting the applied field in the database.
	FieldApplied = "applied"
	// FieldTotal holds the string denoting the total field in the database.
	FieldTotal = "total"
	// FieldExecutedAt holds the string denoting the executed_at field in the database.
	FieldExecutedAt = "executed_at"
	// FieldExecutionTime holds the string denoting the execution_time field in the database.
	FieldExecutionTime = "execution_time"
	// FieldError holds the string denoting the error field in the database.
	FieldError = "error"
	// FieldErrorStmt holds the string denoting the error_stmt field in the database.
	FieldErrorStmt = "error_stmt"
	// FieldHash holds the string denoting the hash field in the database.
	FieldHash = "hash"
	// FieldPartialHashes holds the string denoting the partial_hashes field in the database.
	FieldPartialHashes = "partial_hashes"
	// FieldOperatorVersion holds the string denoting the operator_version field in the database.
	FieldOperatorVersion = "operator_version"
	// Table holds the table name of the revision in the database.
	Table = "atlas_schema_revisions"
)

// Columns holds all SQL columns for revision fields.
var Columns = []string{
	FieldID,
	FieldDescription,
	FieldType,
	FieldApplied,
	FieldTotal,
	FieldExecutedAt,
	FieldExecutionTime,
	FieldError,
	FieldErrorStmt,
	FieldHash,
	FieldPartialHashes,
	FieldOperatorVersion,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultType holds the default value on creation for the "type" field.
	DefaultType migrate.RevisionType
	// DefaultApplied holds the default value on creation for the "applied" field.
	DefaultApplied int
	// AppliedValidator is a validator for the "applied" field. It is called by the builders before save.
	AppliedValidator func(int) error
	// DefaultTotal holds the default value on creation for the "total" field.
	DefaultTotal int
	// TotalValidator is a validator for the "total" field. It is called by the builders before save.
	TotalValidator func(int) error
)

// OrderOption defines the ordering options for the Revision queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// ByApplied orders the results by the applied field.
func ByApplied(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldApplied, opts...).ToFunc()
}

// ByTotal orders the results by the total field.
func ByTotal(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTotal, opts...).ToFunc()
}

// ByExecutedAt orders the results by the executed_at field.
func ByExecutedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExecutedAt, opts...).ToFunc()
}

// ByExecutionTime orders the results by the execution_time field.
func ByExecutionTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExecutionTime, opts...).ToFunc()
}

// ByError orders the results by the error field.
func ByError(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldError, opts...).ToFunc()
}

// ByErrorStmt orders the results by the error_stmt field.
func ByErrorStmt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldErrorStmt, opts...).ToFunc()
}

// ByHash orders the results by the hash field.
func ByHash(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHash, opts...).ToFunc()
}

// ByOperatorVersion orders the results by the operator_version field.
func ByOperatorVersion(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOperatorVersion, opts...).ToFunc()
}
