package rowgetter

import (
	"database/sql"
	"errors"
	"math/big"
)

//Errors
var (
	ErrNotBigRat = errors.New("src is not *big.Rat")
)

//BigCast - Cast big.Rat to float64
type BigCast struct {
	sql.NullFloat64
	sql.Scanner
}

//Scan - Scanner is an interface used by Scan.
func (b *BigCast) Scan(src interface{}) error {
	rat, ok := src.(*big.Rat)
	if !ok {
		return ErrNotBigRat
	}
	b.Float64, _ = rat.Float64()
	b.Valid = true
	return nil
}
