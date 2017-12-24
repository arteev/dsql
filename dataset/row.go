package dataset

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
)

type Row struct {
	Num     int        `json:"num" xml:"num,attr"`
	DataRow []*DataRow `json:"data" xml:"data>value"`
}

type DataRow struct {
	Column string      `json:"-" xml:"column,attr"`
	Value  interface{} `json:"-" xml:",chardata"`
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

func (r *Row) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString("{")
	fmt.Fprintf(buffer, "\"num\":%d,", r.Num)
	buffer.WriteString("\"data\":{")
	corr := false
	for i, d := range r.DataRow {

		if d.Column == "_CODE_" {
			corr = true
			continue
		}
		jsonVal, err := json.Marshal(d.Value)
		if err != nil {
			return nil, err
		}
		fmt.Fprintf(buffer, "\"%s\":%s", d.Column, string(jsonVal))
		if i < len(r.DataRow)-1 {
			if (i < len(r.DataRow)-2) || (i == len(r.DataRow)-2 && corr) {
				buffer.WriteString(",")
			}
		}

	}
	buffer.WriteString("}}")
	return buffer.Bytes(), nil
}
