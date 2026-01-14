package util

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

func ParseNumeric(s string) (pgtype.Numeric, error) {
	var n pgtype.Numeric
	if err := n.Scan(s); err != nil {
		return pgtype.Numeric{}, err
	}
	return n, nil
}

func ParseOptionalNumeric(s *string) (pgtype.Numeric, error) {
	if s == nil {
		return pgtype.Numeric{Valid: false}, nil
	}
	return ParseNumeric(*s)
}

func ParseTimestamptz(t time.Time) pgtype.Timestamptz {
	return pgtype.Timestamptz{
		Time:  t,
		Valid: true,
	}
}
