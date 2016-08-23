package rowgetter

import (
	"database/sql"
	"strconv"
	"strings"
	"time"

	"github.com/arteev/logger"
)

//A RowGetter wrap *sql.Rows with Columns. Fix when convert from time,big, etc.
type RowGetter struct {
	prows        *sql.Rows
	corrected    bool
	correct      map[int]string
	countColumns int
}

//MustRowGetter returns RowGetter
func MustRowGetter(r *sql.Rows) *RowGetter {
	result := &RowGetter{}
	result.prows = r
	result.correct = make(map[int]string)
	cols, e := result.prows.Columns()
	if e != nil {
		panic(e)
	}
	result.countColumns = len(cols)
	return result
}

func (r *RowGetter) createValue(index int, def *sql.NullString) interface{} {
	if stype, ok := r.correct[index]; ok {
		switch stype {
		case "time":
			return &time.Time{}
		case "big":
			return &BigCast{} // &big.Rat{}//&sql.NullFloat64{}

		default:
			return def
		}
	}

	return def
}

func (r *RowGetter) getTryValue(raw sql.NullString, val interface{}, column int) interface{} {
	if typecor, ok := r.correct[column]; ok {
		if typecor == "time" {
			t := val.(*time.Time)
			//logger.Trace.Println(t)
			if t.Year() == 0 {
				return t.Format("15:04:05")
			}
			return t.Format("2006-01-02 15:04:05")
		}
		if typecor == "big" {
			t, ok := val.(*BigCast)
			if !ok {
				panic("not float")
			}
			if !t.Valid {
				return nil
			}
			return strconv.FormatFloat(t.Float64, 'f', -1, 64)
			//return strings.TrimSpace(raw.String)

		}
	} else {
		t := val.(*sql.NullString)
		if !t.Valid {
			return nil
		}
		//todo: убрать trim
		return strings.TrimSpace(t.String)

	}
	return nil
}

func isErrorScanTime(e error) string {
	if e == nil {
		return ""
	}
	if strings.Contains(e.Error(), "unsupported driver -> Scan pair: time.Time -> *string") {
		return "time"
	}
	if strings.Contains(e.Error(), "unsupported driver -> Scan pair: *big.Rat -> *string") ||
		strings.Contains(e.Error(), "type *big.Rat into type") {
		return "big"
	}
	return ""
}

func getColumnErrorScanTime(e error) int {
	emsg := e.Error()
	idxStart := strings.Index(emsg, "index")
	idxEnd := strings.Index(emsg, ": unsupported")
	idx, ec := strconv.Atoi(strings.TrimSpace(emsg[idxStart+len("index") : idxEnd]))
	if ec == nil {
		return idx
	}
	return -1
}

func (r *RowGetter) createDestRaw() ([]interface{}, []sql.NullString) {
	destination := make([]interface{}, r.countColumns)
	resRaw := make([]sql.NullString, r.countColumns)
	for i := range resRaw {
		destination[i] = r.createValue(i, &resRaw[i])
	}
	return destination, resRaw
}

func (r *RowGetter) getDefaultTypeCorrect(cortype string) (string, interface{}, bool) {
	switch cortype {
	case "time":
		return "time", &time.Time{}, true
	case "big":
		return "big", &BigCast{}, true
	default:
	}
	return "", nil, false
}

//Next and scan from sql.Rows
func (r *RowGetter) Next() ([]interface{}, bool) {
	if r.prows.Next() {
		//prepare out arrays
		destination, resRaw := r.createDestRaw()
		for {
			err := r.prows.Scan(destination...)
			if err == nil {
				break
			}
			logger.Error.Println(err)
			cortype := isErrorScanTime(err)
			if r.corrected || cortype == "" {
				logger.Error.Println(err)
				break //for
			}

			idx := getColumnErrorScanTime(err)
			if idx != -1 {
				if ct, val, ok := r.getDefaultTypeCorrect(cortype); ok {
					r.correct[idx] = ct
					destination[idx] = val
				}
			}
		}
		r.corrected = r.corrected || len(r.correct) != 0
		res := make([]interface{}, r.countColumns)
		for i, raw := range resRaw {
			res[i] = r.getTryValue(raw, destination[i], i)
		}
		return res, true
	}
	return nil, false
}
