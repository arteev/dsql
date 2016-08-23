package commands

import (
	"github.com/arteev/dsql/cmd/db"
	"github.com/codegangsta/cli"
)

type cliFlagMode byte

const (
	modeFlagUnUsed = cliFlagMode(0)
	modeFlagSingle = cliFlagMode(1)
	modeFlagMulti  = cliFlagMode(2)
)

type cliOption struct {
	Databases        cliFlagMode
	ExcludeDatabases cliFlagMode
	Engines          cliFlagMode
	Tags             cliFlagMode
	ExcludeTags      cliFlagMode
}

type cliFlags struct {
	opt   cliOption
	ctx   *cli.Context
	flags []cli.Flag
}

func newCliFlags(opt cliOption) *cliFlags {
	result := &cliFlags{
		opt: opt,
	}
	result.flags = result.genFlags()
	return result
}

//SetContext use before get Database,Engines,Tags
func (f *cliFlags) SetContext(ctx *cli.Context) {
	f.ctx = ctx
}

func getFlagByMode(mode cliFlagMode, Name, Usage string) cli.Flag {
	switch mode {
	case modeFlagUnUsed:
		return nil
	case modeFlagSingle:
		return &cli.StringFlag{
			Name:  Name,
			Usage: Usage,
		}
	case modeFlagMulti:
		return &cli.StringSliceFlag{
			Name:  Name,
			Usage: Usage,
		}
	}
	return nil
}

func (f *cliFlags) genFlags() (flags []cli.Flag) {
	df := getFlagByMode(f.opt.Databases, "databases,d", "use for the concrete database(s)")
	if df != nil {
		flags = append(flags, df)
	}

	nd := getFlagByMode(f.opt.ExcludeDatabases, "nd", "Exclude database(s)")
	if nd != nil {
		flags = append(flags, nd)
	}

	ef := getFlagByMode(f.opt.Engines, "engine,e", "user for the concrete engine(s)")
	if ef != nil {
		flags = append(flags, ef)
	}
	tf := getFlagByMode(f.opt.Tags, "tag,t", "use filter by tag")
	if tf != nil {
		flags = append(flags, tf)
	}
	te := getFlagByMode(f.opt.ExcludeTags, "nt", "exclude tag(s)")
	if te != nil {
		flags = append(flags, te)
	}
	return
}

//Flags returns flags for cli.app
func (f *cliFlags) Flags() []cli.Flag {
	return f.flags
}

func (f cliFlags) checkContext() {
	if f.ctx == nil {
		panic("CliFlags: cli.context is nil")
	}

}

//Databases returns list of databases from cli flags
func (f *cliFlags) Databases() []string {
	return f.exclude(
		f.getvalue(f.opt.Databases, "databases", "d"),
		f.getvalue(f.opt.ExcludeDatabases, "nd", ""))
}
//ExDatabases returns excluded list of databases from cli flags
func (f *cliFlags) ExDatabases() []string {
	return f.getvalue(f.opt.ExcludeDatabases, "nd", "")
}

//Engines returns list of engines from cli flags
func (f *cliFlags) Engines() []string {
	return f.getvalue(f.opt.Engines, "engines", "e")
}

//Tags returns list of tags from cli flags
func (f *cliFlags) Tags() []string {
	return f.exclude(f.getvalue(f.opt.Tags, "tags", "t"),
		f.getvalue(f.opt.ExcludeTags, "nt", ""))

}
//ExTags returns excluded list of tags from cli flags
func (f *cliFlags) ExTags() []string {
	return f.getvalue(f.opt.Engines, "nt", "")
}


func (f *cliFlags) getvalue(mode cliFlagMode, ltr1, ltr2 string) (res []string) {
	f.checkContext()
	switch mode {
	case modeFlagUnUsed:
		return
	case modeFlagSingle:
		if ltr1 != "" && f.ctx.IsSet(ltr1) {
			res = append(res, f.ctx.String(ltr1))
		}
		if ltr2 != "" && f.ctx.IsSet(ltr2) {
			res = append(res, f.ctx.String(ltr2))
		}
		return
	case modeFlagMulti:
		if ltr1 != "" && f.ctx.IsSet(ltr1) {
			res = f.ctx.StringSlice(ltr1)
		}
		if ltr2 != "" && f.ctx.IsSet(ltr2) {
			res = f.ctx.StringSlice(ltr2)
		}
		return
	}
	return
}

func (f *cliFlags) exclude(value []string, exvalue []string) (result []string) {
	exists := func(chkval string) bool {
		for _, s := range exvalue {
			if s == chkval {
				return true
			}
		}
		return false
	}
	for _, v := range value {
		if !exists(v) {
			result = append(result, v)
		}
	}
	return
}

func (f *cliFlags) ApplyTo(dbs db.CollectionRepositoryDB) {
	dbs.AddFilterIncludeDB(f.Databases()...).
		AddFilterExcludeDB(f.ExDatabases()...).
		AddFilterIncludeEngine(f.Engines()...).
		AddFilterTag(f.Tags()...).
        AddFilterExcludeTag(f.ExTags()...)
}
