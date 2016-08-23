package paramsreplace

import (
	"strings"

	"github.com/arteev/dsql/cmd/parameters"
)

//Replace from source string by params. ${param} -> value
func Replace(source string, repParams []parameters.Parameter) (string, error) {
	if !strings.Contains(source, "{") || !strings.Contains(source, "}") {
		return source, nil
	}
	for _, p := range repParams {

		source = strings.Replace(source, "{"+p.Name+"}", p.Value, -1)
	}
	return source, nil
}
