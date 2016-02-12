package commands

import (
	"fmt"
	"os"

	"github.com/arteev/dsql/db"
	"github.com/arteev/dsql/rdb"
	"github.com/arteev/fmttab"
	"github.com/arteev/logger"
	"github.com/codegangsta/cli"
)

func stringFlag(name, usage string) cli.Flag {
	return cli.StringFlag{
		Name:  name,
		Usage: usage,
	}
}

func listDatabase() cli.Command {

	dbFilterFlags := newCliFlags(cliOption{
		Databases:        modeFlagMulti,
		ExcludeDatabases: modeFlagMulti,
		Engines:          modeFlagMulti,
		Tags:             modeFlagMulti,
		ExcludeTags:      modeFlagMulti,
	})

	return cli.Command{
		Name:  "list",
		Usage: "list of databases",
		Flags: dbFilterFlags.Flags(),
		Action: func(ctx *cli.Context) {
			logger.Trace.Println("command list database")
			dbFilterFlags.SetContext(ctx)
			d := db.GetInstance()
			dbs, err := d.All()
			if err != nil {
				panic(err)
			}			
			for _, e := range dbFilterFlags.Engines() {
				rdb.CheckCodeEngine(e)
			}
			dbFilterFlags.ApplyTo(dbs)

			tab := fmttab.New("List of databases", fmttab.BorderThin, nil)
			tab.AddColumn("Id", 4, fmttab.AlignRight)
			tab.AddColumn("On", 2, fmttab.AlignLeft)
			tab.AddColumn("Code", 10, fmttab.AlignLeft)
			tab.AddColumn("Engine", 11, fmttab.AlignLeft)
			tab.AddColumn("URI", 40, fmttab.AlignLeft)
			tab.AddColumn("Tags", 25, fmttab.AlignLeft)
			for _, curd := range dbs.Get() {
				rec := make(map[string]interface{})
				rec["Id"] = curd.ID
				if curd.Enabled {
					rec["On"] = "+"
				}
				rec["Code"] = curd.Code
				rec["URI"] = curd.ConnectionString
				rec["Engine"] = curd.Engine
				rec["Tags"] = curd.TagsComma(";")
				tab.AppendData(rec)
			}
			_, err = tab.WriteTo(os.Stdout)
			if err != nil {
				panic(err)
			}
		},
	}
}

func tagDatabase() cli.Command {
	dbFilterFlags := newCliFlags(cliOption{
		Databases:        modeFlagMulti,
		ExcludeDatabases: modeFlagMulti,
		Engines:          modeFlagMulti,
		Tags:             modeFlagUnUsed,
		ExcludeTags:      modeFlagUnUsed,
	})

	flags := dbFilterFlags.Flags()
	flags = append(flags, cli.StringSliceFlag{
		Name:  "add",
		Usage: "new tag(s)",
	})
	flags = append(flags, cli.StringSliceFlag{
		Name:  "remove",
		Usage: "remove tag(s)",
	})

	return cli.Command{
		Name:  "tag",
		Usage: "add or remove tag for database",
		Flags: flags,
		Action: func(ctx *cli.Context) {
			logger.Trace.Println("command db tag")
			defer logger.Trace.Println("command db tag done")

			var add, remove = ctx.StringSlice("add"), ctx.StringSlice("remove")
			if len(add) == 0 && len(remove) == 0 {
				panic(fmt.Errorf("must be set: new tag or del tag"))
			}
			dbFilterFlags.SetContext(ctx)

			logger.Debug.Printf("updating new:%s remove:%s\n", add, remove)

			d := db.GetInstance()
			col, err := d.All()
			if err != nil {
				panic(err)
			}
			for _, e := range dbFilterFlags.Engines() {
				rdb.CheckCodeEngine(e)
			}
			dbFilterFlags.ApplyTo(col)

			dbs := col.Get()
			if len(dbs) == 0 {
				panic("databases not found")
			}

			for _, curdb := range dbs {
				logger.Trace.Printf("process tag: %q\n", curdb.Code)

				if err := d.AddTags(&curdb, add...); err != nil {
					panic(err)
				}

				if err := d.RemoveTags(&curdb, remove...); err != nil {
					panic(err)
				}
			}

		},
	}
}
func addDatabase() cli.Command {
	return cli.Command{
		Name:  "add",
		Usage: "Add new database",
		Flags: []cli.Flag{
			stringFlag("code", ""),
			stringFlag("uri", ""),
			stringFlag("engine", ""),
		},
		Action: func(ctx *cli.Context) {
			logger.Trace.Println("command db add")
			for _, flag := range ctx.FlagNames() {
				if !ctx.IsSet(flag) {
					panic(fmt.Errorf("option %q must be set", flag))
				}
			}
			d := db.GetInstance()
			engine := ctx.String("engine")
			rdb.CheckCodeEngine(engine)
			newdb := db.Database{
				Code:             ctx.String("code"),
				ConnectionString: ctx.String("uri"),
				Enabled:          true,
				Engine:           engine,
			}
			logger.Debug.Println("Adding ", newdb.Code, newdb.ConnectionString)
			err := d.Add(newdb)
			if err != nil {
				panic(err)
			}
			logger.Info.Println("Added ", newdb.Code)

		},
	}

}

