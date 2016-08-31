package tabex

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/arteev/fmttab"
	"github.com/arteev/fmttab/columns"
)

const (
	sepGroups       = ";"
	sepParams       = ":"
	sepColumnParams = ","
)

func findTabColumns(table *fmttab.Table, column string) *columns.Column {
	for i, c := range table.Columns {
		if c.Name == column {
			return table.Columns[i]
		}
	}
	return nil
}

func getStrByIdx(params []string, idx int) (string, bool) {
	if len(params) <= idx {
		return "", false
	}
	return params[idx], true
}

func makeMapFromparams(params []string) (result map[string]string) {
	result = make(map[string]string)
	for _, p := range params {
		keypair := strings.SplitN(p, "=", 2)
		if len(keypair) < 2 {
			result[keypair[0]] = ""
		} else {
			result[keypair[0]] = keypair[1]
		}
	}
	return
}

func setTable(tab *fmttab.Table, params []string) {
	if tableParams, ok := getStrByIdx(params, 1); ok {
		mcp := makeMapFromparams(strings.Split(tableParams, sepColumnParams))
		if header, ok := mcp["header"]; ok {
			if header == "n" || header == "false" {
				tab.VisibleHeader = false
			} else if header == "y" || header == "true" {
				tab.VisibleHeader = true
			}
		}
		if border, ok := mcp["border"]; ok {
			switch strings.ToLower(border) {
			case "thin":
				tab.SetBorder(fmttab.BorderThin)
			case "double":
				tab.SetBorder(fmttab.BorderDouble)
			case "simple":
				tab.SetBorder(fmttab.BorderSimple)
			case "none", "false", "n":
				tab.SetBorder(fmttab.BorderNone)
			}
		}
	} //if
}

func setColumn(tab *fmttab.Table, params []string) {
	if columnParams, ok := getStrByIdx(params, 1); ok {
		mcp := makeMapFromparams(strings.Split(columnParams, sepColumnParams))
		if name, ok := mcp["name"]; ok {
			if column := findTabColumns(tab, name); column != nil {
				if align, ok := mcp["align"]; ok {
					switch align {
					case "left":
						column.Aling = fmttab.AlignLeft
					case "right":
						column.Aling = fmttab.AlignRight
					default:
						panic(fmt.Errorf("Unknow value: %s", align))
					}
				} //align
				if width, ok := mcp["width"]; ok {
					if width == "auto" {
						column.Width = fmttab.WidthAuto
					} else {
						iwidth, err := strconv.Atoi(width)
						if err != nil {
							panic(err)
						}
						column.Width = iwidth
					}
				}

				if visible, ok := mcp["visible"]; ok {
					if visible == "n" || visible == "false" {
						column.Visible = false
					} else if visible == "y" || visible == "true" {
						column.Visible = true
					}
				}

				if caption, ok := mcp["caption"]; ok {
					column.Caption = caption
				}

			}
		}
	} //if
}

//SetTableSubFormat set table options from subformat
func SetTableSubFormat(tab *fmttab.Table, subformat string) {
	// refactor this
	if subformat != "" {
		//column:name=string,width=auto|N,align=left|right,visible=y|n,caption=string;
		groups := strings.Split(subformat, sepGroups)
		for _, g := range groups {

			params := strings.Split(g, sepParams)
			if n, ok := getStrByIdx(params, 0); ok && (strings.Contains(n, "=") || n == "") {
				params = strings.Split("table:"+g, sepParams)
			}
			if subject, ok := getStrByIdx(params, 0); ok {
				switch subject {
				case "table":
					setTable(tab, params)
				case "column":
					setColumn(tab, params)
				} //switch
			}
		}
	}
}
