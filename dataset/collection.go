package dataset

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
)

//CollectionDataset stored collection of datasets
type CollectionDataset struct {
	dsmap map[string]*Dataset
	//Status   Status
	Datasets []*Dataset `json:"Databases" xml:"Databases>Database"`
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

//ToJSON retruns as JSON
func (c CollectionDataset) ToJSON() string {

	j, err := json.Marshal(c)
	if err != nil {
		panic(err)
	}
	var out bytes.Buffer
	json.Indent(&out, j, "", "\t")
	return out.String()
}

func (c CollectionDataset) ToXml() string {
	output, err := xml.MarshalIndent(c, "  ", "    ")
	if err != nil {
		panic(err)
	}
	return string(output)
}