func updateDatabase() cli.Command {
	return cli.Command{
		Name:  "update",
		Usage: "Update database",
		Flags: []cli.Flag{
			stringFlag("code", ""),
			stringFlag("newcode", ""),
			stringFlag("uri", ""),
			stringFlag("engine", ""),
			cli.BoolFlag{
				Name:  "enabled",
				Usage: "enabled or disable database",
			},
		},
		Action: func(ctx *cli.Context) {
			logger.Trace.Println("command db update")
			if !ctx.IsSet("code") {
				panic(fmt.Errorf("option code must be set"))
			}
			code := ctx.String("code")
			logger.Debug.Printf("updating %s, new values(code:%s; uri:%s; enabled:%v; engine:%v)\n", code, ctx.String("code"), ctx.String("uri"), ctx.Bool("enabled"), ctx.String("engine"))
			d := db.GetInstance()
			dbFind, err := d.FindByCode(code)
			logger.Debug.Println(dbFind)
			if err != nil {
				panic(err)
			}
			if ctx.IsSet("newcode") {
				dbFind.Code = ctx.String("newcode")
			}
			if ctx.IsSet("uri") {
				dbFind.ConnectionString = ctx.String("uri")
			}
			if ctx.IsSet("enabled") {
				dbFind.Enabled = ctx.Bool("enabled")
			}

			if ctx.IsSet("engine") {
				dbFind.Engine = ctx.String("engine")
				rdb.CheckCodeEngine(dbFind.Engine)
			}
			if err := d.Update(dbFind); err != nil {
				panic(err)
			}

			logger.Info.Println("updated ", code)

		},
	}

}

func deleteDatabase() cli.Command {
	return cli.Command{
		Name:  "delete",
		Usage: "Delete database by code",
		Flags: []cli.Flag{
			stringFlag("code", ""),
		},
		Action: func(ctx *cli.Context) {
			logger.Trace.Println("command database delete")
			if !ctx.IsSet("code") {
				panic(fmt.Errorf("option code must be set"))
			}
			code := ctx.String("code")
			logger.Debug.Printf("database deleting %q\n", code)
			d := db.GetInstance()
			dbfind, err := d.FindByCode(code)
			if err != nil {
				panic(err)
			}
			if err := d.Delete(dbfind); err != nil {
				panic(err)
			}
			logger.Info.Printf("database %q deleted\n", code)
		},
	}
}

//GetCommandsListDB define cli command for DB
func GetCommandsListDB() []cli.Command {
	return []cli.Command{
		cli.Command{
			Name:  "db",
			Usage: "list or managment of the list remote databases",
			Subcommands: []cli.Command{
				listDatabase(),
				addDatabase(),
				updateDatabase(),
				deleteDatabase(),
				tagDatabase(),
			},
		},
	}
}
