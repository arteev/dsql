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
	{"AutoFitWidthColumns", "true", "auto fit of width for table"},
    {"Fit", "true", "Use for fit table by width window of terminal"},
	{"BorderTable", "Thin", "None,Thin, Double"},
}
