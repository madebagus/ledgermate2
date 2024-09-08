package model

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/types"
	"github.com/lib/pq"
)

const (
	RECTypeEmailRecovery    = 1
	RECTypePhoneRecovery    = 2
	RECTypePasswordRecovery = 3
	RECTypeGARecovery       = 4
)

// SQLTimeFormat : Format string for golang to output SQL standar time
const SQLTimeFormat = "2006-01-02 15:04:05"
const SQLDateFormat = "2006-01-02"

// TimeRange :
type TimeRange struct {
	Valid bool
	Start time.Time
	End   time.Time
}

// RetSuccess :
type RetSuccess struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// AString :
type AString []string

// JSONMap :
type JSONMap map[string]interface{}

// SQLDB : sqlx db pointer
type SQLDB struct {
	DB []*sqlx.DB
}

// DBInit : initialize DB -> pointer is transfered into model from database (localize)
func DBInit(db []*sqlx.DB) SQLDB {
	return SQLDB{
		DB: db,
	}
}

// SRowCount : DB Row count
type SRowCount struct {
	Count int64 `db:"row_count"`
}

// MarshalJSON : overload convert SMember struct
func (a SRowCount) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Count int64 `json:"row_count"`
	}{
		Count: a.Count,
	})
}

/********************************************************
					Type of JSONB
********************************************************/

// JSONB : Type of JsonB
type JSONB []byte

// Value : get value of JSONB
func (j JSONB) Value() (driver.Value, error) {
	if j.IsNull() {
		//      log.Trace("returning null")
		return nil, nil
	}
	return string(j), nil
}

// Scan : scan JSONB
func (j *JSONB) Scan(value interface{}) error {
	if value == nil {
		*j = nil
		return nil
	}
	s, ok := value.([]byte)
	if !ok {
		return errors.New("Scan source was not string")
	}
	// I think I need to make a copy of the bytes.
	// It seems the byte slice passed in is re-used
	*j = append((*j)[0:0], s...)

	return nil
}

// MarshalJSON : returns *m as the JSON encoding of m.
func (j JSONB) MarshalJSON() ([]byte, error) {
	if j == nil {
		return []byte("null"), nil
	}
	return j, nil
}

// UnmarshalJSON :  sets *m to a copy of data.
func (j *JSONB) UnmarshalJSON(data []byte) error {
	if j == nil {
		return errors.New("json.RawMessage: UnmarshalJSON on nil pointer")
	}
	*j = append((*j)[0:0], data...)
	return nil
}

// IsNull : is JSONB null
func (j JSONB) IsNull() bool {
	return len(j) == 0 || string(j) == "null"
}

// Equals : Compare JSONB
func (j JSONB) Equals(j1 JSONB) bool {
	return bytes.Equal([]byte(j), []byte(j1))
}

// NullInt64 is an alias for sql.NullInt64 data type
type NullInt64 struct {
	sql.NullInt64
}

// MarshalJSON for NullInt64
func (ni *NullInt64) MarshalJSON() ([]byte, error) {
	if !ni.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ni.Int64)
}

// UnmarshalJSON for NullInt64
// func (ni *NullInt64) UnmarshalJSON(b []byte) error {
// 	err := json.Unmarshal(b, &ni.Int64)
// 	ni.Valid = (err == nil)
// 	return err
// }

// NullBool is an alias for sql.NullBool data type
type NullBool struct {
	sql.NullBool
}

// MarshalJSON for NullBool
func (nb *NullBool) MarshalJSON() ([]byte, error) {
	if !nb.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(nb.Bool)
}

// UnmarshalJSON for NullBool
// func (nb *NullBool) UnmarshalJSON(b []byte) error {
// 	err := json.Unmarshal(b, &nb.Bool)
// 	nb.Valid = (err == nil)
// 	return err
// }

// NullFloat64 is an alias for sql.NullFloat64 data type
type NullFloat64 struct {
	sql.NullFloat64
}

// MarshalJSON for NullFloat64
func (nf *NullFloat64) MarshalJSON() ([]byte, error) {
	if !nf.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(nf.Float64)
}

// UnmarshalJSON for NullFloat64
// func (nf *NullFloat64) UnmarshalJSON(b []byte) error {
// 	err := json.Unmarshal(b, &nf.Float64)
// 	nf.Valid = (err == nil)
// 	return err
// }

// NullString is an alias for sql.NullString data type
type NullString struct {
	sql.NullString
}

// MarshalJSON for NullString
func (ns *NullString) MarshalJSON() ([]byte, error) {
	if !ns.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ns.String)
}

// UnmarshalJSON for NullString
// func (ns *NullString) UnmarshalJSON(b []byte) error {
// 	err := json.Unmarshal(b, &ns.String)
// 	ns.Valid = (err == nil)
// 	return err
// }

// PQNullTime is an alias for pq.NullTime data type
type PQNullTime struct {
	pq.NullTime
}

// MarshalJSON for NullTime
func (nt *PQNullTime) MarshalJSON() ([]byte, error) {
	if !nt.Valid {
		return []byte("null"), nil
	}
	val := fmt.Sprintf("\"%s\"", nt.Time.Format(time.RFC3339))
	return []byte(val), nil
}

// UnmarshalJSON for NullTime
// func (nt *NullTime) UnmarshalJSON(b []byte) error {
// 	err := json.Unmarshal(b, &nt.Time)
// 	nt.Valid = (err == nil)
// 	return err
// }

// JSONTags :
type JSONTags []string

// Scan :
func (tags *JSONTags) Scan(src interface{}) error {
	var jt types.JSONText

	if err := jt.Scan(src); err != nil {
		return err
	}

	if err := jt.Unmarshal(tags); err != nil {
		return err
	}

	return nil
}

// Value :
func (tags *JSONTags) Value() (driver.Value, error) {
	var jt types.JSONText

	data, err := json.Marshal((*[]string)(tags))
	if err != nil {
		return nil, err
	}

	if err := jt.UnmarshalJSON(data); err != nil {
		return nil, err
	}

	return jt.Value()
}

// MarshalJSON :
func (tags *JSONTags) MarshalJSON() ([]byte, error) {
	return json.Marshal((*[]string)(tags))
}

// UnmarshalJSON :
func (tags *JSONTags) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*[]string)(tags)); err != nil {
		return err
	}

	return nil
}
