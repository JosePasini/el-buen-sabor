package database

import (
	"database/sql"
	"time"
)

//ToBoolP :: Normalización datos nulos
func ToBoolP(in sql.NullBool) *bool {
	if in.Valid {
		b := in.Bool
		return &b
	}
	return nil
}

//ToStringP :: Normalización datos nulos
func ToStringP(in sql.NullString) *string {
	if in.Valid {
		s := in.String
		return &s
	}
	return nil
}

//ToTimeP :: Normalización datos nulos
func ToTimeP(in sql.NullTime) *time.Time {
	if in.Valid {
		t := in.Time
		return &t
	}
	return nil
}

//ToIntP :: Normalización datos nulos
func ToIntP(in sql.NullInt32) *int {
	if in.Valid {
		i := int(in.Int32)
		return &i
	}
	return nil
}

//ToInt64P :: Normalización datos nulos
func ToInt64P(in sql.NullInt64) *int64 {
	if in.Valid {
		i := int64(in.Int64)
		return &i
	}
	return nil
}

//ToFloat64P :: Normalización datos nulos
func ToFloat64P(in sql.NullFloat64) *float64 {
	if in.Valid {
		f := in.Float64
		return &f
	}

	return nil
}
