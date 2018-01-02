package commands

import (
	"errors"
	"fmt"
	"os"

	"github.com/arteev/dsql/parameters"
	"github.com/arteev/dsql/parameters/parametergetter"
	"github.com/arteev/fmttab"
	"github.com/arteev/logger"
	"github.com/nsf/termbox-go"
	"github.com/urfave/cli"
)

//Errors
var (
	ErrParamRemovePreDefined = errors.New("Can not delete predefined parameter")
)

func checkErr(fn func() error) {
	err := fn()
	if err != nil {
		panic(err)
	}
}
func listParams() cli.Command {
	return cli.Command{
		Name:  "list",
		Usage: "list of parametrs",
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:  "fit",
				Usage: "use for fit table by width window of terminal",
			},
			cli.StringFlag{
				Name:  "border",
				Usage: "set type of border table: Thin,Double,Simple or None. Default:Thin",
			},
		},
		Action: func(ctx *cli.Context) {
			logger.Trace.Println("command params list")
			defer logger.Trace.Println("command params list done")
			d := parameters.GetInstance()
			defer checkErr(d.Close)
			params, err := d.All()
			if err != nil {
				panic(err)
			}
			tab := fmttab.New("List of databases", fmttab.BorderThin, nil)

			tab.AddColumn("Id", 4, fmttab.AlignRight)
			tab.AddColumn("Name", fmttab.WidthAuto, fmttab.AlignLeft)
			tab.AddColumn("Value", fmttab.WidthAuto, fmttab.AlignLeft)
			tab.AddColumn("Description", fmttab.WidthAuto, fmttab.AlignLeft)
			for _, curd := range params.Get() {
				rec := make(map[string]interface{})
				rec["Id"] = curd.ID
				rec["Name"] = curd.Name
				rec["Value"] = curd.ValueStr()
				rec["Description"] = curd.Description
				tab.AppendData(rec)
			}
			pget := parametergetter.New(ctx, parameters.GetInstance())
			if pget.GetDef(parametergetter.Fit, false).(bool) {
				if e := termbox.Init(); e != nil {
					panic(e)
				}
				tw, _ := termbox.Size()
				tab.AutoSize(true, tw)
				termbox.Close()
			}
			switch pget.GetDef(parametergetter.BorderTable, "").(string) {
			case "Thin":
				tab.SetBorder(fmttab.BorderThin)
			case "Double":
				tab.SetBorder(fmttab.BorderDouble)
			case "None":
				tab.SetBorder(fmttab.BorderNone)
			case "Simple":
				tab.SetBorder(fmttab.BorderSimple)
			}

			if _, err := tab.WriteTo(os.Stdout); err != nil {
				panic(err)
			}
		},
	}
}

func deleteParam() cli.Command {
	return cli.Command{
		Name:  "remove",
		Usage: "remove parameter",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "name",
				Usage: "name of a parameter",
			},
		},
		Action: func(ctx *cli.Context) {
			logger.Trace.Println("command params remove")
			defer logger.Trace.Println("command params remove done")
			name := ctx.String("name")
			if !ctx.IsSet("name") || name == "" {
				panic(fmt.Errorf("option name must be set"))
			}
			logger.Debug.Printf("remove parameter %s", name)
			d := parameters.GetInstance()
			defer checkErr(d.Close)
			param, err := d.FindByName(name)
			if err != nil {
				panic(err)
			}
			if param.IsPreDefined() {
				panic(ErrParamRemovePreDefined)
			}
			if err := d.Delete(param); err != nil {
				panic(err)
			}
			logger.Info.Printf("Parameter removed: %s\n", param)
		},
	}
}

func getParam() cli.Command {
	return cli.Command{
		Name:  "get",
		Usage: "get parameter by name",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "name",
				Usage: "name of a parameter",
			},
			cli.BoolFlag{
				Name:  "full,f",
				Usage: "use full output with name,value,description",
			},
		},
		Action: func(ctx *cli.Context) {
			logger.Trace.Println("command params get")
			defer logger.Trace.Println("command params get done")
			if !ctx.IsSet("name") || ctx.String("name") == "" {
				panic(fmt.Errorf("option name must be set"))
			}
			name := ctx.String("name")
			logger.Debug.Printf("get parameter:%q\n", name)
			d := parameters.GetInstance()
			defer checkErr(d.Close)

			param, err := d.FindByName(name)
			if err != nil {
				panic(err)
			}
			logger.Debug.Printf("parameter value:%q\n", param.ValueStr())
			if ctx.Bool("full") || ctx.Bool("f") {
				fmt.Printf("parameter:%s\nvalue:%s\ndescription:%s\n", param.Name, param.ValueStr(), param.Description)
			} else {
				fmt.Println(param.ValueStr())
			}
		},
	}
}

func setParam() cli.Command {
	return cli.Command{
		Name:  "set",
		Usage: "set parameter",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "name",
				Usage: "name of a parameter",
			},
			cli.StringFlag{
				Name:  "value",
				Usage: "name of a parameter",
			},
			cli.StringFlag{
				Name:  "desc",
				Usage: "description of a parameter",
			},
		},
		Action: func(ctx *cli.Context) {
			logger.Trace.Println("command params set")
			defer logger.Trace.Println("command params set done")
			if !ctx.IsSet("name") || ctx.String("name") == "" {
				panic(fmt.Errorf("option name must be set"))
			}

			name := ctx.String("name")
			value := ctx.String("value")
			description := ctx.String("desc")
			logger.Debug.Printf("set parameter:%q value:%q description=%q\n", name, value, description)
			d := parameters.GetInstance()
			defer checkErr(d.Close)

			param, err := d.FindByName(name)
			if err == parameters.ErrNotFound {
				param = parameters.MustParameter(name, value, description)
				if err := d.Add(param); err != nil {
					panic(err)
				}
			} else if err != nil {
				panic(err)
			}
			if ctx.IsSet("value") {
				param.Value = value
			}
			if ctx.IsSet("desc") && !param.IsPreDefined() {
				param.Description = description
			}
			if err := d.Update(param); err != nil {
				panic(err)
			}

			logger.Info.Printf("Parameter set: %s\n", param)
		},
	}
}

func init() {
	Register(
		[]cli.Command{
			cli.Command{
				Name:  "param",
				Usage: "list or managment of the parameters",
				Subcommands: []cli.Command{
					listParams(),
					setParam(),
					getParam(),
					deleteParam(),
				},
			},
		})
}
