package typeconv

import (
	"database/sql"
	"time"
)

func NullBoolToBoolRef(v sql.NullBool) *bool {
	if v.Valid {
		return &v.Bool
	}
	return nil
}

func BoolRefToNullBool(p *bool) sql.NullBool {
	if p != nil {
		return sql.NullBool{Bool: *p, Valid: true}
	}
	return sql.NullBool{}
}

func NullByteToByteRef(v sql.NullByte) *byte {
	if v.Valid {
		return &v.Byte
	}
	return nil
}

func ByteRefToNullByte(p *byte) sql.NullByte {
	if p != nil {
		return sql.NullByte{Byte: *p, Valid: true}
	}
	return sql.NullByte{}
}

func NullStringToStringRef(v sql.NullString) *string {
	if v.Valid {
		return &v.String
	}
	return nil
}

func StringRefToNullString(p *string) sql.NullString {
	if p != nil {
		return sql.NullString{String: *p, Valid: true}
	}
	return sql.NullString{}
}

func NullInt16ToInt16Ref(v sql.NullInt16) *int16 {
	if v.Valid {
		return &v.Int16
	}
	return nil
}

func Int16RefToNullInt16(p *int16) sql.NullInt16 {
	if p != nil {
		return sql.NullInt16{Int16: *p, Valid: true}
	}
	return sql.NullInt16{}
}

func NullInt32ToInt32Ref(v sql.NullInt32) *int32 {
	if v.Valid {
		return &v.Int32
	}
	return nil
}

func Int32RefToNullInt32(p *int32) sql.NullInt32 {
	if p != nil {
		return sql.NullInt32{Int32: *p, Valid: true}
	}
	return sql.NullInt32{}
}

func NullInt64ToInt64Ref(v sql.NullInt64) *int64 {
	if v.Valid {
		return &v.Int64
	}
	return nil
}

func Int64RefToNullInt64(p *int64) sql.NullInt64 {
	if p != nil {
		return sql.NullInt64{Int64: *p, Valid: true}
	}
	return sql.NullInt64{}
}

func NullFloat64ToFloat64Ref(v sql.NullFloat64) *float64 {
	if v.Valid {
		return &v.Float64
	}
	return nil
}

func Float64RefToNullFloat64(p *float64) sql.NullFloat64 {
	if p != nil {
		return sql.NullFloat64{Float64: *p, Valid: true}
	}
	return sql.NullFloat64{}
}

func NullTimeToTimeRef(v sql.NullTime) *time.Time {
	if v.Valid {
		return &v.Time
	}
	return nil
}

func TimeRefToNullTime(p *time.Time) sql.NullTime {
	if p != nil {
		return sql.NullTime{Time: *p, Valid: true}
	}
	return sql.NullTime{}
}
