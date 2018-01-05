package tabex

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/arteev/dsql/format"
	"github.com/arteev/dsql/rdb/action"
	"github.com/arteev/fmttab"
	"github.com/arteev/fmttab/columns"
)

func findTabColumns(table *fmttab.Table, column string) *columns.Column {
	for i, c := range table.Columns {
		if c.Name == column {
			return table.Columns[i]
		}
	}
	return nil
}

func formatTab(ctx *action.Context, tab *fmttab.Table) {
	f := ctx.Get("format").(*format.Format)
	if header, ok := f.Root().Get("header"); ok {
		if header == "n" || header == "false" {
			tab.VisibleHeader = false
		} else if header == "y" || header == "true" {
			tab.VisibleHeader = true
		}
	}
	if border, ok := f.Root().Get("border"); ok {
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
}

func formatCols(ctx *action.Context, tab *fmttab.Table) error {
	f := ctx.Get("format").(*format.Format)

	//table:columnN:name=string,width=auto|N,align=left|right,visible=y|n,caption=string;

	n := 1
	for {
		gname := "column" + strconv.Itoa(n)
		gr, exist := f.Group(gname)
		if !exist {
			break
		}
		if columnName, ok := gr.Get("name"); ok {
			column := findTabColumns(tab, columnName)
			if column == nil {
				return fmt.Errorf("Unknow column: %s", columnName)
			}
			if align, ok := gr.Get("align"); ok {
				switch align {
				case "left":
					column.Aling = fmttab.AlignLeft
				case "right":
					column.Aling = fmttab.AlignRight
				default:
					return fmt.Errorf("Unknow value: %s", align)
				}
			} //align

			if width, ok := gr.Get("width"); ok {
				if width == "auto" {
					column.Width = fmttab.WidthAuto
				} else {
					iwidth, err := strconv.Atoi(width)
					if err != nil {
						return err
					}
					column.Width = iwidth
				}
			}
			if visible, ok := gr.Get("visible"); ok {
				if visible == "n" || visible == "false" {
					column.Visible = false
				} else if visible == "y" || visible == "true" {
					column.Visible = true
				}
			}
			if caption, ok := gr.Get("caption"); ok {
				column.Caption = caption
			}
		}
		n++
	}
	return nil
}

//SetTableSubFormat set table options from subformat
func SetTableSubFormat(ctx *action.Context, tab *fmttab.Table) error {
	//table:header=true|y|false|n,border=thin|double|simple|none|false|n
	formatTab(ctx, tab)
	//table:column:name=string,width=auto|N,align=left|right,visible=y|n,caption=string;
	err := formatCols(ctx, tab)
	return err
}
