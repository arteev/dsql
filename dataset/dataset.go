package dataset

//A Dataset for export to json,xml 
type Dataset struct {
	Error   bool
	TextError string `json:",omitempty" xml:",omitempty"`
	Name    string `xml:"name,attr"`
	Columns []*Column `xml:"columns>column"`
	Rows    []*Row `xml:"rows>row"`
}

//NewDataSet returns new dataset by name
func NewDataSet(name string) *Dataset {
	return &Dataset{
		Name: name,
	}
}

//AddColumn add column in dataset and returns new column
func (d *Dataset) AddColumn(name string) *Column {
	if col := d.findcolumn(name); col != nil {
		return col
	}
	col := &Column{
		Name: name,
	}
	d.Columns = append(d.Columns, col)
	return col
}

//AddColumns add columns in dataset and returns dataset
func (d *Dataset) AddColumns(names ...string) *Dataset {
	for _, s := range names {
		d.AddColumn(s)
	}
	return d
}

func (d Dataset) findcolumn(name string) *Column {
	for i, col := range d.Columns {
		if col.Name == name {
			return d.Columns[i]
		}
	}
	return nil
}

//GetColumnsNames returns slice string of names columns
func (d Dataset) GetColumnsNames() (result []string) {
	for _, col := range d.Columns {
		result = append(result, col.Name)
	}
	return
}

func (d Dataset) LastRow() *Row {
    l := len(d.Rows);
    if l==0 {
        return nil
    }
	return d.Rows[l-1]
}

func (d *Dataset) Append(data map[string]interface{}) {
	i := 0
	last := d.LastRow()
	if last != nil {
		i = last.Num + 1
	}
    r := &Row{
		Num:  i,		
	}
    r.SetDataValues(data)
	d.Rows = append(d.Rows, r)
}

//RowsCount returns count rows dataset
func (d Dataset) RowsCount() int {
	return len(d.Rows)
}

