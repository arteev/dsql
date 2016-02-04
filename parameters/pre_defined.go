package parameters

type preDefinedParams struct {
	Name        string
	Default     string
	Description string
}

var definedParams = []preDefinedParams{
	{
		"QueryStatistic", "true", "Output summary statistics for query",
	},
}
