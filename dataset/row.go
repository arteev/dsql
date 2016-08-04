package dataset

type Row struct {
	Num     int        `xml:"num,attr"`
	DataRow []*DataRow `xml:"data>value"`
}

type DataRow struct {
	Column string      `xml:"column,attr"`
	Value  interface{} `xml:",chardata"`
}

func (r *Row) SetDataValues(data map[string]interface{}) {
	r.DataRow = make([]*DataRow, 0)
	for key, pair := range data {
		r.DataRow = append(r.DataRow, &DataRow{
			Column: key,
			Value:  pair,
		})
	}
}

func (r *Row) GetDataMap() (result map[string]interface{}) {
	result = make(map[string]interface{})
	for _, d := range r.DataRow {
		result[d.Column] = d.Value
	}
	return
}
