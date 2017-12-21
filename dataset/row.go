package dataset

import (
	"encoding/xml"
)

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

func (d *DataRow) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if d.Column == "_CODE_" {
		return nil
	}
	start.Attr = []xml.Attr{xml.Attr{Name: xml.Name{Local: "column"}, Value: d.Column}}
	e.EncodeToken(start)
	e.EncodeToken(xml.CharData(d.Value.(string)))
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
