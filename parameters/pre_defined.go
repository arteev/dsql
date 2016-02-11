package parameters

type preDefinedParams struct {
	Name        string
	Default     string
	Description string
}

var definedParams = [...]preDefinedParams{
	{"QueryStatistic", "true", "Output summary statistics for query"},
	{"Statistic", "true", "Output summary statistics for query"},
	{"Silent", "false", "silent mode when error query"},
}
