// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package clientdb

import (
	"database/sql"
)

type Log struct {
	ID         int64
	TraceID    sql.NullString
	SpanID     sql.NullString
	Timestamp  int64
	Severity   int64
	Body       []byte
	Attributes []byte
}

type Span struct {
	ID                     int64
	TraceID                string
	SpanID                 string
	TraceState             string
	ParentSpanID           sql.NullString
	Flags                  int64
	Name                   string
	Kind                   string
	StartTime              int64
	EndTime                sql.NullInt64
	Attributes             []byte
	DroppedAttributesCount int64
	Events                 []byte
	DroppedEventsCount     int64
	Links                  []byte
	DroppedLinksCount      int64
	StatusCode             int64
	StatusMessage          string
	InstrumentationScope   []byte
	Resource               []byte
}