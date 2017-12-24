package dataset

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"io"
)

//CollectionDataset stored collection of datasets
type CollectionDataset struct {
	dsmap map[string]*Dataset
	//Status   Status
	Datasets []*Dataset `json:"databases" xml:"databases>database"`
}

//NewColllection create new collection of datasets
func NewColllection() *CollectionDataset {
	return &CollectionDataset{
		dsmap:    make(map[string]*Dataset),
		Datasets: make([]*Dataset, 0),
		/*Status: Status{
		    Error:false,
		    Message:"",
		},*/
	}
}

//GetOrCreateDataset get or created new dataset in collection by name of dataset
func (c *CollectionDataset) GetOrCreateDataset(name string) *Dataset {
	if d, ok := c.dsmap[name]; ok {
		return d
	}
	ds := NewDataSet(name)
	c.dsmap[name] = ds
	c.Datasets = append(c.Datasets, ds)
	return ds
}

func (c CollectionDataset) GetUniqueColumnsNames() (result []string) {
	names := make(map[string]interface{})
	for _, ds := range c.Datasets {
		for _, col := range ds.Columns {
			if _, ok := names[col.Name]; !ok {
				names[col.Name] = nil
				result = append(result, col.Name)
			}
		}
	}
	return
}

//GetDatasets returns datasets as slice
func (c *CollectionDataset) GetDatasets() (result []*Dataset) {
	result = append(result, c.Datasets[:]...)
	return
}

//ToJSON returns as JSON
func (c CollectionDataset) ToJSON(indent string) string {
	var out bytes.Buffer
	_, err := c.WriteJSON(&out, indent)
	if err != nil {
		panic(err)
	}
	return out.String()
}

//ToXML returns as XML
func (c CollectionDataset) ToXML() string {
	var out bytes.Buffer
	_, err := c.WriteXML(&out, "    ")
	if err != nil {
		panic(err)
	}
	return out.String()
}

//WriteJSON write result in io.Writer
func (c CollectionDataset) WriteJSON(buf io.Writer, indent string) (int, error) {
	var j []byte
	var err error
	if indent != "" {
		j, err = json.MarshalIndent(c, "", indent)
	} else {
		j, err = json.Marshal(c)
	}
	if err != nil {
		return 0, err
	}
	return buf.Write(j)
}

//WriteXML write result in io.Writer
func (c CollectionDataset) WriteXML(buf io.Writer, indent string) (int64, error) {
	var (
		output []byte
		err    error
	)
	if indent != "" {
		output, err = xml.MarshalIndent(c, "", indent)
	} else {
		output, err = xml.Marshal(c)
	}
	if err != nil {
		return 0, err
	}
	n, err := buf.Write(output)
	return int64(n), err
}
