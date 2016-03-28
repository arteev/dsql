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

//ToJSON returns as JSON
func (c CollectionDataset) ToJSON() string {    
	var out bytes.Buffer
	_,err:=c.WriteJSON(&out);
    if err!=nil{
        panic(err)
    }
	return out.String()
}

//ToXML returns as XML
func (c CollectionDataset) ToXML() string {	
    var out bytes.Buffer
	_,err:=c.WriteXML(&out);
    if err!=nil{
        panic(err)
    }
	return out.String();
}

//WriteJSON write result in io.Writer
func (c CollectionDataset) WriteJSON(buf io.Writer) (int64, error) {
	j, err := json.Marshal(c)
	if err != nil {
		return 0,err
	}
	var out bytes.Buffer
	err = json.Indent(&out, j, "", "\t")
    if err!=nil {
        return 0,err
    }    
	return out.WriteTo(buf) 
}

//WriteXML write result in io.Writer
func (c CollectionDataset) WriteXML(buf io.Writer) (int64, error) {   
    output, err := xml.MarshalIndent(c, "  ", "    ")
	if err != nil {
		return 0,err
	}      
    n,err:=buf.Write(output)  	        
    return int64(n),err
}
