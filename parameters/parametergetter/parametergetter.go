package parametergetter

import (
	"github.com/arteev/dsql/parameters"
	"github.com/codegangsta/cli"
)

type parameterGetterMixed struct {
	context *cli.Context
	params  parameters.RepositoryParams
}

//New ParameterGetter (ParameterGetterMixed). Get param from cli app and repository
func New(ctx *cli.Context, params parameters.RepositoryParams) ParameterGetter {
	return &parameterGetterMixed{
		context: ctx,
		params:  params,
	}
}

func (p *parameterGetterMixed) Get(name string) interface{} {

	GetBool := func() interface{} {

		repValue, e := p.params.FindByName(name)
		if e != nil {
			return nil
		}
		return repValue.Bool()
	}

	switch name {
	case Statistic:
		if p.context.GlobalIsSet("stat") {
			return p.context.GlobalBool("stat")
		} else if p.context.GlobalIsSet("s") {
			return p.context.GlobalBool("s")
		}
		return GetBool()
	case QueryStatistic:
		if p.context.GlobalIsSet("statquery") {
			return p.context.GlobalBool("statquery")
		} else if p.context.GlobalIsSet("sq") {
			return p.context.GlobalBool("sq")
		}
		return GetBool()
	case Silent:
		if p.context.GlobalIsSet("silent") {
			return p.context.GlobalBool("silent")
		}
		return GetBool()
	case AutoFitWidthColumns:
		if p.context.IsSet("fit") {
			return p.context.Bool("fit")
		}
	case BorderTable:
		if p.context.IsSet("border") {
			return p.context.String("border")
		}
		if repValue, e := p.params.FindByName(name); e == nil {          
            return repValue.ValueStr()
        }
	}

	return nil
}

func (p *parameterGetterMixed) GetDef(name string, def interface{}) interface{} {
	val := p.Get(name)
	if val != nil {
		return val
	}
	return def
}
